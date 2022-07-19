package usecase

import (
	"errors"
	assert "github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/edxApi"
	mock_edxApi "github.com/skinnykaen/robbo_student_personal_account.git/package/edxApi/mocks"
	"testing"
)

func TestGetUser(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockapi := mock_edxApi.NewMockEdxApiUseCase(ctrl)
		expect := []byte("{\n\t\"username\": \"edxsom\"\n}")
		mockapi.EXPECT().GetUser().Return([]byte("{\n\t\"username\": \"edxsom\"\n}"), nil)
		testApi := &EdxApiUseCaseModule{EdxApiUseCase: mockapi}
		correct, _ := testApi.GetUser()
		assert.Equal(t, expect, correct)
	})
}

func TestEdxApiUseCaseImpl_GetCourseContent(t *testing.T) {
	type mockBehavior func(s *mock_edxApi.MockEdxApiUseCase, courseId string)
	testTable := []struct {
		name         string
		courseId     string
		expectedBody []byte
		mockBehavior
	}{
		{
			name:         "Ok",
			courseId:     "course-v1:Test_org+01+2022",
			expectedBody: []byte("{\"blocks_url\":\"https://edx-test.ru/api/courses/v2/blocks/?course_id=course-v1%3ATest_org%2B01%2B2022\",\"effort\":null,\"end\":\"2023-02-02T00:00:00Z\",\"enrollment_start\":null,\"enrollment_end\n\":null,\"id\":\"course-v1:Test_org+01+2022\",\"media\":{\"banner_image\":{\"uri\":\"/asset-v1:Test_org+01+2022+type@asset+block@images_course_image.jpg\",\"uri_absolute\":\"https://edx-test.ru/asset-\nv1:Test_org+01+2022+type@asset+block@images_course_image.jpg\"},\"course_image\":{\"uri\":\"/asset-v1:Test_org+01+2022+type@asset+block@images_course_image.jpg\"},\"course_video\":{\"uri\":null},\n\"image\":{\"raw\":\"https://edx-test.ru/asset-v1:Test_org+01+2022+type@asset+block@images_course_image.jpg\",\"small\":\"https://edx-test.ru/asset-v1:Test_org+01+2022+type@asset+block@images_c\nourse_image.jpg\",\"large\":\"https://edx-test.ru/asset-v1:Test_org+01+2022+type@asset+block@images_course_image.jpg\"}},\"name\":\"Cert test\",\"number\":\"01\",\"org\":\"Test_org\",\"short_description\n\":\"\",\"start\":\"2022-01-01T00:00:00Z\",\"start_display\":\"Jan. 1, 2022\",\"start_type\":\"timestamp\",\"pacing\":\"instructor\",\"mobile_available\":false,\"hidden\":false,\"invitation_only\":false,\"cours\ne_id\":\"course-v1:Test_org+01+2022\",\"overview\":\"<section class=\\\"about\\\">\\n  <h2>About This Course</h2>\\n  <p>Include your long course description here. The long course description shou\nld contain 150-400 words.</p>\\n\\n  <p>This is paragraph 2 of the long course description. Add more paragraphs as needed. Make sure to enclose them in paragraph tags.</p>\\n</section>\\n\\\nn<section class=\\\"prerequisites\\\">\\n  <h2>Requirements</h2>\\n  <p>Add information about the skills and knowledge students need to take this course.</p>\\n</section>\\n\\n<section class=\\\"\ncourse-staff\\\">\\n  <h2>Course Staff</h2>\\n  <article class=\\\"teacher\\\">\\n    <div class=\\\"teacher-image\\\">\\n      <img src=\\\"/static/images/placeholder-faculty.png\\\" align=\\\"left\\\" sty\nle=\\\"margin:0 20 px 0\\\" alt=\\\"Course Staff Image #1\\\">\\n    </div>\\n\\n    <h3>Staff Member #1</h3>\\n    <p>Biography of instructor/staff member #1</p>\\n  </article>\\n\\n  <article class\n=\\\"teacher\\\">\\n    <div class=\\\"teacher-image\\\">\\n      <img src=\\\"/static/images/placeholder-faculty.png\\\" align=\\\"left\\\" style=\\\"margin:0 20 px 0\\\" alt=\\\"Course Staff Image #2\\\">\\n  \n  </div>\\n\\n    <h3>Staff Member #2</h3>\\n    <p>Biography of instructor/staff member #2</p>\\n  </article>\\n</section>\\n\\n<section class=\\\"faq\\\">\\n  <section class=\\\"responses\\\">\\n    \n<h2>Frequently Asked Questions</h2>\\n    <article class=\\\"response\\\">\\n      <h3>What web browser should I use?</h3>\\n      <p>The Open edX platform works best with current versions of\n Chrome, Edge, Firefox, Internet Explorer, or Safari.</p>\\n      <p>See our <a href=\\\"https://edx.readthedocs.org/projects/open-edx-learner-guide/en/latest/front_matter/browsers.html\\\"\n>list of supported browsers</a> for the most up-to-date information.</p>\\n    </article>\\n\\n    <article class=\\\"response\\\">\\n      <h3>Question #2</h3>\\n      <p>Your answer would be displayed here.</p>\\n    </article>\\n  </section>\\n</section>\\n\"}\n"),
			mockBehavior: func(s *mock_edxApi.MockEdxApiUseCase, courseId string) {
				s.EXPECT().GetCourseContent(courseId).Return([]byte("{\"blocks_url\":\"https://edx-test.ru/api/courses/v2/blocks/?course_id=course-v1%3ATest_org%2B01%2B2022\",\"effort\":null,\"end\":\"2023-02-02T00:00:00Z\",\"enrollment_start\":null,\"enrollment_end\n\":null,\"id\":\"course-v1:Test_org+01+2022\",\"media\":{\"banner_image\":{\"uri\":\"/asset-v1:Test_org+01+2022+type@asset+block@images_course_image.jpg\",\"uri_absolute\":\"https://edx-test.ru/asset-\nv1:Test_org+01+2022+type@asset+block@images_course_image.jpg\"},\"course_image\":{\"uri\":\"/asset-v1:Test_org+01+2022+type@asset+block@images_course_image.jpg\"},\"course_video\":{\"uri\":null},\n\"image\":{\"raw\":\"https://edx-test.ru/asset-v1:Test_org+01+2022+type@asset+block@images_course_image.jpg\",\"small\":\"https://edx-test.ru/asset-v1:Test_org+01+2022+type@asset+block@images_c\nourse_image.jpg\",\"large\":\"https://edx-test.ru/asset-v1:Test_org+01+2022+type@asset+block@images_course_image.jpg\"}},\"name\":\"Cert test\",\"number\":\"01\",\"org\":\"Test_org\",\"short_description\n\":\"\",\"start\":\"2022-01-01T00:00:00Z\",\"start_display\":\"Jan. 1, 2022\",\"start_type\":\"timestamp\",\"pacing\":\"instructor\",\"mobile_available\":false,\"hidden\":false,\"invitation_only\":false,\"cours\ne_id\":\"course-v1:Test_org+01+2022\",\"overview\":\"<section class=\\\"about\\\">\\n  <h2>About This Course</h2>\\n  <p>Include your long course description here. The long course description shou\nld contain 150-400 words.</p>\\n\\n  <p>This is paragraph 2 of the long course description. Add more paragraphs as needed. Make sure to enclose them in paragraph tags.</p>\\n</section>\\n\\\nn<section class=\\\"prerequisites\\\">\\n  <h2>Requirements</h2>\\n  <p>Add information about the skills and knowledge students need to take this course.</p>\\n</section>\\n\\n<section class=\\\"\ncourse-staff\\\">\\n  <h2>Course Staff</h2>\\n  <article class=\\\"teacher\\\">\\n    <div class=\\\"teacher-image\\\">\\n      <img src=\\\"/static/images/placeholder-faculty.png\\\" align=\\\"left\\\" sty\nle=\\\"margin:0 20 px 0\\\" alt=\\\"Course Staff Image #1\\\">\\n    </div>\\n\\n    <h3>Staff Member #1</h3>\\n    <p>Biography of instructor/staff member #1</p>\\n  </article>\\n\\n  <article class\n=\\\"teacher\\\">\\n    <div class=\\\"teacher-image\\\">\\n      <img src=\\\"/static/images/placeholder-faculty.png\\\" align=\\\"left\\\" style=\\\"margin:0 20 px 0\\\" alt=\\\"Course Staff Image #2\\\">\\n  \n  </div>\\n\\n    <h3>Staff Member #2</h3>\\n    <p>Biography of instructor/staff member #2</p>\\n  </article>\\n</section>\\n\\n<section class=\\\"faq\\\">\\n  <section class=\\\"responses\\\">\\n    \n<h2>Frequently Asked Questions</h2>\\n    <article class=\\\"response\\\">\\n      <h3>What web browser should I use?</h3>\\n      <p>The Open edX platform works best with current versions of\n Chrome, Edge, Firefox, Internet Explorer, or Safari.</p>\\n      <p>See our <a href=\\\"https://edx.readthedocs.org/projects/open-edx-learner-guide/en/latest/front_matter/browsers.html\\\"\n>list of supported browsers</a> for the most up-to-date information.</p>\\n    </article>\\n\\n    <article class=\\\"response\\\">\\n      <h3>Question #2</h3>\\n      <p>Your answer would be displayed here.</p>\\n    </article>\\n  </section>\\n</section>\\n\"}\n"), nil)
			},
		},

		{
			name:         "Bad courseId",
			courseId:     "Ddssadad",
			expectedBody: []byte("\n\n\n\n\n\n\n\n\n<!DOCTYPE html>\n<!--[if lte IE 9]><html class=\"ie ie9 lte9\" lang=\"en\"><![endif]-->\n<!--[if !IE]><!--><html lang=\"en\"><!--<![endif]-->\n<head dir=\"ltr\">\n    <meta charset=\"UTF-8\">\n    <meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n    <meta http-equiv=\"origin-trial\" content=\"ArNBN7d1AkvMhJTGWXlJ8td/AN4lOokzOnqKRNkTnLqaqx0HpfYvmx8JePPs/emKh6O5fckx14LeZIGJ1AQYjgAAAABzeyJvcmlnaW4iOiJodHRwOi8vbG9jYWxob3N0OjE4MDAwIiwiZmVhdHVyZSI6IkRpc2FibGVEaWZmZXJlbnRPcmlnaW5TdWJmcmFtZURpYWxvZ1N1cHByZXNzaW9uIiwiZXhwaXJ5IjoxNjM5NTI2Mzk5fQ==\">\n\n\n\n\n      <title>\n       Page Not Found | Robbo OpenEdx Test\n      </title>\n\n\n      <script type=\"text/javascript\">\n        /* immediately break out of an iframe if coming from the marketing website */\n        (function(window) {\n          if (window.location !== window.top.location) {\n            window.top.location = window.location;\n          }\n        })(this);\n      </script>\n\n\n\n  <script type=\"text/javascript\" src=\"/static/js/i18n/en/djangojs.2623d59dd64d.js\"></script>\n  <script type=\"text/javascript\" src=\"/static/js/ie11_find_array.bd1c6dc7a133.js\"></script>\n\n  <link rel=\"icon\" type=\"image/x-icon\" href=\"/static/images/favicon.03ffbbf95a0d.ico\"/>\n\n\n\n\n\n\n    <link href=\"/static/css/lms-style-vendor.68e48093f5dd.css\" rel=\"stylesheet\" type=\"text/css\" />\n\n\n\n\n\n\n    <link href=\"/static/css/lms-main-v1.5f06a95d4cfd.css\" rel=\"stylesheet\" type=\"text/css\" />\n\n\n\n\n\n\n<script type=\"text/javascript\" src=\"/static/js/lms-main_vendor.3c3c9a2604d6.js\" charset=\"utf-8\"></script>\n\n\n\n\n<script type=\"text/javascript\" src=\"/static/js/lms-application.0e3fae472a20.js\" charset=\"utf-8\"></script>\n\n\n\n\n\n\n    <script type=\"text/javascript\" src=\"/static/bundles/commons.8f14f343047f03c88fc2.d5486becec92.js\" ></script>\n\n\n\n  <script>\n    window.baseUrl = \"/static/\";\n    (function (require) {\n      require.config({\n          baseUrl: window.baseUrl\n      });\n    }).call(this, require || RequireJS.require);\n  </script>\n  <script type=\"text/javascript\" src=\"/static/lms/js/require-config.38226099c6ad.js\"></script>\n\n    <script type=\"text/javascript\">\n        (function (require) {\n          require.config({\n              paths: {\n                'course_bookmarks/js/views/bookmark_button': 'course_bookmarks/js/views/bookmark_button.d4cfaf3361fa',\n'js/views/message_banner': 'js/views/message_banner.141974fd4f5d',\n'moment': 'common/js/vendor/moment-with-locales.43ec67e44612',\n'moment-timezone': 'common/js/vendor/moment-timezone-with-data.60142e6c4416',\n'js/courseware/course_info_events': 'js/courseware/course_info_events.2fc35b57627f',\n'js/courseware/accordion_events': 'js/courseware/accordion_events.6064c7809de5',\n'js/dateutil_factory': 'js/dateutil_factory.be68acdff619',\n'js/courseware/link_clicked_events': 'js/courseware/link_clicked_events',\n'js/courseware/toggle_element_visibility': 'js/courseware/toggle_element_visibility.474ff5ba9de3',\n'js/student_account/logistration_factory': 'js/student_account/logistration_factory.983820206227',\n'js/courseware/courseware_factory': 'js/courseware/courseware_factory.1504fc10caef',\n'js/groups/views/cohorts_dashboard_factory': 'js/groups/views/cohorts_dashboard_factory.ca68388d81d6',\n'js/groups/discussions_management/discussions_dashboard_factory': 'js/discussions_management/views/discussions_dashboard_factory.2e10d9097343',\n'draggabilly': 'js/vendor/draggabilly.26caba6f7187',\n'hls': 'common/js/vendor/hls.5e0c7e1b3bfd'\n            }\n          });\n        }).call(this, require || RequireJS.require);\n    </script>\n\n\n\n\n\n\n\n\n\n\n\n<script type=\"application/json\" id=\"user-metadata\">\n    null\n</script>\n\n\n\n\n\n\n\n\n\n<!-- dummy Segment -->\n<script type=\"text/javascript\">\n  var analytics = {\n    track: function() { return; },\n    trackLink: function() { return; },\n    pageview: function() { return; },\n    page: function() { return; }\n  };\n</script>\n<!-- end dummy Segment -->\n\n\n  <meta name=\"path_prefix\" content=\"\">\n\n\n  <meta name=\"openedx-release-line\" content=\"lilac\" />\n\n\n\n\n\n</head>\n\n<body class=\"ltr  lang_en\">\n\n\n<div id=\"page-prompt\"></div>\n  <div class=\"window-wrap\" dir=\"ltr\">\n\n    <a class=\"nav-skip sr-only sr-only-focusable\" href=\"#main\">Skip to main content</a>\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n<header class=\"global-header \">\n    <div class=\"main-header\">\n\n\n\n\n\n\n\n\n<h1 class=\"header-logo\">\n    <a href=\"/dashboard\">\n\n        <img  class=\"logo\" src=\"/static/images/logo.b6c374d66d57.png\" alt=\"Robbo OpenEdx Test  Home Page\"/>\n\n    </a>\n</h1>\n\n        <div class=\"hamburger-menu\" role=\"button\" aria-label=Options Menu aria-expanded=\"false\" aria-controls=\"mobile-menu\" tabindex=\"0\">\n            <span class=\"line\"></span>\n            <span class=\"line\"></span>\n            <span class=\"line\"></span>\n            <span class=\"line\"></span>\n        </div>\n\n\n\n\n\n\n\n\n\n<nav class=\"nav-links\" aria-label=Supplemental Links>\n  <div class=\"main\">\n      <div class=\"mobile-nav-item hidden-mobile nav-item\">\n        <a href=\"/courses\">Explore courses</a>\n      </div>\n  </div>\n  <div class=\"secondary\">\n    <div>\n            <div class=\"mobile-nav-item hidden-mobile nav-item\">\n                <a class=\"register-btn btn\" href=\"/register?next=%2Fapi%2Fcourses%2Fv1%2Fcourses%2Fdsada\">Register</a>\n            </div>\n          <div class=\"mobile-nav-item hidden-mobile nav-item\">\n              <a class=\"sign-in-btn btn\" href=\"/login?next=%2Fapi%2Fcourses%2Fv1%2Fcourses%2Fdsada\">Sign in</a>\n          </div>\n    </div>\n  </div>\n</nav>\n\n    </div>\n    <div class=\"mobile-menu hidden\" aria-label=More Options role=\"menu\" id=\"mobile-menu\"></div>\n</header>\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n    <div class=\"marketing-hero\"></div>\n\n    <div class=\"content-wrapper main-container\" id=\"content\" dir=\"ltr\">\n\n\n\n\n\n\n\n<main id=\"main\" aria-label=\"Content\" tabindex=\"-1\">\n    <section class=\"outside-app\">\n        <h1>\n            Page not found\n        </h1>\n        <p>\n\n                    The page that you were looking for was not found. Go back to the <a href=\"/\">homepage</a> .\n\n        </p>\n    </section>\n</main>\n\n\n    </div>\n\n\n\n\n\n\n\n  <div class=\"wrapper wrapper-footer\">\n    <footer id=\"footer-openedx\" class=\"grid-container\"\n    >\n      <div class=\"colophon\">\n        <nav class=\"nav-colophon\" aria-label=\"About\">\n          <ol>\n              <li class=\"nav-colophon-01\">\n                <a id=\"about\" href=\"/about\">About</a>\n              </li>\n              <li class=\"nav-colophon-02\">\n                <a id=\"blog\" href=\"/blog\">Blog</a>\n              </li>\n              <li class=\"nav-colophon-03\">\n                <a id=\"contact\" href=\"/support/contact_us\">Contact</a>\n              </li>\n              <li class=\"nav-colophon-04\">\n                <a id=\"donate\" href=\"/donate\">Donate</a>\n              </li>\n          </ol>\n        </nav>\n\n\n        <div class=\"wrapper-logo\">\n          <p>\n            <a href=\"/\">\n              <img alt=\"organization logo\" src=\"https://edx-test.ru/static/images/logo.b6c374d66d57.png\">\n            </a>\n          </p>\n        </div>\n\n          <p class=\"copyright\">© Robbo OpenEdx Test .  All rights reserved except where noted.  edX, Open edX and their respective logos are registered trademarks of edX Inc.\n          </p>\n\n        <nav class=\"nav-legal\" aria-label=\"Legal\">\n          <ul>\n              <li class=\"nav-legal-01\">\n                <a href=\"/tos_and_honor\">Terms of Service &amp; Honor Code</a>\n              </li>\n              <li class=\"nav-legal-02\">\n                <a href=\"/privacy\">Privacy Policy</a>\n              </li>\n            <li><a href=\"https://www.edx.org/?utm_medium=affiliate_partner&amp;utm_source=opensource-partner&amp;utm_content=open-edx-partner-footer-link&amp;utm_campaign=open-edx-footer\">Take free online courses at edX.org</a></li>\n          </ul>\n        </nav>\n      </div>\n\n      <div class=\"footer-about-openedx\">\n        <p>\n          <a href=\"https://open.edx.org\">\n            <img src=\"https://files.edx.org/openedx-logos/open-edx-logo-tag.png\" alt=\"Powered by Open edX\" width=\"175\" />\n          </a>\n        </p>\n      </div>\n    </footer>\n  </div>\n\n\n  </div>\n\n\n\n\n\n\n\n\n  <script type=\"text/javascript\" src=\"/static/js/vendor/noreferrer.aa62a3e70ffa.js\" charset=\"utf-8\"></script>\n  <script type=\"text/javascript\" src=\"/static/js/utils/navigation.08930e16ab3d.js\" charset=\"utf-8\"></script>\n  <script type=\"text/javascript\" src=\"/static/js/header/header.309a1243e175.js\"></script>\n\n  <script type=\"text/javascript\" src=\"/static/js/src/jquery_extend_patch.54dddef28d15.js\"></script>\n</body>\n</html>\n"),
			mockBehavior: func(s *mock_edxApi.MockEdxApiUseCase, courseId string) {
				s.EXPECT().GetCourseContent(courseId).Return([]byte("\n\n\n\n\n\n\n\n\n<!DOCTYPE html>\n<!--[if lte IE 9]><html class=\"ie ie9 lte9\" lang=\"en\"><![endif]-->\n<!--[if !IE]><!--><html lang=\"en\"><!--<![endif]-->\n<head dir=\"ltr\">\n    <meta charset=\"UTF-8\">\n    <meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n    <meta http-equiv=\"origin-trial\" content=\"ArNBN7d1AkvMhJTGWXlJ8td/AN4lOokzOnqKRNkTnLqaqx0HpfYvmx8JePPs/emKh6O5fckx14LeZIGJ1AQYjgAAAABzeyJvcmlnaW4iOiJodHRwOi8vbG9jYWxob3N0OjE4MDAwIiwiZmVhdHVyZSI6IkRpc2FibGVEaWZmZXJlbnRPcmlnaW5TdWJmcmFtZURpYWxvZ1N1cHByZXNzaW9uIiwiZXhwaXJ5IjoxNjM5NTI2Mzk5fQ==\">\n\n\n\n\n      <title>\n       Page Not Found | Robbo OpenEdx Test\n      </title>\n\n\n      <script type=\"text/javascript\">\n        /* immediately break out of an iframe if coming from the marketing website */\n        (function(window) {\n          if (window.location !== window.top.location) {\n            window.top.location = window.location;\n          }\n        })(this);\n      </script>\n\n\n\n  <script type=\"text/javascript\" src=\"/static/js/i18n/en/djangojs.2623d59dd64d.js\"></script>\n  <script type=\"text/javascript\" src=\"/static/js/ie11_find_array.bd1c6dc7a133.js\"></script>\n\n  <link rel=\"icon\" type=\"image/x-icon\" href=\"/static/images/favicon.03ffbbf95a0d.ico\"/>\n\n\n\n\n\n\n    <link href=\"/static/css/lms-style-vendor.68e48093f5dd.css\" rel=\"stylesheet\" type=\"text/css\" />\n\n\n\n\n\n\n    <link href=\"/static/css/lms-main-v1.5f06a95d4cfd.css\" rel=\"stylesheet\" type=\"text/css\" />\n\n\n\n\n\n\n<script type=\"text/javascript\" src=\"/static/js/lms-main_vendor.3c3c9a2604d6.js\" charset=\"utf-8\"></script>\n\n\n\n\n<script type=\"text/javascript\" src=\"/static/js/lms-application.0e3fae472a20.js\" charset=\"utf-8\"></script>\n\n\n\n\n\n\n    <script type=\"text/javascript\" src=\"/static/bundles/commons.8f14f343047f03c88fc2.d5486becec92.js\" ></script>\n\n\n\n  <script>\n    window.baseUrl = \"/static/\";\n    (function (require) {\n      require.config({\n          baseUrl: window.baseUrl\n      });\n    }).call(this, require || RequireJS.require);\n  </script>\n  <script type=\"text/javascript\" src=\"/static/lms/js/require-config.38226099c6ad.js\"></script>\n\n    <script type=\"text/javascript\">\n        (function (require) {\n          require.config({\n              paths: {\n                'course_bookmarks/js/views/bookmark_button': 'course_bookmarks/js/views/bookmark_button.d4cfaf3361fa',\n'js/views/message_banner': 'js/views/message_banner.141974fd4f5d',\n'moment': 'common/js/vendor/moment-with-locales.43ec67e44612',\n'moment-timezone': 'common/js/vendor/moment-timezone-with-data.60142e6c4416',\n'js/courseware/course_info_events': 'js/courseware/course_info_events.2fc35b57627f',\n'js/courseware/accordion_events': 'js/courseware/accordion_events.6064c7809de5',\n'js/dateutil_factory': 'js/dateutil_factory.be68acdff619',\n'js/courseware/link_clicked_events': 'js/courseware/link_clicked_events',\n'js/courseware/toggle_element_visibility': 'js/courseware/toggle_element_visibility.474ff5ba9de3',\n'js/student_account/logistration_factory': 'js/student_account/logistration_factory.983820206227',\n'js/courseware/courseware_factory': 'js/courseware/courseware_factory.1504fc10caef',\n'js/groups/views/cohorts_dashboard_factory': 'js/groups/views/cohorts_dashboard_factory.ca68388d81d6',\n'js/groups/discussions_management/discussions_dashboard_factory': 'js/discussions_management/views/discussions_dashboard_factory.2e10d9097343',\n'draggabilly': 'js/vendor/draggabilly.26caba6f7187',\n'hls': 'common/js/vendor/hls.5e0c7e1b3bfd'\n            }\n          });\n        }).call(this, require || RequireJS.require);\n    </script>\n\n\n\n\n\n\n\n\n\n\n\n<script type=\"application/json\" id=\"user-metadata\">\n    null\n</script>\n\n\n\n\n\n\n\n\n\n<!-- dummy Segment -->\n<script type=\"text/javascript\">\n  var analytics = {\n    track: function() { return; },\n    trackLink: function() { return; },\n    pageview: function() { return; },\n    page: function() { return; }\n  };\n</script>\n<!-- end dummy Segment -->\n\n\n  <meta name=\"path_prefix\" content=\"\">\n\n\n  <meta name=\"openedx-release-line\" content=\"lilac\" />\n\n\n\n\n\n</head>\n\n<body class=\"ltr  lang_en\">\n\n\n<div id=\"page-prompt\"></div>\n  <div class=\"window-wrap\" dir=\"ltr\">\n\n    <a class=\"nav-skip sr-only sr-only-focusable\" href=\"#main\">Skip to main content</a>\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n<header class=\"global-header \">\n    <div class=\"main-header\">\n\n\n\n\n\n\n\n\n<h1 class=\"header-logo\">\n    <a href=\"/dashboard\">\n\n        <img  class=\"logo\" src=\"/static/images/logo.b6c374d66d57.png\" alt=\"Robbo OpenEdx Test  Home Page\"/>\n\n    </a>\n</h1>\n\n        <div class=\"hamburger-menu\" role=\"button\" aria-label=Options Menu aria-expanded=\"false\" aria-controls=\"mobile-menu\" tabindex=\"0\">\n            <span class=\"line\"></span>\n            <span class=\"line\"></span>\n            <span class=\"line\"></span>\n            <span class=\"line\"></span>\n        </div>\n\n\n\n\n\n\n\n\n\n<nav class=\"nav-links\" aria-label=Supplemental Links>\n  <div class=\"main\">\n      <div class=\"mobile-nav-item hidden-mobile nav-item\">\n        <a href=\"/courses\">Explore courses</a>\n      </div>\n  </div>\n  <div class=\"secondary\">\n    <div>\n            <div class=\"mobile-nav-item hidden-mobile nav-item\">\n                <a class=\"register-btn btn\" href=\"/register?next=%2Fapi%2Fcourses%2Fv1%2Fcourses%2Fdsada\">Register</a>\n            </div>\n          <div class=\"mobile-nav-item hidden-mobile nav-item\">\n              <a class=\"sign-in-btn btn\" href=\"/login?next=%2Fapi%2Fcourses%2Fv1%2Fcourses%2Fdsada\">Sign in</a>\n          </div>\n    </div>\n  </div>\n</nav>\n\n    </div>\n    <div class=\"mobile-menu hidden\" aria-label=More Options role=\"menu\" id=\"mobile-menu\"></div>\n</header>\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n    <div class=\"marketing-hero\"></div>\n\n    <div class=\"content-wrapper main-container\" id=\"content\" dir=\"ltr\">\n\n\n\n\n\n\n\n<main id=\"main\" aria-label=\"Content\" tabindex=\"-1\">\n    <section class=\"outside-app\">\n        <h1>\n            Page not found\n        </h1>\n        <p>\n\n                    The page that you were looking for was not found. Go back to the <a href=\"/\">homepage</a> .\n\n        </p>\n    </section>\n</main>\n\n\n    </div>\n\n\n\n\n\n\n\n  <div class=\"wrapper wrapper-footer\">\n    <footer id=\"footer-openedx\" class=\"grid-container\"\n    >\n      <div class=\"colophon\">\n        <nav class=\"nav-colophon\" aria-label=\"About\">\n          <ol>\n              <li class=\"nav-colophon-01\">\n                <a id=\"about\" href=\"/about\">About</a>\n              </li>\n              <li class=\"nav-colophon-02\">\n                <a id=\"blog\" href=\"/blog\">Blog</a>\n              </li>\n              <li class=\"nav-colophon-03\">\n                <a id=\"contact\" href=\"/support/contact_us\">Contact</a>\n              </li>\n              <li class=\"nav-colophon-04\">\n                <a id=\"donate\" href=\"/donate\">Donate</a>\n              </li>\n          </ol>\n        </nav>\n\n\n        <div class=\"wrapper-logo\">\n          <p>\n            <a href=\"/\">\n              <img alt=\"organization logo\" src=\"https://edx-test.ru/static/images/logo.b6c374d66d57.png\">\n            </a>\n          </p>\n        </div>\n\n          <p class=\"copyright\">© Robbo OpenEdx Test .  All rights reserved except where noted.  edX, Open edX and their respective logos are registered trademarks of edX Inc.\n          </p>\n\n        <nav class=\"nav-legal\" aria-label=\"Legal\">\n          <ul>\n              <li class=\"nav-legal-01\">\n                <a href=\"/tos_and_honor\">Terms of Service &amp; Honor Code</a>\n              </li>\n              <li class=\"nav-legal-02\">\n                <a href=\"/privacy\">Privacy Policy</a>\n              </li>\n            <li><a href=\"https://www.edx.org/?utm_medium=affiliate_partner&amp;utm_source=opensource-partner&amp;utm_content=open-edx-partner-footer-link&amp;utm_campaign=open-edx-footer\">Take free online courses at edX.org</a></li>\n          </ul>\n        </nav>\n      </div>\n\n      <div class=\"footer-about-openedx\">\n        <p>\n          <a href=\"https://open.edx.org\">\n            <img src=\"https://files.edx.org/openedx-logos/open-edx-logo-tag.png\" alt=\"Powered by Open edX\" width=\"175\" />\n          </a>\n        </p>\n      </div>\n    </footer>\n  </div>\n\n\n  </div>\n\n\n\n\n\n\n\n\n  <script type=\"text/javascript\" src=\"/static/js/vendor/noreferrer.aa62a3e70ffa.js\" charset=\"utf-8\"></script>\n  <script type=\"text/javascript\" src=\"/static/js/utils/navigation.08930e16ab3d.js\" charset=\"utf-8\"></script>\n  <script type=\"text/javascript\" src=\"/static/js/header/header.309a1243e175.js\"></script>\n\n  <script type=\"text/javascript\" src=\"/static/js/src/jquery_extend_patch.54dddef28d15.js\"></script>\n</body>\n</html>\n"), nil)
			},
		},
		{
			name:         "Empty courseId",
			courseId:     "",
			expectedBody: []byte("\n\n\n\n\n\n\n\n\n<!DOCTYPE html>\n<!--[if lte IE 9]><html class=\"ie ie9 lte9\" lang=\"en\"><![endif]-->\n<!--[if !IE]><!--><html lang=\"en\"><!--<![endif]-->\n<head dir=\"ltr\">\n    <meta charset=\"UTF-8\">\n    <meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n    <meta http-equiv=\"origin-trial\" content=\"ArNBN7d1AkvMhJTGWXlJ8td/AN4lOokzOnqKRNkTnLqaqx0HpfYvmx8JePPs/emKh6O5fckx14LeZIGJ1AQYjgAAAABzeyJvcmlnaW4iOiJodHRwOi8vbG9jYWxob3N0OjE4MDAwIiwiZmVhdHVyZSI6IkRpc2FibGVEaWZmZXJlbnRPcmlnaW5TdWJmcmFtZURpYWxvZ1N1cHByZXNzaW9uIiwiZXhwaXJ5IjoxNjM5NTI2Mzk5fQ==\">\n\n\n\n\n      <title>\n       Page Not Found | Robbo OpenEdx Test\n      </title>\n\n\n      <script type=\"text/javascript\">\n        /* immediately break out of an iframe if coming from the marketing website */\n        (function(window) {\n          if (window.location !== window.top.location) {\n            window.top.location = window.location;\n          }\n        })(this);\n      </script>\n\n\n\n  <script type=\"text/javascript\" src=\"/static/js/i18n/en/djangojs.2623d59dd64d.js\"></script>\n  <script type=\"text/javascript\" src=\"/static/js/ie11_find_array.bd1c6dc7a133.js\"></script>\n\n  <link rel=\"icon\" type=\"image/x-icon\" href=\"/static/images/favicon.03ffbbf95a0d.ico\"/>\n\n\n\n\n\n\n    <link href=\"/static/css/lms-style-vendor.68e48093f5dd.css\" rel=\"stylesheet\" type=\"text/css\" />\n\n\n\n\n\n\n    <link href=\"/static/css/lms-main-v1.5f06a95d4cfd.css\" rel=\"stylesheet\" type=\"text/css\" />\n\n\n\n\n\n\n<script type=\"text/javascript\" src=\"/static/js/lms-main_vendor.3c3c9a2604d6.js\" charset=\"utf-8\"></script>\n\n\n\n\n<script type=\"text/javascript\" src=\"/static/js/lms-application.0e3fae472a20.js\" charset=\"utf-8\"></script>\n\n\n\n\n\n\n    <script type=\"text/javascript\" src=\"/static/bundles/commons.8f14f343047f03c88fc2.d5486becec92.js\" ></script>\n\n\n\n  <script>\n    window.baseUrl = \"/static/\";\n    (function (require) {\n      require.config({\n          baseUrl: window.baseUrl\n      });\n    }).call(this, require || RequireJS.require);\n  </script>\n  <script type=\"text/javascript\" src=\"/static/lms/js/require-config.38226099c6ad.js\"></script>\n\n    <script type=\"text/javascript\">\n        (function (require) {\n          require.config({\n              paths: {\n                'course_bookmarks/js/views/bookmark_button': 'course_bookmarks/js/views/bookmark_button.d4cfaf3361fa',\n'js/views/message_banner': 'js/views/message_banner.141974fd4f5d',\n'moment': 'common/js/vendor/moment-with-locales.43ec67e44612',\n'moment-timezone': 'common/js/vendor/moment-timezone-with-data.60142e6c4416',\n'js/courseware/course_info_events': 'js/courseware/course_info_events.2fc35b57627f',\n'js/courseware/accordion_events': 'js/courseware/accordion_events.6064c7809de5',\n'js/dateutil_factory': 'js/dateutil_factory.be68acdff619',\n'js/courseware/link_clicked_events': 'js/courseware/link_clicked_events',\n'js/courseware/toggle_element_visibility': 'js/courseware/toggle_element_visibility.474ff5ba9de3',\n'js/student_account/logistration_factory': 'js/student_account/logistration_factory.983820206227',\n'js/courseware/courseware_factory': 'js/courseware/courseware_factory.1504fc10caef',\n'js/groups/views/cohorts_dashboard_factory': 'js/groups/views/cohorts_dashboard_factory.ca68388d81d6',\n'js/groups/discussions_management/discussions_dashboard_factory': 'js/discussions_management/views/discussions_dashboard_factory.2e10d9097343',\n'draggabilly': 'js/vendor/draggabilly.26caba6f7187',\n'hls': 'common/js/vendor/hls.5e0c7e1b3bfd'\n            }\n          });\n        }).call(this, require || RequireJS.require);\n    </script>\n\n\n\n\n\n\n\n\n\n\n\n<script type=\"application/json\" id=\"user-metadata\">\n    null\n</script>\n\n\n\n\n\n\n\n\n\n<!-- dummy Segment -->\n<script type=\"text/javascript\">\n  var analytics = {\n    track: function() { return; },\n    trackLink: function() { return; },\n    pageview: function() { return; },\n    page: function() { return; }\n  };\n</script>\n<!-- end dummy Segment -->\n\n\n  <meta name=\"path_prefix\" content=\"\">\n\n\n  <meta name=\"openedx-release-line\" content=\"lilac\" />\n\n\n\n\n\n</head>\n\n<body class=\"ltr  lang_en\">\n\n\n<div id=\"page-prompt\"></div>\n  <div class=\"window-wrap\" dir=\"ltr\">\n\n    <a class=\"nav-skip sr-only sr-only-focusable\" href=\"#main\">Skip to main content</a>\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n<header class=\"global-header \">\n    <div class=\"main-header\">\n\n\n\n\n\n\n\n\n<h1 class=\"header-logo\">\n    <a href=\"/dashboard\">\n\n        <img  class=\"logo\" src=\"/static/images/logo.b6c374d66d57.png\" alt=\"Robbo OpenEdx Test  Home Page\"/>\n\n    </a>\n</h1>\n\n        <div class=\"hamburger-menu\" role=\"button\" aria-label=Options Menu aria-expanded=\"false\" aria-controls=\"mobile-menu\" tabindex=\"0\">\n            <span class=\"line\"></span>\n            <span class=\"line\"></span>\n            <span class=\"line\"></span>\n            <span class=\"line\"></span>\n        </div>\n\n\n\n\n\n\n\n\n\n<nav class=\"nav-links\" aria-label=Supplemental Links>\n  <div class=\"main\">\n      <div class=\"mobile-nav-item hidden-mobile nav-item\">\n        <a href=\"/courses\">Explore courses</a>\n      </div>\n  </div>\n  <div class=\"secondary\">\n    <div>\n            <div class=\"mobile-nav-item hidden-mobile nav-item\">\n                <a class=\"register-btn btn\" href=\"/register?next=%2Fapi%2Fcourses%2Fv1%2Fcourses%2Fdsada\">Register</a>\n            </div>\n          <div class=\"mobile-nav-item hidden-mobile nav-item\">\n              <a class=\"sign-in-btn btn\" href=\"/login?next=%2Fapi%2Fcourses%2Fv1%2Fcourses%2Fdsada\">Sign in</a>\n          </div>\n    </div>\n  </div>\n</nav>\n\n    </div>\n    <div class=\"mobile-menu hidden\" aria-label=More Options role=\"menu\" id=\"mobile-menu\"></div>\n</header>\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n    <div class=\"marketing-hero\"></div>\n\n    <div class=\"content-wrapper main-container\" id=\"content\" dir=\"ltr\">\n\n\n\n\n\n\n\n<main id=\"main\" aria-label=\"Content\" tabindex=\"-1\">\n    <section class=\"outside-app\">\n        <h1>\n            Page not found\n        </h1>\n        <p>\n\n                    The page that you were looking for was not found. Go back to the <a href=\"/\">homepage</a> .\n\n        </p>\n    </section>\n</main>\n\n\n    </div>\n\n\n\n\n\n\n\n  <div class=\"wrapper wrapper-footer\">\n    <footer id=\"footer-openedx\" class=\"grid-container\"\n    >\n      <div class=\"colophon\">\n        <nav class=\"nav-colophon\" aria-label=\"About\">\n          <ol>\n              <li class=\"nav-colophon-01\">\n                <a id=\"about\" href=\"/about\">About</a>\n              </li>\n              <li class=\"nav-colophon-02\">\n                <a id=\"blog\" href=\"/blog\">Blog</a>\n              </li>\n              <li class=\"nav-colophon-03\">\n                <a id=\"contact\" href=\"/support/contact_us\">Contact</a>\n              </li>\n              <li class=\"nav-colophon-04\">\n                <a id=\"donate\" href=\"/donate\">Donate</a>\n              </li>\n          </ol>\n        </nav>\n\n\n        <div class=\"wrapper-logo\">\n          <p>\n            <a href=\"/\">\n              <img alt=\"organization logo\" src=\"https://edx-test.ru/static/images/logo.b6c374d66d57.png\">\n            </a>\n          </p>\n        </div>\n\n          <p class=\"copyright\">© Robbo OpenEdx Test .  All rights reserved except where noted.  edX, Open edX and their respective logos are registered trademarks of edX Inc.\n          </p>\n\n        <nav class=\"nav-legal\" aria-label=\"Legal\">\n          <ul>\n              <li class=\"nav-legal-01\">\n                <a href=\"/tos_and_honor\">Terms of Service &amp; Honor Code</a>\n              </li>\n              <li class=\"nav-legal-02\">\n                <a href=\"/privacy\">Privacy Policy</a>\n              </li>\n            <li><a href=\"https://www.edx.org/?utm_medium=affiliate_partner&amp;utm_source=opensource-partner&amp;utm_content=open-edx-partner-footer-link&amp;utm_campaign=open-edx-footer\">Take free online courses at edX.org</a></li>\n          </ul>\n        </nav>\n      </div>\n\n      <div class=\"footer-about-openedx\">\n        <p>\n          <a href=\"https://open.edx.org\">\n            <img src=\"https://files.edx.org/openedx-logos/open-edx-logo-tag.png\" alt=\"Powered by Open edX\" width=\"175\" />\n          </a>\n        </p>\n      </div>\n    </footer>\n  </div>\n\n\n  </div>\n\n\n\n\n\n\n\n\n  <script type=\"text/javascript\" src=\"/static/js/vendor/noreferrer.aa62a3e70ffa.js\" charset=\"utf-8\"></script>\n  <script type=\"text/javascript\" src=\"/static/js/utils/navigation.08930e16ab3d.js\" charset=\"utf-8\"></script>\n  <script type=\"text/javascript\" src=\"/static/js/header/header.309a1243e175.js\"></script>\n\n  <script type=\"text/javascript\" src=\"/static/js/src/jquery_extend_patch.54dddef28d15.js\"></script>\n</body>\n</html>\n"),
			mockBehavior: func(s *mock_edxApi.MockEdxApiUseCase, courseId string) {
				s.EXPECT().GetCourseContent(courseId).Return([]byte("\n\n\n\n\n\n\n\n\n<!DOCTYPE html>\n<!--[if lte IE 9]><html class=\"ie ie9 lte9\" lang=\"en\"><![endif]-->\n<!--[if !IE]><!--><html lang=\"en\"><!--<![endif]-->\n<head dir=\"ltr\">\n    <meta charset=\"UTF-8\">\n    <meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n    <meta http-equiv=\"origin-trial\" content=\"ArNBN7d1AkvMhJTGWXlJ8td/AN4lOokzOnqKRNkTnLqaqx0HpfYvmx8JePPs/emKh6O5fckx14LeZIGJ1AQYjgAAAABzeyJvcmlnaW4iOiJodHRwOi8vbG9jYWxob3N0OjE4MDAwIiwiZmVhdHVyZSI6IkRpc2FibGVEaWZmZXJlbnRPcmlnaW5TdWJmcmFtZURpYWxvZ1N1cHByZXNzaW9uIiwiZXhwaXJ5IjoxNjM5NTI2Mzk5fQ==\">\n\n\n\n\n      <title>\n       Page Not Found | Robbo OpenEdx Test\n      </title>\n\n\n      <script type=\"text/javascript\">\n        /* immediately break out of an iframe if coming from the marketing website */\n        (function(window) {\n          if (window.location !== window.top.location) {\n            window.top.location = window.location;\n          }\n        })(this);\n      </script>\n\n\n\n  <script type=\"text/javascript\" src=\"/static/js/i18n/en/djangojs.2623d59dd64d.js\"></script>\n  <script type=\"text/javascript\" src=\"/static/js/ie11_find_array.bd1c6dc7a133.js\"></script>\n\n  <link rel=\"icon\" type=\"image/x-icon\" href=\"/static/images/favicon.03ffbbf95a0d.ico\"/>\n\n\n\n\n\n\n    <link href=\"/static/css/lms-style-vendor.68e48093f5dd.css\" rel=\"stylesheet\" type=\"text/css\" />\n\n\n\n\n\n\n    <link href=\"/static/css/lms-main-v1.5f06a95d4cfd.css\" rel=\"stylesheet\" type=\"text/css\" />\n\n\n\n\n\n\n<script type=\"text/javascript\" src=\"/static/js/lms-main_vendor.3c3c9a2604d6.js\" charset=\"utf-8\"></script>\n\n\n\n\n<script type=\"text/javascript\" src=\"/static/js/lms-application.0e3fae472a20.js\" charset=\"utf-8\"></script>\n\n\n\n\n\n\n    <script type=\"text/javascript\" src=\"/static/bundles/commons.8f14f343047f03c88fc2.d5486becec92.js\" ></script>\n\n\n\n  <script>\n    window.baseUrl = \"/static/\";\n    (function (require) {\n      require.config({\n          baseUrl: window.baseUrl\n      });\n    }).call(this, require || RequireJS.require);\n  </script>\n  <script type=\"text/javascript\" src=\"/static/lms/js/require-config.38226099c6ad.js\"></script>\n\n    <script type=\"text/javascript\">\n        (function (require) {\n          require.config({\n              paths: {\n                'course_bookmarks/js/views/bookmark_button': 'course_bookmarks/js/views/bookmark_button.d4cfaf3361fa',\n'js/views/message_banner': 'js/views/message_banner.141974fd4f5d',\n'moment': 'common/js/vendor/moment-with-locales.43ec67e44612',\n'moment-timezone': 'common/js/vendor/moment-timezone-with-data.60142e6c4416',\n'js/courseware/course_info_events': 'js/courseware/course_info_events.2fc35b57627f',\n'js/courseware/accordion_events': 'js/courseware/accordion_events.6064c7809de5',\n'js/dateutil_factory': 'js/dateutil_factory.be68acdff619',\n'js/courseware/link_clicked_events': 'js/courseware/link_clicked_events',\n'js/courseware/toggle_element_visibility': 'js/courseware/toggle_element_visibility.474ff5ba9de3',\n'js/student_account/logistration_factory': 'js/student_account/logistration_factory.983820206227',\n'js/courseware/courseware_factory': 'js/courseware/courseware_factory.1504fc10caef',\n'js/groups/views/cohorts_dashboard_factory': 'js/groups/views/cohorts_dashboard_factory.ca68388d81d6',\n'js/groups/discussions_management/discussions_dashboard_factory': 'js/discussions_management/views/discussions_dashboard_factory.2e10d9097343',\n'draggabilly': 'js/vendor/draggabilly.26caba6f7187',\n'hls': 'common/js/vendor/hls.5e0c7e1b3bfd'\n            }\n          });\n        }).call(this, require || RequireJS.require);\n    </script>\n\n\n\n\n\n\n\n\n\n\n\n<script type=\"application/json\" id=\"user-metadata\">\n    null\n</script>\n\n\n\n\n\n\n\n\n\n<!-- dummy Segment -->\n<script type=\"text/javascript\">\n  var analytics = {\n    track: function() { return; },\n    trackLink: function() { return; },\n    pageview: function() { return; },\n    page: function() { return; }\n  };\n</script>\n<!-- end dummy Segment -->\n\n\n  <meta name=\"path_prefix\" content=\"\">\n\n\n  <meta name=\"openedx-release-line\" content=\"lilac\" />\n\n\n\n\n\n</head>\n\n<body class=\"ltr  lang_en\">\n\n\n<div id=\"page-prompt\"></div>\n  <div class=\"window-wrap\" dir=\"ltr\">\n\n    <a class=\"nav-skip sr-only sr-only-focusable\" href=\"#main\">Skip to main content</a>\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n<header class=\"global-header \">\n    <div class=\"main-header\">\n\n\n\n\n\n\n\n\n<h1 class=\"header-logo\">\n    <a href=\"/dashboard\">\n\n        <img  class=\"logo\" src=\"/static/images/logo.b6c374d66d57.png\" alt=\"Robbo OpenEdx Test  Home Page\"/>\n\n    </a>\n</h1>\n\n        <div class=\"hamburger-menu\" role=\"button\" aria-label=Options Menu aria-expanded=\"false\" aria-controls=\"mobile-menu\" tabindex=\"0\">\n            <span class=\"line\"></span>\n            <span class=\"line\"></span>\n            <span class=\"line\"></span>\n            <span class=\"line\"></span>\n        </div>\n\n\n\n\n\n\n\n\n\n<nav class=\"nav-links\" aria-label=Supplemental Links>\n  <div class=\"main\">\n      <div class=\"mobile-nav-item hidden-mobile nav-item\">\n        <a href=\"/courses\">Explore courses</a>\n      </div>\n  </div>\n  <div class=\"secondary\">\n    <div>\n            <div class=\"mobile-nav-item hidden-mobile nav-item\">\n                <a class=\"register-btn btn\" href=\"/register?next=%2Fapi%2Fcourses%2Fv1%2Fcourses%2Fdsada\">Register</a>\n            </div>\n          <div class=\"mobile-nav-item hidden-mobile nav-item\">\n              <a class=\"sign-in-btn btn\" href=\"/login?next=%2Fapi%2Fcourses%2Fv1%2Fcourses%2Fdsada\">Sign in</a>\n          </div>\n    </div>\n  </div>\n</nav>\n\n    </div>\n    <div class=\"mobile-menu hidden\" aria-label=More Options role=\"menu\" id=\"mobile-menu\"></div>\n</header>\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n    <div class=\"marketing-hero\"></div>\n\n    <div class=\"content-wrapper main-container\" id=\"content\" dir=\"ltr\">\n\n\n\n\n\n\n\n<main id=\"main\" aria-label=\"Content\" tabindex=\"-1\">\n    <section class=\"outside-app\">\n        <h1>\n            Page not found\n        </h1>\n        <p>\n\n                    The page that you were looking for was not found. Go back to the <a href=\"/\">homepage</a> .\n\n        </p>\n    </section>\n</main>\n\n\n    </div>\n\n\n\n\n\n\n\n  <div class=\"wrapper wrapper-footer\">\n    <footer id=\"footer-openedx\" class=\"grid-container\"\n    >\n      <div class=\"colophon\">\n        <nav class=\"nav-colophon\" aria-label=\"About\">\n          <ol>\n              <li class=\"nav-colophon-01\">\n                <a id=\"about\" href=\"/about\">About</a>\n              </li>\n              <li class=\"nav-colophon-02\">\n                <a id=\"blog\" href=\"/blog\">Blog</a>\n              </li>\n              <li class=\"nav-colophon-03\">\n                <a id=\"contact\" href=\"/support/contact_us\">Contact</a>\n              </li>\n              <li class=\"nav-colophon-04\">\n                <a id=\"donate\" href=\"/donate\">Donate</a>\n              </li>\n          </ol>\n        </nav>\n\n\n        <div class=\"wrapper-logo\">\n          <p>\n            <a href=\"/\">\n              <img alt=\"organization logo\" src=\"https://edx-test.ru/static/images/logo.b6c374d66d57.png\">\n            </a>\n          </p>\n        </div>\n\n          <p class=\"copyright\">© Robbo OpenEdx Test .  All rights reserved except where noted.  edX, Open edX and their respective logos are registered trademarks of edX Inc.\n          </p>\n\n        <nav class=\"nav-legal\" aria-label=\"Legal\">\n          <ul>\n              <li class=\"nav-legal-01\">\n                <a href=\"/tos_and_honor\">Terms of Service &amp; Honor Code</a>\n              </li>\n              <li class=\"nav-legal-02\">\n                <a href=\"/privacy\">Privacy Policy</a>\n              </li>\n            <li><a href=\"https://www.edx.org/?utm_medium=affiliate_partner&amp;utm_source=opensource-partner&amp;utm_content=open-edx-partner-footer-link&amp;utm_campaign=open-edx-footer\">Take free online courses at edX.org</a></li>\n          </ul>\n        </nav>\n      </div>\n\n      <div class=\"footer-about-openedx\">\n        <p>\n          <a href=\"https://open.edx.org\">\n            <img src=\"https://files.edx.org/openedx-logos/open-edx-logo-tag.png\" alt=\"Powered by Open edX\" width=\"175\" />\n          </a>\n        </p>\n      </div>\n    </footer>\n  </div>\n\n\n  </div>\n\n\n\n\n\n\n\n\n  <script type=\"text/javascript\" src=\"/static/js/vendor/noreferrer.aa62a3e70ffa.js\" charset=\"utf-8\"></script>\n  <script type=\"text/javascript\" src=\"/static/js/utils/navigation.08930e16ab3d.js\" charset=\"utf-8\"></script>\n  <script type=\"text/javascript\" src=\"/static/js/header/header.309a1243e175.js\"></script>\n\n  <script type=\"text/javascript\" src=\"/static/js/src/jquery_extend_patch.54dddef28d15.js\"></script>\n</body>\n</html>\n"), nil)
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockapi := mock_edxApi.NewMockEdxApiUseCase(ctrl)
			expect := testCase.expectedBody
			testCase.mockBehavior(mockapi, testCase.courseId)
			testApi := &EdxApiUseCaseModule{EdxApiUseCase: mockapi}
			correct, _ := testApi.GetCourseContent(testCase.courseId)
			assert.Equal(t, expect, correct)
		})
	}
}

