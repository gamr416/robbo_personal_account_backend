package coursePacket

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type UseCase interface {
	CreateCoursePacket(coursePacket *models.CoursePacketCore) (id string, err error)
	DeleteCoursePacket(coursePacketId string) (err error)
	UpdateCoursePacket(coursePacket *models.CoursePacketCore) (err error)
	GetAllCoursePackets() (coursePackets []*models.CoursePacketCore, err error)
	GetCoursePacketById(coursePacketId string) (coursePacket *models.CoursePacketCore, err error)
}
