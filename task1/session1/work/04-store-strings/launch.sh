#!/bin/bash

if test $# -ne 1; then
    echo "Usage: $0 <path_to_kvm_image>" 1>&2
    exit 1
fi

KVM_IMAGE="$1"
BRIDGE_IFACE="virbr0"
BRIDGE_IP="172.44.0.1"
VM_IP="172.44.0.2"

echo "Creating bridge $BRIDGE_IFACE with IP address $BRIDGE_IP ..."
sudo brctl addbr "$BRIDGE_IFACE" || true
sudo ifconfig "$BRIDGE_IFACE" "$BRIDGE_IP"

echo "Starting KVM image connected to bridge interface $BRIDGE_IFACE ..."
sudo qemu-system-x86_64  -fsdev local,id=myid,path=$(pwd)/fs0,security_model=none \
                         -device virtio-9p-pci,fsdev=myid,mount_tag=rootfs,disable-modern=on,disable-legacy=off \
                         -netdev bridge,id=en0,br=virbr0 \
                         -device virtio-net-pci,netdev=en0 \
                         -kernel "$KVM_IMAGE" \
                         -append "netdev.ipv4_addr=$VM_IP netdev.ipv4_gw_addr=$BRIDGE_IP netdev.ipv4_subnet_mask=255.255.255.0 --" \
                         -cpu host \
                         -enable-kvm \
                         -nographic

sudo ip l set dev virbr0 down
sudo brctl delbr virbr0
sudo brctl addbr virbr0
sudo ip a a  172.44.0.1/24 dev virbr0
sudo ip l set dev virbr0 up
sudo qemu-system-x86_64  -fsdev local,id=myid,path=$(pwd)/fs0,security_model=none \
                         -device virtio-9p-pci,fsdev=myid,mount_tag=rootfs,disable-modern=on,disable-legacy=off \
                         -netdev bridge,id=en0,br=virbr0 \
                         -device virtio-net-pci,netdev=en0 \
                         -kernel build/04-store-strings_qemu-x86_64 \
                         -append "netdev.ipv4_addr=172.44.0.2 netdev.ipv4_gw_addr=172.44.0.1 netdev.ipv4_subnet_mask=255.255.255.0 --" \
                         -nographic