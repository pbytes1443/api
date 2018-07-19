package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"log"
	"io/ioutil"
	"github.com/gorilla/mux"
	"strconv"
	//"flags"
	)
type Person struct{
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Bal string `json:"bal,omitempty"`
	}
var people []Person

func getBalance(w http.ResponseWriter, req *http.Request){
	params:= mux.Vars(req)
	for _,item:=range people{
		if item.Id == params["id"]{
			
			json.NewEncoder(w).Encode(item)
			return
			}
		}
		json.NewEncoder(w).Encode(people)
	}

func addUser(w http.ResponseWriter, req *http.Request){
	var person Person;
	body, err:= ioutil.ReadAll(req.Body)
	if err!=nil{
		return
		}
	if err := json.Unmarshal(body, &person); err != nil{
        panic(err)
   	 }
	people=append(people,person)
	 fmt.Println(person)
	json.NewEncoder(w).Encode(people)
}
func deposit(w http.ResponseWriter, req *http.Request){
	params:= mux.Vars(req)
	var person Person;
	body, err:= ioutil.ReadAll(req.Body)
	if err!=nil{
		panic(err)
		}
	if err := json.Unmarshal(body, &person); err != nil{
        panic(err)
   	 }
	fmt.Println(people[1])
	ind:=0
	for _,pers:=range people{
		if pers.Id == params["id"]{
			b, err := strconv.Atoi(person.Bal)
   				 if err == nil {
					fmt.Println(b)
				    }
			b1, err := strconv.Atoi(pers.Bal)
   				 if err == nil {
					fmt.Println(b1)
				    }
			bal:=b1+b;
				
			bl:=string(bal)
			people[ind].Bal=bl
			json.NewEncoder(w).Encode(pers) 
			}
		ind++
		}
	people=append(people,person)
	fmt.Println(people)
		json.NewEncoder(w).Encode(people)
	}
func withDraw(w http.ResponseWriter, req *http.Request){
	params:= mux.Vars(req)
	var person Person;
	body, err:= ioutil.ReadAll(req.Body)
	if err!=nil{
		panic(err)
		}
	if err := json.Unmarshal(body, &person); err != nil{
        panic(err)
   	 }
	ind:=0
	for _,pers:=range people{
		if pers.Id == params["id"]{
			b, err := strconv.Atoi(person.Bal)
   				 if err == nil {
					fmt.Println(b)
				    }
			b1, err := strconv.Atoi(pers.Bal)
   				 if err == nil {
					fmt.Println(b1)
				    }
			bal:=string(b1-b);
			people[ind].Bal=bal
			json.NewEncoder(w).Encode(pers)
			}
		ind++
		}
		json.NewEncoder(w).Encode(people)
	}

func main(){
	router:= mux.NewRouter()
	fmt.Println("Server is runnning on port 8080\n")
	people=append(people,Person{Id:"121443",Name:"Prashanth",Bal:"1000"})
	people=append(people,Person{Id:"121560",Name:"Rambabu",Bal:"500"})
	router.HandleFunc("/adduser",addUser).Methods("POST")
	router.HandleFunc("/getbal/{id}",getBalance).Methods("GET")
	router.HandleFunc("/getbal",getBalance).Methods("GET")
	router.HandleFunc("/deposit/{id}",deposit).Methods("POST")
	router.HandleFunc("/withdraw/{id}",withDraw).Methods("POST")
	//	cdrouter.HandleFunc("/withdraw/{id}",getRemainBalance).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080",router))
	}
