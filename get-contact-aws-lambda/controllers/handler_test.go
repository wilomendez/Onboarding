package controllers

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/wilomendez/Onboarding/get-contact-aws-lambda/mocks"
	"github.com/wilomendez/Onboarding/get-contact-aws-lambda/models"
)

func TestFind(t *testing.T) {
	mockController := gomock.NewController(t)
	mockHandler := mocks.NewMockIGetContactHandler(mockController)
	defer mockController.Finish()

	// FIXTURES
	var (
		req = models.Request{
			Id: "76334e6d-3707-4c82-aa68-c6897140b5ef",
		}
		response = models.Contacts{
			Id:        "76334e6d-3707-4c82-aa68-c6897140b5ef",
			FirstName: "John",
			LastName:  "English",
			Status:    "CREATED",
		}
		empty_request = models.Request{
			Id: " ",
		}
		empty_response = models.Contacts{}
		errorMsg       = errEmptyData
		ctx            = context.Background()
		fake_id        = models.Request{
			Id: "76334e6d-3707-4c82-aa68-c6897140b5sx",
		}
		errorFakeID = HandleCustomValidationError("Item Not Found!")
	)

	t.Run("When you sent a correct params", func(t *testing.T) {
		gomock.InOrder(
			mockHandler.EXPECT().Find(ctx, req).Return(response, nil).Times(1),
		)

		result, err := mockHandler.Find(ctx, req)
		if assert.NoError(t, err) {
			assert.Equal(t, response, result)
		}
	})

	t.Run("When you sent a empty params", func(t *testing.T) {
		gomock.InOrder(
			mockHandler.EXPECT().Find(ctx, empty_request).Return(empty_response, errorMsg).Times(1),
		)

		result, err := mockHandler.Find(ctx, empty_request)
		if assert.Error(t, err) {
			assert.Equal(t, empty_response, result)
			assert.Equal(t, errorMsg.Error(), err.Error())
		}
	})

	t.Run("When you sent a empty params", func(t *testing.T) {
		gomock.InOrder(
			mockHandler.EXPECT().Find(ctx, fake_id).Return(empty_response, errorFakeID).Times(1),
		)

		result, err := mockHandler.Find(ctx, fake_id)
		if assert.Error(t, err) {
			assert.Equal(t, empty_response, result)
			assert.Equal(t, errorFakeID.Error(), err.Error())
		}
	})

}
