package edxApiCohortsUsecase

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/edxApi"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/edxApi/edxApiUseCase"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type EdxApiCohortImpl struct {
}
type EdxApiCohortModule struct {
	fx.Out
	edxApi.EdxApiCohort
}

func SetupEdxApiCohort() EdxApiCohortModule {
	return EdxApiCohortModule{EdxApiCohort: &EdxApiCohortImpl{}}
}

func (p *EdxApiCohortImpl) CreateCohort(courseId string, cohortParams map[string]interface{}) (respBody []byte, err error) {
	urlAddr := viper.GetString("api_urls.postCohort") + courseId + "/cohorts/"
	return edxApiUseCase.PostWithAuth(urlAddr, cohortParams)
}

func (p *EdxApiCohortImpl) AddStudent(username, courseId string, cohortId int) (respBody []byte, err error) {
	err = edxApiUseCase.RefreshToken()
	if err != nil {
		log.Println("token not refresh")
		return nil, edxApi.ErrTknNotRefresh

	}

	var bearer = "Bearer " + viper.GetString("api.token")
	urlAddr := viper.GetString("api_urls.postCohort") + courseId + "/cohorts/" + strconv.Itoa(cohortId) + "/users/" + username
	request, err := http.NewRequest("POST", urlAddr, nil)
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
