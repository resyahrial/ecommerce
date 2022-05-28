package authentication_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	auth_dom "github.com/resyahrial/go-commerce/internal/domains/authentication"
	auth_dom_mock "github.com/resyahrial/go-commerce/internal/domains/authentication/mocks"
	"github.com/resyahrial/go-commerce/internal/usecases/authentication"
	"github.com/stretchr/testify/suite"
)

type authenticationUsecaseSuite struct {
	suite.Suite
	authRepo *auth_dom_mock.MockAuthenticationRepo
	ucase    authentication.AuthenticationUsecaseInterface
}

func TestAuthenticationUsecase(t *testing.T) {
	suite.Run(t, new(authenticationUsecaseSuite))
}

func (s *authenticationUsecaseSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	s.authRepo = auth_dom_mock.NewMockAuthenticationRepo(ctrl)
	s.ucase = authentication.New(s.authRepo)
}

func (s *authenticationUsecaseSuite) Login_Success() {
	_, err := s.ucase.Login(context.Background(), auth_dom.Login{})
	s.Nil(err)
}
