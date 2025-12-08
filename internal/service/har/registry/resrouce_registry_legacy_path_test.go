package registry_test

import (
	"fmt"
	"github.com/harness/terraform-provider-harness/internal/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"os"
	"testing"
)

// Virtual Docker Registry
func TestAccResourceVirtualDockerRegistryLegacyPath(t *testing.T) {
	id := fmt.Sprintf("tf_auto_virtual_docker_registry")
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config: testAccResourceVirtualDockerRegistryLegacyPath(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
		},
	})
}
func TestOrgResourceVirtualDockerRegistryLegacyPath(t *testing.T) {
	id := fmt.Sprintf("tf_auto_virtual_docker_registry")
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config: testOrgResourceVirtualDockerRegistryLegacyPath(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
		},
	})
}
func TestProjectResourceVirtualDockerRegistryLegacyPath(t *testing.T) {
	id := fmt.Sprintf("tf_auto_virtual_docker_registry")
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config: testProjResourceVirtualDockerRegistryLegacyPath(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
		},
	})
}

// Upstream Docker Registry --- UserPassword
func TestAccResourceUpstreamDockerRegistryLegacyPath(t *testing.T) {
	id := fmt.Sprintf("tf_auto_upstream_docker_registry")
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config: testAccResourceUpstreamDockerRegistryLegacyPath(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
		},
	})
}
func TestOrgResourceUpstreamDockerRegistryLegacyPath(t *testing.T) {
	id := fmt.Sprintf("tf_auto_upstream_docker_registry")
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config: testOrgResourceUpstreamDockerRegistryLegacyPath(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
		},
	})
}
func TestProjectResourceUpstreamDockerRegistryLegacyPath(t *testing.T) {
	id := fmt.Sprintf("tf_auto_upstream_docker_registry")
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config: testProjResourceUpstreamDockerRegistryLegacyPath(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
		},
	})
}

func testAccResourceVirtualDockerRegistryLegacyPath(id string, accId string) string {
	return fmt.Sprintf(`

 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s"
   package_type = "DOCKER"

   config {
    type = "VIRTUAL"
   }
   parent_ref = "%[2]s"
 }
`, id, accId)
}

func testOrgResourceVirtualDockerRegistryLegacyPath(id string, accId string) string {
	return fmt.Sprintf(`
 resource "harness_platform_organization" "test" {
  identifier = "%[1]s_org"
  name = "%[1]s"
 }

 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s/${harness_platform_organization.test.identifier}"
   package_type = "DOCKER"

   config {
    type = "VIRTUAL"
   }
   parent_ref = "%[2]s/${harness_platform_organization.test.identifier}"
 }
`, id, accId)
}

func testProjResourceVirtualDockerRegistryLegacyPath(id string, accId string) string {
	return fmt.Sprintf(`
 resource "harness_platform_organization" "test" {
  identifier = "%[1]s_org"
  name = "%[1]s"
 }

 resource "harness_platform_project" "test" {
  identifier = "%[1]s_project"
  name = "%[1]s"
  org_id = harness_platform_organization.test.id
  color = "#472848"
 }
 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s/${harness_platform_organization.test.identifier}/${harness_platform_project.test.identifier}"
   package_type = "DOCKER"

   config {
    type = "VIRTUAL"
   }
   parent_ref = "%[2]s/${harness_platform_organization.test.identifier}/${harness_platform_project.test.identifier}"
 }
`, id, accId)
}

func testAccResourceUpstreamDockerRegistryLegacyPath(id string, accId string) string {
	return fmt.Sprintf(`
 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s"
   package_type = "DOCKER"

   config {
		type = "UPSTREAM"
		auth_type = "UserPassword"
		source = "Dockerhub"
		auth {
			auth_type = "UserPassword"
			user_name = "username"
			secret_identifier = "Secret_Token"
			secret_space_path = "%[2]s"
		}
   }
   parent_ref = "%[2]s"
 }
`, id, accId)
}

func testOrgResourceUpstreamDockerRegistryLegacyPath(id string, accId string) string {
	return fmt.Sprintf(`
 resource "harness_platform_organization" "test" {
  identifier = "%[1]s_org"
  name = "%[1]s"
 }

 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s/${harness_platform_organization.test.identifier}"
   package_type = "DOCKER"

   config {
		type = "UPSTREAM"
		auth_type = "UserPassword"
		source = "Dockerhub"
		auth {
			auth_type = "UserPassword"
			user_name = "username"
			secret_identifier = "Secret_Token"
			secret_space_path = "%[2]s"
		}
   }
   parent_ref = "%[2]s/${harness_platform_organization.test.identifier}"
 }
`, id, accId)
}

