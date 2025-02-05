package dependenciesJutsu

import (
	"API-HEXAGONAL/src/jutsu/application/useCaseJutsu"
	"API-HEXAGONAL/src/jutsu/domain"
	"API-HEXAGONAL/src/jutsu/infrastructureJutsu/controllers"
	"github.com/gin-gonic/gin"
)

type JutsuDependencies struct {
	CreateJutsuController  *controllers.CreateJutsuController
	UpdateJutsuController  *controllers.UpdateJutsuController
	DeleteJutsuController  *controllers.DeleteJutsuController
	GetAllJutsuController  gin.HandlerFunc
	GetByIdJutsuController gin.HandlerFunc
}

func NewJutsuDependencies(repo domain.IJutsu) *JutsuDependencies {
	createJutsuUseCase := useCaseJutsu.NewCreateJutsu(repo)
	updateJutsuUseCase := useCaseJutsu.NewUpdateJutsu(repo)
	deleteJutsuUseCase := useCaseJutsu.NewDeleteJutsu(repo)

	return &JutsuDependencies{
		CreateJutsuController:  controllers.NewCreateJutsuController(createJutsuUseCase),
		UpdateJutsuController:  controllers.NewUpdateJutsuController(updateJutsuUseCase),
		DeleteJutsuController:  controllers.NewDeleteJutsuController(deleteJutsuUseCase),
		GetAllJutsuController:  controllers.GetAllJutsuController(repo),
		GetByIdJutsuController: controllers.GetByIdJutsuController(repo),
	}
}
