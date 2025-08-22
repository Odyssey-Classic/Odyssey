import path from 'path'

export default {
    mode: "development",
    entry: './src/index.ts',
    devtool: 'inline-source-map',
    watch: true,
    module: {
        rules: [
            {
                test: /\.tsx?$/,
                use: 'ts-loader',
                exclude: [
                    /node_modules/, // Exclude node_modules
                    /\.test\.tsx?$/, // Exclude files ending with .test.ts or .test.tsx
                ]
            },
        ],
    },
    resolve: {
        extensions: ['.tsx', '.ts', '.js'],
    },
    output: {
        filename: 'bundle.js',
        path: path.resolve('./dist'),
    },
}
