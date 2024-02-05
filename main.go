package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/tawoe/obp-go-sdk"
	"log/slog"
	"obp-sdk-examples/config"
	"obp-sdk-examples/direct_login"
)

func main() {
	err := godotenv.Load()
	configuration := config.GetConfig()
	if err != nil {
		slog.Warn("Error loading .env file, relying on System Environment Variables")
	}
	var directLoginToken = direct_login.GetDirectLoginToken(configuration.Host, configuration.User, configuration.Password, configuration.ConsumerKey)
	client_conf := obp_golang.Configuration{
		BasePath:      configuration.Host,
		DefaultHeader: make(map[string]string),
		UserAgent:     "Swagger-Codegen/1.0.0/go",
	}
	client_conf.AddDefaultHeader("Authorization", fmt.Sprintf("DirectLogin token=%s", directLoginToken))
	client := obp_golang.NewAPIClient(&client_conf)
	user_api := client.UserApi
	user, _, err := user_api.GetCurrentUser(nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user.UserId)
}
