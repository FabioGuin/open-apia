#!/usr/bin/env php
<?php

/**
 * APAI Validator CLI - PHP Implementation
 * 
 * Command-line interface for validating APAI specifications.
 * 
 * @package APAI
 * @version 0.1.0
 * @license Apache-2.0
 */

require_once __DIR__ . '/vendor/autoload.php';

use APAI\APAIValidator;

/**
 * CLI entry point
 */
function main(): void
{
    global $argv;
    $args = array_slice($argv ?? [], 1);
    
    if (empty($args) || in_array($args[0], ['-h', '--help', 'help'])) {
        showHelp();
        exit(0);
    }
    
    $command = $args[0];
    $options = array_slice($args, 1);
    
    switch ($command) {
        case 'validate':
            handleValidate($options);
            break;
        case 'tree':
            handleTree($options);
            break;
        case 'merge':
            handleMerge($options);
            break;
        default:
            echo "Unknown command: {$command}\n";
            showHelp();
            exit(1);
    }
}

function handleValidate(array $options): void
{
    $filePath = $options[0] ?? null;
    $hierarchical = in_array('--hierarchical', $options);
    
    if (!$filePath) {
        echo "Error: No file specified\n";
        echo "Usage: php cli.php validate <file> [--hierarchical]\n";
        exit(1);
    }
    
    echo "Validating APAI specification" . ($hierarchical ? " with inheritance" : "") . ": {$filePath}\n";
    echo str_repeat('-', 60) . "\n";
    
    $validator = new APAIValidator();
    
    if ($hierarchical) {
        $isValid = $validator->validateWithInheritance($filePath);
    } else {
        $isValid = $validator->validateFile($filePath);
    }
    
    if ($isValid) {
        echo "✅ Validation successful!\n";
    } else {
        echo "❌ Validation failed!\n";
        echo "\nErrors:\n";
        foreach ($validator->getErrors() as $error) {
            echo "  • {$error}\n";
        }
    }
    
    if (!empty($validator->getWarnings())) {
        echo "\nWarnings:\n";
        foreach ($validator->getWarnings() as $warning) {
            echo "  ⚠️  {$warning}\n";
        }
    }
    
    exit($isValid ? 0 : 1);
}

function handleTree(array $options): void
{
    $filePath = $options[0] ?? null;
    
    if (!$filePath) {
        echo "Error: No file specified\n";
        echo "Usage: php cli.php tree <file>\n";
        exit(1);
    }
    
    echo "APAI Specification Hierarchy Tree\n";
    echo str_repeat('=', 50) . "\n";
    
    $validator = new APAIValidator();
    $validator->printHierarchyTree($filePath);
}

function handleMerge(array $options): void
{
    if (count($options) < 2) {
        echo "Error: Missing required arguments\n";
        echo "Usage: php cli.php merge <output> <file1> [file2] ...\n";
        exit(1);
    }
    
    $outputPath = $options[0];
    $inputFiles = array_slice($options, 1);
    
    echo "Merging APAI specifications...\n";
    echo "Output: {$outputPath}\n";
    echo "Input files: " . implode(', ', $inputFiles) . "\n";
    echo str_repeat('-', 60) . "\n";
    
    $validator = new APAIValidator();
    $specs = [];
    
    foreach ($inputFiles as $file) {
        if (!file_exists($file)) {
            echo "Error: File not found: {$file}\n";
            exit(1);
        }
        
        $spec = $validator->loadSpec($file);
        if ($spec === null) {
            echo "Error: Cannot load specification: {$file}\n";
            exit(1);
        }
        
        $specs[] = $spec;
        echo "✅ Loaded: {$file}\n";
    }
    
    $format = pathinfo($outputPath, PATHINFO_EXTENSION) === 'json' ? 'json' : 'yaml';
    $success = $validator->mergeSpecifications($specs, $outputPath, $format);
    
    if ($success) {
        echo "\n✅ Merge completed successfully!\n";
        echo "Merged specification saved to: {$outputPath}\n";
    } else {
        echo "\n❌ Merge failed!\n";
        foreach ($validator->getErrors() as $error) {
            echo "  • {$error}\n";
        }
        exit(1);
    }
}

/**
 * Show help information
 */
function showHelp(): void
{
    echo "APAI Validator CLI - PHP Implementation\n";
    echo "============================================\n\n";
    
    echo "USAGE:\n";
    echo "  php cli.php <command> [options]\n\n";
    
    echo "COMMANDS:\n";
    echo "  validate <file> [--hierarchical]  Validate APAI specification\n";
    echo "  tree <file>                       Show hierarchy tree for specification\n";
    echo "  merge <output> <files...>         Merge multiple specifications\n\n";
    
    echo "OPTIONS:\n";
    echo "  --hierarchical                   Use hierarchical validation with inheritance\n";
    echo "  -h, --help                       Show this help message\n\n";
    
    echo "EXAMPLES:\n";
    echo "  php cli.php validate spec.yaml\n";
    echo "  php cli.php validate spec.yaml --hierarchical\n";
    echo "  php cli.php tree spec.yaml\n";
    echo "  php cli.php merge output.yaml spec1.yaml spec2.yaml\n\n";
    
    echo "For more information, visit: https://github.com/FabioGuin/APAI\n";
}

// Run CLI
main();
