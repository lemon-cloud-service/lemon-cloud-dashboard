package handler

import (
	"context"
	"fmt"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-utils/lccu_log"
	"github.com/lemon-cloud-service/lemon-cloud-dashboard/lemon-cloud-dashboard-common/adm_service"
	"github.com/lemon-cloud-service/lemon-cloud-dashboard/lemon-cloud-dashboard-service/grpc_adm_service_impl"
	"github.com/lemon-cloud-service/lemon-cloud-dashboard/lemon-cloud-dashboard-service/manager"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net"
	"net/http"
)

func StartGrpcServer() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", manager.ConfigManagerInstance().GeneralConfig().Service.Port))
	if err != nil {
		lccu_log.Error("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	//service.RegisterUserLoginServiceServer(s, &service_impl.LoginServiceImpl{})
	if err = s.Serve(listen); err != nil {
		lccu_log.Error("failed to start server: %v", err)
	}
}

func StartGrpcWebServer() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", manager.ConfigManagerInstance().GeneralConfig().Service.Port))
	if err != nil {
		lccu_log.Error("failed to listen: %v", err)
	}
	// 注册gRPC的拦截器
	var opts []grpc.ServerOption
	var interceptor grpc.UnaryServerInterceptor
	interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			fmt.Println("authorization: ", md.Get("authorization"))
		}
		return handler(ctx, req)
	}
	opts = append(opts, grpc.UnaryInterceptor(interceptor))
	s := grpc.NewServer(opts...)
	// 绑定所有服务实现
	registerAllGrpcServiceImpl(s)
	// 包装gRPC-web服务
	grpcWebServer := grpcweb.WrapServer(s)
	httpServer := &http.Server{
		Handler: h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor == 2 {
				grpcWebServer.ServeHTTP(w, r)
			} else {
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
				w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-User-Agent, X-Grpc-Web")
				if grpcWebServer.IsGrpcWebRequest(r) {
					grpcWebServer.ServeHTTP(w, r)
				}
			}
		}), &http2.Server{}),
	}
	if err = httpServer.Serve(listen); err != nil {
		lccu_log.Error("failed to start server: %v", err)
	}
}

func registerAllGrpcServiceImpl(server *grpc.Server) {
	adm_service.RegisterAdministratorServiceServer(server, &grpc_adm_service_impl.AdminUsrServiceImpl{})
	//adm_service.RegisterServiceServiceServer(server, &grpc_adm_service_impl.ServiceServiceImpl{})
}
