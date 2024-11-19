package services

import (
	"fmt"
	. "servidorhttp/app/structs"
)

type ServiceHome struct{}

func (s ServiceHome) Welcome() Welcome {
	w := Welcome{Message: "Welcome"}
	fmt.Println(w.Message)
	return w
}
