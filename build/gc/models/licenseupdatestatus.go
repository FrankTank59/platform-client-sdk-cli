package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LicenseupdatestatusMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LicenseupdatestatusDud struct { 
    


    


    

}

// Licenseupdatestatus
type Licenseupdatestatus struct { 
    // UserId
    UserId string `json:"userId"`


    // LicenseId
    LicenseId string `json:"licenseId"`


    // Result
    Result string `json:"result"`

}

// String returns a JSON representation of the model
func (o *Licenseupdatestatus) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Licenseupdatestatus) MarshalJSON() ([]byte, error) {
    type Alias Licenseupdatestatus

    if LicenseupdatestatusMarshalled {
        return []byte("{}"), nil
    }
    LicenseupdatestatusMarshalled = true

    return json.Marshal(&struct { 
        UserId string `json:"userId"`
        
        LicenseId string `json:"licenseId"`
        
        Result string `json:"result"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

