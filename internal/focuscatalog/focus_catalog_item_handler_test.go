package focuscatalog

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/juanmabm/focusio-core/internal/entity"
	"github.com/juanmabm/focusio-core/internal/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func setup(t *testing.T) (*gomock.Controller, *mocks.MockFocusCatalogItemRepository, *gin.Engine) {

	ctrl := gomock.NewController(t)
	cr := mocks.NewMockFocusCatalogItemRepository(ctrl)
	r := gin.Default()
	RegisterHandlers(r, cr)

	return ctrl, cr, r
}

func performTestRequest(r *gin.Engine, method string, url string, body any) *httptest.ResponseRecorder {

	requestBody := &bytes.Buffer{}
	if jsonBody, err := json.Marshal(body); err == nil {
		requestBody = bytes.NewBuffer(jsonBody)
	}
	req, _ := http.NewRequest(method, url, requestBody)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return w
}

func TestFindByNameShouldNotFoundIfCatalogNotExists(t *testing.T) {

	ctrl, cr, r := setup(t)
	defer ctrl.Finish()

	cr.EXPECT().FindByName("quarkus-api").Return(entity.FocusCatalogItem{}, errors.New("Mock error"))

	w := performTestRequest(r, "GET", "/catalog/quarkus-api", nil)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestUpdateNotExistingCatalogShouldReturnNotFound(t *testing.T) {

	ctrl, cr, r := setup(t)
	defer ctrl.Finish()

	cr.EXPECT().FindByName("quarkus-api").Return(entity.FocusCatalogItem{}, errors.New("Mock error"))

	w := performTestRequest(r, "PUT", "/catalog/quarkus-api", nil)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestUpdateShouldReturnInternalServerErrorIfDatabaseError(t *testing.T) {

	ctrl, cr, r := setup(t)
	defer ctrl.Finish()

	catalog := entity.FocusCatalogItem{}
	cr.EXPECT().FindByName("quarkus-api").Return(entity.FocusCatalogItem{}, nil)
	cr.EXPECT().Update("quarkus-api", catalog).Return(errors.New("Mock error"))

	w := performTestRequest(r, "PUT", "/catalog/quarkus-api", catalog)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestUpdateShouldReturnOkStatus(t *testing.T) {

	ctrl, cr, r := setup(t)
	defer ctrl.Finish()

	catalog := entity.FocusCatalogItem{}
	cr.EXPECT().FindByName("quarkus-api").Return(entity.FocusCatalogItem{}, nil)
	cr.EXPECT().Update("quarkus-api", catalog).Return(nil)

	w := performTestRequest(r, "PUT", "/catalog/quarkus-api", catalog)

	assert.Equal(t, http.StatusOK, w.Code)
}

// /////
func TestCreateAlreadyCatalogShouldReturnUnprocessable(t *testing.T) {

	ctrl, cr, r := setup(t)
	defer ctrl.Finish()

	catalog := entity.FocusCatalogItem{Name: "quarkus-api"}
	cr.EXPECT().FindByName("quarkus-api").Return(entity.FocusCatalogItem{Name: "quarkus-api"}, nil)

	w := performTestRequest(r, "POST", "/catalog", catalog)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestCreateDatabaseErrorShouldReturnInternalServerError(t *testing.T) {

	ctrl, cr, r := setup(t)
	defer ctrl.Finish()

	catalog := entity.FocusCatalogItem{Name: "quarkus-api"}
	cr.EXPECT().FindByName("quarkus-api").Return(entity.FocusCatalogItem{}, nil)
	cr.EXPECT().Insert(catalog).Return(errors.New("Mock error"))

	w := performTestRequest(r, "POST", "/catalog", catalog)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestCreateShouldReturnResponseCatalog(t *testing.T) {

	ctrl, cr, r := setup(t)
	defer ctrl.Finish()

	catalog := entity.FocusCatalogItem{Name: "quarkus-api"}
	cr.EXPECT().FindByName("quarkus-api").Return(entity.FocusCatalogItem{}, nil)
	cr.EXPECT().Insert(catalog).Return(nil)

	w := performTestRequest(r, "POST", "/catalog", catalog)
	response := entity.FocusCatalogItem{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, catalog, response)
}
