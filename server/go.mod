module github.com/aixoio/amionline/server

go 1.21.4

require github.com/gorilla/mux v1.8.1

require (
	github.com/aixoio/amionline/logger v0.0.0-20231118080850-8ee983c5f65c
	github.com/go-sql-driver/mysql v1.7.1
	github.com/redis/go-redis/v9 v9.3.0
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
)
