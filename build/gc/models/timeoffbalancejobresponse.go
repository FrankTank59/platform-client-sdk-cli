package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TimeoffbalancejobresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TimeoffbalancejobresponseDud struct { 
    


    

}

// Timeoffbalancejobresponse
type Timeoffbalancejobresponse struct { 
    // Entities
    Entities []Timeoffbalanceresponse `json:"entities"`


    // Status - The status of the time off balance job
    Status string `json:"status"`

}

// String returns a JSON representation of the model
func (o *Timeoffbalancejobresponse) String() string {
     o.Entities = []Timeoffbalanceresponse{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Timeoffbalancejobresponse) MarshalJSON() ([]byte, error) {
    type Alias Timeoffbalancejobresponse

    if TimeoffbalancejobresponseMarshalled {
        return []byte("{}"), nil
    }
    TimeoffbalancejobresponseMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Timeoffbalanceresponse `json:"entities"`
        
        Status string `json:"status"`
        *Alias
    }{

        
        Entities: []Timeoffbalanceresponse{{}},
        


        

        Alias: (*Alias)(u),
    })
}

