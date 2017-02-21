package main

import "log"

// First command should be `dog deploy`.
// Flow:
//   - The command would create an archive from current directory
//   - Send the archive to engine api
//   - Print out response
func main() {
	log.Println("Future command line app for communicating with main engine api")
}