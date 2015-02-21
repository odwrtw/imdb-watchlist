## Features

Get IMDB ids from a public watchlist URL

## Example

```
package main

import (
	"fmt"
	"log"

	"gitlab.quimbo.fr/odwrtw/imdb-watchlist"
)

func main() {
	wl := "http://www.imdb.com/user/ur27482023/watchlist"

	fmt.Println("Fetching watchlist:", wl)
	ids, err := watchlist.GetWatchlistIDs(wl)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(ids)
}
```
