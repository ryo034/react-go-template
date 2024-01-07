package config

const (
	serverPort Key = "PORT"
)

func (r *reader) ServerPort() string {
	var port = r.fromEnv(serverPort)
	if port == "" {
		port = "8080"
	}
	return port
}
