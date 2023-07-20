sudo brctl addbr kraft0
sudo ifconfig kraft0 172.44.0.1
sudo ifconfig kraft0 up

sudo dnsmasq -d \
             -log-queries \
             --bind-dynamic \
             --interface=kraft0 \
             --listen-addr=172.44.0.1 \
             --dhcp-range=172.44.0.2,172.44.0.254,255.255.255.0,12h &> dnsmasq.logs &

./qemu-guest.sh -k ./build/app-nginx_qemu-x86_64 \
                -a "" \
                -b kraft0 \
                -e ./nginx_files \
                -m 100

sudo qemu-system-x86_64 -kernel build/app-nginx_qemu-x86_64 \
                   -fsdev local,id=myid,path=$(pwd)/nginx_files,security_model=none \
                   -device virtio-9p-pci,fsdev=myid,mount_tag=fs0,disable-modern=on,disable-legacy=off \
                   -netdev bridge,id=en0,br=kraft0 \
                   -device virtio-net-pci,netdev=en0 \
                   -append "" \
                   -m 100 \
                   -nographic