func testProjResourceUpstreamDockerRegistryLegacyPath(id string, accId string) string {
	return fmt.Sprintf(`
 resource "harness_platform_organization" "test" {
  identifier = "%[1]s_org"
  name = "%[1]s"
 }

 resource "harness_platform_project" "test" {
  identifier = "%[1]s_project"
  name = "%[1]s"
  org_id = harness_platform_organization.test.id
  color = "#472848"
 }
 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s/${harness_platform_organization.test.identifier}/${harness_platform_project.test.identifier}"
   package_type = "DOCKER"

   config {
		type = "UPSTREAM"
		auth_type = "UserPassword"
		source = "Dockerhub"
		auth {
			auth_type = "UserPassword"
			user_name = "username"
			secret_identifier = "Secret_Token"
			secret_space_path = "%[2]s"
		}
   }
   parent_ref = "%[2]s/${harness_platform_organization.test.identifier}/${harness_platform_project.test.identifier}"
 }
`, id, accId)
}

// Virtual Helm Registry
func TestAccResourceVirtualHelmRegistryLegacyPath(t *testing.T) {
	id := fmt.Sprintf("tf_auto_virtual_helm_registry")
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config: testAccResourceVirtualHelmRegistryLegacyPath(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
		},
	})
}
func TestOrgResourceVirtualHelmRegistryLegacyPath(t *testing.T) {
	id := fmt.Sprintf("tf_auto_virtual_helm_registry")
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config: testOrgResourceVirtualHelmRegistryLegacyPath(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
		},
	})
}
func TestProjectResourceVirtualHelmRegistryLegacyPath(t *testing.T) {
	id := fmt.Sprintf("tf_auto_virtual_helm_registry")
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config: testProjResourceVirtualHelmRegistryLegacyPath(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
		},
	})
}

// Upstream Helm Registry --- UserPass
func TestAccResourceUpstreamHelmRegistryLegacyPath(t *testing.T) {
	id := fmt.Sprintf("tf_auto_upstream_helm_registry")
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config: testAccResourceUpstreamHelmRegistryLegacyPath(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
		},
	})
}
func TestOrgResourceUpstreamHelmRegistryLegacyPath(t *testing.T) {
	id := fmt.Sprintf("tf_auto_upstream_helm_registry")
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config: testOrgResourceUpstreamHelmRegistryLegacyPath(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
		},
	})
}
func TestProjectResourceUpstreamHelmRegistryLegacyPath(t *testing.T) {
	id := fmt.Sprintf("tf_auto_upstream_helm_registry")
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config: testProjResourceUpstreamHelmRegistryLegacyPath(id, accountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
		},
	})
}

func testAccResourceVirtualHelmRegistryLegacyPath(id string, accId string) string {
	return fmt.Sprintf(`

 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s"
   package_type = "HELM"

   config {
    type = "VIRTUAL"
   }
   parent_ref = "%[2]s"
 }
`, id, accId)
}

func testOrgResourceVirtualHelmRegistryLegacyPath(id string, accId string) string {
	return fmt.Sprintf(`
 resource "harness_platform_organization" "test" {
  identifier = "%[1]s_org"
  name = "%[1]s"
 }

 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s/${harness_platform_organization.test.identifier}"
   package_type = "HELM"

   config {
    type = "VIRTUAL"
   }
   parent_ref = "%[2]s/${harness_platform_organization.test.identifier}"
 }
`, id, accId)
}

func testProjResourceVirtualHelmRegistryLegacyPath(id string, accId string) string {
	return fmt.Sprintf(`
 resource "harness_platform_organization" "test" {
  identifier = "%[1]s_org"
  name = "%[1]s"
 }

 resource "harness_platform_project" "test" {
  identifier = "%[1]s_project"
  name = "%[1]s"
  org_id = harness_platform_organization.test.id
  color = "#472848"
 }
 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s/${harness_platform_organization.test.identifier}/${harness_platform_project.test.identifier}"
   package_type = "HELM"

   config {
    type = "VIRTUAL"
   }
   parent_ref = "%[2]s/${harness_platform_organization.test.identifier}/${harness_platform_project.test.identifier}"
 }
`, id, accId)
}

