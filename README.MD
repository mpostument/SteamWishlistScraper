# SteamWishlistScraper

Scrape your steam wishlist.

## Table of Contents

- [SteamWishlistScraper](#steamwishlistscraper)
  - [Table of Contents](#table-of-contents)
  - [Installing](#installing)
  - [Getting Started](#getting-started)
    - [Scrape wishlist](#scrape-wishlist)
  - [Global parameters](#global-parameters)
  - [Contributing](#contributing)
  - [License](#license)

## Installing

Download the latest binary from [releases](https://github.com/mpostument/SteamWishlistScraper)

## Getting Started

### Scrape wishlist

1. Generate api key on [steam](https://steamcommunity.com/dev/apikey)
2. Run SteamWishlistScraper
   `SteamWishlistScraper --username <SteamUserName> --apikey <ApiKey>`

Note: Api key can be added to $HOME/.SteamWishlistScraper.yaml as `apikey: <ApiKey>`

## Global parameters

`username` - Steam user name. Default `""`  
`apikey` - Steam api key. Default `""`  

## Contributing

1. Fork it
2. Download your fork to your PC ( `git clone https://github.com/mpostument/SteamWishlistScraper && cd SteamWishlistScraper` )
3. Create your feature branch ( `git checkout -b my-new-feature` )
4. Make changes and add them ( `git add .` )
5. Commit your changes ( `git commit -m 'Add some feature'` )
6. Push to the branch ( `git push origin my-new-feature` )
7. Create new pull request

## License

SteamWishlistScraper is released under the Apache 2.0 license. See [LICENSE.txt](https://github.com/mpostument/SteamWishlistScraper/blob/master/LICENSE)