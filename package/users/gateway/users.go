package gateway

import (
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/db_client"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/users"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"log"
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

func (r *UsersGatewayImpl) GetStudentsByRobboUnitId(robboUnitId string) (students []*models.StudentCore, err error) {
	var studentsDb []*models.StudentDB

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("robbo_unit_id = ?", robboUnitId).Find(&studentsDb).Error; err != nil {
			return
		}
		return
	})

	for _, studentDb := range studentsDb {
		students = append(students, studentDb.ToCore())
	}
	return
}

func (r *UsersGatewayImpl) AddStudentToRobboGroup(studentId, robboGroupId, robboUnitId string) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Model(&models.StudentDB{}).Where("id = ?", studentId).
			Update("robbo_group_id", gorm.Expr(robboGroupId)).
			Update("robbo_unit_id", gorm.Expr(robboUnitId)).Error
		return
	})
	return
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
	return
}

func (r *UsersGatewayImpl) SearchStudentsByEmail(email string, parentId string) (students []*models.StudentCore, err error) {
	var studentsDb []*models.StudentDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Limit(10).Where("email LIKE ?", email).Find(&studentsDb).Error; err != nil {
			err = auth.ErrUserNotFound
			log.Println(err)
			return
		}
		return
	})

	//student = studentsDb.ToCore()
	for _, studentDb := range studentsDb {
		var countRow int64
		err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
			tx.Where("parent_id = ? AND child_id = ?", parentId, studentDb.ID).Model(&models.ChildrenOfParentDB{}).Count(&countRow)
			return
		})
		if countRow != 0 {
			continue
		}

		var countParents int64
		err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
			tx.Where("child_id = ?", studentDb.ID).Model(&models.ChildrenOfParentDB{}).Count(&countParents)
			return
		})
		if countParents >= 2 {
			continue
		}

		students = append(students, studentDb.ToCore())
	}
	return
}

func (r *UsersGatewayImpl) CreateStudent(student *models.StudentCore) (newStudent *models.StudentCore, err error) {
	studentDb := models.StudentDB{}
	studentDb.FromCore(student)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&studentDb).Error
		var duplicateEntryError = &pgconn.PgError{Code: "23505"}
		if errors.As(err, &duplicateEntryError) {
			return auth.ErrUserAlreadyExist
		}
		return
	})

	newStudent = studentDb.ToCore()
	return
}

func (r *UsersGatewayImpl) DeleteStudent(studentId string) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Model(&models.StudentDB{}).Where("id = ?", studentId).First(&models.StudentDB{}).Delete(&models.StudentDB{}).Error; err != nil {
			err = auth.ErrUserNotFound
			return
		}
		return
	})
	return
}

func (r *UsersGatewayImpl) GetStudentById(studentId string) (student *models.StudentCore, err error) {
	var studentDb models.StudentDB

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("id = ?", studentId).First(&studentDb).Error; err != nil {
			err = auth.ErrUserNotFound
			log.Println(err)
			return
		}
		return
	})
	student = studentDb.ToCore()

	return
}

func (r *UsersGatewayImpl) GetStudentsByRobboGroupId(robboGroupId string) (students []*models.StudentCore, err error) {
	var studentsDb []*models.StudentDB

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("robbo_group_id = ?", robboGroupId).Find(&studentsDb).Error; err != nil {
			return
		}
		return
	})

	for _, studentDb := range studentsDb {
		students = append(students, studentDb.ToCore())
	}
	return
}

func (r *UsersGatewayImpl) UpdateStudent(student *models.StudentCore) (studentUpdated *models.StudentCore, err error) {
	studentDb := models.StudentDB{}
	studentDb.FromCore(student)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Model(&studentDb).Where("id = ?", studentDb.ID).First(&models.StudentDB{}).Updates(&studentDb).Error; err != nil {
			err = auth.ErrUserNotFound
			return
		}
		return
	})
	fmt.Println(studentDb)
	studentUpdated = studentDb.ToCore()
	return
}

