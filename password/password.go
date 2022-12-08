package app

type Password struct {
  Length int
  IncNums bool
  IncSpecialChars bool
  IncUpper bool
}

func (p Password) generate() (string) {
  return "hello"
}
