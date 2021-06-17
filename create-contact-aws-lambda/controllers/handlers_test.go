package controllers

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/wilomendez/Onboarding/create-contact-aws-lambda/mocks"
	"github.com/wilomendez/Onboarding/create-contact-aws-lambda/models"
)

func TestCreate(t *testing.T) {
	mockController := gomock.NewController(t)
	mockHandler := mocks.NewMockICreateContactHandler(mockController)
	defer mockController.Finish()

	// Fixtures
	var (
		request = models.Request{
			FirstName: "John",
			LastName:  "Doe",
		}
		ctx      = context.Background()
		response = models.Response{
			ContactNumber: "76334e6d-3707-4c82-aa68-c6897140b5ef",
		}
		emptyResponse = models.Response{}
		badRequest    = models.Request{
			LastName: "Doe",
		}
		errorMessage = HandleCustomValidationError("Input validation FAILED, First Name or Last name is empty")
	)

	t.Run("When you sent correct data", func(t *testing.T) {
		gomock.InOrder(
			mockHandler.EXPECT().Create(ctx, request).Return(response, nil).Times(1),
		)
		resp, err := mockHandler.Create(ctx, request)
		if assert.NoError(t, err) {
			assert.Equal(t, response, resp)
		}
	})

	t.Run("When you sent empty data", func(t *testing.T) {
		gomock.InOrder(
			mockHandler.EXPECT().Create(ctx, gomock.Any()).Return(emptyResponse, errorMessage),
		)

		resp, err := mockHandler.Create(ctx, badRequest)
		if assert.Error(t, err) {
			assert.Empty(t, resp)
			assert.Equal(t, errorMessage.Error(), err.Error())
		}
	})
}
