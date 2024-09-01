package artisan

import (
	"fmt"
	"strings"
)

const (
	TABLE createAttribute = iota
	TRIGGER
	INDEX
	VIEW
)

type createAttribute int
type createColumns struct {
	builder   *Builder
	tableName string
}
type create struct {
	builder *Builder
	columns []Column
}
type createOptions struct {
	attribute createAttribute
}
type CreateOption func(*createOptions)

func getCreateOptions(opts ...CreateOption) *createOptions {
	co := &createOptions{}
	for _, opt := range opts {
		opt(co)
	}
	return co
}

func (b *Builder) Create(tableName string) *createColumns {
	// always reset string if calling create
	b.Reset()
	b.stmt += fmt.Sprintf("CREATE TABLE %s", tableName)

	return &createColumns{builder: b, tableName: tableName}
}

func (c *createColumns) WithAttribute(attribute createAttribute) *createColumns {
	switch attribute {
	case TRIGGER:
		// b.stmt += fmt.Sprintf("CREATE TRIGGER %s")
		/* TODO */
	case VIEW:
		// b.stmt += fmt.Sprintf("CREATE VIEW %s")
		/* TODO */
	case INDEX:
		/* TODO */
	case TABLE:
		c.builder.stmt = fmt.Sprintf("CREATE TABLE %s", c.tableName)
	default:
		c.builder.stmt = fmt.Sprintf("CREATE TABLE %s", c.tableName)
		// TODO: do reflection and add to stmt based on *supported* data type
	}
	return c
}

func (c *createColumns) AddColumn(name string, typ Type) *create {
	cc := &create{builder: c.builder}
	return cc.AddColumn(name, typ)
}

func (c *create) AddColumn(name string, typ Type) *create {
	c.columns = append(c.columns, Column{name, typ})
	return c
}

func (c *createColumns) AddColumns(columns Columns) *create {
	cc := &create{builder: c.builder}
	return cc.AddColumns(columns)
}

func (c *create) AddColumns(columns Columns) *create {
	for k, v := range columns {
		c.columns = append(c.columns, Column{k, v})
	}
	return c
}

func (c *create) Build() string {
	var columns, values string
	for _, col := range c.columns {
		columns += fmt.Sprintf("%s %s, ", col.Name, col.Type.Name())
	}

	// trim off trailing delimiter
	columns = strings.TrimRight(columns, ", ")
	values = strings.TrimRight(values, ", ")

	return fmt.Sprintf("%s(%s);", c.builder.stmt, columns)
}

// func WithColumn(name string, _type Type) CreateOption {
// 	return func(co *createOptions) {
// 		co.columns = append(co.columns, Column{name: name, _type: _type})
// 	}
// }

// func WithColumns(columns Columns) CreateOption {
// 	return func(co *createOptions) {
// 		for name, _type := range columns {
// 			WithColumn(name, _type)
// 		}
// 	}
// }
