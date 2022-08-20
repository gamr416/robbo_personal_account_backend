package usecase

import (
	"log"

	"github.com/skinnykaen/robbo_student_personal_account.git/package/coursePacket"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"go.uber.org/fx"
)

type CoursePacketUseCaseImpl struct {
	coursePacket.Gateway
}

type CoursePacketUseCaseModule struct {
	fx.Out
	coursePacket.UseCase
}

func SetupCoursePacketUseCase(gateway coursePacket.Gateway) CoursePacketUseCaseModule {
	return CoursePacketUseCaseModule{
		UseCase: &CoursePacketUseCaseImpl{
			Gateway: gateway,
		},
	}
}

func (p *CoursePacketUseCaseImpl) CreateCoursePacket(coursePacket *models.CoursePacketCore) (id string, err error) {
	CoursePacketId, err := p.Gateway.CreateCoursePacket(coursePacket)
	if err != nil {
		log.Println("Error creating Course Packet")
		return "", err
	}

	return CoursePacketId, nil
}

func (p *CoursePacketUseCaseImpl) UpdateCoursePacket(course *models.CoursePacketCore) (err error) {

	err = p.Gateway.UpdateCoursePacket(course)
	if err != nil {
		log.Println("Error update Course")
		return
	}

	return nil
}

func (p *CoursePacketUseCaseImpl) DeleteCoursePacket(courseId string) (err error) {
	id, err := p.Gateway.DeleteCoursePacket(courseId)
	if err != nil {
		log.Println("Error delete Course")
		return
	}

	log.Println(id)

	return nil
}

func (p *CoursePacketUseCaseImpl) GetAllCoursePackets() (coursePackets []*models.CoursePacketCore, err error) {
	return p.Gateway.GetAllCoursePackets()
}

func (p *CoursePacketUseCaseImpl) GetCoursePacketById(coursePacketId string) (coursePacket *models.CoursePacketCore, err error) {
	return p.Gateway.GetCoursePacketById(coursePacketId)
}
