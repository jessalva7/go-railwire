package models

import (
	"github.com/jessalva/go-railwire/api/plans"
	"gorm.io/gorm"
)

type FUPPlan struct {
	gorm.Model
	PortSpeed     int32
	PortSpeedUOM  string
	PlanRentalINR int32
	DataUsage   int32
	DataUsageUOM string
	AfterFUPSpeed int32
	AfterFUPSpeedUOM string
}


func ( fupPlan *FUPPlan ) FromFUPPlanRequest( plan *plans.SavePlanRequest ){

	fupPlan.PlanRentalINR = plan.GetFupPlanRequest().GetPlanRentalINR()
	fupPlan.PortSpeed = plan.GetFupPlanRequest().GetPortSpeed().GetSpeed()
	fupPlan.PortSpeedUOM = plan.GetFupPlanRequest().GetPortSpeed().GetDataUnitType().String()
	fupPlan.DataUsage = plan.GetFupPlanRequest().GetDataUsage().GetDataAmount()
	fupPlan.DataUsageUOM = plan.GetFupPlanRequest().GetDataUsage().GetDataUnitType().String()
	fupPlan.AfterFUPSpeed = plan.GetFupPlanRequest().GetAfterFUP().GetSpeed()
	fupPlan.AfterFUPSpeedUOM = plan.GetFupPlanRequest().GetAfterFUP().GetDataUnitType().String()

}