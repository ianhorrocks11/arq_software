//Escriba un test para la siguiente funcion

package main

import "fmt"

func Division(a int, b int)(int, error) {
	if b==0 {
		return 0, errors.New("No puedo dividir por 0")}	
	return a/b, nil
}
func main() {
	fmt.Println(Division(8, 4))
}
