// Package handlers define the core behaviors of each API
package handlers

import (
	notification "github.com/ilovelili/dongfeng-notification"
	"github.com/ilovelili/dongfeng-physique/services/server/core/controllers"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// Facade api facade
type Facade struct{}

// NewFacade constructor
func NewFacade() *Facade {
	return new(Facade)
}

// syslog save notification
func (f *Facade) syslog(notification *notification.Notification) {
	go func() {
		notificationcontroller := controllers.NewNotificationController()
		notificationcontroller.Save(&proto.Notification{
			UserId:     notification.UserID,
			CustomCode: notification.CustomCode,
			Details:    notification.Details,
			Link:       notification.Link,
			CategoryId: int64(notification.CategoryID),
		})
	}()
}
