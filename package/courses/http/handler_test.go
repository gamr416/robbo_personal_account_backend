package http

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/courses/http/mocks"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/edxApi"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"net/http/httptest"
	"testing"
)

func TestHandler_GetCourseContent(t *testing.T) {
	type mockBehavior func(s *mock_courses.MockDelegate, courseId string)

	testTable := []struct {
		name               string
		inputParam         string
		mockBehavior       mockBehavior
		expectedStatusCode int
	}{
		{
			name:       "OK",
			inputParam: "course-v1:Test_org+01+2022",
			mockBehavior: func(s *mock_courses.MockDelegate, courseId string) {
				s.EXPECT().GetCourseContent(courseId).Return([]byte("{\n\t\"blocks_url\": \"https://edx-test.ru/api/courses/v2/blocks/?course_id=course-v1%3ATest_org%2B01%2B2022\",\n\t\"effort\": null,\n\t\"end\": \"2023-02-02T00:00:00Z\",\n\t\"enrollment_start\": null,\n\t\"enrollment_end\": null,\n\t\"id\": \"course-v1:Test_org+01+2022\",\n\t\"media\": {\n\t\t\"banner_image\": {\n\t\t\t\"uri\": \"/asset-v1:Test_org+01+2022+type@asset+block@images_course_image.jpg\",\n\t\t\t\"uri_absolute\": \"https://edx-test.ru/asset-v1:Test_org+01+2022+type@asset+block@images_course_image.jpg\"\n\t\t},\n\t\t\"course_image\": {\n\t\t\t\"uri\": \"/asset-v1:Test_org+01+2022+type@asset+block@images_course_image.jpg\"\n\t\t},\n\t\t\"course_video\": {\n\t\t\t\"uri\": null\n\t\t},\n\t\t\"image\": {\n\t\t\t\"raw\": \"https://edx-test.ru/asset-v1:Test_org+01+2022+type@asset+block@images_course_image.jpg\",\n\t\t\t\"small\": \"https://edx-test.ru/asset-v1:Test_org+01+2022+type@asset+block@images_course_image.jpg\",\n\t\t\t\"large\": \"https://edx-test.ru/asset-v1:Test_org+01+2022+type@asset+block@images_course_image.jpg\"\n\t\t}\n\t},\n\t\"name\": \"Cert test\",\n\t\"number\": \"01\",\n\t\"org\": \"Test_org\",\n\t\"short_description\": \"\",\n\t\"start\": \"2022-01-01T00:00:00Z\",\n\t\"start_display\": \"Jan. 1, 2022\",\n\t\"start_type\": \"timestamp\",\n\t\"pacing\": \"instructor\",\n\t\"mobile_available\": false,\n\t\"hidden\": false,\n\t\"invitation_only\": false,\n\t\"course_id\": \"course-v1:Test_org+01+2022\",\n\t\"overview\": \"<section class=\\\"about\\\">\\n  <h2>About This Course</h2>\\n  <p>Include your long course description here. The long course description should contain 150-400 words.</p>\\n\\n  <p>This is paragraph 2 of the long course description. Add more paragraphs as needed. Make sure to enclose them in paragraph tags.</p>\\n</section>\\n\\n<section class=\\\"prerequisites\\\">\\n  <h2>Requirements</h2>\\n  <p>Add information about the skills and knowledge students need to take this course.</p>\\n</section>\\n\\n<section class=\\\"course-staff\\\">\\n  <h2>Course Staff</h2>\\n  <article class=\\\"teacher\\\">\\n    <div class=\\\"teacher-image\\\">\\n      <img src=\\\"/static/images/placeholder-faculty.png\\\" align=\\\"left\\\" style=\\\"margin:0 20 px 0\\\" alt=\\\"Course Staff Image #1\\\">\\n    </div>\\n\\n    <h3>Staff Member #1</h3>\\n    <p>Biography of instructor/staff member #1</p>\\n  </article>\\n\\n  <article class=\\\"teacher\\\">\\n    <div class=\\\"teacher-image\\\">\\n      <img src=\\\"/static/images/placeholder-faculty.png\\\" align=\\\"left\\\" style=\\\"margin:0 20 px 0\\\" alt=\\\"Course Staff Image #2\\\">\\n    </div>\\n\\n    <h3>Staff Member #2</h3>\\n    <p>Biography of instructor/staff member #2</p>\\n  </article>\\n</section>\\n\\n<section class=\\\"faq\\\">\\n  <section class=\\\"responses\\\">\\n    <h2>Frequently Asked Questions</h2>\\n    <article class=\\\"response\\\">\\n      <h3>What web browser should I use?</h3>\\n      <p>The Open edX platform works best with current versions of Chrome, Edge, Firefox, Internet Explorer, or Safari.</p>\\n      <p>See our <a href=\\\"https://edx.readthedocs.org/projects/open-edx-learner-guide/en/latest/front_matter/browsers.html\\\">list of supported browsers</a> for the most up-to-date information.</p>\\n    </article>\\n\\n    <article class=\\\"response\\\">\\n      <h3>Question #2</h3>\\n      <p>Your answer would be displayed here.</p>\\n    </article>\\n  </section>\\n</section>\\n\"\n}"), nil)
			},
			expectedStatusCode: 200,
		},
		{
			name:       "Incorrect course id",
			inputParam: "fsd",
			mockBehavior: func(s *mock_courses.MockDelegate, courseId string) {
				s.EXPECT().GetCourseContent(courseId).Return(nil, edxApi.ErrIncorrectInputParam)
			},
			expectedStatusCode: 500,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			edx := mock_courses.NewMockDelegate(c)
			testCase.mockBehavior(edx, testCase.inputParam)

			handler := NewCoursesHandler(edx)

			r := gin.New()
			r.GET("/getCourseContent/:courseId", handler.GetCourseContent, func(c *gin.Context) {
				c.Param(testCase.inputParam)
			})

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/getCourseContent/"+testCase.inputParam, nil)
			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
		})
	}
}

