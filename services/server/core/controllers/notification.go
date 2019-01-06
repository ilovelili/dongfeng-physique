package controllers

import (
	"github.com/ilovelili/dongfeng-physique/services/server/core/clients"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// NotificationController notification controller
type NotificationController struct {
	client *clients.Client
}

// NewNotificationController new controller
func NewNotificationController() *NotificationController {
	return &NotificationController{
		client: clients.New(),
	}
}

// Save save Notifications
func (c *NotificationController) Save(notification *proto.Notification) error {
	req := &proto.SaveNotificationRequest{
		Notification: notification,
	}

	_, err := c.client.SaveNotification(req)
	return err
}
