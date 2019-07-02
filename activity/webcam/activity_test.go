package webcam

import (
	"testing"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestPlain(t *testing.T) {
	settings := &Settings{DeviceID: "0", ImageWidth: 1024, ImageHeigth: 720, Compression: 10}

	iCtx := test.NewActivityInitContext(settings, nil)
	act, err := New(iCtx)
	assert.Nil(t, err)

	tc := test.NewActivityContext(act.Metadata())

	done, err := act.Eval(tc)
	assert.Nil(t, err)
	assert.True(t, done)

	output := &Output{}
	err = tc.GetOutputObject(output)
	t.Log("Image byte array output: ", output.Image)
	t.Log("Base64String output: ", output.Base64String)

	/*
		file, err := os.Create("base64string.txt")
		if err != nil {
			t.Log("Error creating file", err)
		}
		defer file.Close()

		_, err = io.WriteString(file, output.Base64String)
		if err != nil {
			t.Log("Error writing to file", err)
		}
	*/

}
