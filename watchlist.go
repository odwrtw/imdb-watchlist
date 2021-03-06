package imdbwatchlist

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"time"
)

var (
	imdbIDRegexp = regexp.MustCompile(`tt\d{7}`)
	imdbUserRe   = regexp.MustCompile(`ur\d{8}`) // DEPRECIATED
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

func getIdsPage(userid string, filter string, sort string, page int) (*[]string, error) {
	URL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	URL.Path += userid
	URL.Path += "/watchlist"

	parameters := url.Values{}
	parameters.Add("title_type", filter)
	parameters.Add("sort", sort)
	parameters.Add("page", strconv.Itoa(page))
	URL.RawQuery = parameters.Encode()

	var httpClient = &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := httpClient.Get(URL.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("got http error %q", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	matches := imdbIDRegexp.FindAllString(string(body), -1)
	if matches != nil {
		unique(&matches)
		return &matches, nil
	}
	return nil, errors.New("empty page")
}

func getIds(userid string, filter string, sort string) (*[]string, error) {
	var matches []string
	page := 1

	var previousListCount = 0
	for {
		curMatches, err := getIdsPage(userid, filter, sort, page)
		if err != nil {
			return nil, err
		}

		matches = append(matches, *curMatches...)
		unique(&matches)
		currentListCount := len(matches)
		if currentListCount == previousListCount {
			// Nothing new, let's break the loop
			break
		}
		previousListCount = currentListCount
		page++
	}

	return &matches, nil
}

// GetMovies from userid return slice of ids
func GetMovies(userid string) (*[]string, error) {
	return getIds(userid, "movie", "date_added,desc")
}

// GetTvSeries from userid return slice of ids
func GetTvSeries(userid string) (*[]string, error) {
	return getIds(userid, "tvSeries", "date_added,desc")
}

// GetWatchlistIDs return all imdbid DEPRECIATED
func GetWatchlistIDs(watchlist string) (*[]string, error) {
	userid := imdbUserRe.FindString(watchlist)
	return getIds(userid, "movie,tvSeries", "date_added,desc")
}
