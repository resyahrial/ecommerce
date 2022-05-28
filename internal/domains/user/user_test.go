package user_test

import (
	"testing"

	"github.com/resyahrial/go-commerce/internal/domains/user"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/suite"
)

type userDomainSuite struct {
	suite.Suite
}

func TestUserDomain(t *testing.T) {
	suite.Run(t, new(userDomainSuite))
}

func (s *userDomainSuite) TestUserToBuyer_Success() {
	generalUser := user.User{
		ID:   ksuid.New(),
		Role: user.BUYER,
	}

	buyer, ok := generalUser.ToBuyer()
	s.True(ok)
	s.Equal(generalUser.ID, buyer.ID)
}

func (s *userDomainSuite) TestUserToBuyer_Fail() {
	generalUser := user.User{
		ID:   ksuid.New(),
		Role: user.SELLER,
	}

	buyer, ok := generalUser.ToBuyer()
	s.False(ok)
	s.Equal(ksuid.Nil, buyer.ID)
}

func (s *userDomainSuite) TestUserToSeller_Success() {
	generalUser := user.User{
		ID:   ksuid.New(),
		Role: user.SELLER,
	}

	seller, ok := generalUser.ToSeller()
	s.True(ok)
	s.Equal(generalUser.ID, seller.ID)
}

func (s *userDomainSuite) TestUserToSeller_Fail() {
	generalUser := user.User{
		ID:   ksuid.New(),
		Role: user.BUYER,
	}

	seller, ok := generalUser.ToSeller()
	s.False(ok)
	s.Equal(ksuid.Nil, seller.ID)
}
