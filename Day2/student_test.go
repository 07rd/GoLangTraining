package main

import "testing"

func TestStudentInfo(t *testing.T)  {
	s:= Student{"Rohan","Dube",25,ContactInfo{"0987654321","abc@gmail.com"}}
	if s.age < 20{
		t.Errorf("Not eligible")
	}
	if s.email ==""{
		t.Errorf("EmailId not found")
	}
}