func TestHandler_CreateCourse(t *testing.T) {
	type mockBehavior func(s *mock_courses.MockDelegate, course models.CourseHTTP, courseId string)

	testTable := []struct {
		name               string
		inputParam         string
		inputCourse        models.CourseHTTP
		mockBehavior       mockBehavior
		expectedStatusCode int
	}{
		{
			name:        "OK",
			inputParam:  "course-v1:Test_org+01+2022",
			inputCourse: models.CourseHTTP{},
			mockBehavior: func(s *mock_courses.MockDelegate, course models.CourseHTTP, courseId string) {
				s.EXPECT().CreateCourse(&course, courseId).Return("1", nil)
			},
			expectedStatusCode: 200,
		},
		{
			name:        "Incorrect course id",
			inputParam:  "fsd",
			inputCourse: models.CourseHTTP{},
			mockBehavior: func(s *mock_courses.MockDelegate, course models.CourseHTTP, courseId string) {
				s.EXPECT().CreateCourse(&course, courseId).Return("0", edxApi.ErrIncorrectInputParam)
			},
			expectedStatusCode: 500,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			edx := mock_courses.NewMockDelegate(c)
			testCase.mockBehavior(edx, testCase.inputCourse, testCase.inputParam)

			handler := NewCoursesHandler(edx)

			r := gin.New()
			r.POST("/createCourse/:courseId", handler.CreateCourse, func(c *gin.Context) {
				c.Param(testCase.inputParam)
			})

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/createCourse/"+testCase.inputParam, nil)
			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
		})
	}
}

