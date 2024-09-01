package artisan

import (
	"fmt"
	"reflect"
)

type Column struct {
	Name string
	Type Type
}
type Columns map[string]Type
type Values map[string]any

// define abstract interfaces
type Type interface {
	Name() string
	Value() string
}

// define functions
func AllColumns() selectArg {
	return "*"
}

// define concrete SQLite types
type Base struct{ value string }
type Text struct{ Base }
type Integer struct{ Base }
type Real struct{ Base }
type Bool struct{ Base }

// define type name/value functions
func (b Base) Value() string { return b.value }
func (Base) Name() string    { return "TEXT" }
func (Text) Name() string    { return "TEXT" }
func (Integer) Name() string { return "INTEGER" }
func (Real) Name() string    { return "REAL" }
func (Bool) Name() string    { return "BOOLEAN" }

// define string interface functions for printing
func (t Text) String() string    { return t.Value() }
func (i Integer) String() string { return i.Value() }
func (r Real) String() string    { return r.Value() }

type IntegerType interface{ ~int | ~int32 | ~int64 }

func ConvertValue(value any) Type {
	s := fmt.Sprint(value)
	switch reflect.TypeOf(value).Kind() {
	case reflect.String:
		return Text{Base{fmt.Sprintf("\"%s\"", s)}}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return Integer{Base{s}}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return Integer{Base{s}}
	case reflect.Float32, reflect.Float64:
		return Real{Base{s}}
	case reflect.Bool:
		return Bool{Base{s}}
	}
	return Base{s}
}
