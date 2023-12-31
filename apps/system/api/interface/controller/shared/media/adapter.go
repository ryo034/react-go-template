package media

import (
	mediaPb "github.com/ryo034/react-go-template/apps/system/api/schema/pb/shared/media/v1"
)

type Adapter interface {
}

type adapter struct {
	isLocal bool
}

func NewAdapter() Adapter {
	return &adapter{}
}

func (a *adapter) isVideo(t mediaPb.MediaType) bool {
	switch t {
	case mediaPb.MediaType_MEDIA_TYPE_MP4:
		return true
	case mediaPb.MediaType_MEDIA_TYPE_MOV:
		return true
	}
	return false
}