func TestEdxApiUseCaseImpl_GetEnrollments(t *testing.T) {
	type mockBehavior func(s *mock_edxApi.MockEdxApiUseCase, username string)
	testTable := []struct {
		name         string
		username     string
		expectedBody []byte
		mockBehavior
	}{
		{
			name:         "Ok",
			username:     "edxsom",
			expectedBody: []byte("{\"next\":null,\"previous\":null,\"results\":[{\"created\":\"2022-06-13T03:00:12.571664Z\",\"mode\":\"honor\",\"is_active\":false,\"user\":\"edxsom\",\"course_id\":\"course-v1:TestOrg+02+2022\"},{\"created\":\"2022-06-13T01:16:45.374794Z\",\"mode\":\"honor\",\"is_active\":true,\"user\":\"edxsom\",\"course_id\":\"course-v1:Test_org+01+2022\"}]}\n2022/07/08 21:58:52 {<nil> <nil> [{2022-06-13 03:00:12.571664 +0000 UTC honor false edxsom course-v1:TestOrg+02+2022} {2022-06-13 01:16:45.374794 +0000 UTC honor true edxsom course-v1:Test_org+01+2022}]}\n"),
			mockBehavior: func(s *mock_edxApi.MockEdxApiUseCase, username string) {
				s.EXPECT().GetEnrollments(username).Return([]byte("{\"next\":null,\"previous\":null,\"results\":[{\"created\":\"2022-06-13T03:00:12.571664Z\",\"mode\":\"honor\",\"is_active\":false,\"user\":\"edxsom\",\"course_id\":\"course-v1:TestOrg+02+2022\"},{\"created\":\"2022-06-13T01:16:45.374794Z\",\"mode\":\"honor\",\"is_active\":true,\"user\":\"edxsom\",\"course_id\":\"course-v1:Test_org+01+2022\"}]}\n2022/07/08 21:58:52 {<nil> <nil> [{2022-06-13 03:00:12.571664 +0000 UTC honor false edxsom course-v1:TestOrg+02+2022} {2022-06-13 01:16:45.374794 +0000 UTC honor true edxsom course-v1:Test_org+01+2022}]}\n"), nil)
			},
		},

		{
			name:         "Bad username",
			username:     "dsad",
			expectedBody: []byte("{\"next\":null,\"previous\":null,\"results\":[]}"),
			mockBehavior: func(s *mock_edxApi.MockEdxApiUseCase, username string) {
				s.EXPECT().GetEnrollments(username).Return([]byte("{\"next\":null,\"previous\":null,\"results\":[]}"), nil)
			},
		},
		{
			name:         "Empty username",
			username:     "",
			expectedBody: []byte("{\"next\":null,\"previous\":null,\"results\":[]}"),
			mockBehavior: func(s *mock_edxApi.MockEdxApiUseCase, username string) {
				s.EXPECT().GetEnrollments(username).Return([]byte("{\"next\":null,\"previous\":null,\"results\":[]}"), nil)
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockapi := mock_edxApi.NewMockEdxApiUseCase(ctrl)
			expect := testCase.expectedBody
			testCase.mockBehavior(mockapi, testCase.username)
			testApi := &EdxApiUseCaseModule{EdxApiUseCase: mockapi}
			correct, _ := testApi.GetEnrollments(testCase.username)
			assert.Equal(t, expect, correct)
		})
	}
}

