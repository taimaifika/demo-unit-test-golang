# code coverage
> Code coverage is a measure used to describe the degree to which the source code of a program is executed when a particular test suite runs.

## Run tests with coverage
```bash
go test -cover ./...
```
- The `-cover` flag tells the `go test` command to run the tests and generate a coverage report. The coverage report will show the percentage of code that is covered by the tests.

## Run tests with coverage and generate a coverage profile
```bash
go test -coverprofile=coverage.out ./...
```
- The `-coverprofile` flag tells the `go test` command to generate a coverage profile file named `coverage.out`.

## Generate an HTML coverage report
```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o=coverage.html
```
