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

const APAIValidator = require('./src/APAIValidator');
const path = require('path');

/**
 * CLI entry point
 */
async function main() {
    const args = process.argv.slice(2);

    if (args.length === 0 || args.includes('-h') || args.includes('--help') || args[0] === 'help') {
        showHelp();
        process.exit(0);
    }

    const command = args[0];
    const options = args.slice(1);

    switch (command) {
        case 'validate':
            await handleValidate(options);
            break;
        case 'tree':
            handleTree(options);
            break;
        case 'merge':
            await handleMerge(options);
            break;
        default:
            console.log(`Unknown command: ${command}`);
            showHelp();
            process.exit(1);
    }
}

async function handleValidate(options) {
    const filePath = options[0];
    const hierarchical = options.includes('--hierarchical');

    if (!filePath) {
        console.log('Error: No file specified');
        console.log('Usage: node cli.js validate <file> [--hierarchical]');
        process.exit(1);
    }

    console.log(`Validating OpenAPIA specification${hierarchical ? ' with inheritance' : ''}: ${filePath}`);
    console.log('-'.repeat(60));

    const validator = new APAIValidator();

    const isValid = hierarchical ?
        await validator.validateWithInheritance(filePath) :
        await validator.validateFile(filePath);

    if (isValid) {
        console.log('✅ Validation successful!');
    } else {
        console.log('❌ Validation failed!');
        console.log('\nErrors:');
        validator.getErrors().forEach(error => {
            console.log(`  • ${error}`);
        });
    }

    const warnings = validator.getWarnings();
    if (warnings.length > 0) {
        console.log('\nWarnings:');
        warnings.forEach(warning => {
            console.log(`  ⚠️  ${warning}`);
        });
    }

    process.exit(isValid ? 0 : 1);
}

function handleTree(options) {
    const filePath = options[0];

    if (!filePath) {
        console.log('Error: No file specified');
        console.log('Usage: node cli.js tree <file>');
        process.exit(1);
    }

    console.log('OpenAPIA Specification Hierarchy Tree');
    console.log('='.repeat(50));

    const validator = new APAIValidator();
    validator.printHierarchyTree(filePath);
}

async function handleMerge(options) {
    if (options.length < 2) {
        console.log('Error: Missing required arguments');
        console.log('Usage: node cli.js merge <output> <file1> [file2] ...');
        process.exit(1);
    }

    const outputPath = options[0];
    const inputFiles = options.slice(1);

    console.log('Merging OpenAPIA specifications...');
    console.log(`Output: ${outputPath}`);
    console.log(`Input files: ${inputFiles.join(', ')}`);
    console.log('-'.repeat(60));

    const validator = new APAIValidator();
    const specs = [];

    for (const file of inputFiles) {
        if (!require('fs').existsSync(file)) {
            console.log(`Error: File not found: ${file}`);
            process.exit(1);
        }

        const spec = validator.loadSpec(file);
        if (!spec) {
            console.log(`Error: Cannot load specification: ${file}`);
            process.exit(1);
        }

        specs.push(spec);
        console.log(`✅ Loaded: ${file}`);
    }

    const format = require('path').extname(outputPath) === '.json' ? 'json' : 'yaml';
    const success = validator.mergeSpecifications(specs, outputPath, format);

    if (success) {
        console.log('\n✅ Merge completed successfully!');
        console.log(`Merged specification saved to: ${outputPath}`);
    } else {
        console.log('\n❌ Merge failed!');
        validator.getErrors().forEach(error => {
            console.log(`  • ${error}`);
        });
        process.exit(1);
    }
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
    console.log('==================================================');
    console.log('');

    console.log('USAGE:');
    console.log('  node cli.js <command> [options]');
    console.log('');

    console.log('COMMANDS:');
    console.log('  validate <file> [--hierarchical]  Validate OpenAPIA specification');
    console.log('  tree <file>                       Show hierarchy tree for specification');
    console.log('  merge <output> <files...>         Merge multiple specifications');
    console.log('');

    console.log('OPTIONS:');
    console.log('  --hierarchical                   Use hierarchical validation with inheritance');
    console.log('  -h, --help                       Show this help message');
    console.log('');

    console.log('EXAMPLES:');
    console.log('  node cli.js validate spec.yaml');
    console.log('  node cli.js validate spec.yaml --hierarchical');
    console.log('  node cli.js tree spec.yaml');
    console.log('  node cli.js merge output.yaml spec1.yaml spec2.yaml');
    console.log('');

    console.log('For more information, visit: https://github.com/openapia/openapia');
}

// Run CLI
if (require.main === module) {
    main().catch(error => {
        console.error('Unexpected error:', error.message);
        process.exit(1);
    });
}

module.exports = { main, parseArgs, showHelp };
