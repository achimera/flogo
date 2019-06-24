package webcam

import (
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
	deviceID string `md:"deviceID,required"` // The webcam deviceId

}
type Input struct {
}

func (i *Input) ToMap() map[string]interface{} {
	return nil
}

func (i *Input) FromMap(values map[string]interface{}) error {

	return nil
}

type Output struct {
	Image  []byte `md:"image"`  // Documents the partition that the message was placed on
	Status string `md:"status"` // Status of the webcam
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"image":  o.Image,
		"status": o.Status,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {

	var err error
	o.Image, err = coerce.ToBytes(values["image"])
	if err != nil {
		return err
	}

	o.Status, err = coerce.ToString(values["status"])
	if err != nil {
		return err
	}

	return nil
}
