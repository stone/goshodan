package goshodan

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
)

const (
	SHODANBASEURL = "http://www.shodanhq.com/api"
)

// This is the result after decoding the json sent by shodan, it is 1:1 mapping
// of the json with the exeption using capitalization and CamelCase where keys
// got an underscore
type Result struct {
	Matches   []Match   `json:"matches"`
	Total     int       `json:"total"`
	Cities    []City    `json:"cities"`
	Countries []Country `json:"country"`
}

// Match represents a shodan Match and contains data about the host
type Match struct {
	Updated      string   `json:"updated"`
	RegionName   string   `json:"region_name"`
	IP           string   `json:"ip"`
	AreaCode     int      `json:"area_code"`
	CountryName  string   `json:"country_name"`
	Hostnames    []string `json:"hostnames,omitempty"`
	PostalCode   string   `json:"postal_code"`
	DmaCode      int      `json:"dma_code"`
	Org          string   `json:"org"`
	Data         string   `json:"data"`
	Port         int      `json:"port"`
	City         string   `json:"city"`
	Title        string   `json:"title"`
	Isp          string   `json:"isp"`
	Longitude    float32  `json:"longitude"`
	CountryCode3 string   `json:"country_code3"`
	Html         string   `json:"html"`
	Latitude     float32  `json:"latitude"`
	Os           string   `json:"os"`
}

// City represents statistics about citites in the query
type City struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
}

// Country represents statistics about contries in the query
type Country struct {
	Count int    `json:"country"`
	Code  string `json:"code"`
	Name  string `json:"name"`
}

// Represents a Shodanhq query type
type Shodan struct {
	apikey string
	debug  bool
	page   int
	query  string
}

// Search shodanhq.com using "query" and "page" query uses the
// same syntax as the serarch engine on the homepage:
// http://www.shodanhq.com/help/filters and page is
// which page of results to return.
func (s *Shodan) Search(query string, page int) (*Result, error) {

	var shurl *url.URL
	shurl, err := url.Parse(SHODANBASEURL)
	if err != nil {
		panic("BOOM, this should _not_ happen")
	}
	shurl.Path += "/search"
	param := url.Values{}
	param.Add("key", s.apikey)
	param.Add("q", query)
	param.Add("p", strconv.Itoa(page))
	shurl.RawQuery = param.Encode()
	//fmt.Printf("Encoded URL is %q\n", shurl.String())

	client := &http.Client{}
	resp, err := client.Get(shurl.String())
	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(resp.Body)
	defer resp.Body.Close()

	res := &Result{}
	if err := dec.Decode(res); err != nil {
		return nil, err
	}
	s.page = page
	return res, nil
}

// Return next page of results on the previous query
func (s *Shodan) NextPage() (*Result, error) {
	if s.query == "" {
		return nil, errors.New("Error getting next page, no previous query")
	}
	return s.Search(s.query, s.page+1)
}

//Returns the total number of search results for the query.
// returns error and -1 on error
func (s *Shodan) Count(query string) (num int, err error) {
	res, err := s.Search(query, 1)
	if err != nil {
		return -1, err
	}
	return res.Total, nil
}

// NewWebAPI creates a new Shodan which can be used for querying.
func NewWebAPI(apikey string) (shodan *Shodan) {
	return &Shodan{apikey: apikey}
}

// Not used yet
func (s *Shodan) Debug(on bool) {
	s.debug = on
}
