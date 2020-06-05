/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/sahil-lakhwani/kubectl-creyaml/pkg/kube"
	"github.com/sahil-lakhwani/kubectl-creyaml/pkg/parse"
	"github.com/sahil-lakhwani/kubectl-creyaml/pkg/yaml"
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kubectl-creyaml",
	Short: "Generate CR example YAML from CRD",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(generateCR(args[0]))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
}

func generateCR(name string) (string, error) {
	unstructured, err := kube.GetUnstructeredCRD(name)
	if err != nil {
		return "", err
	}

	crd, err := parse.TypedCRD(*unstructured)
	if err != nil {
		return "", err
	}

	cr := parse.GenerateCR(*crd)

	crYaml, err := yaml.CRYaml(*cr)
	if err != nil {
		return "", err
	}
	return crYaml, nil
}
