package main

import "github.com/MHSaeedkia/store-by-gin-redis-postgress/internal/server"

func main() {
	e := server.GetEngine()
	e.Run()
}
