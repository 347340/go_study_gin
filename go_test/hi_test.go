package main
import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestHi(t *testing.T) {
  expectedResult := 35
  returnResult := Hi(5, 7)
  assert.Equal(t, expectedResult, returnResult)
}
