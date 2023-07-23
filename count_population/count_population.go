package count_population

import (
	"github.com/gennadyterekhov/tribal-relations-core/entity"
)

func CountPopulation(fertility int, availableFood int, currentPopulation int) int {
	tempTribe := createTribe()

	count := tempTribe.GetNewPopulationCount(fertility)

	return count
}

func createTribe() entity.Tribe {
	pop := entity.Population{}
	trb := entity.Tribe{
		Population: pop,
	}
	return trb
}
