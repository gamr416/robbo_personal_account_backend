package gateway

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/db_client"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/users"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"log"
	"strconv"
)

type UsersGatewayImpl struct {
	PostgresClient *db_client.PostgresClient
}

type UsersGatewayModule struct {
	fx.Out
	users.Gateway
}

func SetupUsersGateway(postgresClient db_client.PostgresClient) UsersGatewayModule {
	return UsersGatewayModule{
		Gateway: &UsersGatewayImpl{PostgresClient: &postgresClient},
	}
}

func (r *UsersGatewayImpl) GetUnitAdminById(unitAdminId uint) (unitAdmin *models.UnitAdminCore, err error) {
	var unitAdminDb models.UnitAdminDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("id = ?", unitAdminId).First(&unitAdminDb).Error; err != nil {
			err = auth.ErrUserNotFound
			log.Println(err)
			return
		}
		return
	})
	unitAdmin = unitAdminDb.ToCore()
	return
}

func (r *UsersGatewayImpl) GetParentById(parentId uint) (parent *models.ParentCore, err error) {
	var parentDb models.ParentDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("id = ?", parentId).First(&parentDb).Error; err != nil {
			err = auth.ErrUserNotFound
			log.Println(err)
			return
		}
		return
	})
	parent = parentDb.ToCore()
	return
}

func (r *UsersGatewayImpl) CreateUnitAdmin(unitAdmin *models.UnitAdminCore) (id string, err error) {
	unitAdminDb := models.UnitAdminDB{}
	unitAdminDb.FromCore(unitAdmin)
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&unitAdminDb).Error
		if err != nil {
			log.Println(err)
			return
		}
		return
	})
	if err != nil {
		log.Println(err)
		return "", err
	}
	id = strconv.FormatUint(uint64(unitAdminDb.ID), 10)
	return id, nil
}

func (r *UsersGatewayImpl) UpdateUnitAdmin(unitAdmin *models.UnitAdminCore) (err error) {
	unitAdminDb := models.UnitAdminDB{}
	unitAdminDb.FromCore(unitAdmin)
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Where("id = ?", unitAdminDb.ID).First(&models.UnitAdminDB{}).Error
		if err != nil {
			err = auth.ErrUserNotFound
			log.Println(err)
			return
		}
		err = tx.Model(&unitAdminDb).Where("id = ?", unitAdminDb.ID).Updates(unitAdminDb).Error
		if err != nil {
			log.Println(err)
			return
		}
		return
	})
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (r *UsersGatewayImpl) GetSuperAdminById(superAdminId uint) (superAdmin *models.SuperAdminCore, err error) {
	var superAdminDb models.SuperAdminDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("id = ?", superAdminId).First(&superAdminDb).Error; err != nil {
			err = auth.ErrUserNotFound
			log.Println(err)
			return
		}
		return
	})
	superAdmin = superAdminDb.ToCore()
	return
}

func (r *UsersGatewayImpl) CreateParent(parent *models.ParentCore) (id string, err error) {
	//students := parent.StudentsID
	parentDb := models.ParentDB{}
	parentDb.FromCore(parent)
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&parentDb).Error
		/*for _, studentId := range students {
			parentStudent := models.ParentStudent{
				ParentID:  parentDb.ID,
				StudentID: studentId,
			}
			err = tx.Create(&parentStudent).Error
		}
		*/
		return
	})
	id = strconv.FormatUint(uint64(parentDb.ID), 10)
	return
}

func (r *UsersGatewayImpl) DeleteParent(parentId uint) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Where("id = ?", parentId).Delete(&models.ParentDB{}).Error
		if err != nil {
			err = auth.ErrUserNotFound
			log.Println(err)
			return
		}
		return
	})
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (r *UsersGatewayImpl) DeleteUnitAdmin(unitAdminId uint) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Where("id = ?", unitAdminId).Delete(&models.UnitAdminDB{}).Error
		if err != nil {
			err = auth.ErrUserNotFound
			log.Println(err)
			return
		}
		return
	})
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (r *UsersGatewayImpl) CreateTeacher(teacher *models.TeacherCore) (id string, err error) {
	/*	groups := teacher.TeachersID
		courses := teacher.CoursesID*/
	teacherDb := models.TeacherDB{}
	teacherDb.FromCore(teacher)
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&teacherDb).Error
		/*for _, groupId := range groups {
			teacherGroup := models.TeacherGroup{
				TeacherID: teacherDb.ID,
				GroupID:   groupId,
			}
			err = tx.Create(&teacherGroup).Error
		}
		for _, courseId := range courses {
			teacherCourse := models.TeacherCourse{
				TeacherID: teacherDb.ID,
				CourseID:  courseId,
			}
			err = tx.Create(&teacherCourse).Error
		}*/
		return
	})
	id = strconv.FormatUint(uint64(teacherDb.ID), 10)
	return
}

