package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationentitylistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationentitylistingDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Conversationentitylisting
type Conversationentitylisting struct { 
    // Entities
    Entities []Conversation `json:"entities"`


    // PageSize
    PageSize int `json:"pageSize"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // Total
    Total int `json:"total"`


    // FirstUri
    FirstUri string `json:"firstUri"`


    // NextUri
    NextUri string `json:"nextUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // LastUri
    LastUri string `json:"lastUri"`


    // PageCount
    PageCount int `json:"pageCount"`

}

// String returns a JSON representation of the model
func (o *Conversationentitylisting) String() string {
     o.Entities = []Conversation{{}} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationentitylisting) MarshalJSON() ([]byte, error) {
    type Alias Conversationentitylisting

    if ConversationentitylistingMarshalled {
        return []byte("{}"), nil
    }
    ConversationentitylistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Conversation `json:"entities"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Total int `json:"total"`
        
        FirstUri string `json:"firstUri"`
        
        NextUri string `json:"nextUri"`
        
        PreviousUri string `json:"previousUri"`
        
        SelfUri string `json:"selfUri"`
        
        LastUri string `json:"lastUri"`
        
        PageCount int `json:"pageCount"`
        *Alias
    }{

        
        Entities: []Conversation{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

