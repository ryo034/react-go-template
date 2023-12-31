package config

const (
	firebaseStorageBucket Key = "FIREBASE_STORAGE_BUCKET"
)

func (r *reader) FirebaseStorageBucket() string {
	return r.fromEnv(firebaseStorageBucket)
}
