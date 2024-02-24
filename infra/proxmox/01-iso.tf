resource "proxmox_virtual_environment_download_file" "talos" {
  content_type       = "iso"
  datastore_id       = "local"
  node_name          = "proxmox"
  url                = "https://github.com/siderolabs/talos/releases/download/v1.6.4/metal-amd64.iso"
  file_name          = "tf-talos-1.6.4.iso"
  checksum           = "b3786227765ab09301e10f30a0808dd8728e03c9ea51e694dbd91370bb3c2b21"
  checksum_algorithm = "sha256"
}
