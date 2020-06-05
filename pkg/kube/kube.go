package kube

import (
	"context"
	"os/user"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

func GetUnstructeredCRD(name string) (*unstructured.Unstructured, error) {
	kubepath, err := userConfig()
	if err != nil {
		return nil, err
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubepath)
	if err != nil {
		return nil, err
	}
	dyn, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	dr := dyn.Resource(schema.GroupVersionResource{
		Group:    "apiextensions.k8s.io",
		Version:  "v1",
		Resource: "customresourcedefinitions",
	})

	o, err := dr.Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return o, nil
}

func userConfig() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	return filepath.Join(usr.HomeDir, ".kube", "config"), nil
}
