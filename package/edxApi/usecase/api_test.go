package usecase

import (
	assert "github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	mock_edxApi "github.com/skinnykaen/robbo_student_personal_account.git/package/edxApi/mocks"
	"testing"
)

func TestGetUser(t *testing.T) {

	t.Run("OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockapi := mock_edxApi.NewMockEdxApiUseCase(ctrl)

		expect := "{\n\t\"username\": \"edxsom\"\n}"
		mockapi.EXPECT().GetUser().Return("{\n\t\"username\": \"edxsom\"\n}", nil)
		testApi := &EdxApiUseCaseModule{EdxApiUseCase: mockapi}
		correct, _ := testApi.GetUser()
		t.Logf(correct)
		assert.Equal(t, expect, correct)
	})
}
