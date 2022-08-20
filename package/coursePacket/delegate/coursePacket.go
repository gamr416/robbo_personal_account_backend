package delegate

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/coursePacket"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"go.uber.org/fx"
)

type CoursePacketDelegateImpl struct {
	coursePacket.UseCase
}

type CoursePacketDelegateModule struct {
	fx.Out
	coursePacket.Delegate
}

func SetupCoursePacketDelegate(usecase coursePacket.UseCase) CoursePacketDelegateModule {
	return CoursePacketDelegateModule{
		Delegate: &CoursePacketDelegateImpl{
			usecase,
		},
	}
}

func (p *CoursePacketDelegateImpl) CreateCoursePacket(coursePacket *models.CoursePacketHTTP, coursePacketId string) (id string, err error) {

	coursePacketCore := coursePacket.ToCore()
	return p.UseCase.CreateCoursePacket(coursePacketCore)
}

func (p *CoursePacketDelegateImpl) DeleteCoursePacket(coursePacketId string) (err error) {
	return p.UseCase.DeleteCoursePacket(coursePacketId)
}

func (p *CoursePacketDelegateImpl) UpdateCoursePacket(coursePacket *models.CoursePacketHTTP) (err error) {
	coursePacketCore := coursePacket.ToCore()
	return p.UseCase.UpdateCoursePacket(coursePacketCore)
}

func (p *CoursePacketDelegateImpl) GetCoursePacketById(coursePacketId string) (coursePacket models.CoursePacketHTTP, err error) {
	coursePacketCore, getCoursePacketErr := p.UseCase.GetCoursePacketById(coursePacketId)
	if getCoursePacketErr != nil {
		err = getCoursePacketErr
		return
	}
	coursePacket.FromCore(coursePacketCore)
	return
}

func (p *CoursePacketDelegateImpl) GetAllCoursePackets() (coursePackets []*models.CoursePacketHTTP, err error) {

	coursePacketsCore, getCoursePacketsErr := p.UseCase.GetAllCoursePackets()
	if getCoursePacketsErr != nil {
		err = getCoursePacketsErr
		return
	}
	for _, coursePacketCore := range coursePacketsCore {
		var coursePacketTemp models.CoursePacketHTTP
		coursePacketTemp.FromCore(coursePacketCore)
		coursePackets = append(coursePackets, &coursePacketTemp)
	}
	return
}
