package tmplext

import (
	"bytes"

	"gopkg.in/yaml.v3"
)

func toYaml(obj any) string {
	dst := new(bytes.Buffer)
	encoder := yaml.NewEncoder(dst)
	encoder.SetIndent(2)
	encoder.Encode(obj)
	encoder.Close()

	return dst.String()
}
