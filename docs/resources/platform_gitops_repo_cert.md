---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_platform_gitops_repo_cert Resource - terraform-provider-harness"
subcategory: "Next Gen"
description: |-
  Resource for creating a Harness Gitops Repositories Certificates.
---

# harness_platform_gitops_repo_cert (Resource)

Resource for creating a Harness Gitops Repositories Certificates.

## Example Usage

```terraform
resource "harness_platform_gitops_repo_cert" "example" {
  account_id = "account_id"
  agent_id   = "agent_id"

  request {
    upsert = true
    certificates {
      metadata {

      }
      items {
        server_name = "serverName"
        cert_type   = "https"
        cert_data   = "yourcertdata"
      }
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String) account identifier of the Repository Certificates.
- `agent_id` (String) agent identifier of the Repository Certificates.
- `request` (Block List, Min: 1) Repository Certificates create/Update request. (see [below for nested schema](#nestedblock--request))

### Optional

- `org_id` (String) organization identifier of the Repository Certificates.
- `project_id` (String) project identifier of the Repository Certificates.

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--request"></a>
### Nested Schema for `request`

Optional:

- `certificates` (Block List) certificates details. (see [below for nested schema](#nestedblock--request--certificates))
- `upsert` (Boolean) if the Repository Certificates should be upserted.

<a id="nestedblock--request--certificates"></a>
### Nested Schema for `request.certificates`

Optional:

- `items` (Block List) List of certificates to be processed. (see [below for nested schema](#nestedblock--request--certificates--items))
- `metadata` (Block List) metadata details (see [below for nested schema](#nestedblock--request--certificates--metadata))

<a id="nestedblock--request--certificates--items"></a>
### Nested Schema for `request.certificates.items`

Optional:

- `cert_data` (String) CertData contains the actual certificate data, dependent on the certificate type.
- `cert_info` (String) CertInfo will hold additional certificate info, depdendent on the certificate type .
- `cert_sub_type` (String) CertSubType specifies the sub type of the cert, i.e. ssh-rsa.
- `cert_type` (String) CertType specifies the type of the certificate - currently one of https or ssh.
- `server_name` (String) ServerName specifies the DNS name of the server this certificate is intended.


<a id="nestedblock--request--certificates--metadata"></a>
### Nested Schema for `request.certificates.metadata`

Optional:

- `continue` (String) continue may be set if the user set a limit on the number of items returned.
- `remaining_item_count` (String) subsequent items in the list.
- `resource_version` (String) dentifies the server's internal version.
- `self_link` (String) selfLink is a URL representing this object.

## Import

Import is supported using the following syntax:

```shell
# Import a Account level Gitops Repository Certificate
terraform import harness_platform_gitops_repo_cert.example <repocert_id>
```
