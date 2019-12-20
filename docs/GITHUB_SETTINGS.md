## GitHub Settings
This section explain about settings GitHub.
[notifier](https://github.com/bannzai/notifier) can be hook some GitHub events and notify to other service(e.g Slack).

1. Create [webhook](https://developer.github.com/webhooks/). And setting **Payload URL** to your hosted **notifier** application URL.
1. Select the event you want to be notified. See also [Current Support GitHub Event](https://github.com/bannzai/notifier/tree/master/docs/GITHUB_SETTINGS.md#Current-Support-GitHub-Event).
1. [Configure Yaml File](https://github.com/bannzai/notifier/tree/master/docs/NOTIFIER_YAML_SETTINGS.md)

### Current Support GitHub Event
- Issue comments
- Issues
- Pull requests
- Pull request reviews
- Pull request review comments

