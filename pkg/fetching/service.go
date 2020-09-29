package fetching

import (
	"context"
	"github.com/jessalva/go-railwire/api/plans"
	"log"
)

type service struct {
	log *log.Logger
}

func (s *service) GetPlan(context context.Context, plan *plans.Plan) (*plans.PlanResponse, error) {

	s.log.Print("Got Request ", plan.StateCode)

	plansForState := &plans.PlanResponse{FupPlans: []*plans.FUPPlanResponse{
			&plans.FUPPlanResponse{
				AfterFUP:      &plans.PortSpeedType{DataUnitType: plans.SpeedType_MegaBitPerSecond, Speed: 10},
				DataUsage:     &plans.DataUsageType{DataUnitType: plans.DataUnitType_GB, DataAmount: 500},
				PortSpeed:     &plans.PortSpeedType{DataUnitType: plans.SpeedType_MegaBitPerSecond, Speed: 1},
				PlanRentalINR: 449,
			},
			&plans.FUPPlanResponse{
				AfterFUP:      &plans.PortSpeedType{DataUnitType: plans.SpeedType_MegaBitPerSecond, Speed: 10},
				DataUsage:     &plans.DataUsageType{DataUnitType: plans.DataUnitType_GB, DataAmount: 1000},
				PortSpeed:     &plans.PortSpeedType{DataUnitType: plans.SpeedType_MegaBitPerSecond, Speed: 2},
				PlanRentalINR: 499,
			},
			&plans.FUPPlanResponse{
				AfterFUP:      &plans.PortSpeedType{DataUnitType: plans.SpeedType_MegaBitPerSecond, Speed: 20},
				DataUsage:     &plans.DataUsageType{DataUnitType: plans.DataUnitType_GB, DataAmount: 1000},
				PortSpeed:     &plans.PortSpeedType{DataUnitType: plans.SpeedType_MegaBitPerSecond, Speed: 4},
				PlanRentalINR: 699,
			},
		},
	}

	s.log.Printf("%v", plansForState)

	return plansForState, nil

}

func NewService(logger *log.Logger) *plans.PlansService {
	fetchingService := &service{log: logger}
	return &plans.PlansService{GetPlan: fetchingService.GetPlan}
}
