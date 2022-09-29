package strategies

type responseVK struct {
	Response []struct {
		ID        int    `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	} `json:"response"`
}

type responseOK struct {
	Age         int64  `json:"age"`
	Birthday    string `json:"birthday"`
	BirthdaySet bool   `json:"birthdaySet"`
	FirstName   string `json:"first_name"`
	Gender      string `json:"gender"`
	Email       string `json:"email"`
	HasEmail    bool   `json:"has_email"`
	LastName    string `json:"last_name"`
	Locale      string `json:"locale"`
	Location    struct {
		City        string `json:"city"`
		Country     string `json:"country"`
		CountryCode string `json:"countryCode"`
		CountryName string `json:"countryName"`
	} `json:"location"`
	Name    string `json:"name"`
	Online  string `json:"online"`
	PhotoID string `json:"photo_id"`
	Pic1    string `json:"pic_1"`
	Pic2    string `json:"pic_2"`
	Pic3    string `json:"pic_3"`
	UID     string `json:"uid"`
}

type responseYandex struct {
}
