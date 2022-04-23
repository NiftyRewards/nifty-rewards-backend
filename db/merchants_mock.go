package db

//var err error
//
//const SecondAddressW3a = "0x123"
//const FirstAddressB = "0x456"
//
//func MockGetUer(db *pg.DB) bool {
//	user, err := GetUser(db, SecondAddressW3a)
//	if err != nil {
//		log.Printf("[TestGetUser] err: %v", err)
//	}
//
//	if user.AddressW3a != SecondAddressW3a {
//		log.Printf("user.MerchantId != SecondAddressW3a")
//		return false
//	}
//	if user.Address_B != FirstAddressB {
//		log.Printf("user.MerchantName != FirstAddressB")
//		return false
//	}
//
//	log.Printf("MockGetUser passed")
//	return true
//}
//
//func MockCreateUsr(db *pg.DB) bool {
//	_, err := CreateUser(db, SecondAddressW3a)
//	log.Printf("[MockCreateUser] CreateUser err: %v", err)
//
//	user, err := GetUser(db, SecondAddressW3a)
//	if err != nil {
//		log.Printf("[MockCreateUser] GetUser: %v", err)
//	}
//
//	if user.AddressW3a != SecondAddressW3a {
//		log.Printf("user.MerchantId != SecondAddressW3a")
//		return false
//	}
//	if user.Address_B != FirstAddressB {
//		log.Printf("user.MerchantId != SecondAddressW3a")
//		return false
//	}
//
//	log.Printf("MockCreateUser passed")
//	return true
//}
