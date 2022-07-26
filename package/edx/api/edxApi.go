package api

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/edx"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type EdxApiAuthImpl struct {
}

type EdxApiAuthModule struct {
	fx.Out
	edx.AuthUseCase
}

func (p *EdxApiAuthImpl) GetWithAuth(url string) (respBody []byte, err error) {
	err = p.RefreshToken()

	if err != nil {
		log.Println("Token not refresh.\n[ERROR] -", err)
		return nil, edx.ErrTknNotRefresh
	}
	var bearer = "Bearer " + viper.GetString("api.token")

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error on request.\n[ERROR] -", err)
		return nil, edx.ErrOnReq
	}

	request.Header.Add("Authorization", bearer)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
		return nil, edx.ErrOnResp
	}
	if response.StatusCode != http.StatusOK {
		return nil, edx.ErrIncorrectInputParam
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
		return nil, edx.ErrReadRespBody
	}
	return body, nil
}

func (p *EdxApiAuthImpl) PostWithAuth(url string, params map[string]interface{}) (respBody []byte, err error) {
	err = p.RefreshToken()
	if err != nil {
		log.Println("token not refresh")
		return nil, edx.ErrTknNotRefresh

	}

	data, err := json.Marshal(params)

	if err != nil {
		log.Println(err)
		return nil, edx.ErrJsonMarshal
	}

	var bearer = "Bearer " + viper.GetString("api.token")

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
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

func (p *EdxApiAuthImpl) RefreshToken() (err error) {
	if viper.GetInt64("api.token_expiration_time") < time.Now().Unix() {
		urlAddr := viper.GetString("api_urls.refreshToken")
		response, err := http.PostForm(urlAddr, url.Values{
			"grant_type":    {"client_credentials"},
			"client_id":     {viper.GetString("api.client_id")},
			"client_secret": {viper.GetString("api.client_secret")},
		})
		if err != nil {
			log.Println(err)
			return edx.ErrOnReq
		}
		if response.StatusCode != http.StatusOK {
			return edx.ErrIncorrectInputParam
		}

		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Println(err)
			return edx.ErrIncorrectInputParam
		}

		newtkn := &edx.NewToken{}
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
