package main

import (
	"fmt"
        "net/http"
	"io/ioutil"
	"encoding/json"
)

type Provider struct {
	Name string `json:"name"`
	Id int `json:"id"`
}
 
type WordExample struct {
	Year int `json:"year"`
	Provider Provider `json:"provider"`
	Url string `json:"url"`
	Word string `json:"word"`
	Text string `json:"text"`
	Title string `json:"title"`
	DocumentId int `json:"documentId"`
	ExampleId int `json:"exampleId"`
	Rating float64 `json:"rating"`         
}

type WordExamples struct {
	Examples []WordExample `json:"examples"`	
}

func main() {
	
	res, err := http.Get("http://api.wordnik.com:80/v4/word.json/wonderful/examples?includeDuplicates=false&useCanonical=false&skip=0&limit=5&api_key=a2a73e7b926c924fad7001ca3111acd55af2ffabf50eb4ae5")
	
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()
	
	if err != nil {
           panic(err.Error())
        }
	
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
           panic(err.Error())
        }

	
	wordexamples, err := ParseWordExamples(body)

	for _, example := range wordexamples.Examples {
		fmt.Println(example.Text)
		fmt.Println("=======================")
	}
		
}

func ParseWordExamples(wordexamplesbody []byte) (*WordExamples, error) {
    var wordexamples WordExamples
    err := json.Unmarshal(wordexamplesbody, &wordexamples)
    if(err != nil){
        fmt.Println("Error ", err)
    }
    return &wordexamples, err
}







