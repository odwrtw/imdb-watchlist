package watchlist

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

var imdbIDRegexp *regexp.Regexp

func init() {
	imdbIDRegexp = regexp.MustCompile("tt\\d{7}")
}

// GetWatchlistIDs returns the imdb ids included in a public watchlist
func GetWatchlistIDs(watchlist string) (*[]string, error) {
	resp, err := http.Get(watchlist)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	ids := []string{}
	matches := imdbIDRegexp.FindAll(body, -1)
	if matches != nil {
		movieIDs := map[string]bool{}
		for _, m := range matches {
			movieIDs[string(m)] = true
		}
		for id := range movieIDs {
			ids = append(ids, id)
		}
	}

	return &ids, nil
}
