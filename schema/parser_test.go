package schema

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

const goodYAMLInput = `
id: schema-123
version: 321
fields:
  field1: string
  field2: int
  field3: bool
`

func TestSerializedSchemaYAML_Schema(t *testing.T) {
	schema, err := YAMLParser{}.Parse(bytes.NewBufferString(goodYAMLInput))

	assert.Nil(t, err)
	assert.Equal(t, "schema-123", schema.ID)
	assert.Equal(t, "321", schema.Version)
	assert.Contains(t, schema.Fields, "field1")
	assert.Equal(t, StringField, schema.Fields["field1"])
	assert.Contains(t, schema.Fields, "field2")
	assert.Equal(t, IntField, schema.Fields["field2"])
	assert.Contains(t, schema.Fields, "field3")
	assert.Equal(t, BoolField, schema.Fields["field3"])
}

func TestSerializedSchemaYAML_Schema_Errors(t *testing.T) {
	expectError := func(expected string, input string) {
		_, err := YAMLParser{}.Parse(bytes.NewBufferString(input))
		assert.Error(t, err, expected)
		assert.Contains(t, err.Error(), expected)
	}

	expectError("failed decoding YAML input", `not a real yaml string`)
	expectError("no schema ID found", `{"version": "123", "fields": {"abc": "int"}}`)
	expectError("no schema version found", `{"id": "123", "fields": {"abc": "int"}}`)
	expectError("no schema fields", `{"id": "123", "version": "321"}`)
	expectError("no schema fields", `{"id": "123", "version": "321", "fields": {}}`)
	expectError("schema contains field of empty name", `{"id": "123", "version": "321", "fields": {"": "int"}}`)
	expectError("schema contains invalid field type", `{"id": "123", "version": "321", "fields": {"foo": "bar"}}`)
}

func TestProvideYAMLParser(t *testing.T) {
	// kind of a dumb test, since it's an empty struct for now
	assert.Equal(t, YAMLParser{}, ProvideYAMLParser())
}
