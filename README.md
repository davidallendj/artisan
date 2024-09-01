# Artisan

The `artisan` package is an experimental, LINQ-inspired SQL string query builder that is meant to be simple and intuitive to create SQL queries. The library leans into IDE auto-suggestions and makes full use of Go's type system to reduce mistakes and increase productivity by catching common SQL errors that are hard to debug when using plain strings.

## Building the Executable

Clone the repository, download the dependencies, and build the binary.

```bash
git clone github.com/davidallendj/artisan
cd artisan
go mod tidy && go build
```

## Running the Example

For now, the `artisan` only has some example code. Run it with `./artisan`.

## Strict Functional API

The `artisan` tool uses an internal state-machine to create a guided experience through a strict API. Restricting the API in this manner helps to prevent small mistakes early and eliminates the need to check SQL strings at runtime.

By default, nothing from in the `artisan` package is exported with the exception of the `Builder` struct. The struct serves as an entry point for constructing a new SQL statement by limiting API exposure. After construction, a new collection of functions is exposed such as `Create()`, `Select()`, `Insert()` and so on.

For example, using the `Builder` would look something like the following example snippet below:

```go
var (
    b     = artisan.Builder{}
    name  = "david"
    count = 1
    sql   = ""
)

// sql = "CREATE TABLE test(count INTEGER, names TEXT);"
sql = b.Create("test").AddColumn("count", artisan.Integer{}).AddColumn("names", artisan.Text{}).Build()
fmt.Println(sql)

// sql = "SELECT count,name FROM test;"
sql = b.Select("count", "name").From("test").Build()
fmt.Println(sql)

// sql = "SELECT * FROM test WHERE count>=10;"
sql = b.Select().From("test").Where(artisan.IsGreaterThanOrEqual("count", 10)).Build()
fmt.Println(sql)
```

The beauty of this design is most noticeable when this is being typed out in VSCode in realtime. Every function call invocation prompts the editor's autocomplete feature to only show specific parts of the API are immediately available as mentioned above.

But what's the point? By designing the strict APIs to only be exposed in specific states, we're able to completely bypass the need for error checking (assuming that the library is built correctly).

**Note: If there's another name for this type of design, I'd love to know!**

## Known Issues

## TODO

- [ ] Expand API to cover more SQL
- [ ] Create interface to handle different SQL syntax

## License

See [LICENSE.md](./LICENSE.md)
