# pushbulletnote
This activity provides your flogo application the ability to send a notification via Pushbullet.
If no email and channel target is specified, then the note is sent to all devices. 


## Installation

```bash
flogo add activity github.com/achimera/flogo/activity/pushbulletnote
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
      "name": "note",
      "type": "string",
      "required": true
    },
    {
      "name": "noteTitle",
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
| note      	 | The note to send |
| noteTitle | The title of the note |
| emailTarget    	| The email address where to send the note. Only one target is allowed |
| channelTarget    	| The channel name where to send the note. Only one target is allowed |

If emailTarget and channelTarget are empty, then the note is sent to all devices. 

In the 'status' output, you may get the following values:
- 'OK' : the note was correctly sent
- 'PUSH_ERR' : an error on sending the note via Pushbullet
- 'CONNECT_ERR' : if there was an error connecting to Pushbullet
- 'NO_NOTE_ERR' : if the input 'note' field is empty
- 'TOO_MANY_TARGETS_ERR' : if multiple targets are specified


## Configuration Example

```json
            {  
            	"id": 2,
            	"name": "Pushbullet Note Notification",
            	"type":1,
            	"activityType":"pushbulletnote",
            	"attributes":[  
    				{
      					"name": "accessToken",
      					"value": "YOUR_ACCESS_TOKEN",
      					"type": "string"
    				},
    				{
      					"name": "note",
      					"value": "YOUR_NOTE",
      					"type": "string"
    				},
    				{
      					"name": "noteTitle",
      					"value": "YOUR NOTE TITLE",
      					"type": "string"
    				}
            	]
         	},
```

