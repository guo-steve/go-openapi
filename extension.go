package openapi

// SpecificationExtensions provides support for OpenAPI extensions.
// It reads/writes all properties that begin with "x-".
type SpecificationExtensions map[string]interface{}

// Get returns required Extention object. If there is not given name,
// this function returns nil
func (x SpecificationExtensions) Get(name string) interface{} {
	val, ok := x[name]
	if !ok {
		return nil
	}
	return val
}
