package raspicamera

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/dhowden/raspicam"
	"time"
	"os"
	//"path"
)

// log is the default package logger
var log = logger.GetLogger("activity-raspicamera")

const (
	ivTimeout 		= "timeout" //delay before the image is taken
	ivSharpness 	= "sharpness"
	ivBrightness 	= "brightness"
	ivContrast		= "contrast"
	ivSaturation 	= "saturation"
	ivISO			= "iso"
	ivFilename		= "filename"

	ovStatus        = "status"
)

// RaspicameraActivity is a stub for your Activity implementation
type RaspicameraActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &RaspicameraActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *RaspicameraActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *RaspicameraActivity) Eval(context activity.Context) (done bool, err error) {
	timeout := context.GetInput(ivTimeout)
	sharpness := context.GetInput(ivSharpness)
	brightness := context.GetInput(ivBrightness)
	contrast := context.GetInput(ivContrast)
	saturation := context.GetInput(ivSaturation)
	iso := context.GetInput(ivISO)
	filename := context.GetInput(ivFilename)

	// Check if mandatory credentials are set in config
	if filename == nil {
		log.Error("Missing output filename")
		err := activity.NewError("Raspicam filename config not specified", "", nil)
		return false, err
	}

	// Create a client for raspicam.
	still := raspicam.NewStill()

	//preview := still.Preview
	//still.Preview = preview
	preview := raspicam.Preview { Mode: raspicam.PreviewDisabled }
	still.Preview = preview

	if timeout != nil {
		still.Timeout = time.Duration(timeout.(int))
		log.Debug("Camera timeout set to %v", timeout)
	}
	if sharpness != nil {
		still.Camera.Sharpness = sharpness.(int)
		log.Debug("Camera sharpness set to %v", sharpness)
	}
	if brightness != nil {
		still.Camera.Brightness = brightness.(int)
		log.Debug("Camera brightness set to %v", brightness)
	}
	if contrast != nil {
		still.Camera.Contrast = contrast.(int)
		log.Debug("Camera contrast set to %v", contrast)
	}
	if saturation != nil {
		still.Camera.Saturation = saturation.(int)
		log.Debug("Camera saturation set to %v", saturation)
	}
	if iso != nil {
		still.Camera.ISO = iso.(int)
		log.Debug("Camera iso set to %v", iso)
	}

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
	f, err := os.Create(filename.(string))
	if err != nil {
		log.Error("Raspicam error on creating the image file: %v", err)
		context.SetOutput(ovStatus, "IMAGE_CREATE__ERR")
		return true, nil
		//fmt.Fprintf(os.Stderr, "create file: %v", err)

	}
	defer f.Close()

	errCh := make(chan error)
	go func() {
		for x := range errCh {
			//fmt.Fprintf(os.Stderr, "%v\n", x)
			log.Error("%v\n", x)
		}
	}()

	//cmd := exec.Command("raspistill", "-vf", "-hf", "-a", "1024", "-a", "8", "-a", "achimera| %F %r", "-o", imageFile)
	raspicam.Capture(still, f, errCh)
	log.Info("Raspicam created image file: %v", filename)

	context.SetOutput(ovStatus, "OK")

	return true, nil
}
