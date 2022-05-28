package buyer

import "github.com/segmentio/ksuid"

type Buyer struct {
	ID      ksuid.KSUID `json:"ID"`
	Email   string      `json:"email"`
	Name    string      `json:"name"`
	Address string      `json:"address"`
}
