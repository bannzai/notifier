## NOTIFIER_YAML_SETTINGS
**notifier** necessary yaml file about defined cooperation service ids.

### Example
```yaml
id: owata
github:
  login: owata
slack:
  id: YYYYYYYYY
id: bannzai
github:
  login: bannzai
slack:
  id: XXXXXXXXX
```
Top level `id` is free word. It should be set unique name.

## GitHub
GitHub settings should configure `github.login`.


## Slack
Slack settings should configure `slack.id`. You can get user identifier for slack on workspace from Slack API tester for [users.list](https://api.slack.com/methods/users.list/test).
