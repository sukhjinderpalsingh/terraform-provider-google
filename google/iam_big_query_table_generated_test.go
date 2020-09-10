// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccBigQueryTableIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/bigquery.dataOwner",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccBigQueryTableIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_bigquery_table_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/datasets/%s/tables/%s roles/bigquery.dataOwner", getTestProjectFromEnv(), fmt.Sprintf("tf_test_dataset_id%s", context["random_suffix"]), fmt.Sprintf("tf_test_table_id%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccBigQueryTableIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_bigquery_table_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/datasets/%s/tables/%s roles/bigquery.dataOwner", getTestProjectFromEnv(), fmt.Sprintf("tf_test_dataset_id%s", context["random_suffix"]), fmt.Sprintf("tf_test_table_id%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccBigQueryTableIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/bigquery.dataOwner",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccBigQueryTableIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_bigquery_table_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/datasets/%s/tables/%s roles/bigquery.dataOwner user:admin@hashicorptest.com", getTestProjectFromEnv(), fmt.Sprintf("tf_test_dataset_id%s", context["random_suffix"]), fmt.Sprintf("tf_test_table_id%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccBigQueryTableIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/bigquery.dataOwner",
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccBigQueryTableIamPolicy_basicGenerated(context),
			},
			{
				ResourceName:      "google_bigquery_table_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/datasets/%s/tables/%s", getTestProjectFromEnv(), fmt.Sprintf("tf_test_dataset_id%s", context["random_suffix"]), fmt.Sprintf("tf_test_table_id%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccBigQueryTableIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_bigquery_table_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/datasets/%s/tables/%s", getTestProjectFromEnv(), fmt.Sprintf("tf_test_dataset_id%s", context["random_suffix"]), fmt.Sprintf("tf_test_table_id%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccBigQueryTableIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_bigquery_dataset" "test" {
	dataset_id = "tf_test_dataset_id%{random_suffix}"
}

resource "google_bigquery_table" "test" {
	table_id   = "tf_test_table_id%{random_suffix}"
	dataset_id = google_bigquery_dataset.test.dataset_id
	time_partitioning {
		type = "DAY"
	}
	schema = <<EOH
[
	{
		"name": "city",
		"type": "RECORD",
		"fields": [
	{
		"name": "id",
		"type": "INTEGER"
	},
	{
		"name": "coord",
		"type": "RECORD",
		"fields": [
		{
			"name": "lon",
			"type": "FLOAT"
		},
		{
			"name": "lat",
			"type": "FLOAT"
		}
		]
	}
		]
	},
	{
		"name": "country",
		"type": "RECORD",
		"fields": [
	{
		"name": "id",
		"type": "INTEGER"
	},
	{
		"name": "name",
		"type": "STRING"
	}
		]
	}
]
EOH
}

resource "google_bigquery_table_iam_member" "foo" {
  project = google_bigquery_table.test.project
  dataset_id = google_bigquery_table.test.dataset_id
  table_id = google_bigquery_table.test.table_id
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccBigQueryTableIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_bigquery_dataset" "test" {
	dataset_id = "tf_test_dataset_id%{random_suffix}"
}

resource "google_bigquery_table" "test" {
	table_id   = "tf_test_table_id%{random_suffix}"
	dataset_id = google_bigquery_dataset.test.dataset_id
	time_partitioning {
		type = "DAY"
	}
	schema = <<EOH
[
	{
		"name": "city",
		"type": "RECORD",
		"fields": [
	{
		"name": "id",
		"type": "INTEGER"
	},
	{
		"name": "coord",
		"type": "RECORD",
		"fields": [
		{
			"name": "lon",
			"type": "FLOAT"
		},
		{
			"name": "lat",
			"type": "FLOAT"
		}
		]
	}
		]
	},
	{
		"name": "country",
		"type": "RECORD",
		"fields": [
	{
		"name": "id",
		"type": "INTEGER"
	},
	{
		"name": "name",
		"type": "STRING"
	}
		]
	}
]
EOH
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_bigquery_table_iam_policy" "foo" {
  project = google_bigquery_table.test.project
  dataset_id = google_bigquery_table.test.dataset_id
  table_id = google_bigquery_table.test.table_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccBigQueryTableIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_bigquery_dataset" "test" {
	dataset_id = "tf_test_dataset_id%{random_suffix}"
}

resource "google_bigquery_table" "test" {
	table_id   = "tf_test_table_id%{random_suffix}"
	dataset_id = google_bigquery_dataset.test.dataset_id
	time_partitioning {
		type = "DAY"
	}
	schema = <<EOH
[
	{
		"name": "city",
		"type": "RECORD",
		"fields": [
	{
		"name": "id",
		"type": "INTEGER"
	},
	{
		"name": "coord",
		"type": "RECORD",
		"fields": [
		{
			"name": "lon",
			"type": "FLOAT"
		},
		{
			"name": "lat",
			"type": "FLOAT"
		}
		]
	}
		]
	},
	{
		"name": "country",
		"type": "RECORD",
		"fields": [
	{
		"name": "id",
		"type": "INTEGER"
	},
	{
		"name": "name",
		"type": "STRING"
	}
		]
	}
]
EOH
}

data "google_iam_policy" "foo" {
}

resource "google_bigquery_table_iam_policy" "foo" {
  project = google_bigquery_table.test.project
  dataset_id = google_bigquery_table.test.dataset_id
  table_id = google_bigquery_table.test.table_id
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccBigQueryTableIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_bigquery_dataset" "test" {
	dataset_id = "tf_test_dataset_id%{random_suffix}"
}

resource "google_bigquery_table" "test" {
	table_id   = "tf_test_table_id%{random_suffix}"
	dataset_id = google_bigquery_dataset.test.dataset_id
	time_partitioning {
		type = "DAY"
	}
	schema = <<EOH
[
	{
		"name": "city",
		"type": "RECORD",
		"fields": [
	{
		"name": "id",
		"type": "INTEGER"
	},
	{
		"name": "coord",
		"type": "RECORD",
		"fields": [
		{
			"name": "lon",
			"type": "FLOAT"
		},
		{
			"name": "lat",
			"type": "FLOAT"
		}
		]
	}
		]
	},
	{
		"name": "country",
		"type": "RECORD",
		"fields": [
	{
		"name": "id",
		"type": "INTEGER"
	},
	{
		"name": "name",
		"type": "STRING"
	}
		]
	}
]
EOH
}

resource "google_bigquery_table_iam_binding" "foo" {
  project = google_bigquery_table.test.project
  dataset_id = google_bigquery_table.test.dataset_id
  table_id = google_bigquery_table.test.table_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccBigQueryTableIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_bigquery_dataset" "test" {
	dataset_id = "tf_test_dataset_id%{random_suffix}"
}

resource "google_bigquery_table" "test" {
	table_id   = "tf_test_table_id%{random_suffix}"
	dataset_id = google_bigquery_dataset.test.dataset_id
	time_partitioning {
		type = "DAY"
	}
	schema = <<EOH
[
	{
		"name": "city",
		"type": "RECORD",
		"fields": [
	{
		"name": "id",
		"type": "INTEGER"
	},
	{
		"name": "coord",
		"type": "RECORD",
		"fields": [
		{
			"name": "lon",
			"type": "FLOAT"
		},
		{
			"name": "lat",
			"type": "FLOAT"
		}
		]
	}
		]
	},
	{
		"name": "country",
		"type": "RECORD",
		"fields": [
	{
		"name": "id",
		"type": "INTEGER"
	},
	{
		"name": "name",
		"type": "STRING"
	}
		]
	}
]
EOH
}

resource "google_bigquery_table_iam_binding" "foo" {
  project = google_bigquery_table.test.project
  dataset_id = google_bigquery_table.test.dataset_id
  table_id = google_bigquery_table.test.table_id
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:paddy@hashicorp.com"]
}
`, context)
}