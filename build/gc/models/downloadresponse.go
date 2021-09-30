package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DownloadresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DownloadresponseDud struct { 
    


    


    

}

// Downloadresponse
type Downloadresponse struct { 
    // ContentLocationUri
    ContentLocationUri string `json:"contentLocationUri"`


    // ImageUri
    ImageUri string `json:"imageUri"`


    // Thumbnails
    Thumbnails []Documentthumbnail `json:"thumbnails"`

}

// String returns a JSON representation of the model
func (o *Downloadresponse) String() string {
    
    
    
    
    
    
    
    
    
    
     o.Thumbnails = []Documentthumbnail{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Downloadresponse) MarshalJSON() ([]byte, error) {
    type Alias Downloadresponse

    if DownloadresponseMarshalled {
        return []byte("{}"), nil
    }
    DownloadresponseMarshalled = true

    return json.Marshal(&struct { 
        ContentLocationUri string `json:"contentLocationUri"`
        
        ImageUri string `json:"imageUri"`
        
        Thumbnails []Documentthumbnail `json:"thumbnails"`
        
        *Alias
    }{
        

        

        

        

        

        
        Thumbnails: []Documentthumbnail{{}},
        

        
        Alias: (*Alias)(u),
    })
}

