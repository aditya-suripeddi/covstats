package repository

import "github.com/aditya-suripeddi/covstats/model"

// RegionInfoRepository interface is a list of methods of RegionInfo model
type RegionInfoRepository interface {
	Save(*model.RegionInfo) error
	Update(string, *model.RegionInfo) error
	Delete(string) error
	FindByRegion(string) (*model.RegionInfo, error)
	FindAll() (model.Region, error)
}
