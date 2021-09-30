package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningassignmentaggregateparamMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningassignmentaggregateparamDud struct { 
    


    


    


    

}

// Learningassignmentaggregateparam
type Learningassignmentaggregateparam struct { 
    // Interval - Specifies the range of due dates to be used for filtering. Milliseconds will be truncated. A maximum of 1 year can be specified in the range. End date is not inclusive. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    Interval string `json:"interval"`


    // Metrics - The list of metrics to be returned. If omitted, all metrics are returned.
    Metrics []string `json:"metrics"`


    // GroupBy - Specifies if the aggregated data is combined into a single set of metrics (groupBy is empty or not specified), or contains an element per attendeeId (groupBy is \"attendeeId\")
    GroupBy []string `json:"groupBy"`


    // Filter - The filter applied to the data.  This is ANDed with the interval parameter. 
    Filter Learningassignmentaggregatequeryrequestfilter `json:"filter"`

}

// String returns a JSON representation of the model
func (o *Learningassignmentaggregateparam) String() string {
    
    
    
    
    
    
     o.Metrics = []string{""} 
    
    
    
     o.GroupBy = []string{""} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningassignmentaggregateparam) MarshalJSON() ([]byte, error) {
    type Alias Learningassignmentaggregateparam

    if LearningassignmentaggregateparamMarshalled {
        return []byte("{}"), nil
    }
    LearningassignmentaggregateparamMarshalled = true

    return json.Marshal(&struct { 
        Interval string `json:"interval"`
        
        Metrics []string `json:"metrics"`
        
        GroupBy []string `json:"groupBy"`
        
        Filter Learningassignmentaggregatequeryrequestfilter `json:"filter"`
        
        *Alias
    }{
        

        

        

        
        Metrics: []string{""},
        

        

        
        GroupBy: []string{""},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

