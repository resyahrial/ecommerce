package seller

import "github.com/segmentio/ksuid"

type Seller struct {
	ID      ksuid.KSUID `json:"ID"`
	Email   string      `json:"email"`
	Name    string      `json:"name"`
	Address string      `json:"address"`
}
