package postgres

import (
	"github.com/jessalva/go-railwire/api/plans"
	"github.com/jessalva/go-railwire/pkg/repository/models"
	"log"
)

func (p postgresDB) Save(plan *plans.SavePlanRequest) ( *plans.SavePlanResponse, error) {


	response := &plans.SavePlanResponse{
		StateCode: *plans.StateCode(plans.StateCode_value[plan.StateCode.String()]).Enum(),
		PlanType:  plans.PlanType_FUP,
	}

	if plan.GetFupPlan() != nil {

		fupPlan := &models.FUPPlan{}
		fupPlan.FromFUPPlanRequest( plan )
		result := p.db.Create( fupPlan )
		if result.Error != nil {
			log.Print( result.Error )
		}
		log.Print("Created FUP Plan ", fupPlan.StateCode)

		response.PlanResponse = &plans.SavePlanResponse_FupPlan{

			FupPlan: fupPlan.ToFupPlan(),

		}

	}


	return response,nil

}