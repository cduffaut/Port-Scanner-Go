package port

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/cduffaut/Port-Scanner-Go/utils"
)

type ScanResult struct {
	Port  string `yaml:"Port"`
	State string `yaml:"State"`
}

func ScanPort(protocol, hostname string, port int) ScanResult {
	result := ScanResult{Port: protocol + "/" + strconv.Itoa(port)}
	address := hostname + ":" + strconv.Itoa(port)
	// net.DialTimeout establish a network connection
	conn, err := net.DialTimeout(protocol, address, 3*time.Second)

	if err != nil {
		result.State = "Close"
		return result
	}

	result.State = "Open"
	defer conn.Close()
	return result
}

func InitialScan(bag utils.Bag) []ScanResult {
	var results []ScanResult

	// scanning the list port
	for i, port := range bag.PortList {
		results = append(results, ScanPort("tcp", bag.VarEnv.HOSTNAME, port))
		results = append(results, ScanPort("udp", bag.VarEnv.HOSTNAME, port))
		state := float64(i) * 100 / float64(len(bag.PortList))
		if bag.VarEnv.VERBOSE == "True" || bag.VarEnv.VERBOSE == "true" || bag.VarEnv.VERBOSE == "TRUE" {
			fmt.Printf("[port %d] : %.2f%% ...\n", port, state)
		}
	}
	// scanning the range port
	for i := bag.PortRange.Start; i <= bag.PortRange.End; i++ {
		results = append(results, ScanPort("tcp", bag.VarEnv.HOSTNAME, i))
		results = append(results, ScanPort("udp", bag.VarEnv.HOSTNAME, i))
		state := float64(i) * 100 / float64(bag.PortRange.End)
		if bag.VarEnv.VERBOSE == "True" || bag.VarEnv.VERBOSE == "true" || bag.VarEnv.VERBOSE == "TRUE" {
			fmt.Printf("[port %d/%d] : %.2f%% ...\n", i, bag.PortRange.End, state)
		}
	}

	return results
}
