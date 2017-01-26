package main

import (
	"fmt"
        "net/http"
	"flag"
	"io/ioutil"
	"encoding/json"
	"github.com/sachin147/constants"
	"github.com/sachin147/keys"
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
	
	wordexample := flag.String("wordexamples", "", "Get examples of a word")
	flag.Parse()
	word := *wordexample
	url := constants.API_URL + constants.WORD_JSON + word + constants.EXAMPLES_PATH + constants.QUERY_PATH + keys.WORDNIK_API_KEY
	res, err := http.Get(url)
	
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







