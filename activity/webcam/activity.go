package webcam

import (
	"fmt"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"gocv.io/x/gocv"
)

// log is the default package logger
var log = logger.GetLogger("activity-webcam")

const (
	ivDeviceID = "deviceID"
	ivFilename = "fileName"

	ovImage  = "image"
	ovStatus = "status"
)

// WebcamActivity is a stub for your Activity implementation
type WebcamActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &WebcamActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *WebcamActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *WebcamActivity) Eval(context activity.Context) (done bool, err error) {
	deviceID := context.GetInput(ivDeviceID)
	fileName := context.GetInput(ivFilename).(string)

	// Check if mandatory credentials are set in config
	if fileName == "" {
		log.Error("Missing output fileName")
		err := activity.NewError("Raspicam filename config not specified", "", nil)
		return false, err
	}

	webcam, err := gocv.OpenVideoCapture(deviceID)
	webcam.Set(gocv.VideoCaptureFrameHeight, 1280)
	webcam.Set(gocv.VideoCaptureFrameWidth, 720)

	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", deviceID)
		return
	}
	defer webcam.Close()

	img := gocv.NewMat()
	defer img.Close()

	log.Info("Webcam capturing image...")
	if ok := webcam.Read(&img); !ok {
		fmt.Printf("cannot read device %v\n", deviceID)
		return
	}
	if img.Empty() {
		fmt.Printf("no image on device %v\n", deviceID)
		return
	}
	log.Info("Done. Image captured.")

	imgByte := img.ToBytes()
	log.Info(imgByte)
	gocv.IMWrite(fileName, img)

	context.SetOutput(ovImage, imgByte)
	context.SetOutput(ovStatus, "OK")

	return true, nil
}
