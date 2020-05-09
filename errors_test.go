package error_collector

import (
  "testing"

  "github.com/pkg/errors"
  "github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
  ec := NewErrorCollector()

  ec.Add(errors.New("Issue here"))
  ec.Add(nil)

  assert.Equal(t, 1, ec.Count(), "Count mismatch, nil errors should not be added")
}

func TestGet(t *testing.T) {

  ec := NewErrorCollector()

  ec.Add(errors.New("Issue here"))
  ec.Add(errors.New("Issue There"))

  assert.Contains(t,ec.Error(), "----Error 1-----","Error string mismatch, Get should combine errors")
  assert.Contains(t,ec.Error(), "----Error 2-----","Error string mismatch, Get should combine errors")

}
