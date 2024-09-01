package artisan

import (
	"fmt"
	"strings"
)

const (
	NONE selectAttribute = iota
	DISTINCT
	ALL
)

type selectAttribute int
type selectArg string
type selectClause struct{ builder *Builder }
type selectOptional struct{ builder *Builder }

// Select() initiates a new SQL using a builder object. Calling this function
// will call the builders Reset() function and clear the current statement.
func (b *Builder) Select(what ...selectArg) *selectOptional {
	b.Reset()
	if len(what) == 0 {
		b.stmt += "SELECT *"
	} else {
		var output string
		for _, s := range what {
			output += string(s) + ", "
		}
		output = strings.TrimRight(output, ", ")
		b.stmt += fmt.Sprintf("SELECT %s", output)
	}
	return &selectOptional{builder: b}
}

func (s *selectOptional) WithAttribute(attribute selectAttribute) *selectClause {
	switch attribute {
	case NONE:
		s.builder.stmt += ""
	case DISTINCT:
		s.builder.stmt += " DISTINCT"
	case ALL:
		s.builder.stmt += " ALL"
	default:
	}
	return &selectClause{builder: s.builder}
}

func (s *selectOptional) From(table tablename) *fromClause {
	sc := &selectClause{builder: s.builder}
	return sc.From(table)
}

func (s *selectClause) From(table tablename) *fromClause {
	s.builder.stmt += fmt.Sprintf(" FROM %s", table)
	return &fromClause{builder: s.builder}
}

func (s *fromClause) Build() string {
	return s.builder.stmt + ";"
}
