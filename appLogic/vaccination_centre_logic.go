package appLogic

import (
	"awesomeProject/constants"
	"awesomeProject/entity"
	"fmt"
)

var VaccinationCentres = []entity.VaccinationCentre{}

func GetNewVaccinationCentre(ID string, state string, district string, capacity [10]int) entity.VaccinationCentre {
	return entity.VaccinationCentre{
		CentreID: ID,
		State:    state,
		District: district,
		Capacity: capacity,
	}
}

func AddNewVaccinationCentre(newVaccinationCentre entity.VaccinationCentre) {
	VaccinationCentres = append(VaccinationCentres, newVaccinationCentre)
	fmt.Println("Added successfully")
}

func AddCapacity(centreId string, capacity int, day int) {
	success := constants.FALSE
	fmt.Println()
	for _, centre := range VaccinationCentres {
		if centre.CentreID == centreId {
			centre.Capacity[day-1] += capacity
			success = constants.TRUE
			fmt.Println(centre)
			break
		}
	}
	if !success {
		fmt.Println("No such centre exists")
	} else {
		fmt.Println("Capacity added successfully")
	}
}

func ListAllVaccinationCentres(district string) {
	for _, centre := range VaccinationCentres {
		if centre.District == district {
			fmt.Println(centre.CentreID, centre.State, centre.District)
		}
	}
}
