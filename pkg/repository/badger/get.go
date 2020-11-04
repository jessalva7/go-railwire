package badger
//
//import (
//	"encoding/json"
//	"github.com/dgraph-io/badger/v2"
//	"github.com/jessalva/go-railwire/api/plans"
//)
//
//func (b badgerStore) Get(plan plans.Plan) ([]plans.FUPPlanResponse, error) {
//
//	txn := b.db.NewTransaction(false)
//
//	var currentPlans []plans.FUPPlanResponse
//	currentPlansItems, err := txn.Get([]byte("FUP_Plans"))
//	if err != nil && err != badger.ErrKeyNotFound {
//		return nil, err
//	}
//
//	if err == nil {
//		currentPlansBytes, err := currentPlansItems.ValueCopy(nil)
//		err = json.Unmarshal(currentPlansBytes, &currentPlans)
//		if err != nil {
//			return nil, err
//		}
//	} else {
//		currentPlans = make([]plans.FUPPlanResponse, 0)
//	}
//
//	return currentPlans, nil
//}
