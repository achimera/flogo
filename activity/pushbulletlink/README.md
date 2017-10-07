# pushbulletlink
This activity provides your flogo application the ability to send a link via Pushbullet.
If no email and channel target is specified, then the link is sent to all devices. 


## Installation

```bash
flogo add activity github.com/achimera/flogo/activity/pushbulletlink
```

## Schema
Inputs and Outputs:

```json
{
  "inputs":[
    {
      "name": "accessToken",
      "type": "string",
      "required": true
    },
    {
      "name": "linkTitle",
      "type": "string",
      "required": true
    },
    {
      "name": "linkMsg",
      "type": "string",
      "required": true
    },
    {
      "name": "linkUrl",
      "type": "string",
      "required": true
    },
    {
      "name": "emailTarget",
      "type": "string",
      "required": false
    },
    {
      "name": "channelTarget",
      "type": "string",
      "required": false
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
| accessToken  | The Pushbullet access token allocated for your app |
| linkTitle | The title of the link |
| linkMsg      	 | The message associated with the link to send |
| linkUrl      	 | The URL to open |
| emailTarget    	| The email address where to send the link. Only one target is allowed |
| channelTarget    	| The channel name where to send the link. Only one target is allowed |

If emailTarget and channelTarget are empty, then the link is send to all devices. 

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

