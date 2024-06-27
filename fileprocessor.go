package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func processFiles(files []string) (map[string][]string, error) {
	keysInFiles := make(map[string][]string)
	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			return nil, fmt.Errorf("error reading file %s: %w", file, err)
		}

		var result map[string]interface{}
		if err := json.Unmarshal(data, &result); err != nil {
			return nil, fmt.Errorf("error unmarshalling JSON from file %s: %w", file, err)
		}

		processKeys(result, keysInFiles, file)
	}
	return keysInFiles, nil
}

func processKeys(data interface{}, keysInFiles map[string][]string, fileName string) {
	switch value := data.(type) {
	case map[string]interface{}:
		for k, v := range value {
			if _, err := strconv.Atoi(k); err == nil {
				continue
			}
			keysInFiles[k] = appendUnique(keysInFiles[k], fileName)
			processKeys(v, keysInFiles, fileName)
		}
	case []interface{}:
		for _, v := range value {
			processKeys(v, keysInFiles, fileName)
		}
	}
}

func appendUnique(slice []string, item string) []string {
	for _, elem := range slice {
		if elem == item {
			return slice
		}
	}
	return append(slice, item)
}
