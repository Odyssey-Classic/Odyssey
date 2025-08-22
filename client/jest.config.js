module.exports = {
    preset: 'ts-jest', // Use ts-jest for TypeScript support
    testEnvironment: '', // Set the test environment to Node.js
    moduleFileExtensions: ['ts', 'js'], // Recognize TypeScript and JavaScript files
    moduleDirectories: ['node_modules', 'src'],
    testMatch: ['**/*.test.ts'], // Match test files with .test.ts extension
    roots: ['<rootDir>/src'], // Define the root directory for tests
    globals: {
        'ts-jest': {
            tsconfig: '<rootDir>/tsconfig.json', // Path to the TypeScript configuration file
        },
    },
};
