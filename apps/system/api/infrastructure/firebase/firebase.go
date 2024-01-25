package firebase

import (
	"context"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"firebase.google.com/go/v4/storage"
)

type Firebase struct {
	Auth    *auth.Client
	Storage *storage.Client
}

func setEmulatorEnv() error {
	if err := os.Setenv("FIREBASE_AUTH_EMULATOR_HOST", "host.docker.internal:9099"); err != nil {
		panic(err)
	}
	return os.Setenv("STORAGE_EMULATOR_HOST", "host.docker.internal:9199")
}

var firebaseLocalProjectID = "test"

func NewFirebase(isLocal bool, firebaseStorageBucket string) (inst *Firebase, err error) {
	fConf := &firebase.Config{
		StorageBucket: firebaseStorageBucket,
	}

	if isLocal {
		if err = setEmulatorEnv(); err != nil {
			return nil, err
		}
		fConf.ProjectID = firebaseLocalProjectID
	}

	ctx := context.Background()
	inst = new(Firebase)
	app, err := firebase.NewApp(ctx, fConf)
	if err != nil {
		return nil, err
	}

	authInst, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	storageClient, err := app.Storage(ctx)
	if err != nil {
		return nil, err
	}

	inst.Storage = storageClient
	inst.Auth = authInst
	return inst, nil
}
