package helper

import "strings"

func ChannelMentionToChannelId(channel string) string {
	channelId := strings.Replace(strings.Replace(channel, "<#", "", -1), ">", "", -1)
	return channelId
}
