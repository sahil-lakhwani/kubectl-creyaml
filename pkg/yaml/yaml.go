package yaml

import (

	"gopkg.in/yaml.v2"
	"github.com/pkg/errors"
	"github.com/sahil-lakhwani/kubectl-creyaml/pkg/schema"
)

func CRYaml(cr schema.CR) (string, error) {
	bytes, err := yaml.Marshal(&cr)
	if err != nil {
		return "", errors.Wrap(err, "Error converting to YAML")
	}
	return string(bytes), nil
}
