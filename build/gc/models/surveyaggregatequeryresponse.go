package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SurveyaggregatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SurveyaggregatequeryresponseDud struct { 
    

}

// Surveyaggregatequeryresponse
type Surveyaggregatequeryresponse struct { 
    // Results
    Results []Surveyaggregatedatacontainer `json:"results"`

}

// String returns a JSON representation of the model
func (o *Surveyaggregatequeryresponse) String() string {
    
    
     o.Results = []Surveyaggregatedatacontainer{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Surveyaggregatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Surveyaggregatequeryresponse

    if SurveyaggregatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    SurveyaggregatequeryresponseMarshalled = true

    return json.Marshal(&struct { 
        Results []Surveyaggregatedatacontainer `json:"results"`
        
        *Alias
    }{
        

        
        Results: []Surveyaggregatedatacontainer{{}},
        

        
        Alias: (*Alias)(u),
    })
}

