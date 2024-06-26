package cdk

import (
	"os"
	"path/filepath"
)

func detectLanguage() (lang string, err error) {
	root := "."
	patterns := map[string]string{
		"*.csproj":     "dotnet",
		"package.json": "typescript",
	}

	fileNames := []string{}
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fileNames = append(fileNames, path)

		return nil

	})
	if err != nil {
		return "", err
	}

	for _, fileName := range fileNames {
		for pattern := range patterns {
			matched, err := filepath.Match(pattern, filepath.Base(fileName))
			if err != nil {
				return "", err
			}
			if matched {
				return patterns[pattern], nil
			}
		}
	}
	return "", nil
}
