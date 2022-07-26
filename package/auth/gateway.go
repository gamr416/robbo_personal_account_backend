package auth

type Gateway interface {
<<<<<<< HEAD
	GetUser(email, password string) (user *models.UserCore, err error)
	CreateUser(user *models.UserCore) (id string, err error)
=======
	//GetUser(email, password string) (user *models.UserCore, err error)
>>>>>>> b51413c19b53a2b40776b2746be8d694d6f8e40e
}
