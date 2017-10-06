package pushbullet

import (
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
)

func TestRegistered(t *testing.T) {
	act := activity.Get("pushbullet")

	if act == nil {
		t.Error("Activity Not Registered")
		t.Fail()
		return
	}
}



func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	md := activity.NewMetadata(jsonMetadata)
	act := &PushbulletActivity{metadata: md}

	tc := test.NewTestActivityContext(md)
	tc.SetInput(ivAccToken, "o.AYA7hnpZIIoiPlkr7j8clD0OEaLHcF2u")
	tc.SetInput(ivMessageTitle, "Flogo")
	tc.SetInput(ivMessage, "Go Flogo")
	//setup attrs

	act.Eval(tc)

	//check result attr
}
