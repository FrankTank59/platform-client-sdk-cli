package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FaxdocumententitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FaxdocumententitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Faxdocumententitylisting
type Faxdocumententitylisting struct { 
    // Entities
    Entities []Faxdocument `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // NextUri
    NextUri string `json:"nextUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // LastUri
    LastUri string `json:"lastUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Faxdocumententitylisting) String() string {
     o.Entities = []Faxdocument{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Faxdocumententitylisting) MarshalJSON() ([]byte, error) {
    type Alias Faxdocumententitylisting

    if FaxdocumententitylistingMarshalled {
        return []byte("{}"), nil
    }
    FaxdocumententitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Faxdocument `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        FirstUri string `json:"firstUri"`
        
        NextUri string `json:"nextUri"`
        
        PreviousUri string `json:"previousUri"`
        
        SelfUri string `json:"selfUri"`
        
        LastUri string `json:"lastUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Faxdocument{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

