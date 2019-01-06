package clients

import (
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// SaveNotification save notification
func (c *Client) SaveNotification(req *proto.SaveNotificationRequest) (*proto.SaveNotificationResponse, error) {
	return c.client.SaveNotification(ctx(), req)
}
