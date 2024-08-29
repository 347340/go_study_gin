package main
import "fmt"

func Hi (a, b int) int {
  fmt.Println("Hi:123456")
  return a * b
}

func main() {
 var result = Hi(5, 7)
 fmt.Println("result:", result)
}
