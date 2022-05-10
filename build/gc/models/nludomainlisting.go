package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NludomainlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NludomainlistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Nludomainlisting
type Nludomainlisting struct { 
    // Entities
    Entities []Nludomain `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // LastUri
    LastUri string `json:"lastUri"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // NextUri
    NextUri string `json:"nextUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Nludomainlisting) String() string {
    
    
     o.Entities = []Nludomain{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Nludomainlisting) MarshalJSON() ([]byte, error) {
    type Alias Nludomainlisting

    if NludomainlistingMarshalled {
        return []byte("{}"), nil
    }
    NludomainlistingMarshalled = true

    return json.Marshal(&struct { 
        Entities []Nludomain `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        LastUri string `json:"lastUri"`
        
        FirstUri string `json:"firstUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        NextUri string `json:"nextUri"`
        
        PageCount int `json:"pageCount"`
        
        *Alias
    }{
        

        
        Entities: []Nludomain{{}},
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

