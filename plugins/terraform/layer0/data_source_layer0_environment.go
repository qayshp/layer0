package layer0

import (
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/quintilesims/layer0/client"
)

func dataSourceLayer0Environment() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLayer0EnvironmentRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"instance_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"min_scale": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"current_scale": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"max_scale": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"os": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ami": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"security_group_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceLayer0EnvironmentRead(d *schema.ResourceData, meta interface{}) error {
	apiClient := meta.(client.Client)

	query := url.Values{}
	query.Set(client.TagQueryParamType, "environment")
	query.Set(client.TagQueryParamName, d.Get("name").(string))

	tags, err := apiClient.ListTags(query)
	if err != nil {
		return err
	}

	if len(tags) == 0 {
		return fmt.Errorf("No environment found matching %#v", query)
	}

	environmentID := tags[0].EntityID
	environment, err := apiClient.ReadEnvironment(environmentID)
	if err != nil {
		return err
	}

	d.SetId(environment.EnvironmentID)
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