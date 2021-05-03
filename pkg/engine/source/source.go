// Package source (go:generate go run -mod=mod github.com/golang/mock/mockgen -package mock -source=./$GOFILE -destination=../mock/$GOFILE)
package source

import (
	"sort"

	"github.com/Checkmarx/kics/pkg/model"
)

const (
	// QueryFileName The default query file name
	QueryFileName = "query.rego"
	// MetadataFileName The default metadata file name
	MetadataFileName = "metadata.json"
	// LibraryFileName The default library file name
	LibraryFileName = "library.rego"
	// LibrariesDefaultBasePath the path to rego libraries
	LibrariesDefaultBasePath = "./assets/libraries/"
)

var (
	supportedPlatforms = map[string]string{
		"Ansible":        "ansible",
		"CloudFormation": "cloudformation",
		"Dockerfile":     "dockerfile",
		"Kubernetes":     "k8s",
		"Terraform":      "terraform",
		"OpenAPI":        "openapi",
	}
)

// ExcludeQueries represents a struct with options to exclude queries and a list for each option
type ExcludeQueries struct {
	ByIDs        []string
	ByCategories []string
}

// QueriesSource wraps an interface that contains basic methods: GetQueries and GetQueryLibrary
// GetQueries gets all queries from a QueryMetadata list
// GetQueryLibrary gets a library of rego functions given a plataform's name
type QueriesSource interface {
	GetQueries(excludeQueries ExcludeQueries) ([]model.QueryMetadata, error)
	GetQueryLibrary(platform string) (string, error)
}

// ListSupportedPlatforms returns a list of supported platforms
func ListSupportedPlatforms() []string {
	keys := make([]string, len(supportedPlatforms))
	i := 0
	for k := range supportedPlatforms {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	return keys
}
