package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/aditya-suripeddi/covstats/helpers/wrapper"
	"github.com/aditya-suripeddi/covstats/model"
	"github.com/aditya-suripeddi/covstats/repository"

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

// @Summary Get Covid Stats for all States in India from mohfw
// @Tags root
// @Accept application/json
// @Produce json
// @Success 200 {object}  wrapper.Props{Data=model.Region}
// @Failure 500 {object}  wrapper.Props{code=int,Data=string,Success=boolean}
// @Router /states [get]
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
	now := time.Now()

	
	// As per mohfw data set  { "sname" : "" } respresents covid stats of India
	for _, value := range state_data {
		if value.Sname == "" {
			value.Sname = "India"
			value.Sno = "-1"
			value.StateCode = "-1"
		}

		region_info := model.AsRegion(value, now)
		region_data = append(region_data, region_info)
		err := cshandler.regionInfoRepo.Save(&region_info)

		if err != nil {
			log.Fatal(err)
			return wrapper.Error(http.StatusInternalServerError, err.Error(), c)
		}
	}

	return wrapper.Data(http.StatusOK, region_data, "Data Source mohfw.gov.in ", c)

}
