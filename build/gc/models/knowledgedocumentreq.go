package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgedocumentreqMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgedocumentreqDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Knowledgedocumentreq
type Knowledgedocumentreq struct { 
    


    // Title - Document title.
    Title string `json:"title"`


    // Visible - Indicates if the knowledge document should be included in search results.
    Visible bool `json:"visible"`


    // Alternatives - List of alternate phrases related to the title which improves search results.
    Alternatives []Knowledgedocumentalternative `json:"alternatives"`


    // CategoryId - The category associated with the document.
    CategoryId string `json:"categoryId"`


    // LabelIds - The ids of labels associated with the document.
    LabelIds []string `json:"labelIds"`


    

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentreq) String() string {
    
    
     o.Alternatives = []Knowledgedocumentalternative{{}} 
    
     o.LabelIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgedocumentreq) MarshalJSON() ([]byte, error) {
    type Alias Knowledgedocumentreq

    if KnowledgedocumentreqMarshalled {
        return []byte("{}"), nil
    }
    KnowledgedocumentreqMarshalled = true

    return json.Marshal(&struct {
        
        Title string `json:"title"`
        
        Visible bool `json:"visible"`
        
        Alternatives []Knowledgedocumentalternative `json:"alternatives"`
        
        CategoryId string `json:"categoryId"`
        
        LabelIds []string `json:"labelIds"`
        *Alias
    }{

        


        


        


        
        Alternatives: []Knowledgedocumentalternative{{}},
        


        


        
        LabelIds: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

