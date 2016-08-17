package ol

import(
	"time"
	
	"github.com/jinzhu/gorm"
)

type Business struct {
	gorm.Model
	Id			uint		`json:"id"`
	Uuid		string		`json:"uuid"`
	Name		string		`json:"name"`
	Address		string		`json:"address"`
	Address2	string		`json:"address2"`
	City		string		`json:"city"`
	State		string		`json:"state"`
	Zip			uint		`json:"zip"`
	Country		string		`json:"country"`
	Phone		uint		`json:"phone"`
	Website		string		`json:"website"`
	Created_At	time.Time   `json:"created_at"`
}

type Businesses []Business