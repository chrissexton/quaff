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

type Checkin struct {
	CheckinID      int     `json:"checkin_id"`
	CreatedAt      string  `json:"created_at"`
	CheckinComment string  `json:"checkin_comment"`
	RatingScore    float64 `json:"rating_score"`
	User           struct {
		UID          int    `json:"uid"`
		UserName     string `json:"user_name"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		Location     string `json:"location"`
		IsSupporter  int    `json:"is_supporter"`
		URL          string `json:"url"`
		Bio          string `json:"bio"`
		Relationship string `json:"relationship"`
		UserAvatar   string `json:"user_avatar"`
		IsPrivate    int    `json:"is_private"`
		Contact      struct {
			Foursquare int    `json:"foursquare"`
			Twitter    string `json:"twitter"`
			Facebook   int    `json:"facebook"`
		} `json:"contact"`
	} `json:"user"`
	Beer struct {
		Bid        int     `json:"bid"`
		BeerName   string  `json:"beer_name"`
		BeerLabel  string  `json:"beer_label"`
		BeerStyle  string  `json:"beer_style"`
		BeerAbv    float64 `json:"beer_abv"`
		AuthRating float64 `json:"auth_rating"`
		WishList   bool    `json:"wish_list"`
		BeerActive int     `json:"beer_active"`
	} `json:"beer"`
	Brewery struct {
		BreweryID    int    `json:"brewery_id"`
		BreweryName  string `json:"brewery_name"`
		BrewerySlug  string `json:"brewery_slug"`
		BreweryLabel string `json:"brewery_label"`
		CountryName  string `json:"country_name"`
		Contact      struct {
			Twitter   string `json:"twitter"`
			Facebook  string `json:"facebook"`
			Instagram string `json:"instagram"`
			URL       string `json:"url"`
		} `json:"contact"`
		Location struct {
			BreweryCity  string  `json:"brewery_city"`
			BreweryState string  `json:"brewery_state"`
			Lat          float64 `json:"lat"`
			Lng          float64 `json:"lng"`
		} `json:"location"`
		BreweryActive int `json:"brewery_active"`
	} `json:"brewery"`
	venue struct {
		VenueID          int    `json:"venue_id"`
		VenueName        string `json:"venue_name"`
		PrimaryCategory  string `json:"primary_category"`
		ParentCategoryID string `json:"parent_category_id"`
		Categories       struct {
			Count int `json:"count"`
			Items []struct {
				CategoryName string `json:"category_name"`
				CategoryID   string `json:"category_id"`
				IsPrimary    bool   `json:"is_primary"`
			} `json:"items"`
		} `json:"categories"`
		Location struct {
			VenueAddress string  `json:"venue_address"`
			VenueCity    string  `json:"venue_city"`
			VenueState   string  `json:"venue_state"`
			VenueCountry string  `json:"venue_country"`
			Lat          float64 `json:"lat"`
			Lng          float64 `json:"lng"`
		} `json:"location"`
		Contact struct {
			Twitter  string `json:"twitter"`
			VenueURL string `json:"venue_url"`
		} `json:"contact"`
		PublicVenue bool `json:"public_venue"`
		Foursquare  struct {
			FoursquareID  string `json:"foursquare_id"`
			FoursquareURL string `json:"foursquare_url"`
		} `json:"foursquare"`
		VenueIcon struct {
			Sm string `json:"sm"`
			Md string `json:"md"`
			Lg string `json:"lg"`
		} `json:"venue_icon"`
	} // unexported. Sometimes a [], sometimes not `json:"venue"`
	Comments struct {
		TotalCount int           `json:"total_count"`
		Count      int           `json:"count"`
		Items      []interface{} `json:"items"`
	} `json:"comments"`
	Toasts struct {
		TotalCount int           `json:"total_count"`
		Count      int           `json:"count"`
		AuthToast  bool          `json:"auth_toast"`
		Items      []interface{} `json:"items"`
	} `json:"toasts"`
	Media struct {
		Count int           `json:"count"`
		Items []interface{} `json:"items"`
	} `json:"media"`
	Source struct {
		AppName    string `json:"app_name"`
		AppWebsite string `json:"app_website"`
	} `json:"source"`
	Badges struct {
		Count int `json:"count"`
		Items []struct {
			BadgeID          int    `json:"badge_id"`
			UserBadgeID      int    `json:"user_badge_id"`
			BadgeName        string `json:"badge_name"`
			BadgeDescription string `json:"badge_description"`
			CreatedAt        string `json:"created_at"`
			BadgeImage       struct {
				Sm string `json:"sm"`
				Md string `json:"md"`
				Lg string `json:"lg"`
			} `json:"badge_image"`
		} `json:"items"`
	} `json:"badges"`
}

type Meta struct {
	Code         int `json:"code"`
	ResponseTime struct {
		Time    float64 `json:"time"`
		Measure string  `json:"measure"`
	} `json:"response_time"`
	InitTime struct {
		Time    int    `json:"time"`
		Measure string `json:"measure"`
	} `json:"init_time"`
}

type feedResp struct {
	Meta     Meta `json:"meta"`
	Response struct {
		Checkins struct {
			Count int       `json:"count"`
			Items []Checkin `json:"items"`
		} `json:"checkins"`
	} `json:"response"`
}

type Toasts struct {
	AuthToken bool `json:"auth_token"`
	Count     int  `json:"count"`
	Items     struct {
		UID  int `json:"uid"`
		User struct {
			UID            int           `json:"uid"`
			UserName       string        `json:"user_name"`
			FirstName      string        `json:"first_name"`
			LastName       string        `json:"last_name"`
			Bio            string        `json:"bio"`
			Location       string        `json:"location"`
			UserAvatar     string        `json:"user_avatar"`
			UserLink       string        `json:"user_link"`
			AccountType    string        `json:"account_type"`
			BreweryDetails []interface{} `json:"brewery_details"`
		} `json:"user"`
		LikeID    int    `json:"like_id"`
		LikeOwner bool   `json:"like_owner"`
		CreatedAt string `json:"created_at"`
	} `json:"items"`
}

type toastResponse struct {
	Meta     Meta `json:"meta"`
	Response struct {
		Result   string `json:"result"`
		LikeType string `json:"like_type"`
		Toasts   Toasts `json:"toasts"`
	} `json:"response"`
}

func (u *Untappd) PullFeed() ([]Checkin, error) {
	url, _ := url.Parse("https://api.untappd.com/v4/checkin/recent/")
	q := url.Query()
	q.Set("access_token", u.Token)
	q.Set("limit", strconv.Itoa(u.Limit))
	url.RawQuery = q.Encode()

	resp, err := http.Get(url.String())
	if err != nil {
		return []Checkin{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Checkin{}, err
	}

	if resp.StatusCode == 500 {
		log.Printf("Error querying untappd: %s, %s", resp.Status, body)
		return []Checkin{}, errors.New(resp.Status)
	}

	var beers feedResp
	err = json.Unmarshal(body, &beers)
	if err != nil {
		err = fmt.Errorf("Error:\n%s\n\nPayload:\n %s", err, body)
		return []Checkin{}, err
	}
	return beers.Response.Checkins.Items, nil
}

func (u *Untappd) Toast(checkinID int) (Toasts, error) {
	url, err := url.Parse("https://api.untappd.com/v4/checkin/toast/")
	if err != nil {
		return Toasts{}, err
	}
	url, err = url.Parse(strconv.Itoa(checkinID))
	if err != nil {
		return Toasts{}, err
	}
	q := url.Query()
	q.Set("access_token", u.Token)
	url.RawQuery = q.Encode()

	log.Fatal(url.String())

	resp, err := http.Get(url.String())
	if err != nil {
		return Toasts{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Toasts{}, err
	}

	if resp.StatusCode == 500 {
		log.Printf("Error querying untappd: %s, %s", resp.Status, body)
		return Toasts{}, errors.New(resp.Status)
	}

	var toasts toastResponse
	err = json.Unmarshal(body, &toasts)
	if err != nil {
		log.Println(err)
		return Toasts{}, err
	}
	return toasts.Response.Toasts, nil
}
