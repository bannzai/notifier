package sender

import (
	"fmt"

	"github.com/bannzai/notifier/pkg/logger"
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
	ids, err := sender.Mapper.MapIDs(content, SlackContentType)
	for _, slackUserID := range ids {
		if err != nil {
			return errors.Wrapf(err, "Slack sender.Mapper.MapID error with %v", content)
		}

		user, err := sender.GetUserInfo(slackUserID)
		if err != nil {
			return errors.Wrapf(err, "sender.GetUserInfo(%s) is error", slackUserID)
		}

		logger.Logf("start post message to slack from github. for user name is %s and user id is %s", user.Name, user.ID)

		dmChannel, err := sender.getDirectMessageChannelID(slackUserID)
		if err != nil {
			return errors.Wrapf(err, "sender.GetDirectMessageChannelID(%s) is error", slackUserID)
		}

		messageOption := slack.MsgOptionText(sender.buildContent(content), false)
		responseChannel, responseTimestamp, err := sender.PostMessage(dmChannel, messageOption)
		if err != nil {
			return errors.Wrapf(err, "sender.postmessage(%s) is error", slackUserID)
		}
		logger.Logf("Postmessage channel id: %s, timestamp: %s", responseChannel, responseTimestamp)
	}
	if err != nil {
		return errors.Wrapf(err, "slack.Send is error. but this ids: %v alrady post message", ids)
	}
	return nil
}

func (sender Slack) buildContent(content parser.Content) string {
	switch content.ContentType {
	case parser.GitHubMentionContent:
		return fmt.Sprintf("You got mention from: %s", content.LinkURL)
	case parser.GitHubAssignedContent:
		return fmt.Sprintf("You got contact for assignee this PR: %s", content.LinkURL)
	case parser.GitHubRequestReviewedContent:
		return fmt.Sprintf("You got request for reviewer this PR: %s", content.LinkURL)
	default:
		return fmt.Sprintf("Unknown event %s", content.LinkURL)
	}
}

func (sender Slack) getDirectMessageChannelID(userID string) (string, error) {
	noop, alreadyOpen, channelID, err := sender.OpenIMChannel(userID)
	if err != nil {
		return "", errors.Wrapf(err, "sender.OpenIMChannel(%s) is error", userID)
	}
	fmt.Printf("open id status noop = %+v, already open = %v, channelID = %v\n", noop, alreadyOpen, channelID)
	return channelID, nil
}
