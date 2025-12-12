# TextFmt - A Cobra CLI Demonstration

A simple command-line text formatter and analyzer built with Go and Cobra.

## Description

TextFmt is a command-line tool that provides basic text formatting and analysis
capabilities. This project was created as a demonstration for learning Go and
the Cobra library.

## Features

- Convert text to lowercase
- Convert text to uppercase
- Count characters in text
- Count words in text
- Accepts input from arguments or standard input (stdin)

## Installation

### Building from source

Build and install directly:

```bash
go install github.com/br-lemes/textfmt@latest
```

Or clone the source code first:

```bash
git clone https://github.com/br-lemes/textfmt.git
cd textfmt
go build
```

## Usage

```bash
# Process text from arguments:
./textfmt "Hello World" --upper
# Output: HELLO WORLD
```

```bash
# Process text from stdin (pipe):
echo "Hello World" | ./textfmt --lower --words
# Output: hello world
# Word count: 2
```

```bash
# Count characters and words:
./textfmt "This is a test." --chars --words
# Output: This is a test.
# Character count: 15
# Word count: 4
```

## Contributing

Contributions are welcome! Feel free to open issues or pull requests.

## License

This project is licensed under the terms of the [LICENSE](LICENSE) file.
