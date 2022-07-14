package models

type CohortCore struct {
	Name           string
	ID             uint
	UserCount      uint
	AssignmentType string
}

type CohortHTTP struct {
	Name           string `json:"name"`
	ID             uint   `json:"id"`
	UserCount      uint   `json:"user_count"`
	AssignmentType string `json:"assignment_type"`
}

type CreateCohortHTTP struct {
	Message map[string]interface{} `json:"message"`
}

type CohortDB struct {
	Name           string
	ID             uint
	UserCount      uint
	AssignmentType string
}

func (ht *CohortHTTP) ToCore() *CohortCore {
	return &CohortCore{
		Name:           ht.Name,
		ID:             ht.ID,
		UserCount:      ht.UserCount,
		AssignmentType: ht.AssignmentType,
	}
}

func (ht *CohortHTTP) FromCore(cohortCore *CohortCore) {
	ht.Name = cohortCore.Name
	ht.ID = cohortCore.ID
	ht.UserCount = cohortCore.UserCount
	ht.AssignmentType = cohortCore.AssignmentType
}
