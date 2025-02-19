package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LabelcreaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LabelcreaterequestDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Labelcreaterequest
type Labelcreaterequest struct { 
    


    // Name - The name of the label.
    Name string `json:"name"`


    // Color - The color for the label.
    Color string `json:"color"`


    // ExternalId - The external id associated with the label.
    ExternalId string `json:"externalId"`


    

}

// String returns a JSON representation of the model
func (o *Labelcreaterequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Labelcreaterequest) MarshalJSON() ([]byte, error) {
    type Alias Labelcreaterequest

    if LabelcreaterequestMarshalled {
        return []byte("{}"), nil
    }
    LabelcreaterequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Color string `json:"color"`
        
        ExternalId string `json:"externalId"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

