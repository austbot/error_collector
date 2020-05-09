package error_collector

import (
  "fmt"
  "sync"

  "github.com/pkg/errors"
)

type ErrorsCollector struct {
  sync.RWMutex
  errorSlice []error
}

func NewErrorCollector() ErrorsCollector {
  return ErrorsCollector{
    errorSlice: []error{},
  }
}

func (e ErrorsCollector) Count() int {
  return len(e.errorSlice)
}

func (e *ErrorsCollector) New(err string) {
  if err != "" {
    e.errorSlice = append(e.errorSlice, errors.New(err))
  }
}

func (e *ErrorsCollector) Add(err error) {
  if err != nil {
    e.Lock()
    defer e.Unlock()
    e.errorSlice = append(e.errorSlice, errors.Wrap(err, ""))
  }
}

func (e *ErrorsCollector) Error() string {
  var errorMessage string
  e.Lock()
  defer e.Unlock()
  if len(e.errorSlice) > 0 {
    for i, specError := range e.errorSlice {
      if errorMessage != "" {
        errorMessage = fmt.Sprintf("%s\n----Error %d-----\n%+v", errorMessage, i+1, specError)
      } else {
        errorMessage = fmt.Sprintf("----Error %d-----\n%+v", i+1, specError)
      }
    }
  } else {
    return "----No Errors----"
  }
  return errorMessage
}
