package registry

import (
	"github.com/harness/harness-go-sdk/harness/har"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strings"
)

// the origional implamentation took two differnet but idential parameters, parent_ref and space_ref
// this function will use the origional parameters if they are given, otherwise build it using the standard 
// org_id and project_id parameters found in other standard multi-hierarchy resources
func buildHARPathRef(d *schema.ResourceData, c *har.APIClient) (parentRef string) {
	if attr, ok := d.GetOk("parent_ref"); ok {
		parentRef = attr.(string)
	} else {
		if attr, ok := d.GetOk("space_ref"); ok {
			parentRef = attr.(string)
		} else {
			result := []string{c.AccountId}
			if org_id, ok := d.GetOk("org_id"); ok {
				result = append(result, org_id.(string))
			}
			if project_id, ok := d.GetOk("project_id"); ok {
				result = append(result, project_id.(string))
			}
			parentRef = strings.Join(result, "/")
		}
	}
	return
}
