package gateway

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/courses"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/db_client"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"log"
	"strconv"
)

type CoursesGatewayImpl struct {
	PostgresClient *db_client.PostgresClient
}

type CoursesGatewayModule struct {
	fx.Out
	courses.Gateway
}

func SetupCoursesGateway(postgresClient db_client.PostgresClient) CoursesGatewayModule {
	return CoursesGatewayModule{
		Gateway: &CoursesGatewayImpl{PostgresClient: &postgresClient},
	}
}

func (r *CoursesGatewayImpl) DeleteAccessCourseRelationsByRobboUnitId(robboUnitId string) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Where("parameter = ? AND object_id", "robbo_unit", robboUnitId).Delete(&models.CourseRelationDB{}).Error
		return
	})
	return
}

func (r *CoursesGatewayImpl) DeleteAccessCourseRelationsByRobboGroupId(robboGroupId string) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Where("parameter = ? AND object_id", "robbo_group", robboGroupId).Delete(&models.CourseRelationDB{}).Error
		return
	})
	return
}

func (r *CoursesGatewayImpl) DeleteAccessCourseRelationsByStudentId(studentId string) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Where("parameter = ? AND object_id", "student", studentId).Delete(&models.CourseRelationDB{}).Error
		return
	})
	return
}

func (r *CoursesGatewayImpl) DeleteAccessCourseRelationsByTeacherId(teacherId string) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Where("parameter = ? AND object_id", "teacher", teacherId).Delete(&models.CourseRelationDB{}).Error
		return
	})
	return
}

func (r *CoursesGatewayImpl) DeleteAccessCourseRelationsByUnitAdminId(unitAdminId string) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Where("parameter = ? AND object_id", "unit_admin", unitAdminId).Delete(&models.CourseRelationDB{}).Error
		return
	})
	return
}

func (r *CoursesGatewayImpl) GetAccessCourseRelationsByCourseId(courseId string) (courseRelations []*models.CourseRelationCore, err error) {
	var courseRelationsDb []*models.CourseRelationDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("course_id", courseId).Find(&courseRelationsDb).Error; err != nil {
			return
		}
		return
	})

	for _, courseRelationDb := range courseRelationsDb {
		var courseRelationCore *models.CourseRelationCore

		courseRelationCore = courseRelationDb.ToCore()
		courseRelations = append(courseRelations, courseRelationCore)
	}
	return
}

func (r *CoursesGatewayImpl) GetAccessCourseRelationsRobboGroups() (courseRelations []*models.CourseRelationCore, err error) {
	var courseRelationsDb []*models.CourseRelationDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("parameter = ?", "robbo_group").Find(&courseRelationsDb).Error; err != nil {
			return
		}
		return
	})

	for _, courseRelationDb := range courseRelationsDb {
		var courseRelationCore *models.CourseRelationCore

		courseRelationCore = courseRelationDb.ToCore()
		courseRelations = append(courseRelations, courseRelationCore)
	}
	return
}

func (r *CoursesGatewayImpl) GetAccessCourseRelationsRobboUnits() (courseRelations []*models.CourseRelationCore, err error) {
	var courseRelationsDb []*models.CourseRelationDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("parameter = ?", "robbo_unit").Find(&courseRelationsDb).Error; err != nil {
			return
		}
		return
	})

	for _, courseRelationDb := range courseRelationsDb {
		var courseRelationCore *models.CourseRelationCore

		courseRelationCore = courseRelationDb.ToCore()
		courseRelations = append(courseRelations, courseRelationCore)
	}
	return
}

func (r *CoursesGatewayImpl) GetAccessCourseRelationsStudents() (courseRelations []*models.CourseRelationCore, err error) {
	var courseRelationsDb []*models.CourseRelationDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("parameter = ?", "student").Find(&courseRelationsDb).Error; err != nil {
			return
		}
		return
	})

	for _, courseRelationDb := range courseRelationsDb {
		var courseRelationCore *models.CourseRelationCore

		courseRelationCore = courseRelationDb.ToCore()
		courseRelations = append(courseRelations, courseRelationCore)
	}
	return
}

func (r *CoursesGatewayImpl) GetAccessCourseRelationsTeachers() (courseRelations []*models.CourseRelationCore, err error) {
	var courseRelationsDb []*models.CourseRelationDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("parameter = ?", "teacher").Find(&courseRelationsDb).Error; err != nil {
			return
		}
		return
	})

	for _, courseRelationDb := range courseRelationsDb {
		var courseRelationCore *models.CourseRelationCore

		courseRelationCore = courseRelationDb.ToCore()
		courseRelations = append(courseRelations, courseRelationCore)
	}
	return
}

func (r *CoursesGatewayImpl) GetAccessCourseRelationsUnitAdmins() (courseRelations []*models.CourseRelationCore, err error) {
	var courseRelationsDb []*models.CourseRelationDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("parameter = ?", "unit_admin").Find(&courseRelationsDb).Error; err != nil {
			return
		}
		return
	})

	for _, courseRelationDb := range courseRelationsDb {
		var courseRelationCore *models.CourseRelationCore

		courseRelationCore = courseRelationDb.ToCore()
		courseRelations = append(courseRelations, courseRelationCore)
	}
	return
}