func TestEdxApiUseCaseImpl_GetAllPublicCourses(t *testing.T) {
	type mockBehavior func(s *mock_edxApi.MockEdxApiUseCase, pageNumber int)
	testTable := []struct {
		name         string
		pageNumber   int
		expectedBody []byte
		mockBehavior
	}{
		{
			name:         "Ok",
			pageNumber:   1,
			expectedBody: []byte("{[{https://courses.edx.org/api/courses/v2/blocks/?course_id=ccx-v1%3Aadam%2BMac_APccx%2Be0d%2Bccx%403 2:00 0001-01-01 00:00:00 +0000 UTC 2015-07-21 10:00:00 +0000 U\nTC 2016-06-29 10:00:00 +0000 UTC ccx-v1:adam+Mac_APccx+e0d+ccx@3 map[banner_image:map[uri:/asset-v1:adam+Mac_APccx+e0d+type@asset+block@images_course_image.jpg uri_absolute:https://cou\nrses.edx.org/asset-v1:adam+Mac_APccx+e0d+type@asset+block@images_course_image.jpg] course_image:map[uri:/asset-v1:adam+Mac_APccx+e0d+type@asset+block@Davidson_EdX-68_1_.png] course_vid\neo:map[uri:http://www.youtube.com/watch?v=sAnHwOL8aAs] image:map[large:https://courses.edx.org/asset-v1:adam+Mac_APccx+e0d+type@thumbnail+block@Davidson_EdX-68_1_-png-750x400.jpg raw:h\nttps://courses.edx.org/asset-v1:adam+Mac_APccx+e0d+type@asset+block@Davidson_EdX-68_1_.png small:https://courses.edx.org/asset-v1:adam+Mac_APccx+e0d+type@thumbnail+block@Davidson_EdX-6\n8_1_-png-375x200.jpg]] ntest Mac_APccx adam  2016-06-27 14:10:44 +0000 UTC June 27, 2016 timestamp instructor false false false ccx-v1:adam+Mac_APccx+e0d+ccx@3 <nil>} {https://courses.\nedx.org/api/courses/v2/blocks/?course_id=ccx-v1%3Aadam%2BMac_APccx%2Be0d%2Bccx%404 2:00 0001-01-01 00:00:00 +0000 UTC 2015-07-21 10:00:00 +0000 UTC 2016-06-29 10:00:00 +0000 UTC ccx-v1\n:adam+Mac_APccx+e0d+ccx@4 map[banner_image:map[uri:/asset-v1:adam+Mac_APccx+e0d+type@asset+block@images_course_image.jpg uri_absolute:https://courses.edx.org/asset-v1:adam+Mac_APccx+e0\nd+type@asset+block@images_course_image.jpg] course_image:map[uri:/asset-v1:adam+Mac_APccx+e0d+type@asset+block@Davidson_EdX-68_1_.png] course_video:map[uri:http://www.youtube.com/watch\n?v=sAnHwOL8aAs] image:map[large:https://courses.edx.org/asset-v1:adam+Mac_APccx+e0d+type@thumbnail+block@Davidson_EdX-68_1_-png-750x400.jpg raw:https://courses.edx.org/asset-v1:adam+Ma\nc_APccx+e0d+type@asset+block@Davidson_EdX-68_1_.png small:https://courses.edx.org/asset-v1:adam+Mac_APccx+e0d+type@thumbnail+block@Davidson_EdX-68_1_-png-375x200.jpg]] ntest Mac_APccx \nadam  2016-06-27 14:10:48 +0000 UTC June 27, 2016 timestamp instructor false false false ccx-v1:adam+Mac_APccx+e0d+ccx@4 <nil>} {https://courses.edx.org/api/courses/v2/blocks/?course_i\nd=ccx-v1%3ABerkeleyX%2BCS169.1x%2B3T2015SP%2Bccx%4011 12 hours/week 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC 2016-07-18 00:00:00 +0000 UTC ccx-v1:BerkeleyX+CS169.1x+\n3T2015SP+ccx@11 map[banner_image:map[uri:/asset-v1:BerkeleyX+CS169.1x+3T2015SP+type@asset+block@images_course_image.jpg uri_absolute:https://courses.edx.org/asset-v1:BerkeleyX+CS169.1x\n+3T2015SP+type@asset+block@images_course_image.jpg] course_image:map[uri:/asset-v1:BerkeleyX+CS169.1x+3T2015SP+type@asset+block@ruby-rails-programming-100354527-primary.idge.jpg] cours\ne_video:map[uri:http://www.youtube.com/watch?v=V36LpHqtcDY] image:map[large:https://courses.edx.org/asset-v1:BerkeleyX+CS169.1x+3T2015SP+type@thumbnail+block@ruby-rails-programming-100\n354527-primary.idge-750x400.jpg raw:https://courses.edx.org/asset-v1:BerkeleyX+CS169.1x+3T2015SP+type@asset+block@ruby-rails-programming-100354527-primary.idge.jpg small:https://course\ns.edx.org/asset-v1:BerkeleyX+CS169.1x+3T2015SP+type@thumbnail+block@ruby-rails-programming-100354527-primary.idge-375x200.jpg]] Grinnell CSC321 2016F CS169.1x BerkeleyX CS169.1x teache\ns the fundamentals for engineering long-lived software using Agile techniques to develop Software as a Service (SaaS) using Ruby on Rails. 2016-06-07 20:00:00 +0000 UTC June 7, 2016 ti\nmestamp instructor true false false ccx-v1:BerkeleyX+CS169.1x+3T2015SP+ccx@11 <nil>} {https://courses.edx.org/api/courses/v2/blocks/?course_id=ccx-v1%3ABerkeleyX%2BCS169.1x%2B3T2015SP%\n2Bccx%405 12 hours/week 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC 2016-07-18 00:00:00 +0000 UTC ccx-v1:BerkeleyX+CS169.1x+3T2015SP+ccx@5 map[banner_image:map[uri:/ass\net-v1:BerkeleyX+CS169.1x+3T2015SP+type@asset+block@images_course_image.jpg uri_absolute:https://courses.edx.org/asset-v1:BerkeleyX+CS169.1x+3T2015SP+type@asset+block@images_course_imag\ne.jpg] course_image:map[uri:/asset-v1:BerkeleyX+CS169.1x+3T2015SP+type@asset+block@ruby-rails-programming-100354527-primary.idge.jpg] course_video:map[uri:http://www.youtube.com/watch?\nv=V36LpHqtcDY] image:map[large:https://courses.edx.org/asset-v1:BerkeleyX+CS169.1x+3T2015SP+type@thumbnail+block@ruby-rails-programming-100354527-primary.idge-750x400.jpg raw:https://c\nourses.edx.org/asset-v1:BerkeleyX+CS169.1x+3T2015SP+type@asset+block@ruby-rails-programming-100354527-primary.idge.jpg small:https://courses.edx.org/asset-v1:BerkeleyX+CS169.1x+3T2015S\nP+type@thumbnail+block@ruby-rails-programming-100354527-primary.idge-375x200.jpg]] Test CS169.1x BerkeleyX CS169.1x teaches the fundamentals for engineering long-lived software using A\ngile techniques to develop Software as a Service (SaaS) using Ruby on Rails. 2016-05-24 20:00:00 +0000 UTC May 24, 2016 timestamp instructor true false false ccx-v1:BerkeleyX+CS169.1x+\n3T2015SP+ccx@5 <nil>} {https://courses.edx.org/api/courses/v2/blocks/?course_id=ccx-v1%3ABUx%2BPY1x%2B3T2016%2Bccx%4012  0001-01-01 00:00:00 +0000 UTC 2016-08-08 18:00:00 +0000 UTC 000\n1-01-01 00:00:00 +0000 UTC ccx-v1:BUx+PY1x+3T2016+ccx@12 map[banner_image:map[uri:/asset-v1:BUx+PY1x+3T2016+type@asset+block@images_course_image.jpg uri_absolute:https://courses.edx.or\ng/asset-v1:BUx+PY1x+3T2016+type@asset+block@images_course_image.jpg] course_image:map[uri:/asset-v1:BUx+PY1x+3T2016+type@asset+block@AP_Physics_Dashboard_Image.png] course_video:map[ur\ni:<nil>] image:map[large:https://courses.edx.org/asset-v1:BUx+PY1x+3T2016+type@thumbnail+block@AP_Physics_Dashboard_Image-png-750x400.jpg raw:https://courses.edx.org/asset-v1:BUx+PY1x+\n3T2016+type@asset+block@AP_Physics_Dashboard_Image.png small:https://courses.edx.org/asset-v1:BUx+PY1x+3T2016+type@thumbnail+block@AP_Physics_Dashboard_Image-png-375x200.jpg]] AHS AP P\nhysics PY1x BUx  2016-09-12 16:14:32 +0000 UTC Sept. 12, 2016 timestamp instructor true false false ccx-v1:BUx+PY1x+3T2016+ccx@12 <nil>}] {https://courses.edx.org/api/courses/v1/courses/?page=2&page_size=5 <nil> 20489 4098}}"),
			mockBehavior: func(s *mock_edxApi.MockEdxApiUseCase, pageNumber int) {
				s.EXPECT().GetAllPublicCourses(pageNumber).Return([]byte("{[{https://courses.edx.org/api/courses/v2/blocks/?course_id=ccx-v1%3Aadam%2BMac_APccx%2Be0d%2Bccx%403 2:00 0001-01-01 00:00:00 +0000 UTC 2015-07-21 10:00:00 +0000 U\nTC 2016-06-29 10:00:00 +0000 UTC ccx-v1:adam+Mac_APccx+e0d+ccx@3 map[banner_image:map[uri:/asset-v1:adam+Mac_APccx+e0d+type@asset+block@images_course_image.jpg uri_absolute:https://cou\nrses.edx.org/asset-v1:adam+Mac_APccx+e0d+type@asset+block@images_course_image.jpg] course_image:map[uri:/asset-v1:adam+Mac_APccx+e0d+type@asset+block@Davidson_EdX-68_1_.png] course_vid\neo:map[uri:http://www.youtube.com/watch?v=sAnHwOL8aAs] image:map[large:https://courses.edx.org/asset-v1:adam+Mac_APccx+e0d+type@thumbnail+block@Davidson_EdX-68_1_-png-750x400.jpg raw:h\nttps://courses.edx.org/asset-v1:adam+Mac_APccx+e0d+type@asset+block@Davidson_EdX-68_1_.png small:https://courses.edx.org/asset-v1:adam+Mac_APccx+e0d+type@thumbnail+block@Davidson_EdX-6\n8_1_-png-375x200.jpg]] ntest Mac_APccx adam  2016-06-27 14:10:44 +0000 UTC June 27, 2016 timestamp instructor false false false ccx-v1:adam+Mac_APccx+e0d+ccx@3 <nil>} {https://courses.\nedx.org/api/courses/v2/blocks/?course_id=ccx-v1%3Aadam%2BMac_APccx%2Be0d%2Bccx%404 2:00 0001-01-01 00:00:00 +0000 UTC 2015-07-21 10:00:00 +0000 UTC 2016-06-29 10:00:00 +0000 UTC ccx-v1\n:adam+Mac_APccx+e0d+ccx@4 map[banner_image:map[uri:/asset-v1:adam+Mac_APccx+e0d+type@asset+block@images_course_image.jpg uri_absolute:https://courses.edx.org/asset-v1:adam+Mac_APccx+e0\nd+type@asset+block@images_course_image.jpg] course_image:map[uri:/asset-v1:adam+Mac_APccx+e0d+type@asset+block@Davidson_EdX-68_1_.png] course_video:map[uri:http://www.youtube.com/watch\n?v=sAnHwOL8aAs] image:map[large:https://courses.edx.org/asset-v1:adam+Mac_APccx+e0d+type@thumbnail+block@Davidson_EdX-68_1_-png-750x400.jpg raw:https://courses.edx.org/asset-v1:adam+Ma\nc_APccx+e0d+type@asset+block@Davidson_EdX-68_1_.png small:https://courses.edx.org/asset-v1:adam+Mac_APccx+e0d+type@thumbnail+block@Davidson_EdX-68_1_-png-375x200.jpg]] ntest Mac_APccx \nadam  2016-06-27 14:10:48 +0000 UTC June 27, 2016 timestamp instructor false false false ccx-v1:adam+Mac_APccx+e0d+ccx@4 <nil>} {https://courses.edx.org/api/courses/v2/blocks/?course_i\nd=ccx-v1%3ABerkeleyX%2BCS169.1x%2B3T2015SP%2Bccx%4011 12 hours/week 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC 2016-07-18 00:00:00 +0000 UTC ccx-v1:BerkeleyX+CS169.1x+\n3T2015SP+ccx@11 map[banner_image:map[uri:/asset-v1:BerkeleyX+CS169.1x+3T2015SP+type@asset+block@images_course_image.jpg uri_absolute:https://courses.edx.org/asset-v1:BerkeleyX+CS169.1x\n+3T2015SP+type@asset+block@images_course_image.jpg] course_image:map[uri:/asset-v1:BerkeleyX+CS169.1x+3T2015SP+type@asset+block@ruby-rails-programming-100354527-primary.idge.jpg] cours\ne_video:map[uri:http://www.youtube.com/watch?v=V36LpHqtcDY] image:map[large:https://courses.edx.org/asset-v1:BerkeleyX+CS169.1x+3T2015SP+type@thumbnail+block@ruby-rails-programming-100\n354527-primary.idge-750x400.jpg raw:https://courses.edx.org/asset-v1:BerkeleyX+CS169.1x+3T2015SP+type@asset+block@ruby-rails-programming-100354527-primary.idge.jpg small:https://course\ns.edx.org/asset-v1:BerkeleyX+CS169.1x+3T2015SP+type@thumbnail+block@ruby-rails-programming-100354527-primary.idge-375x200.jpg]] Grinnell CSC321 2016F CS169.1x BerkeleyX CS169.1x teache\ns the fundamentals for engineering long-lived software using Agile techniques to develop Software as a Service (SaaS) using Ruby on Rails. 2016-06-07 20:00:00 +0000 UTC June 7, 2016 ti\nmestamp instructor true false false ccx-v1:BerkeleyX+CS169.1x+3T2015SP+ccx@11 <nil>} {https://courses.edx.org/api/courses/v2/blocks/?course_id=ccx-v1%3ABerkeleyX%2BCS169.1x%2B3T2015SP%\n2Bccx%405 12 hours/week 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC 2016-07-18 00:00:00 +0000 UTC ccx-v1:BerkeleyX+CS169.1x+3T2015SP+ccx@5 map[banner_image:map[uri:/ass\net-v1:BerkeleyX+CS169.1x+3T2015SP+type@asset+block@images_course_image.jpg uri_absolute:https://courses.edx.org/asset-v1:BerkeleyX+CS169.1x+3T2015SP+type@asset+block@images_course_imag\ne.jpg] course_image:map[uri:/asset-v1:BerkeleyX+CS169.1x+3T2015SP+type@asset+block@ruby-rails-programming-100354527-primary.idge.jpg] course_video:map[uri:http://www.youtube.com/watch?\nv=V36LpHqtcDY] image:map[large:https://courses.edx.org/asset-v1:BerkeleyX+CS169.1x+3T2015SP+type@thumbnail+block@ruby-rails-programming-100354527-primary.idge-750x400.jpg raw:https://c\nourses.edx.org/asset-v1:BerkeleyX+CS169.1x+3T2015SP+type@asset+block@ruby-rails-programming-100354527-primary.idge.jpg small:https://courses.edx.org/asset-v1:BerkeleyX+CS169.1x+3T2015S\nP+type@thumbnail+block@ruby-rails-programming-100354527-primary.idge-375x200.jpg]] Test CS169.1x BerkeleyX CS169.1x teaches the fundamentals for engineering long-lived software using A\ngile techniques to develop Software as a Service (SaaS) using Ruby on Rails. 2016-05-24 20:00:00 +0000 UTC May 24, 2016 timestamp instructor true false false ccx-v1:BerkeleyX+CS169.1x+\n3T2015SP+ccx@5 <nil>} {https://courses.edx.org/api/courses/v2/blocks/?course_id=ccx-v1%3ABUx%2BPY1x%2B3T2016%2Bccx%4012  0001-01-01 00:00:00 +0000 UTC 2016-08-08 18:00:00 +0000 UTC 000\n1-01-01 00:00:00 +0000 UTC ccx-v1:BUx+PY1x+3T2016+ccx@12 map[banner_image:map[uri:/asset-v1:BUx+PY1x+3T2016+type@asset+block@images_course_image.jpg uri_absolute:https://courses.edx.or\ng/asset-v1:BUx+PY1x+3T2016+type@asset+block@images_course_image.jpg] course_image:map[uri:/asset-v1:BUx+PY1x+3T2016+type@asset+block@AP_Physics_Dashboard_Image.png] course_video:map[ur\ni:<nil>] image:map[large:https://courses.edx.org/asset-v1:BUx+PY1x+3T2016+type@thumbnail+block@AP_Physics_Dashboard_Image-png-750x400.jpg raw:https://courses.edx.org/asset-v1:BUx+PY1x+\n3T2016+type@asset+block@AP_Physics_Dashboard_Image.png small:https://courses.edx.org/asset-v1:BUx+PY1x+3T2016+type@thumbnail+block@AP_Physics_Dashboard_Image-png-375x200.jpg]] AHS AP P\nhysics PY1x BUx  2016-09-12 16:14:32 +0000 UTC Sept. 12, 2016 timestamp instructor true false false ccx-v1:BUx+PY1x+3T2016+ccx@12 <nil>}] {https://courses.edx.org/api/courses/v1/courses/?page=2&page_size=5 <nil> 20489 4098}}"), nil)
			},
		},

		{
			name:         "Page number is 0",
			pageNumber:   0,
			expectedBody: nil,
			mockBehavior: func(s *mock_edxApi.MockEdxApiUseCase, pageNumber int) {
				s.EXPECT().GetAllPublicCourses(pageNumber).Return(nil, errors.New("Page number is zero or more then page count"))
			},
		},
		{
			name:         "Page number more then page count",
			pageNumber:   423423423,
			expectedBody: nil,
			mockBehavior: func(s *mock_edxApi.MockEdxApiUseCase, pageNumber int) {
				s.EXPECT().GetAllPublicCourses(pageNumber).Return(nil, errors.New("Page number is zero or more then page count"))
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockapi := mock_edxApi.NewMockEdxApiUseCase(ctrl)
			expect := testCase.expectedBody
			testCase.mockBehavior(mockapi, testCase.pageNumber)
			testApi := &EdxApiUseCaseModule{EdxApiUseCase: mockapi}
			correct, _ := testApi.GetAllPublicCourses(testCase.pageNumber)
			assert.Equal(t, expect, correct)
		})
	}
}

