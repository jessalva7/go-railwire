package postgres

import (
	"github.com/jessalva/go-railwire/api/plans"
	"github.com/jessalva/go-railwire/pkg/repository/models"
	"log"
)

func (p postgresDB) Save(plan *plans.SavePlanRequest) ( *plans.SavePlanResponse, error) {


	if plan.GetFupPlan() != nil {

		fupPlan := &models.FUPPlan{}
		fupPlan.FromFUPPlanRequest( plan )
		result := p.db.Create( fupPlan )
		if result.Error != nil {
			log.Print( result.Error )
		}
		log.Print("Created FUP Plan ", fupPlan.StateCode)

		return fupPlan.ToSavePlanResponse(),nil
	}


	return &plans.SavePlanResponse{},nil

}