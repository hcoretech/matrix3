package main 

import (
	"fmt"
	"matrix3"
)


func main() {
	identityMatrix := matrix3.NewIdentity()
	fmt.Println("Identity Matrix:" , identityMatrix)	
}