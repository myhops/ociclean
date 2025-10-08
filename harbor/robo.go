package harbor

import (
	"encoding/json"
	"os"
	"time"

	"github.com/google/go-containerregistry/pkg/authn"
)

type RoboAccount struct {
	CreationTime time.Time `json:"creation_time"`
	Name         string    `json:"name"`
	Secret       string    `json:"secret"`
}

func ReadSecret(filename string) (*RoboAccount, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	res := new(RoboAccount)
	if err := json.Unmarshal(b, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func BasicAuth(rb *RoboAccount) *authn.Basic {
	return &authn.Basic{
		Username: rb.Name,
		Password: rb.Secret,
	}
}
