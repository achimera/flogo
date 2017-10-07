package pushbulletnote

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/mitsuse/pushbullet-go"
	"github.com/mitsuse/pushbullet-go/requests"
)

// log is the default package logger
var log = logger.GetLogger("activity-pushbulletnote")

const (
	ivAccessToken     	= "accessToken"
	ivNoteTitle 		= "noteTitle"
	ivNote				= "note"
	ivEmailTarget 		= "emailTarget"
	ivChannelTarget		= "channelTarget"
	ovStatus       		= "status"
)

// PushbulletNoteActivity is a stub for your Activity implementation
type PushbulletNoteActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &PushbulletNoteActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *PushbulletNoteActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *PushbulletNoteActivity) Eval(context activity.Context) (done bool, err error) {
	accessToken 	:= context.GetInput(ivAccessToken)
	note 			:= context.GetInput(ivNote)
	noteTitle 		:= context.GetInput(ivNoteTitle)

	emailTarget 	:= context.GetInput(ivEmailTarget)
	channelTarget 	:= context.GetInput(ivChannelTarget)

	// Check if mandatory credentials are set in config
	if accessToken == nil {
		log.Error("Missing Pushbullet Access Token")
		err := activity.NewError("Pushbullet Access Token config not specified", "", nil)
		return false, err
	}

	// Check if there is a note to send
	if note == nil {
		log.Error("No Pushbullet note to send")
		context.SetOutput(ovStatus, "NO_NOTE")
		return true, nil
	}

	// Create a client for Pushbullet.
	pb := pushbullet.New(accessToken.(string))

	// Create a request. The following codes create a note, which is of note types.
	n := requests.NewNote()
	n.Title = noteTitle.(string)
	n.Body = note.(string)

	if emailTarget != nil {
		n.Email = emailTarget.(string)
		log.Info("Send Pushbullet note to email %v", emailTarget)
	}

	if channelTarget != nil {
		n.ChannelTag = channelTarget.(string)
		log.Info("Send Pushbullet note to channel %v", channelTarget)
	}

	// Send the note via Pushbullet.
	if _, err = pb.PostPushesNote(n); err != nil {
		log.Error("Pushbullet Connection Error : ", err)
		context.SetOutput(ovStatus, "CONNECT_ERR")
		return true, nil
	}

	context.SetOutput(ovStatus, "OK")

	return true, nil
}