/*
  Inspired by
    https://github.com/prometheus/client_golang/blob/master/examples/simple/main.go
*/

package /* main */ server

import (
	//"flag"
	//"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//var addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")

func /* main */ serveProm(mux *http.ServeMux) {
	//flag.Parse()
	//http.Handle("/metrics", promhttp.Handler())
	//log.Fatal(http.ListenAndServe(*addr, nil))
	mux.Handle("/metrics", promhttp.Handler())
}
