package utils

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"runtime"
)

//GenericJSONParsing Parse into struct
func GenericJSONParsing[T any](filePath string) (T, error) {

	var x T

	rawContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return x, fmt.Errorf("Error when reading the configuration file: %w", err)
	}

	err = json.Unmarshal(rawContent, &x)
	if err != nil {
		return x, fmt.Errorf("Error when unmarshalling the JSON file: %w", err)
	}

	return x, err
}

//GenericYAMLParsing Parse into struct
func GenericYAMLParsing[T any](filePath string) (T, error) {

	var x T

	rawContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return x, fmt.Errorf("Error when reading the configuration file: %w", err)
	}

	err = yaml.Unmarshal(rawContent, &x)
	if err != nil {
		return x, fmt.Errorf("Error when unmarshalling the YAML file: %w", err)
	}

	return x, err
}

//ValueEqual check that 2 pointers to comparable items point to values that are equal
func ValueEqual[T comparable](item1, item2 *T) bool {
	if item1 == nil {
		if item2 != nil {
			return false
		}
	} else {
		if item2 == nil {
			return false
		} else {
			if *item1 != *item2 {
				return false
			}
		}
	}
	return true
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
