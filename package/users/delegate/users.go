package delegate

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/edxApi"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/users"
	"go.uber.org/fx"
	"log"
)

type UsersDelegateImpl struct {
	users.UseCase
	edxApi.EdxApiUseCase
}

type UsersDelegateModule struct {
	fx.Out
	users.Delegate
}

func SetupUsersDelegate(usecase users.UseCase, edxApi edxApi.EdxApiUseCase) UsersDelegateModule {
	return UsersDelegateModule{
		Delegate: &UsersDelegateImpl{
			usecase,
			edxApi,
		},
	}
}

func (p *UsersDelegateImpl) GetUserById(userId int) (user models.UserHttp, err error) {
	userCore, err := p.UseCase.GetUserById(userId)
	if err != nil {
		log.Println("User not found")
		return user, auth.ErrUserNotFound
	}
	user.FromCore(userCore)
	return
}

func (p *UsersDelegateImpl) CreateUser(user *models.UserHttp) (id string, err error) {
	userCore := user.ToCore()
	return p.UseCase.CreateUser(userCore)
}

func (p *UsersDelegateImpl) GetUsersByRole(role *models.Role) (users *[]models.UserHttp, err error) {

	usersCore, err := p.UseCase.GetUsersByRole(role)

	var usersToHttp []models.UserHttp
	for _, userCore := range *usersCore {
		var userHttp models.UserHttp
		userHttp.FromCore(&userCore)
		usersToHttp = append(usersToHttp, userHttp)
	}
	return &usersToHttp, err
}

func (p *UsersDelegateImpl) LoginUser(email, password string) (err error) {
	_, err = p.EdxApiUseCase.Login(email, password)
	if err != nil {
		log.Println(err)
		return err
	}
	return
}

func (p *UsersDelegateImpl) UpdateUser(userHTTP *models.UserHttp) (err error) {
	userCore := userHTTP.ToCore()
	return p.UseCase.UpdateUser(userCore)
}

func (p *UsersDelegateImpl) PostEnrollment(postEnrollmentHTTP *models.PostEnrollmentHTTP) (err error) {
	_, err = p.EdxApiUseCase.PostEnrollment(postEnrollmentHTTP.Message)
	if err != nil {
		log.Println(err)
		return err
	}
	return
}

func (p *UsersDelegateImpl) PostUnenroll(postUnenrollHTTP *models.PostEnrollmentHTTP) (err error) {
	_, err = p.EdxApiUseCase.PostEnrollment(postUnenrollHTTP.Message)
	if err != nil {
		log.Println(err)
		return err
	}
	return
}

func (p *UsersDelegateImpl) DeleteUser(userId int) (err error) {
	return p.UseCase.DeleteUser(userId)
}

func (p *UsersDelegateImpl) RegistrationUser(userForm *edxApi.RegistrationForm) (err error) {
	_, err = p.EdxApiUseCase.PostRegistration(userForm)
	if err != nil {
		log.Println(err)
		return err
	}
	return
}
