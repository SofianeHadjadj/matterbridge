// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AttributeMappingSource undocumented
type AttributeMappingSource struct {
	// Object is the base model of AttributeMappingSource
	Object
	// Expression undocumented
	Expression *string `json:"expression,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Parameters undocumented
	Parameters []StringKeyAttributeMappingSourceValuePair `json:"parameters,omitempty"`
	// Type undocumented
	Type *AttributeMappingSourceType `json:"type,omitempty"`
}