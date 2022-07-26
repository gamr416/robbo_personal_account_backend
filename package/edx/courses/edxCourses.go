package courses

import (
	"errors"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/edx"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

//go:generate mockgen -source=edx.go -destination=mocks/mock.go

type EdxApiCourseImpl struct {
	edx.AuthUseCase
}
type EdxApiCourseModule struct {
	fx.Out
	edx.CourseUseCase
}

func (p *EdxApiCourseImpl) GetAllPublicCourses(pageNumber int) (respBody []byte, err error) {
	if pageNumber <= 0 && pageNumber >= 5000 {
		return nil, errors.New("Page number is zero or more then page count")
	}
	resp, err := http.Get(viper.GetString("api_urls.getAllPublicCourses") + strconv.Itoa(pageNumber) + "&page_size=5")
	if err != nil {
		log.Println(err)
		return nil, edx.ErrOnReq
	}
	if resp.StatusCode != http.StatusOK {
		return nil, edx.ErrIncorrectInputParam
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, edx.ErrReadRespBody
	}
	return body, nil
}

func (p *EdxApiCourseImpl) GetCoursesByUser() (respBody []byte, err error) {
	response, err := http.Get(viper.GetString("api_urls.getCourses"))
	if err != nil {
		log.Println(err)
		return nil, edx.ErrOnReq
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return nil, edx.ErrReadRespBody
	}

	return body, nil
}

func (p *EdxApiCourseImpl) GetEnrollments(username string) (respBody []byte, err error) {
	return p.GetWithAuth(viper.GetString("api_urls.getEnrollment") + username)
}

func (p *EdxApiCourseImpl) GetUser() (respBody []byte, err error) {
	return p.GetWithAuth(viper.GetString("api_urls.getUser"))
}

func (p *EdxApiCourseImpl) GetCourseContent(courseId string) (respBody []byte, err error) {
	return p.GetWithAuth(viper.GetString("api_urls.getCourse") + courseId)
}

func (p *EdxApiCourseImpl) PostEnrollment(message map[string]interface{}) (respBody []byte, err error) {
	return p.PostWithAuth(viper.GetString("api_urls.postEnrollment"), message)
}
