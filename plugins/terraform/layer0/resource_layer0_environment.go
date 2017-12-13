package layer0

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/quintilesims/layer0/client"
	"github.com/quintilesims/layer0/common/config"
	"github.com/quintilesims/layer0/common/errors"
	"github.com/quintilesims/layer0/common/models"
)

func resourceLayer0Environment() *schema.Resource {
	return &schema.Resource{
		Create: resourceLayer0EnvironmentCreate,
		Read:   resourceLayer0EnvironmentRead,
		Update: resourceLayer0EnvironmentUpdate,
		Delete: resourceLayer0EnvironmentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"instance_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  config.DefaultEnvironmentInstanceType,
				ForceNew: true,
			},
			"min_scale": {
				Type:     schema.TypeInt,
				Default:  0,
				Optional: true,
			},
			"max_scale": {
				Type:     schema.TypeInt,
				Default:  config.DefaultEnvironmentMaxScale,
				Optional: true,
			},
			"user_data": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"os": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  config.DefaultEnvironmentOS,
				ForceNew: true,
			},
			"ami": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"current_scale": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"security_group_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceLayer0EnvironmentCreate(d *schema.ResourceData, meta interface{}) error {
	apiClient := meta.(client.Client)

	req := models.CreateEnvironmentRequest{
		EnvironmentName:  d.Get("name").(string),
		InstanceType:     d.Get("instance_type").(string),
		UserDataTemplate: []byte(d.Get("user_data").(string)),
		MinScale:         d.Get("min_scale").(int),
		MaxScale:         d.Get("max_scale").(int),
		OperatingSystem:  d.Get("os").(string),
		AMIID:            d.Get("ami").(string),
	}

	jobID, err := apiClient.CreateEnvironment(req)
	if err != nil {
		return err
	}

	job, err := client.WaitForJob(apiClient, jobID, config.DefaultTimeout)
	if err != nil {
		return err
	}

	environmentID := job.Result
	d.SetId(environmentID)

	return resourceLayer0EnvironmentRead(d, meta)
}

func resourceLayer0EnvironmentRead(d *schema.ResourceData, meta interface{}) error {
	apiClient := meta.(client.Client)
	environmentID := d.Id()

	environment, err := apiClient.ReadEnvironment(environmentID)
	if err != nil {
		if err, ok := err.(*errors.ServerError); ok && err.Code == errors.EnvironmentDoesNotExist {
			d.SetId("")
			log.Printf("[WARN] Error Reading Environment (%s), environment does not exist", environmentID)
			return nil
		}

		return err
	}

	d.Set("name", environment.EnvironmentName)
	d.Set("instance_type", environment.InstanceType)
	d.Set("min_scale", environment.MinScale)
	d.Set("current_scale", environment.CurrentScale)
	d.Set("max_scale", environment.MaxScale)
	d.Set("security_group_id", environment.SecurityGroupID)
	d.Set("os", environment.OperatingSystem)
	d.Set("ami", environment.AMIID)

	return nil
}

func resourceLayer0EnvironmentUpdate(d *schema.ResourceData, meta interface{}) error {
	apiClient := meta.(client.Client)
	environmentID := d.Id()

	req := models.UpdateEnvironmentRequest{}
	if d.HasChange("min_scale") {
		minScale := d.Get("min_scale").(int)
		req.MinScale = &minScale
	}

	if d.HasChange("max_scale") {
		maxScale := d.Get("max_scale").(int)
		req.MaxScale = &maxScale
	}

	if req.MinScale != nil || req.MaxScale != nil {
		jobID, err := apiClient.UpdateEnvironment(environmentID, req)
		if err != nil {
			return err
		}

		if _, err := client.WaitForJob(apiClient, jobID, config.DefaultTimeout); err != nil {
			return err
		}
	}

	return resourceLayer0EnvironmentRead(d, meta)
}

func resourceLayer0EnvironmentDelete(d *schema.ResourceData, meta interface{}) error {
	apiClient := meta.(client.Client)
	environmentID := d.Id()

	jobID, err := apiClient.DeleteEnvironment(environmentID)
	if err != nil {
		return err
	}

	if _, err := client.WaitForJob(apiClient, jobID, config.DefaultTimeout); err != nil {
		if err, ok := err.(*errors.ServerError); ok && err.Code == errors.EnvironmentDoesNotExist {
			return nil
		}

		return err
	}

	return nil
}
