/* Removes Byte Order Marks. */

package main

import (
 "fmt"
 "io/ioutil"
 "os"
 "bytes"
)

func main() {
  bom := []byte{0xef, 0xbb, 0xbf} // UTF-8

  if len(os.Args) < 2 {
    fmt.Println("Include file name to parse on command-line")
    return
  }
  fileName := os.Args[1]
  contents, err := ioutil.ReadFile(fileName)
    if err != nil {
      fmt.Println("Error reading file")
      fmt.Println(err)
      return
    }

  if !bytes.Equal(contents[:3], bom) {
    fmt.Println("No byte-order mark found")
    return
  }

  err = os.Rename(fileName, fileName + ".bak")
  if err != nil {
    fmt.Println("Error renaming file")
    fmt.Println(err)
    return
  }

  err = ioutil.WriteFile(fileName, contents[3:], 0644)
  if err != nil {
    fmt.Println("Error re-writing file")
    fmt.Println(err)
    return
  }
}
