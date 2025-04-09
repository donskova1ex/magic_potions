package processors

import (
	"context"
	"errors"
	"fmt"
	"github.com/donskova1ex/magic_potions/internal/domain"
	"github.com/donskova1ex/magic_potions/internal/processors/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type IngredientsTestSuite struct {
	suite.Suite
	ctrl       *gomock.Controller
	processor  *Ingredients
	repository *mocks.IngredientsRepository
	logger     *mocks.IngredientsLogger
	ctx        context.Context
}

func (s *IngredientsTestSuite) SetupSuite() {
	s.ctrl = gomock.NewController(s.T())
	s.logger = mocks.NewIngredientsLogger(s.ctrl)
	s.repository = mocks.NewIngredientsRepository(s.ctrl)
	s.processor = NewIngredient(s.repository, s.logger)
	s.ctx = context.Background()
}

func (s *IngredientsTestSuite) TearDownSuite() {
	s.ctrl.Finish()
}

func (s *IngredientsTestSuite) TestIngredientsListSuccess() {
	expectedIngredients := []*domain.Ingredient{
		{
			ID:       1,
			UUID:     "1",
			Name:     "FirstIngredient",
			Quantity: 6,
		},
	}

	gomock.InOrder(
		s.repository.
			EXPECT().
			IngredientsAll(s.ctx).
			Return(expectedIngredients, nil),
	)
	actualIngredients, err := s.processor.IngredientsList(s.ctx)
	require.NoError(s.T(), err)
	require.Equal(s.T(), expectedIngredients, actualIngredients)
}

func (s *IngredientsTestSuite) TestIngredientsListError() {
	dbError := errors.New("db error")

	expectedError := fmt.Errorf("it is impossible to get a ingredients list: %w", dbError)

	gomock.InOrder(
		s.repository.
			EXPECT().
			IngredientsAll(gomock.Any()).
			Return(nil, dbError),
		s.logger.
			EXPECT().
			Error(gomock.Any(), gomock.Any()),
	)
	actualIngredients, err := s.processor.IngredientsList(s.ctx)
	require.Nil(s.T(), actualIngredients)
	require.EqualError(s.T(), err, expectedError.Error())
}

func (s *IngredientsTestSuite) TestCreateIngredientSuccess() {
	expectedIngredient := &domain.Ingredient{
		ID:       1,
		UUID:     "1",
		Name:     "FirstIngredient",
		Quantity: 16,
	}

	gomock.InOrder(
		s.repository.
			EXPECT().
			CreateIngredient(gomock.Any(), gomock.Any()).
			Return(expectedIngredient, nil),
	)

	actualIngredient, err := s.processor.CreateIngredient(s.ctx, expectedIngredient)
	require.NoError(s.T(), err)
	require.Equal(s.T(), expectedIngredient, actualIngredient)
}

func (s *IngredientsTestSuite) TestCreateIngredientError() {

}

func TestIngredientsTestSuite(t *testing.T) {
	suite.Run(t, new(IngredientsTestSuite))
}
