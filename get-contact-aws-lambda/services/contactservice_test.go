package services

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/wilomendez/Onboarding/get-contact-aws-lambda/mocks"
	"github.com/wilomendez/Onboarding/get-contact-aws-lambda/models"
	"github.com/wilomendez/Onboarding/get-contact-aws-lambda/utils"
)

func TestFind(t *testing.T) {
	mockController := gomock.NewController(t)
	mockContactRepo := mocks.NewMockIContactsRepo(mockController)
	mockContactService := New(mockContactRepo)
	defer mockController.Finish()

	var (
		id      = "76334e6d-3707-4c82-aa68-c6897140b5ef"
		fake_id = "76334e6d-3707-4c82-aa68-c6897140b5sx"
		contact = models.Contacts{
			Id:        "76334e6d-3707-4c82-aa68-c6897140b5ef",
			FirstName: "John",
			LastName:  "English",
			Status:    "CREATED",
		}
		empty_contact = models.Contacts{}
		errorMsg      = utils.ValidationError("Not found!")
	)

	t.Run("When found the contact", func(t *testing.T) {
		gomock.InOrder(
			mockContactRepo.EXPECT().Find(id).Return(contact, nil).Times(1),
		)

		result, err := mockContactService.Find(id)
		if assert.NoError(t, err) {
			assert.Equal(t, contact, result)
		}
	})

	t.Run("When the id not found", func(t *testing.T) {
		gomock.InOrder(
			mockContactRepo.EXPECT().Find(fake_id).Return(empty_contact, errorMsg).Times(1),
		)

		result, err := mockContactService.Find(fake_id)
		if assert.Error(t, err) {
			assert.Equal(t, empty_contact, result)
			assert.Equal(t, errorMsg.Error(), err.Error())
		}
	})
}
