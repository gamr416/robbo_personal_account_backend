package gateway

import (
	"fmt"
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

func (r *UsersGatewayImpl) SearchStudentByEmail(email string) (students []*models.StudentCore, err error) {
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
		students = append(students, studentDb.ToCore())
	}
	return
}

func (r *UsersGatewayImpl) CreateStudent(student *models.StudentCore) (id string, err error) {
	studentDb := models.StudentDB{}
	studentDb.FromCore(student)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&studentDb).Error
		return
	})

	id = strconv.FormatUint(uint64(studentDb.ID), 10)
	return
}

func (r *UsersGatewayImpl) DeleteStudent(studentId uint) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Delete(&models.StudentDB{}, studentId).Error
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

func (r *UsersGatewayImpl) UpdateStudent(student *models.StudentCore) (err error) {
	studentDb := models.StudentDB{}
	studentDb.FromCore(student)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Model(&studentDb).Where("id = ?", studentDb.ID).Updates(studentDb).Error
		return
	})
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

func (r *UsersGatewayImpl) GetAllTeachers() (teachers []*models.TeacherCore, err error) {
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

func (r *UsersGatewayImpl) GetTeacherById(userId string) (teacher *models.TeacherCore, err error) {
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

func (r *UsersGatewayImpl) CreateTeacher(teacher *models.TeacherCore) (id string, err error) {
	teacherDb := models.TeacherDB{}
	teacherDb.FromCore(teacher)
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&teacherDb).Error
		return
	})
	id = strconv.FormatUint(uint64(teacherDb.ID), 10)
	return
}

func (r *UsersGatewayImpl) DeleteTeacher(teacherId uint) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Delete(&models.TeacherDB{}, teacherId).Error
		return
	})
	return
}

func (r *UsersGatewayImpl) UpdateTeacher(teacher *models.TeacherCore) (err error) {
	teacherDb := models.TeacherDB{}
	teacherDb.FromCore(teacher)
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
			err = tx.Model(&teacherDb).Where("id = ?", teacherDb.ID).Updates(teacherDb).Error
			return
		})
		return
	})

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

func (r *UsersGatewayImpl) CreateParent(parent *models.ParentCore) (id string, err error) {
	parentDb := models.ParentDB{}
	parentDb.FromCore(parent)
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&parentDb).Error
		return
	})
	id = strconv.FormatUint(uint64(parentDb.ID), 10)
	return
}

func (r *UsersGatewayImpl) DeleteParent(parentId uint) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Delete(&models.ParentDB{}, parentId).Error
		return
	})
	return
}

func (r *UsersGatewayImpl) UpdateParent(parent *models.ParentCore) (err error) {
	parentDb := models.ParentDB{}
	parentDb.FromCore(parent)
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
			err = tx.Model(&parentDb).Where("id = ?", parentDb.ID).Updates(parentDb).Error
			return
		})
		return
	})

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

func (r *UsersGatewayImpl) CreateFreeListener(freeListener *models.FreeListenerCore) (id string, err error) {
	freeListenerDb := models.FreeListenerDB{}
	freeListenerDb.FromCore(freeListener)
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&freeListenerDb).Error
		return
	})

	id = strconv.FormatUint(uint64(freeListenerDb.ID), 10)
	return id, nil
}

func (r *UsersGatewayImpl) DeleteFreeListener(freeListenerId uint) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Delete(&models.FreeListenerDB{}, freeListenerId).Error
		return
	})
	return
}

func (r *UsersGatewayImpl) UpdateFreeListener(freeListener *models.FreeListenerCore) (err error) {
	freeListenerDb := models.FreeListenerDB{}
	freeListenerDb.FromCore(freeListener)
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
			err = tx.Model(&freeListenerDb).Where("id = ?", freeListenerDb.ID).Updates(freeListenerDb).Error
			return
		})
		return
	})
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

func (r *UsersGatewayImpl) CreateUnitAdmin(unitAdmin *models.UnitAdminCore) (id string, err error) {
	unitAdminDb := models.UnitAdminDB{}
	unitAdminDb.FromCore(unitAdmin)
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&unitAdminDb).Error
		return
	})

	id = strconv.FormatUint(uint64(unitAdminDb.ID), 10)
	return id, nil
}

func (r *UsersGatewayImpl) DeleteUnitAdmin(unitAdminId uint) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Delete(&models.UnitAdminDB{}, unitAdminId).Error
		return
	})
	return
}

func (r *UsersGatewayImpl) UpdateUnitAdmin(unitAdmin *models.UnitAdminCore) (err error) {
	unitAdminDb := models.UnitAdminDB{}
	unitAdminDb.FromCore(unitAdmin)
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
			err = tx.Model(&unitAdminDb).Where("id = ?", unitAdminDb.ID).Updates(unitAdminDb).Error
			return
		})
		return
	})
	return
}

func (r *UsersGatewayImpl) SearchUnitAdminByEmail(email string) (unitAdmins []*models.UnitAdminCore, err error) {
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

func (r *UsersGatewayImpl) UpdateSuperAdmin(superAdmin *models.SuperAdminCore) (err error) {
	superAdminDb := models.SuperAdminDB{}
	superAdminDb.FromCore(superAdmin)
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
			err = tx.Model(&superAdminDb).Where("id = ?", superAdminDb.ID).Updates(superAdminDb).Error
			return
		})
		return
	})
	return
}

func (r *UsersGatewayImpl) DeleteSuperAdmin(superAdminId uint) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Delete(&models.SuperAdminDB{}, superAdminId).Error
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
