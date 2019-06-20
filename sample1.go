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

type Message struct {
	Status string
}

func helloGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		m := Message{"404 Not found"}
		b, _ := json.Marshal(m)
		fmt.Fprintf(w, string(b))
		return
	}

	var list []person = make([]person, 4, 10)

	list[0] = person{123, "Anh", 21}
	list[1] = person{Id: 124, Name: "Nam", Age: 20}
	list[2] = person{Id: 125, Name: "Thuong", Age: 19}
	list[3] = person{Id: 126, Name: "Oanh", Age: 21}

	if list[0].Id == 0 {
		m := Message{"404 Not found"}
		b, _ := json.Marshal(m)
		fmt.Fprintf(w, string(b))
		return
	}

	res2, _ := json.Marshal(list[0])
	fmt.Fprintln(w, string(res2))

}

type dataFjson struct {
	Count int
}

func helloPost(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		m := Message{"404 Not found"}
		b, _ := json.Marshal(m)
		fmt.Fprintf(w, string(b))
		return
	}

	var list []person = make([]person, 4, 10)
	list[0] = person{Id: 123, Name: "Anh", Age: 21}
	list[1] = person{Id: 124, Name: "Nam", Age: 20}
	list[2] = person{Id: 125, Name: "Thuong", Age: 19}
	list[3] = person{Id: 126, Name: "Oanh", Age: 21}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error reading body data: %v", err)
		http.Error(w, "{'code':'error'}", http.StatusBadRequest)
		return
	}

	var dat dataFjson

	err1 := json.Unmarshal(body, &dat)
	if err1 != nil {
		log.Println(err1)
	}
	// len(lst) == 0
	//fmt.Fprintln(w, "\n", dat.Count, " người đầu tiên trong danh sách:")
	if len(list) == 0 {
		m := Message{"404 Not found"}
		b, _ := json.Marshal(m)
		fmt.Fprintf(w, string(b))
		return
	}

	if dat.Count < 0 || dat.Count > 10 {
		m := Message{"400 Bad request"}
		b, _ := json.Marshal(m)
		fmt.Fprintf(w, string(b))
		return
	}

	res2, _ := json.Marshal(list[:dat.Count])
	fmt.Fprintln(w, string(res2))

}

func helloPut(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PUT" {
		m := Message{"404 Not found"}
		b, _ := json.Marshal(m)
		fmt.Fprintf(w, string(b))
		return
	}

	var list []person = make([]person, 4, 10)
	list[0] = person{Id: 123, Name: "Anh", Age: 21}
	list[1] = person{Id: 124, Name: "Nam", Age: 20}
	list[2] = person{Id: 125, Name: "Thuong", Age: 19}
	list[3] = person{Id: 126, Name: "Oanh", Age: 21}

	var list1Cpy []person = make([]person, 4, 10)
	list1Cpy = list
	//Không có dữ liệu trong list
	if list[0].Id == 0 {
		m := Message{"404 Not found"}
		b, _ := json.Marshal(m)
		fmt.Fprintf(w, string(b))
		return
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
			m := Message{"406 Not Accept. Id have already in system"}
			b, _ := json.Marshal(m)
			fmt.Fprintf(w, string(b))
			return
		}
	}
	if check == true {
		list = append(list, dat)
	}

	var listOfList [][]person = make([][]person, 2, 10)
	listOfList[0] = list1Cpy
	listOfList[1] = list
	res2, _ := json.Marshal(listOfList)
	fmt.Fprintln(w, string(res2))

}

func helloDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		m := Message{"404 Not found"}
		b, _ := json.Marshal(m)
		fmt.Fprintf(w, string(b))
		return
	}

	var list []person = make([]person, 4, 10)
	list[0] = person{Id: 123, Name: "Anh", Age: 21}
	list[1] = person{Id: 124, Name: "Nam", Age: 20}
	list[2] = person{Id: 125, Name: "Thuong", Age: 19}
	list[3] = person{Id: 126, Name: "Oanh", Age: 21}

	//fmt.Fprintf(w, "Xin chao ban. Du lieu ban dau\n")
	if list[0].Id == 0 {
		m := Message{"404 Not found"}
		b, _ := json.Marshal(m)
		fmt.Fprintf(w, string(b))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error reading body data: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	var dat dataFjson

	err1 := json.Unmarshal(body, &dat)
	if err1 != nil {
		log.Println(err1)
	}

	//Kiểm tra ID có tồn tại k
	check := false
	for i := 0; i < len(list); i++ {
		if list[i].Id == dat.Count {
			check = true
			break

		}
	}

	if check == false {
		m := Message{"406 Not Accept. Your ID not in system"}
		b, _ := json.Marshal(m)
		fmt.Fprintf(w, string(b))
		return
	}
	//fmt.Fprintf(w, "Xoa du lieu có id = 123 ")
	for i := 0; i < len(list); i++ {
		if list[i].Id == dat.Count {
			copy(list[i:], list[i+1:])
			list = list[:len(list)-1]
		}
	}

	res2, _ := json.Marshal(list)
	fmt.Fprintln(w, string(res2))

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

//Change nhé!
