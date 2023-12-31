package config

var localValues = map[Key]string{
	// Database Setting
	dbUser:        "root",
	dbPass:        "password",
	dbName:        "main",
	dbSourceHost:  "host.docker.internal",
	dbSourcePort:  "15432",
	dbReplicaHost: "host.docker.internal",
	dbReplicaPort: "25432",
	// gRPC
	grpcPort: "19004",
	// firebase
	firebaseStorageBucket: "",
	// cors
	allowOrigin: "http://localhost:5173,http://localhost",
}
