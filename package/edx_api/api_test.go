package edx_api

import (
	"strconv"
	"testing"
)

func TestGetUser(t *testing.T) {
	var token string
	RefreshToken(&token)
	testTable := []struct {
		assert string
		tkn    string
	}{
		{
			assert: "{\"username\":\"edxsom\"}",
			tkn:    token,
		},
		{
			assert: "{\"error_code\":\"token_not_provided\",\"developer_message\":\"Invalid token header. No credentials provided.\"}",
			tkn:    "",
		},
		{
			assert: "{\"error_code\":\"token_nonexistent\",\"developer_message\":\"The provided access token does not match any valid tokens.\"}",
			tkn:    "blablabla",
		},
	}
	for _, testCase := range testTable {

		expect, _ := GetUser(testCase.tkn)
		if testCase.assert != expect {
			t.Errorf("Assert: " + testCase.assert + "\n Expect: " + expect + "\n Not equal")
		}
	}
}

func TestGetCourse(t *testing.T) {
	var token string
	RefreshToken(&token)
	testTable := []struct {
		assert   int
		courseId string
		tkn      string
	}{
		{
			assert:   200,
			courseId: "course-v1:TestOrg+02+2022",
			tkn:      token,
		},
		{
			assert:   404,
			courseId: "course-v1:Test",
			tkn:      token,
		},
		{
			assert:   200,
			courseId: "",
			tkn:      token,
		},
		{
			assert:   401,
			courseId: "course-v1:TestOrg+02+2022",
			tkn:      "",
		},
		{
			assert:   401,
			courseId: "course-v1:TestOrg+02+2022",
			tkn:      "dsda",
		},
	}
	for _, testCase := range testTable {
		_, expect := GetCourse(testCase.courseId, testCase.tkn)
		if testCase.assert != expect {
			t.Errorf("Assert: " + strconv.Itoa(testCase.assert) + "\n Expect: " + strconv.Itoa(expect) + "\n Not equal")
		}
	}
}

func TestGetEnrollment(t *testing.T) {
	var token string
	RefreshToken(&token)
	testTable := []struct {
		assert   string
		userName string
		tkn      string
	}{
		{
			assert:   "{\"next\":null,\"previous\":null,\"results\":[{\"created\":\"2022-06-18T21:59:34.558581Z\",\"mode\":\"audit\",\"is_active\":true,\"user\":\"tesr_user\",\"course_id\":\"course-v1:Test_org+01+2022\"}]}",
			userName: "tesr_user",
			tkn:      token,
		},
		{
			assert:   "{\"next\":null,\"previous\":null,\"results\":[{\"created\":\"2022-06-18T21:59:34.558581Z\",\"mode\":\"audit\",\"is_active\":true,\"user\":\"tesr_user\",\"course_id\":\"course-v1:Test_org+01+2022\"},{\"created\":\"2022-06-13T03:00:12.571664Z\",\"mode\":\"honor\",\"is_active\":true,\"user\":\"edxsom\",\"course_id\":\"course-v1:TestOrg+02+2022\"},{\"created\":\"2022-06-13T01:16:45.374794Z\",\"mode\":\"honor\",\"is_active\":true,\"user\":\"edxsom\",\"course_id\":\"course-v1:Test_org+01+2022\"}]}",
			userName: "",
			tkn:      token,
		},
		{
			assert:   "{\"next\":null,\"previous\":null,\"results\":[]}",
			userName: "blablabla",
			tkn:      token,
		},
		{
			assert:   "{\"developer_message\":{\"error_code\":\"token_not_provided\",\"developer_message\":\"Invalid token header. No credentials provided.\"}}",
			userName: "blablabla",
			tkn:      "",
		},
		{
			assert:   "{\"developer_message\":{\"error_code\":\"token_nonexistent\",\"developer_message\":\"The provided access token does not match any valid tokens.\"}}",
			userName: "blablabla",
			tkn:      "blablabla",
		},
	}
	for _, testCase := range testTable {
		expect, _ := GetEnrollment(testCase.userName, testCase.tkn)
		if testCase.assert != expect {
			t.Errorf("Assert: " + testCase.assert + "\n Expect: " + expect + "\n Not equal")
		}
	}
}

func TestPostEnrollment(t *testing.T) {
	var token string
	RefreshToken(&token)
	testTable := []struct {
		assert            int
		EnrollmentMessage map[string]interface{}
		tkn               string
	}{
		{
			assert: 200,
			EnrollmentMessage: map[string]interface{}{
				"course_details": map[string]string{
					"course_id": "course-v1:Test_org+01+2022",
				},
				"user": "tesr_user",
			},
			tkn: token,
		},
		{
			assert: 406,
			EnrollmentMessage: map[string]interface{}{
				"course_details": map[string]string{
					"course_id": "course-v1:Test_org+01+2022",
				},
				"user": "tesr_u",
			},
			tkn: token,
		},
		{
			assert: 400,
			EnrollmentMessage: map[string]interface{}{
				"course_details": map[string]string{
					"course_id": "course-v1:T",
				},
				"user": "tesr_user",
			},
			tkn: token,
		},
		{
			assert: 401,
			EnrollmentMessage: map[string]interface{}{
				"course_details": map[string]string{
					"course_id": "course-v1:T",
				},
				"user": "tesr_user",
			},
			tkn: "",
		},
		{
			assert: 401,
			EnrollmentMessage: map[string]interface{}{
				"course_details": map[string]string{
					"course_id": "course-v1:T",
				},
				"user": "tesr_user",
			},
			tkn: "blablabla",
		},
	}
	url := "https://edx-test.ru/api/enrollment/v1/enrollment"
	for _, testCase := range testTable {
		_, expect := PostEnrollment(url, testCase.tkn, testCase.EnrollmentMessage)
		if testCase.assert != expect {
			t.Errorf("Assert: " + strconv.Itoa(testCase.assert) + "\n Expect: " + strconv.Itoa(expect) + "\n Not equal")
		}
	}
}