func testAccResourceUpstreamHelmRegistryLegacyPath(id string, accId string) string {
	return fmt.Sprintf(`
 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s"
   package_type = "HELM"

   config {
		type = "UPSTREAM"
		auth_type = "UserPassword"
		source = "Custom"
		url = "https://har-registry.default.svc.cluster.local"
		auth {
			auth_type = "UserPassword"
			user_name = "username"
			secret_identifier = "Secret_Token"
			secret_space_path = "%[2]s"
		}
   }
   parent_ref = "%[2]s"
 }
`, id, accId)
}

func testOrgResourceUpstreamHelmRegistryLegacyPath(id string, accId string) string {
	return fmt.Sprintf(`
 resource "harness_platform_organization" "test" {
  identifier = "%[1]s_org"
  name = "%[1]s"
 }

 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s/${harness_platform_organization.test.identifier}"
   package_type = "HELM"

   config {
		type = "UPSTREAM"
		auth_type = "UserPassword"
		source = "Custom"
		url = "https://har-registry.default.svc.cluster.local"
		auth {
			auth_type = "UserPassword"
			user_name = "username"
			secret_identifier = "Secret_Token"
			secret_space_path = "%[2]s"
		}
   }
   parent_ref = "%[2]s/${harness_platform_organization.test.identifier}"
 }
`, id, accId)
}

func testProjResourceUpstreamHelmRegistryLegacyPath(id string, accId string) string {
	return fmt.Sprintf(`
 resource "harness_platform_organization" "test" {
  identifier = "%[1]s_org"
  name = "%[1]s"
 }

 resource "harness_platform_project" "test" {
  identifier = "%[1]s_project"
  name = "%[1]s"
  org_id = harness_platform_organization.test.id
  color = "#472848"
 }
 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s/${harness_platform_organization.test.identifier}/${harness_platform_project.test.identifier}"
   package_type = "HELM"

   config {
		type = "UPSTREAM"
		auth_type = "UserPassword"
		source = "Custom"
		url = "https://har-registry.default.svc.cluster.local"
		auth {
			auth_type = "UserPassword"
			user_name = "username"
			secret_identifier = "Secret_Token"
			secret_space_path = "%[2]s"
		}
   }
   parent_ref = "%[2]s/${harness_platform_organization.test.identifier}/${harness_platform_project.test.identifier}"
 }
`, id, accId)
}

// Upstream Docker Registry --- Anonymous
func TestAccResourceUpstreamDockerAnonymousRegistryLegacyPath(t *testing.T) {
	id := fmt.Sprintf("tf_auto_upstream_docker_registry")
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config:             testAccResourceUpstreamDockerAnonymousRegistry(id, accountId),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
		},
	})
}
func TestOrgResourceUpstreamDockerAnonymousRegistryLegacyPath(t *testing.T) {
	id := fmt.Sprintf("tf_auto_upstream_docker_registry")
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config:             testOrgResourceUpstreamDockerAnonymousRegistry(id, accountId),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
		},
	})
}
func TestProjectResourceUpstreamDockerAnonymousRegistryLegacyPath(t *testing.T) {
	id := fmt.Sprintf("tf_auto_upstream_docker_registry")
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config:             testProjResourceUpstreamDockerAnonymousRegistry(id, accountId),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
		},
	})
}

func testAccResourceUpstreamDockerAnonymousRegistryLegacyPath(id string, accId string) string {
	return fmt.Sprintf(`
 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s"
   package_type = "DOCKER"

   config {
		type = "UPSTREAM"
		auth_type = "Anonymous"
		source = "Dockerhub"
		auth {
			auth_type = "Anonymous"
		}
   }
   parent_ref = "%[2]s"
 }
`, id, accId)
}

func testOrgResourceUpstreamDockerAnonymousRegistryLegacyPath(id string, accId string) string {
	return fmt.Sprintf(`
 resource "harness_platform_organization" "test" {
  identifier = "%[1]s_org"
  name = "%[1]s"
 }

 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s/${harness_platform_organization.test.identifier}"
   package_type = "DOCKER"

   config {
		type = "UPSTREAM"
		auth_type = "UserPassword"
		source = "Dockerhub"
		auth {
			auth_type = "Anonymous"
		}
   }
   parent_ref = "%[2]s/${harness_platform_organization.test.identifier}"
 }
`, id, accId)
}

