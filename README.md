# notifier
**notifier** connect to Slack when any actions on GitHub

## Install
Install with git clone.
```shell
$ git clone git@github.com:bannzai/notifier.git
```

Or install via go get
```shell
$ go get -u github.com/bannzai/notifier
```

## Deploy
### Settings
Setting up other services is necessary to run **notifier** .
Look below.

- [Slack Settings Document](https://github.com/bannzai/notifier/tree/master/docs/SLACK_SETTINGS.md)
- [GitHub Settings Document](https://github.com/bannzai/notifier/tree/master/docs/GITHUB_SETTINGS.md)

### Recommended
notifier prepared heroku configurations.
If you have heroku account and already login, you can exec make heroku command and deploy it. 

```
$ make heroku
```

### Other
notifier prepared Deockerfile.
It mean about run on Docker anywhere.

### Environment
notifier is necessary some environment variables

|  Key  |  Description  | 
| ---- | ---- | 
|  YAML_FILE_PATH  |  Yaml file path about relation of user for GitHub and Slack accounts. Bot local file path or remote file path(http or https)| 
|  NOTIFIER_SLACK_TOKEN  |  notifier used Slack API Token. See also https://api.slack.com/ | 


## LICENSE
**notifiier** is available under the MIT license. See the LICENSE file for more info.

