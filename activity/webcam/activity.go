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
	//img := gocv.NewMatFromBytes

	img := gocv.NewMat()
	defer img.Close()

	if ok := webcam.Read(&img); !ok {
		fmt.Printf("cannot read device %v\n", deviceID)
		return
	}
	if img.Empty() {
		fmt.Printf("no image on device %v\n", deviceID)
		return
	}

	imgByte := img.ToBytes()
	gocv.IMWrite(fileName, img)

	/*
		imageDirectory, imageFile := path.Split(filename.(string))
		if imageFile == "" {
			context.SetOutput(ovStatus, "NO_FILENAME_ERR")
			return true, nil
		}
		if imageDirectory == "" {
			if _, err := os.Stat(imageDirectory); os.IsNotExist(err) {
				os.MkdirAll(imageDirectory, 0777)
			}
		}*/

	// create the folder for the image
	/*
		f, err := os.Create(filename.(string))
		if err != nil {
			log.Error("Raspicam error on creating the image file: ", err)
			context.SetOutput(ovStatus, "IMAGE_CREATE__ERR")
			return true, nil
			//fmt.Fprintf(os.Stderr, "create file: %v", err)

		}
		defer f.Close()

		errCh := make(chan error)
		go func() {
			for x := range errCh {
				//fmt.Fprintf(os.Stderr, "%v\n", x)
				log.Error("Error %v\n", x)
			}
		}()
	*/
	log.Info("Webcam capturing image...")

	//cmd := exec.Command("raspistill", "-vf", "-hf", "-a", "1024", "-a", "8", "-a", "TIBCO - %d-%m-%Y %X %r", "-o", filename.(string))
	//var stderr bytes.Buffer
	//cmd.Stderr = &stderr

	//myErr := cmd.Run()
	// Check for errors
	//if myErr != nil {
	//	log.Error(fmt.Sprint(myErr) + ": " + stderr.String())
	//}

	context.SetOutput(ovImage, imgByte)
	context.SetOutput(ovStatus, "OK")

	return true, nil
}
