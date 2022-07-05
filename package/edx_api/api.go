package edx_api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type NewToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

var EnrollmentMessage = map[string]interface{}{
	"course_details": map[string]string{
		"course_id": "course-v1:Test_org+01+2022",
	},
	"user": "tesr_user",
}

type registrationForm struct {
	email            string
	username         string
	name             string
	password         string
	terms_of_service string
}

func GetAllCourses(pageNumber int) (string, int) {
	resp, err := http.Get("https://courses.edx.org/api/courses/v1/courses/?page=" + strconv.Itoa(pageNumber) + "&page_size=10")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//log.Println("All Courses: " + string(body))
	return string(body), resp.StatusCode
}

func GetCourses() (string, int) {
	response, err := http.Get("https://edx-test.ru/api/courses/v1/courses/")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(body), response.StatusCode
}

func GetWithAuth(funcName, url, token string) (respBody string, statusCode int) {
	var bearer = "Bearer " + token

	request, err := http.NewRequest("GET", url, nil)

	request.Header.Add("Authorization", bearer)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	//log.Println(funcName + ": " + string([]byte(body)))

	return string([]byte(body)), response.StatusCode
}

func GetEnrollment(username, token string) (string, int) {
	return GetWithAuth("Get Enrollment", "https://edx-test.ru/api/enrollment/v1/enrollments?username="+username, token)
}
func GetUser(token string) (string, int) {

	return GetWithAuth("Get User", "https://edx-test.ru/api/user/v1/me", token)
}

func GetCourse(coueseId, token string) (string, int) {
	return GetWithAuth("Get Course", "https://edx-test.ru/api/courses/v1/courses/"+coueseId, token)
}

func PostEnrollment(url, token string, message map[string]interface{}) (string, int) {

	data, err := json.Marshal(message)

	if err != nil {
		log.Fatal(err)
	}

	var bearer = "Bearer " + token

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Add("Authorization", bearer)
	request.Header.Add("Content-Type", "application/json;charset=utf-8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return string(body), response.StatusCode
}

func PostRegistration(urlAddr string, registrationMessage registrationForm) (string, int) {

	response, err := http.PostForm(urlAddr, url.Values{
		"email":            {registrationMessage.email},
		"username":         {registrationMessage.username},
		"name":             {registrationMessage.name},
		"password":         {registrationMessage.password},
		"terms_of_service": {registrationMessage.terms_of_service}})

	if err != nil {
		log.Fatalln(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalln(err)
	}

	return string(body), response.StatusCode
}
func RefreshToken(token *string) {
	urlAddr := "https://edx-test.ru/oauth2/access_token"
	response, err := http.PostForm(urlAddr, url.Values{
		"grant_type":    {"client_credentials"},
		"client_id":     {"uYSlyzAlmR2dEuukRWkqjcYjK6BplIYyXo8SawJl"},
		"client_secret": {"vWXLUQf9KsOFtl5RhxkNozjIVV3gWZDcUERpLXd8ypAtrHmiCo1tLG0N9TV4ryIgnOcZZt3Cnc4oUXLUpfacXB3ocSBJCIJAK5cax6xm9ElhBjHlc6dKPE8rd72lZWQM"},
	})
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalln(err)
	}

	newtkn := &NewToken{}
	err = json.Unmarshal(body, newtkn)

	*token = newtkn.AccessToken
}

func TestApi() {
	var token string
	RefreshToken(&token)

	//fmt.Println(GetCourse("course-v1:TestOrg+02+2022", token))//+
	//fmt.Println(GetEnrollment("edxsom", token))//+
	//fmt.Println(GetUser(token))//+
	//fmt.Println(GetCourses())//+
	//fmt.Println(PostEnrollment("https://edx-test.ru/api/enrollment/v1/enrollment", token, EnrollmentMessage)) //+

	//fmt.Println(PostRegistration("https://edx-test.ru/api/user/v1/account/registration/", regMsg)) //
	//fmt.Println(GetAllCourses(1, 1))//+
}
