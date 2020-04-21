package kube // import "helm.sh/helm/v3/pkg/kube"

import (
	openshift "github.com/openshift/client-go/apps/clientset/versioned"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/tools/clientcmd"
)

type OpenShiftFactory interface {
	OpenShiftClientSet() (*openshift.Clientset, error)
}

func NewOpenShiftFactory(clientGetter genericclioptions.RESTClientGetter) OpenShiftFactory {
	if clientGetter == nil {
		panic("attempt to instantiate client_access_factory with nil clientGetter")
	}

	f := &openShiftFactoryImpl{
		clientGetter: clientGetter,
	}

	return f
}

type openShiftFactoryImpl struct {
	OpenShiftFactory
	clientGetter genericclioptions.RESTClientGetter
}

func (f *openShiftFactoryImpl) ToRawKubeConfigLoader() clientcmd.ClientConfig {
	return f.clientGetter.ToRawKubeConfigLoader()
}

func (f *openShiftFactoryImpl) OpenShiftClientSet() (*openshift.Clientset, error) {
	clientConfig, err := f.ToRawKubeConfigLoader().ClientConfig()
	if err != nil {
		return nil, err
	}
	cs, err := openshift.NewForConfig(clientConfig)
	if err != nil {
		return nil, err
	}
	return cs, nil
}
