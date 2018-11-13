// Code generated by counterfeiter. DO NOT EDIT.
package csi_fake

import (
	sync "sync"

	csishim "code.cloudfoundry.org/csishim"
	grpcshim "code.cloudfoundry.org/goshims/grpcshim"
	csi "github.com/container-storage-interface/spec/lib/go/csi"
	grpc "google.golang.org/grpc"
)

type FakeCsi struct {
	NewControllerClientStub        func(grpcshim.ClientConn) csi.ControllerClient
	newControllerClientMutex       sync.RWMutex
	newControllerClientArgsForCall []struct {
		arg1 grpcshim.ClientConn
	}
	newControllerClientReturns struct {
		result1 csi.ControllerClient
	}
	newControllerClientReturnsOnCall map[int]struct {
		result1 csi.ControllerClient
	}
	NewIdentityClientStub        func(grpcshim.ClientConn) csi.IdentityClient
	newIdentityClientMutex       sync.RWMutex
	newIdentityClientArgsForCall []struct {
		arg1 grpcshim.ClientConn
	}
	newIdentityClientReturns struct {
		result1 csi.IdentityClient
	}
	newIdentityClientReturnsOnCall map[int]struct {
		result1 csi.IdentityClient
	}
	NewNodeClientStub        func(grpcshim.ClientConn) csi.NodeClient
	newNodeClientMutex       sync.RWMutex
	newNodeClientArgsForCall []struct {
		arg1 grpcshim.ClientConn
	}
	newNodeClientReturns struct {
		result1 csi.NodeClient
	}
	newNodeClientReturnsOnCall map[int]struct {
		result1 csi.NodeClient
	}
	RegisterControllerServerStub        func(*grpc.Server, csi.ControllerServer)
	registerControllerServerMutex       sync.RWMutex
	registerControllerServerArgsForCall []struct {
		arg1 *grpc.Server
		arg2 csi.ControllerServer
	}
	RegisterIdentityServerStub        func(*grpc.Server, csi.IdentityServer)
	registerIdentityServerMutex       sync.RWMutex
	registerIdentityServerArgsForCall []struct {
		arg1 *grpc.Server
		arg2 csi.IdentityServer
	}
	RegisterNodeServerStub        func(*grpc.Server, csi.NodeServer)
	registerNodeServerMutex       sync.RWMutex
	registerNodeServerArgsForCall []struct {
		arg1 *grpc.Server
		arg2 csi.NodeServer
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCsi) NewControllerClient(arg1 grpcshim.ClientConn) csi.ControllerClient {
	fake.newControllerClientMutex.Lock()
	ret, specificReturn := fake.newControllerClientReturnsOnCall[len(fake.newControllerClientArgsForCall)]
	fake.newControllerClientArgsForCall = append(fake.newControllerClientArgsForCall, struct {
		arg1 grpcshim.ClientConn
	}{arg1})
	fake.recordInvocation("NewControllerClient", []interface{}{arg1})
	fake.newControllerClientMutex.Unlock()
	if fake.NewControllerClientStub != nil {
		return fake.NewControllerClientStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.newControllerClientReturns
	return fakeReturns.result1
}

func (fake *FakeCsi) NewControllerClientCallCount() int {
	fake.newControllerClientMutex.RLock()
	defer fake.newControllerClientMutex.RUnlock()
	return len(fake.newControllerClientArgsForCall)
}

func (fake *FakeCsi) NewControllerClientCalls(stub func(grpcshim.ClientConn) csi.ControllerClient) {
	fake.newControllerClientMutex.Lock()
	defer fake.newControllerClientMutex.Unlock()
	fake.NewControllerClientStub = stub
}

func (fake *FakeCsi) NewControllerClientArgsForCall(i int) grpcshim.ClientConn {
	fake.newControllerClientMutex.RLock()
	defer fake.newControllerClientMutex.RUnlock()
	argsForCall := fake.newControllerClientArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCsi) NewControllerClientReturns(result1 csi.ControllerClient) {
	fake.newControllerClientMutex.Lock()
	defer fake.newControllerClientMutex.Unlock()
	fake.NewControllerClientStub = nil
	fake.newControllerClientReturns = struct {
		result1 csi.ControllerClient
	}{result1}
}

func (fake *FakeCsi) NewControllerClientReturnsOnCall(i int, result1 csi.ControllerClient) {
	fake.newControllerClientMutex.Lock()
	defer fake.newControllerClientMutex.Unlock()
	fake.NewControllerClientStub = nil
	if fake.newControllerClientReturnsOnCall == nil {
		fake.newControllerClientReturnsOnCall = make(map[int]struct {
			result1 csi.ControllerClient
		})
	}
	fake.newControllerClientReturnsOnCall[i] = struct {
		result1 csi.ControllerClient
	}{result1}
}

func (fake *FakeCsi) NewIdentityClient(arg1 grpcshim.ClientConn) csi.IdentityClient {
	fake.newIdentityClientMutex.Lock()
	ret, specificReturn := fake.newIdentityClientReturnsOnCall[len(fake.newIdentityClientArgsForCall)]
	fake.newIdentityClientArgsForCall = append(fake.newIdentityClientArgsForCall, struct {
		arg1 grpcshim.ClientConn
	}{arg1})
	fake.recordInvocation("NewIdentityClient", []interface{}{arg1})
	fake.newIdentityClientMutex.Unlock()
	if fake.NewIdentityClientStub != nil {
		return fake.NewIdentityClientStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.newIdentityClientReturns
	return fakeReturns.result1
}

func (fake *FakeCsi) NewIdentityClientCallCount() int {
	fake.newIdentityClientMutex.RLock()
	defer fake.newIdentityClientMutex.RUnlock()
	return len(fake.newIdentityClientArgsForCall)
}

func (fake *FakeCsi) NewIdentityClientCalls(stub func(grpcshim.ClientConn) csi.IdentityClient) {
	fake.newIdentityClientMutex.Lock()
	defer fake.newIdentityClientMutex.Unlock()
	fake.NewIdentityClientStub = stub
}

func (fake *FakeCsi) NewIdentityClientArgsForCall(i int) grpcshim.ClientConn {
	fake.newIdentityClientMutex.RLock()
	defer fake.newIdentityClientMutex.RUnlock()
	argsForCall := fake.newIdentityClientArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCsi) NewIdentityClientReturns(result1 csi.IdentityClient) {
	fake.newIdentityClientMutex.Lock()
	defer fake.newIdentityClientMutex.Unlock()
	fake.NewIdentityClientStub = nil
	fake.newIdentityClientReturns = struct {
		result1 csi.IdentityClient
	}{result1}
}

func (fake *FakeCsi) NewIdentityClientReturnsOnCall(i int, result1 csi.IdentityClient) {
	fake.newIdentityClientMutex.Lock()
	defer fake.newIdentityClientMutex.Unlock()
	fake.NewIdentityClientStub = nil
	if fake.newIdentityClientReturnsOnCall == nil {
		fake.newIdentityClientReturnsOnCall = make(map[int]struct {
			result1 csi.IdentityClient
		})
	}
	fake.newIdentityClientReturnsOnCall[i] = struct {
		result1 csi.IdentityClient
	}{result1}
}

func (fake *FakeCsi) NewNodeClient(arg1 grpcshim.ClientConn) csi.NodeClient {
	fake.newNodeClientMutex.Lock()
	ret, specificReturn := fake.newNodeClientReturnsOnCall[len(fake.newNodeClientArgsForCall)]
	fake.newNodeClientArgsForCall = append(fake.newNodeClientArgsForCall, struct {
		arg1 grpcshim.ClientConn
	}{arg1})
	fake.recordInvocation("NewNodeClient", []interface{}{arg1})
	fake.newNodeClientMutex.Unlock()
	if fake.NewNodeClientStub != nil {
		return fake.NewNodeClientStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.newNodeClientReturns
	return fakeReturns.result1
}

func (fake *FakeCsi) NewNodeClientCallCount() int {
	fake.newNodeClientMutex.RLock()
	defer fake.newNodeClientMutex.RUnlock()
	return len(fake.newNodeClientArgsForCall)
}

func (fake *FakeCsi) NewNodeClientCalls(stub func(grpcshim.ClientConn) csi.NodeClient) {
	fake.newNodeClientMutex.Lock()
	defer fake.newNodeClientMutex.Unlock()
	fake.NewNodeClientStub = stub
}

func (fake *FakeCsi) NewNodeClientArgsForCall(i int) grpcshim.ClientConn {
	fake.newNodeClientMutex.RLock()
	defer fake.newNodeClientMutex.RUnlock()
	argsForCall := fake.newNodeClientArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCsi) NewNodeClientReturns(result1 csi.NodeClient) {
	fake.newNodeClientMutex.Lock()
	defer fake.newNodeClientMutex.Unlock()
	fake.NewNodeClientStub = nil
	fake.newNodeClientReturns = struct {
		result1 csi.NodeClient
	}{result1}
}

func (fake *FakeCsi) NewNodeClientReturnsOnCall(i int, result1 csi.NodeClient) {
	fake.newNodeClientMutex.Lock()
	defer fake.newNodeClientMutex.Unlock()
	fake.NewNodeClientStub = nil
	if fake.newNodeClientReturnsOnCall == nil {
		fake.newNodeClientReturnsOnCall = make(map[int]struct {
			result1 csi.NodeClient
		})
	}
	fake.newNodeClientReturnsOnCall[i] = struct {
		result1 csi.NodeClient
	}{result1}
}

func (fake *FakeCsi) RegisterControllerServer(arg1 *grpc.Server, arg2 csi.ControllerServer) {
	fake.registerControllerServerMutex.Lock()
	fake.registerControllerServerArgsForCall = append(fake.registerControllerServerArgsForCall, struct {
		arg1 *grpc.Server
		arg2 csi.ControllerServer
	}{arg1, arg2})
	fake.recordInvocation("RegisterControllerServer", []interface{}{arg1, arg2})
	fake.registerControllerServerMutex.Unlock()
	if fake.RegisterControllerServerStub != nil {
		fake.RegisterControllerServerStub(arg1, arg2)
	}
}

func (fake *FakeCsi) RegisterControllerServerCallCount() int {
	fake.registerControllerServerMutex.RLock()
	defer fake.registerControllerServerMutex.RUnlock()
	return len(fake.registerControllerServerArgsForCall)
}

func (fake *FakeCsi) RegisterControllerServerCalls(stub func(*grpc.Server, csi.ControllerServer)) {
	fake.registerControllerServerMutex.Lock()
	defer fake.registerControllerServerMutex.Unlock()
	fake.RegisterControllerServerStub = stub
}

func (fake *FakeCsi) RegisterControllerServerArgsForCall(i int) (*grpc.Server, csi.ControllerServer) {
	fake.registerControllerServerMutex.RLock()
	defer fake.registerControllerServerMutex.RUnlock()
	argsForCall := fake.registerControllerServerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCsi) RegisterIdentityServer(arg1 *grpc.Server, arg2 csi.IdentityServer) {
	fake.registerIdentityServerMutex.Lock()
	fake.registerIdentityServerArgsForCall = append(fake.registerIdentityServerArgsForCall, struct {
		arg1 *grpc.Server
		arg2 csi.IdentityServer
	}{arg1, arg2})
	fake.recordInvocation("RegisterIdentityServer", []interface{}{arg1, arg2})
	fake.registerIdentityServerMutex.Unlock()
	if fake.RegisterIdentityServerStub != nil {
		fake.RegisterIdentityServerStub(arg1, arg2)
	}
}

func (fake *FakeCsi) RegisterIdentityServerCallCount() int {
	fake.registerIdentityServerMutex.RLock()
	defer fake.registerIdentityServerMutex.RUnlock()
	return len(fake.registerIdentityServerArgsForCall)
}

func (fake *FakeCsi) RegisterIdentityServerCalls(stub func(*grpc.Server, csi.IdentityServer)) {
	fake.registerIdentityServerMutex.Lock()
	defer fake.registerIdentityServerMutex.Unlock()
	fake.RegisterIdentityServerStub = stub
}

func (fake *FakeCsi) RegisterIdentityServerArgsForCall(i int) (*grpc.Server, csi.IdentityServer) {
	fake.registerIdentityServerMutex.RLock()
	defer fake.registerIdentityServerMutex.RUnlock()
	argsForCall := fake.registerIdentityServerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCsi) RegisterNodeServer(arg1 *grpc.Server, arg2 csi.NodeServer) {
	fake.registerNodeServerMutex.Lock()
	fake.registerNodeServerArgsForCall = append(fake.registerNodeServerArgsForCall, struct {
		arg1 *grpc.Server
		arg2 csi.NodeServer
	}{arg1, arg2})
	fake.recordInvocation("RegisterNodeServer", []interface{}{arg1, arg2})
	fake.registerNodeServerMutex.Unlock()
	if fake.RegisterNodeServerStub != nil {
		fake.RegisterNodeServerStub(arg1, arg2)
	}
}

func (fake *FakeCsi) RegisterNodeServerCallCount() int {
	fake.registerNodeServerMutex.RLock()
	defer fake.registerNodeServerMutex.RUnlock()
	return len(fake.registerNodeServerArgsForCall)
}

func (fake *FakeCsi) RegisterNodeServerCalls(stub func(*grpc.Server, csi.NodeServer)) {
	fake.registerNodeServerMutex.Lock()
	defer fake.registerNodeServerMutex.Unlock()
	fake.RegisterNodeServerStub = stub
}

func (fake *FakeCsi) RegisterNodeServerArgsForCall(i int) (*grpc.Server, csi.NodeServer) {
	fake.registerNodeServerMutex.RLock()
	defer fake.registerNodeServerMutex.RUnlock()
	argsForCall := fake.registerNodeServerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCsi) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newControllerClientMutex.RLock()
	defer fake.newControllerClientMutex.RUnlock()
	fake.newIdentityClientMutex.RLock()
	defer fake.newIdentityClientMutex.RUnlock()
	fake.newNodeClientMutex.RLock()
	defer fake.newNodeClientMutex.RUnlock()
	fake.registerControllerServerMutex.RLock()
	defer fake.registerControllerServerMutex.RUnlock()
	fake.registerIdentityServerMutex.RLock()
	defer fake.registerIdentityServerMutex.RUnlock()
	fake.registerNodeServerMutex.RLock()
	defer fake.registerNodeServerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCsi) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ csishim.Csi = new(FakeCsi)
