package coursePacket

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type Gateway interface {
	CreateCoursePacket(coursePacket *models.CoursePacketCore) (id string, err error)
	DeleteCoursePacket(coursePacketId string) (id string, err error)
	UpdateCoursePacket(crsPacket *models.CoursePacketCore) (err error)
	GetAllCoursePackets() (coursePackets []*models.CoursePacketCore, err error)
	GetCoursePacketById(coursePacketId string) (crsPacket *models.CoursePacketCore, err error)
}
