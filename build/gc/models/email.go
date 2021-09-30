package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EmailMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EmailDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Email
type Email struct { 
    // State - The connection state of this communication.
    State string `json:"state"`


    // Id - A globally unique identifier for this communication.
    Id string `json:"id"`


    // Held - True if this call is held and the person on this side hears silence.
    Held bool `json:"held"`


    // Subject - The subject for the initial email that started this conversation.
    Subject string `json:"subject"`


    // MessagesSent - The number of email messages sent by this participant.
    MessagesSent int `json:"messagesSent"`


    // Segments - The time line of the participant's email, divided into activity segments.
    Segments []Segment `json:"segments"`


    // Direction - The direction of the email
    Direction string `json:"direction"`


    // RecordingId - A globally unique identifier for the recording associated with this call.
    RecordingId string `json:"recordingId"`


    // ErrorInfo
    ErrorInfo Errorbody `json:"errorInfo"`


    // DisconnectType - System defined string indicating what caused the communication to disconnect. Will be null until the communication disconnects.
    DisconnectType string `json:"disconnectType"`


    // StartHoldTime - The timestamp the email was placed on hold in the cloud clock if the email is currently on hold. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartHoldTime time.Time `json:"startHoldTime"`


    // StartAlertingTime - The timestamp the communication has when it is first put into an alerting state. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartAlertingTime time.Time `json:"startAlertingTime"`


    // ConnectedTime - The timestamp when this communication was connected in the cloud clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ConnectedTime time.Time `json:"connectedTime"`


    // DisconnectedTime - The timestamp when this communication disconnected from the conversation in the provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DisconnectedTime time.Time `json:"disconnectedTime"`


    // AutoGenerated - Indicates that the email was auto-generated like an Out of Office reply.
    AutoGenerated bool `json:"autoGenerated"`


    // Provider - The source provider for the email.
    Provider string `json:"provider"`


    // ScriptId - The UUID of the script to use.
    ScriptId string `json:"scriptId"`


    // PeerId - The id of the peer communication corresponding to a matching leg for this communication.
    PeerId string `json:"peerId"`


    // MessageId - A globally unique identifier for the stored content of this communication.
    MessageId string `json:"messageId"`


    // DraftAttachments - A list of uploaded attachments on the email draft.
    DraftAttachments []Attachment `json:"draftAttachments"`


    // Spam - Indicates if the inbound email was marked as spam.
    Spam bool `json:"spam"`


    // Wrapup - Call wrap up or disposition data.
    Wrapup Wrapup `json:"wrapup"`


    // AfterCallWork - After-call work for the communication.
    AfterCallWork Aftercallwork `json:"afterCallWork"`


    // AfterCallWorkRequired - Indicates if after-call work is required for a communication. Only used when the ACW Setting is Agent Requested.
    AfterCallWorkRequired bool `json:"afterCallWorkRequired"`

}

// String returns a JSON representation of the model
func (o *Email) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Segments = []Segment{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.DraftAttachments = []Attachment{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Email) MarshalJSON() ([]byte, error) {
    type Alias Email

    if EmailMarshalled {
        return []byte("{}"), nil
    }
    EmailMarshalled = true

    return json.Marshal(&struct { 
        State string `json:"state"`
        
        Id string `json:"id"`
        
        Held bool `json:"held"`
        
        Subject string `json:"subject"`
        
        MessagesSent int `json:"messagesSent"`
        
        Segments []Segment `json:"segments"`
        
        Direction string `json:"direction"`
        
        RecordingId string `json:"recordingId"`
        
        ErrorInfo Errorbody `json:"errorInfo"`
        
        DisconnectType string `json:"disconnectType"`
        
        StartHoldTime time.Time `json:"startHoldTime"`
        
        StartAlertingTime time.Time `json:"startAlertingTime"`
        
        ConnectedTime time.Time `json:"connectedTime"`
        
        DisconnectedTime time.Time `json:"disconnectedTime"`
        
        AutoGenerated bool `json:"autoGenerated"`
        
        Provider string `json:"provider"`
        
        ScriptId string `json:"scriptId"`
        
        PeerId string `json:"peerId"`
        
        MessageId string `json:"messageId"`
        
        DraftAttachments []Attachment `json:"draftAttachments"`
        
        Spam bool `json:"spam"`
        
        Wrapup Wrapup `json:"wrapup"`
        
        AfterCallWork Aftercallwork `json:"afterCallWork"`
        
        AfterCallWorkRequired bool `json:"afterCallWorkRequired"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        
        Segments: []Segment{{}},
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        DraftAttachments: []Attachment{{}},
        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