func testProjResourceUpstreamDockerAnonymousRegistryLegacyPath(id string, accId string) string {
	return fmt.Sprintf(`
 resource "harness_platform_organization" "test" {
  identifier = "%[1]s_org"
  name = "%[1]s"
 }

 resource "harness_platform_project" "test" {
  identifier = "%[1]s_project"
  name = "%[1]s"
  org_id = harness_platform_organization.test.id
  color = "#472848"
 }
 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s/${harness_platform_organization.test.identifier}/${harness_platform_project.test.identifier}"
   package_type = "DOCKER"

   config {
		type = "UPSTREAM"
		auth_type = "UserPassword"
		source = "Dockerhub"
		auth {
			auth_type = "Anonymous"
		}
   }
   parent_ref = "%[2]s/${harness_platform_organization.test.identifier}/${harness_platform_project.test.identifier}"
 }
`, id, accId)
}

// Upstream Helm Registry --- Anonymous
func TestAccResourceUpstreamHelmAnonymousRegistryLegacyPath(t *testing.T) {
	id := fmt.Sprintf("tf_auto_upstream_helm_registry")
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config:             testAccResourceUpstreamHelmAnonymousRegistry(id, accountId),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
		},
	})
}
func TestOrgResourceUpstreamHelmAnonymousRegistryLegacyPath(t *testing.T) {
	id := fmt.Sprintf("tf_auto_upstream_helm_registry")
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config:             testOrgResourceUpstreamHelmAnonymousRegistry(id, accountId),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
		},
	})
}
func TestProjectResourceUpstreamHelmAnonymousRegistryLegacyPath(t *testing.T) {
	id := fmt.Sprintf("tf_auto_upstream_helm_registry")
	resourceName := "harness_platform_har_registry.test"
	accountId := os.Getenv("HARNESS_ACCOUNT_ID")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { acctest.TestAccPreCheck(t) },
		ProviderFactories: acctest.ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					acctest.TestAccConfigureProvider()
					_, _ = acctest.TestAccGetHarClientWithContext()
				},
				Config:             testProjResourceUpstreamHelmAnonymousRegistry(id, accountId),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "identifier", id),
				),
			},
		},
	})
}

func testAccResourceUpstreamHelmAnonymousRegistryLegacyPath(id string, accId string) string {
	return fmt.Sprintf(`
 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s"
   package_type = "HELM"

   config {
		type = "UPSTREAM"
		auth_type = "UserPassword"
		source = "Custom"
		url = "https://har-registry.default.svc.cluster.local"
		auth {
			auth_type = "Anonymous"
		}
   }
   parent_ref = "%[2]s"
 }
`, id, accId)
}

func testOrgResourceUpstreamHelmAnonymousRegistryLegacyPath(id string, accId string) string {
	return fmt.Sprintf(`
 resource "harness_platform_organization" "test" {
  identifier = "%[1]s_org"
  name = "%[1]s"
 }

 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s/${harness_platform_organization.test.identifier}"
   package_type = "HELM"

   config {
		type = "UPSTREAM"
		auth_type = "UserPassword"
		source = "Custom"
		url = "https://har-registry.default.svc.cluster.local"
		auth {
			auth_type = "Anonymous"
		}
   }
   parent_ref = "%[2]s/${harness_platform_organization.test.identifier}"
 }
`, id, accId)
}

func testProjResourceUpstreamHelmAnonymousRegistryLegacyPath(id string, accId string) string {
	return fmt.Sprintf(`
 resource "harness_platform_organization" "test" {
  identifier = "%[1]s_org"
  name = "%[1]s"
 }

 resource "harness_platform_project" "test" {
  identifier = "%[1]s_project"
  name = "%[1]s"
  org_id = harness_platform_organization.test.id
  color = "#472848"
 }
 resource "harness_platform_har_registry" "test" {
   identifier   = "%[1]s"
   space_ref    = "%[2]s/${harness_platform_organization.test.identifier}/${harness_platform_project.test.identifier}"
   package_type = "HELM"

   config {
		type = "UPSTREAM"
		auth_type = "UserPassword"
		source = "Custom"
		url = "https://har-registry.default.svc.cluster.local"
		auth {
			auth_type = "Anonymous"
		}
   }
   parent_ref = "%[2]s/${harness_platform_organization.test.identifier}/${harness_platform_project.test.identifier}"
 }
`, id, accId)
}
