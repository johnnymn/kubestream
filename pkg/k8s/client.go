package k8s

import (
	"path/filepath"

	"k8s.io/client-go/dynamic"
	k8s "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// Returns a k8s dynamic client that we can use
// to make calls to the k8s API using a generics
// like workflow, instead of having to instantiate
// a typed client for each of the object kinds that
// we need to operate with.
func GetDynamicClient(kubeconfig string) (*dynamic.Interface, error) {
	home := homedir.HomeDir()
	if kubeconfig == "" && home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	client, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &client, nil
}

// Returns a k8s ClientSet that we can
// use to discover the API types that
// exists on the cluster.
func GetClientSet(kubeconfig string) (*k8s.Clientset, error) {
	home := homedir.HomeDir()
	if kubeconfig == "" && home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	clientSet, err := k8s.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientSet, nil
}
