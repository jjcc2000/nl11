package main

import (
	"encoding/json"
	"fmt"
	
)

type Person struct{
	First string
	Last string
	Age int
}

func main(){
	ej3()
}

func ej(){
	p1 := Person{
		First: "Johan",
		Last: "Chacon",
		Age: 12,
	}
	bs,err :=json.Marshal(p1)

	if err!= nil{
		fmt.Println("This is the error:\t",err)
	}else{
		fmt.Println(string(bs))
	}
}
func ej2(){
	type pperson struct{
		First string
		Last string
		Age int 
	}
	p := pperson{
		First: "Sergio",
		Last: "Chacon",
		Age: 8,
	}
	bs, err:= func (p pperson) ([]byte, error){
		bs,err := json.Marshal(p)
		if err != nil{
			fmt.Println("This is the error:",err)
			return	[]byte{}, fmt.Errorf("This is the error: %v",err)
		}
			return bs, nil
	}(p)
	if err!= nil{
		fmt.Println(err)
	}else {
		fmt.Println(string(bs))
	}
}
//FIXME:///////////////////////////


type customErr struct{
	info string
}
func (c customErr) Error() string{
	return fmt.Sprintf("This is the error: %v",c.info)	
}

func foo(e error){
	fmt.Println("foo ran-",e)
}


func ej3(){
	//Implement the build it interface ERROR
	
	c := customErr{
		info: "Info from the customer",
	}
	foo(c)
}
