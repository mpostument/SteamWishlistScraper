package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

type Game struct {
	Name string `json:"name"`
}

func ScrapeWishlist(steamId string) []string{
	pageNumber := 0
	var gameList []string

	for {
		page, err := http.Get(fmt.Sprintf("https://store.steampowered.com/wishlist/profiles/%s/wishlistdata/?p=%d", steamId, pageNumber))
		if err != nil {
			log.Fatalln(err)
		}

		var result map[string]Game

		err = json.NewDecoder(page.Body).Decode(&result)
		if err != nil {
			res, _ := ioutil.ReadAll(page.Body)
			fmt.Println(res)
			if len(res) <= 2 {
				break
			}
			log.Fatalln(err)
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