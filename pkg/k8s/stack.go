package k8s

import (
	"github.com/relingan/kubestream/stacks"
	apiErrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	k8s "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/restmapper"
)

// This helper function iterates through a map
// of stacks.Object, checks if they exists in
// the cluster, and creates them if they don't.
// The second parameter is the k8s dynamic client
// we want to use to make the API calls.
func DeployStack(
	objects map[int]stacks.Object,
	namespace string,
	dynClient dynamic.Interface,
	clientSet *k8s.Clientset) error {

	for i := 0; i < len(objects); i++ {
		obj := objects[i]

		gk := schema.GroupKind{
			Group: obj.GroupVersionKind.Group,
			Kind:  obj.GroupVersionKind.Kind,
		}

		// We need to use the ClientSet to discover
		// the available API types on the cluster.
		grs, err := restmapper.GetAPIGroupResources(clientSet.Discovery())
		if err != nil {
			return err
		}

		rm := restmapper.NewDiscoveryRESTMapper(grs)
		mapping, err := rm.RESTMapping(gk, obj.GroupVersionKind.Version)
		if err != nil {
			return err
		}

		// Convert the object to unstructured.Unstructured
		// so that we can pass it to the dynamic client.
		uns, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj.Object)
		if err != nil {
			return err
		}
		unstructured := unstructured.Unstructured{Object: uns}

		// If the unstructured object has a default
		// namespace attached that means the object
		// is namespaced.
		ns := unstructured.GetNamespace()

		// And if the namespace argument is not empty
		// we use it instead of the default.
		if ns != "" && namespace != "" {
			ns = namespace
			unstructured.SetNamespace(namespace)
		}

		// IMPORTANT: A really edgy case happens when we have
		// to create a namespace, which is not a namespaced object,
		// but that we need to rename if we want to substitute the
		// defaults. Given that all Kubestream resources are contained
		// in a single namespace for a given installation we can solve
		// it like this:
		if ns == "" && namespace != "" && unstructured.GetKind() == "Namespace" {
			unstructured.SetName(namespace)
		}

		// We don't care about the resulting
		// object as long as the operation
		// succeeds.
		if ns != "" {
			_, err = dynClient.
				Resource(mapping.Resource).
				Namespace(ns).
				Create(&unstructured, metav1.CreateOptions{})
		} else {
			_, err = dynClient.
				Resource(mapping.Resource).
				Create(&unstructured, metav1.CreateOptions{})
		}

		// If we get an AlreadyExists
		// error just ignore it.
		if err != nil && !apiErrors.IsAlreadyExists(err) {
			return err
		}
	}

	return nil
}
