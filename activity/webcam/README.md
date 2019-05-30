# webcam
This activity provides your flogo application the ability to take pictures with a USB connected Webcam.
It is mandatory to specify the filename. If also the path is specified, then it will create the directory structure if it doesn't exist.

This activity uses the GoCV package that provides a binding for the OpenCV 4 computer vision library. Based
on the specific OS, you need to install OpenCV 4 on your system. 

The installation instructions can be found here: https://github.com/hybridgroup/gocv#how-to-install

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
      "name": "filename",
      "type": "string",
      "required": true
    }
  ],
  "outputs": [
    {
      "name": "image",
      "type": "[]byte"
    },
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
| filename   | The Pushbullet access token allocated for your app |


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

