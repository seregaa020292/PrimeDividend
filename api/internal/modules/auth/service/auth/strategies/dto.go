package strategies

type vkDTO struct {
	Response []struct {
		ID        int    `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	} `json:"response"`
}

type okDTO struct {
}

type yandexDTO struct {
}
