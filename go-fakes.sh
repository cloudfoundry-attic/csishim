#!/usr/bin/env bash

# csi_fake.go
go generate ./...

# fake_controllerclient.go
counterfeiter -o csi_fake/fake_controllerclient.go ../../github.com/container-storage-interface/spec/lib/go/csi/csi.pb.go ControllerClient

# fake_identityclient.go
counterfeiter -o csi_fake/fake_identityclient.go ../../github.com/container-storage-interface/spec/lib/go/csi/csi.pb.go IdentityClient

# fake_nodeclient.go
counterfeiter -o csi_fake/fake_nodeclient.go ../../github.com/container-storage-interface/spec/lib/go/csi/csi.pb.go NodeClient

