package main

import(
	"fmt"
	"net/http"
	"errors"
	"ioutil"
	"encoding/json"
)

type Categories []Category //toma a Category y lo reivindica a arrays de cada categoria
type Category struct {
	ID string `json: "id"`
	name string `json: "name"`
}

func main(){
	cats, err := GetCategories("MLA")

	if err != nil {
		fmt.Println("Error:", err.Error())
	}
	fmt.Println("Las categorias son: ....")
}

func GetCategories(siteID string) (Categories ,error) {

	response := http.Get("https://api.mercadolibre.com/sites/MLA/categories") //completar

	bytes := ioutil.ReadAll(response.Bytes) //completar


	var cats Categories
	err =: json.Unmarshal(bytes, &cats) //hay que poner el if etc

	


	return cats, nil

}