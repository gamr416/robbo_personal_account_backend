package factory

import (
	"github.com/hasura/go-graphql-client"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/courses"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
)

func DataCreateCourseAccessRelationRobboGroup() (data []DataTest) {
	data = []DataTest{
		{
			Name: "Ok",
			Variables: map[string]interface{}{
				"NewCourseAccessRelationRobboGroup": models.NewAccessCourseRelationRobboGroup{
					CourseID:     "1",
					RobboGroupID: "1",
				},
			},
			ExpectedError: nil,
		},
		{
			Name: "Incorrect course id",
			Variables: map[string]interface{}{
				"NewCourseAccessRelationRobboGroup": models.NewAccessCourseRelationRobboGroup{
					CourseID:     "test",
					RobboGroupID: "1",
				},
			},
			ExpectedError: graphql.Error{
				Message:   courses.ErrIncorrectInputParam.Error(),
				Locations: nil,
			},
		},
		{
			Name: "Incorrect robbo group id",
			Variables: map[string]interface{}{
				"NewCourseAccessRelationRobboGroup": models.NewAccessCourseRelationRobboGroup{
					CourseID:     "1",
					RobboGroupID: "test",
				},
			},
			ExpectedError: graphql.Error{
				Message:   courses.ErrIncorrectInputParam.Error(),
				Locations: nil,
			},
		},
	}
	return
}