func (r *UsersGatewayImpl) GetTeacher(email, password string) (teacher models.TeacherCore, err error) {
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

func (r *UsersGatewayImpl) GetAllTeachers() (teachers []models.TeacherCore, err error) {
	var teachersDB []*models.TeacherDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Find(&teachersDB).Error; err != nil {
			return
		}
		return
	})

	for _, teacherDb := range teachersDB {
		teachers = append(teachers, teacherDb.ToCore())
	}
	return
}

func (r *UsersGatewayImpl) GetTeacherById(userId string) (teacher models.TeacherCore, err error) {
	var teacherDb models.TeacherDB

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("id = ?", userId).First(&teacherDb).Error; err != nil {
			err = auth.ErrUserNotFound
			log.Println(err)
			return
		}
		return
	})
	teacher = teacherDb.ToCore()
	return teacher, err
}

func (r *UsersGatewayImpl) CreateTeacher(teacher *models.TeacherCore) (newTeacher models.TeacherCore, err error) {
	teacherDb := models.TeacherDB{}
	teacherDb.FromCore(teacher)
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&teacherDb).Error
		var duplicateEntryError = &pgconn.PgError{Code: "23505"}
		if errors.As(err, &duplicateEntryError) {
			return auth.ErrUserAlreadyExist
		}
		return
	})
	newTeacher = teacherDb.ToCore()
	return
}

func (r *UsersGatewayImpl) DeleteTeacher(teacherId string) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Model(&models.TeacherDB{}).Where("id = ?", teacherId).First(&models.TeacherDB{}).Delete(&models.TeacherDB{}).Error; err != nil {
			err = auth.ErrUserNotFound
			return
		}
		return
	})
	return
}

func (r *UsersGatewayImpl) UpdateTeacher(teacher *models.TeacherCore) (teacherUpdated models.TeacherCore, err error) {
	teacherDb := models.TeacherDB{}
	teacherDb.FromCore(teacher)
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
			if err = tx.Model(&teacherDb).Where("id = ?", teacherDb.ID).First(&models.TeacherDB{}).Updates(teacherDb).Error; err != nil {
				err = auth.ErrUserNotFound
				return
			}
			return
		})
		return
	})
	teacherUpdated = teacherDb.ToCore()
	return
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

func (r *UsersGatewayImpl) GetAllParent() (parents []*models.ParentCore, err error) {
	var parentsDB []*models.ParentDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Find(&parentsDB).Error; err != nil {
			return
		}
		return
	})
	fmt.Println(parentsDB)

	for _, parentDb := range parentsDB {
		parents = append(parents, parentDb.ToCore())
	}
	return
}

func (r *UsersGatewayImpl) GetParentById(parentId string) (parent *models.ParentCore, err error) {
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

func (r *UsersGatewayImpl) CreateParent(parent *models.ParentCore) (newParent *models.ParentCore, err error) {
	parentDb := models.ParentDB{}
	parentDb.FromCore(parent)
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&parentDb).Error
		var duplicateEntryError = &pgconn.PgError{Code: "23505"}
		if errors.As(err, &duplicateEntryError) {
			return auth.ErrUserAlreadyExist
		}
		return
	})
	newParent = parentDb.ToCore()
	return
}

func (r *UsersGatewayImpl) DeleteParent(parentId string) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Model(&models.ParentDB{}).Where("id = ?", parentId).First(&models.ParentDB{}).Delete(&models.ParentDB{}).Error; err != nil {
			err = auth.ErrUserNotFound
			return
		}
		return
	})
	return
}

func (r *UsersGatewayImpl) UpdateParent(parent *models.ParentCore) (parentUpdated *models.ParentCore, err error) {
	parentDb := models.ParentDB{}
	parentDb.FromCore(parent)
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
			if err = tx.Model(&parentDb).Where("id = ?", parentDb.ID).First(&models.ParentDB{}).Updates(parentDb).Error; err != nil {
				err = auth.ErrUserNotFound
				return
			}
			return
		})
		return
	})
	parentUpdated = parentDb.ToCore()
	return
}

func (r *UsersGatewayImpl) GetFreeListener(email, password string) (freeListener *models.FreeListenerCore, err error) {
	var freeListenerDb models.FreeListenerDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("email = ? AND  password = ?", email, password).First(&freeListenerDb).Error; err != nil {
			err = auth.ErrUserNotFound
			log.Println(err)
			return
		}
		return
	})
	freeListener = freeListenerDb.ToCore()
	return
}

