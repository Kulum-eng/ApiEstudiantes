package dependencies

import (
	"github.com/gin-gonic/gin"

	"API-HEXAGONAL/src/estudiantes/application/useCaseShinobi"
	"API-HEXAGONAL/src/estudiantes/domain"
	"API-HEXAGONAL/src/estudiantes/infrastructure/controllers"
)

type ShinobiDependencies struct {
	CreateShinobiController  *controllers.CreateShinobiController
	UpdateShinobiController  *controllers.UpdateShinobiController
	DeleteShinobiController  *controllers.DeleteShinobiController
	GetAllShinobisController gin.HandlerFunc
	GetByIdShinobiController gin.HandlerFunc
}

func NewShinobiDependencies(repo domain.IShinobi) *ShinobiDependencies {
	createShinobiUseCase := useCaseShinobi.NewCreateShinobi(repo)
	updateShinobiUseCase := useCaseShinobi.NewUpdateShinobi(repo)
	deleteShinobiUseCase := useCaseShinobi.NewDeleteShinobi(repo)

	return &ShinobiDependencies{
		CreateShinobiController:  controllers.NewCreateShinobiController(createShinobiUseCase),
		UpdateShinobiController:  controllers.NewUpdateShinobiController(updateShinobiUseCase),
		DeleteShinobiController:  controllers.NewDeleteShinobiController(deleteShinobiUseCase),
		GetAllShinobisController: controllers.GetAllShinobisController(repo),
		GetByIdShinobiController: controllers.GetByIdShinobiController(repo),
	}
}