func TestHandler_GetAllPublicCourses(t *testing.T) {
	type mockBehavior func(s *mock_courses.MockDelegate, pageNumber int)

	testTable := []struct {
		name                 string
		inputParamForHandler string
		inputParamForMock    int
		mockBehavior         mockBehavior
		expectedStatusCode   int
	}{
		{
			name:                 "OK",
			inputParamForHandler: "1",
			inputParamForMock:    1,
			mockBehavior: func(s *mock_courses.MockDelegate, pageNumber int) {
				s.EXPECT().GetAllPublicCourses(pageNumber).Return([]byte("{\n\t\"results\": [\n\t\t{\n\t\t\t\"blocks_url\": \"https://courses.edx.org/api/courses/v2/blocks/?course_id=ccx-v1%3Aadam%2BMac_APccx%2Be0d%2Bccx%403\",\n\t\t\t\"effort\": \"2:00\",\n\t\t\t\"end\": null,\n\t\t\t\"enrollment_start\": \"2015-07-21T10:00:00Z\",\n\t\t\t\"enrollment_end\": \"2016-06-29T10:00:00Z\",\n\t\t\t\"id\": \"ccx-v1:adam+Mac_APccx+e0d+ccx@3\",\n\t\t\t\"media\": {\n\t\t\t\t\"banner_image\": {\n\t\t\t\t\t\"uri\": \"/asset-v1:adam+Mac_APccx+e0d+type@asset+block@images_course_image.jpg\",\n\t\t\t\t\t\"uri_absolute\": \"https://courses.edx.org/asset-v1:adam+Mac_APccx+e0d+type@asset+block@images_course_image.jpg\"\n\t\t\t\t},\n\t\t\t\t\"course_image\": {\n\t\t\t\t\t\"uri\": \"/asset-v1:adam+Mac_APccx+e0d+type@asset+block@Davidson_EdX-68_1_.png\"\n\t\t\t\t},\n\t\t\t\t\"course_video\": {\n\t\t\t\t\t\"uri\": \"http://www.youtube.com/watch?v=sAnHwOL8aAs\"\n\t\t\t\t},\n\t\t\t\t\"image\": {\n\t\t\t\t\t\"raw\": \"https://courses.edx.org/asset-v1:adam+Mac_APccx+e0d+type@asset+block@Davidson_EdX-68_1_.png\",\n\t\t\t\t\t\"small\": \"https://courses.edx.org/asset-v1:adam+Mac_APccx+e0d+type@thumbnail+block@Davidson_EdX-68_1_-png-375x200.jpg\",\n\t\t\t\t\t\"large\": \"https://courses.edx.org/asset-v1:adam+Mac_APccx+e0d+type@thumbnail+block@Davidson_EdX-68_1_-png-750x400.jpg\"\n\t\t\t\t}\n\t\t\t},\n\t\t\t\"name\": \"ntest\",\n\t\t\t\"number\": \"Mac_APccx\",\n\t\t\t\"org\": \"adam\",\n\t\t\t\"short_description\": \"\",\n\t\t\t\"start\": \"2016-06-27T14:10:44Z\",\n\t\t\t\"start_display\": \"June 27, 2016\",\n\t\t\t\"start_type\": \"timestamp\",\n\t\t\t\"pacing\": \"instructor\",\n\t\t\t\"mobile_available\": false,\n\t\t\t\"hidden\": false,\n\t\t\t\"invitation_only\": false,\n\t\t\t\"course_id\": \"ccx-v1:adam+Mac_APccx+e0d+ccx@3\"\n\t\t},\n\t\t{\n\t\t\t\"blocks_url\": \"https://courses.edx.org/api/courses/v2/blocks/?course_id=ccx-v1%3Aadam%2BMac_APccx%2Be0d%2Bccx%404\",\n\t\t\t\"effort\": \"2:00\",\n\t\t\t\"end\": null,\n\t\t\t\"enrollment_start\": \"2015-07-21T10:00:00Z\",\n\t\t\t\"enrollment_end\": \"2016-06-29T10:00:00Z\",\n\t\t\t\"id\": \"ccx-v1:adam+Mac_APccx+e0d+ccx@4\",\n\t\t\t\"media\": {\n\t\t\t\t\"banner_image\": {\n\t\t\t\t\t\"uri\": \"/asset-v1:adam+Mac_APccx+e0d+type@asset+block@images_course_image.jpg\",\n\t\t\t\t\t\"uri_absolute\": \"https://courses.edx.org/asset-v1:adam+Mac_APccx+e0d+type@asset+block@images_course_image.jpg\"\n\t\t\t\t},\n\t\t\t\t\"course_image\": {\n\t\t\t\t\t\"uri\": \"/asset-v1:adam+Mac_APccx+e0d+type@asset+block@Davidson_EdX-68_1_.png\"\n\t\t\t\t},\n\t\t\t\t\"course_video\": {\n\t\t\t\t\t\"uri\": \"http://www.youtube.com/watch?v=sAnHwOL8aAs\"\n\t\t\t\t},\n\t\t\t\t\"image\": {\n\t\t\t\t\t\"raw\": \"https://courses.edx.org/asset-v1:adam+Mac_APccx+e0d+type@asset+block@Davidson_EdX-68_1_.png\",\n\t\t\t\t\t\"small\": \"https://courses.edx.org/asset-v1:adam+Mac_APccx+e0d+type@thumbnail+block@Davidson_EdX-68_1_-png-375x200.jpg\",\n\t\t\t\t\t\"large\": \"https://courses.edx.org/asset-v1:adam+Mac_APccx+e0d+type@thumbnail+block@Davidson_EdX-68_1_-png-750x400.jpg\"\n\t\t\t\t}\n\t\t\t},\n\t\t\t\"name\": \"ntest\",\n\t\t\t\"number\": \"Mac_APccx\",\n\t\t\t\"org\": \"adam\",\n\t\t\t\"short_description\": \"\",\n\t\t\t\"start\": \"2016-06-27T14:10:48Z\",\n\t\t\t\"start_display\": \"June 27, 2016\",\n\t\t\t\"start_type\": \"timestamp\",\n\t\t\t\"pacing\": \"instructor\",\n\t\t\t\"mobile_available\": false,\n\t\t\t\"hidden\": false,\n\t\t\t\"invitation_only\": false,\n\t\t\t\"course_id\": \"ccx-v1:adam+Mac_APccx+e0d+ccx@4\"\n\t\t},\n\t\t{\n\t\t\t\"blocks_url\": \"https://courses.edx.org/api/courses/v2/blocks/?course_id=ccx-v1%3ABerkeleyX%2BCS169.1x%2B3T2015SP%2Bccx%4011\",\n\t\t\t\"effort\": \"12 hours/week\",\n\t\t\t\"end\": null,\n\t\t\t\"enrollment_start\": null,\n\t\t\t\"enrollment_end\": \"2016-07-18T00:00:00Z\",\n\t\t\t\"id\": \"ccx-v1:BerkeleyX+CS169.1x+3T2015SP+ccx@11\",\n\t\t\t\"media\": {\n\t\t\t\t\"banner_image\": {\n\t\t\t\t\t\"uri\": \"/asset-v1:BerkeleyX+CS169.1x+3T2015SP+type@asset+block@images_course_image.jpg\",\n\t\t\t\t\t\"uri_absolute\": \"https://courses.edx.org/asset-v1:BerkeleyX+CS169.1x+3T2015SP+type@asset+block@images_course_image.jpg\"\n\t\t\t\t},\n\t\t\t\t\"course_image\": {\n\t\t\t\t\t\"uri\": \"/asset-v1:BerkeleyX+CS169.1x+3T2015SP+type@asset+block@ruby-rails-programming-100354527-primary.idge.jpg\"\n\t\t\t\t},\n\t\t\t\t\"course_video\": {\n\t\t\t\t\t\"uri\": \"http://www.youtube.com/watch?v=V36LpHqtcDY\"\n\t\t\t\t},\n\t\t\t\t\"image\": {\n\t\t\t\t\t\"raw\": \"https://courses.edx.org/asset-v1:BerkeleyX+CS169.1x+3T2015SP+type@asset+block@ruby-rails-programming-100354527-primary.idge.jpg\",\n\t\t\t\t\t\"small\": \"https://courses.edx.org/asset-v1:BerkeleyX+CS169.1x+3T2015SP+type@thumbnail+block@ruby-rails-programming-100354527-primary.idge-375x200.jpg\",\n\t\t\t\t\t\"large\": \"https://courses.edx.org/asset-v1:BerkeleyX+CS169.1x+3T2015SP+type@thumbnail+block@ruby-rails-programming-100354527-primary.idge-750x400.jpg\"\n\t\t\t\t}\n\t\t\t},\n\t\t\t\"name\": \"Grinnell CSC321 2016F\",\n\t\t\t\"number\": \"CS169.1x\",\n\t\t\t\"org\": \"BerkeleyX\",\n\t\t\t\"short_description\": \"CS169.1x teaches the fundamentals for engineering long-lived software using Agile techniques to develop Software as a Service (SaaS) using Ruby on Rails.\",\n\t\t\t\"start\": \"2016-06-07T20:00:00Z\",\n\t\t\t\"start_display\": \"June 7, 2016\",\n\t\t\t\"start_type\": \"timestamp\",\n\t\t\t\"pacing\": \"instructor\",\n\t\t\t\"mobile_available\": true,\n\t\t\t\"hidden\": false,\n\t\t\t\"invitation_only\": false,\n\t\t\t\"course_id\": \"ccx-v1:BerkeleyX+CS169.1x+3T2015SP+ccx@11\"\n\t\t},\n\t\t{\n\t\t\t\"blocks_url\": \"https://courses.edx.org/api/courses/v2/blocks/?course_id=ccx-v1%3ABerkeleyX%2BCS169.1x%2B3T2015SP%2Bccx%405\",\n\t\t\t\"effort\": \"12 hours/week\",\n\t\t\t\"end\": null,\n\t\t\t\"enrollment_start\": null,\n\t\t\t\"enrollment_end\": \"2016-07-18T00:00:00Z\",\n\t\t\t\"id\": \"ccx-v1:BerkeleyX+CS169.1x+3T2015SP+ccx@5\",\n\t\t\t\"media\": {\n\t\t\t\t\"banner_image\": {\n\t\t\t\t\t\"uri\": \"/asset-v1:BerkeleyX+CS169.1x+3T2015SP+type@asset+block@images_course_image.jpg\",\n\t\t\t\t\t\"uri_absolute\": \"https://courses.edx.org/asset-v1:BerkeleyX+CS169.1x+3T2015SP+type@asset+block@images_course_image.jpg\"\n\t\t\t\t},\n\t\t\t\t\"course_image\": {\n\t\t\t\t\t\"uri\": \"/asset-v1:BerkeleyX+CS169.1x+3T2015SP+type@asset+block@ruby-rails-programming-100354527-primary.idge.jpg\"\n\t\t\t\t},\n\t\t\t\t\"course_video\": {\n\t\t\t\t\t\"uri\": \"http://www.youtube.com/watch?v=V36LpHqtcDY\"\n\t\t\t\t},\n\t\t\t\t\"image\": {\n\t\t\t\t\t\"raw\": \"https://courses.edx.org/asset-v1:BerkeleyX+CS169.1x+3T2015SP+type@asset+block@ruby-rails-programming-100354527-primary.idge.jpg\",\n\t\t\t\t\t\"small\": \"https://courses.edx.org/asset-v1:BerkeleyX+CS169.1x+3T2015SP+type@thumbnail+block@ruby-rails-programming-100354527-primary.idge-375x200.jpg\",\n\t\t\t\t\t\"large\": \"https://courses.edx.org/asset-v1:BerkeleyX+CS169.1x+3T2015SP+type@thumbnail+block@ruby-rails-programming-100354527-primary.idge-750x400.jpg\"\n\t\t\t\t}\n\t\t\t},\n\t\t\t\"name\": \"Test\",\n\t\t\t\"number\": \"CS169.1x\",\n\t\t\t\"org\": \"BerkeleyX\",\n\t\t\t\"short_description\": \"CS169.1x teaches the fundamentals for engineering long-lived software using Agile techniques to develop Software as a Service (SaaS) using Ruby on Rails.\",\n\t\t\t\"start\": \"2016-05-24T20:00:00Z\",\n\t\t\t\"start_display\": \"May 24, 2016\",\n\t\t\t\"start_type\": \"timestamp\",\n\t\t\t\"pacing\": \"instructor\",\n\t\t\t\"mobile_available\": true,\n\t\t\t\"hidden\": false,\n\t\t\t\"invitation_only\": false,\n\t\t\t\"course_id\": \"ccx-v1:BerkeleyX+CS169.1x+3T2015SP+ccx@5\"\n\t\t},\n\t\t{\n\t\t\t\"blocks_url\": \"https://courses.edx.org/api/courses/v2/blocks/?course_id=ccx-v1%3ABUx%2BPY1x%2B3T2016%2Bccx%4012\",\n\t\t\t\"effort\": null,\n\t\t\t\"end\": null,\n\t\t\t\"enrollment_start\": \"2016-08-08T18:00:00Z\",\n\t\t\t\"enrollment_end\": null,\n\t\t\t\"id\": \"ccx-v1:BUx+PY1x+3T2016+ccx@12\",\n\t\t\t\"media\": {\n\t\t\t\t\"banner_image\": {\n\t\t\t\t\t\"uri\": \"/asset-v1:BUx+PY1x+3T2016+type@asset+block@images_course_image.jpg\",\n\t\t\t\t\t\"uri_absolute\": \"https://courses.edx.org/asset-v1:BUx+PY1x+3T2016+type@asset+block@images_course_image.jpg\"\n\t\t\t\t},\n\t\t\t\t\"course_image\": {\n\t\t\t\t\t\"uri\": \"/asset-v1:BUx+PY1x+3T2016+type@asset+block@AP_Physics_Dashboard_Image.png\"\n\t\t\t\t},\n\t\t\t\t\"course_video\": {\n\t\t\t\t\t\"uri\": null\n\t\t\t\t},\n\t\t\t\t\"image\": {\n\t\t\t\t\t\"raw\": \"https://courses.edx.org/asset-v1:BUx+PY1x+3T2016+type@asset+block@AP_Physics_Dashboard_Image.png\",\n\t\t\t\t\t\"small\": \"https://courses.edx.org/asset-v1:BUx+PY1x+3T2016+type@thumbnail+block@AP_Physics_Dashboard_Image-png-375x200.jpg\",\n\t\t\t\t\t\"large\": \"https://courses.edx.org/asset-v1:BUx+PY1x+3T2016+type@thumbnail+block@AP_Physics_Dashboard_Image-png-750x400.jpg\"\n\t\t\t\t}\n\t\t\t},\n\t\t\t\"name\": \"AHS AP Physics\",\n\t\t\t\"number\": \"PY1x\",\n\t\t\t\"org\": \"BUx\",\n\t\t\t\"short_description\": \"\",\n\t\t\t\"start\": \"2016-09-12T16:14:32Z\",\n\t\t\t\"start_display\": \"Sept. 12, 2016\",\n\t\t\t\"start_type\": \"timestamp\",\n\t\t\t\"pacing\": \"instructor\",\n\t\t\t\"mobile_available\": true,\n\t\t\t\"hidden\": false,\n\t\t\t\"invitation_only\": false,\n\t\t\t\"course_id\": \"ccx-v1:BUx+PY1x+3T2016+ccx@12\"\n\t\t},\n\t\t{\n\t\t\t\"blocks_url\": \"https://courses.edx.org/api/courses/v2/blocks/?course_id=ccx-v1%3ABUx%2BPY1x%2B3T2016%2Bccx%4013\",\n\t\t\t\"effort\": null,\n\t\t\t\"end\": null,\n\t\t\t\"enrollment_start\": \"2016-08-08T18:00:00Z\",\n\t\t\t\"enrollment_end\": null,\n\t\t\t\"id\": \"ccx-v1:BUx+PY1x+3T2016+ccx@13\",\n\t\t\t\"media\": {\n\t\t\t\t\"banner_image\": {\n\t\t\t\t\t\"uri\": \"/asset-v1:BUx+PY1x+3T2016+type@asset+block@images_course_image.jpg\",\n\t\t\t\t\t\"uri_absolute\": \"https://courses.edx.org/asset-v1:BUx+PY1x+3T2016+type@asset+block@images_course_image.jpg\"\n\t\t\t\t},\n\t\t\t\t\"course_image\": {\n\t\t\t\t\t\"uri\": \"/asset-v1:BUx+PY1x+3T2016+type@asset+block@AP_Physics_Dashboard_Image.png\"\n\t\t\t\t},\n\t\t\t\t\"course_video\": {\n\t\t\t\t\t\"uri\": null\n\t\t\t\t},\n\t\t\t\t\"image\": {\n\t\t\t\t\t\"raw\": \"https://courses.edx.org/asset-v1:BUx+PY1x+3T2016+type@asset+block@AP_Physics_Dashboard_Image.png\",\n\t\t\t\t\t\"small\": \"https://courses.edx.org/asset-v1:BUx+PY1x+3T2016+type@thumbnail+block@AP_Physics_Dashboard_Image-png-375x200.jpg\",\n\t\t\t\t\t\"large\": \"https://courses.edx.org/asset-v1:BUx+PY1x+3T2016+type@thumbnail+block@AP_Physics_Dashboard_Image-png-750x400.jpg\"\n\t\t\t\t}\n\t\t\t},\n\t\t\t\"name\": \"AHS AP Physics\",\n\t\t\t\"number\": \"PY1x\",\n\t\t\t\"org\": \"BUx\",\n\t\t\t\"short_description\": \"\",\n\t\t\t\"start\": \"2016-09-12T06:00:00Z\",\n\t\t\t\"start_display\": \"Sept. 12, 2016\",\n\t\t\t\"start_type\": \"timestamp\",\n\t\t\t\"pacing\": \"instructor\",\n\t\t\t\"mobile_available\": true,\n\t\t\t\"hidden\": false,\n\t\t\t\"invitation_only\": false,\n\t\t\t\"course_id\": \"ccx-v1:BUx+PY1x+3T2016+ccx@13\"\n\t\t},\n\t\t{\n\t\t\t\"blocks_url\": \"https://courses.edx.org/api/courses/v2/blocks/?course_id=ccx-v1%3ABUx%2BPY1x%2B3T2016%2Bccx%4014\",\n\t\t\t\"effort\": null,\n\t\t\t\"end\": null,\n\t\t\t\"enrollment_start\": \"2016-08-08T18:00:00Z\",\n\t\t\t\"enrollment_end\": null,\n\t\t\t\"id\": \"ccx-v1:BUx+PY1x+3T2016+ccx@14\",\n\t\t\t\"media\": {\n\t\t\t\t\"banner_image\": {\n\t\t\t\t\t\"uri\": \"/asset-v1:BUx+PY1x+3T2016+type@asset+block@images_course_image.jpg\",\n\t\t\t\t\t\"uri_absolute\": \"https://courses.edx.org/asset-v1:BUx+PY1x+3T2016+type@asset+block@images_course_image.jpg\"\n\t\t\t\t},\n\t\t\t\t\"course_image\": {\n\t\t\t\t\t\"uri\": \"/asset-v1:BUx+PY1x+3T2016+type@asset+block@AP_Physics_Dashboard_Image.png\"\n\t\t\t\t},\n\t\t\t\t\"course_video\": {\n\t\t\t\t\t\"uri\": null\n\t\t\t\t},\n\t\t\t\t\"image\": {\n\t\t\t\t\t\"raw\": \"https://courses.edx.org/asset-v1:BUx+PY1x+3T2016+type@asset+block@AP_Physics_Dashboard_Image.png\",\n\t\t\t\t\t\"small\": \"https://courses.edx.org/asset-v1:BUx+PY1x+3T2016+type@thumbnail+block@AP_Physics_Dashboard_Image-png-375x200.jpg\",\n\t\t\t\t\t\"large\": \"https://courses.edx.org/asset-v1:BUx+PY1x+3T2016+type@thumbnail+block@AP_Physics_Dashboard_Image-png-750x400.jpg\"\n\t\t\t\t}\n\t\t\t},\n\t\t\t\"name\": \"GIS AP Physics 1\",\n\t\t\t\"number\": \"PY1x\",\n\t\t\t\"org\": \"BUx\",\n\t\t\t\"short_description\": \"\",\n\t\t\t\"start\": \"2016-09-12T18:00:00Z\",\n\t\t\t\"start_display\": \"Sept. 12, 2016\",\n\t\t\t\"start_type\": \"timestamp\",\n\t\t\t\"pacing\": \"instructor\",\n\t\t\t\"mobile_available\": true,\n\t\t\t\"hidden\": false,\n\t\t\t\"invitation_only\": false,\n\t\t\t\"course_id\": \"ccx-v1:BUx+PY1x+3T2016+ccx@14\"\n\t\t},\n\t\t{\n\t\t\t\"blocks_url\": \"https://courses.edx.org/api/courses/v2/blocks/?course_id=ccx-v1%3ABUx%2BPY1x%2B3T2016%2Bccx%4017\",\n\t\t\t\"effort\": null,\n\t\t\t\"end\": null,\n\t\t\t\"enrollment_start\": \"2016-08-08T18:00:00Z\",\n\t\t\t\"enrollment_end\": null,\n\t\t\t\"id\": \"ccx-v1:BUx+PY1x+3T2016+ccx@17\",\n\t\t\t\"media\": {\n\t\t\t\t\"banner_image\": {\n\t\t\t\t\t\"uri\": \"/asset-v1:BUx+PY1x+3T2016+type@asset+block@images_course_image.jpg\",\n\t\t\t\t\t\"uri_absolute\": \"https://courses.edx.org/asset-v1:BUx+PY1x+3T2016+type@asset+block@images_course_image.jpg\"\n\t\t\t\t},\n\t\t\t\t\"course_image\": {\n\t\t\t\t\t\"uri\": \"/asset-v1:BUx+PY1x+3T2016+type@asset+block@AP_Physics_Dashboard_Image.png\"\n\t\t\t\t},\n\t\t\t\t\"course_video\": {\n\t\t\t\t\t\"uri\": null\n\t\t\t\t},\n\t\t\t\t\"image\": {\n\t\t\t\t\t\"raw\": \"https://courses.edx.org/asset-v1:BUx+PY1x+3T2016+type@asset+block@AP_Physics_Dashboard_Image.png\",\n\t\t\t\t\t\"small\": \"https://courses.edx.org/asset-v1:BUx+PY1x+3T2016+type@thumbnail+block@AP_Physics_Dashboard_Image-png-375x200.jpg\",\n\t\t\t\t\t\"large\": \"https://courses.edx.org/asset-v1:BUx+PY1x+3T2016+type@thumbnail+block@AP_Physics_Dashboard_Image-png-750x400.jpg\"\n\t\t\t\t}\n\t\t\t},\n\t\t\t\"name\": \"AP Physics 1\",\n\t\t\t\"number\": \"PY1x\",\n\t\t\t\"org\": \"BUx\",\n\t\t\t\"short_description\": \"\",\n\t\t\t\"start\": \"2016-09-12T18:00:00Z\",\n\t\t\t\"start_display\": \"Sept. 12, 2016\",\n\t\t\t\"start_type\": \"timestamp\",\n\t\t\t\"pacing\": \"instructor\",\n\t\t\t\"mobile_available\": true,\n\t\t\t\"hidden\": false,\n\t\t\t\"invitation_only\": false,\n\t\t\t\"course_id\": \"ccx-v1:BUx+PY1x+3T2016+ccx@17\"\n\t\t},\n\t\t{\n\t\t\t\"blocks_url\": \"https://courses.edx.org/api/courses/v2/blocks/?course_id=ccx-v1%3ABUx%2BPY1x%2B3T2016%2Bccx%4018\",\n\t\t\t\"effort\": null,\n\t\t\t\"end\": null,\n\t\t\t\"enrollment_start\": \"2016-08-08T18:00:00Z\",\n\t\t\t\"enrollment_end\": null,\n\t\t\t\"id\": \"ccx-v1:BUx+PY1x+3T2016+ccx@18\",\n\t\t\t\"media\": {\n\t\t\t\t\"banner_image\": {\n\t\t\t\t\t\"uri\": \"/asset-v1:BUx+PY1x+3T2016+type@asset+block@images_course_image.jpg\",\n\t\t\t\t\t\"uri_absolute\": \"https://courses.edx.org/asset-v1:BUx+PY1x+3T2016+type@asset+block@images_course_image.jpg\"\n\t\t\t\t},\n\t\t\t\t\"course_image\": {\n\t\t\t\t\t\"uri\": \"/asset-v1:BUx+PY1x+3T2016+type@asset+block@AP_Physics_Dashboard_Image.png\"\n\t\t\t\t},\n\t\t\t\t\"course_video\": {\n\t\t\t\t\t\"uri\": null\n\t\t\t\t},\n\t\t\t\t\"image\": {\n\t\t\t\t\t\"raw\": \"https://courses.edx.org/asset-v1:BUx+PY1x+3T2016+type@asset+block@AP_Physics_Dashboard_Image.png\",\n\t\t\t\t\t\"small\": \"https://courses.edx.org/asset-v1:BUx+PY1x+3T2016+type@thumbnail+block@AP_Physics_Dashboard_Image-png-375x200.jpg\",\n\t\t\t\t\t\"large\": \"https://courses.edx.org/asset-v1:BUx+PY1x+3T2016+type@thumbnail+block@AP_Physics_Dashboard_Image-png-750x400.jpg\"\n\t\t\t\t}\n\t\t\t},\n\t\t\t\"name\": \"Morckel AP1 Period 5\",\n\t\t\t\"number\": \"PY1x\",\n\t\t\t\"org\": \"BUx\",\n\t\t\t\"short_description\": \"\",\n\t\t\t\"start\": \"2016-09-12T18:00:00Z\",\n\t\t\t\"start_display\": \"Sept. 12, 2016\",\n\t\t\t\"start_type\": \"timestamp\",\n\t\t\t\"pacing\": \"instructor\",\n\t\t\t\"mobile_available\": true,\n\t\t\t\"hidden\": false,\n\t\t\t\"invitation_only\": false,\n\t\t\t\"course_id\": \"ccx-v1:BUx+PY1x+3T2016+ccx@18\"\n\t\t},\n\t\t{\n\t\t\t\"blocks_url\": \"https://courses.edx.org/api/courses/v2/blocks/?course_id=ccx-v1%3ABUx%2BPY1x%2B3T2016%2Bccx%4019\",\n\t\t\t\"effort\": null,\n\t\t\t\"end\": null,\n\t\t\t\"enrollment_start\": \"2016-08-08T18:00:00Z\",\n\t\t\t\"enrollment_end\": null,\n\t\t\t\"id\": \"ccx-v1:BUx+PY1x+3T2016+ccx@19\",\n\t\t\t\"media\": {\n\t\t\t\t\"banner_image\": {\n\t\t\t\t\t\"uri\": \"/asset-v1:BUx+PY1x+3T2016+type@asset+block@images_course_image.jpg\",\n\t\t\t\t\t\"uri_absolute\": \"https://courses.edx.org/asset-v1:BUx+PY1x+3T2016+type@asset+block@images_course_image.jpg\"\n\t\t\t\t},\n\t\t\t\t\"course_image\": {\n\t\t\t\t\t\"uri\": \"/asset-v1:BUx+PY1x+3T2016+type@asset+block@AP_Physics_Dashboard_Image.png\"\n\t\t\t\t},\n\t\t\t\t\"course_video\": {\n\t\t\t\t\t\"uri\": null\n\t\t\t\t},\n\t\t\t\t\"image\": {\n\t\t\t\t\t\"raw\": \"https://courses.edx.org/asset-v1:BUx+PY1x+3T2016+type@asset+block@AP_Physics_Dashboard_Image.png\",\n\t\t\t\t\t\"small\": \"https://courses.edx.org/asset-v1:BUx+PY1x+3T2016+type@thumbnail+block@AP_Physics_Dashboard_Image-png-375x200.jpg\",\n\t\t\t\t\t\"large\": \"https://courses.edx.org/asset-v1:BUx+PY1x+3T2016+type@thumbnail+block@AP_Physics_Dashboard_Image-png-750x400.jpg\"\n\t\t\t\t}\n\t\t\t},\n\t\t\t\"name\": \"Morckel AP1 Period 5\",\n\t\t\t\"number\": \"PY1x\",\n\t\t\t\"org\": \"BUx\",\n\t\t\t\"short_description\": \"\",\n\t\t\t\"start\": \"2016-09-12T18:00:00Z\",\n\t\t\t\"start_display\": \"Sept. 12, 2016\",\n\t\t\t\"start_type\": \"timestamp\",\n\t\t\t\"pacing\": \"instructor\",\n\t\t\t\"mobile_available\": true,\n\t\t\t\"hidden\": false,\n\t\t\t\"invitation_only\": false,\n\t\t\t\"course_id\": \"ccx-v1:BUx+PY1x+3T2016+ccx@19\"\n\t\t}\n\t],\n\t\"pagination\": {\n\t\t\"next\": \"https://courses.edx.org/api/courses/v1/courses/?page=2\",\n\t\t\"previous\": null,\n\t\t\"count\": 20505,\n\t\t\"num_pages\": 2051\n\t}\n}"), nil)
			},
			expectedStatusCode: 200,
		},
		{
			name:                 "Incorrect page number",
			inputParamForHandler: "1000000000000000",
			inputParamForMock:    1000000000000000,
			mockBehavior: func(s *mock_courses.MockDelegate, pageNumber int) {
				s.EXPECT().GetAllPublicCourses(pageNumber).Return(nil, edxApi.ErrIncorrectInputParam)
			},
			expectedStatusCode: 500,
		},
		{
			name:                 "Incorrect page number",
			inputParamForHandler: "0",
			inputParamForMock:    0,
			mockBehavior: func(s *mock_courses.MockDelegate, pageNumber int) {
				s.EXPECT().GetAllPublicCourses(pageNumber).Return(nil, edxApi.ErrIncorrectInputParam)
			},
			expectedStatusCode: 500,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			edx := mock_courses.NewMockDelegate(c)
			testCase.mockBehavior(edx, testCase.inputParamForMock)

			handler := NewCoursesHandler(edx)

			r := gin.New()
			r.GET("/getAllPublicCourses/:pageNumber", handler.GetAllPublicCourses, func(c *gin.Context) {
				c.Param(testCase.inputParamForHandler)
			})

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/getAllPublicCourses/"+testCase.inputParamForHandler, nil)
			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
		})
	}
}

