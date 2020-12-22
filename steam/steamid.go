package steam

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
	baseURL := "http://api.steampowered.com/ISteamUser/ResolveVanityURL/v0001/"
	u, err := url.Parse(baseURL)
	if err != nil {
		log.Fatalln("Not able to parse url", err)
	}
	queryString := u.Query()
	queryString.Set("key", apiKey)
	queryString.Set("vanityurl", userName)
	u.RawQuery = queryString.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		log.Fatalln("Didn't get response from steam", err)
	}
	defer resp.Body.Close()
	var result User

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatalln("Decoding failed", err)
	}
	return result.Response.Steamid
}
