## Testing

1. test functions should start with capital T
## Writing test cases
```go
package main
func Test_isPrime(t *testing.T) {}
```
## Running single test
```bash
go test -run Test_XYZ
```

## Testing coverage
```bash
go test -cover . 
 ```
```bash
ok      primeapp        0.181s  coverage: 72.7% of statements

```
```bash
go test -coverprofile=coverage.out
```

```bash
go tool cover -html=coverage.out
```
