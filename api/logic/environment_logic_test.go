package logic

import (
	"testing"

	"github.com/quintilesims/layer0/api/backend/ecs/id"
	"github.com/quintilesims/layer0/common/models"
	"github.com/quintilesims/layer0/common/testutils"
	"github.com/stretchr/testify/assert"
)

func TestGetEnvironment(t *testing.T) {
	testLogic, ctrl := NewTestLogic(t)
	defer ctrl.Finish()

	retEnvironment := &models.Environment{
		EnvironmentID: "e1",
	}

	testLogic.Backend.EXPECT().
		GetEnvironment("e1").
		Return(retEnvironment, nil)

	testLogic.AddTags(t, []*models.Tag{
		{EntityID: "e1", EntityType: "environment", Key: "name", Value: "env"},
		{EntityID: "e1", EntityType: "environment", Key: "os", Value: "linux"},
		{EntityID: "e1", EntityType: "environment", Key: "link", Value: "e2"},
		{EntityID: "extra", EntityType: "environment", Key: "name", Value: "extra"},
	})

	environmentLogic := NewL0EnvironmentLogic(testLogic.Logic())
	received, err := environmentLogic.GetEnvironment("e1")
	if err != nil {
		t.Fatal(err)
	}

	expected := &models.Environment{
		EnvironmentID:   "e1",
		EnvironmentName: "env",
		OperatingSystem: "linux",
		Links:           []string{"e2"},
	}

	testutils.AssertEqual(t, received, expected)
}

func TestListEnvironments(t *testing.T) {
	testLogic, ctrl := NewTestLogic(t)
	defer ctrl.Finish()

	ecsEnvironmentIDs := []id.ECSEnvironmentID{
		id.L0EnvironmentID("env_id1").ECSEnvironmentID(),
		id.L0EnvironmentID("env_id2").ECSEnvironmentID(),
	}

	testLogic.Backend.EXPECT().
		ListEnvironments().
		Return(ecsEnvironmentIDs, nil)

	testLogic.AddTags(t, []*models.Tag{
		{EntityID: "env_id1", EntityType: "environment", Key: "name", Value: "env_name1"},
		{EntityID: "env_id1", EntityType: "environment", Key: "os", Value: "linux"},
		{EntityID: "env_id2", EntityType: "environment", Key: "name", Value: "env_name2"},
		{EntityID: "env_id2", EntityType: "environment", Key: "os", Value: "windows"},
		{EntityID: "extra", EntityType: "environment", Key: "name", Value: "extra"},
	})

	environmentLogic := NewL0EnvironmentLogic(testLogic.Logic())
	result, err := environmentLogic.ListEnvironments()
	if err != nil {
		t.Fatal(err)
	}

	expected := []models.EnvironmentSummary{
		{
			EnvironmentID:   "env_id1",
			EnvironmentName: "env_name1",
			OperatingSystem: "linux",
		},
		{
			EnvironmentID:   "env_id2",
			EnvironmentName: "env_name2",
			OperatingSystem: "windows",
		},
	}

	assert.Equal(t, expected, result)
}

func TestDeleteEnvironment(t *testing.T) {
	testLogic, ctrl := NewTestLogic(t)
	defer ctrl.Finish()

	testLogic.Backend.EXPECT().
		DeleteEnvironment("eid1").
		Return(nil)

	testLogic.Backend.EXPECT().
		DeleteEnvironmentLink("eid1", "eid2").
		Return(nil)

	testLogic.AddTags(t, []*models.Tag{
		{EntityID: "eid1", EntityType: "environment", Key: "name", Value: "env"},
		{EntityID: "eid1", EntityType: "environment", Key: "os", Value: "linux"},
		{EntityID: "eid1", EntityType: "environment", Key: "link", Value: "eid2"},
		{EntityID: "eid2", EntityType: "environment", Key: "link", Value: "eid1"},
		{EntityID: "extra", EntityType: "environment", Key: "name", Value: "extra"},
	})

	environmentLogic := NewL0EnvironmentLogic(testLogic.Logic())
	if err := environmentLogic.DeleteEnvironment("eid1"); err != nil {
		t.Fatal(err)
	}

	tags, err := testLogic.TagStore.SelectByType("environment")
	if err != nil {
		t.Fatal(err)
	}

	// make sure the 'extra' tag is the only one left
	testutils.AssertEqual(t, len(tags), 1)
}

func TestCanCreateEnvironment(t *testing.T) {
	testLogic, ctrl := NewTestLogic(t)
	defer ctrl.Finish()

	testLogic.AddTags(t, []*models.Tag{
		{EntityID: "e1", EntityType: "environment", Key: "name", Value: "env_1"},
		{EntityID: "e2", EntityType: "environment", Key: "name", Value: "env_2"},
		{EntityID: "extra", EntityType: "environment", Key: "name", Value: "extra"},
	})

	environmentLogic := NewL0EnvironmentLogic(testLogic.Logic())

	cases := map[string]bool{
		"env_1":  false,
		"env_2":  false,
		"env3":   true,
		"env_12": true,
		"env":    true,
	}

	for name, expected := range cases {
		request := models.CreateEnvironmentRequest{EnvironmentName: name}

		received, err := environmentLogic.CanCreateEnvironment(request)
		if err != nil {
			t.Fatal(err)
		}

		if received != expected {
			t.Errorf("Failure on case '%s': response was %v, expected %v", name, received, expected)
		}
	}
}

