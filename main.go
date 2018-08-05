package main

import (
	"fmt"

	"github.com/pankona/di-sandbox/glue"
	"github.com/pankona/di-sandbox/repository"
	"github.com/pankona/di-sandbox/service"
)

func main() {
	r := repository.New()
	u := service.New(glue.ServiceRepoer(r))

	err := u.CreateNewUser()
	if err != nil {
		fmt.Println(err)
	}
}
