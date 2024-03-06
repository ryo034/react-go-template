package media

import (
	"fmt"
	"io"

	"github.com/ryo034/react-go-template/apps/system/api/domain/shared/account"

	minio2 "github.com/ryo034/react-go-template/apps/system/api/infrastructure/storage/minio"

	"github.com/minio/minio-go/v7"
	"golang.org/x/net/context"
)

type Driver interface {
	UploadAvatar(ctx context.Context, aID account.ID, fileName string, body io.Reader, objectSize int64) error
}

type driver struct {
	minioClient *minio.Client
	conf        *minio2.Config
}

const AccountAvatarFolder = "avatar"
const AccountsFolder = "accounts"

func NewDriver(minioClient *minio.Client, conf *minio2.Config) Driver {
	return &driver{minioClient, conf}
}

func (d *driver) UploadAvatar(ctx context.Context, aID account.ID, fileName string, body io.Reader, objectSize int64) error {
	filePath := fmt.Sprintf("/%s/%s/%s/%s", AccountsFolder, aID.ToString(), AccountAvatarFolder, fileName)
	_, err := d.minioClient.PutObject(ctx, d.conf.BucketName, filePath, body, objectSize, minio.PutObjectOptions{})
	return err
}