func TestCreateEnvironment(t *testing.T) {
	testLogic, ctrl := NewTestLogic(t)
	defer ctrl.Finish()

	retEnvironment := &models.Environment{
		EnvironmentID: "e1",
	}

	testLogic.Backend.EXPECT().
		CreateEnvironment("name", "m3.medium", "linux", "amiid", 2, []byte("user_data")).
		Return(retEnvironment, nil)

	request := models.CreateEnvironmentRequest{
		EnvironmentName:  "name",
		InstanceSize:     "m3.medium",
		OperatingSystem:  "linux",
		AMIID:            "amiid",
		MinClusterCount:  2,
		UserDataTemplate: []byte("user_data"),
	}

	environmentLogic := NewL0EnvironmentLogic(testLogic.Logic())
	received, err := environmentLogic.CreateEnvironment(request)
	if err != nil {
		t.Fatal(err)
	}

	expected := &models.Environment{
		EnvironmentID:   "e1",
		EnvironmentName: "name",
		OperatingSystem: "linux",
		Links:           []string{},
	}

	testutils.AssertEqual(t, received, expected)
	testLogic.AssertTagExists(t, models.Tag{EntityID: "e1", EntityType: "environment", Key: "name", Value: "name"})
	testLogic.AssertTagExists(t, models.Tag{EntityID: "e1", EntityType: "environment", Key: "os", Value: "linux"})
}

func TestCreateEnvironmentError_missingRequiredParams(t *testing.T) {
	testLogic, ctrl := NewTestLogic(t)
	defer ctrl.Finish()

	environmentLogic := NewL0EnvironmentLogic(testLogic.Logic())

	cases := map[string]models.CreateEnvironmentRequest{
		"Missing EnvironmentName": {
			OperatingSystem: "linux",
		},
		"Missing OperatingSystem": {
			EnvironmentName: "name",
		},
	}

	for name, request := range cases {
		if _, err := environmentLogic.CreateEnvironment(request); err == nil {
			t.Errorf("Case %s: error was nil!", name)
		}
	}
}

func TestUpdateEnvironment(t *testing.T) {
	testLogic, ctrl := NewTestLogic(t)
	defer ctrl.Finish()

	retEnvironment := &models.Environment{
		EnvironmentID: "e1",
	}

	testLogic.Backend.EXPECT().
		UpdateEnvironment("e1", 2).
		Return(retEnvironment, nil)

	testLogic.AddTags(t, []*models.Tag{
		{EntityID: "e1", EntityType: "environment", Key: "name", Value: "env"},
		{EntityID: "extra", EntityType: "environment", Key: "name", Value: "extra"},
	})

	environmentLogic := NewL0EnvironmentLogic(testLogic.Logic())
	received, err := environmentLogic.UpdateEnvironment("e1", 2)
	if err != nil {
		t.Fatal(err)
	}

	expected := &models.Environment{
		EnvironmentID:   "e1",
		EnvironmentName: "env",
		Links:           []string{},
	}

	testutils.AssertEqual(t, received, expected)
}

func TestCreateEnvironmentLink(t *testing.T) {
	testLogic, ctrl := NewTestLogic(t)
	defer ctrl.Finish()

	testLogic.Backend.EXPECT().
		CreateEnvironmentLink("eid1", "eid2").
		Return(nil)

	environmentLogic := NewL0EnvironmentLogic(testLogic.Logic())
	if err := environmentLogic.CreateEnvironmentLink("eid1", "eid2"); err != nil {
		t.Fatal(err)
	}

	testLogic.AssertTagExists(t, models.Tag{EntityID: "eid1", EntityType: "environment", Key: "link", Value: "eid2"})
	testLogic.AssertTagExists(t, models.Tag{EntityID: "eid2", EntityType: "environment", Key: "link", Value: "eid1"})
}

func TestDeleteEnvironmentLink(t *testing.T) {
	testLogic, ctrl := NewTestLogic(t)
	defer ctrl.Finish()

	testLogic.Backend.EXPECT().
		DeleteEnvironmentLink("eid1", "eid2").
		Return(nil)

	testLogic.AddTags(t, []*models.Tag{
		{EntityID: "eid1", EntityType: "environment", Key: "link", Value: "eid2"},
		{EntityID: "extra", EntityType: "environment", Key: "name", Value: "extra"},
	})

	environmentLogic := NewL0EnvironmentLogic(testLogic.Logic())
	if err := environmentLogic.DeleteEnvironmentLink("eid1", "eid2"); err != nil {
		t.Fatal(err)
	}

	tags, err := testLogic.TagStore.SelectByType("environment")
	if err != nil {
		t.Fatal(err)
	}

	// make sure the 'extra' tag is the only one left
	testutils.AssertEqual(t, len(tags), 1)
}