func TestHandler_GetEnrollments(t *testing.T) {
	type mockBehavior func(s *mock_courses.MockDelegate, username string)

	testTable := []struct {
		name               string
		inputParam         string
		mockBehavior       mockBehavior
		expectedStatusCode int
	}{
		{
			name:       "OK",
			inputParam: "edxsom",
			mockBehavior: func(s *mock_courses.MockDelegate, username string) {
				s.EXPECT().GetEnrollments(username).Return([]byte("{\n\t\"next\": null,\n\t\"previous\": null,\n\t\"results\": [\n\t\t{\n\t\t\t\"created\": \"2022-06-13T03:00:12.571664Z\",\n\t\t\t\"mode\": \"honor\",\n\t\t\t\"is_active\": false,\n\t\t\t\"user\": \"edxsom\",\n\t\t\t\"course_id\": \"course-v1:TestOrg+02+2022\"\n\t\t},\n\t\t{\n\t\t\t\"created\": \"2022-06-13T01:16:45.374794Z\",\n\t\t\t\"mode\": \"honor\",\n\t\t\t\"is_active\": true,\n\t\t\t\"user\": \"edxsom\",\n\t\t\t\"course_id\": \"course-v1:Test_org+01+2022\"\n\t\t}\n\t]\n}"), nil)
			},
			expectedStatusCode: 200,
		},
		{
			name:       "Incorrect username",
			inputParam: "fsd",

			mockBehavior: func(s *mock_courses.MockDelegate, username string) {
				s.EXPECT().GetEnrollments(username).Return([]byte("Somebody once told me the world is gonna roll me"), nil)
			},
			expectedStatusCode: 500,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			edx := mock_courses.NewMockDelegate(c)
			testCase.mockBehavior(edx, testCase.inputParam)

			handler := NewCoursesHandler(edx)

			r := gin.New()
			r.GET("/getEnrollments/:username", handler.GetEnrollments, func(c *gin.Context) {
				c.Param(testCase.inputParam)
			})

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/getEnrollments/"+testCase.inputParam, nil)
			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
		})
	}
}

