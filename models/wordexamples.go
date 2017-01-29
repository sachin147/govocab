package models

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
