package cohorts

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/edx"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type EdxApiCohortImpl struct {
	edx.AuthUseCase
}
type EdxApiCohortModule struct {
	fx.Out
	edx.CohortUseCase
}

func (p *EdxApiCohortImpl) CreateCohort(courseId string, cohortParams map[string]interface{}) (respBody []byte, err error) {
	urlAddr := viper.GetString("api_urls.postCohort") + courseId + "/cohorts/"
	return p.PostWithAuth(urlAddr, cohortParams)
}

func (p *EdxApiCohortImpl) AddStudent(username, courseId string, cohortId int) (respBody []byte, err error) {
	err = p.RefreshToken()
	if err != nil {
		log.Println("token not refresh")
		return nil, edx.ErrTknNotRefresh

	}

	var bearer = "Bearer " + viper.GetString("api.token")
	urlAddr := viper.GetString("api_urls.postCohort") + courseId + "/cohorts/" + strconv.Itoa(cohortId) + "/users/" + username
	request, err := http.NewRequest("POST", urlAddr, nil)
	if err != nil {
		log.Println(err)
		return nil, edx.ErrOnReq
	}

	request.Header.Add("Authorization", bearer)
	request.Header.Add("Content-Type", "application/json;charset=utf-8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return nil, edx.ErrOnResp
	}
	if response.StatusCode != http.StatusOK {
		return nil, edx.ErrIncorrectInputParam
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return nil, edx.ErrReadRespBody
	}
	return body, nil
}