func TestHandler_PostEnrollment(t *testing.T) {
	type mockBehavior func(s *mock_courses.MockDelegate, enrollment models.PostEnrollmentHTTP)

	testTable := []struct {
		name               string
		inputBody          string
		inputModel         models.PostEnrollmentHTTP
		mockBehavior       mockBehavior
		expectedStatusCode int
	}{
		{
			name:       "OK",
			inputBody:  "{\n\t\"course_details\": {\n\t\t\"course_id\": \"course-v1:Test_org+01+2022\"\n\t},\n\t\"user\":\"tesr_user\"\n}",
			inputModel: models.PostEnrollmentHTTP{},
			mockBehavior: func(s *mock_courses.MockDelegate, enrollment models.PostEnrollmentHTTP) {
				s.EXPECT().PostEnrollment(&enrollment).Return(nil)
			},
			expectedStatusCode: 200,
		},
		{
			name:       "Incorrect input body",
			inputBody:  "{\n\t\"course_details\": {\n\t\t\"codsa\": \"course-v1:Test_org+01+2022\"\n\t},\n\t\"user\":\"tesr_user\"\n}",
			inputModel: models.PostEnrollmentHTTP{},
			mockBehavior: func(s *mock_courses.MockDelegate, enrollment models.PostEnrollmentHTTP) {
				s.EXPECT().PostEnrollment(&enrollment).Return(edxApi.ErrIncorrectInputParam)
			},
			expectedStatusCode: 500,
		},
		{
			name:               "Empty input body",
			inputBody:          "",
			inputModel:         models.PostEnrollmentHTTP{},
			mockBehavior:       nil,
			expectedStatusCode: 500,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			edx := mock_courses.NewMockDelegate(c)
			if testCase.mockBehavior != nil {
				testCase.mockBehavior(edx, testCase.inputModel)
			}

			handler := NewCoursesHandler(edx)

			r := gin.New()
			r.POST("/postEnrollment", handler.PostEnrollment)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/postEnrollment", bytes.NewBufferString(testCase.inputBody))
			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
		})
	}
}

