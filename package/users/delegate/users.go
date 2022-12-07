package delegate

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/users"
	"go.uber.org/fx"
	"log"
)

type UsersDelegateImpl struct {
	UseCase users.UseCase
}

func (p *UsersDelegateImpl) GetStudentsByTeacherId(teacherId string) (students []*models.StudentHTTP, err error) {
	studentsCore, getStudentsErr := p.UseCase.GetStudentsByTeacherId(teacherId)
	if getStudentsErr != nil {
		err = getStudentsErr
		return
	}
	for _, studentCore := range studentsCore {
		studentHttpTemp := models.StudentHTTP{
			UserHTTP:     &models.UserHTTP{},
			RobboGroupID: studentCore.RobboGroupId,
			RobboUnitID:  studentCore.RobboUnitId,
		}
		studentHttpTemp.FromCore(studentCore)
		students = append(students, &studentHttpTemp)
	}
	return
}

func (p *UsersDelegateImpl) GetTeachersByStudentId(studentId string) (teachers []*models.TeacherHTTP, err error) {
	teachersCore, getTeacherErr := p.UseCase.GetTeachersByStudentId(studentId)
	if getTeacherErr != nil {
		err = getTeacherErr
		return
	}
	for _, teacherCore := range teachersCore {
		teacherHttpTemp := models.TeacherHTTP{
			UserHTTP: &models.UserHTTP{},
		}
		teacherHttpTemp.FromCore(teacherCore)
		teachers = append(teachers, &teacherHttpTemp)
	}
	return
}

func (p *UsersDelegateImpl) GetStudentsByRobboUnitId(robboUnitId string) (students []*models.StudentHTTP, err error) {
	studentsCore, getStudentsErr := p.UseCase.GetStudentsByRobboUnitId(robboUnitId)
	if getStudentsErr != nil {
		err = getStudentsErr
		return
	}
	for _, studentCore := range studentsCore {
		studentHttpTemp := models.StudentHTTP{
			UserHTTP:     &models.UserHTTP{},
			RobboGroupID: studentCore.RobboGroupId,
			RobboUnitID:  robboUnitId,
		}
		studentHttpTemp.FromCore(studentCore)
		students = append(students, &studentHttpTemp)
	}
	return
}

func (p *UsersDelegateImpl) GetTeacherByRobboGroupId(robboGroupId string) (teachers []*models.TeacherHTTP, err error) {
	teachersCore, getTeacherErr := p.UseCase.GetTeacherByRobboGroupId(robboGroupId)
	if getTeacherErr != nil {
		err = getTeacherErr
		return
	}
	for _, teacherCore := range teachersCore {
		teacherHttpTemp := models.TeacherHTTP{
			UserHTTP: &models.UserHTTP{},
		}
		teacherHttpTemp.FromCore(teacherCore)
		teachers = append(teachers, &teacherHttpTemp)
	}
	return
}

