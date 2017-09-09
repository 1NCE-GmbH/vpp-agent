// Code generated by github.com/ungerik/pkgreflect DO NOT EDIT.

package vpe

import "reflect"

var Types = map[string]reflect.Type{
	"AddNodeNext": reflect.TypeOf((*AddNodeNext)(nil)).Elem(),
	"AddNodeNextReply": reflect.TypeOf((*AddNodeNextReply)(nil)).Elem(),
	"BdIPMacAddDel": reflect.TypeOf((*BdIPMacAddDel)(nil)).Elem(),
	"BdIPMacAddDelReply": reflect.TypeOf((*BdIPMacAddDelReply)(nil)).Elem(),
	"ClassifySetInterfaceIPTable": reflect.TypeOf((*ClassifySetInterfaceIPTable)(nil)).Elem(),
	"ClassifySetInterfaceIPTableReply": reflect.TypeOf((*ClassifySetInterfaceIPTableReply)(nil)).Elem(),
	"ClassifySetInterfaceL2Tables": reflect.TypeOf((*ClassifySetInterfaceL2Tables)(nil)).Elem(),
	"ClassifySetInterfaceL2TablesReply": reflect.TypeOf((*ClassifySetInterfaceL2TablesReply)(nil)).Elem(),
	"Cli": reflect.TypeOf((*Cli)(nil)).Elem(),
	"CliInband": reflect.TypeOf((*CliInband)(nil)).Elem(),
	"CliInbandReply": reflect.TypeOf((*CliInbandReply)(nil)).Elem(),
	"CliReply": reflect.TypeOf((*CliReply)(nil)).Elem(),
	"ControlPing": reflect.TypeOf((*ControlPing)(nil)).Elem(),
	"ControlPingReply": reflect.TypeOf((*ControlPingReply)(nil)).Elem(),
	"CreateLoopback": reflect.TypeOf((*CreateLoopback)(nil)).Elem(),
	"CreateLoopbackInstance": reflect.TypeOf((*CreateLoopbackInstance)(nil)).Elem(),
	"CreateLoopbackInstanceReply": reflect.TypeOf((*CreateLoopbackInstanceReply)(nil)).Elem(),
	"CreateLoopbackReply": reflect.TypeOf((*CreateLoopbackReply)(nil)).Elem(),
	"CreateSubif": reflect.TypeOf((*CreateSubif)(nil)).Elem(),
	"CreateSubifReply": reflect.TypeOf((*CreateSubifReply)(nil)).Elem(),
	"CreateVlanSubif": reflect.TypeOf((*CreateVlanSubif)(nil)).Elem(),
	"CreateVlanSubifReply": reflect.TypeOf((*CreateVlanSubifReply)(nil)).Elem(),
	"DeleteLoopback": reflect.TypeOf((*DeleteLoopback)(nil)).Elem(),
	"DeleteLoopbackReply": reflect.TypeOf((*DeleteLoopbackReply)(nil)).Elem(),
	"DeleteSubif": reflect.TypeOf((*DeleteSubif)(nil)).Elem(),
	"DeleteSubifReply": reflect.TypeOf((*DeleteSubifReply)(nil)).Elem(),
	"FeatureEnableDisable": reflect.TypeOf((*FeatureEnableDisable)(nil)).Elem(),
	"FeatureEnableDisableReply": reflect.TypeOf((*FeatureEnableDisableReply)(nil)).Elem(),
	"GetNextIndex": reflect.TypeOf((*GetNextIndex)(nil)).Elem(),
	"GetNextIndexReply": reflect.TypeOf((*GetNextIndexReply)(nil)).Elem(),
	"GetNodeGraph": reflect.TypeOf((*GetNodeGraph)(nil)).Elem(),
	"GetNodeGraphReply": reflect.TypeOf((*GetNodeGraphReply)(nil)).Elem(),
	"GetNodeIndex": reflect.TypeOf((*GetNodeIndex)(nil)).Elem(),
	"GetNodeIndexReply": reflect.TypeOf((*GetNodeIndexReply)(nil)).Elem(),
	"IP4ArpEvent": reflect.TypeOf((*IP4ArpEvent)(nil)).Elem(),
	"IP4FibCounter": reflect.TypeOf((*IP4FibCounter)(nil)).Elem(),
	"IP4NbrCounter": reflect.TypeOf((*IP4NbrCounter)(nil)).Elem(),
	"IP6FibCounter": reflect.TypeOf((*IP6FibCounter)(nil)).Elem(),
	"IP6NbrCounter": reflect.TypeOf((*IP6NbrCounter)(nil)).Elem(),
	"IP6NdEvent": reflect.TypeOf((*IP6NdEvent)(nil)).Elem(),
	"IPSourceAndPortRangeCheckAddDel": reflect.TypeOf((*IPSourceAndPortRangeCheckAddDel)(nil)).Elem(),
	"IPSourceAndPortRangeCheckAddDelReply": reflect.TypeOf((*IPSourceAndPortRangeCheckAddDelReply)(nil)).Elem(),
	"IPSourceAndPortRangeCheckInterfaceAddDel": reflect.TypeOf((*IPSourceAndPortRangeCheckInterfaceAddDel)(nil)).Elem(),
	"IPSourceAndPortRangeCheckInterfaceAddDelReply": reflect.TypeOf((*IPSourceAndPortRangeCheckInterfaceAddDelReply)(nil)).Elem(),
	"InputACLSetInterface": reflect.TypeOf((*InputACLSetInterface)(nil)).Elem(),
	"InputACLSetInterfaceReply": reflect.TypeOf((*InputACLSetInterfaceReply)(nil)).Elem(),
	"InterfaceNameRenumber": reflect.TypeOf((*InterfaceNameRenumber)(nil)).Elem(),
	"InterfaceNameRenumberReply": reflect.TypeOf((*InterfaceNameRenumberReply)(nil)).Elem(),
	"IoamDisable": reflect.TypeOf((*IoamDisable)(nil)).Elem(),
	"IoamDisableReply": reflect.TypeOf((*IoamDisableReply)(nil)).Elem(),
	"IoamEnable": reflect.TypeOf((*IoamEnable)(nil)).Elem(),
	"IoamEnableReply": reflect.TypeOf((*IoamEnableReply)(nil)).Elem(),
	"L2InterfaceEfpFilter": reflect.TypeOf((*L2InterfaceEfpFilter)(nil)).Elem(),
	"L2InterfaceEfpFilterReply": reflect.TypeOf((*L2InterfaceEfpFilterReply)(nil)).Elem(),
	"L2PatchAddDel": reflect.TypeOf((*L2PatchAddDel)(nil)).Elem(),
	"L2PatchAddDelReply": reflect.TypeOf((*L2PatchAddDelReply)(nil)).Elem(),
	"OamAddDel": reflect.TypeOf((*OamAddDel)(nil)).Elem(),
	"OamAddDelReply": reflect.TypeOf((*OamAddDelReply)(nil)).Elem(),
	"OamEvent": reflect.TypeOf((*OamEvent)(nil)).Elem(),
	"PgCapture": reflect.TypeOf((*PgCapture)(nil)).Elem(),
	"PgCaptureReply": reflect.TypeOf((*PgCaptureReply)(nil)).Elem(),
	"PgCreateInterface": reflect.TypeOf((*PgCreateInterface)(nil)).Elem(),
	"PgCreateInterfaceReply": reflect.TypeOf((*PgCreateInterfaceReply)(nil)).Elem(),
	"PgEnableDisable": reflect.TypeOf((*PgEnableDisable)(nil)).Elem(),
	"PgEnableDisableReply": reflect.TypeOf((*PgEnableDisableReply)(nil)).Elem(),
	"ProxyArpAddDel": reflect.TypeOf((*ProxyArpAddDel)(nil)).Elem(),
	"ProxyArpAddDelReply": reflect.TypeOf((*ProxyArpAddDelReply)(nil)).Elem(),
	"ProxyArpIntfcEnableDisable": reflect.TypeOf((*ProxyArpIntfcEnableDisable)(nil)).Elem(),
	"ProxyArpIntfcEnableDisableReply": reflect.TypeOf((*ProxyArpIntfcEnableDisableReply)(nil)).Elem(),
	"Punt": reflect.TypeOf((*Punt)(nil)).Elem(),
	"PuntReply": reflect.TypeOf((*PuntReply)(nil)).Elem(),
	"ResetFib": reflect.TypeOf((*ResetFib)(nil)).Elem(),
	"ResetFibReply": reflect.TypeOf((*ResetFibReply)(nil)).Elem(),
	"ResetVrf": reflect.TypeOf((*ResetVrf)(nil)).Elem(),
	"ResetVrfReply": reflect.TypeOf((*ResetVrfReply)(nil)).Elem(),
	"SetArpNeighborLimit": reflect.TypeOf((*SetArpNeighborLimit)(nil)).Elem(),
	"SetArpNeighborLimitReply": reflect.TypeOf((*SetArpNeighborLimitReply)(nil)).Elem(),
	"ShowVersion": reflect.TypeOf((*ShowVersion)(nil)).Elem(),
	"ShowVersionReply": reflect.TypeOf((*ShowVersionReply)(nil)).Elem(),
	"SwInterfaceSetL2Bridge": reflect.TypeOf((*SwInterfaceSetL2Bridge)(nil)).Elem(),
	"SwInterfaceSetL2BridgeReply": reflect.TypeOf((*SwInterfaceSetL2BridgeReply)(nil)).Elem(),
	"SwInterfaceSetL2Xconnect": reflect.TypeOf((*SwInterfaceSetL2Xconnect)(nil)).Elem(),
	"SwInterfaceSetL2XconnectReply": reflect.TypeOf((*SwInterfaceSetL2XconnectReply)(nil)).Elem(),
	"SwInterfaceSetMplsEnable": reflect.TypeOf((*SwInterfaceSetMplsEnable)(nil)).Elem(),
	"SwInterfaceSetMplsEnableReply": reflect.TypeOf((*SwInterfaceSetMplsEnableReply)(nil)).Elem(),
	"SwInterfaceSetVpath": reflect.TypeOf((*SwInterfaceSetVpath)(nil)).Elem(),
	"SwInterfaceSetVpathReply": reflect.TypeOf((*SwInterfaceSetVpathReply)(nil)).Elem(),
	"VnetGetSummaryStats": reflect.TypeOf((*VnetGetSummaryStats)(nil)).Elem(),
	"VnetGetSummaryStatsReply": reflect.TypeOf((*VnetGetSummaryStatsReply)(nil)).Elem(),
	"VnetIP4FibCounters": reflect.TypeOf((*VnetIP4FibCounters)(nil)).Elem(),
	"VnetIP4NbrCounters": reflect.TypeOf((*VnetIP4NbrCounters)(nil)).Elem(),
	"VnetIP6FibCounters": reflect.TypeOf((*VnetIP6FibCounters)(nil)).Elem(),
	"VnetIP6NbrCounters": reflect.TypeOf((*VnetIP6NbrCounters)(nil)).Elem(),
	"WantIP4ArpEvents": reflect.TypeOf((*WantIP4ArpEvents)(nil)).Elem(),
	"WantIP4ArpEventsReply": reflect.TypeOf((*WantIP4ArpEventsReply)(nil)).Elem(),
	"WantIP6NdEvents": reflect.TypeOf((*WantIP6NdEvents)(nil)).Elem(),
	"WantIP6NdEventsReply": reflect.TypeOf((*WantIP6NdEventsReply)(nil)).Elem(),
	"WantOamEvents": reflect.TypeOf((*WantOamEvents)(nil)).Elem(),
	"WantOamEventsReply": reflect.TypeOf((*WantOamEventsReply)(nil)).Elem(),
	"WantStats": reflect.TypeOf((*WantStats)(nil)).Elem(),
	"WantStatsReply": reflect.TypeOf((*WantStatsReply)(nil)).Elem(),
}

