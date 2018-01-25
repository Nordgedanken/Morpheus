package messages

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
	"runtime"
	"unsafe"

	"github.com/therecipe/qt"
	std_core "github.com/therecipe/qt/core"
	std_gui "github.com/therecipe/qt/gui"
)

func cGoUnpackString(s C.struct_Moc_PackedString) string {
	if len := int(s.len); len == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}

type Message_ITF interface {
	std_core.QObject_ITF
	Message_PTR() *Message
}

func (ptr *Message) Message_PTR() *Message {
	return ptr
}

func (ptr *Message) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *Message) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromMessage(ptr Message_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.Message_PTR().Pointer()
	}
	return nil
}

func NewMessageFromPointer(ptr unsafe.Pointer) *Message {
	var n *Message
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(Message)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *Message:
			n = deduced

		case *std_core.QObject:
			n = &Message{QObject: *deduced}

		default:
			n = new(Message)
			n.SetPointer(ptr)
		}
	}
	return n
}

//export callbackMessage_Constructor
func callbackMessage_Constructor(ptr unsafe.Pointer) {
	gPtr := NewMessageFromPointer(ptr)
	qt.Register(ptr, gPtr)
}

//export callbackMessage_SetAvatar
func callbackMessage_SetAvatar(ptr unsafe.Pointer, avatar unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "SetAvatar"); signal != nil {
		signal.(func(*std_gui.QPixmap))(std_gui.NewQPixmapFromPointer(avatar))
	}

}

func (ptr *Message) ConnectSetAvatar(f func(avatar *std_gui.QPixmap)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "SetAvatar") {
			C.Message_ConnectSetAvatar(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "SetAvatar"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "SetAvatar", func(avatar *std_gui.QPixmap) {
				signal.(func(*std_gui.QPixmap))(avatar)
				f(avatar)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "SetAvatar", f)
		}
	}
}

func (ptr *Message) DisconnectSetAvatar() {
	if ptr.Pointer() != nil {
		C.Message_DisconnectSetAvatar(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "SetAvatar")
	}
}

func (ptr *Message) SetAvatar(avatar std_gui.QPixmap_ITF) {
	if ptr.Pointer() != nil {
		C.Message_SetAvatar(ptr.Pointer(), std_gui.PointerFromQPixmap(avatar))
	}
}

func Message_QRegisterMetaType() int {
	return int(int32(C.Message_Message_QRegisterMetaType()))
}

func (ptr *Message) QRegisterMetaType() int {
	return int(int32(C.Message_Message_QRegisterMetaType()))
}

func Message_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Message_Message_QRegisterMetaType2(typeNameC)))
}

func (ptr *Message) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Message_Message_QRegisterMetaType2(typeNameC)))
}

func Message_QmlRegisterType() int {
	return int(int32(C.Message_Message_QmlRegisterType()))
}

func (ptr *Message) QmlRegisterType() int {
	return int(int32(C.Message_Message_QmlRegisterType()))
}

func Message_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.Message_Message_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *Message) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.Message_Message_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *Message) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQByteArrayFromPointer(C.Message___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *Message) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.Message___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *Message) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.Message___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *Message) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.Message___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Message) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Message___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Message) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.Message___findChildren_newList2(ptr.Pointer()))
}

func (ptr *Message) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.Message___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Message) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Message___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Message) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.Message___findChildren_newList3(ptr.Pointer()))
}

func (ptr *Message) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.Message___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Message) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Message___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Message) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.Message___findChildren_newList(ptr.Pointer()))
}

func (ptr *Message) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.Message___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Message) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Message___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Message) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.Message___children_newList(ptr.Pointer()))
}

func NewMessage(parent std_core.QObject_ITF) *Message {
	var tmpValue = NewMessageFromPointer(C.Message_NewMessage(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackMessage_DestroyMessage
func callbackMessage_DestroyMessage(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~Message"); signal != nil {
		signal.(func())()
	} else {
		NewMessageFromPointer(ptr).DestroyMessageDefault()
	}
}

func (ptr *Message) ConnectDestroyMessage(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~Message"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~Message", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~Message", f)
		}
	}
}

func (ptr *Message) DisconnectDestroyMessage() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~Message")
	}
}

func (ptr *Message) DestroyMessage() {
	if ptr.Pointer() != nil {
		C.Message_DestroyMessage(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *Message) DestroyMessageDefault() {
	if ptr.Pointer() != nil {
		C.Message_DestroyMessageDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackMessage_Event
func callbackMessage_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMessageFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *Message) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.Message_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackMessage_EventFilter
func callbackMessage_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewMessageFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *Message) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.Message_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackMessage_ChildEvent
func callbackMessage_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewMessageFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *Message) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Message_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackMessage_ConnectNotify
func callbackMessage_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewMessageFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Message) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Message_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackMessage_CustomEvent
func callbackMessage_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewMessageFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *Message) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Message_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackMessage_DeleteLater
func callbackMessage_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewMessageFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *Message) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.Message_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackMessage_Destroyed
func callbackMessage_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackMessage_DisconnectNotify
func callbackMessage_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewMessageFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Message) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Message_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackMessage_ObjectNameChanged
func callbackMessage_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackMessage_TimerEvent
func callbackMessage_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewMessageFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *Message) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Message_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}
