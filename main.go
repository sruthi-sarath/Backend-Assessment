package main

import (
	"fmt"
	"net/http"
	"time"
	"encoding/json"

)

type Request struct {
	Numbers []int `json:"numbers"`     //this will only contain integers
	Target  int   `json:"target"`	   
}

type Response struct {
	Solutions [][]int `json:"solutions"`
}

func main(){

url:="http://localhost:8080"          //local url to run the code
fmt.Println(url)
	go func() {
		http.HandleFunc("/find-pairs", handler)        //calling api
		errhttp := http.ListenAndServe(":8080", nil)   //listening on port 8080
		if errhttp != nil {
			fmt.Println("Failed to start server")
		}
	}()

	time.Sleep(30 * time.Second)				//wait time given is 30sec
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || len(req.Numbers) == 0 {       // checking if there is any error, and length is 0
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	
	solutions := sumofvalue(req.Numbers, req.Target)    //calling the func
	fmt.Println("solutions",solutions)              
	res := Response{Solutions: solutions}               //output of func is saved in solution

	w.Header().Set("Content-Type", "application/json")   
	json.NewEncoder(w).Encode(res)
	
}

func sumofvalue(value []int,target int)[][]int{
	empty:=[][]int{}

	length:=len(value)
	fmt.Println("length",length)

	if length>0{
		for i:=0;i<length;i++{
			for j:=i+1;j<length;j++{
				
				if value[i]!=value[j]{      // checking if there is duplicate value. If it is then its skip that value
				if value[i]+value[j]==target{     //  checking if addition of the values is same as target
				
					empty=append(empty,[]int{i,j})    //saving the value in empty if 2 values=target
				}
			}
			}
		}
		
	}
	return  empty     //if no value is same it will send empty
}

