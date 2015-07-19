## Features

Get IMDB ids from a public watchlist URL

## Example

```
package main

import (
	"fmt"
	"log"

	"github.com/odwrtw/imdb-watchlist"
)

func main() {
	wl := "http://www.imdb.com/user/ur27482023/watchlist"

	fmt.Println("Fetching watchlist:", wl)
	ids, err := imdbwatchlist.GetWatchlistIDs(wl)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(ids)
}
```
