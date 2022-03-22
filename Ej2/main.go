package main

import( "fmt"
		"errors"
)

func main() {
	div, err := division(7, 2)

	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}
	fmt.Println("Resultado: ", div)
}

func division(m ,n int)(int , error){

	if n == 0{
		return -1, errors.New("n no puede ser 0")
	}

	return m / n, nil
}

