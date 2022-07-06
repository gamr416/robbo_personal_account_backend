package usecase

import (
	"bytes"
	"encoding/json"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/courses"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type CourseUseCaseImpl struct {
	courses.Gateway
}

type NewToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

type CourseUseCaseModule struct {
	fx.Out
	courses.UseCase
}

func SetupCourseUseCase(gateway courses.Gateway) CourseUseCaseModule {
	return CourseUseCaseModule{
		UseCase: &CourseUseCaseImpl{
			Gateway: gateway,
		},
	}
}

func (p *CourseUseCaseImpl) CreateCourse(course *models.CourseHTTP, courseId string) (id string, statusCode int, err error) {
	statusCode, err = RefreshToken()
	if err != nil {
		log.Println(err)
		return "", statusCode, err
	}
	body, statusCode, err := p.GetCourseContent(courseId)
	if err != nil {
		log.Println(err)
		return "", statusCode, err
	}
	err = json.Unmarshal([]byte(body), &course)
	if err != nil {
		log.Println(err)
		return "", http.StatusBadRequest, err
	}
	courseCore := course.ToCore()
	return p.Gateway.CreateCourse(courseCore)
}

func (p *CourseUseCaseImpl) UpdateCourse(course *models.CourseCore) (err error) {
	return p.Gateway.UpdateCourse(course)
}

func (p *CourseUseCaseImpl) DeleteCourse(course *models.CourseCore) (err error) {
	return p.Gateway.DeleteCourse(course)
}

func (p *CourseUseCaseImpl) GetAllPublicCourses(pageNumber int) (respBody string, statusCode int, err error) {
	resp, err := http.Get(viper.GetString("api_urls.getAllPublicCourses") + strconv.Itoa(pageNumber) + "&page_size=10")
	if err != nil {
		log.Println(err)
		return "", http.StatusBadRequest, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "", http.StatusBadGateway, err
	}
	return string(body), resp.StatusCode, nil
}

func (p *CourseUseCaseImpl) GetCoursesForUser() (respBody string, statusCode int, err error) {
	response, err := http.Get(viper.GetString("api_urls.getCourses"))
	if err != nil {
		log.Println(err)
		return "", http.StatusBadGateway, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return "", http.StatusBadGateway, err
	}
	return string(body), response.StatusCode, nil
}

func GetWithAuth(url string) (respBody string, statusCode int, err error) {
	statusCode, err = RefreshToken()

	if err != nil {
		log.Println("Token not refresh.\n[ERROR] -", err)
		return "", statusCode, err
	}
	var bearer = "Bearer " + viper.GetString("api.token")

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error on request.\n[ERROR] -", err)
		return "", http.StatusBadRequest, err
	}

	request.Header.Add("Authorization", bearer)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
		return "", http.StatusBadGateway, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
		return "", http.StatusBadGateway, err
	}

	return string(body), response.StatusCode, nil
}

func (p *CourseUseCaseImpl) GetEnrollments(username string) (respBody string, statusCode int, err error) {
	return GetWithAuth(viper.GetString("api_urls.getEnrollment") + username)
}
func (p *CourseUseCaseImpl) GetUser() (respBody string, statusCode int, err error) {

	return GetWithAuth(viper.GetString("api_urls.getUser"))
}

func (p *CourseUseCaseImpl) GetCourseContent(courseId string) (respBody string, statusCode int, err error) {

	return GetWithAuth(viper.GetString("api_urls.getCourse") + courseId)
}

func (p *CourseUseCaseImpl) PostEnrollment(message map[string]interface{}) (respBody string, statusCode int, err error) {
	urlAddr := viper.GetString("api_urls.postEnrollment")
	data, err := json.Marshal(message)

	if err != nil {
		log.Println(err)
		return "", http.StatusBadGateway, err
	}

	var bearer = "Bearer " + viper.GetString("api.token")

	request, err := http.NewRequest("POST", urlAddr, bytes.NewBuffer(data))
	if err != nil {
		log.Println(err)
		return "", http.StatusBadRequest, err
	}

	request.Header.Add("Authorization", bearer)
	request.Header.Add("Content-Type", "application/json;charset=utf-8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return "", http.StatusBadGateway, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return "", http.StatusBadGateway, err
	}
	return string(body), response.StatusCode, nil
}

/*
func (p *CourseUseCaseImpl) PostRegistration(registrationMessage registrationForm) (respBody string, statusCode int) {
	urlAddr := viper.GetString("api_urls.postRegistration")
	response, err := http.PostForm(urlAddr, url.Values{
		"email":            {registrationMessage.email},
		"username":         {registrationMessage.username},
		"name":             {registrationMessage.name},
		"password":         {registrationMessage.password},
		"terms_of_service": {registrationMessage.terms_of_service}})

	if err != nil {
		log.Println(err)
		return "", http.StatusBadRequest, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Println(err)
		return "", http.StatusBadGateway, err
	}

	return string(body), response.StatusCode
}*/
func RefreshToken() (statusCode int, err error) {
	urlAddr := viper.GetString("api_urls.refreshToken")
	response, err := http.PostForm(urlAddr, url.Values{
		"grant_type":    {"client_credentials"},
		"client_id":     {viper.GetString("api.client_id")},
		"client_secret": {viper.GetString("api.client_secret")},
	})
	if err != nil {
		log.Println(err)
		return http.StatusBadRequest, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Println(err)
		return http.StatusBadGateway, err
	}

	newtkn := &NewToken{}
	err = json.Unmarshal(body, newtkn)
	if err != nil {
		log.Println(err)
		return http.StatusBadGateway, err
	}
	viper.Set("api.token", newtkn.AccessToken)
	return http.StatusOK, err
}
