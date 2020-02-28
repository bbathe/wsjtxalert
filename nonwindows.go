// +build !windows

package main

import "log"

// alertSound generates an aubible alert
func alertSound() {
	// satisfy 'unused' linter
	log.Printf("alert")
}
