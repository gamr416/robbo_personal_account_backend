package factory

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
)

func DataCreateRobboUnit() (data []DataTest) {
	data = []DataTest{
		{
			Name: "Ok",
			Variables: map[string]interface{}{
				"NewRobboUnit": models.NewRobboUnit{
					Name: "name_test",
					City: "city_test",
				},
			},
			ExpectedError: nil,
		},
	}
	return
}
