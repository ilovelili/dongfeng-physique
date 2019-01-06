package handlers

import (
	"context"
	"fmt"

	proto "github.com/ilovelili/dongfeng-protobuf"
)

// SaveNotification save notification api marker
func (f *Facade) SaveNotification(ctx context.Context, req *proto.SaveNotificationRequest, rsp *proto.SaveNotificationResponse) error {
	return fmt.Errorf("not implemented")
}
