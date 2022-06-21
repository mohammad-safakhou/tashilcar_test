package adapter

import (
	"tashil/handlers"
)

type notificationHandler struct {
}

func NewNotificationHandler() handlers.Notifications {
	return &notificationHandler{}
}

// RegisterNotification: registering our Notification rules and pushing scheduling
func (e *notificationHandler) SendPushNotification(message string, destination string) error {
	panic("implement me")
}
