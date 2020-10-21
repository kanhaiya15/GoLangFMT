package errs

import (
	"fmt"
	"strings"
)

type Err struct {
	Code    string
	Message string
	URL     string
}

// TODO: Print err url
func (e Err) Error() string {
	return fmt.Sprintf("%s : %s ", e.Code, e.Message)
}

var ERR_DUMMY = Err{
	Code:    "ERR::DUMMY",
	Message: "Dummy error "}

var ERR_INVALID_ENVIRONMENT = Err{
	Code:    "ERR::INV::ENV",
	Message: "Invalid environment specified"}

var ERR_UPD_NF = Err{
	Code:    "ERR::UPD::NF",
	Message: "Tunnel binary update link not found for OS"}

var ERR_CTRL_CONN_MAX_ATTEMPT = Err{
	Code:    "ERR::CTRL::CONN::MAX::ATTEMPT",
	Message: "Control websocket reconnection max attempt reached"}

var ERR_SNK_PRX_MAX_ATTEMPT = Err{
	Code:    "ERR::SNK::PRX::MAX::ATTEMPT",
	Message: "Sink proxy restart max attempt reached"}

var ERR_INF_API_MAX_ATTEMPT = Err{
	Code:    "ERR::INF::API::MAX::ATTEMPT",
	Message: "Info api server restart max attempt reached"}

var ERR_FS_MAX_ATTEMPT = Err{
	Code:    "ERR::FS::MAX::ATTEMPT",
	Message: "File server restart max attempt reached"}

var ERR_INV_WS_DAT_TYPE = Err{
	Code:    "ERR::INV::WS::DAT::TYPE",
	Message: "Invalid data type reader received from websocket"}

func ERR_BIN_UPD(err string) Err {
	return Err{
		Code:    "ERR::BIN::UPD",
		Message: "Unable to update binary " + err}
}

func ERR_WS_CTRL_CONN(err string) Err {
	return Err{
		Code:    "ERR::WS::Conn",
		Message: "Unable to establish control websocket connection " + err}
}

func ERR_WS_CONN(err string) Err {
	return Err{
		Code:    "ERR::WS::Conn",
		Message: "Unable to  establish websocket connection " + err}
}

func ERR_WS_CTRL_CONN_DWN(err string) Err {
	return Err{
		Code:    "ERR::WS::CTRL::CONN::DWN",
		Message: "Control websocket connection closed " + err}
}

func ERR_DAT_CONN_DWN(err string) Err {
	return Err{
		Code:    "ERR::DAT::CONN::DWN",
		Message: "Data websocket connection closed " + err}
}

func ERR_INVALID_WS_URL(err string) Err {
	return Err{
		Code:    "ERR::INV::WS::URL",
		Message: "Invalid websocket url error " + err}
}

func ERR_MARS_TUN_PYLD(err string) Err {
	return Err{
		Code:    "ERR::MARS::TUN::PYLD",
		Message: "Setup tunnel marshalling error " + err}
}

func ERR_LNCH_TUNN_RD_RES(err string) Err {
	return Err{
		Code:    "ERR::LUNCH::TUNN::RD::RES",
		Message: "Launch tunnel failed, Unable to read respone " + err}
}

func ERR_LNCH_TUNN(err string) Err {
	return Err{
		Code:    "ERR::LUNCH::TUNN",
		Message: "Launch tunnel failed " + err}
}

func ERR_STP_TUNN(err string) Err {
	return Err{
		Code:    "ERR::STP::TUNN",
		Message: "Stop tunnel failed " + err}
}

func ERR_SNK_PRX(err string) Err {
	return Err{
		Code:    "ERR::SNK::PRX",
		Message: "Sink proxy failed :  " + err}
}

func ERR_SNK_PRX_CONN(err string) Err {
	return Err{
		Code:    "ERR::SNK::PRX::CONN",
		Message: "Unable to establish connection to local proxy :  " + err}
}

func ERR_WS_WRT(err string) Err {
	return Err{
		Code:    "ERR::WS::WRT",
		Message: "Unable to valid retrieve writer from ws :  " + err}
}

func ERR_WS_RDR(err string) Err {
	return Err{
		Code:    "ERR::WS::RDR",
		Message: "Unable to retrieve valid reader from ws :  " + err}
}

func ERR_ATT_PRX(reqType string, err string) Err {
	return Err{
		Code:    "ERR::ATT::PRX",
		Message: fmt.Sprintf("Unable to attach proxy to [ %s ]request :  %s", reqType, err)}
}

func ERR_DNS_RLV(err string) Err {
	return Err{
		Code:    "ERR::DNS::RLV",
		Message: fmt.Sprintf("Error while resolving dns :  %s", err)}
}

func ERR_VLD_CFG(errs []string) Err {
	return Err{
		Code:    "ERR::CNF::FLD::VLD",
		Message: fmt.Sprintf("Validation errors :  \n%s", strings.Join(errs, "\n"))}
}

func ERR_DAT_WS_RD(err string) Err {
	return Err{
		Code:    "ERR::DAT::WS::RD",
		Message: fmt.Sprintf("Unable to read from websocket :  \n%s", err)}
}

func ERR_SNK_WRT(err string) Err {
	return Err{
		Code:    "ERR::SNK::WRT",
		Message: fmt.Sprintf("Unable to read from websocket :  \n%s", err)}
}

func ERR_API_SRV_STR(err string) Err {
	return Err{
		Code:    "ERR::API::SRV::STR",
		Message: fmt.Sprintf("Unable to start api server :  \n%s", err)}
}

func ERR_FIL_SRV_STR(err string) Err {
	return Err{
		Code:    "ERR::FIL::SRV::STR",
		Message: fmt.Sprintf("Unable to start file server :  \n%s", err)}
}

func ERR_API_WEB_HOK(err string) Err {
	return Err{
		Code:    "ERR::API::WEB::HOK",
		Message: fmt.Sprintf("Unable to call webhook url :  \n%s", err)}
}

var ERR_SNK_RD_WRT_MSM = Err{
	Code:    "ERR::SNK::RD::WRT::MSM",
	Message: "Read write mismatch in sink proxy "}
