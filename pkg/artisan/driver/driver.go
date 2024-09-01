package sql

type Driver interface {
	Build() string
}
