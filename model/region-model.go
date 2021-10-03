package model

import "time"

/*
{
	"rno": "1",
	"region_name": "Andhra Pradesh",
	"active": "11503",
	"positive": "2050324",
	"cured": "2024645",
	"death": "14176",
	"new_active": "11142",
	"new_positive": "2051133",
	"new_cured": "2025805",
	"new_death": "14186",
	"updated_at": "2021-10-03T01:01:13.303+05:30",
	"region_code": "28"
  }
  */

type RegionInfo struct {
	Rno         string `json:"rno" bson:"rno" example:"1"`
	Rname       string `json:"region_name" bson:"region_name" example:"Andhra Pradesh"`
	Active      string `json:"active" bson:"active" example:"11503"`
	Positive    string `json:"positive" bson:"positive" example:"2050324"`
	Cured       string `json:"cured" bson:"cured" example:"2024645"`
	Death       string `json:"death" bson:"death" example:"14176"`
	NewActive   string `json:"new_active" bson:"new_active" example:"11142"`
	NewPositive string `json:"new_positive" bson:"new_positive" example:"20511233"`
	NewCured    string `json:"new_cured" bson:"new_cured" example:"2025805"`
	NewDeath    string `json:"new_death" bson:"new_death" example:"14186"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at" example:"2021-10-03T01:01:13.303+05:30"`
	RegionCode  string `json:"region_code" bson:"region_code" example:"28"`
}

type Region []RegionInfo

func AsRegion(s StateInfo, now time.Time) RegionInfo {

	r := RegionInfo{}
	r.Rno = s.Sno
	r.Rname = s.Sname
	r.Active = s.Active
	r.Death = s.Death
	r.Positive = s.Positive
	r.Cured = s.Cured
	r.NewActive = s.NewActive
	r.NewDeath = s.NewDeath
	r.NewPositive = s.NewPositive
	r.NewCured = s.NewCured
	r.RegionCode = s.StateCode
	r.UpdatedAt = now

	return r
}
