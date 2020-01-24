package stacks

import (
	"bufio"
	"io/ioutil"

	"github.com/markbates/pkger"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
)

// Returns a runtime.Object and API version that it
// renders from the file that is passed as argument.
// It loads the manifests from the pkger vault that
// is code generated and bundled in the binary.
func RenderManifest(file string) (runtime.Object, *schema.GroupVersionKind, error) {
	f, err := pkger.Open(file)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, nil, err
	}

	decode := scheme.Codecs.UniversalDeserializer().Decode

	obj, apiVersion, err := decode([]byte(data), nil, nil)
	if err != nil {
		return nil, nil, err
	}

	// By returning the schema.GroupVersionKind we can
	// enable caller code to directly pass obj to a
	// dynamic client
	return obj, apiVersion, nil
}
