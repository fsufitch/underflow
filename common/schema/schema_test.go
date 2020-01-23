package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldType_String(t *testing.T) {
	assert.Equal(t, "bool", BoolField.String())
	assert.Equal(t, "int", IntField.String())
	assert.Equal(t, "float", FloatField.String())
	assert.Equal(t, "string", StringField.String())
	assert.Equal(t, "unknown", UnknownField.String())
	assert.Equal(t, "unknown", FieldType(123).String())
}

func TestFieldTypeFromString(t *testing.T) {
	var field FieldType
	var err error

	field, err = FieldTypeFromString("bool")
	assert.Nil(t, err)
	assert.Equal(t, BoolField, field)

	field, err = FieldTypeFromString("int")
	assert.Nil(t, err)
	assert.Equal(t, IntField, field)

	field, err = FieldTypeFromString("float")
	assert.Nil(t, err)
	assert.Equal(t, FloatField, field)

	field, err = FieldTypeFromString("string")
	assert.Nil(t, err)
	assert.Equal(t, StringField, field)

	field, err = FieldTypeFromString("unknown")
	assert.NotNil(t, err)
	assert.Equal(t, UnknownField, field)

	field, err = FieldTypeFromString("STRING")
	assert.NotNil(t, err)
	assert.Equal(t, UnknownField, field)
}

func TestSchema_Checksum_Equivalence(t *testing.T) {
	sumBaseline := Schema{ID: "test", Version: "1", Fields: map[string]FieldType{
		"field1": IntField,
		"field2": StringField,
		"field3": BoolField,
	}}.Checksum()
	sumIdentical := Schema{ID: "test", Version: "1", Fields: map[string]FieldType{
		"field1": IntField,
		"field2": StringField,
		"field3": BoolField,
	}}.Checksum()
	sumFieldsReordered := Schema{ID: "test", Version: "1", Fields: map[string]FieldType{
		"field3": BoolField, // Moved field
		"field1": IntField,
		"field2": StringField,
	}}.Checksum()
	sumFieldRemoved := Schema{ID: "test", Version: "1", Fields: map[string]FieldType{
		"field1": IntField,
		"field2": StringField,
		// "field3": BoolField, // Removed field
	}}.Checksum()
	sumFieldAdded := Schema{ID: "test", Version: "1", Fields: map[string]FieldType{
		"field1": IntField,
		"field2": StringField,
		"field3": BoolField,
		"field4": IntField, // New field
	}}.Checksum()
	sumFieldTypeChanged := Schema{ID: "test", Version: "1", Fields: map[string]FieldType{
		"field1": IntField,
		"field2": StringField,
		"field3": FloatField, // Changed field
	}}.Checksum()
	sumIDChanged := Schema{ID: "different", Version: "1", Fields: map[string]FieldType{
		"field1": IntField,
		"field2": StringField,
		"field3": BoolField,
	}}.Checksum()
	sumVersionChanged := Schema{ID: "test", Version: "2", Fields: map[string]FieldType{
		"field1": IntField,
		"field2": StringField,
		"field3": BoolField,
	}}.Checksum()

	assert.Equal(t, sumBaseline, sumIdentical)
	assert.Equal(t, sumBaseline, sumFieldsReordered)
	assert.NotEqual(t, sumBaseline, sumFieldRemoved)
	assert.NotEqual(t, sumBaseline, sumFieldAdded)
	assert.NotEqual(t, sumBaseline, sumFieldTypeChanged)
	assert.NotEqual(t, sumBaseline, sumIDChanged)
	assert.NotEqual(t, sumBaseline, sumVersionChanged)
}

func TestSchema_Checksum_Format(t *testing.T) {
	// Setup
	calledWithValue := []byte{}
	customHash := func(in []byte) string {
		calledWithValue = in
		return "test return"
	}
	sum := Schema{ID: "test", Version: "1", hashFunc: customHash, Fields: map[string]FieldType{
		"field1": IntField,
		"field2": StringField,
		"field3": BoolField,
	}}

	// Tested code
	result := sum.Checksum()

	// Asserts
	assert.Equal(t, "test return", result)
	assert.Equal(t, `underflow.schema::test::1
field::field1::int
field::field2::string
field::field3::bool
`, string(calledWithValue))
}
