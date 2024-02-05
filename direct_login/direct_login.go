package direct_login

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type DirectLoginToken struct {
	Token string `json:"token"`
}

func GetDirectLoginToken(obpApiHost, user, password, consumerKey string) string {
	request, _ := http.NewRequest("POST", obpApiHost+"/my/logins/direct", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "DirectLogin username=\""+user+"\",password=\""+password+"\",consumer_key=\""+consumerKey+"\"")
	response, err := http.DefaultClient.Do(request)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var directLogin DirectLoginToken
	json.Unmarshal(responseData, &directLogin)
	return string(directLogin.Token)
}
