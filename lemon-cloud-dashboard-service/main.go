package main

import (
	"fmt"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/lemon-cloud-service/lemon-cloud-dashboard/lemon-cloud-dashboard-common/usr_service"
	"github.com/lemon-cloud-service/lemon-cloud-dashboard/lemon-cloud-dashboard-service/usr_service_impl"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

func main() {
	//listen, err := net.Listen("tcp", ":33385")
	//if err != nil {
	//	println("failed to listen: %v", err)
	//}

	// 实例化grpc Server
	s := grpc.NewServer()

	// 注册HelloService
	usr_service.RegisterAdminServiceServer(s, &usr_service_impl.AdminUsrServiceImpl{})
	//s.Serve(listen)

	grpcWebServer := grpcweb.WrapServer(s)
	httpServer := &http.Server{
		Handler: h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor == 2 {
				grpcWebServer.ServeHTTP(w, r)
			} else {
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
				w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-User-Agent, X-Grpc-Web")
				//w.Header().Set("grpc-message", "")
				//w.Header().Set("grpc-status", "")
				if grpcWebServer.IsGrpcWebRequest(r) {
					grpcWebServer.ServeHTTP(w, r)
				}
			}
		}), &http2.Server{}),
	}
	listen1, _ := net.Listen("tcp", ":33386")
	fmt.Println("服务启动完毕...")
	httpServer.Serve(listen1)
}
