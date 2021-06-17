package services

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/wilomendez/Onboarding/create-contact-aws-lambda/mocks"
	"github.com/wilomendez/Onboarding/create-contact-aws-lambda/models"
	//"github.com/wilomendez/Onboarding/create-contact-aws-lambda/utils"
)

func TestCreate(t *testing.T) {
	mockController := gomock.NewController(t)
	mockContactRepo := mocks.NewMockIContactsRepo(mockController)
	mockContactService := New(mockContactRepo)
	defer mockController.Finish()

	// Fixtures
	var (
		response = models.Response{
			ContactNumber: "76334e6d-3707-4c82-aa68-c6897140b5ef",
		}
		contact = models.Contacts{
			Id:        "76334e6d-3707-4c82-aa68-c6897140b5ef",
			FirstName: "John",
			LastName:  "English",
			Status:    "CREATED",
		}
		// exceptError   = utils.ValidationError("Input validation FAILED, First Name or Last name is empty")
		// emptyContact  = models.Contacts{}
		// emptyResponse = models.Response{}
	)

	t.Run("CREATE", func(t *testing.T) {
		t.Run("When we receive correct data", func(t *testing.T) {
			gomock.InOrder(
				mockContactRepo.EXPECT().Create(contact).Return(response, nil).Times(1),
			)

			result, err := mockContactService.Create(contact)
			if assert.NoError(t, err) {
				assert.Equal(t, response, result)
			}
		})

		// t.Run("When we receive incorrect data", func(t *testing.T) {
		// 	gomock.InOrder(
		// 		mockContactRepo.EXPECT().Create(emptyContact).Return(emptyResponse, exceptError).Times(1),
		// 	)

		// 	_, err := mockContactService.Create(emptyContact)
		// 	if assert.Error(t, err) {
		// 		assert.Equal(t, exceptError.Error(), err.Error())
		// 	}
		// })
	})
}
