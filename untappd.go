package untappd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type Untappd struct {
	Token string
	Limit int
}

func New(token string) *Untappd {
	return &Untappd{token, 25}
}

func (u *Untappd) query(url *url.URL) ([]byte, error) {
	var empty []byte
	resp, err := http.Get(url.String())
	if err != nil {
		return empty, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return empty, err
	}

	if resp.StatusCode == 500 {
		log.Printf("Error querying untappd: %s, %s", resp.Status, body)
		return empty, errors.New(resp.Status)
	}
	return body, nil
}

func (u *Untappd) PullUserCheckins(userName string) ([]Checkin, error) {
	var empty []Checkin
	url, _ := url.Parse("https://api.untappd.com/v4/user/checkins/")
	url, err := url.Parse(userName)
	if err != nil {
		return empty, err
	}
	q := url.Query()
	q.Set("access_token", u.Token)
	url.RawQuery = q.Encode()

	body, err := u.query(url)
	if err != nil {
		return []Checkin{}, err
	}

	var beers feedResp
	err = json.Unmarshal(body, &beers)
	if err != nil {
		err = fmt.Errorf("Error:\n%s\n\nPayload:\n %s", err, body)
		return []Checkin{}, err
	}
	return beers.Response.Checkins.Items, nil
}

func (u *Untappd) PullFeed() ([]Checkin, error) {
	url, _ := url.Parse("https://api.untappd.com/v4/checkin/recent/")
	q := url.Query()
	q.Set("access_token", u.Token)
	q.Set("limit", strconv.Itoa(u.Limit))
	url.RawQuery = q.Encode()

	body, err := u.query(url)
	if err != nil {
		return []Checkin{}, err
	}

	var beers feedResp
	err = json.Unmarshal(body, &beers)
	if err != nil {
		err = fmt.Errorf("Error:\n%s\n\nPayload:\n %s", err, body)
		return []Checkin{}, err
	}
	return beers.Response.Checkins.Items, nil
}

func (u *Untappd) Toast(checkinID int) ([]Toast, error) {
	var empty []Toast
	url, err := url.Parse("https://api.untappd.com/v4/checkin/toast/")
	if err != nil {
		return empty, err
	}
	url, err = url.Parse(strconv.Itoa(checkinID))
	if err != nil {
		return empty, err
	}
	q := url.Query()
	q.Set("access_token", u.Token)
	url.RawQuery = q.Encode()

	body, err := u.query(url)
	if err != nil {
		return empty, err
	}

	var toasts toastResponse
	err = json.Unmarshal(body, &toasts)
	if err != nil {
		log.Println(err)
		return empty, err
	}
	return toasts.Response.Toasts.Items, nil
}
