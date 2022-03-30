package util

import (
	"fmt"
	"log"
	"regexp"
)

type Array struct {
}

//提取数组
func (s Array) ExtractArray(src map[string][]string, key string) map[string]interface{} {
	result := make(map[string]interface{})

	reg, err := regexp.Compile(fmt.Sprintf(`^(%s)\[([a-z0-9]+)\]$`, key))
	if err != nil {
		log.Fatalf("Error compiling regexp: %v", err)
	}
	var matches [][]string
	for k, v := range src {
		matches = reg.FindAllStringSubmatch(k, -1)
		if len(matches) != 1 {
			continue
		}

		if len(matches[0]) != 3 {
			continue
		}

		if key != "" && matches[0][1] != key {
			continue
		}

		result[matches[0][2]] = v[0]
	}

	return result
}
