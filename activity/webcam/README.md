# webcam
This activity provides your flogo application the ability to take pictures with a USB connected Webcam camera.
It is mandatory to specify the filename. If also the path is specified, then it will create the directory structure if it doesn't exist.

Documentation of the RaspiCam can be found here: https://www.raspberrypi.org/documentation/raspbian/applications/camera.md

## Installation

```bash
flogo add activity github.com/achimera/flogo/activity/webcam
```

## Schema
Inputs and Outputs:

```json
{
  "inputs":[
    {
      "name": "timeout",
      "type": "int",
      "required": false
    },
    {
      "name": "sharpness",
      "type": "int",
      "required": false
    },
    {
      "name": "brightness",
      "type": "int",
      "required": false
    },
    {
      "name": "contrast",
      "type": "int",
      "required": false
    },
    {
      "name": "saturation",
      "type": "int",
      "required": false
    },
    {
      "name": "iso",
      "type": "int",
      "required": false
    },
    {
      "name": "filename",
      "type": "string",
      "required": true
    }
  ],
  "outputs": [
  	{
      "name": "status",
      "type": "string"
    }
  ]
}
```
## Settings
| Setting      | Description    |
|:-------------|:---------------|        
| timeout   | The Pushbullet access token allocated for your app |
| sharpness | The title of the link |
| linkMsg      	 | The message associated with the link to send |
| linkUrl      	 | The URL to open |
| emailTarget    	| The email address where to send the link. Only one target is allowed |
| channelTarget    	| The channel name where to send the link. Only one target is allowed |

If emailTarget and channelTarget are empty, then the link is sent to all devices. 

In the 'status' output, you may get the following values:
- 'OK' : the link was correctly sent
- 'PUSH_ERR' : an error on sending the link via Pushbullet
- 'CONNECT_ERR' : if there was an error connecting to Pushbullet
- 'NO_LINK_URL_ERR' : if the link URL field is empty
- 'TOO_MANY_TARGETS_ERR' : if multiple targets are specified


## Configuration Example

```json
            {  
            	"id": 2,
            	"name": "Pushbullet Link Notification",
            	"type":1,
            	"activityType":"pushbulletlink",
            	"attributes":[  
    				{
      					"name": "accessToken",
      					"value": "YOUR_ACCESS_TOKEN",
      					"type": "string"
            },
            {
      					"name": "linkTitle",
      					"value": "YOUR_LINK_TITLE",
      					"type": "string"
    				},
    				{
      					"name": "linkMsg",
      					"value": "YOUR_LINK_MESSAGE",
      					"type": "string"
    				},
    			  {
      					"name": "linkUrl",
      					"value": "YOUR_URL_TO_OPEN",
      					"type": "string"
    				}
            	]
         	},
```

