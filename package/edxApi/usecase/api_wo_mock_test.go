package usecase

import (
	"errors"
	"github.com/go-playground/assert/v2"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/config"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/edxApi"
	"log"
	"testing"
)

func TestGetUser2(t *testing.T) {
	if err := config.InitForTests(); err != nil {
		log.Fatalf("%s", err.Error())
	}
	edx := SetupEdxApiUseCase()

	expect := []byte("{\"username\":\"edxsom\"}")
	correct, _ := edx.GetUser()
	assert.Equal(t, expect, correct)

}

func TestEdxApiUseCaseImpl_GetCourseContent2(t *testing.T) {
	if err := config.InitForTests(); err != nil {
		log.Fatalf("%s", err.Error())
	}
	edx := SetupEdxApiUseCase()
	testTable := []struct {
		name          string
		courseId      string
		expectedError error
	}{
		{
			name:          "Ok",
			courseId:      "course-v1:Test_org+01+2022",
			expectedError: nil,
		},

		{
			name:          "Bad courseId",
			courseId:      "Ddssadad",
			expectedError: edxApi.ErrIncorrectInputParam,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			expect := testCase.expectedError
			_, correct := edx.GetCourseContent(testCase.courseId)
			assert.Equal(t, expect, correct)
		})
	}
}

func TestEdxApiUseCaseImpl_GetEnrollments2(t *testing.T) {
	if err := config.InitForTests(); err != nil {
		log.Fatalf("%s", err.Error())
	}
	edx := SetupEdxApiUseCase()
	testTable := []struct {
		name         string
		username     string
		expectedBody string
	}{
		{
			name:         "Ok",
			username:     "edxsom",
			expectedBody: "{\"next\":null,\"previous\":null,\"results\":[{\"created\":\"2022-06-13T03:00:12.571664Z\",\"mode\":\"honor\",\"is_active\":true,\"user\":\"edxsom\",\"course_id\":\"course-v1:TestOrg+02+2022\"},{\"created\":\"2022-06-13T01:16:45.374794Z\",\"mode\":\"honor\",\"is_active\":true,\"user\":\"edxsom\",\"course_id\":\"course-v1:Test_org+01+2022\"}]}",
		},

		{
			name:         "Bad username",
			username:     "dsad",
			expectedBody: "{\"next\":null,\"previous\":null,\"results\":[]}",
		},
		{
			name:         "Empty username",
			username:     "",
			expectedBody: "{\"next\":null,\"previous\":null,\"results\":[{\"created\":\"2022-06-18T21:59:34.558581Z\",\"mode\":\"audit\",\"is_active\":true,\"user\":\"tesr_user\",\"course_id\":\"course-v1:Test_org+01+2022\"},{\"created\":\"2022-06-13T03:00:12.571664Z\",\"mode\":\"honor\",\"is_active\":true,\"user\":\"edxsom\",\"course_id\":\"course-v1:TestOrg+02+2022\"},{\"created\":\"2022-06-13T01:16:45.374794Z\",\"mode\":\"honor\",\"is_active\":true,\"user\":\"edxsom\",\"course_id\":\"course-v1:Test_org+01+2022\"}]}",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			expect := testCase.expectedBody

			correct, _ := edx.GetEnrollments(testCase.username)
			assert.Equal(t, expect, string(correct))
		})
	}
}

func TestEdxApiUseCaseImpl_GetAllPublicCourses2(t *testing.T) {
	if err := config.InitForTests(); err != nil {
		log.Fatalf("%s", err.Error())
	}
	edx := SetupEdxApiUseCase()
	testTable := []struct {
		name          string
		pageNumber    int
		expectedError error
	}{
		{
			name:          "Ok",
			pageNumber:    1,
			expectedError: nil,
		},

		{
			name:          "Page number is 0",
			pageNumber:    0,
			expectedError: errors.New("user not found"),
		},
		{
			name:          "Page number more then page count",
			pageNumber:    423423423,
			expectedError: errors.New("user not found"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			expect := testCase.expectedError

			_, correct := edx.GetAllPublicCourses(testCase.pageNumber)
			assert.Equal(t, expect, correct)
		})
	}
}

func TestEdxApiUseCaseImpl_Login2(t *testing.T) {
	if err := config.InitForTests(); err != nil {
		log.Fatalf("%s", err.Error())
	}
	edx := SetupEdxApiUseCase()
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

func TestEdxApiUseCaseImpl_PostEnrollment2(t *testing.T) {
	if err := config.InitForTests(); err != nil {
		log.Fatalf("%s", err.Error())
	}
	edx := SetupEdxApiUseCase()
	testTable := []struct {
		name          string
		message       map[string]interface{}
		expectedError error
	}{
		{
			name: "Ok",
			message: map[string]interface{}{
				"course_details": map[string]string{
					"course_id": "course-v1:Test_org+01+2022",
				},
				"user": "edxsom",
			},
			expectedError: nil,
		},

		{
			name: "Course id incorrect",
			message: map[string]interface{}{
				"course_details": map[string]string{
					"course_id": "dasda",
				},
				"user": "edxsom",
			},
			expectedError: edxApi.ErrIncorrectInputParam,
		},
		{
			name: "Username incorrect",
			message: map[string]interface{}{
				"course_details": map[string]string{
					"course_id": "course-v1:Test_org+01+2022",
				},
				"user": "edm",
			},
			expectedError: edxApi.ErrIncorrectInputParam,
		},
		{
			name: "Empty field courseId",
			message: map[string]interface{}{
				"course_details": map[string]string{
					"course_id": "",
				},
				"user": "edxsom",
			},
			expectedError: edxApi.ErrIncorrectInputParam,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			expect := testCase.expectedError

			_, correct := edx.PostEnrollment(testCase.message)
			assert.Equal(t, expect, correct)
		})
	}
}

func TestEdxApiUseCaseImpl_PostRegistration2(t *testing.T) {
	if err := config.InitForTests(); err != nil {
		log.Fatalf("%s", err.Error())
	}
	edx := SetupEdxApiUseCase()
	testTable := []struct {
		name                string
		registrationMessage edxApi.RegistrationForm
		expectedError       error
	}{
		{
			name: "Ok",
			registrationMessage: edxApi.RegistrationForm{
				Email:            "insomnia_testrrw323fsd22dasf3@fake.email",
				Username:         "InsomniaTest3fsd122fsdfda3",
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
