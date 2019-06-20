package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type person struct {
	Id   int
	Name string
	Age  int
}

func helloGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GET \n\n")
	var list []person = make([]person, 10)
	list[0] = person{Id: 123, Name: "Anh", Age: 21}
	list[1] = person{Id: 124, Name: "Nam", Age: 20}
	list[2] = person{Id: 125, Name: "Thuong", Age: 19}
	list[3] = person{Id: 126, Name: "Oanh", Age: 21}
	fmt.Fprintf(w, "Xin chao ban %s", list[0].Name)
}

type dataFjson struct {
	Count int
}

func helloPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "POST \n")
	var list []person = make([]person, 10)
	list[0] = person{Id: 123, Name: "Anh", Age: 21}
	list[1] = person{Id: 124, Name: "Nam", Age: 20}
	list[2] = person{Id: 125, Name: "Thuong", Age: 19}
	list[3] = person{Id: 126, Name: "Oanh", Age: 21}
	fmt.Fprintf(w, "Xin chao ban.\n ")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error reading body data: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	//fmt.Fprintf(w, "Body data:: %v", string(body))

	var dat dataFjson

	err1 := json.Unmarshal(body, &dat)
	if err1 != nil {
		log.Println(err1)
	}
	fmt.Fprintln(w, "\n", dat.Count, " người đầu tiên trong danh sách:")
	for i := 0; i < dat.Count; i++ {
		fmt.Fprintf(w, "ID: %d, Name: %s, Age: %d\n", list[i].Id, list[i].Name, list[i].Age)
	}
}

func helloPut(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PUT\n")
	var list []person = make([]person, 10)
	list[0] = person{Id: 123, Name: "Anh", Age: 21}
	list[1] = person{Id: 124, Name: "Nam", Age: 20}
	list[2] = person{Id: 125, Name: "Thuong", Age: 19}
	list[3] = person{Id: 126, Name: "Oanh", Age: 21}

	fmt.Fprintf(w, "Xin chao ban. Du lieu ban dau\n")
	for i := 0; i < len(list); i++ {
		if (list[i].Id) != 0 {
			fmt.Fprintf(w, "ID: %d, Name: %s, Age: %d\n", list[i].Id, list[i].Name, list[i].Age)
		}
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error reading body data: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	var dat person
	err1 := json.Unmarshal(body, &dat)
	if err1 != nil {
		log.Println(err1)
	}
	//fmt.Fprintf(w, "\nID: %d, Name: %s, Age: %d\n", dat.Id, dat.Name, dat.Age)
	var check bool
	check = true
	for i := 0; i < len(list); i++ {
		if list[i].Id == dat.Id {
			check = false
			fmt.Fprintln(w, "Id cua ban da ton tai")
		}
	}
	if check == true {
		list = append(list, dat)
	}

	fmt.Fprintf(w, "Du lieu sau khi update\n")
	for i := 0; i < len(list); i++ {
		if (list[i].Id) != 0 {
			fmt.Fprintf(w, "ID: %d, Name: %s, Age: %d\n", list[i].Id, list[i].Name, list[i].Age)
		}
	}
}

func helloDelete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DELETE \n\n")
	var list []person = make([]person, 10)
	list[0] = person{Id: 123, Name: "Anh", Age: 21}
	list[1] = person{Id: 124, Name: "Nam", Age: 20}
	list[2] = person{Id: 125, Name: "Thuong", Age: 19}
	list[3] = person{Id: 126, Name: "Oanh", Age: 21}

	fmt.Fprintf(w, "Xin chao ban. Du lieu ban dau\n")
	for i := 0; i < len(list); i++ {
		if (list[i].Id) != 0 {
			fmt.Fprintf(w, "ID: %d, Name: %s, Age: %d\n", list[i].Id, list[i].Name, list[i].Age)
		}
	}
	fmt.Fprintf(w, "Xoa du lieu có id = 123 ")
	for i := 0; i < len(list); i++ {
		if list[i].Id == 123 {
			copy(list[i:], list[i+1:])
			list = list[:len(list)-1]
		}
	}
	fmt.Fprintf(w, "Du lieu sau khi xoa\n")
	for i := 0; i < len(list); i++ {
		if (list[i].Id) != 0 {
			fmt.Fprintf(w, "ID: %d, Name: %s, Age: %d\n", list[i].Id, list[i].Name, list[i].Age)
		}
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "This is a website server by a Go HTTP server.")
	})

	http.HandleFunc("/helloGet", helloGet)
	http.HandleFunc("/helloPost", helloPost)
	http.HandleFunc("/helloPut", helloPut)
	http.HandleFunc("/helloDelete", helloDelete)

	fmt.Println(":3001")
	http.ListenAndServe(":3001", nil)
}
