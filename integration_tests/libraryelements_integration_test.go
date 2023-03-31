package integration_tests

import (
	"github.com/esnet/gdg/api"
	"github.com/esnet/grafana-swagger-api-golang/goclient/models"
	"github.com/gosimple/slug"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLibraryElementsCRUD(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	apiClient, _ := initTest(t)
	apiClient.DeleteAllDashboards(api.NewDashboardFilter("", "", ""))
	filtersEntity := api.NewDashboardFilter("", "", "")
	log.Info("Exporting all Library Elements")
	apiClient.ExportLibraryElements(filtersEntity)
	log.Info("Listing all library elements")
	boards := apiClient.ListLibraryElements(filtersEntity)
	log.Infof("Imported %d library elements", len(boards))
	var generalBoard *models.LibraryElementDTO
	var otherBoard *models.LibraryElementDTO
	for ndx, board := range boards {
		log.Infof(board.Name)
		if slug.Make(board.Name) == "dashboard-makeover-extra-cleaning-duty-assignment-today" {
			generalBoard = boards[ndx]
		}
		if slug.Make(board.Name) == "extreme-dashboard-makeover-mac-oven" {
			otherBoard = boards[ndx]
		}
	}
	assert.NotNil(t, otherBoard)
	assert.NotNil(t, generalBoard)
	validateLibraryElement(t, generalBoard, map[string]interface{}{"Name": "Dashboard Makeover - Extra Cleaning Duty Assignment Today",
		"Type": "table", "UID": "T47RSwQnz", "Kind": int64(1)})
	validateLibraryElement(t, otherBoard, map[string]interface{}{"Name": "Extreme Dashboard Makeover - Mac Oven",
		"Type": "stat", "UID": "VvzpJ5X7z", "Kind": int64(1)})

	//Import Library Elements
	log.Info("Importing Library Elements")
	list := apiClient.ImportLibraryElements(filtersEntity)
	assert.Equal(t, len(list), len(boards))
	log.Info("Deleting Library Elements")
	deleteList := apiClient.DeleteAllLibraryElements(filtersEntity)
	assert.Equal(t, len(deleteList), len(boards))
	log.Info("List Dashboards again")
	boards = apiClient.ListLibraryElements(filtersEntity)
	assert.Equal(t, len(boards), 0)

}

func validateLibraryElement(t *testing.T, board *models.LibraryElementDTO, data map[string]interface{}) {
	assert.Equal(t, board.Name, data["Name"].(string))
	assert.Equal(t, board.Type, data["Type"].(string))
	assert.Equal(t, board.UID, data["UID"].(string))
	assert.Equal(t, board.Kind, data["Kind"].(int64))

}