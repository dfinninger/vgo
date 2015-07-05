package gate

import (
  "encoding/json"
  "io/ioutil"
)

type Logic interface {
  Func()
}

type Gate struct {
  Name string `json:"name"`
  FuncDef string `json:"function"`
  neg bool `json:"negation"`
  Inputs []bool
  Output bool
}

// Func() calls the Gate's operator on the input slice
// and then modifies the output accordingly
func (g *Gate) Func() {
  switch g.FuncDef {
  case "conjunction":
    acc := true
    for _, x := range g.Inputs {
      acc = acc && x
    }
    g.Output = acc
  case "disjunction":
    acc := false
    for _, x := range g.Inputs {
      acc = acc || x
    }
    g.Output = acc
  }

  if g.neg {
    g.Output = !g.Output
  }
}

func New(def string) Gate {
  file, err := ioutil.ReadFile(def)
  if err != nil {
      panic(err)
  }

  var retGate Gate
  err = json.Unmarshal(file, &retGate)
  if err != nil {
      panic(err)
  }

  return retGate
}

func NewFromByte(data []byte) Gate {
  var retGate Gate
  err := json.Unmarshal(data, &retGate)
  if err != nil {
      panic(err)
  }

  return retGate
}
