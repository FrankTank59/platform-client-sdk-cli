package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TeamactivityqueryfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TeamactivityqueryfilterDud struct { 
    


    


    

}

// Teamactivityqueryfilter
type Teamactivityqueryfilter struct { 
    // VarType - Boolean operation to apply to the provided predicates and clauses
    VarType string `json:"type"`


    // Clauses - Boolean 'and/or' logic with up to two-levels of nesting
    Clauses []Teamactivityqueryclause `json:"clauses"`


    // Predicates - Like a three-word sentence: (attribute-name) (operator) (target-value).
    Predicates []Teamactivityquerypredicate `json:"predicates"`

}

// String returns a JSON representation of the model
func (o *Teamactivityqueryfilter) String() string {
    
     o.Clauses = []Teamactivityqueryclause{{}} 
     o.Predicates = []Teamactivityquerypredicate{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Teamactivityqueryfilter) MarshalJSON() ([]byte, error) {
    type Alias Teamactivityqueryfilter

    if TeamactivityqueryfilterMarshalled {
        return []byte("{}"), nil
    }
    TeamactivityqueryfilterMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Clauses []Teamactivityqueryclause `json:"clauses"`
        
        Predicates []Teamactivityquerypredicate `json:"predicates"`
        *Alias
    }{

        


        
        Clauses: []Teamactivityqueryclause{{}},
        


        
        Predicates: []Teamactivityquerypredicate{{}},
        

        Alias: (*Alias)(u),
    })
}