func TestHandler_PostUnenroll(t *testing.T) {
	type mockBehavior func(s *mock_courses.MockDelegate, unenroll models.PostEnrollmentHTTP)

	testTable := []struct {
		name               string
		inputBody          string
		inputModel         models.PostEnrollmentHTTP
		mockBehavior       mockBehavior
		expectedStatusCode int
	}{
		{
			name:       "OK",
			inputBody:  "{\n\"is_active\" : false,\n\t\"course_details\": {\n\t\t\"course_id\": \"course-v1:Test_org+01+2022\"\n\t},\n\t\"user\":\"tesr_user\"\n}",
			inputModel: models.PostEnrollmentHTTP{},
			mockBehavior: func(s *mock_courses.MockDelegate, unenroll models.PostEnrollmentHTTP) {
				s.EXPECT().PostUnenroll(&unenroll).Return(nil)
			},
			expectedStatusCode: 200,
		},
		{
			name:       "Incorrect input body",
			inputBody:  "{\n\t\"corwerwels\": {\n\t\t\"codsa\": \"course-v1:Test_org+01+2022\"\n\t},\n\t\"user\":\"tesr_user\"\n}",
			inputModel: models.PostEnrollmentHTTP{},
			mockBehavior: func(s *mock_courses.MockDelegate, unenroll models.PostEnrollmentHTTP) {
				s.EXPECT().PostUnenroll(&unenroll).Return(edxApi.ErrIncorrectInputParam)
			},
			expectedStatusCode: 500,
		},
		{
			name:               "Empty input body",
			inputBody:          "",
			inputModel:         models.PostEnrollmentHTTP{},
			mockBehavior:       nil,
			expectedStatusCode: 500,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			edx := mock_courses.NewMockDelegate(c)
			if testCase.mockBehavior != nil {
				testCase.mockBehavior(edx, testCase.inputModel)
			}

			handler := NewCoursesHandler(edx)

			r := gin.New()
			r.POST("/postUnenroll", handler.PostUnenroll)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/postUnenroll", bytes.NewBufferString(testCase.inputBody))
			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
		})
	}
}

