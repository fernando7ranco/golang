package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"
)

type teste struct {
	name     string `json:"name"`
	lastname string `json:"lastname"`
	age      int8   `json:"age"`
}

type teste2 struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Age      int8   `json:"age"`
}

type person struct {
	Name structName `json:"name"`
	Age  int8       `json:"age"`
}

type structName struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

type student struct {
	person  `json:"person"`
	Subject string `json:"subject"`
}

func (s student) printName() {
	fmt.Println(s.Name)
}

func (s *student) setName(first string, last string) {
	s.Name.First = first
	s.Name.Last = last
}

func main() {

	jsonTest()
}

func jsonTest() {

	t := teste2{"nome", "Oi", 121}
	tJson, err := json.Marshal(t)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(tJson)
	fmt.Println(string(tJson))

	s := student{
		person: person{
			Name: structName{
				First: "fernando",
				Last:  "franco",
			},
			Age: 26,
		},
		Subject: "phisic",
	}
	fmt.Println(s)

	studentJson, err := json.Marshal(s)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(studentJson)
	fmt.Println(string(studentJson))

	teste := map[string]string{
		"nome":  "oi",
		"teste": "Oi",
	}

	testJson, err := json.Marshal(teste)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(testJson)
	fmt.Println(bytes.NewBuffer(testJson))
}

func goRutines() {
	var waitgroup sync.WaitGroup
	waitgroup.Add(2)

	loop := func(variante string, seconds int8) {
		for seconds > 0 {
			fmt.Println("loop", variante)
			time.Sleep(time.Second / 2)
			seconds--
		}
		fmt.Println("fim loop ", variante)
		waitgroup.Done()
	}

	go loop("fisrt", 2)
	go loop("secund", 3)

	waitgroup.Wait()

	chanel := make(chan string)

	loopWithChanel := func(variante string, seconds int8, chanel chan string) {
		for seconds > 0 {
			chanel <- "loop with chanel " + variante
			time.Sleep(time.Second / time.Duration(seconds) * 2)
			seconds--
		}
		chanel <- "fim loop " + variante
		close(chanel)
	}

	go loopWithChanel("1", 2, chanel)

	for message := range chanel {
		fmt.Println(message)
	}

	chanels1, chanels2 := make(chan string), make(chan string)

	go loopWithChanel(" -> 1", 20, chanels1)
	go loopWithChanel("-> 2", 15, chanels2)

	for {
		select {
		case m1 := <-chanels1:
			fmt.Println(m1)
		case m2 := <-chanels2:
			fmt.Println(m2)
		}
	}

}

func arrays() {
	var v1 []int
	fmt.Println(v1)
	fmt.Println(len(v1))
	fmt.Println(cap(v1))

	v2 := [4]int{1}
	fmt.Println(v2)
	fmt.Println(len(v2))
	fmt.Println(cap(v2))

	v3 := make([]int, 2, 4)
	fmt.Println(v3)
	fmt.Println(len(v3))
	fmt.Println(cap(v3))
}

func useReferences() {
	value := 10
	value2 := value

	fmt.Println(value, value2)
	value++
	fmt.Println(value, value2)

	// use point
	var pvalue int
	var pvalue2 *int
	fmt.Println(pvalue, pvalue2)

	pvalue = 22
	pvalue2 = &pvalue
	fmt.Println(pvalue, *pvalue2)
	pvalue++
	fmt.Println(pvalue, *pvalue2)
}

func aletoryCode() {
	var (
		var1 string
		var2 string
		var3 int16
	)
	fmt.Println("vars 1 to 3", var1, var2, var3)

	a := func() string {
		return "goood"
	}

	fmt.Println(a())
}

func getName() (first string, last string) {
	first = "fernando"
	last = "franco"
	return
}

func useStruct() {

	var t teste
	t.name = "fernando"
	t.lastname = "franco"
	t.age = 22
	fmt.Println(t)

	tFast := teste{name: "fernando", lastname: "franco", age: 22}
	fmt.Println(tFast)

	p := person{
		Name: structName{
			First: "fernando",
			Last:  "franco",
		},
		Age: 26,
	}
	fmt.Println(p)

	s := student{
		person: person{
			Name: structName{
				First: "fernando",
				Last:  "franco",
			},
			Age: 26,
		},
		Subject: "phisic",
	}
	fmt.Println(s)
	fmt.Println(s.Name.First)

	s.printName()

	s.setName("testinho", "delteste")
	s.printName()
}

func useMaps() {
	usuario := map[string]string{
		"name":     "fernando",
		"lastname": "franco",
	}

	strutura := map[string]map[string]string{
		"name": {
			"fist": "fernando",
			"last": "franco",
		},
	}

	defer fmt.Println(usuario["name"])
	fmt.Println(strutura["name"]["last"])

	for key, value := range usuario {
		fmt.Printf("key = %v value = %v \n", key, value)
	}

	first, last := getName()
	fmt.Println(first, last)

}
