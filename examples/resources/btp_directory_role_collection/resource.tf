resource "btp_directory_role_collection" "my_collection" {
  directory_id = "dd005d8b-1fee-4e6b-b6ff-cb9a197b7fe0"
  name         = "My own role collection"
  description  = "A description of what the role collection is supposed to do."

  role_references = [
    {
      name                 = "Directory Admin"
      role_template_app_id = "cis-central!b13"
      role_template_name   = "Directory_Admin"
    }
  ]
}
