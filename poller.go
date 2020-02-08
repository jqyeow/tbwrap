package tbwrap

import (
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func NewPollerWithAllowedUserAndGroups(pollTimout time.Duration, allowedUsers []int, allowedGroups []int) *tb.MiddlewarePoller {
	poller := &tb.LongPoller{Timeout: pollTimout}

	return tb.NewMiddlewarePoller(poller, func(upd *tb.Update) bool {
		allowedUsersAndGroups := append(allowedUsers, allowedGroups...)

		// allow request if no restrictions are set
		if len(allowedUsersAndGroups) == 0 {
			return true
		}

		if upd.Message != nil {
			return isInList(int(upd.Message.Chat.ID), allowedUsersAndGroups)
		}

		if upd.Callback != nil {
			return isInList(int(upd.Callback.Message.Chat.ID), allowedUsersAndGroups)
		}

		return false
	})
}

func isInList(ID int, list []int) bool {
	for i := range list {
		if list[i] == ID {
			return true
		}
	}
	return false
}
