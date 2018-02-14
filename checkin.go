package untappd

type Checkin struct {
	CheckinID      int     `json:"checkin_id"`
	CreatedAt      string  `json:"created_at"`
	CheckinComment string  `json:"checkin_comment"`
	RatingScore    float64 `json:"rating_score"`
	User           User    `json:"user"`
	Beer           struct {
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
		BreweryID    int     `json:"brewery_id"`
		BreweryName  string  `json:"brewery_name"`
		BrewerySlug  string  `json:"brewery_slug"`
		BreweryLabel string  `json:"brewery_label"`
		CountryName  string  `json:"country_name"`
		Contact      Contact `json:"contact"`
		Location     struct {
			BreweryCity  string  `json:"brewery_city"`
			BreweryState string  `json:"brewery_state"`
			Lat          float64 `json:"lat"`
			Lng          float64 `json:"lng"`
		} `json:"location"`
		BreweryActive int `json:"brewery_active"`
	} `json:"brewery"`
	venue    Venue // unexported. Sometimes a [], sometimes not `json:"venue"`
	Comments struct {
		TotalCount int           `json:"total_count"`
		Count      int           `json:"count"`
		Items      []interface{} `json:"items"`
	} `json:"comments"`
	Toasts struct {
		TotalCount int     `json:"total_count"`
		Count      int     `json:"count"`
		AuthToast  bool    `json:"auth_toast"`
		Items      []Toast `json:"items"`
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

type User struct {
	UID          int       `json:"uid"`
	UserName     string    `json:"user_name"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Location     string    `json:"location"`
	IsSupporter  int       `json:"is_supporter"`
	URL          string    `json:"url"`
	Bio          string    `json:"bio"`
	Relationship string    `json:"relationship"`
	UserAvatar   string    `json:"user_avatar"`
	IsPrivate    int       `json:"is_private"`
	Contact      []Contact `json:"contact"`
}

type Venue struct {
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
	Contact     []Contact `json:"contact"`
	PublicVenue bool      `json:"public_venue"`
	Foursquare  struct {
		FoursquareID  string `json:"foursquare_id"`
		FoursquareURL string `json:"foursquare_url"`
	} `json:"foursquare"`
	VenueIcon struct {
		Sm string `json:"sm"`
		Md string `json:"md"`
		Lg string `json:"lg"`
	} `json:"venue_icon"`
}

type Contact struct {
	Twitter   string `json:"twitter"`
	Facebook  string `json:"facebook"`
	Instagram string `json:"instagram"`
	URL       string `json:"url"`
}
