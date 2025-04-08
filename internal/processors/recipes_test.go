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

type RecipesTestSuite struct {
	suite.Suite
	ctrl       *gomock.Controller
	processor  *Recipes
	repository *mocks.RecipesRepository
	logger     *mocks.RecipesLogger
	ctx        context.Context
}

func (s *RecipesTestSuite) SetupSuite() {
	s.ctrl = gomock.NewController(s.T())
	s.logger = mocks.NewRecipesLogger(s.ctrl)
	s.repository = mocks.NewRecipesRepository(s.ctrl)
	s.processor = NewRecipe(s.repository, s.logger)
	s.ctx = context.Background()
}

func (s *RecipesTestSuite) TearDownSuite() {
	s.ctrl.Finish()
}

func (s *RecipesTestSuite) TestRecipesListSuccess() {
	expectedRecipes := []*domain.Recipe{
		{
			ID:              1,
			UUID:            "1",
			Name:            "FirstRecipe",
			BrewTimeSeconds: 6,
			Ingredients: []*domain.Ingredient{
				{
					ID:       1,
					UUID:     "12",
					Name:     "FirstIngredient",
					Quantity: 3,
				},
			},
		},
	}

	gomock.InOrder(
		s.repository.
			EXPECT().
			RecipesAll(gomock.Any()).
			Return(expectedRecipes, nil),
	)

	actualRecipes, err := s.processor.RecipesList(s.ctx)
	require.NoError(s.T(), err)
	require.Equal(s.T(), expectedRecipes, actualRecipes)
}

func (s *RecipesTestSuite) TestRecipesListError() {
	dbError := errors.New("db error")
	expectedError := fmt.Errorf("recipes list getting error: %w", dbError)
	gomock.InOrder(
		s.repository.
			EXPECT().
			RecipesAll(gomock.Any()).
			Return(nil, dbError),
		s.logger.
			EXPECT().
			Error(gomock.Any(), gomock.Any()),
	)
	actualRecipes, err := s.processor.RecipesList(s.ctx)
	require.Nil(s.T(), actualRecipes)
	require.EqualError(s.T(), err, expectedError.Error())
}

func (s *RecipesTestSuite) TestCreateRecipeSuccess() {
	expectedRecipe := &domain.Recipe{
		ID:              1,
		UUID:            "1",
		Name:            "FirstRecipe",
		BrewTimeSeconds: 10,
		Ingredients:     nil,
	}

	gomock.InOrder(
		s.repository.
			EXPECT().
			CreateRecipe(gomock.Any(), gomock.Any()).
			Return(expectedRecipe, nil),
	)

	recipe := &domain.Recipe{
		ID:              2,
		UUID:            "2",
		Name:            "SecondRecipe",
		BrewTimeSeconds: 20,
		Ingredients:     nil,
	}

	actualRecipes, err := s.processor.CreateRecipe(s.ctx, recipe)
	require.NoError(s.T(), err)
	require.Equal(s.T(), expectedRecipe, actualRecipes)
}

func (s *RecipesTestSuite) TestCreateRecipeError() {
	dbError := errors.New("db error")
	expectedRecipe := &domain.Recipe{
		ID:              1,
		UUID:            "1",
		Name:            "FirstRecipe",
		BrewTimeSeconds: 10,
		Ingredients:     nil,
	}

	expectedError := fmt.Errorf("can not create recipe: %s, error: %w", expectedRecipe.Name, dbError)

	gomock.InOrder(
		s.repository.
			EXPECT().
			CreateRecipe(gomock.Any(), gomock.Any()).
			Return(nil, dbError),
		s.logger.
			EXPECT().
			Error(gomock.Any(), gomock.Any()),
	)

	actualRecipe, err := s.processor.CreateRecipe(s.ctx, expectedRecipe)
	require.Nil(s.T(), actualRecipe)
	require.EqualError(s.T(), err, expectedError.Error())
}

func (s *RecipesTestSuite) TestRecipeByUUIDSuccess() {
	expectedRecipe := &domain.Recipe{
		ID:              1,
		UUID:            "1",
		Name:            "FirstRecipe",
		BrewTimeSeconds: 10,
		Ingredients:     nil,
	}

	gomock.InOrder(
		s.repository.
			EXPECT().
			RecipeByUUID(gomock.Any(), gomock.Any()).
			Return(expectedRecipe, nil),
	)
	actualRecipe, err := s.processor.RecipeByUUID(s.ctx, expectedRecipe.UUID)
	require.NoError(s.T(), err)
	require.Equal(s.T(), expectedRecipe, actualRecipe)
}