func (r *UsersGatewayImpl) GetFreeListenerById(freeListenerId string) (freeListener *models.FreeListenerCore, err error) {
	var freeListenerDb models.FreeListenerDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("id = ?", freeListenerId).First(&freeListenerDb).Error; err != nil {
			err = auth.ErrUserNotFound
			log.Println(err)
			return
		}
		return
	})
	freeListener = freeListenerDb.ToCore()
	return
}

func (r *UsersGatewayImpl) CreateFreeListener(freeListener *models.FreeListenerCore) (newFreeListener *models.FreeListenerCore, err error) {
	freeListenerDb := models.FreeListenerDB{}
	freeListenerDb.FromCore(freeListener)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&freeListenerDb).Error
		var duplicateEntryError = &pgconn.PgError{Code: "23505"}
		if errors.As(err, &duplicateEntryError) {
			return auth.ErrUserAlreadyExist
		}
		return
	})

	newFreeListener = freeListenerDb.ToCore()
	return
}

func (r *UsersGatewayImpl) DeleteFreeListener(freeListenerId string) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Model(&models.FreeListenerDB{}).Where("id = ?", freeListenerId).First(&models.FreeListenerDB{}).Delete(&models.FreeListenerDB{}).Error; err != nil {
			err = auth.ErrUserNotFound
			return
		}
		return
	})
	return
}

func (r *UsersGatewayImpl) UpdateFreeListener(freeListener *models.FreeListenerCore) (freeListenerUpdated *models.FreeListenerCore, err error) {
	freeListenerDb := models.FreeListenerDB{}
	freeListenerDb.FromCore(freeListener)
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
			if err = tx.Model(&freeListenerDb).Where("id = ?", freeListenerDb.ID).First(&models.FreeListenerDB{}).Updates(freeListenerDb).Error; err != nil {
				err = auth.ErrUserNotFound
				return
			}
			return
		})
		return
	})
	freeListenerUpdated = freeListenerDb.ToCore()
	return
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
	return
}

func (r *UsersGatewayImpl) GetUnitAdminById(unitAdminId string) (unitAdmin *models.UnitAdminCore, err error) {
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

func (r *UsersGatewayImpl) GetAllUnitAdmins() (unitAdmins []*models.UnitAdminCore, err error) {
	var unitAdminsDB []*models.UnitAdminDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Find(&unitAdminsDB).Error; err != nil {
			return
		}
		return
	})

	for _, unitAdminDb := range unitAdminsDB {
		unitAdmins = append(unitAdmins, unitAdminDb.ToCore())
	}
	return
}

func (r *UsersGatewayImpl) CreateUnitAdmin(unitAdmin *models.UnitAdminCore) (newUnitAdmin *models.UnitAdminCore, err error) {
	unitAdminDb := models.UnitAdminDB{}
	unitAdminDb.FromCore(unitAdmin)
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&unitAdminDb).Error
		var duplicateEntryError = &pgconn.PgError{Code: "23505"}
		if errors.As(err, &duplicateEntryError) {
			return auth.ErrUserAlreadyExist
		}
		return
	})

	newUnitAdmin = unitAdminDb.ToCore()
	return
}

func (r *UsersGatewayImpl) DeleteUnitAdmin(unitAdminId string) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Model(&models.UnitAdminDB{}).Where("id = ?", unitAdminId).First(&models.UnitAdminDB{}).Delete(&models.UnitAdminDB{}).Error; err != nil {
			err = auth.ErrUserNotFound
			return
		}
		return
	})
	return
}

func (r *UsersGatewayImpl) UpdateUnitAdmin(unitAdmin *models.UnitAdminCore) (unitAdminUpdated *models.UnitAdminCore, err error) {
	unitAdminDb := models.UnitAdminDB{}
	unitAdminDb.FromCore(unitAdmin)
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
			if err = tx.Model(&unitAdminDb).Where("id = ?", unitAdminDb.ID).First(&models.UnitAdminDB{}).Updates(unitAdminDb).Error; err != nil {
				err = auth.ErrUserNotFound
				return
			}
			return
		})
		return
	})
	unitAdminUpdated = unitAdminDb.ToCore()
	return
}

