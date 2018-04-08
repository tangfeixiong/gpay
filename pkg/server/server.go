package server

import (
	"fmt"
	//"mime"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"golang.org/x/net/context"

	"google.golang.org/grpc"

	"github.com/tangfeixiong/gpay/pb"
)

type Config struct {
	SecureAddress   string
	InsecureAddress string
	SecureHTTP      bool
	LogLevel        int
	Kubeconfig      string
	//RuntimeConfig   hadoop.Config
}

type Server struct {
	config *Config
	//ops    map[string]*operator.Operator
	root   http.FileSystem
	signal os.Signal
}

func Start(config *Config) {
	s := &Server{
		config: config,
		//ops:    make(map[string]*operator.Operator),
	}
	//	op, err := operator.Run(config.RuntimeConfig)
	//	if err != nil {
	//		glog.Errorf("Start operator failed: %v", err)
	//		return
	//	}
	//s.ops["hadoop-operator"] = op
	s.start()
}

func (s *Server) start() {
	ch := make(chan string)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.startGRPC(ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.startGateway(ch)
	}()

	/*
	   https://github.com/kubernetes/kubernetes/blob/release-1.1/build/pause/pause.go
	*/
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	// Block until a signal is received.
	s.signal = <-c

	//wg.Wait()
}

func (s *Server) startGRPC(ch chan<- string) {
	gs := grpc.NewServer()

	pb.RegisterSimpleGRpcServiceServer(gs, s)
	host := s.config.SecureAddress

	l, err := net.Listen("tcp", host)
	if err != nil {
		panic(err)
	}

	fmt.Println("Start gRPC into", l.Addr())
	go func() {
		time.Sleep(500)
		ch <- host
	}()
	if err := gs.Serve(l); nil != err {
		panic(err)
	}
}

func (s *Server) startGateway(ch <-chan string) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := http.NewServeMux()
	// mux.HandleFunc("/swagger/", serveSwagger2)
	//	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, req *http.Request) {
	//		io.Copy(w, strings.NewReader(healthcheckerpb.Swagger))
	//	})

	dopts := []grpc.DialOption{grpc.WithInsecure()}

	gwmux := runtime.NewServeMux()
	gRPCHost := <-ch

	host := s.config.InsecureAddress
	if err := pb.RegisterSimpleGRpcServiceHandlerFromEndpoint(ctx, gwmux, gRPCHost, dopts); err != nil {
		fmt.Println("Register gRPC gateway failed.", err)
		return
	}

	mux.Handle("/api", gwmux)
	// serveSwagger(mux)
	serveProm(mux)
	s.serveWebPages(mux)

	lstn, err := net.Listen("tcp", host)
	if nil != err {
		panic(err)
	}

	fmt.Println("Start gRPC Gateway into", lstn.Addr())
	//	if err := http.ListenAndServe(host, allowCORS(mux)); nil != err {
	//		fmt.Fprintf(os.Stderr, "Server died: %s\n", err)
	//	}
	gws := &http.Server{
		Handler: func /*allowCORS*/ (h http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if origin := r.Header.Get("Origin"); origin != "" {
					w.Header().Set("Access-Control-Allow-Origin", origin)
					if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
						func /*preflightHandler*/ (w http.ResponseWriter, r *http.Request) {
							headers := []string{"Content-Type", "Accept"}
							w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
							methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
							w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
							glog.Infof("preflight request for %s", r.URL.Path)
							return
						}(w, r)
						return
					}
				}
				h.ServeHTTP(w, r)
			})
		}(mux),
	}

	if err := gws.Serve(lstn); nil != err {
		fmt.Fprintln(os.Stderr, "Server died.", err.Error())
	}
}
