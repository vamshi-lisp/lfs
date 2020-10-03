package main

import(
	"fmt"
	//"strings"
)

func  READ() string{
	var input string
	fmt.Print("\n-> ")
	fmt.Scanf("%s",&input)
	//add to history
	
	return input
}

func EVAL(ast string ,env string) string{
	return ast
}

func PRINT(exp string) string{
	return exp
}

func rep() string{
	return PRINT(EVAL(READ(),""))
}

func main(){

	fmt.Println("\v\t\t\t\t\tR E P L")
	for i := 0 ; i < 10 ; i++ {
		fmt.Print(rep())
	}
}
