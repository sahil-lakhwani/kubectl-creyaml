package parse

import (
	"github.com/sahil-lakhwani/kubectl-creyaml/pkg/schema"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

func TypedCRD(u unstructured.Unstructured) (*apiextensions.CustomResourceDefinition, error) {
	var crd apiextensions.CustomResourceDefinition
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &crd)
	if err != nil {
		return nil, err
	}
	return &crd, nil
}

func GenerateCR(crd apiextensions.CustomResourceDefinition) *schema.CR {
	cr := &schema.CR{
		APIVersion: crd.Spec.Group + "/" + crd.Spec.Versions[0].Name,
		Kind:       crd.Spec.Names.Kind,
	}

	if crd.Spec.Versions[0].Schema != nil {
		m := make(map[string]interface{})
		cr.Spec = getSpecProperties(crd.Spec.Versions[0].Schema.OpenAPIV3Schema.Properties["spec"].Properties, m)
	}
	return cr
}

func getSpecProperties(props map[string]apiextensions.JSONSchemaProps, m map[string]interface{}) map[string]interface{} {
	for k, v := range props {
		if len(v.Properties) != 0 {
			m[k] = getSpecProperties(v.Properties, make(map[string]interface{}))
		} else {
			m[k] = v.Type
		}
	}

	return m
}
