/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package connectionmanager

import (
	vcfg "k8s.io/cloud-provider-vsphere/pkg/common/config"
	cm "k8s.io/cloud-provider-vsphere/pkg/common/credentialmanager"
	vclib "k8s.io/cloud-provider-vsphere/pkg/common/vclib"
)

type ConnectionManager struct {
	// Maps the VC server to VSphereInstance
	VsphereInstanceMap map[string]*VSphereInstance
	// CredentialsManager
	credentialManager *cm.SecretCredentialManager
}

// VSphereInstance represents a vSphere instance where one or more kubernetes nodes are running.
type VSphereInstance struct {
	Conn *vclib.VSphereConnection
	Cfg  *vcfg.VirtualCenterConfig
}
