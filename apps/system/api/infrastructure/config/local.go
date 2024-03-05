package config

var localValues = map[Key]string{
	serviceName: "system",
	// Database Setting
	dbUser:        "root",
	dbPass:        "password",
	dbName:        "main",
	dbSourceHost:  "host.docker.internal",
	dbSourcePort:  "15432",
	dbReplicaHost: "host.docker.internal",
	dbReplicaPort: "25432",
	// gRPC
	serverPort: "19004",
	// firebase
	firebaseStorageBucket: "",
	// cors
	allowOrigin: "http://localhost:5173,http://localhost",
	// redis
	redisAddr: "host.docker.internal:6379",
	redisDB:   "0",
	redisPass: "root",
	// mail
	mailHost: "mailhog",
	mailPort: "1025",
	// email
	noReplyEmail: "no-reply@example.com",
	// resend api key
	resendAPIKey: "",
	//	storage
	storageHost:       "localhost:9090",
	storageEndpoint:   "host.docker.internal:9090",
	storageAccessKey:  "minio",
	storageSecretKey:  "minio123",
	storageBucketName: "system",
}
