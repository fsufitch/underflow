package schema

import "io"

// Parser is an interface for an object that can parse an Underflow store schema
type Parser interface {
	Parse(io.Reader)
}
