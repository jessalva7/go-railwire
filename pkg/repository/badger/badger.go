package badger

//type badgerStore struct {
//	db *badger.DB
//}
//
//func NewBadger() repository.PlanRepository {
//
//	var badgerDir string
//	if os.Getenv("BADGER_DB_DIR") == "" {
//		badgerDir = "./resources/badger"
//	} else {
//		badgerDir = os.Getenv("BADGER_DB_DIR")
//	}
//
//	opts := badger.DefaultOptions(badgerDir)
//
//	db, err := badger.Open(opts)
//	if err != nil {
//
//		log.Fatal("couldn't open DB")
//
//	}
//
//	return badgerStore{db: db}
//
//}
