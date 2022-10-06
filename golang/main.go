package main

import (
	"fmt"
	"os"

	"sanbertutor/routes"
)

func main() {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatalf("Error read env file with err: %s", err)
	// }

	e := routes.Routes()
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))))
}
