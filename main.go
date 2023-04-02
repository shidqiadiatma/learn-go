package main

import (
	"chapter2-sesi3/database"
	routers "chapter2-sesi3/routes"

	_ "github.com/lib/pq"
)

func main() {
	db := database.GetConnection()
	defer db.Close()

	const PORT = ":8080"

	routers.StartServer().Run(PORT)
}
