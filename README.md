# go-tester

`go-tester` is a Go package designed for generating consistent, human-readable test case names. It provides a function to convert various data types, including complex nested structures, into a string representation that can be used as a test case name.

## Features

- **Consistent Naming:** Converts input objects into a string format, ensuring consistency and readability.
- **Supports Multiple Types:** Handles different types of inputs including integers, slices, arrays, maps, and structs.
- **Automatic String Formatting:** Automatically removes unnecessary whitespaces, replaces slashes with `%`, and truncates names that are too long to maintain readability.

## Installation

You can install the package using `go get`:

```bash
go get github.com/jmnote/go-tester
```

## Usage

To use the `Name` function in your tests, you can create a test function like the following:

```go
import "github.com/jmnote/go-tester/testcase"
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

