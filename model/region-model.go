package model

import "time"

type RegionInfo struct {
	Rno         string `json:"rno" bson:"rno"`
	Rname       string `json:"region_name" bson:"region_name"`
	Active      string `json:"active" bson:"active"`
	Positive    string `json:"positive" bson:"positive"`
	Cured       string `json:"cured" bson:"cured"`
	Death       string `json:"death" bson:"death"`
	NewActive   string `json:"new_active" bson:"new_active"`
	NewPositive string `json:"new_positive" bson:"new_positive"`
	NewCured    string `json:"new_cured" bson:"new_cured"`
	NewDeath    string `json:"new_death" bson:"new_death"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
	RegionCode  string `json:"region_code" bson:"region_code"`
}

type Region []RegionInfo

func AsRegion(s StateInfo) RegionInfo {

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
	r.UpdatedAt = time.Now()

	return r
}
