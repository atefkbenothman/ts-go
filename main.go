package main

import (
  "fmt"
  "strconv"
  "app/password"
)

func main() {
  // ask()
  // read()
  connect()
}

func connect() {
  app.ReadFromConnection("localhost:8080")
}

func read() {
  app.ReadFromFile("password/test.txt")
}

func ask() {
  // length
  var passwordLengthStr string;

  fmt.Print("length of password: ");
  fmt.Scanln(&passwordLengthStr);

  passwordLengthInt, err := strconv.Atoi(passwordLengthStr)
  if err != nil {
    return
  }

  // numbers
  var includeNumbersString string
  var includeNumbersBool bool = false

  fmt.Print("include numbers? ")
  fmt.Scanln(&includeNumbersString)

  if includeNumbersString == "y" {
    includeNumbersBool = true
  }

  // special characters
  var includeSpecialCharsString string
  var includeSpecialCharsBool bool = false

  fmt.Print("include special characters? ");
  fmt.Scanln(&includeSpecialCharsString);

  if includeSpecialCharsString == "y" {
    includeSpecialCharsBool = true
  }

  // upper case
  var includeUppercaseString string
  var includeUppercaseBool bool = false

  fmt.Print("include uppercase? ");
  fmt.Scanln(&includeUppercaseString);

  if includeUppercaseString == "y" {
    includeUppercaseBool = true
  }

  // generate
  p := app.Password{
    Length: passwordLengthInt,
    IncNums: includeNumbersBool,
    IncSpecialChars: includeSpecialCharsBool,
    IncUpper: includeUppercaseBool,
  }
  fmt.Println(p)
}
