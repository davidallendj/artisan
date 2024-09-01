package artisan

import "fmt"

func (s *fromClause) Where(condition conditional) *whereClause {
	s.builder.stmt += fmt.Sprintf(" WHERE %s", condition)
	return &whereClause{builder: s.builder}
}

func (s *whereClause) Build() string {
	return s.builder.stmt + ";"
}
