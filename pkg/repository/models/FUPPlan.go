package models

import (
	"github.com/jessalva/go-railwire/api/plans"
)

type FUPPlan struct {
	StateCode        string `gorm:"primaryKey"`
	PortSpeed        int32
	PortSpeedUOM     string
	PlanRentalINR    int32
	DataUsage        int32
	DataUsageUOM     string
	AfterFUPSpeed    int32
	AfterFUPSpeedUOM string
}


func ( fupPlan *FUPPlan ) FromFUPPlanRequest( plan *plans.SavePlanRequest ){

	fupPlan.StateCode = plan.GetStateCode().String()
	fupPlan.PlanRentalINR = plan.GetFupPlan().GetPlanRentalINR()
	fupPlan.PortSpeed = plan.GetFupPlan().GetPortSpeed().GetSpeed()
	fupPlan.PortSpeedUOM = plan.GetFupPlan().GetPortSpeed().GetDataUnitType().String()
	fupPlan.DataUsage = plan.GetFupPlan().GetDataUsage().GetDataAmount()
	fupPlan.DataUsageUOM = plan.GetFupPlan().GetDataUsage().GetDataUnitType().String()
	fupPlan.AfterFUPSpeed = plan.GetFupPlan().GetAfterFUP().GetSpeed()
	fupPlan.AfterFUPSpeedUOM = plan.GetFupPlan().GetAfterFUP().GetDataUnitType().String()

}

func ( fupPlan *FUPPlan ) ToSavePlanResponse() *plans.SavePlanResponse{



	return &plans.SavePlanResponse{
		StateCode: *plans.StateCode( plans.StateCode_value[fupPlan.StateCode] ).Enum(),
		PlanType: plans.PlanType_FUPPlan,
		PlanResponse: &plans.SavePlanResponse_FupPlan{

			FupPlan: &plans.FupPlanType{
				PortSpeed: &plans.PortSpeedType{
					Speed:        fupPlan.PortSpeed,
					DataUnitType: *plans.SpeedType(plans.SpeedType_value[fupPlan.PortSpeedUOM]).Enum(),
				},
					PlanRentalINR: fupPlan.PlanRentalINR,
					AfterFUP: &plans.PortSpeedType{
					Speed:        fupPlan.AfterFUPSpeed,
					DataUnitType: *plans.SpeedType(plans.SpeedType_value[fupPlan.AfterFUPSpeedUOM]).Enum(),
				},
					DataUsage: &plans.DataUsageType{
					DataAmount:   fupPlan.DataUsage,
					DataUnitType: *plans.DataUnitType(plans.DataUnitType_value[fupPlan.DataUsageUOM]).Enum(),
				},
			},
		},
	}

}
