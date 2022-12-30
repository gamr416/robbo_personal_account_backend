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

func (r *CoursesGatewayImpl) DeleteCourseRelationsByRobboUnitId(robboUnitId string) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Where("parameter = ? AND object_id", "unit", robboUnitId).Delete(&models.CourseRelationsDB{}).Error
		return
	})
	return
}

func (r *CoursesGatewayImpl) DeleteCourseRelationsByRobboGroupId(robboGroupId string) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Where("parameter = ? AND object_id", "group", robboGroupId).Delete(&models.CourseRelationsDB{}).Error
		return
	})
	return
}

func (r *CoursesGatewayImpl) DeleteCourseRelationsByRoleId(roleId string) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Where("parameter = ? AND object_id", "role", roleId).Delete(&models.CourseRelationsDB{}).Error
		return
	})
	return
}

func (r *CoursesGatewayImpl) GetCourseRelationsByCourseId(courseId string) (courseRelations []*models.CourseRelationsCore, err error) {
	var courseRelationsDb []*models.CourseRelationsDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("course_id", courseId).Find(&courseRelationsDb).Error; err != nil {
			return
		}
		return
	})

	for _, courseRelationDb := range courseRelationsDb {
		var courseRelationCore *models.CourseRelationsCore

		courseRelationCore = courseRelationDb.ToCore()
		courseRelations = append(courseRelations, courseRelationCore)
	}
	return
}

func (r *CoursesGatewayImpl) GetCourseRelationsGroups() (courseRelations []*models.CourseRelationsCore, err error) {
	var courseRelationsDb []*models.CourseRelationsDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("parameter = ?", "group").Find(&courseRelationsDb).Error; err != nil {
			return
		}
		return
	})

	for _, courseRelationDb := range courseRelationsDb {
		var courseRelationCore *models.CourseRelationsCore

		courseRelationCore = courseRelationDb.ToCore()
		courseRelations = append(courseRelations, courseRelationCore)
	}
	return
}

func (r *CoursesGatewayImpl) GetCourseRelationsUnits() (courseRelations []*models.CourseRelationsCore, err error) {
	var courseRelationsDb []*models.CourseRelationsDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("parameter = ?", "unit").Find(&courseRelationsDb).Error; err != nil {
			return
		}
		return
	})

	for _, courseRelationDb := range courseRelationsDb {
		var courseRelationCore *models.CourseRelationsCore

		courseRelationCore = courseRelationDb.ToCore()
		courseRelations = append(courseRelations, courseRelationCore)
	}
	return
}

func (r *CoursesGatewayImpl) GetCourseRelationsRoles() (courseRelations []*models.CourseRelationsCore, err error) {
	var courseRelationsDb []*models.CourseRelationsDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("parameter = ?", "role").Find(&courseRelationsDb).Error; err != nil {
			return
		}
		return
	})

	for _, courseRelationDb := range courseRelationsDb {
		var courseRelationCore *models.CourseRelationsCore

		courseRelationCore = courseRelationDb.ToCore()
		courseRelations = append(courseRelations, courseRelationCore)
	}
	return
}

func (r *CoursesGatewayImpl) GetCourseRelationsByRobboUnitId(robboUnitId string) (courseRelations []*models.CourseRelationsCore, err error) {
	var courseRelationsDb []*models.CourseRelationsDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("parameter = ? AND object_id = ?", "unit", robboUnitId).Find(&courseRelationsDb).Error; err != nil {
			return
		}
		return
	})

	for _, courseRelationDb := range courseRelationsDb {
		var courseRelationCore *models.CourseRelationsCore

		courseRelationCore = courseRelationDb.ToCore()
		courseRelations = append(courseRelations, courseRelationCore)
	}
	return
}

func (r *CoursesGatewayImpl) GetCourseRelationByRobboGroupId(robboGroupId string) (courseRelations []*models.CourseRelationsCore, err error) {
	var courseRelationsDb []*models.CourseRelationsDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("parameter = ? AND object_id = ?", "group", robboGroupId).Find(&courseRelationsDb).Error; err != nil {
			return
		}
		return
	})

	for _, courseRelationDb := range courseRelationsDb {
		var courseRelationCore *models.CourseRelationsCore

		courseRelationCore = courseRelationDb.ToCore()
		courseRelations = append(courseRelations, courseRelationCore)
	}
	return
}

func (r *CoursesGatewayImpl) GetCourseRelationByRoleId(roleId string) (courseRelations []*models.CourseRelationsCore, err error) {
	var courseRelationsDb []*models.CourseRelationsDB
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		if err = tx.Where("parameter = ? AND object_id = ?", "role", roleId).Find(&courseRelationsDb).Error; err != nil {
			return
		}
		return
	})

	for _, courseRelationDb := range courseRelationsDb {
		var courseRelationCore *models.CourseRelationsCore

		courseRelationCore = courseRelationDb.ToCore()
		courseRelations = append(courseRelations, courseRelationCore)
	}
	return
}

func (r *CoursesGatewayImpl) CreateCourseRelation(courseRelation *models.CourseRelationsCore) (newCourseRelation *models.CourseRelationsCore, err error) {
	courseRelationDb := models.CourseRelationsDB{}
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

func (r *CoursesGatewayImpl) DeleteCourseRelationById(courseRelationId string) (id string, err error) {
	courseRelationDb := models.CourseRelationsDB{}
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Model(&courseRelationDb).Where("id = ?", courseRelationId).First(&models.CourseRelationsDB{}).Delete(&models.CourseRelationsDB{}).Error
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
