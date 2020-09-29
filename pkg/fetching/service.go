package fetching

import (
	"context"
	"github.com/jessalva/go-railwire/api/plans"
	"log"
)

type service struct {
	log *log.Logger
}

func ( s *service) 	GetPlan(context context.Context, plan *plans.Plan) (*plans.FUPPlanResponse, error) {

	s.log.Print("Got Request ", plan.StateCode)

	mainPlan := &plans.FUPPlanResponse{
		AfterFUP:      &plans.PortSpeedType{DataUnitType: plans.SpeedType_MegaBitPerSecond, Speed: 10},
		DataUsage:     &plans.DataUsageType{DataUnitType: plans.DataUnitType_GB, DataAmount: 500},
		PortSpeed:     &plans.PortSpeedType{DataUnitType: plans.SpeedType_MegaBitPerSecond, Speed: 1},
		PlanRentalINR: 499,
	}

	s.log.Printf("%v", mainPlan)

	return mainPlan, nil

}

func NewService( logger *log.Logger) *plans.PlansService{
	fetchingService := &service{ log: logger}
	return &plans.PlansService{ GetPlan: fetchingService.GetPlan }
}