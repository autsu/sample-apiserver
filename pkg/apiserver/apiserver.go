/*
Copyright 2016 The Kubernetes Authors.

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

package apiserver

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/version"
	"k8s.io/apiserver/pkg/registry/rest"
	genericapiserver "k8s.io/apiserver/pkg/server"

	"k8s.io/sample-apiserver/pkg/apis/transformers"
	transformersinstall "k8s.io/sample-apiserver/pkg/apis/transformers/install"

	//"k8s.io/sample-apiserver/pkg/apis/wardle"
	//"k8s.io/sample-apiserver/pkg/apis/wardle/install"
	"k8s.io/sample-apiserver/pkg/registry"
	autobotstorage "k8s.io/sample-apiserver/pkg/registry/transformers/autobots"
	//fischerstorage "k8s.io/sample-apiserver/pkg/registry/wardle/fischer"
	//flunderstorage "k8s.io/sample-apiserver/pkg/registry/wardle/flunder"
	//decepticonstorage "k8s.io/sample-apiserver/pkg/registry/decepticons"
	//"k8s.io/sample-apiserver/pkg/registry/transformers"
	transform "k8s.io/sample-apiserver/pkg/registry/transformers/transform"
)

var (
	// Scheme defines methods for serializing and deserializing API objects.
	Scheme = runtime.NewScheme()
	// Codecs provides methods for retrieving codecs and serializers for specific
	// versions and content types.
	Codecs = serializer.NewCodecFactory(Scheme)
)

func init() {
	// install.Install(Scheme)
	transformersinstall.Install(Scheme)
	// we need to add the options to empty v1
	// TODO fix the server code to avoid this
	metav1.AddToGroupVersion(Scheme, schema.GroupVersion{Version: "v1"})

	// TODO: keep the generic API server from wanting this
	unversioned := schema.GroupVersion{Group: "", Version: "v1"}
	Scheme.AddUnversionedTypes(unversioned,
		&metav1.Status{},
		&metav1.APIVersions{},
		&metav1.APIGroupList{},
		&metav1.APIGroup{},
		&metav1.APIResourceList{},
	)
}

// ExtraConfig holds custom apiserver config
type ExtraConfig struct {
	// Place you custom config here.
}

// Config defines the config for the apiserver
type Config struct {
	GenericConfig *genericapiserver.RecommendedConfig
	ExtraConfig   ExtraConfig
}

// WardleServer contains state for a Kubernetes cluster master/api server.
type WardleServer struct {
	GenericAPIServer *genericapiserver.GenericAPIServer
}

type completedConfig struct {
	GenericConfig genericapiserver.CompletedConfig
	ExtraConfig   *ExtraConfig
}

// CompletedConfig embeds a private pointer that cannot be instantiated outside of this package.
type CompletedConfig struct {
	*completedConfig
}

// Complete fills in any fields not set that are required to have valid data. It's mutating the receiver.
func (cfg *Config) Complete() CompletedConfig {
	c := completedConfig{
		cfg.GenericConfig.Complete(),
		&cfg.ExtraConfig,
	}

	c.GenericConfig.Version = &version.Info{
		Major: "1",
		Minor: "0",
	}

	return CompletedConfig{&c}
}

// New returns a new instance of WardleServer from the given config.
func (c completedConfig) New() (*WardleServer, error) {
	genericServer, err := c.GenericConfig.New("sample-apiserver", genericapiserver.NewEmptyDelegate())
	if err != nil {
		return nil, err
	}

	s := &WardleServer{
		GenericAPIServer: genericServer,
	}

	apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(transformers.GroupName, Scheme, metav1.ParameterCodec, Codecs)

	v1alpha1storage := map[string]rest.Storage{}

	v1alpha1storage["autobots"] = registry.MustNewRestStorage(func() (rest.Storage, error) {
		//return autobotstorage.NewREST(Scheme, c.GenericConfig.RESTOptionsGetter)
		return &autobotstorage.FakeREST{}, nil
	})
	// 子资源
	v1alpha1storage["autobots/transform"] = registry.MustNewRestStorage(func() (rest.Storage, error) {
		//return autobotstorage.NewREST(Scheme, c.GenericConfig.RESTOptionsGetter)
		return &transform.REST{}, nil
	})
	// v1alpha1storage["autobots/transform"] = transformers.NewTransformerREST()
	// v1alpha1storage["autobots/status"] = autobotstatusREST
	// v1alpha1storage["autobots/scan"] = autobotstatsREST
	// v1alpha1storage["autobots/repair"] = autobotscanstatusREST
	// v1alpha1storage["autobots/battle"] = autobotbattleREST

	// v1alpha1storage["decepticons"] = decepticonstorage.NewREST(Scheme, c.GenericConfig.RESTOptionsGetter)

	apiGroupInfo.VersionedResourcesStorageMap["v1alpha1"] = v1alpha1storage

	v1beta1storage := map[string]rest.Storage{}
	apiGroupInfo.VersionedResourcesStorageMap["v1beta1"] = v1beta1storage

	if err := s.GenericAPIServer.InstallAPIGroup(&apiGroupInfo); err != nil {
		return nil, err
	}

	return s, nil
}
