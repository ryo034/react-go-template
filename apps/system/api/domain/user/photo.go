package user

import "github.com/ryo034/react-go-template/apps/system/api/domain/shared/media"

type Photo struct {
	filePath media.Path
}

func NewPhotoFromString(path string) (Photo, error) {
	p, err := media.NewPath(path)
	if err != nil {
		return Photo{}, err
	}
	return Photo{p}, nil
}

func (p *Photo) FilePath() media.Path {
	return p.filePath
}
