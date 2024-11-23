package structs

import (
	"encoding/json"
	"fmt"
	"log"
)

type HelperJson struct{}

func (h HelperJson) ParseJson(childs interface{}) []byte {
	bytesJson, err := json.Marshal(&childs)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(string(bytesJson))
	return bytesJson
}
