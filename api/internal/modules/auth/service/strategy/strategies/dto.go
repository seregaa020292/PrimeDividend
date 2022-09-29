package strategies

type responseVK struct {
	Response []struct {
		ID        int    `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	} `json:"response"`
}

type responseOK struct {
}

type responseYandex struct {
}
