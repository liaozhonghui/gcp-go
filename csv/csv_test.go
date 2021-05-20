package csv

import (
  "os"
  "testing"
)

func TestWriteCsr(t *testing.T) {
  file, _ := os.OpenFile("test.csv", os.O_WRONLY|os.O_CREATE, os.ModePerm)
  writeCsv(file)
}
func TestReadCsv(t *testing.T) {
  file, _ := os.Open("test.csv")
  readCsv(file)
}
