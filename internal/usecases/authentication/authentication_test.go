package authentication_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	auth_dom "github.com/resyahrial/go-commerce/internal/domains/authentication"
	auth_dom_mock "github.com/resyahrial/go-commerce/internal/domains/authentication/mocks"
	user_dom "github.com/resyahrial/go-commerce/internal/domains/user"
	user_dom_mock "github.com/resyahrial/go-commerce/internal/domains/user/mocks"
	"github.com/resyahrial/go-commerce/internal/usecases/authentication"
	"github.com/resyahrial/go-commerce/pkg/hasher"
	"github.com/stretchr/testify/suite"
)

type authenticationUsecaseSuite struct {
	suite.Suite
	userRepo *user_dom_mock.MockUserRepo
	authRepo *auth_dom_mock.MockAuthenticationRepo
	ucase    authentication.AuthenticationUsecaseInterface
}

func TestAuthenticationUsecase(t *testing.T) {
	suite.Run(t, new(authenticationUsecaseSuite))
}

func (s *authenticationUsecaseSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	s.userRepo = user_dom_mock.NewMockUserRepo(ctrl)
	s.authRepo = auth_dom_mock.NewMockAuthenticationRepo(ctrl)
	s.ucase = authentication.New(s.userRepo, s.authRepo)
}

func (s *authenticationUsecaseSuite) TestLogin_Success() {
	loginInput := auth_dom.Login{
		Email:    "email@email.com",
		Password: "qwerty",
	}

	hashedPassword, _ := hasher.HashPassword(loginInput.Password)

	s.userRepo.EXPECT().GetDetail(gomock.Any(), user_dom.User{
		Email: loginInput.Email,
	}).Return(user_dom.User{
		Email:    loginInput.Email,
		Password: hashedPassword,
	}, nil)

	_, err := s.ucase.Login(context.Background(), loginInput)
	s.Nil(err)
}
