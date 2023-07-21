package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SupportcenterdetailedcategorymodulesidebarMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SupportcenterdetailedcategorymodulesidebarDud struct { 
    

}

// Supportcenterdetailedcategorymodulesidebar
type Supportcenterdetailedcategorymodulesidebar struct { 
    // Enabled - Whether sidebar is enabled or not
    Enabled bool `json:"enabled"`

}

// String returns a JSON representation of the model
func (o *Supportcenterdetailedcategorymodulesidebar) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Supportcenterdetailedcategorymodulesidebar) MarshalJSON() ([]byte, error) {
    type Alias Supportcenterdetailedcategorymodulesidebar

    if SupportcenterdetailedcategorymodulesidebarMarshalled {
        return []byte("{}"), nil
    }
    SupportcenterdetailedcategorymodulesidebarMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

