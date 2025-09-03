#!/usr/bin/env php
<?php

/**
 * OpenAPIA Validator CLI - PHP Implementation
 * 
 * Command-line interface for validating OpenAPIA specifications.
 * 
 * @package OpenAPIA
 * @version 0.1.0
 * @license Apache-2.0
 */

require_once __DIR__ . '/OpenAPIAValidator.php';

use OpenAPIA\OpenAPIAValidator;

/**
 * CLI entry point
 */
function main(): void
{
    $options = getopt('f:j:q:h', ['file:', 'json', 'quiet', 'help']);
    
    // Show help
    if (isset($options['h']) || isset($options['help'])) {
        showHelp();
        exit(0);
    }
    
    // Get file path
    $filePath = $options['f'] ?? $options['file'] ?? null;
    if (!$filePath) {
        echo "Error: No file specified\n";
        echo "Use --help for usage information\n";
        exit(1);
    }
    
    // Validate file
    $validator = new OpenAPIAValidator();
    $isValid = $validator->validateFile($filePath);
    
    // Output results
    if (isset($options['j']) || isset($options['json'])) {
        // JSON output
        $results = $validator->getResults();
        echo json_encode($results, JSON_PRETTY_PRINT) . "\n";
    } elseif (!isset($options['q']) && !isset($options['quiet'])) {
        // Human-readable output
        $validator->printResults();
    }
    
    exit($isValid ? 0 : 1);
}

/**
 * Show help information
 */
function showHelp(): void
{
    echo "OpenAPIA Validator CLI - PHP Implementation\n";
    echo "==========================================\n\n";
    echo "Usage: php cli.php [options] -f <file>\n\n";
    echo "Options:\n";
    echo "  -f, --file <file>    OpenAPIA specification file to validate\n";
    echo "  -j, --json          Output results as JSON\n";
    echo "  -q, --quiet         Only output errors (no warnings)\n";
    echo "  -h, --help          Show this help message\n\n";
    echo "Examples:\n";
    echo "  php cli.php -f spec.yaml\n";
    echo "  php cli.php --file spec.json --json\n";
    echo "  php cli.php -f spec.yaml --quiet\n\n";
    echo "Supported formats: YAML (.yaml, .yml), JSON (.json)\n";
}

// Run CLI
main();
