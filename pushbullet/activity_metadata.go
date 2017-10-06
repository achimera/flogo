package pushbullet

var jsonMetadata = `{
  "name": "pushbullet",
  "version": "0.0.1",
  "description": "Simple Pushbullet Notification Activity",
  "author": "Alessandro Chimera <achimera@tibco.com>",
  "inputs":[
    {
      "name": "accToken",
      "type": "string"
    },
    {
      "name": "message",
      "type": "string"
    },
    {
      "name": "messageTitle",
      "type": "string"
    }
  ],
  "outputs": [
  	{
      "name": "status",
      "type": "string"
    }
  ]
}`