func (r *UsersGatewayImpl) UpdateParent(parent *models.ParentCore) (err error) {
	/*
		students := parent.StudentsID
	*/
	parentDb := models.ParentDB{}
	parentDb.FromCore(parent)
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Where("id = ?", parentDb.ID).First(&models.ParentDB{}).Error
		if err != nil {
			log.Println(err)
			return auth.ErrUserNotFound
		}
		err = tx.Model(&parentDb).Where("id = ?", parentDb.ID).Updates(parentDb).Error
		if err != nil {
			log.Println(err)
			return
		}
		/*
			err = tx.Where("parent_id = ?", parentDb.ID).Delete(&models.ParentStudent{}).Error
			for _, studentId := range students {
				parentStudent := models.ParentStudent{ParentID: parentDb.ID,
					StudentID: studentId,
				}
				err = tx.Create(&parentStudent).Error
			}
		*/
		if err != nil {
			log.Println(err)
			return
		}
		return
	})
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (r *UsersGatewayImpl) DeleteTeacher(teacherId uint) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Where("id = ?", teacherId).Delete(&models.TeacherDB{}).Error
		if err != nil {
			err = auth.ErrUserNotFound
			log.Println(err)
			return
		}
		return
	})
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (r *UsersGatewayImpl) GetTeacherById(userId uint) (teacher *models.TeacherCore, err error) {
	var teacherDb models.TeacherDB
	/*var groups []models.TeacherGroup
	var courses []models.TeacherCourse*/
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("id = ?", userId).First(&teacherDb).Error; err != nil {
			err = auth.ErrUserNotFound
			log.Println(err)
			return
		}
		/*if err = tx.Where("teacher_id = ?", teacherDb.ID).Find(&groups).Error; err != nil {
			return
		}
		if err = tx.Where("teacher_id = ?", teacherDb.ID).Find(&courses).Error; err != nil {
			return
		}
		*/
		return
	})
	teacher = teacherDb.ToCore()
	/*for _, group := range groups {
		teacher.GroupsID = append(teacher.GroupsID, group.GroupID)
	}
	for _, course := range courses {
		teacher.CoursesID = append(teacher.CoursesID, course.CourseID)
	}*/
	return teacher, err
}

func (r *UsersGatewayImpl) UpdateTeacher(teacher *models.TeacherCore) (err error) {
	/*
		courses := teacher.CoursesID
		groups := teacher.GroupsID
	*/
	teacherDb := models.TeacherDB{}
	teacherDb.FromCore(teacher)
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Where("id = ?", teacherDb.ID).First(&models.TeacherDB{}).Error
		if err != nil {
			log.Println(err)
			return auth.ErrUserNotFound
		}
		err = tx.Model(&teacherDb).Where("id = ?", teacherDb.ID).Updates(teacherDb).Error
		if err != nil {
			log.Println(err)
			return
		}
		/*
			err = tx.Where("teacher_id = ?", teacherDb.ID).Delete(&models.TeacherCourse{}).Error
			for _, courseId := range courses {
				teacherCourse := models.TeacherCourse{
					TeacherID: teacherDb.ID,
					CourseID:  courseId,
				}
				err = tx.Create(&teacherCourse).Error
			}
			for _, groupId := range groups {
				teacherGroup := models.TeacherGroup{
					TeacherID: teacherDb.ID,
					GroupID:   groupId,
				}
				err = tx.Create(&teacherGroup).Error
			}
		*/
		if err != nil {
			log.Println(err)
			return
		}
		return
	})
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (r *UsersGatewayImpl) GetTeacher(email, password string) (teacher *models.TeacherCore, err error) {
	var teacherDb models.TeacherDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("email = ? AND  password = ?", email, password).First(&teacherDb).Error; err != nil {
			err = auth.ErrUserNotFound
			log.Println(err)
			return
		}
		return
	})
	teacher = teacherDb.ToCore()
	return teacher, err
}

func (r *UsersGatewayImpl) GetParent(email, password string) (parent *models.ParentCore, err error) {
	var parentDb models.ParentDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("email = ? AND  password = ?", email, password).First(&parentDb).Error; err != nil {
			err = auth.ErrUserNotFound
			log.Println(err)
			return
		}
		return
	})
	parent = parentDb.ToCore()
	return parent, err
}

func (r *UsersGatewayImpl) GetUnitAdmin(email, password string) (unitAdmin *models.UnitAdminCore, err error) {
	var unitAdminDb models.UnitAdminDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("email = ? AND  password = ?", email, password).First(&unitAdminDb).Error; err != nil {
			err = auth.ErrUserNotFound
			log.Println(err)
			return
		}
		return
	})
	unitAdmin = unitAdminDb.ToCore()
	return unitAdmin, err
}

