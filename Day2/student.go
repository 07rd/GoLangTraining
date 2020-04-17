package main

import (
    "fmt"
    "io/ioutil"
    "log"
)

type ContactInfo struct {
    phoneNo string
    email string
}
type Student struct {
    firstName string
    lastName string
    age int
    ContactInfo

}
func (s *Student) updateMail(newMail string)  {
    s.email = newMail
}
func (s Student) print()  {
    fmt.Printf("%+v\n",s)
}
func (s Student) saveToFile(filename string,str string) {
    err :=ioutil.WriteFile(filename,[]byte(str),0644)
    if err != nil{
        log.Fatal(err)
    }
}
func toString(s Student) string  {
    return fmt.Sprintf("%+v\n",s)
}
func main(){
    s:= Student{firstName: "Rohan", lastName: "Dube",age: 25,ContactInfo:ContactInfo{phoneNo: "0987654321",email: "xyz@gamil.com"}}
    s.print()
    fmt.Println("After update")
    s.updateMail("abc@gmail.com")
    s.print()
    s.saveToFile("studentInfo",toString(s))
}
