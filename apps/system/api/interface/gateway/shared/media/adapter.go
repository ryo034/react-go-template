package media

import (
	"github.com/ryo034/react-go-template/packages/go/domain/shared/media"
	models "github.com/ryo034/react-go-template/packages/go/infrastructure/database/sqlboiler/api"
)

type Adapter interface {
	AdaptPhoto(m *models.Photo, order uint32) (media.Media, error)
}

type adapter struct {
}

func NewAdapter() Adapter {
	return &adapter{}
}

func (a *adapter) AdaptPhoto(m *models.Photo, order uint32) (media.Media, error) {
	pID, err := media.NewID(m.PhotoID)
	if err != nil {
		return nil, err
	}
	fp, err := media.NewPath(m.PhotoPath)
	return media.NewPhoto(pID, m.Title, m.Description, fp, m.Width, m.Height, order), nil
}
