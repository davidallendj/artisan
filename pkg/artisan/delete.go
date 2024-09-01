package artisan

import "fmt"

type deleteClause struct{ builder *Builder }

func (b *Builder) Delete(tableName string) *deleteClause {
	b.stmt = fmt.Sprintf("DELETE FROM %s", tableName)
	return &deleteClause{builder: b}
}

func (s *deleteClause) Where(condition conditional) *whereClause {
	s.builder.stmt += fmt.Sprintf(" WHERE %s", condition)
	return &whereClause{builder: s.builder}
}

// defined in `select.go`
// func (s *whereClause) Build() string
