package main

import (
	"log"

	"github.com/LambdaTest/mould/cmd"
)

// Main function just executes root command `ts`
// this project structure is inspired from `cobra` package
func main() {
	if err := cmd.RootCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}
