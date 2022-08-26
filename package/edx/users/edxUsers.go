package users

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/edx"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type EdxApiUserImpl struct {
	edx.AuthUseCase
}
type EdxApiUserModule struct {
	fx.Out
	edx.UserUseCase
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

func (p *EdxApiUserImpl) PostRegistration(registrationMessage edx.RegistrationForm) (respBody []byte, err error) {

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
		return nil, edx.ErrIncorrectInputParam
	}
	return body, nil
}
func (p *EdxApiUserImpl) Login(email, password string) (respBody []byte, err error) {

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
		return nil, edx.ErrIncorrectInputParam
	}
	return body, nil

}
func (p *EdxApiUserImpl) GetUser() (respBody []byte, err error) {
	return p.GetWithAuth(viper.GetString("api_urls.getUser"))
}
