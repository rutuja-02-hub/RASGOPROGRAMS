package main
import "fmt"

func main() {
    var password string
    correctPassword := "Go1234"

    for {
        fmt.Print("Enter password: ")
        fmt.Scan(&password)
        if password == correctPassword {
            fmt.Println("Password correct. Access granted!")
            break
        } else {
            fmt.Println("Incorrect password. Try again.")
        }
    }
}
