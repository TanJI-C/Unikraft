package main

import (
	"fmt"

	"libvirt.org/go/libvirt"
)

func main() {
	fmt.Println("main")
	vName := "shengborun"
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	defer conn.Close()

	// 通过虚拟机名称获取其全部状态
	dob, err := conn.LookupDomainByName(vName)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	dobs := make([]*libvirt.Domain, 0)
	dobs = append(dobs, dob)
	dstats, err := conn.GetAllDomainStats(dobs, 0x3FE, libvirt.CONNECT_GET_ALL_DOMAINS_STATS_ACTIVE)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("dstats", dstats)

	// 获取所有开启的虚拟机
	doms, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Printf("%d running domains:\n", len(doms))
	for _, dom := range doms {
		name, err := dom.GetName()
		if err == nil {
			fmt.Printf("  %s\n", name)
		}
		if name != vName {
			continue
		}
		info, err := dom.GetInfo()
		fmt.Println("GetInfo", info, err)

		//内存利用率
		// get tag 4:剩余 & 5:总共
		meminfo, err := dom.MemoryStats(10, 0)
		fmt.Println("MemoryStats", meminfo, err)

		//磁盘利用率-错误值（都是最大值）
		diskinfo, err := dom.GetBlockInfo("hda", 0)
		fmt.Println("DiskStats", diskinfo, err)

		blockinfo, err := dom.BlockStats("hda")
		fmt.Println("BlockStats", blockinfo, err)

		//go版本不支持
		guestinfo, err := dom.GetGuestInfo(libvirt.DOMAIN_GUEST_INFO_DISKS, 0)
		fmt.Println("GetGuestInfo", guestinfo, err)

		//dom.GetCPUStats()

		dom.Free()
	}
}

