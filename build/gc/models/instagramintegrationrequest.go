package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InstagramintegrationrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InstagramintegrationrequestDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Instagramintegrationrequest
type Instagramintegrationrequest struct { 
    


    // Name - The name of the Instagram Integration
    Name string `json:"name"`


    // SupportedContent - Defines the SupportedContent profile configured for an integration
    SupportedContent Supportedcontentreference `json:"supportedContent"`


    // MessagingSetting - Defines the message settings to be applied for this integration
    MessagingSetting Messagingsettingrequestreference `json:"messagingSetting"`


    // PageAccessToken - The long-lived Page Access Token of Instagram page.  See https://developers.facebook.com/docs/facebook-login/access-tokens.  When a pageAccessToken is provided, pageId and userAccessToken are not required.
    PageAccessToken string `json:"pageAccessToken"`


    // UserAccessToken - The short-lived User Access Token of Instagram user logged into Facebook app.  See https://developers.facebook.com/docs/facebook-login/access-tokens.  When userAccessToken is provided, pageId is mandatory.  When userAccessToken/pageId combination is provided, pageAccessToken is not required.
    UserAccessToken string `json:"userAccessToken"`


    // PageId - The page ID of Instagram page. The pageId is required when userAccessToken is provided.
    PageId string `json:"pageId"`


    // AppId - The app ID of Facebook app. The appId is required when a customer wants to use their own approved Facebook app.
    AppId string `json:"appId"`


    // AppSecret - The app Secret of Facebook app. The appSecret is required when appId is provided.
    AppSecret string `json:"appSecret"`


    

}

// String returns a JSON representation of the model
func (o *Instagramintegrationrequest) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Instagramintegrationrequest) MarshalJSON() ([]byte, error) {
    type Alias Instagramintegrationrequest

    if InstagramintegrationrequestMarshalled {
        return []byte("{}"), nil
    }
    InstagramintegrationrequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        SupportedContent Supportedcontentreference `json:"supportedContent"`
        
        MessagingSetting Messagingsettingrequestreference `json:"messagingSetting"`
        
        PageAccessToken string `json:"pageAccessToken"`
        
        UserAccessToken string `json:"userAccessToken"`
        
        PageId string `json:"pageId"`
        
        AppId string `json:"appId"`
        
        AppSecret string `json:"appSecret"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

