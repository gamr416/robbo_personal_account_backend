package edx_api

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func GetAllPublicCourses(pageNumber int) (respBody string, err error) {
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

func GetCoursesByUser() (respBody string, err error) {
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

func GetWithAuth(url string) (respBody string, err error) {
	err = RefreshToken()

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

func GetEnrollments(username string) (respBody string, err error) {
	return GetWithAuth(viper.GetString("api_urls.getEnrollment") + username)
}
func GetUser() (respBody string, err error) {

	return GetWithAuth(viper.GetString("api_urls.getUser"))
}

func GetCourseContent(courseId string) (respBody string, err error) {

	return GetWithAuth(viper.GetString("api_urls.getCourse") + courseId)
}

func PostEnrollment(message map[string]interface{}) (respBody string, err error) {
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

func PostRegistration(registrationMessage registrationForm) (respBody string, err error) {
	urlAddr := viper.GetString("api_urls.postRegistration")
	response, err := http.PostForm(urlAddr, url.Values{
		"email":            {registrationMessage.email},
		"username":         {registrationMessage.username},
		"name":             {registrationMessage.name},
		"password":         {registrationMessage.password},
		"terms_of_service": {registrationMessage.terms_of_service}})

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
func RefreshToken() (err error) {
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

	newtkn := &NewToken{}
	err = json.Unmarshal(body, newtkn)
	if err != nil {
		log.Println(err)
		return errors.New("Error on json unmarshal")
	}
	viper.Set("api.token", newtkn.AccessToken)
	return nil
}

func Login(email, password string) (respBody string, err error) {
	err = RefreshToken()
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
