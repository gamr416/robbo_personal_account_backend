package delegate

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/users"
	"go.uber.org/fx"
	"log"
)

type UsersDelegateImpl struct {
	users.UseCase
}

type UsersDelegateModule struct {
	fx.Out
	users.Delegate
}

func SetupUsersDelegate(usecase users.UseCase) UsersDelegateModule {
	return UsersDelegateModule{
		Delegate: &UsersDelegateImpl{
			usecase,
		},
	}
}

func (p *UsersDelegateImpl) GetUnitAdminById(unitAdminId uint) (unitAdmin models.UnitAdminHTTP, err error) {
	unAd, err := p.UseCase.GetUnitAdminById(unitAdminId)
	if err != nil {
		log.Println("User not found")
		return unitAdmin, auth.ErrUserNotFound
	}
	unitAdmin.FromCore(unAd)
	return
}

func (p *UsersDelegateImpl) CreateStudent(student *models.StudentHTTP) (id string, err error) {
	studentCore := student.ToCore()
	return p.UseCase.CreateStudent(studentCore)
}

func (p *UsersDelegateImpl) DeleteStudent(studentId uint) (err error) {
	return p.UseCase.DeleteStudent(studentId)
}

func (p *UsersDelegateImpl) GetStudentById(studentId uint) (student models.StudentHTTP, err error) {
	studentCore, err := p.UseCase.GetStudentById(studentId)
	if err != nil {
		log.Println("User not found")
		return student, auth.ErrUserNotFound
	}
	student.FromCore(studentCore)
	return
}

func (p *UsersDelegateImpl) UpdateStudent(studentHTTP *models.StudentHTTP) (err error) {
	studentCore := studentHTTP.ToCore()
	return p.UseCase.UpdateStudent(studentCore)
}

func (p *UsersDelegateImpl) GetStudent(email, password string) (student models.StudentHTTP, err error) {
	studentCore, err := p.UseCase.GetStudent(email, password)
	if err != nil {
		log.Println("User not found")
		return student, auth.ErrUserNotFound
	}
	student.FromCore(studentCore)
	return
}

func (p *UsersDelegateImpl) GetTeacherById(teacherId uint) (teacher models.TeacherHTTP, err error) {
	teacherCore, err := p.UseCase.GetTeacherById(teacherId)
	if err != nil {
		log.Println("User not found")
		return teacher, auth.ErrUserNotFound
	}
	teacher.FromCore(teacherCore)
	return
}

func (p *UsersDelegateImpl) GetTeacher(email, password string) (teacher models.TeacherHTTP, err error) {
	t, err := p.UseCase.GetTeacher(email, password)
	if err != nil {
		log.Println("User not found")
		return teacher, auth.ErrUserNotFound
	}
	teacher.FromCore(t)
	return
}

func (p *UsersDelegateImpl) UpdateTeacher(teacherHTTP *models.TeacherHTTP) (err error) {
	teacherCore := teacherHTTP.ToCore()
	return p.UseCase.UpdateTeacher(teacherCore)
}

func (p *UsersDelegateImpl) CreateTeacher(teacherHTTP *models.TeacherHTTP) (id string, err error) {
	teacherCore := teacherHTTP.ToCore()
	return p.UseCase.CreateTeacher(teacherCore)
}

func (p *UsersDelegateImpl) DeleteTeacher(teacherId uint) (err error) {
	return p.UseCase.DeleteTeacher(teacherId)
}

func (p *UsersDelegateImpl) GetParent(email, password string) (parent models.ParentHTTP, err error) {
	pr, err := p.UseCase.GetParent(email, password)
	if err != nil {
		log.Println("User not found")
		return parent, auth.ErrUserNotFound
	}
	parent.FromCore(pr)
	return
}

func (p *UsersDelegateImpl) GetParentById(parentId uint) (parent models.ParentHTTP, err error) {
	pr, err := p.UseCase.GetParentById(parentId)
	if err != nil {
		log.Println("User not found")
		return parent, auth.ErrUserNotFound
	}
	parent.FromCore(pr)
	return
}

