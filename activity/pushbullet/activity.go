package pushbullet

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/op/go-logging"
	"github.com/mitsuse/pushbullet-go"
	"github.com/mitsuse/pushbullet-go/requests" 
)

// log is the default package logger
var log = logging.MustGetLogger("activity-pushbullet")

const (
	ivAccToken = "accToken"
	ivMessage = "message"
	ivMessageTitle = "messageTitle"
	ovStatus = "status"
)

// PushbulletActivity is a stub for your Activity implementation
type PushbulletActivity struct {
	metadata *activity.Metadata
}

// init create & register activity
func init() {
	md := activity.NewMetadata(jsonMetadata)
	activity.Register(&PushbulletActivity{metadata: md})
}

// Metadata implements activity.Activity.Metadata
func (a *PushbulletActivity) Metadata() *activity.Metadata {
	return a.metadata
}


// Eval implements activity.Activity.Eval
func (a *PushbulletActivity) Eval(context activity.Context) (done bool, err error)  {

	accToken := context.GetInput(ivAccToken)
	message := context.GetInput(ivMessage)
	messageTitle := context.GetInput(ivMessageTitle)

	// Check if mandatory credentials are set in config
	if accToken == nil {
		log.Error("Missing Pushbullet Access Token")
		err := activity.NewError("Access Token config not specified","",nil)
		return false, err
	}

	// Check if there is a message to send
	if message == nil {
		log.Error("No Message to Send")
		context.SetOutput(ovStatus, "NO_MSG")
		return true, nil
	}

	// Create a client for Pushbullet.
	pb := pushbullet.New(accToken.(string))

	// Create a push. The following codes create a note, which is one of push types.
	n := requests.NewNote()
	n.Title = messageTitle.(string)
	n.Body = message.(string)

	// Send the note via Pushbullet.
	if _, err = pb.PostPushesNote(n); err != nil {
		log.Error("Pushbullet Connection Error : ", err)
		context.SetOutput(ovStatus, "CONNECT_ERR")
		return true, nil
	}


	context.SetOutput(ovStatus, "OK")
	
	return true, nil
}
