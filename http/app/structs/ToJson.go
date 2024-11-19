package structes

import (
	"encoding/json"
	"fmt"
	"log"
	intefaces "servidorhttp/app/interfaces"
)

type HelperJson struct{}

func (h HelperJson) ToJson(child intefaces.IStructToJson) []byte {
	bytesJson, err := json.Marshal(child)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(string(bytesJson))
	return bytesJson
}
