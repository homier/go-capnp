module hashtest

go 1.17

require github.com/homier/go-capnp/v3 v3.0.0-alpha.4

require hashes v1.0.0 // indirect

replace hashes v1.0.0 => ./hashes
