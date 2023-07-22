package main

import (
	"github.com/gennadyterekhov/tribal-relations-core/entity"
	// "github.com/gennadyterekhov/tribal-relations/core/population.go"
)

//	"github.com/gennadyterekhov/tribal-relations/core/entity/tribe.go"

func CountPopulation(fertility int, availableFood int, currentPopulation int) int {
	temp_tribe := createTribe()

	count := temp_tribe.GetNewPopulationCount(fertility)

	return count
}

func createTribe() entity.Tribe {
	pop := entity.Population{}
	trb := entity.Tribe{
		Population: pop,
	}
	return trb
}