func TestEdxApiUseCaseImpl_Login(t *testing.T) {
	type mockBehavior func(s *mock_edxApi.MockEdxApiUseCase, email, password string)
	testTable := []struct {
		name         string
		email        string
		password     string
		expectedBody []byte
		mockBehavior
	}{
		{
			name:         "Ok",
			email:        "insomnia_test32133@fake.email",
			password:     "123456",
			expectedBody: []byte("{\n\t\"success\": true,\n\t\"redirect_url\": null\n}"),
			mockBehavior: func(s *mock_edxApi.MockEdxApiUseCase, email, password string) {
				s.EXPECT().Login(email, password).Return([]byte("{\n\t\"success\": true,\n\t\"redirect_url\": null\n}"), nil)
			},
		},

		{
			name:         "Email or password incorrect",
			email:        "dsadddas",
			password:     "dsadad",
			expectedBody: []byte("{\n\t\"success\": false,\n\t\"value\": \"Email or password is incorrect.\",\n\t\"error_code\": \"incorrect-email-or-password\"\n}"),
			mockBehavior: func(s *mock_edxApi.MockEdxApiUseCase, email, password string) {
				s.EXPECT().Login(email, password).Return([]byte("{\n\t\"success\": false,\n\t\"value\": \"Email or password is incorrect.\",\n\t\"error_code\": \"incorrect-email-or-password\"\n}"), nil)
			},
		},
		{
			name:         "Email or password is empty",
			email:        "",
			password:     "",
			expectedBody: []byte("{\n\t\"success\": false,\n\t\"value\": \"Email or password is incorrect.\",\n\t\"error_code\": \"incorrect-email-or-password\"\n}"),
			mockBehavior: func(s *mock_edxApi.MockEdxApiUseCase, email, password string) {
				s.EXPECT().Login(email, password).Return([]byte("{\n\t\"success\": false,\n\t\"value\": \"Email or password is incorrect.\",\n\t\"error_code\": \"incorrect-email-or-password\"\n}"), nil)
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockapi := mock_edxApi.NewMockEdxApiUseCase(ctrl)
			expect := testCase.expectedBody
			testCase.mockBehavior(mockapi, testCase.email, testCase.password)
			testApi := &EdxApiUseCaseModule{EdxApiUseCase: mockapi}
			correct, _ := testApi.Login(testCase.email, testCase.password)
			assert.Equal(t, expect, correct)
		})
	}
}

