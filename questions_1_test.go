package stream

import (
	"fmt"
	"testing"

	"github.com/Pallinder/go-randomdata"
	"github.com/sumory/idgen"
)

func RandomInstance() *Employee {
	_, id := idGen.NextId()
	country := randomdata.Country(randomdata.FullCountry)
	province := randomdata.ProvinceForCountry(country)
	city := randomdata.City()
	age := randomdata.Number(18, 100)
	name := randomdata.SillyName()
	phone := ""
	if randomdata.Boolean() {
		phone = randomdata.PhoneNumber()
	}
	return &Employee{
		Id:    id,
		Name:  &name,
		Age:   &age,
		Phone: &phone,
		Position: &PositionInfo{
			Province: &province,
			Country:  &country,
			City:     &city,
		},
	}
}

var TestEmployees []*Employee

func GenTestData() {
	for i:=0;i<10;i++ {
		TestEmployees = append(TestEmployees, RandomInstance())
	}
}

func TestQuestion1Sub1(t *testing.T) {
	GenTestData()
	for _, e := range TestEmployees {
		fmt.Printf("value of e %+v", e)
	}
}
