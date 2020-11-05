package repository

import "github.com/jessalva/go-railwire/api/plans"

type PlanRepository interface {
	Save(plan *plans.SavePlanRequest) (*plans.SavePlanResponse,error)
	Get(plan plans.Plan) ([]plans.FUPPlanResponse, error)
}
