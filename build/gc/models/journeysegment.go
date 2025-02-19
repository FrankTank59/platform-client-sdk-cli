package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneysegmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneysegmentDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`


    


    

}

// Journeysegment
type Journeysegment struct { 
    


    // IsActive - Whether or not the segment is active.
    IsActive bool `json:"isActive"`


    // DisplayName - The display name of the segment.
    DisplayName string `json:"displayName"`


    // Version - The version of the segment.
    Version int `json:"version"`


    // Description - A description of the segment.
    Description string `json:"description"`


    // Color - The hexadecimal color value of the segment.
    Color string `json:"color"`


    // Scope - The target entity that a segment applies to.
    Scope string `json:"scope"`


    // ShouldDisplayToAgent - Whether or not the segment should be displayed to agent/supervisor users.
    ShouldDisplayToAgent bool `json:"shouldDisplayToAgent"`


    // Context - The context of the segment.
    Context Context `json:"context"`


    // Journey - The pattern of rules defining the segment.
    Journey Journey `json:"journey"`


    // ExternalSegment - Details of an entity corresponding to this segment in an external system.
    ExternalSegment Externalsegment `json:"externalSegment"`


    // AssignmentExpirationDays - Time, in days, from when the segment is assigned until it is automatically unassigned.
    AssignmentExpirationDays int `json:"assignmentExpirationDays"`


    


    // CreatedDate - Timestamp indicating when the segment was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CreatedDate time.Time `json:"createdDate"`


    // ModifiedDate - Timestamp indicating when the the segment was last updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ModifiedDate time.Time `json:"modifiedDate"`

}

// String returns a JSON representation of the model
func (o *Journeysegment) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeysegment) MarshalJSON() ([]byte, error) {
    type Alias Journeysegment

    if JourneysegmentMarshalled {
        return []byte("{}"), nil
    }
    JourneysegmentMarshalled = true

    return json.Marshal(&struct {
        
        IsActive bool `json:"isActive"`
        
        DisplayName string `json:"displayName"`
        
        Version int `json:"version"`
        
        Description string `json:"description"`
        
        Color string `json:"color"`
        
        Scope string `json:"scope"`
        
        ShouldDisplayToAgent bool `json:"shouldDisplayToAgent"`
        
        Context Context `json:"context"`
        
        Journey Journey `json:"journey"`
        
        ExternalSegment Externalsegment `json:"externalSegment"`
        
        AssignmentExpirationDays int `json:"assignmentExpirationDays"`
        
        CreatedDate time.Time `json:"createdDate"`
        
        ModifiedDate time.Time `json:"modifiedDate"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

