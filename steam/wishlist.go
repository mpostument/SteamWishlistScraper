package steam

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strconv"
)

type Game struct {
	Name string `json:"name"`
}

func ScrapeWishlist(steamId string) []string {
	pageNumber := 0
	var gameList []string

	for {
		baseURL := fmt.Sprintf("https://store.steampowered.com/wishlist/profiles/%s/wishlistdata/", steamId)
		u, err := url.Parse(baseURL)
		if err != nil {
			log.Fatalln("Not able to pars url", err)
		}
		queryString := u.Query()
		queryString.Set("p", strconv.Itoa(pageNumber))
		u.RawQuery = queryString.Encode()
		page, err := http.Get(u.String())
		if err != nil {
			log.Fatalln("Didn't get responce from wishlist page", err)
		}

		var result map[string]Game

		err = json.NewDecoder(page.Body).Decode(&result)
		if err != nil {
			res, _ := ioutil.ReadAll(page.Body)
			if len(res) <= 2 {
				break
			}
			log.Fatalln("Wishlist decoding failed", err)
		}
		page.Body.Close()
		for _, v := range result {
			gameList = append(gameList, v.Name)
		}
		pageNumber++
	}
	sort.Strings(gameList)
	return gameList
}
