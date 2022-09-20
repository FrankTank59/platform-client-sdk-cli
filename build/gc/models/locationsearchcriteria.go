package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LocationsearchcriteriaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LocationsearchcriteriaDud struct { 
    


    


    


    


    


    


    


    


    

}

// Locationsearchcriteria
type Locationsearchcriteria struct { 
    // EndValue - The end value of the range. This field is used for range search types.
    EndValue string `json:"endValue"`


    // Values - A list of values for the search to match against
    Values []string `json:"values"`


    // StartValue - The start value of the range. This field is used for range search types.
    StartValue string `json:"startValue"`


    // Value - A value for the search to match against
    Value string `json:"value"`


    // Operator - How to apply this search criteria against other criteria
    Operator string `json:"operator"`


    // Group - Groups multiple conditions
    Group []Locationsearchcriteria `json:"group"`


    // DateFormat - Set date format for criteria values when using date range search type.  Supports Java date format syntax, example yyyy-MM-dd'T'HH:mm:ss.SSSX.
    DateFormat string `json:"dateFormat"`


    // VarType - Search Type
    VarType string `json:"type"`


    // Fields - Field names to search against
    Fields []string `json:"fields"`

}

// String returns a JSON representation of the model
func (o *Locationsearchcriteria) String() string {
    
     o.Values = []string{""} 
    
    
    
     o.Group = []Locationsearchcriteria{{}} 
    
    
     o.Fields = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Locationsearchcriteria) MarshalJSON() ([]byte, error) {
    type Alias Locationsearchcriteria

    if LocationsearchcriteriaMarshalled {
        return []byte("{}"), nil
    }
    LocationsearchcriteriaMarshalled = true

    return json.Marshal(&struct {
        
        EndValue string `json:"endValue"`
        
        Values []string `json:"values"`
        
        StartValue string `json:"startValue"`
        
        Value string `json:"value"`
        
        Operator string `json:"operator"`
        
        Group []Locationsearchcriteria `json:"group"`
        
        DateFormat string `json:"dateFormat"`
        
        VarType string `json:"type"`
        
        Fields []string `json:"fields"`
        *Alias
    }{

        


        
        Values: []string{""},
        


        


        


        


        
        Group: []Locationsearchcriteria{{}},
        


        


        


        
        Fields: []string{""},
        

        Alias: (*Alias)(u),
    })
}

