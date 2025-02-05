package dependencies

import (
	"API-HEXAGONAL/src/shinobi/application/useCaseShinobi"
	"API-HEXAGONAL/src/shinobi/domain"
	"API-HEXAGONAL/src/shinobi/infrastructure/controllers"
	"github.com/gin-gonic/gin"
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
