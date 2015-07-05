package gate

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

// ### Builders
func Setup() Gate {
  // Build up some JSON as a byte array
  data := []byte(`{"name": "and","function": "conjunction","negation": false}`)

  // Make a new gate object
  myGate := NewFromByte(data)
  return myGate
}

// ### Test functions
func TestName(t *testing.T) {
  g := Setup()

  // Make sure the name parsed properly
  assert.Equal(t, g.Name, "and", "check name")
}

func TestTrueTrue(t *testing.T) {
  g := Setup()

  // Check the True/True Condition
  g.Inputs = []bool{true, true}
  g.Func()
  assert.Equal(t, g.Output, true, "Check True && True")
}

func TestTrueFalse(t *testing.T) {
  g := Setup()

  // Check the true and false Condition
  g.Inputs = []bool{true, false}
  g.Func()
  assert.Equal(t, g.Output, false, "Check True && False")
}

func TestFalseTrue(t *testing.T) {
  g := Setup()

  // Check the false and true Condition
  g.Inputs = []bool{false, true}
  g.Func()
  assert.Equal(t, g.Output, false, "Check False && True")
}

func TestFalseFalse(t *testing.T) {
  g := Setup()

  // Check the false and false Condition
  g.Inputs = []bool{false, false}
  g.Func()
  assert.Equal(t, g.Output, false, "Check False && False")
}
