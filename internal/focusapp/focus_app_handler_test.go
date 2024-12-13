package focusapp

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
	"github.com/juanmabm/focusio-core/pkg/argocdclient"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func setup(t *testing.T) (*gomock.Controller, *mocks.MockFocusAppRepository, *mocks.MockFocusCatalogItemRepository, *argocdclient.MockArgoCDClient, *gin.Engine) {

	ctrl := gomock.NewController(t)
	fr := mocks.NewMockFocusAppRepository(ctrl)
	cr := mocks.NewMockFocusCatalogItemRepository(ctrl)
	ac := argocdclient.NewMockArgoCDClient(ctrl)
	r := gin.Default()
	RegisterHandlers(r, fr, cr, ac)

	return ctrl, fr, cr, ac, r
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

func TestFindByName_Success(t *testing.T) {

	ctrl, fr, _, _, r := setup(t)
	defer ctrl.Finish()

	fr.EXPECT().FindByName("App1").Return(entity.FocusApp{Name: "App1"}, nil)
	w := performTestRequest(r, "GET", "/app/App1", nil)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "App1")
}

func TestFindByNameUnSuccess(t *testing.T) {

	ctrl, fr, _, _, r := setup(t)
	defer ctrl.Finish()

	fr.EXPECT().FindByName("App1").Return(entity.FocusApp{}, errors.New(""))
	w := performTestRequest(r, "GET", "/app/App1", nil)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestFindAllShouldReturnAllItems(t *testing.T) {

	ctrl, fr, _, _, r := setup(t)
	defer ctrl.Finish()

	expectedApps := []entity.FocusApp{
		{Name: "app1"},
		{Name: "app2"},
		{Name: "app3"},
	}
	fr.EXPECT().FindAll().Return(expectedApps)
	w := performTestRequest(r, "GET", "/app", nil)
	var actualApps []entity.FocusApp
	err := json.Unmarshal(w.Body.Bytes(), &actualApps)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, len(expectedApps), len(actualApps))
}

func TestDeleteAppShoulInvokeDeleteOne(t *testing.T) {

	ctrl, fr, _, _, r := setup(t)
	defer ctrl.Finish()

	fr.EXPECT().Delete("app1").Times(1)

	performTestRequest(r, "DELETE", "/app/app1", nil)
}

func TestUpdateShouldInvokeUpdateRepositoryMethod(t *testing.T) {

	ctrl, fr, _, _, r := setup(t)
	defer ctrl.Finish()

	app := entity.FocusApp{Name: "app1"}
	fr.EXPECT().FindByName("app1").Return(entity.FocusApp{}, nil)
	fr.EXPECT().Update("app1", gomock.Any()).Times(1)

	performTestRequest(r, "PUT", "/app/app1", app)
}

func TestUpdateShouldReturnNotFoundWhenAppNotExists(t *testing.T) {

	ctrl, fr, _, _, r := setup(t)
	defer ctrl.Finish()

	app := entity.FocusApp{Name: "app1"}
	fr.EXPECT().FindByName("app1").Return(entity.FocusApp{}, errors.New("Mock error"))

	performTestRequest(r, "PUT", "/app/app1", app)
}

func TestCreateAlreadyAppShouldReturnUnprocesable(t *testing.T) {

	ctrl, fr, _, _, r := setup(t)
	defer ctrl.Finish()

	app := entity.FocusApp{Name: "app1"}
	fr.EXPECT().FindByName("app1").Return(entity.FocusApp{Name: "app1"}, nil)

	w := performTestRequest(r, "POST", "/app", app)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestCreateWithCatalogNotExistsShouldReturnUnprocessable(t *testing.T) {

	ctrl, fr, cr, _, r := setup(t)
	defer ctrl.Finish()

	app := entity.FocusApp{Name: "app1", FocusCatalogItem: "quarkus-api"}
	fr.EXPECT().FindByName("app1").Return(entity.FocusApp{}, nil)
	cr.EXPECT().FindByName("quarkus-api").Return(entity.FocusCatalogItem{}, errors.New("Mock error"))

	w := performTestRequest(r, "POST", "/app", app)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestInsertErrorShoultReturnInternalServerError(t *testing.T) {

	ctrl, fr, cr, _, r := setup(t)
	defer ctrl.Finish()

	app := entity.FocusApp{Name: "app1", FocusCatalogItem: "quarkus-api"}
	fr.EXPECT().FindByName("app1").Return(entity.FocusApp{}, nil)
	cr.EXPECT().FindByName("quarkus-api").Return(entity.FocusCatalogItem{}, nil)
	fr.EXPECT().Insert(app).Return(errors.New("Mock error"))

	w := performTestRequest(r, "POST", "/app", app)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestCreateAppShouldReturnAppCreated(t *testing.T) {

	ctrl, fr, cr, _, r := setup(t)
	defer ctrl.Finish()

	app := entity.FocusApp{Name: "app1", FocusCatalogItem: "quarkus-api"}
	fr.EXPECT().FindByName("app1").Return(entity.FocusApp{}, nil)
	cr.EXPECT().FindByName("quarkus-api").Return(entity.FocusCatalogItem{}, nil)
	fr.EXPECT().Insert(app).Return(nil)

	w := performTestRequest(r, "POST", "/app", app)
	var responseApp entity.FocusApp
	json.Unmarshal(w.Body.Bytes(), &responseApp)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, app, responseApp)
}
