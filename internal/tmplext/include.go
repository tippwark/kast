package tmplext

import (
	"os"
	"path/filepath"
	"strings"
)

func createIncludeFunc(runTemplate func(t string) (string, error), asYamlFields bool) func(pattern string) (string, error) {
	return func(pattern string) (string, error) {
		files, err := filepath.Glob(pattern)
		if err != nil {
			return "", err
		}

		var result strings.Builder
		for _, f := range files {
			fi, err := os.Stat(f)
			if err != nil {
				return "", err
			}

			if fi.IsDir() {
				continue
			}

			b, err := os.ReadFile(f)
			if err != nil {
				return "", err
			}

			res, err := runTemplate(string(b))
			if err != nil {
				return "", err
			}

			if asYamlFields {
				result.WriteString(fi.Name() + ": |\n")
				result.WriteString(indent(2, res) + "\n")
			} else {
				result.WriteString(res)
			}
		}

		return result.String(), nil
	}
}
