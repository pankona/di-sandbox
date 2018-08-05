package repository

import "fmt"

type Repo struct{}

func New() *Repo {
	return &Repo{}
}

func (r *Repo) Create() error {
	fmt.Println("Create!")
	return nil
}
