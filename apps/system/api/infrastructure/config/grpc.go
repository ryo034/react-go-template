package config

const (
	grpcPort Key = "PORT"
)

func (r *reader) ServerPort() string {
	var port = r.fromEnv(grpcPort)
	if port == "" {
		port = "8080"
	}
	return port
}
