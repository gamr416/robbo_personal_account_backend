package coursePacket

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type Gateway interface {
	CreateCoursePacket(coursePacketCore *models.CoursePacketCore) (id string, err error)
	DeleteCoursePacket(coursePacketId string) (id string, err error)
	UpdateCoursePacket(coursePacketCore *models.CoursePacketCore) (err error)
	GetAllCoursePackets() (coursePacketsCore []*models.CoursePacketCore, err error)
	GetCoursePacketById(coursePacketId string) (coursePacketCore *models.CoursePacketCore, err error)
}
