package schema

import (
	"io"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// Parser is an interface for an object that can parse an Underflow store schema
type Parser interface {
	Parse(io.Reader) (*Schema, error)
}

// SerializedSchemaYAML is what it says on the tin
type SerializedSchemaYAML struct {
	ID      string            `yaml:"id"`
	Version string            `yaml:"version"`
	Fields  map[string]string `yaml:"fields"`
}

// Schema converts this into a Schema
func (s SerializedSchemaYAML) Schema() (*Schema, error) {
	schema := Schema{ID: s.ID, Version: s.Version, Fields: map[string]FieldType{}}
	if s.ID == "" {
		return nil, errors.New("no schema ID found")
	}
	if s.Version == "" {
		return nil, errors.New("no schema version found")
	}
	if len(s.Fields) == 0 {
		return nil, errors.New("no schema fields found")
	}
	for k, v := range s.Fields {
		if k == "" {
			return nil, errors.New("schema contains field of empty name")
		}
		fieldType, err := FieldTypeFromString(v)
		if err != nil {
			return nil, errors.Wrap(err, "schema contains invalid field type")
		}
		schema.Fields[k] = fieldType
	}

	return &schema, nil
}

// YAMLParser is an implementation of Parser based on YAML
type YAMLParser struct{}

// Parse parses the YAML content of an io.Reader into a Schema
func (p YAMLParser) Parse(in io.Reader) (*Schema, error) {
	serialSchema := SerializedSchemaYAML{}
	decoder := yaml.NewDecoder(in)
	decoder.SetStrict(true)
	err := decoder.Decode(&serialSchema)
	if err != nil {
		return nil, errors.Wrap(err, "failed decoding YAML input")
	}
	schema, err := serialSchema.Schema()
	if err != nil {
		return nil, errors.Wrap(err, "invalid YAML input")
	}
	return schema, nil
}

// ProvideYAMLParser creates a YAMLParser for injection
func ProvideYAMLParser() YAMLParser {
	return YAMLParser{}
}
