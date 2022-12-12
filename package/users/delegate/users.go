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

func (p *UsersDelegateImpl) CreateStudent(student *models.StudentHTTP, parentId string) (newStudent *models.StudentHTTP, err error) {
	studentCore := student.ToCore()
	newStudentCore, err := p.UseCase.CreateStudent(studentCore, parentId)
	if err != nil {
		log.Println(err)
		return
	}
	newStudent = &models.StudentHTTP{
		UserHTTP: &models.UserHTTP{},
	}
	newStudent.FromCore(newStudentCore)
	return
}

func (p *UsersDelegateImpl) DeleteStudent(studentId string) (err error) {
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

func (p *UsersDelegateImpl) UpdateStudent(studentHTTP *models.StudentHTTP) (studentUpdated *models.StudentHTTP, err error) {
	studentCore := studentHTTP.ToCore()
	studentUpdatedCore, err := p.UseCase.UpdateStudent(studentCore)
	if err != nil {
		log.Println(err)
		return
	}
	studentUpdated = &models.StudentHTTP{
		UserHTTP: &models.UserHTTP{},
	}
	studentUpdated.FromCore(studentUpdatedCore)
	return
}

func (p *UsersDelegateImpl) AddStudentToRobboGroup(studentId string, robboGroupId string, robboUnitId string) (err error) {
	return p.UseCase.AddStudentToRobboGroup(studentId, robboGroupId, robboUnitId)
}

func (p *UsersDelegateImpl) CreateStudentTeacherRelation(studentId, teacherId string) (student *models.StudentHTTP, err error) {
	studentCore, err := p.UseCase.CreateStudentTeacherRelation(studentId, teacherId)
	if err != nil {
		return
	}
	student = &models.StudentHTTP{
		UserHTTP: &models.UserHTTP{},
	}
	student.FromCore(studentCore)
	return
}

func (p *UsersDelegateImpl) DeleteStudentTeacherRelation(studentId, teacherId string) (student *models.StudentHTTP, err error) {
	studentCore, err := p.UseCase.DeleteStudentTeacherRelation(studentId, teacherId)
	if err != nil {
		return
	}
	student = &models.StudentHTTP{
		UserHTTP: &models.UserHTTP{},
	}
	student.FromCore(studentCore)
	return
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

func (p *UsersDelegateImpl) UpdateTeacher(teacherHTTP *models.TeacherHTTP) (teacherUpdated models.TeacherHTTP, err error) {
	teacherCore := teacherHTTP.ToCore()
	teacherUpdatedCore, err := p.UseCase.UpdateTeacher(teacherCore)
	if err != nil {
		log.Println(err)
		return
	}
	teacherUpdated = models.TeacherHTTP{
		UserHTTP: &models.UserHTTP{},
	}
	teacherUpdated.FromCore(&teacherUpdatedCore)
	return
}

func (p *UsersDelegateImpl) CreateTeacher(teacherHTTP *models.TeacherHTTP) (newTeacher models.TeacherHTTP, err error) {
	teacherCore := teacherHTTP.ToCore()
	newTeacherCore, err := p.UseCase.CreateTeacher(teacherCore)
	if err != nil {
		log.Println(err)
		return
	}
	newTeacher = models.TeacherHTTP{
		UserHTTP: &models.UserHTTP{},
	}
	newTeacher.FromCore(&newTeacherCore)
	return
}

func (p *UsersDelegateImpl) DeleteTeacher(teacherId string) (err error) {
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

func (p *UsersDelegateImpl) CreateParent(parentHTTP *models.ParentHTTP) (newParent *models.ParentHTTP, err error) {
	parentCore := parentHTTP.ToCore()
	newParentCore, err := p.UseCase.CreateParent(parentCore)
	if err != nil {
		log.Println(err)
		return
	}
	newParent = &models.ParentHTTP{
		UserHTTP: &models.UserHTTP{},
		Children: []*models.StudentHTTP{},
	}
	newParent.FromCore(*newParentCore)
	return
}

func (p *UsersDelegateImpl) DeleteParent(parentId string) (err error) {
	return p.UseCase.DeleteParent(parentId)
}

func (p *UsersDelegateImpl) UpdateParent(parentHTTP *models.ParentHTTP) (parentUpdated *models.ParentHTTP, err error) {
	parentCore := parentHTTP.ToCore()
	parentUpdatedCore, err := p.UseCase.UpdateParent(parentCore)
	if err != nil {
		log.Println(err)
		return
	}
	parentUpdated = &models.ParentHTTP{
		UserHTTP: &models.UserHTTP{},
		Children: []*models.StudentHTTP{},
	}
	parentUpdated.FromCore(*parentUpdatedCore)
	return
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

func (p *UsersDelegateImpl) CreateFreeListener(freeListenerHTTP *models.FreeListenerHttp) (newFreeListener *models.FreeListenerHttp, err error) {
	freeListenerCore := freeListenerHTTP.ToCore()
	newFreeListenerCore, err := p.UseCase.CreateFreeListener(freeListenerCore)
	if err != nil {
		log.Println(err)
		return
	}
	newFreeListener = &models.FreeListenerHttp{
		UserHTTP: models.UserHTTP{},
	}
	newFreeListener.FromCore(newFreeListenerCore)
	return
}

func (p *UsersDelegateImpl) DeleteFreeListener(freeListenerId string) (err error) {
	return p.UseCase.DeleteFreeListener(freeListenerId)
}

func (p *UsersDelegateImpl) UpdateFreeListener(freeListenerHTTP *models.FreeListenerHttp) (freeListenerUpdated *models.FreeListenerHttp, err error) {
	freeListenerCore := freeListenerHTTP.ToCore()
	freeListenerUpdatedCore, err := p.UseCase.UpdateFreeListener(freeListenerCore)
	if err != nil {
		log.Println(err)
		return
	}
	freeListenerUpdated = &models.FreeListenerHttp{
		UserHTTP: models.UserHTTP{},
	}
	freeListenerUpdated.FromCore(freeListenerUpdatedCore)
	return
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

func (p *UsersDelegateImpl) CreateUnitAdmin(unitAdminHTTP *models.UnitAdminHTTP) (newUnitAdmin *models.UnitAdminHTTP, err error) {
	unitAdminCore := unitAdminHTTP.ToCore()
	newUnitAdminCore, err := p.UseCase.CreateUnitAdmin(unitAdminCore)
	if err != nil {
		log.Println(err)
		return
	}
	newUnitAdmin = &models.UnitAdminHTTP{
		UserHTTP: &models.UserHTTP{},
	}
	newUnitAdmin.FromCore(newUnitAdminCore)
	return
}

func (p *UsersDelegateImpl) UpdateUnitAdmin(unitAdminHTTP *models.UnitAdminHTTP) (unitAdminUpdated *models.UnitAdminHTTP, err error) {
	unitAdminCore := unitAdminHTTP.ToCore()
	unitAdminUpdatedCore, err := p.UseCase.UpdateUnitAdmin(unitAdminCore)
	if err != nil {
		log.Println(err)
		return
	}
	unitAdminUpdated = &models.UnitAdminHTTP{
		UserHTTP: &models.UserHTTP{},
	}
	unitAdminUpdated.FromCore(unitAdminUpdatedCore)
	return
}

func (p *UsersDelegateImpl) DeleteUnitAdmin(unitAdminId string) (err error) {
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

func (p *UsersDelegateImpl) UpdateSuperAdmin(superAdminHTTP *models.SuperAdminHTTP) (superAdminUpdated *models.SuperAdminHTTP, err error) {
	superAdminCore := superAdminHTTP.ToCore()
	superAdminUpdatedCore, err := p.UseCase.UpdateSuperAdmin(superAdminCore)
	if err != nil {
		log.Println(err)
		return
	}
	superAdminUpdated = &models.SuperAdminHTTP{
		UserHTTP: &models.UserHTTP{},
	}
	superAdminUpdated.FromCore(superAdminUpdatedCore)
	return
}

func (p *UsersDelegateImpl) DeleteSuperAdmin(superAdminId string) (err error) {
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
