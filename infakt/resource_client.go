package infakt

import (
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceClient() *schema.Resource {
	return &schema.Resource{
		Create: resourceClientCreate,
		Read:   resourceClientRead,
		Update: resourceClientUpdate,
		Delete: resourceClientDelete,

		Schema: map[string]*schema.Schema{
			"uuid_count": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceClientCreate(d *schema.ResourceData, m interface{}) error {
	uuid_count := d.Get("uuid_count").(string)

	d.SetId(uuid_count)

	// https://www.uuidtools.com/api/generate/v1/count/uuid_count
	resp, err := http.Get("https://www.uuidtools.com/api/generate/v1/count/" + uuid_count)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	return resourceClientRead(d, m)
}

func resourceClientRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceClientUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceClientRead(d, m)
}

func resourceClientDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
