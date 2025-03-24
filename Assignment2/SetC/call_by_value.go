package main
import "fmt"

func modifyValue(num int) {
    num = num + 10
    fmt.Println("Inside modifyValue, num =", num)
}

func main() {
    num := 5
    fmt.Println("Before calling modifyValue, num =", num)
    
    modifyValue(num) 
    fmt.Println("After calling modifyValue, num =", num)
}
