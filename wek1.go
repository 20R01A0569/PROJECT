Week-1:
package main

import "fmt"
func gcd (temp1 int,temp2 int) int{ 
	if(temp1==0){ 
		return temp2
	}
	return gcd (temp2%temp1,temp1)
}
func main() {
	var n1,n2 int
	fmt.Println("Enter two positive integers : ")
	fmt.Scan(&n1)
	fmt. Scan(&n2)
	fmt.Printf("GCD of %d and %d is %d\n",n1,n2, gcd(n1,n2))
	fmt.Printf("LCM of %d and %d is %d",n1,n2,(n1*n2)/gcd (n1,n2))

}