package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type SchedkitImage struct {
	ImageURI string `json:"imageURI"`
}

type ManifestMap = map[string]SchedkitImage

func main() {
	entries, err := os.ReadDir(".")
	if err != nil {
		fmt.Fprintf(os.Stderr, "read dir: %v\n", err)
		os.Exit(1)
	}

	m := make(ManifestMap)
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		subEntries, err := os.ReadDir(e.Name())
		if err != nil {
			continue // skip unreadable sub-directories
		}

		if HasDockerfile(subEntries) {
			schedulerName := e.Name()
			schedulerImage := fmt.Sprintf("ghcr.io/schedkit/%s", e.Name())
			m[schedulerName] = SchedkitImage{ImageURI: schedulerImage}
		}
	}

	out, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "json: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(out))
}

func HasDockerfile(subEntries []os.DirEntry) bool {
	for _, se := range subEntries {
		if !se.IsDir() && strings.EqualFold(se.Name(), "dockerfile") {
			return true
		}
	}
	return false
}
