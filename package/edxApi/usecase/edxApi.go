package usecase

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/edxApi"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type EdxApiUseCaseImpl struct {
}
type EdxApiUseCaseModule struct {
	fx.Out
	edxApi.EdxApiUseCase
}

func SetupEdxApiUseCase() EdxApiUseCaseModule {

	return EdxApiUseCaseModule{EdxApiUseCase: &EdxApiUseCaseImpl{}}
}

func (p *EdxApiUseCaseImpl) GetAllPublicCourses(pageNumber int) (respBody string, err error) {
	resp, err := http.Get(viper.GetString("api_urls.getAllPublicCourses") + strconv.Itoa(pageNumber) + "&page_size=10")
	if err != nil {
		log.Println(err)
		return "", errors.New("Error on request")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "", errors.New("Error while reading the response bytes")
	}
	return string(body), nil
}

func (p *EdxApiUseCaseImpl) GetCoursesByUser() (respBody string, err error) {
	response, err := http.Get(viper.GetString("api_urls.getCourses"))
	if err != nil {
		log.Println(err)
		return "", errors.New("Error on request")
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return "", errors.New("Error while reading the response bytes")
	}
	return string(body), nil
}

func (p *EdxApiUseCaseImpl) GetWithAuth(url string) (respBody string, err error) {
	err = p.RefreshToken()

	if err != nil {
		log.Println("Token not refresh.\n[ERROR] -", err)
		return "", errors.New("Token not refresh")
	}
	var bearer = "Bearer " + viper.GetString("api.token")

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error on request.\n[ERROR] -", err)
		return "", errors.New("Error on request")
	}

	request.Header.Add("Authorization", bearer)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
		return "", errors.New("Error on response")
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
		return "", errors.New("Error while reading the response bytes")
	}

	return string(body), nil
}

func (p *EdxApiUseCaseImpl) GetEnrollments(username string) (respBody string, err error) {
	return p.GetWithAuth(viper.GetString("api_urls.getEnrollment") + username)
}
func (p *EdxApiUseCaseImpl) GetUser() (respBody string, err error) {

	return p.GetWithAuth(viper.GetString("api_urls.getUser"))
}

func (p *EdxApiUseCaseImpl) GetCourseContent(courseId string) (respBody string, err error) {

	return p.GetWithAuth(viper.GetString("api_urls.getCourse") + courseId)
}

func (p *EdxApiUseCaseImpl) PostEnrollment(message map[string]interface{}) (respBody string, err error) {
	urlAddr := viper.GetString("api_urls.postEnrollment")
	data, err := json.Marshal(message)

	if err != nil {
		log.Println(err)
		return "", errors.New("Error on json Marshal")
	}

	var bearer = "Bearer " + viper.GetString("api.token")

	request, err := http.NewRequest("POST", urlAddr, bytes.NewBuffer(data))
	if err != nil {
		log.Println(err)
		return "", errors.New("Error on request")
	}

	request.Header.Add("Authorization", bearer)
	request.Header.Add("Content-Type", "application/json;charset=utf-8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return "", errors.New("Error on response")
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return "", errors.New("Error while reading the response bytes")
	}
	return string(body), nil
}

func (p *EdxApiUseCaseImpl) PostRegistration(registrationMessage edxApi.RegistrationForm) (respBody string, err error) {
	urlAddr := viper.GetString("api_urls.postRegistration")
	response, err := http.PostForm(urlAddr, url.Values{
		"email":            {registrationMessage.Email},
		"username":         {registrationMessage.Username},
		"name":             {registrationMessage.Name},
		"password":         {registrationMessage.Password},
		"terms_of_service": {registrationMessage.Terms_of_service}})

	if err != nil {
		log.Println(err)
		return "", errors.New("Error on request")
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Println(err)
		return "", errors.New("Error while reading the response bytes")
	}

	return string(body), nil
}
func (p *EdxApiUseCaseImpl) RefreshToken() (err error) {
	urlAddr := viper.GetString("api_urls.refreshToken")
	response, err := http.PostForm(urlAddr, url.Values{
		"grant_type":    {"client_credentials"},
		"client_id":     {viper.GetString("api.client_id")},
		"client_secret": {viper.GetString("api.client_secret")},
	})
	if err != nil {
		log.Println(err)
		return errors.New("Error on request")
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Println(err)
		return errors.New("Error while reading the response bytes")
	}

	newtkn := &edxApi.NewToken{}
	err = json.Unmarshal(body, newtkn)
	if err != nil {
		log.Println(err)
		return errors.New("Error on json unmarshal")
	}
	viper.Set("api.token", newtkn.AccessToken)
	return nil
}

func (p *EdxApiUseCaseImpl) Login(email, password string) (respBody string, err error) {
	err = p.RefreshToken()
	if err != nil {
		log.Println("Token not refresh.\n[ERROR] -", err)
		return "", errors.New("Token not refresh")
	}

	urlAddr := viper.GetString("api_urls.login")
	response, err := http.PostForm(urlAddr, url.Values{
		"email":    {email},
		"password": {password},
	})
	if err != nil {
		return "", errors.New("Error on request")
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", errors.New("Error while reading the response bytes")
	}

	return string(body), nil
}
