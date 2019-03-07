// Automatically generated by MockGen. DO NOT EDIT!
// Source: common.go

package network

import (
	net "net"

	iptables "github.com/coreos/go-iptables/iptables"
	gomock "github.com/golang/mock/gomock"
	netlink "github.com/vishvananda/netlink"

	v1 "kubevirt.io/client-go/api/v1"
)

// Mock of NetworkHandler interface
type MockNetworkHandler struct {
	ctrl     *gomock.Controller
	recorder *_MockNetworkHandlerRecorder
}

// Recorder for MockNetworkHandler (not exported)
type _MockNetworkHandlerRecorder struct {
	mock *MockNetworkHandler
}

func NewMockNetworkHandler(ctrl *gomock.Controller) *MockNetworkHandler {
	mock := &MockNetworkHandler{ctrl: ctrl}
	mock.recorder = &_MockNetworkHandlerRecorder{mock}
	return mock
}

func (_m *MockNetworkHandler) EXPECT() *_MockNetworkHandlerRecorder {
	return _m.recorder
}

func (_m *MockNetworkHandler) LinkByName(name string) (netlink.Link, error) {
	ret := _m.ctrl.Call(_m, "LinkByName", name)
	ret0, _ := ret[0].(netlink.Link)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetworkHandlerRecorder) LinkByName(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LinkByName", arg0)
}

func (_m *MockNetworkHandler) AddrList(link netlink.Link, family int) ([]netlink.Addr, error) {
	ret := _m.ctrl.Call(_m, "AddrList", link, family)
	ret0, _ := ret[0].([]netlink.Addr)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetworkHandlerRecorder) AddrList(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddrList", arg0, arg1)
}

func (_m *MockNetworkHandler) RouteList(link netlink.Link, family int) ([]netlink.Route, error) {
	ret := _m.ctrl.Call(_m, "RouteList", link, family)
	ret0, _ := ret[0].([]netlink.Route)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetworkHandlerRecorder) RouteList(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RouteList", arg0, arg1)
}

func (_m *MockNetworkHandler) AddrDel(link netlink.Link, addr *netlink.Addr) error {
	ret := _m.ctrl.Call(_m, "AddrDel", link, addr)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockNetworkHandlerRecorder) AddrDel(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddrDel", arg0, arg1)
}

func (_m *MockNetworkHandler) AddrAdd(link netlink.Link, addr *netlink.Addr) error {
	ret := _m.ctrl.Call(_m, "AddrAdd", link, addr)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockNetworkHandlerRecorder) AddrAdd(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddrAdd", arg0, arg1)
}

func (_m *MockNetworkHandler) AddrReplace(link netlink.Link, addr *netlink.Addr) error {
	ret := _m.ctrl.Call(_m, "AddrReplace", link, addr)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockNetworkHandlerRecorder) AddrReplace(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddrReplace", arg0, arg1)
}

func (_m *MockNetworkHandler) LinkSetDown(link netlink.Link) error {
	ret := _m.ctrl.Call(_m, "LinkSetDown", link)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockNetworkHandlerRecorder) LinkSetDown(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LinkSetDown", arg0)
}

func (_m *MockNetworkHandler) LinkSetUp(link netlink.Link) error {
	ret := _m.ctrl.Call(_m, "LinkSetUp", link)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockNetworkHandlerRecorder) LinkSetUp(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LinkSetUp", arg0)
}

func (_m *MockNetworkHandler) LinkSetName(link netlink.Link, name string) error {
	ret := _m.ctrl.Call(_m, "LinkSetName", link, name)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockNetworkHandlerRecorder) LinkSetName(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LinkSetName", arg0, arg1)
}

func (_m *MockNetworkHandler) LinkAdd(link netlink.Link) error {
	ret := _m.ctrl.Call(_m, "LinkAdd", link)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockNetworkHandlerRecorder) LinkAdd(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LinkAdd", arg0)
}

func (_m *MockNetworkHandler) LinkSetLearningOff(link netlink.Link) error {
	ret := _m.ctrl.Call(_m, "LinkSetLearningOff", link)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockNetworkHandlerRecorder) LinkSetLearningOff(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LinkSetLearningOff", arg0)
}

func (_m *MockNetworkHandler) ParseAddr(s string) (*netlink.Addr, error) {
	ret := _m.ctrl.Call(_m, "ParseAddr", s)
	ret0, _ := ret[0].(*netlink.Addr)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetworkHandlerRecorder) ParseAddr(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ParseAddr", arg0)
}

func (_m *MockNetworkHandler) GetHostAndGwAddressesFromCIDR(s string) (string, string, error) {
	ret := _m.ctrl.Call(_m, "GetHostAndGwAddressesFromCIDR", s)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockNetworkHandlerRecorder) GetHostAndGwAddressesFromCIDR(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetHostAndGwAddressesFromCIDR", arg0)
}

func (_m *MockNetworkHandler) SetRandomMac(iface string) (net.HardwareAddr, error) {
	ret := _m.ctrl.Call(_m, "SetRandomMac", iface)
	ret0, _ := ret[0].(net.HardwareAddr)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetworkHandlerRecorder) SetRandomMac(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetRandomMac", arg0)
}

func (_m *MockNetworkHandler) GenerateRandomMac() (net.HardwareAddr, error) {
	ret := _m.ctrl.Call(_m, "GenerateRandomMac")
	ret0, _ := ret[0].(net.HardwareAddr)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetworkHandlerRecorder) GenerateRandomMac() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GenerateRandomMac")
}

