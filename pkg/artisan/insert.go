package artisan

import (
	"fmt"
	"strings"
)

type insertClause struct {
	builder *Builder
	columns []Column
}
type insertColumns struct {
	builder   *Builder
	tableName string
}

func (b *Builder) Insert(tableName string) *insertColumns {
	b.Reset()
	b.stmt += fmt.Sprintf("INSERT INTO %s", tableName)

	return &insertColumns{builder: b}
}

func (i *insertColumns) AddValue(name string, value any) *insertClause {
	ic := &insertClause{builder: i.builder}
	return ic.AddValue(name, value)
}

func (i *insertClause) AddValue(name string, value any) *insertClause {
	i.columns = append(i.columns, Column{name, ConvertValue(value)})
	return i
}

func (i *insertColumns) AddValues(values Values) *insertClause {
	ic := &insertClause{builder: i.builder}
	return ic.AddValues(values)
}

func (i *insertClause) AddValues(values Values) *insertClause {
	for name, value := range values {
		i.columns = append(i.columns, Column{name, ConvertValue(value)})
	}
	return i
}

func (i *insertClause) Build() string {
	var columns, values string
	for _, col := range i.columns {
		columns += fmt.Sprintf("%s, ", col.Name)
		values += fmt.Sprintf("%s, ", col.Type.Value())
	}

	// trim off trailing delimiter
	columns = strings.TrimRight(columns, ", ")
	values = strings.TrimRight(values, ", ")
	return fmt.Sprintf("%s(%s) VALUES(%s);", i.builder.stmt, columns, values)
}
