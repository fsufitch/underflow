package schema

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"sort"
)

// FieldType is an enum type for values an Underflow store field can hold
type FieldType uint8

// Possible values for UFValue
const (
	UnknownField = FieldType(iota + 0)
	BoolField
	IntField
	FloatField
	StringField
)

func (t FieldType) String() string {
	switch t {
	case BoolField:
		return "bool"
	case IntField:
		return "int"
	case FloatField:
		return "float"
	case StringField:
		return "string"
	default:
		return "unknown"
	}
}

// FieldTypeFromString parses a field representation into a FieldType
func FieldTypeFromString(t string) (FieldType, error) {
	switch t {
	case "bool":
		return BoolField, nil
	case "int":
		return IntField, nil
	case "float":
		return FloatField, nil
	case "string":
		return StringField, nil
	default:
		return UnknownField, fmt.Errorf("unknown underflow field type: %s", t)
	}
}

// Schema describes the schema of an Underflow store
type Schema struct {
	ID       string
	Version  string
	Fields   map[string]FieldType
	hashFunc func([]byte) string
}

// DefaultSchemaHashFunc is the default hash function used for calculating schema checksum
// It uses crypto/md5 encoded with encoding/base64
var DefaultSchemaHashFunc = func(in []byte) string {
	binSum := md5.Sum(in)
	return base64.StdEncoding.EncodeToString(binSum[:])
}

// Checksum composes a string uniquely identifying this schema, for easy comparison
func (s Schema) Checksum() string {
	buf := bytes.NewBufferString("")
	buf.WriteString(fmt.Sprintf("underflow.schema::%s::%s\n", s.ID, s.Version))

	fieldIDs := []string{}
	for k := range s.Fields {
		fieldIDs = append(fieldIDs, k)
	}
	sort.Strings(fieldIDs)

	for _, fieldID := range fieldIDs {
		buf.WriteString(fmt.Sprintf("field::%s::%s\n", fieldID, s.Fields[fieldID].String()))
	}

	if s.hashFunc == nil {
		return DefaultSchemaHashFunc(buf.Bytes())
	}
	return s.hashFunc(buf.Bytes())
}
