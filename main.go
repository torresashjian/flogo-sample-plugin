package main


import (
	"fmt"
	"plugin"
	)

func main(){
	fmt.Println("In main method")
	p, err := plugin.Open("./plugin.so")
	if err != nil {
		fmt.Println("Error on loading")
		fmt.Println(err.Error())
		panic(err.Error())
	}
	
	f, err := p.Lookup("Serialize")
	if err != nil {
		fmt.Println("Error on opening")
		fmt.Println(err.Error())
		panic(err.Error())
	}
	
	f.(func())()

}