func (p *UsersDelegateImpl) CreateParent(parentHTTP *models.ParentHTTP) (id string, err error) {
	parentCore := parentHTTP.ToCore()
	return p.UseCase.CreateParent(parentCore)
}

func (p *UsersDelegateImpl) DeleteParent(parentId uint) (err error) {
	return p.UseCase.DeleteParent(parentId)
}

func (p *UsersDelegateImpl) UpdateParent(parentHTTP *models.ParentHTTP) (err error) {
	parentCore := parentHTTP.ToCore()
	return p.UseCase.UpdateParent(parentCore)
}

func (p *UsersDelegateImpl) GetFreeListener(email, password string) (freeListener models.FreeListenerHttp, err error) {
	freeListCore, err := p.UseCase.GetFreeListener(email, password)
	if err != nil {
		return
	}
	freeListener.FromCore(freeListCore)
	return
}

func (p *UsersDelegateImpl) GetFreeListenerById(parentId uint) (freeListener models.FreeListenerHttp, err error) {
	pr, err := p.UseCase.GetParentById(parentId)
	if err != nil {
		log.Println("User not found")
		return parent, auth.ErrUserNotFound
	}
	parent.FromCore(pr)
	return
}

func (p *UsersDelegateImpl) CreateFreeListener(freeListenerHTTP *models.FreeListenerHttp) (id string, err error) {
	parentCore := parentHTTP.ToCore()
	return p.UseCase.CreateParent(parentCore)
}

func (p *UsersDelegateImpl) DeleteFreeListener(freeListenerId uint) (err error) {
	return p.UseCase.DeleteParent(parentId)
}

func (p *UsersDelegateImpl) UpdateFreeListener(freeListenerHTTP *models.FreeListenerHttp) (err error) {
	parentCore := parentHTTP.ToCore()
	return p.UseCase.UpdateParent(parentCore)
}

func (p *UsersDelegateImpl) GetUnitAdmin(email, password string) (unitAdmin models.UnitAdminHTTP, err error) {
	unAd, err := p.UseCase.GetUnitAdmin(email, password)
	if err != nil {
		log.Println("User not found")
		return unitAdmin, auth.ErrUserNotFound
	}
	unitAdmin.FromCore(unAd)
	return
}

func (p *UsersDelegateImpl) GetSuperAdminById(superAdminId uint) (superAdmin models.SuperAdminHTTP, err error) {
	sa, err := p.UseCase.GetSuperAdminById(superAdminId)
	if err != nil {
		log.Println("User not found")
		return superAdmin, auth.ErrUserNotFound
	}
	superAdmin.FromCore(sa)
	return
}

func (p *UsersDelegateImpl) CreateUnitAdmin(unitAdminHTTP *models.UnitAdminHTTP) (id string, err error) {
	unitAdminCore := unitAdminHTTP.ToCore()
	return p.UseCase.CreateUnitAdmin(unitAdminCore)
}

func (p *UsersDelegateImpl) UpdateUnitAdmin(unitAdminHTTP *models.UnitAdminHTTP) (err error) {
	unitAdminCore := unitAdminHTTP.ToCore()
	return p.UseCase.UpdateUnitAdmin(unitAdminCore)
}

func (p *UsersDelegateImpl) DeleteUnitAdmin(unitAdminId uint) (err error) {
	return p.UseCase.DeleteUnitAdmin(unitAdminId)
}

func (p *UsersDelegateImpl) GetSuperAdmin(email, password string) (superAdmin models.SuperAdminHTTP, err error) {
	sa, err := p.UseCase.GetSuperAdmin(email, password)
	if err != nil {
		log.Println("User not found")
		return superAdmin, auth.ErrUserNotFound
	}
	superAdmin.FromCore(sa)
	return
}
