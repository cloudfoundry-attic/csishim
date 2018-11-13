// This file was generated by counterfeiter
// with command: counterfeiter -p /Users/pivotal/workspace/csi-local-volume-release/src/github.com//container-storage-interface/spec/lib/go/csi Csi
package csishim

import (
	"code.cloudfoundry.org/goshims/grpcshim"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc"
)

//go:generate counterfeiter -o csi_fake/fake_csi.go . Csi
type Csi interface {
	NewIdentityClient(cc grpcshim.ClientConn) csi.IdentityClient
	RegisterIdentityServer(s *grpc.Server, srv csi.IdentityServer)
	NewControllerClient(cc grpcshim.ClientConn) csi.ControllerClient
	RegisterControllerServer(s *grpc.Server, srv csi.ControllerServer)
	NewNodeClient(cc grpcshim.ClientConn) csi.NodeClient
	RegisterNodeServer(s *grpc.Server, srv csi.NodeServer)
}
