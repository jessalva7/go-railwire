package postgres

import (
	"errors"
	"fmt"
	"github.com/jessalva/go-railwire/api/plans"
	"github.com/jessalva/go-railwire/pkg/repository/models"
	"log"
)

func (p postgresDB) GetPlansOfSpecificTypeAndState(planRequest *plans.GetPlanRequest) (*plans.GetPlansResponse, error) {

	switch planRequest.PlanType {

		case plans.PlanType_FUP:
			var modelFupPlans []models.FUPPlan
			log.Print( plans.PlanType_FUP.String() ," ", planRequest.GetStateCode().String())
			p.db.Where(&models.FUPPlan{StateCode: planRequest.GetStateCode().String()}).Find(&modelFupPlans)

			var fupPlans []*plans.FupPlan
			fupPlans = make([]*plans.FupPlan,0)
			for _, plan := range modelFupPlans {
				fupPlans = append(fupPlans,plan.ToFupPlan())
			}

			response := &plans.GetPlansResponse{
				Plans: &plans.GetPlansResponse_FupPlans{
						FupPlans: &plans.FUPPlans{FupPlans: fupPlans,
					},
				},
			}
			return response, nil

	}

	return nil,errors.New(fmt.Sprintf("Couldn't find plan for plan type: %v an state: %v", planRequest.PlanType, planRequest.StateCode))
}
