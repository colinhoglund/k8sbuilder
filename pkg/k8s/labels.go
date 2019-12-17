package k8s

// CommonLabelName - The name of the application
const CommonLabelName = "app.kubernetes.io/name"

// CommonLabelInstance - A unique name identifying the instance of an application
const CommonLabelInstance = "app.kubernetes.io/instance"

// CommonLabelVersion - The current version of the application (e.g., a semantic version, revision hash, etc.)
const CommonLabelVersion = "app.kubernetes.io/version"

// CommonLabelComponent - The component within the architecture
const CommonLabelComponent = "app.kubernetes.io/component"

// CommonLabelPartOf - The name of a higher level application this one is part of
const CommonLabelPartOf = "app.kubernetes.io/part-of"

// CommonLabelManagedBy - The tool being used to manage the operation of an application
const CommonLabelManagedBy = "app.kubernetes.io/managed-by"

// CommonLabels - https://kubernetes.io/docs/concepts/overview/working-with-objects/common-labels/
type CommonLabels struct {
	Name      string
	Instance  string
	Version   string
	Component string
	PartOf    string
	ManagedBy string
}

// Map returns CommonLabels as a map that can be used with other k8s libraries or marshaled to json
func (cl *CommonLabels) Map() map[string]string {
	labelMap := map[string]string{
		CommonLabelName:      cl.Name,
		CommonLabelInstance:  cl.Instance,
		CommonLabelVersion:   cl.Version,
		CommonLabelComponent: cl.Component,
		CommonLabelPartOf:    cl.PartOf,
		CommonLabelManagedBy: cl.ManagedBy,
	}

	// remove empty keys
	for k, v := range labelMap {
		if v == "" {
			delete(labelMap, k)
		}
	}

	return labelMap
}
