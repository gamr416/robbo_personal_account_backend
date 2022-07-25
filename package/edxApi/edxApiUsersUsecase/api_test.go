package edxApiUsersUsecase

import (
	"github.com/go-playground/assert/v2"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/config"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/edxApi"
	"log"
	"testing"
)

func TestGetUser(t *testing.T) {
	if err := config.InitForTests(); err != nil {
		log.Fatalf("%s", err.Error())
	}
	edx := SetupEdxApiUser()

	expect := []byte("{\"username\":\"edxsom\"}")
	correct, _ := edx.GetUser()
	assert.Equal(t, expect, correct)

}

func TestEdxApiUseCaseImpl_Login(t *testing.T) {
	if err := config.InitForTests(); err != nil {
		log.Fatalf("%s", err.Error())
	}
	edx := SetupEdxApiUser()
	testTable := []struct {
		name          string
		email         string
		password      string
		expectedError error
	}{
		{
			name:          "Ok",
			email:         "edxsom@test.com",
			password:      "123456",
			expectedError: nil,
		},

		{
			name:          "Email or password incorrect",
			email:         "dsadddas",
			password:      "dsadad",
			expectedError: edxApi.ErrIncorrectInputParam,
		},
		{
			name:          "Email or password is empty",
			email:         "",
			password:      "",
			expectedError: edxApi.ErrIncorrectInputParam,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			expect := testCase.expectedError
			_, correct := edx.Login(testCase.email, testCase.password)

			assert.Equal(t, expect, correct)
		})
	}
}

func TestEdxApiUseCaseImpl_PostRegistration(t *testing.T) {
	if err := config.InitForTests(); err != nil {
		log.Fatalf("%s", err.Error())
	}
	edx := SetupEdxApiUser()
	testTable := []struct {
		name                string
		registrationMessage edxApi.RegistrationForm
		expectedError       error
	}{
		{
			name: "Ok",
			registrationMessage: edxApi.RegistrationForm{
				Email:            "insomnia_testrrwds323fsd22dasf3@fake.email",
				Username:         "InsomniaTedasd122fsdfda3",
				Name:             "SomeTestNafdsme12ddsds3",
				Password:         "123456",
				Terms_of_service: "true",
			},
			expectedError: nil,
		},

		{
			name: "Password is empty",
			registrationMessage: edxApi.RegistrationForm{
				Email:            "insomnia_testrrw223@fake.email",
				Username:         "InsomniaTest31223",
				Name:             "SomeTestName123",
				Password:         "",
				Terms_of_service: "true",
			},
			expectedError: edxApi.ErrIncorrectInputParam,
		},
		{
			name: "Email is empty",
			registrationMessage: edxApi.RegistrationForm{
				Email:            "",
				Username:         "InsomniaTest31223",
				Name:             "SomeTestName123",
				Password:         "123456",
				Terms_of_service: "true",
			},
			expectedError: edxApi.ErrIncorrectInputParam,
		},
		{
			name: "Username is empty",
			registrationMessage: edxApi.RegistrationForm{
				Email:            "nsomnia_testrrw223@fake.email",
				Username:         "",
				Name:             "SomeTestName123",
				Password:         "123456",
				Terms_of_service: "true",
			},
			expectedError: edxApi.ErrIncorrectInputParam,
		},
		{
			name: "Name is empty",
			registrationMessage: edxApi.RegistrationForm{
				Email:            "nsomnia_testrrw223@fake.email",
				Username:         "dsadasd",
				Name:             "",
				Password:         "123456",
				Terms_of_service: "true",
			},
			expectedError: edxApi.ErrIncorrectInputParam,
		},
		{
			name: "Terms_of_service is empty",
			registrationMessage: edxApi.RegistrationForm{
				Email:            "nsomnia_testrrw223@fake.email",
				Username:         "dsadasd",
				Name:             "gdgsdfsfs",
				Password:         "123456",
				Terms_of_service: "",
			},
			expectedError: edxApi.ErrIncorrectInputParam,
		},

		{
			name: "All params is empty",
			registrationMessage: edxApi.RegistrationForm{
				Email:            "",
				Username:         "",
				Name:             "",
				Password:         "",
				Terms_of_service: "",
			},
			expectedError: edxApi.ErrIncorrectInputParam,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			expect := testCase.expectedError
			_, correct := edx.PostRegistration(testCase.registrationMessage)
			assert.Equal(t, expect, correct)
		})
	}

}
