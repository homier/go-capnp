module bookstest1

go 1.17

require github.com/homier/go-capnp/v3 v3.0.0-alpha.1

require books v1.0.0 // indirect

replace books v1.0.0 => ../books
