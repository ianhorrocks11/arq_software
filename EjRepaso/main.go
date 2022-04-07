package main

import(
	"fmt"
	"os"
)


func main(){ //ejercicio en foto 29/03

	func createFile(p string) *os.File {
		fmt.Println("creating")
		f, err := os.Create("/tmp/loquito.txt")
		if err != nil {
		panic(err)
		}
		return f
	}

	func writeFile(f *os.File) {
		fmt.Println("writing")
		fmt.Println(f, "data")
	}

	func closeFile(f *os.File) {
		fmt.Println("closing")
		f.close()
		if err != nil {
			fmt.Fprintf(os.Stderr,
	"error: fv\n", err)
		os.Exit (1)
		}
	}
}

/* ***CREATE AND READ FILES*** (no me salio)

CREATE
	f, _ := os.Create("/tmp/loquito.txt") //*os.File
	fmt.Fprintln(f, "data")
	f.Close()

READ
	f, _ := os.ReadFile("/tmp/loquito.txt")
	fmt.Print(string(f))
	
*/


/* ***DEFER***
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
*/


/* ***CICLOS FOR*** (no me salio)
	
	for pos, char := range "x80" {
		fmt.printf("character %#U starts at byte position %d\n", char, pos)
	}
	*/
