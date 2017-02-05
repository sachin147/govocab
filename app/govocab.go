package main

import (
	"fmt"
        "net/http"
	"flag"
	"io/ioutil"
	"encoding/json"
	"github.com/sachin147/govocab/constants"
	"github.com/sachin147/govocab/models"
	"github.com/sachin147/keys"
)

func main() {
	
	wordexample := flag.String("wordexamples", "", "Get examples of a word")
	worddefinition := flag.String("worddefinitions", "", "Get definitions of a word")
	flag.Parse()
	word := *wordexample
	
	url := constants.API_URL + constants.WORD_JSON + word + constants.EXAMPLES_PATH + constants.QUERY_PATH + keys.WORDNIK_API_KEY
	definitionurl := constants.API_URL + constants.WORD_JSON + wordfordefinition + constants.DEFINITIONS_PATH + constants.QUERY_PATH_DEFINITION + keys.WORDNIK_API_KEY

	 
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
		
	
	wordfordefinition := *worddefinition
	definitionurl := constants.API_URL + constants.WORD_JSON + wordfordefinition + constants.DEFINITIONS_PATH + constants.QUERY_PATH_DEFINITION + keys.WORDNIK_API_KEY

	res, err := http.Get(definitionurl)
	
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

	
	worddefinitions, err := parseWordDefinitions(body)

	

}

func ParseWordExamples(wordexamplesbody []byte) (*models.WordExamples, error) {
    var wordexamples models.WordExamples
    err := json.Unmarshal(wordexamplesbody, &wordexamples)
    if(err != nil){
        fmt.Println("Error ", err)
    }
    return &wordexamples, err
}


parseWordDefinitions(worddefinitionsbody []byte)(*models.Definition, error) {
	
	
}




