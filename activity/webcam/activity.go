package webcam

import (
	"fmt"
	"strconv"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
	"gocv.io/x/gocv"
)

func init() {
	_ = activity.Register(&Activity{}, New)
}

var activityMd = activity.ToMetadata(&Input{}, &Output{})

// Activity is a kafka activity
type Activity struct {
	deviceID int
}

// New create a new  activity
func New(ctx activity.InitContext) (activity.Activity, error) {
	settings := &Settings{}

	err := metadata.MapToStruct(ctx.Settings(), settings, true)
	if err != nil {
		return nil, err
	}

	deviceID, err := strconv.Atoi(settings.deviceID)

	act := &Activity{deviceID: deviceID}
	return act, nil
}

// Metadata returns the metadata for the kafka activity
func (*Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements the evaluation of the kafka activity
func (act *Activity) Eval(ctx activity.Context) (done bool, err error) {
	//input := &Input{}

	webcam, err := gocv.OpenVideoCapture(act.deviceID)
	webcam.Set(gocv.VideoCaptureFrameHeight, 1280)
	webcam.Set(gocv.VideoCaptureFrameWidth, 720)

	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", act.deviceID)
		return
	}
	defer webcam.Close()

	img := gocv.NewMat()
	defer img.Close()

	ctx.Logger().Info("Webcam capturing image...")
	if ok := webcam.Read(&img); !ok {
		fmt.Printf("cannot read device %v\n", act.deviceID)
		return
	}
	if img.Empty() {
		fmt.Printf("no image on device %v\n", act.deviceID)
		return
	}
	ctx.Logger().Info("Done. Image captured.")

	imgByte := img.ToBytes()
	ctx.Logger().Info(imgByte)
	gocv.IMWrite("test.png", img)

	output := &Output{}
	output.Image = imgByte
	output.Status = "OK"

	/*
		if ctx.Logger().DebugEnabled() {
			ctx.Logger().Debugf("Kafka message [%v] sent successfully on partition [%d] and offset [%d]",
				input.Message, partition, offset)
		}
	*/

	err = ctx.SetOutputObject(output)
	if err != nil {
		return false, err
	}

	return true, nil
}