func TestHandler_Registration(t *testing.T) {
	type mockBehavior func(s *mock_courses.MockDelegate, registrationMsg edxApi.RegistrationForm)

	testTable := []struct {
		name               string
		inputBody          string
		inputModel         edxApi.RegistrationForm
		mockBehavior       mockBehavior
		expectedStatusCode int
	}{
		{
			name:      "OK",
			inputBody: "{\n\t\"email\" : \"sasdasddasdaa1@mail.ru\",\n\t\"password\" : \"f2r4d3rdat234\",\n\t\"username\" : \"testtdda1ddasdasatt\",\n\t\"terms_of_service\" : \"true\",\n\t\"name\" : \"tttttttt\"\n}",
			inputModel: edxApi.RegistrationForm{
				Email:          "sasdasddasdaa1@mail.ru",
				Password:       "f2r4d3rdat234",
				Username:       "testtdda1ddasdasatt",
				TermsOfService: "true",
				Name:           "tttttttt",
			},
			mockBehavior: func(s *mock_courses.MockDelegate, registrationMsg edxApi.RegistrationForm) {
				s.EXPECT().Registration(&registrationMsg).Return(nil)
			},
			expectedStatusCode: 200,
		},
		{
			name:       "Incorrect input body",
			inputBody:  "{\n\t\"corwerwels\": {\n\t\t\"codsa\": \"course-v1:Test_org+01+2022\"\n\t},\n\t\"user\":\"tesr_user\"\n}",
			inputModel: edxApi.RegistrationForm{},
			mockBehavior: func(s *mock_courses.MockDelegate, registrationMsg edxApi.RegistrationForm) {
				s.EXPECT().Registration(&registrationMsg).Return(edxApi.ErrIncorrectInputParam)
			},
			expectedStatusCode: 500,
		},
		{
			name:               "Empty input body",
			inputBody:          "",
			inputModel:         edxApi.RegistrationForm{},
			mockBehavior:       nil,
			expectedStatusCode: 500,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			edx := mock_courses.NewMockDelegate(c)
			if testCase.mockBehavior != nil {
				testCase.mockBehavior(edx, testCase.inputModel)
			}

			handler := NewCoursesHandler(edx)

			r := gin.New()
			r.POST("/registration", handler.Registration)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/registration", bytes.NewBufferString(testCase.inputBody))
			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
		})
	}
}

