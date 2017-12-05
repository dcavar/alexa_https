/*
AlexaRequest.go
(C) 2017 by Damir Cavar (http://damir.cavar.me/), Rashmi Bidanta, Prateek Srivastava.

Defining the structrs for the Alexa Request JSON.

TODO
See at the end of the file all the missing implementations of structs that
might be relevant or necessary for your particular implementation of a slot
responder.
*/


package main


// These structs are all taken from the Alexa JSON Requests specification at
// https://developer.amazon.com/docs/custom-skills/request-and-response-json-reference.html


type AlexaRequest struct {
    Version string `json:"version"`
    Session *Session `json:"session,omitempty"`
    Context *Context `json:"context,omitempty"`
    Request *Request `json:"request,omitempty"`
}

type Session struct {
    New bool `json:"new,omitempty"`
    SessionId string `json:"sessionId,omitempty"`
    Application *SessionApplication `json:"application,omitempty"`
    Attributes *SessionAttributes `json:"attributes,omitempty"`
    User *SessionUser `json:"user,omitempty"`
}

type SessionApplication struct {
    ApplicationId string `json:"applicationId,omitempty"`
}

type SessionAttributes struct {
    Key string `json:"key,omitempty"`
}

type SessionUser struct {
    UserID string `json:"userId,omitempty"`
    AccessToken string `json:"accessToken,omitempty"`
    Permissions *SessionUserPermissions `json:"permissions,omitempty"`
}

type SessionUserPermissions struct {
    ConsentToken string `json:"consentToken,omitempty"`
}

type Context struct {
    System *ContextSystem `json:"system,omitempty"`
    AudioPlayer *ContextAudioPlayer `json:"AudioPlayer,omitempty"`
}

type ContextAudioPlayer struct {
    PlayerActivity string `json:"playerActivity,omitempty"`
    Token string `json:"token,omitempty"`
    OffsetInMilliseconds int `json:"offsetInMilliseconds,omitempty"`
}

type ContextSystem struct {
    Device *ContextSystemDevice `json:"device,omitempty"`
    Application *ContextSystemApplication `json:"application,omitempty"`
    User *ContextSystemUser `json:"user,omitempty"`
    ApiEndpoint string `json:"apiEndpoint,omitempty"`
    ApiAccessToken string `json:"apiAccessToken,omitempty"`
}

type ContextSystemUser struct {
    UserID string `json:"userId,omitempty"`
    AccessToken string `json:"accessToken,omitempty"`
    Permissions *ContextSystemUserPermissions `json:"permissions,omitempty"`
}

type ContextSystemUserPermissions struct {
    ConsentToken string `json:"consentToken,omitempty"`
}

type ContextSystemApplication struct {
    ApplicationID string `json:"applicationId,omitempty"`
}

type ContextSystemDevice struct {
    DeviceID string `json:"deviceId,omitempty"`
    SupportedInterfaces *ContextSystemDeviceSupportedInterfaces `json:"supportedInterfaces,omitempty"`
}

type ContextSystemDeviceSupportedInterfaces struct {
    AudioPlayer *ContextSystemDeviceSupportedInterfacesAudioPlayer `json:"AudioPlayer,omitempty"`
}

type ContextSystemDeviceSupportedInterfacesAudioPlayer struct {
  // there are missing elements in this struct here
}

type Request struct {
    Type string `json:"type,omitempty"`
    RequestId string `json:"requestId,omitempty"`
    Timestamp string `json:"timestamp,omitempty"`
    DialogState string `json:"dialogState,omitempty"`
    Locale string `json:"locale,omitempty"`
    Intent *RequestIntent `json:"intent,omitempty"`
}

type RequestIntent struct {
    Name string `json:"name,omitempty"`
    ConfirmationStatus string `json:"confirmationStatus,omitempty"`
    Slots *RequestIntentSlots `json:"slots,omitempty"`
}

type RequestIntentSlots struct {
  // TODO implement your slots here!
}


// TODO implement your slot structs here!
