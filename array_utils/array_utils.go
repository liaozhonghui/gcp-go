package array_utils

import (
  "fmt"
)

func findMaxSeqSum(array []int) int {
  SeqSum := make([]int, 0)
  SeqSum = append(SeqSum, array[0])

  for i := 1; i < len(array); i++ {
    if SeqSum[i-1] > 0 {
      SeqSum = append(SeqSum, SeqSum[i-1]+array[i])
    } else {
      SeqSum = append(SeqSum, array[i])
    }
  }
  max := SeqSum[0]
  for j := 1; j < len(SeqSum); j++ {
    if SeqSum[j] > max {
      max = SeqSum[j]
    }
  }
  fmt.Println(max)
  return max
}

func optimizationFindMaxSum(array []int) int {
  SeqSum := make([]int, len(array))
  for i := 0; i < len(array); i++ {
    SeqSum[i] = array[i]
  }

  for j := 1; j < len(SeqSum); j++ {
    tmp := SeqSum[j-1]
    if tmp < 0 {
      tmp = 0
    }
    SeqSum[j] += tmp
  }
  max := SeqSum[0]
  for k := 1; k < len(SeqSum); k++ {
    if max < SeqSum[k] {
      max = SeqSum[k]
    }
  }
  fmt.Printf("optimization: %d \n", max)
  return max
}