func TestHandler_Login(t *testing.T) {
	type mockBehavior func(s *mock_courses.MockDelegate, loginMsg LoginUser)

	testTable := []struct {
		name               string
		inputBody          string
		inputModel         LoginUser
		mockBehavior       mockBehavior
		expectedStatusCode int
	}{
		{
			name:      "OK",
			inputBody: "{\"email\":\"danisdada\",\"password\":\"f2r43rt234\"}",
			inputModel: LoginUser{
				Email:    "danisdada",
				Password: "f2r43rt234",
			},
			mockBehavior: func(s *mock_courses.MockDelegate, loginMsg LoginUser) {
				s.EXPECT().Login(loginMsg.Email, loginMsg.Password).Return(nil)
			},
			expectedStatusCode: 200,
		},
		{
			name:      "Incorrect input body",
			inputBody: "{\n\t\"email\" : \"danisdadatestru@test.ru\"}",
			inputModel: LoginUser{
				Email: "danisdadatestru@test.ru",
			},
			mockBehavior: func(s *mock_courses.MockDelegate, loginMsg LoginUser) {
				s.EXPECT().Login(loginMsg.Email, loginMsg.Password).Return(edxApi.ErrIncorrectInputParam)
			},
			expectedStatusCode: 500,
		},
		{
			name:               "Empty input body",
			inputBody:          "",
			inputModel:         LoginUser{},
			mockBehavior:       nil,
			expectedStatusCode: 500,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			edx := mock_courses.NewMockDelegate(c)
			if testCase.mockBehavior != nil {
				testCase.mockBehavior(edx, testCase.inputModel)
			}

			handler := NewCoursesHandler(edx)

			r := gin.New()
			r.POST("/login", handler.Login)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/login", bytes.NewBufferString(testCase.inputBody))
			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
		})
	}
}

func TestHandler_DeleteCourse(t *testing.T) {
	type mockBehavior func(s *mock_courses.MockDelegate, courseId string)

	testTable := []struct {
		name               string
		inputParam         string
		mockBehavior       mockBehavior
		expectedStatusCode int
	}{
		{
			name:       "OK",
			inputParam: "course-v1:Test_org+01+2022",
			mockBehavior: func(s *mock_courses.MockDelegate, courseId string) {
				s.EXPECT().DeleteCourse(courseId).Return(nil)
			},
			expectedStatusCode: 200,
		},
		{
			name:       "Incorrect course id",
			inputParam: "fsd",
			mockBehavior: func(s *mock_courses.MockDelegate, courseId string) {
				s.EXPECT().DeleteCourse(courseId).Return(edxApi.ErrIncorrectInputParam)
			},
			expectedStatusCode: 500,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			edx := mock_courses.NewMockDelegate(c)
			testCase.mockBehavior(edx, testCase.inputParam)

			handler := NewCoursesHandler(edx)

			r := gin.New()
			r.DELETE("/deleteCourse/:courseId", handler.DeleteCourse, func(c *gin.Context) {
				c.Param(testCase.inputParam)
			})

			w := httptest.NewRecorder()
			req := httptest.NewRequest("DELETE", "/deleteCourse/"+testCase.inputParam, nil)
			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
		})
	}
}

func TestHandler_UpdateCourse(t *testing.T) {
	type mockBehavior func(s *mock_courses.MockDelegate, course models.CourseHTTP)

	testTable := []struct {
		name               string
		inputBody          string
		inputCourse        models.CourseHTTP
		mockBehavior       mockBehavior
		expectedStatusCode int
	}{
		{
			name:      "OK",
			inputBody: "{\n\n      \"id\": \"1\",\n      \"blocks_url\": \"1970\",\n\t\t\t\"effort\": \"somestring\"}",
			inputCourse: models.CourseHTTP{
				ID:        "1",
				BlocksUrl: "1970",
				Effort:    "somestring",
			},
			mockBehavior: func(s *mock_courses.MockDelegate, course models.CourseHTTP) {
				s.EXPECT().UpdateCourse(&course).Return(nil)
			},
			expectedStatusCode: 200,
		},
		{
			name:      "Incorrect body",
			inputBody: "{\n\n      \"id\": \"somedas\",\n      \"blocks_url\": \"1968\",\n\t\t\t\"effort\": \"somestring\"}",
			inputCourse: models.CourseHTTP{
				ID:        "somedas",
				BlocksUrl: "1968",
				Effort:    "somestring",
			},
			mockBehavior: func(s *mock_courses.MockDelegate, course models.CourseHTTP) {
				s.EXPECT().UpdateCourse(&course).Return(edxApi.ErrIncorrectInputParam)
			},
			expectedStatusCode: 500,
		},
		{
			name:               "Empty Body",
			inputBody:          "",
			inputCourse:        models.CourseHTTP{},
			mockBehavior:       nil,
			expectedStatusCode: 400,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			edx := mock_courses.NewMockDelegate(c)
			if testCase.mockBehavior != nil {
				testCase.mockBehavior(edx, testCase.inputCourse)
			}

			handler := NewCoursesHandler(edx)

			r := gin.New()
			r.PUT("/updateCourse", handler.UpdateCourse)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("PUT", "/updateCourse", bytes.NewBufferString(testCase.inputBody))
			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
		})
	}
}