func TestEdxApiUseCaseImpl_PostEnrollment(t *testing.T) {
	type mockBehavior func(s *mock_edxApi.MockEdxApiUseCase, message map[string]interface{})
	testTable := []struct {
		name         string
		message      map[string]interface{}
		expectedBody []byte
		mockBehavior
	}{
		{
			name: "Ok",
			message: map[string]interface{}{
				"course_details": map[string]string{
					"course_id": "course-v1:Test_org+01+2022",
				},
				"user": "edxsom",
			},
			expectedBody: []byte("{\n\t\"created\": \"2022-06-13T03:00:12.571664Z\",\n\t\"mode\": \"honor\",\n\t\"is_active\": true,\n\t\"course_details\": {\n\t\t\"course_id\": \"course-v1:TestOrg+02+2022\",\n\t\t\"course_name\": \"Cert test 2 \",\n\t\t\"enrollment_start\": null,\n\t\t\"enrollment_end\": null,\n\t\t\"course_start\": \"2022-01-01T00:00:00Z\",\n\t\t\"course_end\": \"2022-06-12T00:00:00Z\",\n\t\t\"invite_only\": false,\n\t\t\"course_modes\": [\n\t\t\t{\n\t\t\t\t\"slug\": \"honor\",\n\t\t\t\t\"name\": \"Honor_2\",\n\t\t\t\t\"min_price\": 0,\n\t\t\t\t\"suggested_prices\": \"\",\n\t\t\t\t\"currency\": \"usd\",\n\t\t\t\t\"expiration_datetime\": null,\n\t\t\t\t\"description\": null,\n\t\t\t\t\"sku\": null,\n\t\t\t\t\"bulk_sku\": null\n\t\t\t}\n\t\t]\n\t},\n\t\"user\": \"edxsom\"\n}"),
			mockBehavior: func(s *mock_edxApi.MockEdxApiUseCase, message map[string]interface{}) {
				s.EXPECT().PostEnrollment(message).Return([]byte("{\n\t\"created\": \"2022-06-13T03:00:12.571664Z\",\n\t\"mode\": \"honor\",\n\t\"is_active\": true,\n\t\"course_details\": {\n\t\t\"course_id\": \"course-v1:TestOrg+02+2022\",\n\t\t\"course_name\": \"Cert test 2 \",\n\t\t\"enrollment_start\": null,\n\t\t\"enrollment_end\": null,\n\t\t\"course_start\": \"2022-01-01T00:00:00Z\",\n\t\t\"course_end\": \"2022-06-12T00:00:00Z\",\n\t\t\"invite_only\": false,\n\t\t\"course_modes\": [\n\t\t\t{\n\t\t\t\t\"slug\": \"honor\",\n\t\t\t\t\"name\": \"Honor_2\",\n\t\t\t\t\"min_price\": 0,\n\t\t\t\t\"suggested_prices\": \"\",\n\t\t\t\t\"currency\": \"usd\",\n\t\t\t\t\"expiration_datetime\": null,\n\t\t\t\t\"description\": null,\n\t\t\t\t\"sku\": null,\n\t\t\t\t\"bulk_sku\": null\n\t\t\t}\n\t\t]\n\t},\n\t\"user\": \"edxsom\"\n}"), nil)
			},
		},

		{
			name: "Course id incorrect",
			message: map[string]interface{}{
				"course_details": map[string]string{
					"course_id": "dasda",
				},
				"user": "edxsom",
			},
			expectedBody: []byte("{\n\t\"message\": \"No course 'course-v1:' found for enrollment\"\n}"),
			mockBehavior: func(s *mock_edxApi.MockEdxApiUseCase, message map[string]interface{}) {
				s.EXPECT().PostEnrollment(message).Return([]byte("{\n\t\"message\": \"No course 'course-v1:' found for enrollment\"\n}"), nil)
			},
		},
		{
			name: "Username incorrect",
			message: map[string]interface{}{
				"course_details": map[string]string{
					"course_id": "course-v1:Test_org+01+2022",
				},
				"user": "edm",
			},
			expectedBody: []byte("{\n\t\"message\": \"No course 'course-v1:' found for enrollment\"\n}"),
			mockBehavior: func(s *mock_edxApi.MockEdxApiUseCase, message map[string]interface{}) {
				s.EXPECT().PostEnrollment(message).Return([]byte("{\n\t\"message\": \"No course 'course-v1:' found for enrollment\"\n}"), nil)
			},
		},
		{
			name: "Empty field courseId",
			message: map[string]interface{}{
				"course_details": map[string]string{
					"course_id": "",
				},
				"user": "edxsom",
			},
			expectedBody: []byte("{\n\t\"message\": \"Course ID must be specified to create a new enrollment.\"\n}"),
			mockBehavior: func(s *mock_edxApi.MockEdxApiUseCase, message map[string]interface{}) {
				s.EXPECT().PostEnrollment(message).Return([]byte("{\n\t\"message\": \"Course ID must be specified to create a new enrollment.\"\n}"), nil)
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockapi := mock_edxApi.NewMockEdxApiUseCase(ctrl)
			expect := testCase.expectedBody
			testCase.mockBehavior(mockapi, testCase.message)
			testApi := &EdxApiUseCaseModule{EdxApiUseCase: mockapi}
			correct, _ := testApi.PostEnrollment(testCase.message)
			assert.Equal(t, expect, correct)
		})
	}
}

