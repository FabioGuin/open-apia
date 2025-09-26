package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 || args[0] == "-h" || args[0] == "--help" || args[0] == "help" {
		showHelp()
		os.Exit(0)
	}

	command := args[0]
	options := args[1:]

	switch command {
	case "validate":
		handleValidate(options)
	case "tree":
		handleTree(options)
	case "merge":
		handleMerge(options)
	default:
		fmt.Printf("Unknown command: %s\n", command)
		showHelp()
		os.Exit(1)
	}
}

func handleValidate(options []string) {
	if len(options) == 0 {
		fmt.Println("Error: No file specified")
		fmt.Println("Usage: go run cli.go validate <file> [--hierarchical]")
		os.Exit(1)
	}

	filePath := options[0]
	hierarchical := false
	for _, opt := range options {
		if opt == "--hierarchical" {
			hierarchical = true
			break
		}
	}

	fmt.Printf("Validating APAI specification")
	if hierarchical {
		fmt.Printf(" with inheritance")
	}
	fmt.Printf(": %s\n", filePath)
	fmt.Println(strings.Repeat("-", 60))

	validator := NewAPAIValidator()
	var isValid bool
	var err error

	if hierarchical {
		isValid, err = validator.ValidateWithInheritance(filePath)
	} else {
		isValid, err = validator.ValidateFile(filePath)
	}

	if err != nil {
		fmt.Printf("❌ Validation error: %v\n", err)
		os.Exit(1)
	}

	if isValid {
		fmt.Println("✅ Validation successful!")
	} else {
		fmt.Println("❌ Validation failed!")
		fmt.Println("\nErrors:")
		for _, error := range validator.Errors {
			fmt.Printf("  • %s\n", error)
		}
	}

	if len(validator.Warnings) > 0 {
		fmt.Println("\nWarnings:")
		for _, warning := range validator.Warnings {
			fmt.Printf("  ⚠️  %s\n", warning)
		}
	}

	os.Exit(func() int {
		if isValid {
			return 0
		}
		return 1
	}())
}

func handleTree(options []string) {
	if len(options) == 0 {
		fmt.Println("Error: No file specified")
		fmt.Println("Usage: go run cli.go tree <file>")
		os.Exit(1)
	}

	filePath := options[0]

	fmt.Println("APAI Specification Hierarchy Tree")
	fmt.Println(strings.Repeat("=", 50))

	validator := NewAPAIValidator()
	validator.PrintHierarchyTree(filePath, 0)
}

func handleMerge(options []string) {
	if len(options) < 2 {
		fmt.Println("Error: Missing required arguments")
		fmt.Println("Usage: go run cli.go merge <output> <file1> [file2] ...")
		os.Exit(1)
	}

	outputPath := options[0]
	inputFiles := options[1:]

	fmt.Println("Merging APAI specifications...")
	fmt.Printf("Output: %s\n", outputPath)
	fmt.Printf("Input files: %s\n", strings.Join(inputFiles, ", "))
	fmt.Println(strings.Repeat("-", 60))

	validator := NewAPAIValidator()
	specs := make([]map[string]interface{}, 0, len(inputFiles))

	for _, file := range inputFiles {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			fmt.Printf("Error: File not found: %s\n", file)
			os.Exit(1)
		}

		spec, err := validator.loadSpec(file)
		if err != nil {
			fmt.Printf("❌ Error loading %s: %v\n", file, err)
			os.Exit(1)
		}

		specs = append(specs, spec)
		fmt.Printf("✅ Loaded: %s\n", file)
	}

	format := "yaml"
	if strings.HasSuffix(outputPath, ".json") {
		format = "json"
	}

	err := validator.MergeSpecifications(specs, outputPath, format)
	if err != nil {
		fmt.Printf("\n❌ Merge failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n✅ Merge completed successfully!")
	fmt.Printf("Merged specification saved to: %s\n", outputPath)
}

func showHelp() {
	fmt.Println("APAI Validator CLI - Go Implementation")
	fmt.Println("==========================================")
	fmt.Println("")
	
	fmt.Println("USAGE:")
	fmt.Println("  go run cli.go <command> [options]")
	fmt.Println("")
	
	fmt.Println("COMMANDS:")
	fmt.Println("  validate <file> [--hierarchical]  Validate APAI specification")
	fmt.Println("  tree <file>                       Show hierarchy tree for specification")
	fmt.Println("  merge <output> <files...>         Merge multiple specifications")
	fmt.Println("")
	
	fmt.Println("OPTIONS:")
	fmt.Println("  --hierarchical                   Use hierarchical validation with inheritance")
	fmt.Println("  -h, --help                       Show this help message")
	fmt.Println("")
	
	fmt.Println("EXAMPLES:")
	fmt.Println("  go run cli.go validate spec.yaml")
	fmt.Println("  go run cli.go validate spec.yaml --hierarchical")
	fmt.Println("  go run cli.go tree spec.yaml")
	fmt.Println("  go run cli.go merge output.yaml spec1.yaml spec2.yaml")
	fmt.Println("")
	
	fmt.Println("For more information, visit: https://github.com/FabioGuin/APAI")
}
