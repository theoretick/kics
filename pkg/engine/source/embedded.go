package source

import (
	"embed"
	"fmt"
	"strings"

	"github.com/Checkmarx/kics/pkg/model"
)

//go:embed assets
var content embed.FS

func GetQueries() ([]model.QueryMetadata, error) {
	baseQueriesPath := "assets/queries"
	queryPlatforms, err := content.ReadDir(baseQueriesPath)
	if err != nil {
		return nil, err
	}
	for _, entry := range queryPlatforms {
		if entry.IsDir() {
			getPlatformQueries(baseQueriesPath, entry.Name())
		}
	}
	return nil, nil
}

func getPlatformQueries(base, platformDir string) {
	strings.Join([]string{base, platformDir}, "/")
}

func GetLibraries() {
	libraries, err := content.ReadDir("assets/libraries")
	if err != nil {
		panic(err)
	}
	for _, lib := range libraries {
		fmt.Println(lib)
	}
}
