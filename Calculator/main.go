package main

import "fmt"

func main() {

calculator()

}

func calculator() {
 var a int
 var b int
 var op string

 fmt.Print("Enter number a:")
 fmt.Scan(&a)

 fmt.Print("Enter number b:")
 fmt.Scan(&b)

 fmt.Printf("Select operation:\na.Add\nb.SUB\nc.MUL\nd.DIV\n")
 fmt.Scan(&op)

 switch op{
 	case "a":
		fmt.Printf("%v+%v =%v",a,b,a+b)
	case "b":
		fmt.Printf("%v-%v =%v",a,b,a-b)
	case "c":
		fmt.Printf("%v*%v =%v",a,b,a*b)
	case "d":
		err := div(a,b)
		if err!=nil{
			fmt.Println(err.Error())
		}
	default:
		fmt.Println("Inalid operation")
		
 }
}

func div[T int|float64](a,b T) error{
	if b == 0 {
    return fmt.Errorf("division by zero")
    }
	fmt.Printf("%v/%v =%v",a,b,a/b)
	return nil
}