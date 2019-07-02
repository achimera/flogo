# webcam
This activity provides your flogo application the ability to take pictures with a USB connected Webcam.
It is mandatory to specify the device Id, the image resolution as width and height and the compression factor. The image
has a png format.

**IMPORTANT**  
This activity uses the GoCV package that provides a binding for the OpenCV 4 computer vision library. Based
on the specific OS, you need to install OpenCV 4 on your system, otherwise the activity will throw an exception.

The installation instructions can be found here: https://github.com/hybridgroup/gocv#how-to-install

## Installation

```bash
flogo install activity github.com/achimera/flogo/activity/webcam
```

## Schema
Settings and Outputs:

```json
{
  "settings":[
    {
      "name": "deviceID",
      "type": "int",
      "required": true
	},
	{
      "name": "imageWidth",
      "type": "int",
      "required": true
	},
	{
      "name": "imageHeigth",
      "type": "int",
      "required": true
	},
	{
      "name": "compression",
      "type": "int",
      "required": true
    }
  ],
  "outputs": [
    {
      "name": "image",
      "type": "[]byte"
	},
	{
      "name": "base64String",
      "type": "string"
    },
  	{
      "name": "status",
      "type": "string"
    }
  ]
}
```
## Settings
| Name         | Type | Description    |
|:-------------|:-----|:---------------|        
| deviceId     | int  | The device id of the webcam. Usually the device id is 0 for onboard cameras |
| imageWidth   | int  | The resolution width of the image. Please check the resolutions that your webcam are supporting |
| imageHeigth  | int  | The resolution height of the image. Please check the resolutions that your webcam are supporting |
| compression  | int  | The compression factor of the png image. Value 0 is no compression, where 16 is the maximum compression and also the smallest size |

## Output
| Name      | Type   | Description    |
|:----------|:-------|:---------------|        
| image     | []byte | The captured image is returned as a byte array. No Base64 encoding is done. |
| base64String | string | The captured image is returned as a Base64 encoded string. |
| status    | string | The status of the operation. "OK" is returned when the webcam successfully captured an image | 

## Configuration Example

```json
		{
			"id": "webcam",
			"name": "Take a picture from a Webcam",
			"description": "Webcam Activity",
			"activity": {
				"ref": "github.com/achimera/flogo/activity/webcam",
				"settings": {
					"deviceID": "0",
					"imageWidth": "1024",
					"imageHeigth": "720",
					"compression": "10"
				}
			}
		}
```

