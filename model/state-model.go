package model

import "fmt"

/*
{
		"sno": "20",
		"state_name": "Lakshadweep",
		"active": "5",
		"positive": "10361",
		"cured": "10305",
		"death": "51",
		"new_active": "5",
		"new_positive": "10361",
		"new_cured": "10305",
		"new_death": "51",
		"state_code": "31"
	}
*/

// State and it's json body from mohfw can be put in models package
// in the end to organize the project
type StateInfo struct {
	Sno         string `json:"sno"`
	Sname       string `json:"state_name"`
	Active      string `json:"active"`
	Positive    string `json:"positive"`
	Cured       string `json:"cured"`
	Death       string `json:"death"`
	NewActive   string `json:"new_active"`
	NewPositive string `json:"new_positive"`
	NewCured    string `json:"new_cured"`
	NewDeath    string `json:"new_death"`
	StateCode   string `json:"state_code"`
}

type State []StateInfo


func (s StateInfo) ToString() string {

	return fmt.Sprintf(  "Sno: %s\nSname : %s\nActive: %s\nPositive: %s\nCured: %s" +
	                     "\nDeath: %s\nNewActive: %s \nNewPositive: %s\nNewCured: %s" +
						 "\nNewDeath: %s\nStateCode: %s\n",
	                 	  s.Sno, s.Sname, s.Active, s.Positive, s.Cured,
						  s.Death, s.NewActive,
		                  s.NewPositive, s.NewCured, s.NewDeath, s.StateCode)

}
