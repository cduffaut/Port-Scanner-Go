package output

import (
	"fmt"
	"os"

	"github.com/cduffaut/Port-Scanner-Go/port"
	"gopkg.in/yaml.v3"
)

// stock the result of port scan in yml file
func CreateFolder(result []port.ScanResult) {
	file, err := os.Create("port-scan-result.yml")
	if err != nil {
		fmt.Println("error: failed to create result ports scan file.")
		return
	}
	defer file.Close()

	// encoding data
	encoder := yaml.NewEncoder(file)
	if err := encoder.Encode(result); err != nil {
		fmt.Println("error: YAML encoding failed.")
		return
	}
	fmt.Println("Process ended!")
}
