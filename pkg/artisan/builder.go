package artisan

import (
	"fmt"

	sql "github.com/davidallendj/artisan/pkg/artisan/driver"
)

type state interface {
	GetBuilder() *Builder
}

type Builder struct {
	Driver sql.Driver
	stmt   string
	err    error
}

type fromClause struct{ builder *Builder }
type whereClause struct{ builder *Builder }

type tablename string
type conditional string

func IsGreaterThan[T any, U any](column T, value U) conditional {
	return conditional(fmt.Sprintf("%v>%v", column, ConvertValue(value)))
}
func IsGreaterThanOrEqual[T any, U any](column T, value U) conditional {
	return conditional(fmt.Sprintf("%v>=%v", column, ConvertValue(value)))
}
func IsLessThan[T any, U any](column T, value U) conditional {
	return conditional(fmt.Sprintf("%v<%v", column, ConvertValue(value)))
}
func IsLessThanOrEqual[T any, U any](column T, value U) conditional {
	return conditional(fmt.Sprintf("%v<=%v", column, ConvertValue(value)))
}
func IsEqual[T any, U any](column T, value U) conditional {
	return conditional(fmt.Sprintf("%v=%v", column, ConvertValue(value)))
}

func (b *Builder) Build() string {
	var s []byte = make([]byte, len(b.stmt))
	copy(s, b.stmt)
	b.Reset()
	return string(s)
}

func (b *Builder) Error() error {
	return b.err
}

func (b *Builder) Reset() {
	b.stmt = ""
}
