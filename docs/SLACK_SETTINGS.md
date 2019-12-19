# Slack Settings
[notifier](https://github.com/bannzai/notifier) to use [Slack API](https://api.slack.com/) when post message with other service event to slack channel.
Follow the steps below.

1. Create [slack app](https://api.slack.com/apps)
1. [Configure OAuth Scopes](https://github.com/bannzai/notifier/tree/master/docs/SLACK_SETTINGS.md#Configure-OAuth-Scopes)
1. [Configure NOTIFIER_SLACK_TOKEN](https://github.com/bannzai/notifier/tree/master/docs/SLACK_SETTINGS.md#Configure-NOTIFIER_SLACK_TOKEN)

### Configure OAuth Scopes
In order to use Slack API, it is necessary to configure some **OAuth Scope** required for the [notifier](https://github.com/bannzai/notifier).

<details> 
<summary>
Guidance Slack App Display
</summary>
<div>
<p><img width=300px src="https://user-images.githubusercontent.com/10897361/71160131-57ace000-228a-11ea-885c-2529a82c9fd3.png" /></p>
<p><img width=300px src="https://user-images.githubusercontent.com/10897361/71159331-c6893980-2288-11ea-833a-b85553eb80e0.png" /></p>
</div>
</details>

#### Necessary OAuth Scopes
- [bot](https://api.slack.com/scopes/bot)
- [chat:write:bot](https://api.slack.com/scopes/chat:write:bot)
- [im:read](https://api.slack.com/scopes/im:read)
- [incoming-webhook](https://api.slack.com/scopes/incoming-webhook)
- [users:read](https://api.slack.com/scopes/users:read)


### Configure NOTIFIER_SLACK_TOKEN
In order to use [notifier](https://github.com/bannzai/notifier), it is necessary to set NOTIFIER_SLACK_TOKEN with Slack API OAuth Access Token.
You can copy OAuth Access Token from your slack app **OAuth & Permissions** pages.
When deploying [notifier](https://github.com/bannzai/notifier), set token and deploy it.

<details> 
<summary>
Guidance Slack App Display
</summary>
<div>
<p><img width=300px src="https://user-images.githubusercontent.com/10897361/71160131-57ace000-228a-11ea-885c-2529a82c9fd3.png" /></p>
<p><img width=300px src="https://user-images.githubusercontent.com/10897361/71160683-5def8c00-228b-11ea-972d-e3c1e0d4adc0.png" /></p>
</div>
</details>
