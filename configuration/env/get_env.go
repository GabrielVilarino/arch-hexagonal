package env

import "os"

func GetAmbiente() bool {
	return os.Getenv("AMBIENTE") == "PROD"
}

func GetTokenApi() string {
	return os.Getenv("NEWS_API_KEY")
}

func GetPort() string {
	return os.Getenv("PORT")
}
