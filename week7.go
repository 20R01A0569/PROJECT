Week-7:
package main
import "fmt"
func main() {
    var str string
    fmt.Print("Enter String: ")
    fmt.Scan(&str)
    var ispalindrome bool
    ispalindrome=true
    length:=len(str)
    for i:=0;i<length;i++{
        if(str[i]!=str[length-1-i]){
            ispalindrome=false
            break
        }
    }
    if(ispalindrome){
        fmt.Printf("Number %s is a Palindrome",str)
    }else{
        fmt.Printf("Number %s is not a Palindrome",str)
    }
}