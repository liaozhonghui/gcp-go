package array_utils

import (
  "fmt"
  "os"
  "testing"
)

func TestMain(m *testing.M) {
  fmt.Println("init resources.")
  result := m.Run()
  fmt.Println("end resources.")
  os.Exit(result)
}
func TestFindMaxSeqSum(t *testing.T) {
  sum := findMaxSeqSum([]int{1, 3, -9, 6, 8, -19})

  if sum == 14 {
    t.Log("test successful.")
  } else {
    t.Log("failed.")
  }
}
func TestOptimizationFindMaxSeqSum(t *testing.T) {
  sum := optimizationFindMaxSum(([]int{1, 3, -9, 8, 8, -19}))

  if sum == 16 {
    t.Log("test successful.")
  } else {
    t.Log("failed.")
  }
}
func BenchmarkFindMaxSeqSum(b *testing.B) {
  for i := 0; i < b.N; i++ {
    findMaxSeqSum([]int{1, 3, -9, 6, 8, -19})
  }
}

func ExampleFindMaxSeqSum() {
  findMaxSeqSum([]int{1, 3, -9, 7, 8, -10})
}
