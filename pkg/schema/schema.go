package schema

type CR struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name        string `yaml:"name"`
		Annotations struct {
		} `yaml:"annotations"`
	} `yaml:"metadata"`
	Spec map[string]interface{} `yaml:"spec"`
}
