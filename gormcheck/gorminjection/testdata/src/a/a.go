package a

import (
	gorm2 "github.com/jinzhu/gorm"
	gorm1 "gorm.io/gorm"
)

func OK() {
	(*gorm1.DB)(nil).Find(nil, "1=1", nil) // OK
	(*gorm2.DB)(nil).Find(nil, "1=1", nil) // OK

	(*gorm1.DB)(nil).First(nil, "1==1", nil) // OK
	(*gorm2.DB)(nil).First(nil, "1==1", nil) // OK

	(*gorm1.DB)(nil).Last(nil, "1==1", nil) // OK
	(*gorm2.DB)(nil).Last(nil, "1==1", nil) // OK

	(*gorm1.DB)(nil).Delete(nil, "1==1", nil) // OK
	(*gorm2.DB)(nil).Delete(nil, "1==1", nil) // OK

	(*gorm1.DB)(nil).Take(nil, "1==1", nil) // OK
	(*gorm2.DB)(nil).Take(nil, "1==1", nil) // OK

	(*gorm1.DB)(nil).Find(nil, struct{}{}) // OK
	(*gorm2.DB)(nil).Find(nil, struct{}{}) // OK

	(*gorm1.DB)(nil).First(nil, struct{}{}) // OK
	(*gorm2.DB)(nil).First(nil, struct{}{}) // OK

	(*gorm1.DB)(nil).Last(nil, struct{}{}) // OK
	(*gorm2.DB)(nil).Last(nil, struct{}{}) // OK

	(*gorm1.DB)(nil).Delete(nil, struct{}{}) // OK
	(*gorm2.DB)(nil).Delete(nil, struct{}{}) // OK

	(*gorm1.DB)(nil).Take(nil, struct{}{}) // OK
	(*gorm2.DB)(nil).Take(nil, struct{}{}) // OK

}

func NG() {
	(*gorm1.DB)(nil).Find(nil, "1==1") // want "it may be SQL injection"
	(*gorm2.DB)(nil).Find(nil, "1==1") // want "it may be SQL injection"

	(*gorm1.DB)(nil).First(nil, "1==1") // want "it may be SQL injection"
	(*gorm2.DB)(nil).First(nil, "1==1") // want "it may be SQL injection"

	(*gorm1.DB)(nil).Last(nil, "1==1") // want "it may be SQL injection"
	(*gorm2.DB)(nil).Last(nil, "1==1") // want "it may be SQL injection"

	(*gorm1.DB)(nil).Delete(nil, "1==1") // want "it may be SQL injection"
	(*gorm2.DB)(nil).Delete(nil, "1==1") // want "it may be SQL injection"

	(*gorm1.DB)(nil).Take(nil, "1==1") // want "it may be SQL injection"
	(*gorm2.DB)(nil).Take(nil, "1==1") // want "it may be SQL injection"

}