func (r *CoursesGatewayImpl) GetAccessCourseRelationsByRobboUnitId(robboUnitId string) (courseRelations []*models.CourseRelationCore, err error) {
	var courseRelationsDb []*models.CourseRelationDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("parameter = ? AND object_id = ?", "robbo_unit", robboUnitId).Find(&courseRelationsDb).Error; err != nil {
			return
		}
		return
	})

	for _, courseRelationDb := range courseRelationsDb {
		var courseRelationCore *models.CourseRelationCore

		courseRelationCore = courseRelationDb.ToCore()
		courseRelations = append(courseRelations, courseRelationCore)
	}
	return
}

func (r *CoursesGatewayImpl) GetAccessCourseRelationsByRobboGroupId(robboGroupId string) (courseRelations []*models.CourseRelationCore, err error) {
	var courseRelationsDb []*models.CourseRelationDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("parameter = ? AND object_id = ?", "robbo_group", robboGroupId).Find(&courseRelationsDb).Error; err != nil {
			return
		}
		return
	})

	for _, courseRelationDb := range courseRelationsDb {
		var courseRelationCore *models.CourseRelationCore

		courseRelationCore = courseRelationDb.ToCore()
		courseRelations = append(courseRelations, courseRelationCore)
	}
	return
}

func (r *CoursesGatewayImpl) GetAccessCourseRelationsByStudentId(studentId string) (courseRelations []*models.CourseRelationCore, err error) {
	var courseRelationsDb []*models.CourseRelationDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("parameter = ? AND object_id = ?", "student", studentId).Find(&courseRelationsDb).Error; err != nil {
			return
		}
		return
	})

	for _, courseRelationDb := range courseRelationsDb {
		var courseRelationCore *models.CourseRelationCore

		courseRelationCore = courseRelationDb.ToCore()
		courseRelations = append(courseRelations, courseRelationCore)
	}
	return
}

func (r *CoursesGatewayImpl) GetAccessCourseRelationsByTeacherId(teacherId string) (courseRelations []*models.CourseRelationCore, err error) {
	var courseRelationsDb []*models.CourseRelationDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("parameter = ? AND object_id = ?", "teacher", teacherId).Find(&courseRelationsDb).Error; err != nil {
			return
		}
		return
	})

	for _, courseRelationDb := range courseRelationsDb {
		var courseRelationCore *models.CourseRelationCore

		courseRelationCore = courseRelationDb.ToCore()
		courseRelations = append(courseRelations, courseRelationCore)
	}
	return
}

func (r *CoursesGatewayImpl) GetAccessCourseRelationsByUnitAdminId(unitAdminId string) (courseRelations []*models.CourseRelationCore, err error) {
	var courseRelationsDb []*models.CourseRelationDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("parameter = ? AND object_id = ?", "unit_admin", unitAdminId).Find(&courseRelationsDb).Error; err != nil {
			return
		}
		return
	})

	for _, courseRelationDb := range courseRelationsDb {
		var courseRelationCore *models.CourseRelationCore

		courseRelationCore = courseRelationDb.ToCore()
		courseRelations = append(courseRelations, courseRelationCore)
	}
	return
}

func (r *CoursesGatewayImpl) CreateAccessCourseRelation(courseRelation *models.CourseRelationCore) (newCourseRelation *models.CourseRelationCore, err error) {
	courseRelationDb := models.CourseRelationDB{}
	courseRelationDb.FromCore(courseRelation)
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&courseRelationDb).Error
		if err != nil {
			log.Println(err)
			return
		}
		return
	})
	newCourseRelation = courseRelationDb.ToCore()
	return
}

func (r *CoursesGatewayImpl) DeleteAccessCourseRelationById(courseRelationId string) (id string, err error) {
	courseRelationDb := models.CourseRelationDB{}
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Model(&courseRelationDb).Where("id = ?", courseRelationId).First(&models.CourseRelationDB{}).Delete(&models.CourseRelationDB{}).Error
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
	id = strconv.FormatUint(uint64(courseRelationDb.ID), 10)
	return
}

func (r *CoursesGatewayImpl) CreateCourse(course *models.CourseCore) (id string, err error) {
	courseDb := models.CourseDB{}
	courseDb.FromCore(course)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&courseDb).Error
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
	id = strconv.FormatUint(uint64(courseDb.ID), 10)
	return id, nil
}

func (r *CoursesGatewayImpl) CreateAbsoluteMedia(absoluteMedia *models.AbsoluteMediaCore) (id string, err error) {
	absoluteMediaDb := models.AbsoluteMediaDB{}
	absoluteMediaDb.FromCore(absoluteMedia)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&absoluteMediaDb).Error
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
	id = strconv.FormatUint(uint64(absoluteMediaDb.ID), 10)
	return id, nil
}

