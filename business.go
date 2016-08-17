package main

import "time"

type Business struct {
	Id			int64		`gorm:"column:id" json:"id"`
	Uuid		string		`gorm:"column:uuid" json:"uuid"`
	Name		string		`gorm:"column:name" json:"name"`
	Address		string		`gorm:"column:address" json:"address"`
	Address2	string		`gorm:"column:address2" json:"address2"`
	City		string		`gorm:"column:city" json:"city"`
	State		string		`gorm:"column:state" json:"state"`
	Zip			uint		`gorm:"column:zip" json:"zip"`
	Country		string		`gorm:"column:country" json:"country"`
	Phone		uint		`gorm:"column:phone" json:"phone"`
	Website		string		`gorm:"column:website" json:"website"`
	Created_At	time.Time  	`gorm:"column:created_at" json:"created_at"`
}

type Businesses []Business