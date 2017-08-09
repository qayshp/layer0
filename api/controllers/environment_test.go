package controllers

import (
	"net/http/httptest"
	"testing"

	restful "github.com/emicklei/go-restful"
	"github.com/golang/mock/gomock"
	"github.com/quintilesims/layer0/api/entity/mock_entity"
	"github.com/quintilesims/layer0/api/scheduler"
	"github.com/quintilesims/layer0/common/job"
	"github.com/quintilesims/layer0/common/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateEnvironment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := models.CreateEnvironmentRequest{
		EnvironmentName: "env",
		InstanceSize:    "m3.medium",
		MinClusterCount: 1,
		OperatingSystem: "linux",
		AMIID:           "ami123",
	}

	environmentModel := models.Environment{
		EnvironmentID:   "e1",
		EnvironmentName: "env",
		InstanceSize:    "m3.medium",
		ClusterCount:    1,
		SecurityGroupID: "sg1",
		OperatingSystem: "linux",
		AMIID:           "ami123",
		Links:           []string{"e2"},
	}

	mockEnvironment := mock_entity.NewMockEnvironment(ctrl)
	mockEnvironment.EXPECT().
		Create(req).
		Return(nil)

	mockEnvironment.EXPECT().
		Model().
		Return(&environmentModel, nil)

	mockProvider := mock_entity.NewMockProvider(ctrl)
	mockProvider.EXPECT().
		GetEnvironment("").
		Return(mockEnvironment)

	controller := NewEnvironmentController(mockProvider, nil)

	recorder := httptest.NewRecorder()
	controller.CreateEnvironment(newRequest(t, req, nil), restful.NewResponse(recorder))

	var response models.Environment
	unmarshalBody(t, recorder.Body.Bytes(), &response)

	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, environmentModel, response)
}

func TestDeleteEnvironment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockJobScheduler := scheduler.ScheduleJobFunc(func(req models.CreateJobRequest) (string, error) {
		assert.Equal(t, job.DeleteEnvironmentJob, req.JobType)
		assert.Equal(t, "e1", req.Request)

		return "j1", nil
	})

	controller := NewEnvironmentController(nil, mockJobScheduler)

	recorder := httptest.NewRecorder()
	request := newRequest(t, nil, map[string]string{"id": "e1"})
	controller.DeleteEnvironment(request, restful.NewResponse(recorder))

	assert.Equal(t, 202, recorder.Code)
	assert.Equal(t, "j1", recorder.HeaderMap.Get("X-JobID"))
	assert.Equal(t, "/job/j1", recorder.HeaderMap.Get("Location"))
}

func TestGetEnvironment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	environmentModel := models.Environment{
		EnvironmentID:   "e1",
		EnvironmentName: "env",
		InstanceSize:    "m3.medium",
		ClusterCount:    1,
		SecurityGroupID: "sg1",
		OperatingSystem: "linux",
		AMIID:           "ami123",
		Links:           []string{"e2"},
	}

	mockEnvironment := mock_entity.NewMockEnvironment(ctrl)
	mockEnvironment.EXPECT().
		Model().
		Return(&environmentModel, nil)

	mockProvider := mock_entity.NewMockProvider(ctrl)
	mockProvider.EXPECT().
		GetEnvironment("e1").
		Return(mockEnvironment)

	controller := NewEnvironmentController(mockProvider, nil)

	recorder := httptest.NewRecorder()
	request := newRequest(t, nil, map[string]string{"id": "e1"})
	controller.GetEnvironment(request, restful.NewResponse(recorder))

	var response models.Environment
	unmarshalBody(t, recorder.Body.Bytes(), &response)

	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, environmentModel, response)
}

func TestListEnvironments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	environmentSummaries := []models.EnvironmentSummary{
		{
			EnvironmentID:   "e1",
			EnvironmentName: "env1",
			OperatingSystem: "linux",
		},
		{
			EnvironmentID:   "e2",
			EnvironmentName: "env2",
			OperatingSystem: "windows",
		},
	}

	mockProvider := mock_entity.NewMockProvider(ctrl)
	mockProvider.EXPECT().
		ListEnvironmentIDs().
		Return([]string{"e1", "e2"}, nil)

	for i := range environmentSummaries {
		summary := environmentSummaries[i]

		mockEnvironment := mock_entity.NewMockEnvironment(ctrl)
		mockEnvironment.EXPECT().
			Summary().
			Return(&summary, nil)

		mockProvider.EXPECT().
			GetEnvironment(summary.EnvironmentID).
			Return(mockEnvironment)
	}

	controller := NewEnvironmentController(mockProvider, nil)

	recorder := httptest.NewRecorder()
	controller.ListEnvironments(newRequest(t, nil, nil), restful.NewResponse(recorder))

	var response []models.EnvironmentSummary
	unmarshalBody(t, recorder.Body.Bytes(), &response)

	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, environmentSummaries, response)
}
