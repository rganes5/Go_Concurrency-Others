package main

import "fmt"

func main() {
	//m := mongo{}
	p := postgres{}
	u := Usecase{
		p,
	}
	u.FindUser()
}

type mongo struct{}

func (m mongo) Find() {
	fmt.Println("Found from mongo")
}

type postgres struct{}

func (p postgres) Find() {
	fmt.Println("Found from postgres")
}

type DB interface {
	Find()
}

type Usecase struct {
	DB
}

func (u Usecase) FindUser() {
	u.Find()
}
