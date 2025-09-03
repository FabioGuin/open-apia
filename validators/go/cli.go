package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func main() {
	var (
		filePath = flag.String("f", "", "OpenAPIA specification file to validate")
		jsonOutput = flag.Bool("j", false, "Output results as JSON")
		quiet = flag.Bool("q", false, "Only output errors (no warnings)")
		help = flag.Bool("h", false, "Show help message")
	)
	flag.Parse()

	// Show help
	if *help {
		showHelp()
		os.Exit(0)
	}

	// Get file path
	if *filePath == "" {
		fmt.Fprintf(os.Stderr, "Error: No file specified\n")
		fmt.Fprintf(os.Stderr, "Use -h for usage information\n")
		os.Exit(1)
	}

	// Validate file
	validator := NewOpenAPIAValidator()
	isValid, err := validator.ValidateFile(*filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Output results
	if *jsonOutput {
		// JSON output
		results := validator.GetResults()
		jsonData, err := json.MarshalIndent(results, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(jsonData))
	} else if !*quiet {
		// Human-readable output
		validator.PrintResults()
	}

	os.Exit(func() int {
		if isValid {
			return 0
		}
		return 1
	}())
}

func showHelp() {
	fmt.Println("OpenAPIA Validator CLI - Go Implementation")
	fmt.Println("==========================================")
	fmt.Println("")
	fmt.Println("Usage: go run . [options] -f <file>")
	fmt.Println("")
	fmt.Println("Options:")
	fmt.Println("  -f <file>    OpenAPIA specification file to validate")
	fmt.Println("  -j           Output results as JSON")
	fmt.Println("  -q           Only output errors (no warnings)")
	fmt.Println("  -h           Show this help message")
	fmt.Println("")
	fmt.Println("Examples:")
	fmt.Println("  go run . -f spec.yaml")
	fmt.Println("  go run . -f spec.json -j")
	fmt.Println("  go run . -f spec.yaml -q")
	fmt.Println("")
	fmt.Println("Supported formats: YAML (.yaml, .yml), JSON (.json)")
}