func (p *UsersDelegateImpl) GetStudentsByRobboGroupId(robboGroupId string) (students []*models.StudentHTTP, err error) {
	studentsCore, err := p.UseCase.GetStudentsByRobboGroupId(robboGroupId)
	if err != nil {
		return
	}
	for _, studentCore := range studentsCore {
		studentHttpTemp := models.StudentHTTP{
			UserHTTP:     &models.UserHTTP{},
			RobboGroupID: studentCore.RobboGroupId,
			RobboUnitID:  studentCore.RobboUnitId,
		}
		studentHttpTemp.FromCore(studentCore)
		students = append(students, &studentHttpTemp)
	}
	return
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

func (p *UsersDelegateImpl) CreateStudent(student *models.StudentHTTP, parentId string) (id string, err error) {
	studentCore := student.ToCore()
	return p.UseCase.CreateStudent(studentCore, parentId)
}

func (p *UsersDelegateImpl) DeleteStudent(studentId uint) (err error) {
	return p.UseCase.DeleteStudent(studentId)
}

func (p *UsersDelegateImpl) GetStudentById(studentId string) (student *models.StudentHTTP, err error) {
	studentCore, err := p.UseCase.GetStudentById(studentId)
	if err != nil {
		log.Println("User not found")
		return student, auth.ErrUserNotFound
	}
	student = &models.StudentHTTP{
		UserHTTP: &models.UserHTTP{},
	}
	student.FromCore(studentCore)
	return
}

func (p *UsersDelegateImpl) SearchStudentByEmail(email string, parentId string) (students []*models.StudentHTTP, err error) {
	studentsCore, err := p.UseCase.SearchStudentByEmail(email, parentId)
	if err != nil {
		return
	}
	for _, studentCore := range studentsCore {
		studentTemp := models.StudentHTTP{
			UserHTTP:     &models.UserHTTP{},
			RobboGroupID: "",
			RobboUnitID:  "",
		}
		studentTemp.FromCore(studentCore)
		students = append(students, &studentTemp)
	}
	return
}

func (p *UsersDelegateImpl) GetStudentByParentId(parentId string) (students []*models.StudentHTTP, err error) {
	studentsCore, err := p.UseCase.GetStudentByParentId(parentId)
	if err != nil {
		return
	}
	for _, studentCore := range studentsCore {
		studentHttpTemp := models.StudentHTTP{
			UserHTTP:     &models.UserHTTP{},
			RobboGroupID: "",
			RobboUnitID:  "",
		}
		studentHttpTemp.FromCore(studentCore)
		students = append(students, &studentHttpTemp)
	}
	return
}

func (p *UsersDelegateImpl) UpdateStudent(studentHTTP *models.StudentHTTP) (err error) {
	studentCore := studentHTTP.ToCore()
	return p.UseCase.UpdateStudent(studentCore)
}

func (p *UsersDelegateImpl) AddStudentToRobboGroup(studentId string, robboGroupId string, robboUnitId string) (err error) {
	return p.UseCase.AddStudentToRobboGroup(studentId, robboGroupId, robboUnitId)
}

func (p *UsersDelegateImpl) CreateStudentTeacherRelation(teacherId, studentId string) (err error) {
	return p.UseCase.CreateStudentTeacherRelation(studentId, teacherId)
}

func (p *UsersDelegateImpl) DeleteStudentTeacherRelation(teacherId, studentId string) (err error) {
	return p.UseCase.DeleteStudentTeacherRelation(studentId, teacherId)
}

func (p *UsersDelegateImpl) GetTeacherById(teacherId string) (teacher *models.TeacherHTTP, err error) {
	teacherCore, err := p.UseCase.GetTeacherById(teacherId)
	if err != nil {
		log.Println("User not found")
		return teacher, auth.ErrUserNotFound
	}
	teacher = &models.TeacherHTTP{
		UserHTTP: &models.UserHTTP{},
	}
	teacher.FromCore(&teacherCore)
	return
}

func (p *UsersDelegateImpl) GetAllTeachers() (teachers []*models.TeacherHTTP, err error) {
	teachersCore, err := p.UseCase.GetAllTeachers()
	if err != nil {
		return
	}
	for _, teacherCore := range teachersCore {
		teacherTemp := models.TeacherHTTP{
			UserHTTP: &models.UserHTTP{},
		}
		teacherTemp.FromCore(&teacherCore)
		teachers = append(teachers, &teacherTemp)
	}
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

func (p *UsersDelegateImpl) GetParentById(parentId string) (parent *models.ParentHTTP, err error) {
	parentCore, err := p.UseCase.GetParentById(parentId)
	if err != nil {
		return parent, auth.ErrUserNotFound
	}
	parent = &models.ParentHTTP{
		UserHTTP: &models.UserHTTP{},
	}
	parent.FromCore(*parentCore)
	return
}

func (p *UsersDelegateImpl) GetAllParent() (parents []*models.ParentHTTP, err error) {
	parentsCore, err := p.UseCase.GetAllParent()
	if err != nil {
		return
	}
	for _, parentCore := range parentsCore {
		parentTemp := models.ParentHTTP{
			UserHTTP: &models.UserHTTP{},
			Children: []*models.StudentHTTP{},
		}
		parentTemp.FromCore(*parentCore)
		parents = append(parents, &parentTemp)
	}
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

func (p *UsersDelegateImpl) GetFreeListenerById(freeListenerId string) (freeListener models.FreeListenerHttp, err error) {
	freeListenerCore, err := p.UseCase.GetFreeListenerById(freeListenerId)
	if err != nil {
		log.Println("free listener not found")
		return freeListener, auth.ErrUserNotFound
	}
	freeListener.FromCore(freeListenerCore)
	return
}

func (p *UsersDelegateImpl) CreateFreeListener(freeListenerHTTP *models.FreeListenerHttp) (id string, err error) {
	freeListenerCore := freeListenerHTTP.ToCore()
	return p.UseCase.CreateFreeListener(freeListenerCore)
}

func (p *UsersDelegateImpl) DeleteFreeListener(freeListenerId uint) (err error) {
	return p.UseCase.DeleteFreeListener(freeListenerId)
}

func (p *UsersDelegateImpl) UpdateFreeListener(freeListenerHTTP *models.FreeListenerHttp) (err error) {
	freeListenerCore := freeListenerHTTP.ToCore()
	return p.UseCase.UpdateFreeListener(freeListenerCore)
}

func (p *UsersDelegateImpl) GetUnitAdminById(unitAdminId string) (unitAdmin models.UnitAdminHTTP, err error) {
	unitAdminCore, err := p.UseCase.GetUnitAdminById(unitAdminId)
	if err != nil {
		log.Println("User not found")
		return unitAdmin, auth.ErrUserNotFound
	}
	unitAdmin = models.UnitAdminHTTP{
		UserHTTP: &models.UserHTTP{},
	}
	unitAdmin.FromCore(unitAdminCore)
	return
}

func (p *UsersDelegateImpl) GetAllUnitAdmins() (unitAdmins []*models.UnitAdminHTTP, err error) {
	unitAdminsCore, err := p.UseCase.GetAllUnitAdmins()
	if err != nil {
		return
	}
	for _, unitAdminCore := range unitAdminsCore {
		unitAdminHttpTemp := models.UnitAdminHTTP{
			UserHTTP: &models.UserHTTP{},
		}
		unitAdminHttpTemp.FromCore(unitAdminCore)
		unitAdmins = append(unitAdmins, &unitAdminHttpTemp)
	}
	return
}

func (p *UsersDelegateImpl) GetUnitAdminByRobboUnitId(robboUnitId string) (unitAdmins []*models.UnitAdminHTTP, err error) {
	unitAdminsCore, err := p.UseCase.GetUnitAdminByRobboUnitId(robboUnitId)
	if err != nil {
		return
	}
	for _, unitAdminCore := range unitAdminsCore {
		unitAdminHttpTemp := models.UnitAdminHTTP{
			UserHTTP: &models.UserHTTP{},
		}
		unitAdminHttpTemp.FromCore(unitAdminCore)
		unitAdmins = append(unitAdmins, &unitAdminHttpTemp)
	}
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

func (p *UsersDelegateImpl) SearchUnitAdminByEmail(email string, robboUnitId string) (unitAdmins []*models.UnitAdminHTTP, err error) {
	unitAdminsCore, err := p.UseCase.SearchUnitAdminByEmail(email, robboUnitId)
	if err != nil {
		return
	}
	for _, unitAdminCore := range unitAdminsCore {
		unitAdminHttpTemp := models.UnitAdminHTTP{
			UserHTTP: &models.UserHTTP{},
		}
		unitAdminHttpTemp.FromCore(unitAdminCore)
		unitAdmins = append(unitAdmins, &unitAdminHttpTemp)
	}
	return
}

func (p *UsersDelegateImpl) GetSuperAdminById(superAdminId string) (superAdmin models.SuperAdminHTTP, err error) {
	superAdminCore, err := p.UseCase.GetSuperAdminById(superAdminId)
	if err != nil {
		log.Println("User not found")
		return superAdmin, auth.ErrUserNotFound
	}
	superAdmin = models.SuperAdminHTTP{
		UserHTTP: &models.UserHTTP{},
	}
	superAdmin.FromCore(superAdminCore)
	return
}

func (p *UsersDelegateImpl) UpdateSuperAdmin(superAdminHTTP *models.SuperAdminHTTP) (err error) {
	superAdminCore := superAdminHTTP.ToCore()
	return p.UseCase.UpdateSuperAdmin(superAdminCore)
}
func (p *UsersDelegateImpl) DeleteSuperAdmin(superAdminId uint) (err error) {
	return p.UseCase.DeleteSuperAdmin(superAdminId)
}

func (p *UsersDelegateImpl) CreateRelation(parentId, childrenId string) (err error) {
	return p.UseCase.CreateRelation(parentId, childrenId)
}

func (p *UsersDelegateImpl) SetNewUnitAdminForRobboUnit(unitAdminId, robboUnitId string) (err error) {
	return p.UseCase.SetNewUnitAdminForRobboUnit(unitAdminId, robboUnitId)
}

func (p *UsersDelegateImpl) DeleteUnitAdminForRobboUnit(unitAdminId, robboUnitId string) (err error) {
	return p.UseCase.DeleteUnitAdminForRobboUnit(unitAdminId, robboUnitId)
}
