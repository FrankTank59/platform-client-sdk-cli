package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RoutingactivitymetricvalueMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RoutingactivitymetricvalueDud struct { 
    


    

}

// Routingactivitymetricvalue
type Routingactivitymetricvalue struct { 
    // Metric - metric
    Metric string `json:"metric"`


    // Count - metric count
    Count int `json:"count"`

}

// String returns a JSON representation of the model
func (o *Routingactivitymetricvalue) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Routingactivitymetricvalue) MarshalJSON() ([]byte, error) {
    type Alias Routingactivitymetricvalue

    if RoutingactivitymetricvalueMarshalled {
        return []byte("{}"), nil
    }
    RoutingactivitymetricvalueMarshalled = true

    return json.Marshal(&struct {
        
        Metric string `json:"metric"`
        
        Count int `json:"count"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

