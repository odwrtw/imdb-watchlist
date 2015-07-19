package imdbwatchlist

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
)

var (
	imdbIDRegexp = regexp.MustCompile("tt\\d{7}")
	imdbUserRe   = regexp.MustCompile("ur\\d{8}") // DEPRECIATED
	baseURL      = "http://www.imdb.com/user/"
)

func unique(strs *[]string) {
	returnslice := []string{}
	t := map[string]bool{}
	for _, s := range *strs {
		t[s] = true
	}
	for s := range t {
		returnslice = append(returnslice, s)
	}
	*strs = returnslice
}

func getIds(userid string, filter string) (*[]string, error) {
	URL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	URL.Path += userid
	URL.Path += "/watchlist"

	parameters := url.Values{}
	parameters.Add("title_type", filter)
	URL.RawQuery = parameters.Encode()

	resp, err := http.Get(URL.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	matches := imdbIDRegexp.FindAllString(string(body), -1)
	if matches != nil {
		unique(&matches)
		return &matches, nil
	}
	return nil, nil
}

// GetMovies from userid return slice of ids
func GetMovies(userid string) (*[]string, error) {
	return getIds(userid, "movie")
}

// GetTvSeries from userid return slice of ids
func GetTvSeries(userid string) (*[]string, error) {
	return getIds(userid, "tvSeries")
}

// GetWatchlistIDs return all imdbid DEPRECIATED
func GetWatchlistIDs(watchlist string) (*[]string, error) {
	userid := imdbUserRe.FindString(watchlist)
	return getIds(userid, "movie,tvSeries")
}
