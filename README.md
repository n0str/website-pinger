# Website pinger

A simple tool to monitor website uptime status with advanced rules set and various notification methods.

# Installation

Clone Git repository and navigate to the source directory
```
git clone https://github.com/n0str/website-pinger.git
cd website_pinger
```

## Build executable

Build single executable
```
go build -o pinger
```

You can also build executable for Linux
```
GOOS=linux go build -o pinger
```

## Prepare directories
* logs - directory for error logging
* rules - databases with URL configuration parameters
```
mkdir logs
mkdir rules
```

## Run

Run server with default settings
```
chmod +x pinger
./pinger
```

### Arguments
Currently support the following arguments

|argument|default value|description
|--|--|--
|`--port`|`8080`|The server listening port
|`--max_queue_size`|`100`|The maximum queue size
|`--max_workers`|`5`|Number of workers

# API

You can control settings with REST API.

|URL|METHOD|description
|--|--|--
|`/api/set`|POST| Add and modify rules
|`/api/get`|POST| Get rule by URL
|`/api/list`|GET| Show all URLs
|`/api/delete`|DELETE| Delete rule by URL
|`/api/reload`|GET| Reload database without restart

## Add and modify rules

|Parameter|Type|description
|--|--|--
|`url`|`string`|URL
|`status_code`|`integer`|Desired valid HTTP status code
|`informer_type`|`integer`|Unique informer code _(see in the next section)_
|`informer_payload`|`string`|Payload for informer

### Example
```
curl -X POST http://localhost:8080/api/set --data "url=http://ifmo.su&status_code=200&informer_type=0&informer_payload=code"
```

### Informers

Currently we support only [CodeX Bot](https://ifmo.su/bot) informer. Notifications are made to Telegram according to the documentation: [https://github.com/codex-bot/Webhooks](https://github.com/codex-bot/Webhooks)

|Informer code|name|description
|--|--|--
|`0`|`codex bot`|CodeX Bot Notifications

## Get rule by URL

Get information about the URL specified

|Parameter|Type|description
|--|--|--
|`url`|`string`|URL

Response example
```
{
	"Url": "https://ifmo.su",
	"DesiredStatusCode": 200,
	"InformerPayload": {
		"Type": 0,
		"Payload": "..."
	}
}
```

## Show all URLs

Show URLs which are on monitoring

|Parameter|Type|description
|--|--|--
|no params|

Response example
```
[
	"https://hawk.so",
	"https://ifmo.su"
]
```

## Delete rule by URL

Delete the specified URL from monitoring

|Parameter|Type|description
|--|--|--
|`url`|`string`|URL

## Reload database without restart

Reload rules from files in the `rules` directory. Can be useful if you update the settings manually.

|Parameter|Type|description
|--|--|--
|no params|

# Monitoring

The bot will send you notifications according to the informer. For example with CodeX Bot it will be like:

![](https://capella.pics/eb1c22ae-2bbe-42ca-b170-d4f5274ae130.jpg)



