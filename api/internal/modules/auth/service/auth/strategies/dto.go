package strategies

type vkBody struct {
	Response []struct {
		ID        int    `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	} `json:"response"`
}

type okBody struct {
}

type yandexBody struct {
}
