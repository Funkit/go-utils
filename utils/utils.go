package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"math/rand"
	"os"
	"runtime"
	"time"
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

func CastNumberAsFloat64(val any) (float64, error) {
	switch v1 := val.(type) {
	case float32:
		return float64(v1), nil
	case float64:
		return float64(v1), nil
	case int:
		return float64(v1), nil
	case int8:
		return float64(v1), nil
	case int16:
		return float64(v1), nil
	case int32:
		return float64(v1), nil
	case int64:
		return float64(v1), nil
	default:
		return 0, fmt.Errorf("item %v cannot be cast as float64", val)
	}
}

// IsALetter Check if char is an ASCII letter or not
func IsALetter(element byte) bool {
	val := ((element >= 'a') && (element <= 'z')) || ((element >= 'A') && (element <= 'Z'))

	return val
}

//AppendIfNew add item to string array if not already in slice
func AppendIfNew(strList []string, newItem string) []string {
	_, found := Find(strList, newItem)
	if !found {
		strList = append(strList, newItem)
	}

	return strList
}

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

//GetFileAsLines reads a file and return the content as a string array.
//
//If some processed lines are longer than 65536 characters, maxLineLength needs to be added as an argument.
func GetFileAsLines(filePath string, maxLineLength ...int) ([]string, error) {
	readFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	var buf []byte
	if len(maxLineLength) != 0 {
		buf = make([]byte, maxLineLength[0])
		fileScanner.Buffer(buf, maxLineLength[0])
	}

	fileScanner.Split(bufio.ScanLines)

	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	if err := fileScanner.Err(); err != nil {
		return nil, err
	}

	output := make([]string, len(fileLines))
	copy(output, fileLines)

	return output, nil
}

func GenerateRandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
