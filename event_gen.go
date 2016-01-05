//
// This file is generated. To change the content of this file, please do not
// apply the change to this file because it will get overwritten. Instead,
// change xenapi.go and execute 'go generate'.
//

package xenAPI

import (
	"fmt"
	"github.com/amfranz/go-xmlrpc-client"
	"reflect"
	"strconv"
	"time"
)

var _ = fmt.Errorf
var _ = xmlrpc.NewClient
var _ = reflect.TypeOf
var _ = strconv.Atoi
var _ = time.UTC

type EventOperation string

const (
	EventOperationAdd EventOperation = "add"
	EventOperationDel EventOperation = "del"
	EventOperationMod EventOperation = "mod"
)

type EventRecord struct {
	ID int
	Timestamp time.Time
	Class string
	Operation EventOperation
	Ref string
	ObjUUID string
}

type EventRef string

type EventClass struct {
	client *Client
}

func (client *Client) Event() EventClass {
	return EventClass{client}
}

func (_class EventClass) Inject(sessionID SessionRef, class string, ref string) (_retval string, _err error) {
	_method := "event.inject"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_classArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "class"), class)
	if _err != nil {
		return
	}
	_refArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "ref"), ref)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _classArg, _refArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

func (_class EventClass) GetCurrentID(sessionID SessionRef) (_retval int, _err error) {
	_method := "event.get_current_id"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}

func (_class EventClass) From(sessionID SessionRef, classes []string, token string, timeout float64) (_retval []EventRecord, _err error) {
	_method := "event.from"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_classesArg, _err := convertStringSetToXen(fmt.Sprintf("%s(%s)", _method, "classes"), classes)
	if _err != nil {
		return
	}
	_tokenArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "token"), token)
	if _err != nil {
		return
	}
	_timeoutArg, _err := convertFloatToXen(fmt.Sprintf("%s(%s)", _method, "timeout"), timeout)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _classesArg, _tokenArg, _timeoutArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEventRecordSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class EventClass) Next(sessionID SessionRef) (_retval []EventRecord, _err error) {
	_method := "event.next"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEventRecordSetToGo(_method + " -> ", _result.Value)
	return
}

func (_class EventClass) Unregister(sessionID SessionRef, classes []string) (_err error) {
	_method := "event.unregister"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_classesArg, _err := convertStringSetToXen(fmt.Sprintf("%s(%s)", _method, "classes"), classes)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _classesArg)
	return
}

func (_class EventClass) Register(sessionID SessionRef, classes []string) (_err error) {
	_method := "event.register"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_classesArg, _err := convertStringSetToXen(fmt.Sprintf("%s(%s)", _method, "classes"), classes)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _classesArg)
	return
}
