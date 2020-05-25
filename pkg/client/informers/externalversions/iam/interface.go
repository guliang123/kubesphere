/*
Copyright 2019 The KubeSphere Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package iam

import (
	v1alpha2 "kubesphere.io/kubesphere/pkg/client/informers/externalversions/iam/v1alpha2"
	internalinterfaces "kubesphere.io/kubesphere/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to each of this group's versions.
type Interface interface {
	// V1alpha2 provides access to shared informers for resources in V1alpha2.
	V1alpha2() v1alpha2.Interface
}

type group struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &group{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// V1alpha2 returns a new v1alpha2.Interface.
func (g *group) V1alpha2() v1alpha2.Interface {
	return v1alpha2.New(g.factory, g.namespace, g.tweakListOptions)
}
