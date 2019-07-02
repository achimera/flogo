package webcam

import (
	"encoding/base64"
	"fmt"
	"strconv"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
	"gocv.io/x/gocv"
)

func init() {
	_ = activity.Register(&Activity{}, New)
}

var activityMd = activity.ToMetadata(&Output{})

// Activity is a kafka activity
type Activity struct {
	deviceID    int
	imageWidth  int
	imageHeight int
	compression int
}

// New create a new  activity
func New(ctx activity.InitContext) (activity.Activity, error) {
	settings := &Settings{}
	ctx.Logger().Info("Settings done.")

	err := metadata.MapToStruct(ctx.Settings(), settings, true)
	if err != nil {
		return nil, err
	}

	deviceID, err := strconv.Atoi(settings.DeviceID)
	ctx.Logger().Info("deviceID has been set: ", deviceID)

	imageWidth := settings.ImageWidth
	imageHeight := settings.ImageHeigth

	ctx.Logger().Info("Image resolution has been set to: ", imageWidth, "x", imageHeight)

	compression := 3
	if (settings.Compression < 0) || (settings.Compression > 9) {
		ctx.Logger().Info("Image compression out of range [0...9] or not set. Using default compression factor of 3")
	} else {
		compression := settings.Compression
		ctx.Logger().Info("Image compression factor has been set to: ", compression)
	}

	act := &Activity{deviceID: deviceID, imageWidth: imageWidth, imageHeight: imageHeight, compression: compression}
	return act, nil
}

// Metadata returns the metadata for the kafka activity
func (*Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements the evaluation of the kafka activity
func (act *Activity) Eval(ctx activity.Context) (done bool, err error) {
	//settings := &Settings{}

	webcam, err := gocv.OpenVideoCapture(act.deviceID)

	webcam.Set(gocv.VideoCaptureFrameWidth, float64(act.imageWidth))
	webcam.Set(gocv.VideoCaptureFrameHeight, float64(act.imageHeight))

	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", act.deviceID)
		return
	}
	defer webcam.Close()

	img := gocv.NewMat()
	defer img.Close()

	ctx.Logger().Info("Webcam is taking a picture...")
	if ok := webcam.Read(&img); !ok {
		fmt.Printf("cannot read device %v\n", act.deviceID)
		return
	}
	if img.Empty() {
		fmt.Printf("no image on device %v\n", act.deviceID)
		return
	}
	ctx.Logger().Info("Done. Image captured.")

	//imgByte := img.ToBytes()
	//ctx.Logger().Info(imgByte)
	//gocv.IMWrite("test.png", img) //just for debug

	imageBuffer, err := gocv.IMEncodeWithParams(gocv.PNGFileExt, img, []int{gocv.IMWritePngCompression, act.compression})

	//gocv.IMEncodeWithParams(gocv.JPEGFileExt, img, []int{gocv.IMWriteJpegQuality, quality})

	base64String := base64.StdEncoding.EncodeToString(imageBuffer)

	output := &Output{}
	output.Image = imageBuffer
	output.Base64String = base64String
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
