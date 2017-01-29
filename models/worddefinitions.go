package models


type Definition struct {	
	ExtendedText string `json:"extendedText"`
	Text string `json:"text"`
	SourceDictionary string `json:"sourceDictionary"`
	Citations struct {
		Citation []struct {
			Cite string `json:"cite"`
			Source string `json:"source"`
		}
	}
	Labels struct {
		Label []struct {
			Text string `json:"text"`
			Type string `json:"type"`
		}
	}
	Score float64 `json:"score"`
	ExampleUses struct {
		ExampleUse []struct {
			Text string `json:"text"`
		}
	}
	AttributionUrl string `json:"attributionUrl"`
	SeqString string `json:"seqString"`
	AttributionText string `json:"attributionText"`
	RelatedWords struct {
		RelatedWord []struct {
			Label1 string `json:"label1"`
			RelationshipType string `json:"relationshipType"`
			Label2 string `json:"label2"`
			Label3 string `json:"label3"`
			Words []string `json:"words"`
			Gram string `json:"gram"`
			Label4 string `json:"label4"`
		}
	}
}
