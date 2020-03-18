// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// UnsupportedDeviceConfigurationDetail undocumented
type UnsupportedDeviceConfigurationDetail struct {
	// Object is the base model of UnsupportedDeviceConfigurationDetail
	Object
	// Message A message explaining why an entity is unsupported.
	Message *string `json:"message,omitempty"`
	// PropertyName If message is related to a specific property in the original entity, then the name of that property.
	PropertyName *string `json:"propertyName,omitempty"`
}