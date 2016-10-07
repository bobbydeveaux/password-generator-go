#Password Generator In Go

Example of a password generator written in Go using flags functionality.

## Command-line flags

You can modify the output with these flags:

- **--length**	Length of passwords. Default is `20`.
- **--upper**	Include upper-case characters `A-Z`. Default is `false`.
- **--lower**	Include lower-case characters `a-z`. Default is `false`.
- **--numbers**	Include numbers `0-9`. Default is `false`.
- **--special**	Include special characters `$%&*@^`. Default is `false`.

If non of the above flags are specified, a password containing a combination of all will be used.

## Usage examples

Generate passwords with defaults:

	$ go run main.go

Generate upper-case passwords only:

	$ go run main.go --upper

Genereate a mixture of lower & special with a length of 24

	$ go run main.go --lower --special --length 24