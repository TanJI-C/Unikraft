#!/bin/bash

if test $# -ne 1; then
    echo "Usage: $0 <kvm_image>" 1>&2
    exit 1
fi

kvm_image="$1"
fs_tag="fs0"
local_fs_dir="./guest_fs/"

echo "Starting QEMU image $kvm_image mounting "$local_fs_dir" ..."
sudo qemu-system-x86_64  -fsdev local,id=myid,path=$(pwd)/guest_fs,security_model=none \
    -device virtio-9p-pci,fsdev=myid,mount_tag=fs0,disable-modern=on,disable-legacy=off \
    -kernel "build/06-adding-filesystems_qemu-x86_64" \
    -nographic
    -cpu host \
    -enable-kvm \
./qemu-guest -e guest_fs/ -k build/06-adding-filesystems_qemu-x86_64