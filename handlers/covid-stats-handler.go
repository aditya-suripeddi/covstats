package handlers

import (
	"covstats/helpers/wrapper"
	"covstats/model"
	"covstats/repository"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CovidStatsHandler struct {
	regionInfoRepo repository.RegionInfoRepository
}

// NewUserHandler is constructor
func NewCovidStatsHandler(e *echo.Echo, repo repository.RegionInfoRepository) {
	cshandler := &CovidStatsHandler{
		regionInfoRepo: repo,
	}

	e.GET("/states", cshandler.CovidStats)
}

// CovidStats - handler method for binding JSON body and scraping for statewise covid data
func (cshandler *CovidStatsHandler) CovidStats(c echo.Context) error {

	URL := "https://www.mohfw.gov.in/data/datanew.json"
	var client http.Client
	var state_data model.State
	var region_data model.Region

	respo, erro := client.Get(URL)

	if erro != nil {
		log.Fatal(erro)
		return wrapper.Error(http.StatusInternalServerError, erro.Error(), c)
	}

	defer respo.Body.Close()

	if respo.StatusCode != http.StatusOK {
		errorMessage := fmt.Sprintf("Received %s http status code from mohfw server", respo.Status)
		return wrapper.Error(http.StatusInternalServerError, errorMessage, c)
	}

	bodyBytes, err := ioutil.ReadAll(respo.Body)

	if err != nil {
		log.Fatal(err)
		return wrapper.Error(http.StatusInternalServerError, err.Error(), c)
	}

	json.Unmarshal([]byte(bodyBytes), &state_data)

	// As per mohfw data set  { "sname" : "" } respresents covid stats of India
	for _, value := range state_data {
		if value.Sname == "" {
			value.Sname = "India"
			value.Sno = "-1"
			value.StateCode = "-1"
		}

		region_info := model.AsRegion(value)
		region_data = append(region_data, region_info)
		err := cshandler.regionInfoRepo.Save(&region_info)

		if err != nil {
			log.Fatal(err)
			return wrapper.Error(http.StatusInternalServerError, err.Error(), c)
		}
	}

	return wrapper.Data(http.StatusOK, region_data, "Data Source mohfw.gov.in ", c)

}