func (r *UsersGatewayImpl) SearchUnitAdminByEmail(email string, robboUnitId string) (unitAdmins []*models.UnitAdminCore, err error) {
	var unitAdminsDb []*models.UnitAdminDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Limit(10).Where("email LIKE ?", email).Find(&unitAdminsDb).Error; err != nil {
			err = auth.ErrUserNotFound
			log.Println(err)
			return
		}
		return
	})
	for _, unitAdminDb := range unitAdminsDb {
		var countRow int64
		err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
			tx.Where("robbo_unit_id = ? AND unit_admin_id = ?", robboUnitId, unitAdminDb.ID).Model(&models.UnitAdminsRobboUnitsDB{}).Count(&countRow)
			return
		})
		if countRow != 0 {
			continue
		}
		unitAdmins = append(unitAdmins, unitAdminDb.ToCore())
	}
	return
}

func (r *UsersGatewayImpl) GetSuperAdminById(superAdminId string) (superAdmin *models.SuperAdminCore, err error) {
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

func (r *UsersGatewayImpl) GetSuperAdmin(email, password string) (superAdmin *models.SuperAdminCore, err error) {
	var superAdminDb models.SuperAdminDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("email = ? AND  password = ?", email, password).First(&superAdminDb).Error; err != nil {
			err = auth.ErrUserNotFound
			return
		}
		return
	})
	superAdmin = superAdminDb.ToCore()
	return
}

func (r *UsersGatewayImpl) UpdateSuperAdmin(superAdmin *models.SuperAdminCore) (superAdminUpdated *models.SuperAdminCore, err error) {
	superAdminDb := models.SuperAdminDB{}
	superAdminDb.FromCore(superAdmin)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
			if err = tx.Model(&superAdminDb).Where("id = ?", superAdminDb.ID).First(&models.SuperAdminDB{}).Updates(superAdminDb).Error; err != nil {
				err = auth.ErrUserNotFound
				return
			}
			return
		})
		return
	})
	superAdminUpdated = superAdminDb.ToCore()
	return
}

func (r *UsersGatewayImpl) DeleteSuperAdmin(superAdminId string) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Model(&models.SuperAdminDB{}).Where("id = ?", superAdminId).First(&models.SuperAdminDB{}).Delete(&models.SuperAdminDB{}).Error; err != nil {
			err = auth.ErrUserNotFound
			return
		}
		return
	})
	return
}

func (r *UsersGatewayImpl) CreateRelation(relation *models.ChildrenOfParentCore) (err error) {
	relationDb := models.ChildrenOfParentDB{}
	relationDb.FromCore(relation)
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&relationDb).Error
		return
	})

	return
}

func (r *UsersGatewayImpl) DeleteRelationByParentId(parentId string) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Where("parent_id = ?", parentId).Delete(&models.ChildrenOfParentDB{}).Error
		return
	})
	return
}

func (r *UsersGatewayImpl) DeleteRelationByChildrenId(childrenId string) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Where("child_id = ?", childrenId).Delete(&models.ChildrenOfParentDB{}).Error
		return
	})
	return
}

func (r *UsersGatewayImpl) DeleteRelation(relation *models.ChildrenOfParentCore) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Delete(&models.ChildrenOfParentDB{}, relation).Error
		return
	})
	return
}

func (r *UsersGatewayImpl) GetRelationByParentId(parentId string) (relations []*models.ChildrenOfParentCore, err error) {
	var relationsDB []*models.ChildrenOfParentDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("parent_id = ?", parentId).Find(&relationsDB).Error; err != nil {
			return
		}
		return
	})

	for _, relationDB := range relationsDB {
		relations = append(relations, relationDB.ToCore())
	}
	return
}

