package fetching

import (
	"context"
	"github.com/jessalva/go-railwire/api/plans"
	"github.com/jessalva/go-railwire/pkg/repository"
	"log"
)

type service struct {
	log            *log.Logger
	planRepository repository.PlanRepository
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

func (s *service) SavePlan(context context.Context, plan *plans.SavePlanRequest) (*plans.SavePlanResponse, error) {

	s.log.Print("Got Request ", plan.PlanType)
	savedPlan, err := s.planRepository.Save(plan)
	if err != nil {
		return nil, err
	}
	return savedPlan, nil

}

func NewService(logger *log.Logger, repo repository.PlanRepository) *plans.PlansService {
	fetchingService := &service{log: logger , planRepository: repo}
	return &plans.PlansService{GetPlan: fetchingService.GetPlan, SavePlan: fetchingService.SavePlan}
}
