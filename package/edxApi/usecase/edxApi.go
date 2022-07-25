package usecase

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/edxApi"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

//go:generate mockgen -source=edxApi.go -destination=mocks/mock.go

type EdxApiUseCaseImpl struct {
}
type EdxApiUseCaseModule struct {
	fx.Out
	edxApi.EdxApiUseCase
}

func SetupEdxApiUseCase() EdxApiUseCaseModule {
	return EdxApiUseCaseModule{EdxApiUseCase: &EdxApiUseCaseImpl{}}
}

type myjar struct {
	jar map[string][]*http.Cookie
}

func (p *myjar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	p.jar[u.Host] = cookies
}

func (p *myjar) Cookies(u *url.URL) []*http.Cookie {
	return p.jar[u.Host]
}
func handleСookies(n []*http.Cookie) (csrfToken string, found bool) {
	for _, cookie := range n {
		if cookie.Name == "csrftoken" {
			return cookie.Value, true
		}
	}
	return "", false
}

func (p *EdxApiUseCaseImpl) GetAllPublicCourses(pageNumber int) (respBody []byte, err error) {
	if pageNumber <= 0 && pageNumber >= 5000 {
		return nil, edxApi.ErrIncorrectInputParam
	}
	resp, err := http.Get(viper.GetString("api_urls.getAllPublicCourses") + strconv.Itoa(pageNumber) + "&page_size=5")
	if err != nil {
		log.Println(err)
		return nil, edxApi.ErrOnReq
	}
	if resp.StatusCode != http.StatusOK {
		return nil, edxApi.ErrOnReq
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, edxApi.ErrReadRespBody
	}
	return body, nil
}

func (p *EdxApiUseCaseImpl) GetCoursesByUser() (respBody []byte, err error) {
	response, err := http.Get(viper.GetString("api_urls.getCourses"))
	if err != nil {
		log.Println(err)
		return nil, edxApi.ErrOnReq
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return nil, edxApi.ErrReadRespBody
	}

	return body, nil
}

func (p *EdxApiUseCaseImpl) GetWithAuth(url string) (respBody []byte, err error) {
	err = p.RefreshToken()

	if err != nil {
		log.Println("Token not refresh.\n[ERROR] -", err)
		return nil, edxApi.ErrTknNotRefresh
	}
	var bearer = "Bearer " + viper.GetString("api.token")

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error on request.\n[ERROR] -", err)
		return nil, edxApi.ErrOnReq
	}

	request.Header.Add("Authorization", bearer)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
		return nil, edxApi.ErrOnResp
	}
	if response.StatusCode != http.StatusOK {
		return nil, edxApi.ErrIncorrectInputParam
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
		return nil, edxApi.ErrReadRespBody
	}
	fmt.Println(body)
	return body, nil
}

func (p *EdxApiUseCaseImpl) GetEnrollments(username string) (respBody []byte, err error) {
	if username == "" {
		return nil, edxApi.ErrIncorrectInputParam
	}
	return p.GetWithAuth(viper.GetString("api_urls.getEnrollment") + username)
}
func (p *EdxApiUseCaseImpl) GetUser() (respBody []byte, err error) {

	return p.GetWithAuth(viper.GetString("api_urls.getUser"))
}

func (p *EdxApiUseCaseImpl) GetCourseContent(courseId string) (respBody []byte, err error) {

	return p.GetWithAuth(viper.GetString("api_urls.getCourse") + courseId)
}

func (p *EdxApiUseCaseImpl) PostEnrollment(message map[string]interface{}) (respBody []byte, err error) {
	err = p.RefreshToken()
	if err != nil {
		log.Println("token not refresh")
		return nil, edxApi.ErrTknNotRefresh

	}
	urlAddr := viper.GetString("api_urls.postEnrollment")
	data, err := json.Marshal(message)

	if err != nil {
		log.Println(err)
		return nil, edxApi.ErrJsonMarshal
	}

	var bearer = "Bearer " + viper.GetString("api.token")

	request, err := http.NewRequest("POST", urlAddr, bytes.NewBuffer(data))
	if err != nil {
		log.Println(err)
		return nil, edxApi.ErrOnReq
	}

	request.Header.Add("Authorization", bearer)
	request.Header.Add("Content-Type", "application/json;charset=utf-8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return nil, edxApi.ErrOnResp
	}
	if response.StatusCode != http.StatusOK {
		return nil, edxApi.ErrIncorrectInputParam
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return nil, edxApi.ErrReadRespBody
	}
	return body, nil
}

