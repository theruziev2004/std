package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"

)

type jFile struct {
	Name string
	Login string
	Password string
}

func main() {
	file, err := ioutil.ReadFile("config-settings.json")
	if err != nil {
		fmt.Printf("Error %e",err)
		//os.Exit(1)
	}
	var dataObject jFile
	err = json.Unmarshal(file,&dataObject)
	if err != nil {
		fmt.Println("error:",err)
	}
	fmt.Printf("Name:%s\n", dataObject.Name)
	fmt.Printf("Login:%s\n", dataObject.Login)
	fmt.Printf("Password:%s\n", dataObject.Password)


}
