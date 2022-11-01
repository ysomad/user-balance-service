package pagetoken

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

var ErrInvalid = errors.New("invalid token, must store uuid and RFC3339 datetime separated with comma")

// Encode encodes uuid and t into base64 string separated with ",".
func Encode(pk uuid.UUID, t time.Time) string {
	s := fmt.Sprintf("%s,%s", pk, t.Format(time.RFC3339Nano))
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// Decode decodes page token into uuid and created at time,
// token must be a string splitted with "," and encoded into base64,
// not decoded token example: `ef6e33ec-5b1d-4ade-8dcc-b508262ee859,2006-01-02T15:04:05.999999999Z07:00`
func Decode(token string) (uuid.UUID, time.Time, error) {
	b, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return uuid.UUID{}, time.Time{}, err
	}

	parts := strings.Split(string(b), ",")
	if len(parts) != 2 {
		return uuid.UUID{}, time.Time{}, ErrInvalid
	}

	t, err := time.Parse(time.RFC3339Nano, parts[1])
	if err != nil {
		return uuid.UUID{}, time.Time{}, err
	}

	id, err := uuid.Parse(parts[0])
	if err != nil {
		return uuid.UUID{}, time.Time{}, err
	}

	return id, t, nil
}
