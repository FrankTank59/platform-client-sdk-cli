package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgedocumentversionvariationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgedocumentversionvariationDud struct { 
    Id string `json:"id"`


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    


    SelfUri string `json:"selfUri"`


    DocumentVersion Addressableentityref `json:"documentVersion"`

}

// Knowledgedocumentversionvariation
type Knowledgedocumentversionvariation struct { 
    


    // Body - The content for the variation.
    Body Documentbody `json:"body"`


    


    


    // Contexts - The context values associated with the variation.
    Contexts []Documentvariationcontext `json:"contexts"`


    


    

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentversionvariation) String() string {
    
     o.Contexts = []Documentvariationcontext{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgedocumentversionvariation) MarshalJSON() ([]byte, error) {
    type Alias Knowledgedocumentversionvariation

    if KnowledgedocumentversionvariationMarshalled {
        return []byte("{}"), nil
    }
    KnowledgedocumentversionvariationMarshalled = true

    return json.Marshal(&struct {
        
        Body Documentbody `json:"body"`
        
        Contexts []Documentvariationcontext `json:"contexts"`
        *Alias
    }{

        


        


        


        


        
        Contexts: []Documentvariationcontext{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

