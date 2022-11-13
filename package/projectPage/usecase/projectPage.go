package usecase

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/projectPage"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/projects"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type ProjectPageUseCaseImpl struct {
	projectPageGateway projectPage.Gateway
	projectGateway     projects.Gateway
}

type ProjectPageUseCaseModule struct {
	fx.Out
	projectPage.UseCase
}

func SetupProjectPageUseCase(projectPageGateway projectPage.Gateway, projectGateway projects.Gateway) ProjectPageUseCaseModule {
	return ProjectPageUseCaseModule{
		UseCase: &ProjectPageUseCaseImpl{
			projectPageGateway: projectPageGateway,
			projectGateway:     projectGateway,
		},
	}
}

const emptyProjectJson = "{\"targets\":[{\"isStage\":true,\"name\":\"Stage\",\"variables\":{\"`jEk@4|i[#Fk?(8x)AV." +
	"-my variable\":[\"my variable\",0]},\"lists\":{},\"broadcasts\":{},\"blocks\":{},\"comments\":{},\"currentCos" +
	"tume\":0,\"costumes\":[{\"assetId\":\"cd21514d0531fdffb22204e0ec5ed84a\",\"name\":\"backdrop1\",\"md5ext\":\"" +
	"cd21514d0531fdffb22204e0ec5ed84a.svg\",\"dataFormat\":\"svg\",\"rotationCenterX\":240,\"rotationCenterY\":180" +
	"}],\"sounds\":[{\"assetId\":\"83a9787d4cb6f3b7632b4ddfebf74367\",\"name\":\"pop\",\"dataFormat\":\"wav\",\"fo" +
	"rmat\":\"\",\"rate\":44100,\"sampleCount\":1032,\"md5ext\":\"83a9787d4cb6f3b7632b4ddfebf74367.wav\"}],\"volum" +
	"e\":100,\"layerOrder\":0,\"tempo\":60,\"videoTransparency\":50,\"videoState\":\"on\",\"textToSpeechLanguage\"" +
	":null},{\"isStage\":false,\"name\":\"Sprite1\",\"variables\":{},\"lists\":{},\"broadcasts\":{},\"blocks\":{}," +
	"\"comments\":{},\"currentCostume\":0,\"costumes\":[{\"assetId\":\"bcf454acf82e4504149f7ffe07081dbc\",\"name\"" +
	":\"costume1\",\"bitmapResolution\":1,\"md5ext\":\"bcf454acf82e4504149f7ffe07081dbc.svg\",\"dataFormat\":\"svg" +
	"\",\"rotationCenterX\":48,\"rotationCenterY\":50},{\"assetId\":\"0fb9be3e8397c983338cb71dc84d0b25\",\"name\":" +
	"\"costume2\",\"bitmapResolution\":1,\"md5ext\":\"0fb9be3e8397c983338cb71dc84d0b25.svg\",\"dataFormat\":\"svg\"" +
	",\"rotationCenterX\":46,\"rotationCenterY\":53}],\"sounds\":[{\"assetId\":\"83c36d806dc92327b9e7049a565c6bff\"" +
	",\"name\":\"Meow\",\"dataFormat\":\"wav\",\"format\":\"\",\"rate\":44100,\"sampleCount\":37376,\"md5ext\":\"8" +
	"3c36d806dc92327b9e7049a565c6bff.wav\"}],\"volume\":100,\"layerOrder\":1,\"visible\":true,\"x\":0,\"y\":0,\"si" +
	"ze\":100,\"direction\":90,\"draggable\":false,\"rotationStyle\":\"all around\"}],\"monitors\":[],\"extensions" +
	"\":[],\"meta\":{\"semver\":\"3.0.0\",\"vm\":\"0.2.0-prerelease.20220519142410\",\"agent\":\"Mozilla/5.0 (X11;" +
	" Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36\"}}"

func (p *ProjectPageUseCaseImpl) CreateProjectPage(authorId string) (projectId string, err error) {
	project := models.ProjectCore{}
	project.AuthorId = authorId
	project.Json = emptyProjectJson
	project.Name = "Untitled"

	projectId, err = p.projectGateway.CreateProject(&project)
	if err != nil {
		return "", err
	}

	projectPage := &models.ProjectPageCore{
		Title:       "Untitled",
		ProjectId:   projectId,
		Instruction: "",
		Notes:       "",
		Preview:     "",
		LinkScratch: viper.GetString("projectPage.scratchLink") + "?#" + projectId,
		IsShared:    false,
	}
	_, err = p.projectPageGateway.CreateProjectPage(projectPage)
	if err != nil {
		return "", err
	}
	return
}

func (p *ProjectPageUseCaseImpl) UpdateProjectPage(projectPage *models.ProjectPageCore) (err error) {
	return p.projectPageGateway.UpdateProjectPage(projectPage)
}

func (p *ProjectPageUseCaseImpl) DeleteProjectPage(projectId string) (err error) {
	err = p.projectGateway.DeleteProject(projectId)
	if err != nil {
		return
	}
	return p.projectPageGateway.DeleteProjectPage(projectId)
}

func (p *ProjectPageUseCaseImpl) GetAllProjectPageByUserId(authorId string) (projectPages []*models.ProjectPageCore, err error) {
	projects, err := p.projectGateway.GetProjectsByAuthorId(authorId)
	if err != nil {
		return
	}
	for _, project := range projects {
		projectPage, errGetProjectPageById := p.projectPageGateway.GetProjectPageByProjectId(project.ID)
		if errGetProjectPageById != nil {
			return []*models.ProjectPageCore{}, errGetProjectPageById
		}
		projectPages = append(projectPages, projectPage)
	}
	return
}

func (p *ProjectPageUseCaseImpl) GetProjectPageById(projectPageId string) (projectPage *models.ProjectPageCore, err error) {
	return p.projectPageGateway.GetProjectPageById(projectPageId)
}
