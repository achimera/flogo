package pushbulletlink

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/mitsuse/pushbullet-go"
	"github.com/mitsuse/pushbullet-go/requests"
)

// log is the default package logger
var log = logger.GetLogger("activity-pushbulletlink")

const (
	ivAccessToken		= "accessToken"
	ivLinkTitle			= "linkTitle"
	ivLinkMsg			= "linkMsg"
	ivLinkUrl			= "linkUrl"
	ivEmailTarget		= "emailTarget"
	ivChannelTarget		= "channelTarget"
	ovStatus			= "status"
)

// PushbulletLinkActivity is a stub for your Activity implementation
type PushbulletLinkActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &PushbulletLinkActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *PushbulletLinkActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *PushbulletLinkActivity) Eval(context activity.Context) (done bool, err error) {
	accessToken 	:= context.GetInput(ivAccessToken)
	linkTitle 		:= context.GetInput(ivLinkTitle)
	linkMsg 		:= context.GetInput(ivLinkMsg)
	linkUrl 		:= context.GetInput(ivLinkUrl)

	emailTarget 	:= context.GetInput(ivEmailTarget)
	channelTarget 	:= context.GetInput(ivChannelTarget)

	// Check if mandatory credentials are set in config
	if accessToken == nil {
		log.Error("Missing Pushbullet Access Token")
		err := activity.NewError("Pushbullet Access Token config not specified", "", nil)
		return false, err
	}

	// Check if there is a link URL to send
	if linkUrl == nil {
		log.Error("No Pushbullet link URL to send")
		context.SetOutput(ovStatus, "NO_LINK_URL_ERR")
		return true, nil
	}

	// Create a client for Pushbullet.
	pb := pushbullet.New(accessToken.(string))

	// Create a request. The following code creates a link message, which is of type "link".
	n := requests.NewLink()
	n.Title = linkTitle.(string)
	n.Body = linkMsg.(string)
	n.Url = linkUrl.(string)

	if emailTarget != nil {
		n.Email = emailTarget.(string)
		log.Info("Send Pushbullet link to email %v", emailTarget)
	}

	if channelTarget != nil {
		n.ChannelTag = channelTarget.(string)
		log.Info("Send Pushbullet link to channel %v", channelTarget)
	}

	if emailTarget != "" && channelTarget != "" {
		log.Error("Can't send link to email and channel. Please choose only one target")
		context.SetOutput(ovStatus, "TOO_MANY_TARGETS_ERR")
		return true, nil
	}

	// Send the link via Pushbullet.
	if _, err = pb.PostPushesLink(n); err != nil {
		log.Error("Pushbullet Connection Error : ", err)
		context.SetOutput(ovStatus, "CONNECT_ERR")
		return true, nil
	}

	context.SetOutput(ovStatus, "OK")

	return true, nil
}