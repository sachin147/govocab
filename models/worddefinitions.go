package models


type Definition struct {	
	ExtendedText string `json:"extendedText,omitempty"`
	Text string `json:"text,omitempty"`
	SourceDictionary string `json:"sourceDictionary,omitempty"`
	Citations struct {
		Citation []struct {
			Cite string `json:"cite,omitempty"`
			Source string `json:"source,omitempty"`
		} 
	} `json:"citations,omitempty"`
	Labels struct {
		Label []struct {
			Text string `json:"text,omitempty"`
			Type string `json:"type,omitempty"`
		}
	} `json:"labels,omitempty"`
	Score float64 `json:"score,omitempty"`
	ExampleUses struct {
		ExampleUse []struct {
			Text string `json:"text,omitempty"`
		}
	} `json:"exampleUses,omitempty"`
	AttributionUrl string `json:"attributionUrl,omitempty"`
	SeqString string `json:"seqString,omitempty"`
	AttributionText string `json:"attributionText,omitempty"`
	RelatedWords struct {
		RelatedWord []struct {
			Label1 string `json:"label1,omitempty"`
			RelationshipType string `json:"relationshipType,omitempty"`
			Label2 string `json:"label2,omitempty"`
			Label3 string `json:"label3,omitempty"`
			Words []string `json:"words,omitempty"`
			Gram string `json:"gram,omitempty"`
			Label4 string `json:"label4,omitempty"`
		}
	} `json:"relatedWords,omitempty"`
	Sequence string `json:"sequence,omitempty"`
	Word string `json:"word,omitempty"`
	Notes struct {
		Note []struct {
			Notetype string `json:"noteType,omitempty"`
			AppliesTo []string `json:"appliesTo,omitempty"`
			Value string `json:"value,omitempty"`
			Pos int `json:"pos,omitempty"`
		}
	} `json:"notes,omitempty"`
	TextProns struct {
		TextPron []struct {
			Raw string `json:"raw,omitempty"`
			Seq int `json:"seq,omitempty"`
			RawType string `json:"string,omitempty"`
		}
	} `json:"textProns,omitempty"`
	PartOfSpeech string `json:"partOfSpeech,omitempty"`	
}