func (r *CoursesGatewayImpl) CreateMedia(media *models.MediaCore) (id string, err error) {
	mediaDb := models.MediaDB{}
	mediaDb.FromCore(media)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&mediaDb).Error
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
	id = strconv.FormatUint(uint64(mediaDb.ID), 10)
	return id, nil
}

func (r *CoursesGatewayImpl) CreateImage(image *models.ImageCore) (id string, err error) {
	imageDb := models.ImageDB{}
	imageDb.FromCore(image)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&imageDb).Error
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
	id = strconv.FormatUint(uint64(imageDb.ID), 10)
	return id, nil
}

func (r *CoursesGatewayImpl) CreateCourseApiMediaCollection(courseApiMediaCollection *models.CourseApiMediaCollectionCore) (id string, err error) {
	courseApiMediaCollectionDb := models.CourseApiMediaCollectionDB{}
	courseApiMediaCollectionDb.FromCore(courseApiMediaCollection)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(&courseApiMediaCollectionDb).Error
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
	id = strconv.FormatUint(uint64(courseApiMediaCollectionDb.ID), 10)
	return id, nil
}

func (r *CoursesGatewayImpl) DeleteCourseApiMediaCollection(courseId string) (id string, err error) {
	courseApiMediaCollectionDb := models.CourseApiMediaCollectionDB{}
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Model(&courseApiMediaCollectionDb).Where("course_id = ?", courseId).First(&models.CourseApiMediaCollectionDB{}).Delete(&models.CourseApiMediaCollectionDB{}).Error
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
	id = strconv.FormatUint(uint64(courseApiMediaCollectionDb.ID), 10)
	return
}

func (r *CoursesGatewayImpl) DeleteAbsoluteMedia(courseApiMediaCollectionId string) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Model(&models.AbsoluteMediaDB{}).Where("course_api_media_collection_id = ?", courseApiMediaCollectionId).First(&models.AbsoluteMediaDB{}).Delete(&models.AbsoluteMediaDB{}).Error
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

func (r *CoursesGatewayImpl) DeleteImage(courseApiMediaCollectionId string) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Model(&models.ImageDB{}).Where("course_api_media_collection_id = ?", courseApiMediaCollectionId).First(&models.ImageDB{}).Delete(&models.ImageDB{}).Error
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

func (r *CoursesGatewayImpl) DeleteMedia(courseApiMediaCollectionId string) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Model(&models.MediaDB{}).Where("course_api_media_collection_id = ?", courseApiMediaCollectionId).First(&models.MediaDB{}).Delete(&models.MediaDB{}).Error
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

func (r *CoursesGatewayImpl) DeleteCourse(courseId string) (id string, err error) {
	courseDb := models.CourseDB{}
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Model(&courseDb).Where("id = ?", courseId).First(&models.CourseDB{}).Delete(&models.CourseDB{}).Error
		if err != nil {
			err = courses.ErrCourseNotFound
			log.Println(err)
			return
		}
		return
	})
	if err != nil {
		log.Println(err)
		return
	}
	id = strconv.FormatUint(uint64(courseDb.ID), 10)
	return
}

func (r *CoursesGatewayImpl) UpdateCourse(course *models.CourseCore) (err error) {
	courseDb := models.CourseDB{}
	courseDb.FromCore(course)
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Model(&courseDb).Where("ID = ?", courseDb.ID).First(&models.CourseDB{}).Updates(courseDb).Error
		if err != nil {
			err = courses.ErrCourseNotFound
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

func (r *CoursesGatewayImpl) UpdateCourseApiMediaCollection(courseApiMediaCollection *models.CourseApiMediaCollectionCore) (err error) {
	courseApiMediaCollectionDb := models.CourseApiMediaCollectionDB{}
	courseApiMediaCollectionDb.FromCore(courseApiMediaCollection)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Model(&courseApiMediaCollectionDb).Where("ID = ?", courseApiMediaCollectionDb.ID).First(&models.CourseApiMediaCollectionDB{}).Updates(courseApiMediaCollectionDb).Error
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

func (r *CoursesGatewayImpl) UpdateAbsoluteMedia(absoluteMedia *models.AbsoluteMediaCore) (err error) {
	absoluteMediaDb := models.AbsoluteMediaDB{}
	absoluteMediaDb.FromCore(absoluteMedia)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Model(&absoluteMediaDb).Where("ID = ?", absoluteMediaDb.ID).First(&models.AbsoluteMediaDB{}).Updates(absoluteMediaDb).Error
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

func (r *CoursesGatewayImpl) UpdateMedia(media *models.MediaCore) (err error) {
	mediaDb := models.MediaDB{}
	mediaDb.FromCore(media)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Model(&mediaDb).Where("ID = ?", mediaDb.ID).First(&models.MediaDB{}).Updates(mediaDb).Error
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

func (r *CoursesGatewayImpl) UpdateImage(image *models.ImageCore) (err error) {
	imageDb := models.ImageDB{}
	imageDb.FromCore(image)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Model(&imageDb).Where("ID = ?", imageDb.ID).First(&models.ImageDB{}).Updates(imageDb).Error
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
