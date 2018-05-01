package system

import (
	"log"
	"testing"
	"time"

	"github.com/quintilesims/layer0/common/testutils"
	"github.com/quintilesims/layer0/tests/clients"
)

// Test Resources:
// This test creates an environment named 'dsr' that has a
// SystemTestService named 'sts'
func TestDeadServiceRecreated(t *testing.T) {
	t.Parallel()

	s := NewSystemTest(t, "cases/dead_service_recreated", nil)
	s.Terraform.Apply()
	defer s.Terraform.Destroy()

	serviceID := s.Terraform.Output("service_id")
	serviceURL := s.Terraform.Output("service_url")

	Client := clients.NewSTSTestClient(t, serviceURL)
	Client.WaitForHealthy(time.Minute * 3)
	Client.SetHealth("die")

	testutils.WaitFor(t, time.Second*10, time.Minute, func() bool {
		log.Printf("[DEBUG] Waiting for service to die")
		service := s.Layer0.ReadService(serviceID)
		return service.RunningCount == 0
	})

	testutils.WaitFor(t, time.Second*10, time.Minute*2, func() bool {
		log.Printf("[DEBUG] Waiting for service to recreate")
		service := s.Layer0.ReadService(serviceID)
		return service.RunningCount == 1
	})
}
