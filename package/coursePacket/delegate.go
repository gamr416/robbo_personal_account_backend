package coursePacket

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
)

type Delegate interface {
	CreateCoursePacket(course *models.CoursePacketHTTP, coursePacketId string) (id string, err error)
	DeleteCoursePacket(courseId string) (err error)
	UpdateCoursePacket(course *models.CoursePacketHTTP) (err error)
	GetCoursePacketById(coursePacketId string) (coursePacket models.CoursePacketHTTP, err error)
	GetAllCoursePackets() (coursePackets []*models.CoursePacketHTTP, err error)
}
