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
	hasher_mock "github.com/resyahrial/go-commerce/pkg/hasher/mocks"
	token_manager_mock "github.com/resyahrial/go-commerce/pkg/token-manager/mocks"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/suite"
)

type authenticationUsecaseSuite struct {
	suite.Suite
	authRepo     *auth_dom_mock.MockAuthenticationRepo
	hashHandler  *hasher_mock.MockHasher
	tokenManager *token_manager_mock.MockTokenManager
	userRepo     *user_dom_mock.MockUserRepo
	ucase        authentication.AuthenticationUsecaseInterface
}

func TestAuthenticationUsecase(t *testing.T) {
	suite.Run(t, new(authenticationUsecaseSuite))
}

func (s *authenticationUsecaseSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	s.authRepo = auth_dom_mock.NewMockAuthenticationRepo(ctrl)
	s.hashHandler = hasher_mock.NewMockHasher(ctrl)
	s.tokenManager = token_manager_mock.NewMockTokenManager(ctrl)
	s.userRepo = user_dom_mock.NewMockUserRepo(ctrl)
	s.ucase = authentication.New(s.authRepo, s.hashHandler, s.tokenManager, s.userRepo)
}

func (s *authenticationUsecaseSuite) TestLogin_Success() {
	loginInput := auth_dom.Login{
		Email:    "email@email.com",
		Password: "qwerty",
	}

	s.userRepo.EXPECT().GetDetail(gomock.Any(), user_dom.User{
		Email: loginInput.Email,
	}).Return(user_dom.User{
		ID:       ksuid.New(),
		Email:    loginInput.Email,
		Password: "hashedPassword",
	}, nil)

	s.hashHandler.EXPECT().Compare(loginInput.Password, "hashedPassword").Return(true)

	_, err := s.ucase.Login(context.Background(), loginInput)
	s.Nil(err)
}