var Functions = map[string]reflect.Value{
	"NewAddNodeNext": reflect.ValueOf(NewAddNodeNext),
	"NewAddNodeNextReply": reflect.ValueOf(NewAddNodeNextReply),
	"NewBdIPMacAddDel": reflect.ValueOf(NewBdIPMacAddDel),
	"NewBdIPMacAddDelReply": reflect.ValueOf(NewBdIPMacAddDelReply),
	"NewClassifySetInterfaceIPTable": reflect.ValueOf(NewClassifySetInterfaceIPTable),
	"NewClassifySetInterfaceIPTableReply": reflect.ValueOf(NewClassifySetInterfaceIPTableReply),
	"NewClassifySetInterfaceL2Tables": reflect.ValueOf(NewClassifySetInterfaceL2Tables),
	"NewClassifySetInterfaceL2TablesReply": reflect.ValueOf(NewClassifySetInterfaceL2TablesReply),
	"NewCli": reflect.ValueOf(NewCli),
	"NewCliInband": reflect.ValueOf(NewCliInband),
	"NewCliInbandReply": reflect.ValueOf(NewCliInbandReply),
	"NewCliReply": reflect.ValueOf(NewCliReply),
	"NewControlPing": reflect.ValueOf(NewControlPing),
	"NewControlPingReply": reflect.ValueOf(NewControlPingReply),
	"NewCreateLoopback": reflect.ValueOf(NewCreateLoopback),
	"NewCreateLoopbackInstance": reflect.ValueOf(NewCreateLoopbackInstance),
	"NewCreateLoopbackInstanceReply": reflect.ValueOf(NewCreateLoopbackInstanceReply),
	"NewCreateLoopbackReply": reflect.ValueOf(NewCreateLoopbackReply),
	"NewCreateSubif": reflect.ValueOf(NewCreateSubif),
	"NewCreateSubifReply": reflect.ValueOf(NewCreateSubifReply),
	"NewCreateVlanSubif": reflect.ValueOf(NewCreateVlanSubif),
	"NewCreateVlanSubifReply": reflect.ValueOf(NewCreateVlanSubifReply),
	"NewDeleteLoopback": reflect.ValueOf(NewDeleteLoopback),
	"NewDeleteLoopbackReply": reflect.ValueOf(NewDeleteLoopbackReply),
	"NewDeleteSubif": reflect.ValueOf(NewDeleteSubif),
	"NewDeleteSubifReply": reflect.ValueOf(NewDeleteSubifReply),
	"NewFeatureEnableDisable": reflect.ValueOf(NewFeatureEnableDisable),
	"NewFeatureEnableDisableReply": reflect.ValueOf(NewFeatureEnableDisableReply),
	"NewGetNextIndex": reflect.ValueOf(NewGetNextIndex),
	"NewGetNextIndexReply": reflect.ValueOf(NewGetNextIndexReply),
	"NewGetNodeGraph": reflect.ValueOf(NewGetNodeGraph),
	"NewGetNodeGraphReply": reflect.ValueOf(NewGetNodeGraphReply),
	"NewGetNodeIndex": reflect.ValueOf(NewGetNodeIndex),
	"NewGetNodeIndexReply": reflect.ValueOf(NewGetNodeIndexReply),
	"NewIP4ArpEvent": reflect.ValueOf(NewIP4ArpEvent),
	"NewIP6NdEvent": reflect.ValueOf(NewIP6NdEvent),
	"NewIPSourceAndPortRangeCheckAddDel": reflect.ValueOf(NewIPSourceAndPortRangeCheckAddDel),
	"NewIPSourceAndPortRangeCheckAddDelReply": reflect.ValueOf(NewIPSourceAndPortRangeCheckAddDelReply),
	"NewIPSourceAndPortRangeCheckInterfaceAddDel": reflect.ValueOf(NewIPSourceAndPortRangeCheckInterfaceAddDel),
	"NewIPSourceAndPortRangeCheckInterfaceAddDelReply": reflect.ValueOf(NewIPSourceAndPortRangeCheckInterfaceAddDelReply),
	"NewInputACLSetInterface": reflect.ValueOf(NewInputACLSetInterface),
	"NewInputACLSetInterfaceReply": reflect.ValueOf(NewInputACLSetInterfaceReply),
	"NewInterfaceNameRenumber": reflect.ValueOf(NewInterfaceNameRenumber),
	"NewInterfaceNameRenumberReply": reflect.ValueOf(NewInterfaceNameRenumberReply),
	"NewIoamDisable": reflect.ValueOf(NewIoamDisable),
	"NewIoamDisableReply": reflect.ValueOf(NewIoamDisableReply),
	"NewIoamEnable": reflect.ValueOf(NewIoamEnable),
	"NewIoamEnableReply": reflect.ValueOf(NewIoamEnableReply),
	"NewL2InterfaceEfpFilter": reflect.ValueOf(NewL2InterfaceEfpFilter),
	"NewL2InterfaceEfpFilterReply": reflect.ValueOf(NewL2InterfaceEfpFilterReply),
	"NewL2PatchAddDel": reflect.ValueOf(NewL2PatchAddDel),
	"NewL2PatchAddDelReply": reflect.ValueOf(NewL2PatchAddDelReply),
	"NewOamAddDel": reflect.ValueOf(NewOamAddDel),
	"NewOamAddDelReply": reflect.ValueOf(NewOamAddDelReply),
	"NewOamEvent": reflect.ValueOf(NewOamEvent),
	"NewPgCapture": reflect.ValueOf(NewPgCapture),
	"NewPgCaptureReply": reflect.ValueOf(NewPgCaptureReply),
	"NewPgCreateInterface": reflect.ValueOf(NewPgCreateInterface),
	"NewPgCreateInterfaceReply": reflect.ValueOf(NewPgCreateInterfaceReply),
	"NewPgEnableDisable": reflect.ValueOf(NewPgEnableDisable),
	"NewPgEnableDisableReply": reflect.ValueOf(NewPgEnableDisableReply),
	"NewProxyArpAddDel": reflect.ValueOf(NewProxyArpAddDel),
	"NewProxyArpAddDelReply": reflect.ValueOf(NewProxyArpAddDelReply),
	"NewProxyArpIntfcEnableDisable": reflect.ValueOf(NewProxyArpIntfcEnableDisable),
	"NewProxyArpIntfcEnableDisableReply": reflect.ValueOf(NewProxyArpIntfcEnableDisableReply),
	"NewPunt": reflect.ValueOf(NewPunt),
	"NewPuntReply": reflect.ValueOf(NewPuntReply),
	"NewResetFib": reflect.ValueOf(NewResetFib),
	"NewResetFibReply": reflect.ValueOf(NewResetFibReply),
	"NewResetVrf": reflect.ValueOf(NewResetVrf),
	"NewResetVrfReply": reflect.ValueOf(NewResetVrfReply),
	"NewSetArpNeighborLimit": reflect.ValueOf(NewSetArpNeighborLimit),
	"NewSetArpNeighborLimitReply": reflect.ValueOf(NewSetArpNeighborLimitReply),
	"NewShowVersion": reflect.ValueOf(NewShowVersion),
	"NewShowVersionReply": reflect.ValueOf(NewShowVersionReply),
	"NewSwInterfaceSetL2Bridge": reflect.ValueOf(NewSwInterfaceSetL2Bridge),
	"NewSwInterfaceSetL2BridgeReply": reflect.ValueOf(NewSwInterfaceSetL2BridgeReply),
	"NewSwInterfaceSetL2Xconnect": reflect.ValueOf(NewSwInterfaceSetL2Xconnect),
	"NewSwInterfaceSetL2XconnectReply": reflect.ValueOf(NewSwInterfaceSetL2XconnectReply),
	"NewSwInterfaceSetMplsEnable": reflect.ValueOf(NewSwInterfaceSetMplsEnable),
	"NewSwInterfaceSetMplsEnableReply": reflect.ValueOf(NewSwInterfaceSetMplsEnableReply),
	"NewSwInterfaceSetVpath": reflect.ValueOf(NewSwInterfaceSetVpath),
	"NewSwInterfaceSetVpathReply": reflect.ValueOf(NewSwInterfaceSetVpathReply),
	"NewVnetGetSummaryStats": reflect.ValueOf(NewVnetGetSummaryStats),
	"NewVnetGetSummaryStatsReply": reflect.ValueOf(NewVnetGetSummaryStatsReply),
	"NewVnetIP4FibCounters": reflect.ValueOf(NewVnetIP4FibCounters),
	"NewVnetIP4NbrCounters": reflect.ValueOf(NewVnetIP4NbrCounters),
	"NewVnetIP6FibCounters": reflect.ValueOf(NewVnetIP6FibCounters),
	"NewVnetIP6NbrCounters": reflect.ValueOf(NewVnetIP6NbrCounters),
	"NewWantIP4ArpEvents": reflect.ValueOf(NewWantIP4ArpEvents),
	"NewWantIP4ArpEventsReply": reflect.ValueOf(NewWantIP4ArpEventsReply),
	"NewWantIP6NdEvents": reflect.ValueOf(NewWantIP6NdEvents),
	"NewWantIP6NdEventsReply": reflect.ValueOf(NewWantIP6NdEventsReply),
	"NewWantOamEvents": reflect.ValueOf(NewWantOamEvents),
	"NewWantOamEventsReply": reflect.ValueOf(NewWantOamEventsReply),
	"NewWantStats": reflect.ValueOf(NewWantStats),
	"NewWantStatsReply": reflect.ValueOf(NewWantStatsReply),
}

var Variables = map[string]reflect.Value{
}

var Consts = map[string]reflect.Value{
	"VlAPIVersion": reflect.ValueOf(VlAPIVersion),
}

