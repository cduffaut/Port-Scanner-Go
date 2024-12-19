package main

import (
	"log"
	"os"

	"github.com/cduffaut/Port-Scanner-Go/output"
	"github.com/cduffaut/Port-Scanner-Go/parsing"
	"github.com/cduffaut/Port-Scanner-Go/port"
	"github.com/cduffaut/Port-Scanner-Go/utils"
	"github.com/joho/godotenv"
)

func GetEnv() utils.VarEnv {
	var s_env utils.VarEnv

	godotenv.Load(".env")
	s_env.HOSTNAME = os.Getenv("HOSTNAME")
	s_env.FOR_PORTS = os.Getenv("FOR_PORTS")
	s_env.FOR_RANGE = os.Getenv("FOR_RANGE")
	s_env.VERBOSE = os.Getenv("VERBOSE")
	return s_env
}

// add a maximum count to 1024 max
func main() {
	var bag utils.Bag

	bag.VarEnv = GetEnv()
	parsing.ParseEnv(&bag)

	if bag.VarEnv.HOSTNAME != "" {
		if parsing.ParseIp(bag.VarEnv.HOSTNAME) || parsing.ParseHostname(bag.VarEnv.HOSTNAME) {
			results := port.InitialScan(bag)
			output.CreateFolder(results)
		} else {
			log.Fatalf("error: given hostname: %s is not valid.", bag.VarEnv.HOSTNAME)
		}
	} else {
		log.Fatalf("error: please fill the field \"HOSTNAME=\" in the '.env' file.")
	}
}
