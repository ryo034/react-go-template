package media

import (
	"io"

	minio2 "github.com/ryo034/react-go-template/apps/system/api/infrastructure/storage/minio"

	"github.com/minio/minio-go/v7"
	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/media"
	"golang.org/x/net/context"
)

type Driver interface {
	UploadAvatar(ctx context.Context, id media.ID, body io.Reader) error
}

type driver struct {
	minioClient *minio.Client
	conf        *minio2.Config
}

func NewDriver(minioClient *minio.Client, conf *minio2.Config) Driver {
	return &driver{minioClient, conf}
}

func (d *driver) UploadAvatar(ctx context.Context, id media.ID, body io.Reader) error {
	_, err := d.minioClient.PutObject(ctx, d.conf.BucketName, id.String(), body, -1, minio.PutObjectOptions{})
	return err
}
