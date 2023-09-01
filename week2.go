Week-2:
package main

import "fmt"

func main() {
	var rows, temp int
	fmt.Print("Enter number of rows :")
	fmt.Scan(&rows)
	for i := 1; i <= rows; i++ {
		temp=i
		for j := 1; j <= rows-i; j++ {
			fmt.Print("  ")
		}
		for j := 1; j <= i; j++{ 
			fmt.Print(" ", temp)
			temp++;
		}
		temp--;
		for{
			if(temp==i){
				break
			}
			temp--;
			fmt.Print(" ",temp)
		}
		fmt.Println()
	}

}