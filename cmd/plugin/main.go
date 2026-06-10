// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2026 The semrel Authors

package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	plugin "github.com/SemRels/updater-homebrew/internal/plugin"
)

const pluginSchemaVersion = 1

func main() {
	os.Exit(run(os.Stdout, os.Stderr, os.Getenv))
}

func run(stdout, stderr io.Writer, getenv func(string) string) int {
	fmt.Fprintf(stderr, "plugin_schema_version=%d\n", pluginSchemaVersion)
	version := getenv("SEMREL_VERSION")
	if version == "" {
		version = getenv("SEMREL_NEXT_VERSION")
	}
	if version == "" {
		fmt.Fprintln(stderr, "updater-homebrew: SEMREL_VERSION is required")
		return 1
	}
	version = strings.TrimPrefix(version, "v")

	file := getenv("SEMREL_PLUGIN_FILE")
	if file == "" {
		file = "Formula.rb"
	}

	if getenv("SEMREL_DRY_RUN") == "true" {
		fmt.Fprintf(stdout, "updater-homebrew: [dry-run] would update %s to version %s\n", file, version)
		return 0
	}

	if err := plugin.NewUpdater().Update(file, version); err != nil {
		fmt.Fprintln(stderr, "updater-homebrew:", err)
		return 1
	}

	fmt.Fprintf(stdout, "updater-homebrew: updated %s to version %s\n", file, version)
	return 0
}
