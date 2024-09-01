package artisan

import "fmt"

type drop struct {
	builder *Builder
}
type dropOptional struct {
	tableName string
	builder   *Builder
}

func (b *Builder) Drop(tableName string) *dropOptional {
	b.Reset()
	b.stmt += fmt.Sprintf("DROP TABLE %s", tableName)
	return &dropOptional{builder: b, tableName: tableName}
}

func (d *dropOptional) IfExists() *drop {
	d.builder.stmt = fmt.Sprintf("DROP TABLE IF EXISTS %s", d.tableName)
	return &drop{builder: d.builder}
}

func (d *dropOptional) Build() string {
	return d.builder.stmt + ";"
}

func (d *drop) Build() string {
	return d.builder.stmt + ";"
}
