package coursePacket

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
)

type Delegate interface {
	CreateCoursePacket(coursePacket *models.CoursePacketHTTP, coursePacketId string) (id string, err error)
	DeleteCoursePacket(coursePacketId string) (err error)
	UpdateCoursePacket(coursePacket *models.CoursePacketHTTP) (err error)
	GetCoursePacketById(coursePacketId string) (coursePacket models.CoursePacketHTTP, err error)
	GetAllCoursePackets() (coursePackets []*models.CoursePacketHTTP, err error)
}
