# main.tf

provider "virtualbox" {
}

resource "virtualbox_vm" "example" {
  name   = "example-vm"
  memory = 1024
  cpus   = 1

  network_adapter {
    bridge_mode = "natnetwork"
  }

  storage_controller {
    name           = "SATA Controller"
    bus            = "sata"
    disk_interface = "disk"
  }

  storage {
    controller = "SATA Controller"
    disk       = "example-disk.vdi"
  }

  guest_additions_mode = "upload"
  guest_additions_path = "./VBoxGuestAdditions.iso"

  provisioner "file" {
    source      = "script.sh"
    destination = "/tmp/script.sh"
  }

  provisioner "remote-exec" {
    inline = ["chmod +x /tmp/script.sh", "/tmp/script.sh"]
  }
}
