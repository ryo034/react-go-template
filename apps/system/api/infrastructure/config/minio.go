package config

import (
	"github.com/ryo034/react-go-template/apps/system/api/infrastructure/storage/minio"
)

const (
	storageHost       Key = "STORAGE_HOST"
	storageEndpoint   Key = "STORAGE_ENDPOINT"
	storageAccessKey  Key = "STORAGE_ACCESS_KEY"
	storageSecretKey  Key = "STORAGE_SECRET"
	storageBucketName Key = "STORAGE_BUCKET_NAME"
)

func (r *reader) MinioConfig() *minio.Config {
	return &minio.Config{
		Host:            r.fromEnv(storageHost),
		Endpoint:        r.fromEnv(storageEndpoint),
		AccessKeyID:     r.fromEnv(storageAccessKey),
		SecretAccessKey: r.fromEnv(storageSecretKey),
		BucketName:      r.fromEnv(storageBucketName),
		UseSSL:          r.IsNotLocal(),
	}
}
