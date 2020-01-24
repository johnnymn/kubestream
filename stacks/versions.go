package stacks

// Use int keys so that we can represent the order
// in which the manifests have to be applied.
var manifestFiles = map[string]map[int]string{
	"v1": map[int]string{
		0: "/stacks/templates/v1/namespace.yaml",
		1: "/stacks/templates/v1/registry-deployment.yaml",
		2: "/stacks/templates/v1/registry-service.yaml",
	},
}
