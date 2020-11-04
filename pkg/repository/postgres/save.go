package postgres

import (
	"github.com/jessalva/go-railwire/api/plans"
	"github.com/jessalva/go-railwire/pkg/repository/models"
	"log"
)

func (p postgresDB) Save(plan *plans.SavePlanRequest) error {


	if plan.GetFupPlanRequest() != nil {

		fupPlan := &models.FUPPlan{}
		fupPlan.FromFUPPlanRequest( plan )
		result := p.db.Create( fupPlan )
		if result.Error != nil {
			log.Print( result.Error )
		}
		log.Print("Created Plan ", fupPlan.ID)

	}


	return nil

}