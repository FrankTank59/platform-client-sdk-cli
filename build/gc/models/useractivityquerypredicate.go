package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UseractivityquerypredicateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UseractivityquerypredicateDud struct { 
    


    


    


    

}

// Useractivityquerypredicate
type Useractivityquerypredicate struct { 
    // VarType - Optional type, can usually be inferred
    VarType string `json:"type"`


    // Dimension - Left hand side for dimension predicates
    Dimension string `json:"dimension"`


    // Operator - Optional operator, default is matches
    Operator string `json:"operator"`


    // Value - Right hand side for dimension predicates
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Useractivityquerypredicate) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Useractivityquerypredicate) MarshalJSON() ([]byte, error) {
    type Alias Useractivityquerypredicate

    if UseractivityquerypredicateMarshalled {
        return []byte("{}"), nil
    }
    UseractivityquerypredicateMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Dimension string `json:"dimension"`
        
        Operator string `json:"operator"`
        
        Value string `json:"value"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

