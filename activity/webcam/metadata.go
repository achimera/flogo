package webcam

import (
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
	DeviceID    string `md:"deviceID, required"`    // The webcam deviceId
	ImageWidth  int    `md:"imageWidth, required"`  // The image resolution width
	ImageHeigth int    `md:"imageHeigth, required"` // The image resolution heigth
	Compression int    `md:"compression, required"` // The image compression factor
}

/*
type Input struct {
	in string `md:"in"`
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"in": i.in,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.in, err = coerce.ToString(values["in"])
	if err != nil {
		return err
	}
	return nil
}
*/

type Output struct {
	Image        []byte `md:"image"`        // The byte array containing the image
	Base64String string `md:"base64String"` // The Base64 encoded byte array returned as a string
	Status       string `md:"status"`       // Status of the webcam
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"image":        o.Image,
		"base64String": o.Base64String,
		"status":       o.Status,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {

	var err error

	o.Image, err = coerce.ToBytes(values["image"])
	if err != nil {
		return err
	}

	o.Base64String, err = coerce.ToString(values["base64String"])
	if err != nil {
		return err
	}

	o.Status, err = coerce.ToString(values["status"])
	if err != nil {
		return err
	}

	return nil
}
