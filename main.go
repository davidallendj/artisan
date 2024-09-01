package main

import (
	"fmt"

	artisan "github.com/davidallendj/artisan/pkg/artisan"
)

func main() {
	// the builder is always the starting point with artisan
	var (
		b     = artisan.Builder{}
		name  = "david"
		count = 1
	)

	// do 'create' with two different ways to add columns
	fmt.Printf("create.1: %s\n", b.Create("test").WithAttribute(artisan.TABLE).AddColumns(artisan.Columns{
		"count": artisan.Integer{},
		"name":  artisan.Text{},
	}).Build())
	fmt.Printf("create.2: %s\n", b.Create("test").AddColumn("count", artisan.Integer{}).AddColumn("name", artisan.Text{}).Build())
	fmt.Printf("create.3: %s\n", b.Create("test").AddColumns(artisan.Columns{"count": artisan.Integer{}, "name": artisan.Text{}}).Build())

	// do 'insert' with two different ways to add values
	fmt.Printf("insert.1: %s\n", b.Insert("test").AddValue("count", count).Build())
	fmt.Printf("insert.2: %s\n", b.Insert("test").AddValues(artisan.Values{"count": count, "name": name}).Build())

	// do 'select' both with and without 'where' clause
	fmt.Printf("select.1: %s\n", b.Select().From("test").Where(artisan.IsEqual("name", name)).Build())
	fmt.Printf("select.2: %s\n", b.Select(artisan.AllColumns()).From("test").Build())
	fmt.Printf("select.3: %s\n", b.Select("count", "name").From("test").Where("count>10").Build())
	fmt.Printf("select.4: %s\n", b.Select("name").WithAttribute(artisan.DISTINCT).From("test").Where(artisan.IsGreaterThanOrEqual("count", 10)).Build())

	// do 'update' to set existing values
	fmt.Printf("update.1: %s\n", b.Update("test").Set(artisan.Values{"count": 10, "name": "joe"}).Where("count>3").Build())
	fmt.Printf("update.2: %s\n", b.Update("test").Set(artisan.Values{"count": 10, "name": "joe"}).Where(artisan.IsGreaterThan("count", 3)).OrderBy("count").Limit(10).Offset(2).Build())

	// do 'delete' and delete a single and multiple records
	fmt.Printf("delete.1: %s\n", b.Delete("test").Where(artisan.IsLessThanOrEqual("count", 1)).Build())

	// do 'drop' to remove a table
	fmt.Printf("drop.1: %s\n", b.Drop("test").Build())
	fmt.Printf("drop.2: %s\n", b.Drop("test").IfExists().Build())
}
