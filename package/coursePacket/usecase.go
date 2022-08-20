package coursePacket

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type UseCase interface {
	CreateCoursePacket(course *models.CoursePacketCore) (id string, err error)
	DeleteCoursePacket(coursePacketId string) (err error)
	UpdateCoursePacket(course *models.CoursePacketCore) (err error)
	GetAllCoursePackets() (coursePackets []*models.CoursePacketCore, err error)
	GetCoursePacketById(coursePacketId string) (coursePacket *models.CoursePacketCore, err error)
}
