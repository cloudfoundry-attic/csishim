// This file was generated by counterfeiter
// with command: counterfeiter -p /Users/pivotal/workspace/csi-local-volume-release/src/github.com//container-storage-interface/spec/lib/go/csi Csi
package csishim

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc"
)

type CsiShim struct{}

func (sh *CsiShim) NewIdentityClient(cc *grpc.ClientConn) csi.IdentityClient {
	return csi.NewIdentityClient(cc)
}

func (sh *CsiShim) RegisterIdentityServer(s *grpc.Server, srv csi.IdentityServer) {
	csi.RegisterIdentityServer(s, srv)
}

func (sh *CsiShim) NewControllerClient(cc *grpc.ClientConn) csi.ControllerClient {
	return csi.NewControllerClient(cc)
}

func (sh *CsiShim) RegisterControllerServer(s *grpc.Server, srv csi.ControllerServer) {
	csi.RegisterControllerServer(s, srv)
}

func (sh *CsiShim) NewNodeClient(cc *grpc.ClientConn) csi.NodeClient {
	return csi.NewNodeClient(cc)
}

func (sh *CsiShim) RegisterNodeServer(s *grpc.Server, srv csi.NodeServer) {
	csi.RegisterNodeServer(s, srv)
}