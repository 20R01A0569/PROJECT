Week-6:
package main
import "fmt"
func main() {
    fmt.Print("Enter First String: ") 
    var first,second string    
    fmt.Scanln(&first)
    fmt.Print("Enter Second String: ")
    fmt.Scanln(&second)
    fmt.Print("The Added String is:"+(first + second)) 
}