func (r *UsersGatewayImpl) GetRelationByChildrenId(childrenId string) (relations []*models.ChildrenOfParentCore, err error) {
	var relationsDB []*models.ChildrenOfParentDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("children_id = ?", childrenId).Find(&relationsDB).Error; err != nil {
			return
		}
		return
	})

	for _, relationDB := range relationsDB {
		relations = append(relations, relationDB.ToCore())
	}
	return
}

func (r *UsersGatewayImpl) SetUnitAdminForRobboUnit(relation *models.UnitAdminsRobboUnitsCore) (err error) {
	relationDb := models.UnitAdminsRobboUnitsDB{}
	relationDb.FromCore(relation)
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&relationDb).Error
		return
	})

	return
}

func (r *UsersGatewayImpl) DeleteRelationByRobboUnitId(robboUnitId string) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Where("robbo_unit_id = ?", robboUnitId).Delete(&models.UnitAdminsRobboUnitsDB{}).Error
		return
	})
	return
}

func (r *UsersGatewayImpl) DeleteRelationByUnitAdminId(unitAdminId string) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Where("unit_admin_id = ?", unitAdminId).Delete(&models.UnitAdminsRobboUnitsDB{}).Error
		return
	})
	return
}

func (r *UsersGatewayImpl) GetRelationByRobboUnitId(robboUnitId string) (relations []*models.UnitAdminsRobboUnitsCore, err error) {
	var relationsDB []*models.UnitAdminsRobboUnitsDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("robbo_unit_id = ?", robboUnitId).Find(&relationsDB).Error; err != nil {
			return
		}
		return
	})

	for _, relationDB := range relationsDB {
		relations = append(relations, relationDB.ToCore())
	}
	return
}

func (r *UsersGatewayImpl) GetRelationByUnitAdminId(unitAdminId string) (relations []*models.UnitAdminsRobboUnitsCore, err error) {
	var relationsDB []*models.UnitAdminsRobboUnitsDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("unit_admin_id = ?", unitAdminId).Find(&relationsDB).Error; err != nil {
			return
		}
		return
	})

	for _, relationDB := range relationsDB {
		relations = append(relations, relationDB.ToCore())
	}
	return
}

func (r *UsersGatewayImpl) DeleteUnitAdminForRobboUnit(relation *models.UnitAdminsRobboUnitsCore) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Delete(&models.UnitAdminsRobboUnitsDB{}, relation).Error
		return
	})
	return
}

func (r *UsersGatewayImpl) CreateStudentTeacherRelation(relation *models.StudentsOfTeacherCore) (err error) {
	relationDb := models.StudentsOfTeacherDB{}
	relationDb.FromCore(relation)
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&relationDb).Error
		return
	})
	return
}

func (r *UsersGatewayImpl) DeleteStudentTeacherRelation(relation *models.StudentsOfTeacherCore) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Delete(&models.StudentsOfTeacherDB{}, relation).Error
		return
	})
	return
}

func (r *UsersGatewayImpl) DeleteStudentTeacherRelationByTeacherId(teacherId string) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Where("teacher_id = ?", teacherId).Delete(&models.StudentsOfTeacherDB{}).Error
		return
	})
	return
}

func (r *UsersGatewayImpl) DeleteStudentTeacherRelationByStudentId(studentId string) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Where("student_id = ?", studentId).Delete(&models.StudentsOfTeacherDB{}).Error
		return
	})
	return
}

func (r *UsersGatewayImpl) GetStudentTeacherRelationsByTeacherId(teacherId string) (relations []*models.StudentsOfTeacherCore, err error) {
	var relationsDB []*models.StudentsOfTeacherDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("teacher_id = ?", teacherId).Find(&relationsDB).Error; err != nil {
			return
		}
		return
	})

	for _, relationDB := range relationsDB {
		relations = append(relations, relationDB.ToCore())
	}
	return
}

func (r *UsersGatewayImpl) GetStudentTeacherRelationsByStudentId(studentId string) (relations []*models.StudentsOfTeacherCore, err error) {
	var relationsDB []*models.StudentsOfTeacherDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("student_id = ?", studentId).Find(&relationsDB).Error; err != nil {
			return
		}
		return
	})

	for _, relationDB := range relationsDB {
		relations = append(relations, relationDB.ToCore())
	}
	return
}
