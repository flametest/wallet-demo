package grpc

import (
	"net"

	"github.com/flametest/vita/vserver"
	"github.com/flametest/wallet-demo/internal/container"
	"github.com/flametest/wallet-demo/internal/service"
	"github.com/flametest/wallet-demo/proto"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	listener net.Listener
	server   *grpc.Server
}

func NewGrpcServer(c container.Container) (*GrpcServer, error) {
	gRPCServer := vserver.NewGrpcServer()
	walletSvc := service.NewWalletService(c)
	walletServer := &WalletServer{
		UnimplementedWalletDemoServiceServer: proto.UnimplementedWalletDemoServiceServer{},
		WalletService:                        walletSvc,
	}
	proto.RegisterWalletDemoServiceServer(gRPCServer, walletServer)

	return &GrpcServer{
		server: gRPCServer,
	}, nil
}

func (s *GrpcServer) Start(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	s.listener = lis
	return s.server.Serve(lis)
}

func (s *GrpcServer) Stop() {
	if s.server != nil {
		s.server.GracefulStop()
	}
	if s.listener != nil {
		_ = s.listener.Close()
	}
}
