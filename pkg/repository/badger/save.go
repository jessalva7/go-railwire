package badger
//
//import (
//	"encoding/json"
//	"github.com/dgraph-io/badger/v2"
//	"github.com/jessalva/go-railwire/api/plans"
//	"log"
//)
//
//func (b badgerStore) Save(fupPlan plans.FUPPlanResponse) error {
//
//	txn := b.db.NewTransaction(true)
//
//	var currentPlans []plans.FUPPlanResponse
//	currentPlansItems, err := txn.Get([]byte("FUP_Plans"))
//	if err != nil && err != badger.ErrKeyNotFound {
//		return err
//	}
//
//	if err == nil {
//		currentPlansBytes, err := currentPlansItems.ValueCopy(nil)
//		err = json.Unmarshal(currentPlansBytes, &currentPlans)
//		if err != nil {
//			return err
//		}
//	} else {
//		currentPlans = make([]plans.FUPPlanResponse, 0)
//	}
//	log.Printf("%v", currentPlans)
//
//	currentPlans = append(currentPlans, fupPlan)
//	updatedPlans, err := json.Marshal(currentPlans)
//	if err != nil {
//		return err
//	}
//
//	err = txn.Set([]byte("FUP_Plans"), updatedPlans)
//	if err != nil {
//		return err
//	}
//	err = txn.Commit()
//	if err != nil {
//		return err
//	}
//	return nil
//}
