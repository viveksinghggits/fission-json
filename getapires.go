package main
import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)
type todo struct{
	UserId int 
	Id int 
	Title string
	Completed bool 
}

func Handler(w http.ResponseWriter, r *http.Request){
	var aToDo todo
	URL:="https://jsonplaceholder.typicode.com/todos/1"
	fmt.Println("We will be trying to hit the the URI ", URL)	

	response, err:=http.Get(URL)
	if err!=nil{
		fmt.Println("Error while calling the API ", err)
	} else{
		data, _:= ioutil.ReadAll(response.Body)
		msg:=string(data)
		fmt.Println("Data that we have got from the API is ", string(data))
		json.Unmarshal([]byte(msg), &aToDo)
		w.Write([]byte(aToDo.Title))
	}


}
