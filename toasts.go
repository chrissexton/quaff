package untappd

type Toast struct {
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
}
