package raspicamera

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/dhowden/raspicam"
	"time"
	
	"os/exec"
	"bytes"
	"fmt"
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
	log.Info("Camera timeout set to ", timeout)
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

	//myPreview := raspicam.Preview { Mode: raspicam.PreviewDisabled, Opacity: 0, Rect: raspicam.Rect { X:0, Y:0, Width: 0, Height: 0}, }
	//still.BaseStill.Preview = myPreview

	preview := still.Preview
	//preview.Mode = raspicam.PreviewMode(raspicam.PreviewDisabled)
	preview.Mode = raspicam.PreviewDisabled
	
	log.Info("Preview Mode %v  ", preview.Mode)

	still.Preview = preview
	
	//preview := raspicam.Preview { Mode: raspicam.PreviewDisabled }

	if timeout != nil {
		still.Timeout = time.Duration(timeout.(int))
		log.Info("Camera timeout set to %v", timeout)
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
	log.Info("Raspicam capturing image...")
	
	cmd := exec.Command("raspistill", "-vf", "-hf", "-a", "1024", "-a", "8", "-a", "achimera| %F %r", "-o", filename.(string))
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	myErr := cmd.Run()
	// Check for errors
	if myErr != nil {
		log.Error(fmt.Sprint(myErr) + ": " + stderr.String())
	}

	//raspicam.Capture(still, f, errCh)
	log.Info("Raspicam created image file: ", filename)

	context.SetOutput(ovStatus, "OK")

	return true, nil
}
