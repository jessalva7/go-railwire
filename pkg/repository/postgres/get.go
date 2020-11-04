package postgres

import "github.com/jessalva/go-railwire/api/plans"

func (p postgresDB) Get(plan plans.Plan) ([]plans.FUPPlanResponse, error) {
	panic("implement me")
}
