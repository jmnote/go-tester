# tester

[![pull-request](https://github.com/jmnote/tester/actions/workflows/pull-request.yml/badge.svg)](https://github.com/jmnote/tester/actions/workflows/pull-request.yml)
[![Coverage Status](https://coveralls.io/repos/github/jmnote/tester/badge.svg?branch=main)](https://coveralls.io/github/jmnote/tester?branch=main)
[![GitHub license](https://img.shields.io/github/license/jmnote/tester.svg)](https://github.com/jmnote/tester/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/jmnote/tester)](https://goreportcard.com/report/github.com/jmnote/tester)

`tester` is a Go package designed for generating consistent, human-readable test case names. It provides a function to convert various data types, including complex nested structures, into a string representation that can be used as a test case name.

## Features

- **Consistent Naming:** Converts input objects into a string format, ensuring consistency and readability.
- **Supports Multiple Types:** Handles different types of inputs including integers, slices, arrays, maps, and structs.
- **Automatic String Formatting:** Automatically removes unnecessary whitespaces, replaces slashes with `%`, and truncates names that are too long to maintain readability.

## Installation

You can install the package using `go get`:

```bash
go get github.com/jmnote/tester
```

## Usage

To use the `Name` function in your tests, you can create a test function like the following:

```go
import "github.com/jmnote/tester/testcase"
...
	for i, tc := range testCases {
		t.Run(testcase.Name(i, tc.input), func(t *testing.T) {
			...
	}
```

## License

This project is licensed under the Apache License 2.0. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request or open an Issue to discuss any changes or improvements.

## Acknowledgments

Thanks to all the contributors who have helped make this project better.