func (_m *MockNetworkHandler) GetMacDetails(iface string) (net.HardwareAddr, error) {
	ret := _m.ctrl.Call(_m, "GetMacDetails", iface)
	ret0, _ := ret[0].(net.HardwareAddr)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetworkHandlerRecorder) GetMacDetails(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetMacDetails", arg0)
}

func (_m *MockNetworkHandler) LinkSetMaster(link netlink.Link, master *netlink.Bridge) error {
	ret := _m.ctrl.Call(_m, "LinkSetMaster", link, master)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockNetworkHandlerRecorder) LinkSetMaster(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LinkSetMaster", arg0, arg1)
}

func (_m *MockNetworkHandler) StartDHCP(nic *VIF, serverAddr net.IP, bridgeInterfaceName string, dhcpOptions *v1.DHCPOptions, filterByMAC bool) error {
	ret := _m.ctrl.Call(_m, "StartDHCP", nic, serverAddr, bridgeInterfaceName, dhcpOptions, filterByMAC)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockNetworkHandlerRecorder) StartDHCP(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "StartDHCP", arg0, arg1, arg2, arg3, arg4)
}

func (_m *MockNetworkHandler) HasNatIptables(proto iptables.Protocol) bool {
	ret := _m.ctrl.Call(_m, "HasNatIptables", proto)
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockNetworkHandlerRecorder) HasNatIptables(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HasNatIptables", arg0)
}

func (_m *MockNetworkHandler) IsIpv6Enabled(interfaceName string) (bool, error) {
	ret := _m.ctrl.Call(_m, "IsIpv6Enabled", interfaceName)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetworkHandlerRecorder) IsIpv6Enabled(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsIpv6Enabled", arg0)
}

func (_m *MockNetworkHandler) IsIpv4Primary() (bool, error) {
	ret := _m.ctrl.Call(_m, "IsIpv4Primary")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockNetworkHandlerRecorder) IsIpv4Primary() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsIpv4Primary")
}

func (_m *MockNetworkHandler) ConfigureIpv6Forwarding() error {
	ret := _m.ctrl.Call(_m, "ConfigureIpv6Forwarding")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockNetworkHandlerRecorder) ConfigureIpv6Forwarding() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ConfigureIpv6Forwarding")
}

func (_m *MockNetworkHandler) ConfigureIpv4ArpIgnore() error {
	ret := _m.ctrl.Call(_m, "ConfigureIpv4ArpIgnore")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockNetworkHandlerRecorder) ConfigureIpv4ArpIgnore() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ConfigureIpv4ArpIgnore")
}

func (_m *MockNetworkHandler) IptablesNewChain(proto iptables.Protocol, table string, chain string) error {
	ret := _m.ctrl.Call(_m, "IptablesNewChain", proto, table, chain)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockNetworkHandlerRecorder) IptablesNewChain(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IptablesNewChain", arg0, arg1, arg2)
}

func (_m *MockNetworkHandler) IptablesAppendRule(proto iptables.Protocol, table string, chain string, rulespec ...string) error {
	_s := []interface{}{proto, table, chain}
	for _, _x := range rulespec {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "IptablesAppendRule", _s...)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockNetworkHandlerRecorder) IptablesAppendRule(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IptablesAppendRule", _s...)
}

func (_m *MockNetworkHandler) NftablesNewChain(proto iptables.Protocol, table string, chain string) error {
	ret := _m.ctrl.Call(_m, "NftablesNewChain", proto, table, chain)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockNetworkHandlerRecorder) NftablesNewChain(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NftablesNewChain", arg0, arg1, arg2)
}

func (_m *MockNetworkHandler) NftablesAppendRule(proto iptables.Protocol, table string, chain string, rulespec ...string) error {
	_s := []interface{}{proto, table, chain}
	for _, _x := range rulespec {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "NftablesAppendRule", _s...)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockNetworkHandlerRecorder) NftablesAppendRule(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NftablesAppendRule", _s...)
}

func (_m *MockNetworkHandler) NftablesLoad(fnName string) error {
	ret := _m.ctrl.Call(_m, "NftablesLoad", fnName)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockNetworkHandlerRecorder) NftablesLoad(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NftablesLoad", arg0)
}

func (_m *MockNetworkHandler) GetNFTIPString(proto iptables.Protocol) string {
	ret := _m.ctrl.Call(_m, "GetNFTIPString", proto)
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockNetworkHandlerRecorder) GetNFTIPString(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetNFTIPString", arg0)
}

func (_m *MockNetworkHandler) CreateTapDevice(tapName string, queueNumber uint32, launcherPID int, mtu int) error {
	ret := _m.ctrl.Call(_m, "CreateTapDevice", tapName, queueNumber, launcherPID, mtu)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockNetworkHandlerRecorder) CreateTapDevice(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateTapDevice", arg0, arg1, arg2, arg3)
}

func (_m *MockNetworkHandler) BindTapDeviceToBridge(tapName string, bridgeName string) error {
	ret := _m.ctrl.Call(_m, "BindTapDeviceToBridge", tapName, bridgeName)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockNetworkHandlerRecorder) BindTapDeviceToBridge(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "BindTapDeviceToBridge", arg0, arg1)
}

func (_m *MockNetworkHandler) DisableTXOffloadChecksum(ifaceName string) error {
	ret := _m.ctrl.Call(_m, "DisableTXOffloadChecksum", ifaceName)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockNetworkHandlerRecorder) DisableTXOffloadChecksum(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DisableTXOffloadChecksum", arg0)
}
