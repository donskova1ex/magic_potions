package processors

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/donskova1ex/magic_potions/internal/domain"
	"github.com/donskova1ex/magic_potions/internal/processors/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type WitchesTestSuite struct {
	suite.Suite
	ctrl       *gomock.Controller
	processor  *Witches
	repository *mocks.WitchesRepository
	logger     *mocks.WitchesLogger
	ctx        context.Context
}

func (s *WitchesTestSuite) SetupSuite() {
	s.ctrl = gomock.NewController(s.T())
	s.logger = mocks.NewWitchesLogger(s.ctrl)
	s.repository = mocks.NewWitchesRepository(s.ctrl)
	s.processor = NewWitch(s.repository, s.logger)
	s.ctx = context.Background()
}

func (s *WitchesTestSuite) TeardownSuite() {
	s.ctrl.Finish()
}

func (s *WitchesTestSuite) TestWitchesListSuccess() {
	expectedWitches := []*domain.Witch{
		{UUID: "1", Name: "TestWitch"},
	}

	gomock.InOrder(
		s.repository.
			EXPECT().
			WitchesAll(gomock.AssignableToTypeOf(s.ctx)).
			Return(expectedWitches, nil),
	)

	actualWitches, err := s.processor.WitchesList(s.ctx)
	require.NoError(s.T(), err)
	require.Equal(s.T(), expectedWitches, actualWitches)
}

func (s *WitchesTestSuite) TestWitchesListError() {
	dbError := errors.New("db error")
	expectedError := fmt.Errorf("witches list getting error: %w", dbError)

	gomock.InOrder(
		s.repository.
			EXPECT().
			WitchesAll(gomock.AssignableToTypeOf(s.ctx)).
			Return(nil, dbError),
		s.logger.
			EXPECT().
			Error(gomock.Any(), gomock.Any()),
	)

	actualWitches, err := s.processor.WitchesList(s.ctx)
	require.Nil(s.T(), actualWitches)
	require.EqualError(s.T(), err, expectedError.Error())
}

func TestWitchesTestSuite(t *testing.T) {
	suite.Run(t, new(WitchesTestSuite))
}
