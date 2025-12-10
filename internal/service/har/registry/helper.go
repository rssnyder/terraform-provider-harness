package registry

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func buildHARPathRef(d *schema.ResourceData, accountID string) (registryRef string) {
	registryRef = accountID
	if attr, ok := d.GetOk("org_id"); ok {
		registryRef += "/" + attr.(string)
	}
	if attr, ok := d.GetOk("project_id"); ok {
		registryRef += "/" + attr.(string)
	}
	return
}