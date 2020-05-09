# Error Collector
Error Collector. Useful when you want to delay error handling to the end, such as in go routines, data processing pipelines and minor error modes.

## Install
`go get github.com/austbot/error_collector`
## Use
```go

import (
  "github.com/austbot/error_collector"
)

func main() {
  errorsC := NewErrorCollector()
  f, err := os.Open("/test.txt")
  // Add
  errorsC.add(err)
  
  //...
  errorsC.Error() // Returns an string with stack traces and reasons for all embedded errors.
}
```
