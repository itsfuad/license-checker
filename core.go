package main

import (
	"regexp"
	"strings"
)

// Analyze summarizes a stream of input without modifying user files.
func Analyze(input string) map[string]interface{} {

	found := map[string]string{}
	missing := []string{}
	lines := strings.Split(input, "\n")
	spdx := regexp.MustCompile(`SPDX-License-Identifier:\s*([A-Za-z0-9.+-]+)`)
	for _, line := range lines {
		parts := strings.SplitN(line, "|", 2)
		if len(parts) != 2 {
			continue
		}
		if match := spdx.FindStringSubmatch(parts[1]); len(match) > 1 {
			found[parts[0]] = match[1]
		} else {
			missing = append(missing, parts[0])
		}
	}
	return map[string]interface{}{"licenses": found, "missing": missing}
}
