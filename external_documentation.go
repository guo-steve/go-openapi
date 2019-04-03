package openapi

// codebeat:disable[TOO_MANY_IVARS]

// ExternalDocumentation Object
type ExternalDocumentation struct {
	Description string
	URL         string
	Extensions  SpecificationExtensions `yaml:",inline"`
}

// Validate the values of ExternalDocumentaion object.
func (externalDocumentation ExternalDocumentation) Validate() error {
	return mustURL("externalDocumentation.url", externalDocumentation.URL)
}
