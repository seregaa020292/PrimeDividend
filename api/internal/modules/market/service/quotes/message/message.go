package message

type Message struct {
	Provider string  `json:"provider"`
	Identity string  `json:"identity"`
	Price    float64 `json:"price"`
}
