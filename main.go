package main

import "chapter2-sesi2/routers"

func main() {
	const PORT = ":8080"

	routers.StartServer().Run(PORT)
}
