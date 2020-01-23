package util

import (
	"path/filepath"

	k8s "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// GetClientSet Returns a k8s Clientset
// object that we can use to instantiate
// the clients for all the specific APIs.
func GetClientSet(kubeconfig string) (*k8s.Clientset, error) {
	home := homedir.HomeDir()
	if kubeconfig == "" && home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return &k8s.Clientset{}, err
	}

	clientset, err := k8s.NewForConfig(config)
	if err != nil {
		return &k8s.Clientset{}, err
	}

	return clientset, nil
}
