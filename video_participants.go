package twilio

import (
	"context"
	"fmt"
	"net/url"
)

type VideoParticipantService struct {
	client *Client
}

const videoParticipantsPathPart = "Rooms/%s/Participants"

type VideoParticipant struct {
	Sid         string            `json:"sid"`
	Duration    uint              `json:"duration"`
	Status      Status            `json:"status"`
	DateCreated TwilioTime        `json:"date_created"`
	EndTime     TwilioTime        `json:"end_time"`
	RoomSid     string            `json:"room_sid"`
	URL         string            `json:"url"`
	Size        uint              `json:"size"`
	Identity    string            `json:"identity"`
	Links       map[string]string `json:"links"`
}

type VideoParticipantPage struct {
	Meta         Meta                `json:"meta"`
	Participants []*VideoParticipant `json:"participants"`
}

type VideoParticipantPageIterator struct {
	p *PageIterator
}

// GetPageForRoom Returns a list of participants for a given room. For more information on valid values,
// see https://www.twilio.com/docs/api/video/participants#participant-list-resource
func (vr *VideoParticipantService) GetPageForRoom(ctx context.Context, roomSid string, data url.Values) (*VideoParticipantPage, error) {
	return vr.GetPageIterator(roomSid, data).Next(ctx)
}

// GetPageIterator returns an iterator which can be used to retrieve pages.
func (vr *VideoParticipantService) GetPageIterator(roomSid string, data url.Values) *VideoParticipantPageIterator {
	path := fmt.Sprintf(videoParticipantsPathPart, roomSid)
	iter := NewPageIterator(vr.client, data, path)
	return &VideoParticipantPageIterator{
		p: iter,
	}
}

// Next returns the next page of resources. If there are no more resources,
// NoMoreResults is returned.
func (vr *VideoParticipantPageIterator) Next(ctx context.Context) (*VideoParticipantPage, error) {
	vrp := new(VideoParticipantPage)
	err := vr.p.Next(ctx, vrp)
	if err != nil {
		return nil, err
	}
	vr.p.SetNextPageURI(vrp.Meta.NextPageURL)
	return vrp, nil
}
