package kata

import "fmt"

func countSheep(num int) string {
	var result string

  for i := 1; i <= num; i++ {
    result += fmt.Sprintf("%d", i) + " sheep..."
  }
  
  return result
}