func (p *EdxApiUseCaseImpl) PostRegistration(registrationMessage *edxApi.RegistrationForm) (respBody []byte, err error) {

	urlAddr := viper.GetString("api_urls.postRegistration")

	client := &http.Client{}
	jar := &myjar{}
	jar.jar = make(map[string][]*http.Cookie)
	client.Jar = jar

	resp, err := client.Get(viper.GetString("api_urls.getRegistration"))
	token, flag := handleСookies(resp.Cookies())
	if flag == false {
		return nil, errors.New("csrf token not found")
	}

	buffer := new(bytes.Buffer)
	params := url.Values{}
	params.Set("email", registrationMessage.Email)
	params.Set("password", registrationMessage.Password)
	params.Set("name", registrationMessage.Name)
	params.Set("username", registrationMessage.Username)
	params.Set("terms_of_service", registrationMessage.TermsOfService)
	buffer.WriteString(params.Encode())

	request, err := http.NewRequest("POST", urlAddr, buffer)
	request.Header.Add("x-csrftoken", token)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	request.Header.Add("Referer", "https://edx-test.ru/login?next=%2F'")
	resp, err = client.Do(request)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Error while reading the response bytes")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, edxApi.ErrIncorrectInputParam
	}
	return body, nil
}
func (p *EdxApiUseCaseImpl) Login(email, password string) (respBody []byte, err error) {

	urlAddr := viper.GetString("api_urls.login")
	client := &http.Client{}
	jar := &myjar{}
	jar.jar = make(map[string][]*http.Cookie)
	client.Jar = jar

	resp, err := client.Get(viper.GetString("api_urls.getLogin"))
	token, flag := handleСookies(resp.Cookies())
	if flag == false {
		return nil, errors.New("csrf token not found")
	}

	buffer := new(bytes.Buffer)
	params := url.Values{}
	params.Set("email", email)
	params.Set("password", password)
	buffer.WriteString(params.Encode())

	request, err := http.NewRequest("POST", urlAddr, buffer)
	request.Header.Add("x-csrftoken", token)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	request.Header.Add("Referer", "https://edx-test.ru/login?next=%2F'")
	resp, err = client.Do(request)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	fmt.Println(string(body))
	if resp.StatusCode != http.StatusOK {
		return nil, edxApi.ErrIncorrectInputParam
	}
	return body, nil

}

func (p *EdxApiUseCaseImpl) RefreshToken() (err error) {
	if viper.GetInt64("api.token_expiration_time") < time.Now().Unix() {
		urlAddr := viper.GetString("api_urls.refreshToken")
		response, err := http.PostForm(urlAddr, url.Values{
			"grant_type":    {"client_credentials"},
			"client_id":     {viper.GetString("api.client_id")},
			"client_secret": {viper.GetString("api.client_secret")},
		})
		if err != nil {
			log.Println(err)
			return edxApi.ErrOnReq
		}
		if response.StatusCode != http.StatusOK {
			return edxApi.ErrIncorrectInputParam
		}

		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Println(err)
			return edxApi.ErrIncorrectInputParam
		}

		newtkn := &edxApi.NewToken{}
		err = json.Unmarshal(body, newtkn)
		if err != nil {
			log.Println(err)
			return errors.New("Error on json unmarshal")
		}

		expirationTime := time.Now().Unix() + int64(newtkn.ExpiresIn)
		viper.Set("api.token_expiration_time", expirationTime)
		viper.Set("api.token", newtkn.AccessToken)
		return nil
	} else {
		return nil
	}
}
