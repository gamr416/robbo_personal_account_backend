package edxApiCohortsUsecase

import (
	"github.com/go-playground/assert/v2"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/config"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/edxApi"
	"log"
	"testing"
)

func TestEdxApiCohortImpl_CreateCohort(t *testing.T) {
	if err := config.InitForTests(); err != nil {
		log.Fatalf("%s", err.Error())
	}
	edx := SetupEdxApiCohort()
	testTable := []struct {
		name          string
		message       map[string]interface{}
		courseId      string
		expectedError error
	}{
		{
			name:     "Ok",
			courseId: "course-v1:TestOrg+02+2022",
			message: map[string]interface{}{
				"name":            "cohorTestName23243",
				"assignment_type": "Manual",
			},
			expectedError: nil,
		},
		{
			name:     "Empty cohort Name",
			courseId: "course-v1:TestOrg+02+2022",
			message: map[string]interface{}{
				"name":            "",
				"assignment_type": "Manual",
			},
			expectedError: edxApi.ErrIncorrectInputParam,
		},
		{
			name:     "Assignment type is empty",
			courseId: "course-v1:TestOrg+02+2022",
			message: map[string]interface{}{
				"name":            "cohorTestName2323",
				"assignment_type": "",
			},
			expectedError: edxApi.ErrIncorrectInputParam,
		},
		{
			name:     "Name is taken",
			courseId: "course-v1:TestOrg+02+2022",
			message: map[string]interface{}{
				"name":            "cohorTestName",
				"assignment_type": "Manual",
			},
			expectedError: edxApi.ErrIncorrectInputParam,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			expect := testCase.expectedError
			_, correct := edx.CreateCohort(testCase.courseId, testCase.message)
			assert.Equal(t, expect, correct)
		})
	}
}

func TestEdxApiCohortImpl_AddStudent(t *testing.T) {
	if err := config.InitForTests(); err != nil {
		log.Fatalf("%s", err.Error())
	}
	edx := SetupEdxApiCohort()
	testTable := []struct {
		name          string
		courseId      string
		username      string
		cohortId      int
		expectedError error
	}{
		{
			name:          "Ok",
			courseId:      "course-v1:TestOrg+02+2022",
			username:      "edxsom",
			cohortId:      1,
			expectedError: nil,
		},
		{
			name:          "Invalid cohortId",
			courseId:      "course-v1:TestOrg+02+2022",
			username:      "edxsom",
			cohortId:      100,
			expectedError: edxApi.ErrIncorrectInputParam,
		},
		{
			name:          "Invalid courseId",
			courseId:      "course-v1:T+02+2022",
			username:      "edxsom",
			cohortId:      1,
			expectedError: edxApi.ErrIncorrectInputParam,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			expect := testCase.expectedError
			_, correct := edx.AddStudent(testCase.username, testCase.courseId, testCase.cohortId)
			assert.Equal(t, expect, correct)
		})
	}
}
