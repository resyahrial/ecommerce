package gvalidator_test

import (
	"testing"

	"github.com/resyahrial/go-commerce/pkg/gvalidator"
	"github.com/stretchr/testify/suite"
)

type gvalidatorSuite struct {
	suite.Suite
}

func TestGValidator(t *testing.T) {
	suite.Run(t, new(gvalidatorSuite))
}

type person struct {
	Name                string `validate:"required,min=4,max=15"`
	Email               string `validate:"required,email"`
	Age                 int    `validate:"required,numeric,min=18"`
	DriverLicenseNumber string `validate:"omitempty,len=12,numeric"`
}

func (s *gvalidatorSuite) TestValidate() {
	p := person{
		Name:                "Joe",
		Email:               "dummyemail",
		Age:                 0,
		DriverLicenseNumber: "",
	}

	errDesc, ok := gvalidator.Validate(p)
	s.False(ok)
	s.NotEmpty(errDesc)
}
