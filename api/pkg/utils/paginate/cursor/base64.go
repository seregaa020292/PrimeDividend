package cursor

import (
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

func Encode(cursor Cursor) (string, error) {
	type Payload struct {
		ID        uuid.UUID `json:"id"`
		Time      time.Time `json:"time"`
		Direction Direction `json:"direction"`
	}

	payload := Payload(cursor)

	bytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(bytes), nil
}

func Decode(encodedCursor string) (Cursor, error) {
	byt, err := base64.StdEncoding.DecodeString(encodedCursor)
	if err != nil {
		return Cursor{}, err
	}

	var cursor Cursor
	if err := json.Unmarshal(byt, &cursor); err != nil {
		return Cursor{}, err
	}

	return cursor, nil
}
