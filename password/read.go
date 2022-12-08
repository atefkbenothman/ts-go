package app

import (
  "bufio"
  "fmt"
  "io"
  "strings"
  "os"
)

func ReadFromFile(fileName string) {
  // open file
  file, err := os.Open(fileName)
  if err != nil {
    fmt.Printf("could not find file: %s\n", fileName)
    return
  }

  // create buffered reader
  reader := bufio.NewReader(file)

  // read file line by line
  for {
    line, err := reader.ReadString('\n')
    line = strings.TrimSuffix(line, "\n")
    if err != nil {
      if err == io.EOF {
        break
      }
    }
    // print
    fmt.Println(line)
  }
}
