package repository

import "github.com/jessalva/go-railwire/api/plans"

type PlanRepository interface {
	Save(plan *plans.SavePlanRequest) (*plans.SavePlanResponse,error)
	GetPlansOfSpecificTypeAndState(plan *plans.GetPlanRequest) (*plans.GetPlansResponse, error)
}
