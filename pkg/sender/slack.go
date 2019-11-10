package sender

import (
	"fmt"

	"github.com/bannzai/notifier/pkg/parser"
	"github.com/nlopes/slack"
	"github.com/pkg/errors"
)

type Slack struct {
	*slack.Client
	Mapper
}

func NewSlack(apiToken string, mapper Mapper) Slack {
	return Slack{
		Client: slack.New(apiToken),
		Mapper: mapper,
	}
}

func (sender Slack) Send(content parser.Content) error {
	slackUserID, err := sender.Mapper.MapID(content, parser.SlackContent)
	if err != nil {
		return errors.Wrapf(err, "Slack sender.Mapper.MapID error with %s", content)
	}

	user, err := sender.GetUserInfo(slackUserID)
	if err != nil {
		return errors.Wrapf(err, "sender.GetUserInfo(%s) is error", slackUserID)
	}

	fmt.Printf("start post message to slack from github. for user name is %s and user id is %s", user.Name, user.ID)

	dmChannel, err := sender.getDirectMessageChannelID(slackUserID)
	if err != nil {
		return errors.Wrapf(err, "sender.GetDirectMessageChannelID(%s) is error", slackUserID)
	}

	messageOption := slack.MsgOptionText(sender.buildContent(content), false)
	responseChannel, responseTimestamp, err := sender.PostMessage(dmChannel, messageOption)
	if err != nil {
		return errors.Wrapf(err, "sender.postmessage(%s) is error", slackUserID)
	}
	fmt.Printf("Postmessage channel id: %s, timestamp: %s", responseChannel, responseTimestamp)
	return nil
}

func (sender Slack) buildContent(content parser.Content) string {
	switch content.ContentType {
	case parser.GitHubMentionContent:
		return fmt.Sprintf("You got mention from: %s", content.LinkURL)
	case parser.GitHubAssignedContent:
		return fmt.Sprintf("You assigned this PR: %s", content.LinkURL)
	default:
		return fmt.Sprintf("Unknown event %s", content.LinkURL)
	}
}

func (sender Slack) getDirectMessageChannelID(userID string) (string, error) {
	imList, err := sender.GetIMChannels()
	if err != nil {
		return "", errors.Wrap(err, "sender.GetIMChannels() is error")
	}
	for _, im := range imList {
		if im.User == userID {
			return im.ID, nil
		}
	}
	return "", fmt.Errorf("Can not find im for this user id %s", userID)
}
