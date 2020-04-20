package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func updateSlice(varSlice []string,index int, value string)  {
	varSlice[index] = value
}

func sliceQue()  {
	varSlice:=[]string{"a","b","c","d","e"}
	varSliceCpy:=[]string{}
	fmt.Println("slice value = ",varSlice)
	fmt.Println("Slice length = ",len(varSlice))
	fmt.Println("After Update")
	updateSlice(varSlice,4,"e1")
	fmt.Println("slice value = ",varSlice)
	fmt.Println("After Update 2")
	varSliceCpy = varSlice
	varSlice[4] = "e2"
	fmt.Println("slice value = ",varSlice)
	fmt.Println("Copy slice value = ",varSliceCpy)
}

func mapQue()  {
	student := map[int]string{1:"Rohan Dube",2:"Mayuri Pathare",3:"Devendra Kulkarni",4:"Yogesh Pendse",5:"abc xyz"}
	for key,value := range student{
		fmt.Println("Roll number = ",key,"Full name = ",value)
	}
	fmt.Println("After update")
	student[1] = "John Cena"
	fmt.Println("New map = ",student)
	fmt.Println("After Deleting failed students")
	delete(student,5)
	fmt.Println("New map = ",student)
}

func pingSite(url string, pt bool)int  {
	resp, err :=http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	if pt {
		respBody, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("Response body = ", string(respBody))
	}
	return resp.StatusCode
}

func main() {
	defer sliceQue()
	defer fmt.Println("Slice**********")
	defer fmt.Println()
	defer mapQue()
	defer fmt.Println("Map**********")

	fmt.Println("Assignments Day3")
	fmt.Println("HTTP********")
	responseCode := pingSite("https://www.google.com/",false) // Make pt true for printing body
	fmt.Println("Response code = ", responseCode)
	fmt.Println()

}
/*
Assignments Day3
HTTP********
Response code =  200

Map**********
Roll number =  2 Full name =  Mayuri Pathare
Roll number =  3 Full name =  Devendra Kulkarni
Roll number =  4 Full name =  Yogesh Pendse
Roll number =  5 Full name =  abc xyz
Roll number =  1 Full name =  Rohan Dube
After update
New map =  map[1:John Cena 2:Mayuri Pathare 3:Devendra Kulkarni 4:Yogesh Pendse 5:abc xyz]
After Deleting failed students
New map =  map[1:John Cena 2:Mayuri Pathare 3:Devendra Kulkarni 4:Yogesh Pendse]

Slice**********
slice value =  [a b c d e]
Slice length =  5
After Update
slice value =  [a b c d e1]
After Update 2
slice value =  [a b c d e2]
Copy slice value =  [a b c d e2]

 */