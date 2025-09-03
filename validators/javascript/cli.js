#!/usr/bin/env node

/**
 * OpenAPIA Validator CLI - JavaScript Implementation
 * 
 * Command-line interface for validating OpenAPIA specifications.
 * 
 * @package OpenAPIA
 * @version 0.1.0
 * @license Apache-2.0
 */

const OpenAPIAValidator = require('./src/OpenAPIAValidator');
const path = require('path');

/**
 * CLI entry point
 */
async function main() {
    const args = process.argv.slice(2);
    const options = parseArgs(args);

    // Show help
    if (options.help) {
        showHelp();
        process.exit(0);
    }

    // Get file path
    if (!options.file) {
        console.error('Error: No file specified');
        console.error('Use --help for usage information');
        process.exit(1);
    }

    // Validate file
    const validator = new OpenAPIAValidator();
    const isValid = await validator.validateFile(options.file);

    // Output results
    if (options.json) {
        // JSON output
        const results = validator.getResults();
        console.log(JSON.stringify(results, null, 2));
    } else if (!options.quiet) {
        // Human-readable output
        validator.printResults();
    }

    process.exit(isValid ? 0 : 1);
}

/**
 * Parse command line arguments
 */
function parseArgs(args) {
    const options = {
        file: null,
        json: false,
        quiet: false,
        help: false
    };

    for (let i = 0; i < args.length; i++) {
        const arg = args[i];

        switch (arg) {
            case '-f':
            case '--file':
                options.file = args[++i];
                break;
            case '-j':
            case '--json':
                options.json = true;
                break;
            case '-q':
            case '--quiet':
                options.quiet = true;
                break;
            case '-h':
            case '--help':
                options.help = true;
                break;
            default:
                if (!options.file && !arg.startsWith('-')) {
                    options.file = arg;
                }
                break;
        }
    }

    return options;
}

/**
 * Show help information
 */
function showHelp() {
    console.log('OpenAPIA Validator CLI - JavaScript Implementation');
    console.log('================================================');
    console.log('');
    console.log('Usage: node cli.js [options] -f <file>');
    console.log('');
    console.log('Options:');
    console.log('  -f, --file <file>    OpenAPIA specification file to validate');
    console.log('  -j, --json          Output results as JSON');
    console.log('  -q, --quiet         Only output errors (no warnings)');
    console.log('  -h, --help          Show this help message');
    console.log('');
    console.log('Examples:');
    console.log('  node cli.js -f spec.yaml');
    console.log('  node cli.js --file spec.json --json');
    console.log('  node cli.js -f spec.yaml --quiet');
    console.log('');
    console.log('Supported formats: YAML (.yaml, .yml), JSON (.json)');
}

// Run CLI
if (require.main === module) {
    main().catch(error => {
        console.error('Unexpected error:', error.message);
        process.exit(1);
    });
}

module.exports = { main, parseArgs, showHelp };
