package artisan

import (
	"fmt"
	"strings"
)

type updateClause struct{ builder *Builder }
type setClause struct{ builder *Builder }

func (b *Builder) Update(table tablename) *updateClause {
	b.stmt = fmt.Sprintf("UPDATE %s", table)
	return &updateClause{builder: b}
}

func (s *updateClause) Set(values Values) *setClause {
	columns := ""
	for name, value := range values {
		columns += fmt.Sprintf("%s = %v, ", name, value)
	}
	s.builder.stmt += strings.TrimRight(columns, ", ")
	return &setClause{builder: s.builder}
}

func (s *setClause) Where(condition conditional) *whereClause {
	s.builder.stmt += fmt.Sprintf(" WHERE %s", condition)
	return &whereClause{builder: s.builder}
}

func (s *whereClause) OrderBy(column string) *whereClause {
	s.builder.stmt += fmt.Sprintf(" ORDER BY %s", column)
	return s
}

func (s *whereClause) Limit(rowcount int) *whereClause {
	s.builder.stmt += fmt.Sprintf(" LIMIT %d", rowcount)
	return s
}

func (s *whereClause) Offset(offset int) *whereClause {
	s.builder.stmt += fmt.Sprintf(" OFFSET %d", offset)
	return s
}
