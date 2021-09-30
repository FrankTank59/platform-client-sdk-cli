package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WritabledivisionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WritabledivisionDud struct { 
    


    


    SelfUri string `json:"selfUri"`

}

// Writabledivision
type Writabledivision struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    

}

// String returns a JSON representation of the model
func (o *Writabledivision) String() string {
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Writabledivision) MarshalJSON() ([]byte, error) {
    type Alias Writabledivision

    if WritabledivisionMarshalled {
        return []byte("{}"), nil
    }
    WritabledivisionMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

