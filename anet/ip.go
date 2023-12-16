package anet

import (
    "net"
)

//获取本机ip4地址
func GetIP4List() (ipList []string) {
    interfaceAddrs, err := net.InterfaceAddrs()
    if err != nil {
        return ipList
    }

    for _, item := range interfaceAddrs {
        ipNet, isValidIpNet := item.(*net.IPNet)
        if isValidIpNet && !ipNet.IP.IsLoopback() {
            if ipNet.IP.To4() != nil {
                ipList = append(ipList, ipNet.IP.String())
            }
        }
    }
    return ipList
}

//是否本机ip地址
func IsIP4(ip string) bool {
    ipList := GetIP4List()
    for _, item := range ipList {
        if item == ip {
            return true
        }
    }
    return false
}
