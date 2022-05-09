package delegate

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/usecase"
	"go.uber.org/fx"
)

type Module struct {
	fx.Out
	AuthDelegate
}

func Setup(AuthUseCase usecase.AuthUseCase) Module {
	return Module{
		AuthDelegate: AuthDelegate{
			AuthUseCase,
		},
	}
}