func TestEdxApiUseCaseImpl_PostRegistration(t *testing.T) {
	type mockBehavior func(s *mock_edxApi.MockEdxApiUseCase, registrMsg edxApi.RegistrationForm)
	testTable := []struct {
		name                string
		registrationMessage edxApi.RegistrationForm
		expectedBody        []byte
		mockBehavior
	}{
		{
			name: "Ok",
			registrationMessage: edxApi.RegistrationForm{
				Email:            "insomnia_testrrw223@fake.email",
				Username:         "InsomniaTest31223",
				Name:             "SomeTestName123",
				Password:         "123456",
				Terms_of_service: "true",
			},
			expectedBody: []byte("{\n\t\"success\": true,\n\t\"redirect_url\": \"https://edx-test.ru/dashboard\"\n}"),
			mockBehavior: func(s *mock_edxApi.MockEdxApiUseCase, registrMsg edxApi.RegistrationForm) {
				s.EXPECT().PostRegistration(registrMsg).Return([]byte("{\n\t\"success\": true,\n\t\"redirect_url\": \"https://edx-test.ru/dashboard\"\n}"), nil)
			},
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
			expectedBody: []byte("{\n\t\"password\": [\n\t\t{\n\t\t\t\"user_message\": \"This field is required.\"\n\t\t}\n\t],\n\t\"error_code\": \"validation-error\"\n}"),
			mockBehavior: func(s *mock_edxApi.MockEdxApiUseCase, registrMsg edxApi.RegistrationForm) {
				s.EXPECT().PostRegistration(registrMsg).Return([]byte("{\n\t\"password\": [\n\t\t{\n\t\t\t\"user_message\": \"This field is required.\"\n\t\t}\n\t],\n\t\"error_code\": \"validation-error\"\n}"), nil)
			},
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
			expectedBody: []byte("{\n\t\"email\": [\n\t\t{\n\t\t\t\"user_message\": \"A properly formatted e-mail is required\"\n\t\t}\n\t],\n\t\"error_code\": \"validation-error\"\n}"),
			mockBehavior: func(s *mock_edxApi.MockEdxApiUseCase, registrMsg edxApi.RegistrationForm) {
				s.EXPECT().PostRegistration(registrMsg).Return([]byte("{\n\t\"email\": [\n\t\t{\n\t\t\t\"user_message\": \"A properly formatted e-mail is required\"\n\t\t}\n\t],\n\t\"error_code\": \"validation-error\"\n}"), nil)
			},
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
			expectedBody: []byte("{\n\t\"username\": [\n\t\t{\n\t\t\t\"user_message\": \"Username must be between 2 and 30 characters long.\"\n\t\t}\n\t],\n\t\"error_code\": \"validation-error\"\n}"),
			mockBehavior: func(s *mock_edxApi.MockEdxApiUseCase, registrMsg edxApi.RegistrationForm) {
				s.EXPECT().PostRegistration(registrMsg).Return([]byte("{\n\t\"username\": [\n\t\t{\n\t\t\t\"user_message\": \"Username must be between 2 and 30 characters long.\"\n\t\t}\n\t],\n\t\"error_code\": \"validation-error\"\n}"), nil)
			},
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
			expectedBody: []byte("{\n\t\"name\": [\n\t\t{\n\t\t\t\"user_message\": \"Your legal name must be a minimum of one character long\"\n\t\t}\n\t],\n\t\"error_code\": \"validation-error\"\n}"),
			mockBehavior: func(s *mock_edxApi.MockEdxApiUseCase, registrMsg edxApi.RegistrationForm) {
				s.EXPECT().PostRegistration(registrMsg).Return([]byte("{\n\t\"name\": [\n\t\t{\n\t\t\t\"user_message\": \"Your legal name must be a minimum of one character long\"\n\t\t}\n\t],\n\t\"error_code\": \"validation-error\"\n}"), nil)
			},
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
			expectedBody: []byte("{\n\t\"terms_of_service\": [\n\t\t{\n\t\t\t\"user_message\": \"You must accept the terms of service.\"\n\t\t}\n\t],\n\t\"error_code\": \"validation-error\"\n}"),
			mockBehavior: func(s *mock_edxApi.MockEdxApiUseCase, registrMsg edxApi.RegistrationForm) {
				s.EXPECT().PostRegistration(registrMsg).Return([]byte("{\n\t\"terms_of_service\": [\n\t\t{\n\t\t\t\"user_message\": \"You must accept the terms of service.\"\n\t\t}\n\t],\n\t\"error_code\": \"validation-error\"\n}"), nil)
			},
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
			expectedBody: []byte("{\n\t\t\tname: \"Terms_of_service is empty\",\n\t\t\tregistrationMessage: edxApi.RegistrationForm{\n\t\t\t\tEmail:            \"nsomnia_testrrw223@fake.email\",\n\t\t\t\tUsername:         \"dsadasd\",\n\t\t\t\tName:             \"gdgsdfsfs\",\n\t\t\t\tPassword:         \"123456\",\n\t\t\t\tTerms_of_service: \"\",\n\t\t\t},\n\t\t\texpectedBody: []byte(\"{\\n\\t\\\"terms_of_service\\\": [\\n\\t\\t{\\n\\t\\t\\t\\\"user_message\\\": \\\"You must accept the terms of service.\\\"\\n\\t\\t}\\n\\t],\\n\\t\\\"error_code\\\": \\\"validation-error\\\"\\n}\"),\n\t\t\tmockBehavior: func(s *mock_edxApi.MockEdxApiUseCase, registrMsg edxApi.RegistrationForm) {\n\t\t\t\ts.EXPECT().PostRegistration(registrMsg).Return([]byte(\"{\\n\\t\\\"terms_of_service\\\": [\\n\\t\\t{\\n\\t\\t\\t\\\"user_message\\\": \\\"You must accept the terms of service.\\\"\\n\\t\\t}\\n\\t],\\n\\t\\\"error_code\\\": \\\"validation-error\\\"\\n}\"), nil)\n\t\t\t},\n\t\t}"),
			mockBehavior: func(s *mock_edxApi.MockEdxApiUseCase, registrMsg edxApi.RegistrationForm) {
				s.EXPECT().PostRegistration(registrMsg).Return([]byte("{\n\t\t\tname: \"Terms_of_service is empty\",\n\t\t\tregistrationMessage: edxApi.RegistrationForm{\n\t\t\t\tEmail:            \"nsomnia_testrrw223@fake.email\",\n\t\t\t\tUsername:         \"dsadasd\",\n\t\t\t\tName:             \"gdgsdfsfs\",\n\t\t\t\tPassword:         \"123456\",\n\t\t\t\tTerms_of_service: \"\",\n\t\t\t},\n\t\t\texpectedBody: []byte(\"{\\n\\t\\\"terms_of_service\\\": [\\n\\t\\t{\\n\\t\\t\\t\\\"user_message\\\": \\\"You must accept the terms of service.\\\"\\n\\t\\t}\\n\\t],\\n\\t\\\"error_code\\\": \\\"validation-error\\\"\\n}\"),\n\t\t\tmockBehavior: func(s *mock_edxApi.MockEdxApiUseCase, registrMsg edxApi.RegistrationForm) {\n\t\t\t\ts.EXPECT().PostRegistration(registrMsg).Return([]byte(\"{\\n\\t\\\"terms_of_service\\\": [\\n\\t\\t{\\n\\t\\t\\t\\\"user_message\\\": \\\"You must accept the terms of service.\\\"\\n\\t\\t}\\n\\t],\\n\\t\\\"error_code\\\": \\\"validation-error\\\"\\n}\"), nil)\n\t\t\t},\n\t\t}"), nil)
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockapi := mock_edxApi.NewMockEdxApiUseCase(ctrl)
			expect := testCase.expectedBody
			testCase.mockBehavior(mockapi, testCase.registrationMessage)
			testApi := &EdxApiUseCaseModule{EdxApiUseCase: mockapi}
			correct, _ := testApi.PostRegistration(testCase.registrationMessage)
			assert.Equal(t, expect, correct)
		})
	}
}
