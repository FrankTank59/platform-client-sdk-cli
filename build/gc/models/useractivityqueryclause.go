package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UseractivityqueryclauseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UseractivityqueryclauseDud struct { 
    


    

}

// Useractivityqueryclause
type Useractivityqueryclause struct { 
    // VarType - Boolean operation to apply to the provided predicates
    VarType string `json:"type"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Useractivityquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Useractivityqueryclause) String() string {
    
     o.Predicates = []Useractivityquerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Useractivityqueryclause) MarshalJSON() ([]byte, error) {
    type Alias Useractivityqueryclause

    if UseractivityqueryclauseMarshalled {
        return []byte("{}"), nil
    }
    UseractivityqueryclauseMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Predicates []Useractivityquerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Predicates: []Useractivityquerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

