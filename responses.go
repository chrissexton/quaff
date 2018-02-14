package untappd

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

type toastResponse struct {
	Meta     Meta `json:"meta"`
	Response struct {
		Result   string `json:"result"`
		LikeType string `json:"like_type"`
		Toasts   struct {
			AuthToken  bool    `json:"auth_token"`
			Count      int     `json:"count"`
			Items      []Toast `json:"items"`
			manyToasts []Toast
			oneToast   Toast
		} `json:"toasts"`
	} `json:"response"`
}
