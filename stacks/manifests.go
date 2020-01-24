package stacks

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// Represents a Kubernetes object that is rendered
// from a YAML filed bundled with Kubestream, and
// that is intended to be use with the go-client
// dynamic client. This simple helper is used to
// improve the readability of the metods that
// operate on the kubestream/stacks.
type Object struct {
	Object           runtime.Object
	GroupVersionKind *schema.GroupVersionKind
}

// Iterates over the map of YAML files for
// the stack, and returns a map of stack.Object
// that contains the rendered API object and the
// GroupVersionKind that it gets from the RenderManifest
// utility function. The returned map uses ints as keys
// so that caller code can order the manifests by apply order.
func GetManifests(pkg string) (map[int]Object, error) {
	manifests := make(map[int]Object)

	for i, f := range manifestFiles[pkg] {
		obj, version, err := RenderManifest(f)
		if err != nil {
			return nil, err
		}

		manifests[i] = Object{
			Object:           obj,
			GroupVersionKind: version,
		}
	}

	return manifests, nil
}
