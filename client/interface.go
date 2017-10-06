package client

import (
	"github.com/quintilesims/layer0/common/models"
)

type Client interface {
	CreateEnvironment(req models.CreateEnvironmentRequest) (string, error)
	DeleteEnvironment(environmentID string) (string, error)
	ListEnvironments() ([]*models.EnvironmentSummary, error)
	ReadEnvironment(environmentID string) (*models.Environment, error)
	UpdateEnvironment(req models.UpdateEnvironmentRequest) (string, error)

	CreateDeploy(req models.CreateDeployRequest) (string, error)
	DeleteDeploy(deployID string) (string, error)
	ListDeploys() ([]*models.DeploySummary, error)
	ReadDeploy(deployID string) (*models.Deploy, error)

	DeleteJob(jobID string) error
	ReadJob(jobID string) (*models.Job, error)
	ListJobs() ([]*models.Job, error)

	CreateLoadBalancer(req models.CreateLoadBalancerRequest) (string, error)
	DeleteLoadBalancer(loadBalancerID string) (string, error)
	ListLoadBalancers() ([]*models.LoadBalancerSummary, error)
	ReadLoadBalancer(loadBalancerID string) (*models.LoadBalancer, error)
	UpdateLoadBalancer(req models.UpdateLoadBalancerRequest) (string, error)

	CreateService(req models.CreateServiceRequest) (string, error)
	DeleteService(serviceID string) (string, error)
	ListServices() ([]*models.ServiceSummary, error)
	ReadService(serviceID string) (*models.Service, error)
	UpdateService(req models.UpdateServiceRequest) (string, error)
}
