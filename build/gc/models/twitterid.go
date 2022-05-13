package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TwitteridMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TwitteridDud struct { 
    


    


    


    Verified bool `json:"verified"`


    ProfileUrl string `json:"profileUrl"`

}

// Twitterid - User information for a twitter account
type Twitterid struct { 
    // Id - twitter user.id_str
    Id string `json:"id"`


    // Name - twitter user.name
    Name string `json:"name"`


    // ScreenName - twitter user.screen_name
    ScreenName string `json:"screenName"`


    


    

}

// String returns a JSON representation of the model
func (o *Twitterid) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Twitterid) MarshalJSON() ([]byte, error) {
    type Alias Twitterid

    if TwitteridMarshalled {
        return []byte("{}"), nil
    }
    TwitteridMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        ScreenName string `json:"screenName"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

