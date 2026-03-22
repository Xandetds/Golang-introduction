package main

import("fmt")

func main(){
	x:=55

	if x>5{
		fmt.Println("x is big")
	} else {
		fmt.Println("x is small")
	}



	if x > 5 && x < 15{
		fmt.Println("x is just right")
	}

	if x < 20 || x > 30{
		fmt.Println("x is out of range")
	}

	a:= 11.0
	b:= 20.0

	if frac := a / b; frac > 0.5{
		fmt.Println("a is more than half of b")
	} 



	n:= 2

	switch n{
	case 1:
		fmt.Println("One")
	case 2: 
		fmt. Println("Two")
	case 3:
		fmt.Println("Three")
	default: 
		fmt. Println("Many")
	}



	// naked switch

	switch{
	case n > 100:
		fmt.Println("n is very big")
	case n > 10:
		fmt.Println("n is big")
	default:
		fmt.Println("n is small")
	}
}

