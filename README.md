# gpay
An open source micro-service of China online payment, e.g. Alipay, Unionpay &amp; WeChatPay

## How to contribute

### Development


#### Protocol buffer

For example
```
[vagrant@kubedev-172-17-4-59 gpay]$ GOPATH=/Users/fanhongling/Downloads/workspace:/home/vagrant/go make protoc-grpc
```

#### Web

For example
```
[vagrant@kubedev-172-17-4-59 gpay]$ GOPATH=/Users/fanhongling/Downloads/workspace:/home/vagrant/go make go-bindata-web
```

### Vendoring

For example
```
[vagrant@kubedev-172-17-4-59 gpay]$ GOPATH=/Users/fanhongling/go GOBIN=/Users/fanhongling/Downloads/99-mirror/linux-bin go get -u github.com/golang/dep/cmd/dep
```

```
[vagrant@kubedev-172-17-4-59 gpay]$ dep version  
dep:
 version     : devel
 build date  : 
 git hash    : 
 go version  : go1.10
 go compiler : gc
 platform    : linux/amd64
 features    : ImportDuringSolve=false
```