func (s *RecipesTestSuite) TestRecipeByUUIDError() {
	expectedRecipe := &domain.Recipe{
		ID:              1,
		UUID:            "1",
		Name:            "FirstRecipe",
		BrewTimeSeconds: 10,
		Ingredients:     nil,
	}
	dbError := errors.New("db error")
	expectedError := fmt.Errorf("can not get recipe by uuid: %s, error: %w",
		expectedRecipe.UUID,
		dbError,
	)

	gomock.InOrder(
		s.repository.
			EXPECT().
			RecipeByUUID(gomock.Any(), gomock.Any()).
			Return(nil, dbError),
		s.logger.
			EXPECT().
			Error(gomock.Any(), gomock.Any()),
	)
	actualRecipe, err := s.processor.RecipeByUUID(s.ctx, expectedRecipe.UUID)
	require.Nil(s.T(), actualRecipe)
	require.EqualError(s.T(), err, expectedError.Error())
}

func (s *RecipesTestSuite) TestDeleteRecipeByUUIDSuccess() {
	expectedRecipe := &domain.Recipe{
		ID:              1,
		UUID:            "1",
		Name:            "FirstRecipe",
		BrewTimeSeconds: 10,
		Ingredients:     nil,
	}
	gomock.InOrder(
		s.repository.
			EXPECT().
			DeleteRecipeByUUID(gomock.Any(), gomock.Any()).
			Return(nil),
	)

	err := s.processor.DeleteRecipeByUUID(s.ctx, expectedRecipe.UUID)
	require.NoError(s.T(), err)
}

func (s *RecipesTestSuite) TestDeleteRecipeByUUIDError() {
	dbError := errors.New("db error")
	expectedRecipe := &domain.Recipe{
		ID:              1,
		UUID:            "1",
		Name:            "FirstRecipe",
		BrewTimeSeconds: 10,
		Ingredients:     nil,
	}
	expectedError := fmt.Errorf("unable to delete recipe by uuid: %s, error: %w", expectedRecipe.UUID, dbError)

	gomock.InOrder(
		s.repository.
			EXPECT().
			DeleteRecipeByUUID(gomock.Any(), gomock.Any()).
			Return(dbError),
		s.logger.
			EXPECT().
			Error(gomock.Any(), gomock.Any()),
	)

	err := s.processor.DeleteRecipeByUUID(s.ctx, expectedRecipe.UUID)
	require.EqualError(s.T(), err, expectedError.Error())

}

func (s *RecipesTestSuite) TestUpdateRecipeByUUIDSuccess() {
	expectedRecipe := &domain.Recipe{
		ID:              1,
		UUID:            "1",
		Name:            "FirstRecipe",
		BrewTimeSeconds: 10,
		Ingredients:     nil,
	}

	gomock.InOrder(
		s.repository.
			EXPECT().
			UpdateRecipeByUUID(gomock.Any(), gomock.Any()).
			Return(expectedRecipe, nil),
	)
	recipe := &domain.Recipe{
		ID:              2,
		UUID:            "2",
		Name:            "SecondRecipe",
		BrewTimeSeconds: 20,
		Ingredients:     nil,
	}
	actualRecipe, err := s.processor.UpdateRecipeByID(s.ctx, recipe)
	require.NoError(s.T(), err)
	require.Equal(s.T(), expectedRecipe, actualRecipe)
}

func (s *RecipesTestSuite) TestUpdateRecipeByUUIDError() {
	dbError := errors.New("db error")
	expectedRecipe := &domain.Recipe{
		ID:              1,
		UUID:            "1",
		Name:            "FirstRecipe",
		BrewTimeSeconds: 10,
		Ingredients:     nil,
	}
	expectedError := fmt.Errorf("can not update recipe: %s, error: %w",
		expectedRecipe.Name,
		dbError,
	)

	gomock.InOrder(
		s.repository.
			EXPECT().
			UpdateRecipeByUUID(gomock.Any(), gomock.Any()).
			Return(nil, dbError),
		s.logger.
			EXPECT().
			Error(gomock.Any(), gomock.Any()),
	)
	actualRecipe, err := s.processor.UpdateRecipeByID(s.ctx, expectedRecipe)
	require.Nil(s.T(), actualRecipe)
	require.EqualError(s.T(), err, expectedError.Error())
}

func TestRecipesTestSuite(t *testing.T) {
	suite.Run(t, new(RecipesTestSuite))
}
