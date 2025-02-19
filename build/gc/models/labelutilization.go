package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LabelutilizationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LabelutilizationDud struct { 
    


    

}

// Labelutilization
type Labelutilization struct { 
    // MaximumCapacity - Defines the maximum number of interactions with this label that an agent can handle at one time.
    MaximumCapacity int `json:"maximumCapacity"`


    // InterruptingLabelIds - Defines other labels that can interrupt an interaction with this label.
    InterruptingLabelIds []string `json:"interruptingLabelIds"`

}

// String returns a JSON representation of the model
func (o *Labelutilization) String() string {
    
     o.InterruptingLabelIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Labelutilization) MarshalJSON() ([]byte, error) {
    type Alias Labelutilization

    if LabelutilizationMarshalled {
        return []byte("{}"), nil
    }
    LabelutilizationMarshalled = true

    return json.Marshal(&struct {
        
        MaximumCapacity int `json:"maximumCapacity"`
        
        InterruptingLabelIds []string `json:"interruptingLabelIds"`
        *Alias
    }{

        


        
        InterruptingLabelIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

