package pkg

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type User struct {
	Response `json:"response"`
}

type Response struct {
	Steamid string `json:"steamid"`
	Success int    `json:"success"`
}

func GetSteamId(userName string, apiKey string) string {
	baseUrl := "http://api.steampowered.com/ISteamUser/ResolveVanityURL/v0001/"
	u, err := url.Parse(baseUrl)
	if err != nil {
		log.Fatalln(err)
	}
	queryString := u.Query()
	queryString.Set("key", apiKey)
	queryString.Set("vanityurl", userName)
	u.RawQuery = queryString.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	var result User

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatalln(err)
	}
	return result.Response.Steamid
}
