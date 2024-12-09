package main

import (
	"fmt"
	"os"

	"github.com/cduffaut/Port-Scanner-Go/port"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	hostname := os.Getenv("HOSTNAME")
	results := port.InitialScan(hostname)
	fmt.Println(results)
}