func (r *UsersGatewayImpl) GetSuperAdmin(email, password string) (superAdmin *models.SuperAdminCore, err error) {
	var superAdminDb models.SuperAdminDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("email = ? AND  password = ?", email, password).First(&superAdminDb).Error; err != nil {
			err = auth.ErrUserNotFound
			log.Println(err)
			return
		}
		return
	})
	superAdmin = superAdminDb.ToCore()
	return superAdmin, err
}

func (r *UsersGatewayImpl) GetStudent(email, password string) (student *models.StudentCore, err error) {
	var studentDb models.StudentDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("email = ? AND  password = ?", email, password).First(&studentDb).Error; err != nil {
			err = auth.ErrUserNotFound
			log.Println(err)
			return
		}
		return
	})
	student = studentDb.ToCore()
	return student, err
}

func (r *UsersGatewayImpl) CreateStudent(student *models.StudentCore) (id string, err error) {
	/*groups := student.GroupsID
	teachers := student.TeachersID
	projects := student.ProjectsID*/
	studentDb := models.StudentDB{}
	studentDb.FromCore(student)
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&studentDb).Error
		/*for _, groupId := range groups {
			studentGroup := models.StudentGroup{StudentID: studentDb.ID,
				GroupID: groupId,
			}
			err = tx.Create(&studentGroup).Error
		}
		for _, teacherId := range teachers {
			studentTeacher := models.StudentTeacher{StudentID: studentDb.ID,
				TeacherID: teacherId,
			}
			err = tx.Create(&studentTeacher).Error
		}
		for _, projectId := range projects {
			teacherCourse := models.StudentProject{StudentID: studentDb.ID,
				ProjectID: projectId,
			}
			err = tx.Create(&teacherCourse).Error
		}*/
		return
	})
	id = strconv.FormatUint(uint64(studentDb.ID), 10)
	return
}

func (r *UsersGatewayImpl) DeleteStudent(studentId uint) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Where("id = ?", studentId).Delete(&models.StudentDB{}).Error
		if err != nil {
			log.Println(err)
			return
		}
		return
	})
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (r *UsersGatewayImpl) GetStudentById(studentId uint) (student *models.StudentCore, err error) {
	var studentDb models.StudentDB
	/*var teachers []models.StudentTeacher
	var projects []models.StudentProject
	var groups []models.StudentGroup*/
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("id = ?", studentId).First(&studentDb).Error; err != nil {
			err = auth.ErrUserNotFound
			log.Println(err)
			return
		}
		/*if err = tx.Where("student_id = ?", studentDb.ID).Find(&groups).Error; err != nil {
			return
		}
		if err = tx.Where("student_id = ?", studentDb.ID).Find(&teachers).Error; err != nil {
			return
		}
		if err = tx.Where("student_id = ?", studentDb.ID).Find(&projects).Error; err != nil {
			return
		}

		*/
		return
	})
	student = studentDb.ToCore()
	/*for _, group := range groups {
		student.GroupsID = append(student.GroupsID, group.GroupID)
	}
	for _, project := range projects {
		student.ProjectsID = append(student.ProjectsID, project.ProjectID)
	}
	for _, teacher := range teachers {
		student.TeachersID = append(student.TeachersID, teacher.TeacherID)
	}*/
	return student, err
}

func (r *UsersGatewayImpl) UpdateStudent(student *models.StudentCore) (err error) {
	/*groups := student.GroupsID
	projects := student.ProjectsID
	teachers := student.TeachersID*/
	studentDb := models.StudentDB{}
	studentDb.FromCore(student)
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Where("id = ?", student.ID).First(&models.StudentDB{}).Error
		if err != nil {
			log.Println(err)
			return auth.ErrUserNotFound
		}
		err = tx.Model(&studentDb).Where("id = ?", studentDb.ID).Updates(studentDb).Error
		if err != nil {
			log.Println(err)
			return
		}
		/*
			err = tx.Where("student_id = ?", studentDb.ID).Delete(&models.StudentTeacher{}).Error
			for _, teacherId := range teachers {
				teacherCourse := models.StudentTeacher{StudentID: studentDb.ID,
					TeacherID: teacherId,
				}
				err = tx.Create(&teacherCourse).Error
			}

			err = tx.Where("student_id = ?", studentDb.ID).Delete(&models.StudentProject{}).Error
			for _, projectId := range projects {
				teacherCourse := models.StudentProject{StudentID: studentDb.ID,
					ProjectID: projectId,
				}
				err = tx.Create(&teacherCourse).Error
			}

			err = tx.Where("student_id = ?", studentDb.ID).Delete(&models.StudentGroup{}).Error
			for _, groupId := range groups {
				studentGroup := models.StudentGroup{StudentID: studentDb.ID,
					GroupID: groupId,
				}
				err = tx.Create(&studentGroup).Error
			}
		*/
		if err != nil {
			log.Println(err)
			return
		}
		return
	})
	if err != nil {
		log.Println(err)
		return
	}
	return
}
