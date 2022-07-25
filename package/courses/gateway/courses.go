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
	courseApiMediaCollection := models.CourseApiMediaCollectionDB{}
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Where("course_id = ?", courseId).First(&courseApiMediaCollection).Error
		if err != nil {
			log.Println(err)
			return
		}
		err = tx.Where("course_id = ?", courseId).Delete(&models.CourseApiMediaCollectionDB{}).Error
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
	id = strconv.FormatUint(uint64(courseApiMediaCollection.ID), 10)
	return
}

func (r *CoursesGatewayImpl) DeleteAbsoluteMedia(courseApiMediaCollectionId string) (err error) {
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Where("course_api_media_collection_id = ?", courseApiMediaCollectionId).Delete(&models.AbsoluteMediaDB{}).Error
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
		err = tx.Where("course_api_media_collection_id = ?", courseApiMediaCollectionId).Delete(&models.ImageDB{}).Error
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
		err = tx.Where("course_api_media_collection_id = ?", courseApiMediaCollectionId).Delete(&models.MediaDB{}).Error
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

func (r CoursesGatewayImpl) DeleteCourse(courseId string) (id string, err error) {
	course := models.CourseDB{}
	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Where("id = ?", courseId).First(&course).Error
		if err != nil {
			log.Println(err)
			return
		}
		err = tx.Model(&course).Where("id = ?", courseId).Delete(&models.CourseDB{}).Error
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
	id = strconv.FormatUint(uint64(course.ID), 10)
	return
}

func (r *CoursesGatewayImpl) UpdateCourse(course *models.CourseCore) (err error) {
	courseDb := models.CourseDB{}
	courseDb.FromCore(course)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Where("id = ?", courseDb.ID).First(&models.CourseDB{}).Error
		if err != nil {
			log.Println(err)
			return
		}
		err = tx.Model(&courseDb).Where("ID = ?", courseDb.ID).Updates(courseDb).Error
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

func (r *CoursesGatewayImpl) UpdateCourseApiMediaCollection(courseApiMediaCollection *models.CourseApiMediaCollectionCore) (err error) {
	courseApiMediaCollectionDb := models.CourseApiMediaCollectionDB{}
	courseApiMediaCollectionDb.FromCore(courseApiMediaCollection)

	err = r.PostgresClient.Db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Model(&courseApiMediaCollectionDb).Where("ID = ?", courseApiMediaCollectionDb.ID).Updates(courseApiMediaCollectionDb).Error
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
		err = tx.Model(&absoluteMediaDb).Where("ID = ?", absoluteMediaDb.ID).Updates(absoluteMediaDb).Error
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
		err = tx.Model(&mediaDb).Where("ID = ?", mediaDb.ID).Updates(mediaDb).Error
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
		err = tx.Model(&imageDb).Where("ID = ?", imageDb.ID).Updates(imageDb).Error
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
