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
	fmt.Println(student)
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

func (r *UsersGatewayImpl) GetStudentById(studentId uint) (student *models.StudentCore, err error) {
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

func (r *UsersGatewayImpl) GetTeacherById(userId uint) (teacher *models.TeacherCore, err error) {
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

func (r *UsersGatewayImpl) GetFreeListenerById(freeListenerId uint) (freeListener *models.FreeListenerCore, err error) {
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
