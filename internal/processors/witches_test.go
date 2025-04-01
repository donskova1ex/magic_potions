package processors

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type WitchesTestSuite struct {
	suite.Suite
}

func (suite *WitchesTestSuite) SetupSuite() {
}

func (suite *WitchesTestSuite) TeardownSuite() {
}

func (suite *WitchesTestSuite) TestWitchesListSuccess() {

}

func TestWitchesTestSuite(t *testing.T) {
	suite.Run(t, new(WitchesTestSuite))
}
