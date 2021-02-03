package client

import (
	"flag"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var (
	clientSet  *kubernetes.Clientset
	kubeConfig *rest.Config
)

// NewKubeClient creates the kube client
func NewKubeClient() (*kubernetes.Clientset, error) {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	var config *rest.Config
	var err error
	if _, err = os.Stat(*kubeconfig); os.IsNotExist(err) || *kubeconfig == "" {
		config, err = rest.InClusterConfig()
	} else {
		config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
	}
	if err != nil {
		return nil, err
	}

	clientSet, err = kubernetes.NewForConfig(config)
	return clientSet, err
}

// GetKubeConfig gets the kube client
func GetKubeConfig() *kubernetes.Clientset {
	if clientSet == nil {
		clientSet, _ = NewKubeClient()
	}

	return clientSet
}
