// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2026 The semrel Authors

package main

import (
	"log"

	plugin "github.com/SemRels/updater-homebrew/internal/plugin"
)

func main() {
	updater := plugin.NewUpdater(plugin.FormulaConfig{})
	log.Printf("updater-homebrew plugin ready: updates Homebrew formulae (%T)", updater)
}
