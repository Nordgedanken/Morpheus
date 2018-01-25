package rooms

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

type Room_ITF interface {
	std_core.QObject_ITF
	Room_PTR() *Room
}

func (ptr *Room) Room_PTR() *Room {
	return ptr
}

func (ptr *Room) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *Room) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromRoom(ptr Room_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.Room_PTR().Pointer()
	}
	return nil
}

func NewRoomFromPointer(ptr unsafe.Pointer) *Room {
	var n *Room
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(Room)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *Room:
			n = deduced

		case *std_core.QObject:
			n = &Room{QObject: *deduced}

		default:
			n = new(Room)
			n.SetPointer(ptr)
		}
	}
	return n
}

//export callbackRoom_Constructor
func callbackRoom_Constructor(ptr unsafe.Pointer) {
	gPtr := NewRoomFromPointer(ptr)
	qt.Register(ptr, gPtr)
}

//export callbackRoom_SetAvatar
func callbackRoom_SetAvatar(ptr unsafe.Pointer, roomAvatar unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "SetAvatar"); signal != nil {
		signal.(func(*std_gui.QPixmap))(std_gui.NewQPixmapFromPointer(roomAvatar))
	}

}

func (ptr *Room) ConnectSetAvatar(f func(roomAvatar *std_gui.QPixmap)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "SetAvatar") {
			C.Room_ConnectSetAvatar(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "SetAvatar"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "SetAvatar", func(roomAvatar *std_gui.QPixmap) {
				signal.(func(*std_gui.QPixmap))(roomAvatar)
				f(roomAvatar)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "SetAvatar", f)
		}
	}
}

func (ptr *Room) DisconnectSetAvatar() {
	if ptr.Pointer() != nil {
		C.Room_DisconnectSetAvatar(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "SetAvatar")
	}
}

func (ptr *Room) SetAvatar(roomAvatar std_gui.QPixmap_ITF) {
	if ptr.Pointer() != nil {
		C.Room_SetAvatar(ptr.Pointer(), std_gui.PointerFromQPixmap(roomAvatar))
	}
}

func Room_QRegisterMetaType() int {
	return int(int32(C.Room_Room_QRegisterMetaType()))
}

func (ptr *Room) QRegisterMetaType() int {
	return int(int32(C.Room_Room_QRegisterMetaType()))
}

func Room_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Room_Room_QRegisterMetaType2(typeNameC)))
}

func (ptr *Room) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.Room_Room_QRegisterMetaType2(typeNameC)))
}

func Room_QmlRegisterType() int {
	return int(int32(C.Room_Room_QmlRegisterType()))
}

func (ptr *Room) QmlRegisterType() int {
	return int(int32(C.Room_Room_QmlRegisterType()))
}

func Room_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.Room_Room_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *Room) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.Room_Room_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *Room) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQByteArrayFromPointer(C.Room___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *Room) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.Room___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *Room) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.Room___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *Room) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.Room___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Room) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Room___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Room) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.Room___findChildren_newList2(ptr.Pointer()))
}

func (ptr *Room) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.Room___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Room) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Room___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Room) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.Room___findChildren_newList3(ptr.Pointer()))
}

func (ptr *Room) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.Room___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Room) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Room___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Room) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.Room___findChildren_newList(ptr.Pointer()))
}

func (ptr *Room) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.Room___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *Room) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.Room___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *Room) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.Room___children_newList(ptr.Pointer()))
}

func NewRoom(parent std_core.QObject_ITF) *Room {
	var tmpValue = NewRoomFromPointer(C.Room_NewRoom(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackRoom_DestroyRoom
func callbackRoom_DestroyRoom(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~Room"); signal != nil {
		signal.(func())()
	} else {
		NewRoomFromPointer(ptr).DestroyRoomDefault()
	}
}

func (ptr *Room) ConnectDestroyRoom(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~Room"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~Room", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~Room", f)
		}
	}
}

func (ptr *Room) DisconnectDestroyRoom() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~Room")
	}
}

func (ptr *Room) DestroyRoom() {
	if ptr.Pointer() != nil {
		C.Room_DestroyRoom(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *Room) DestroyRoomDefault() {
	if ptr.Pointer() != nil {
		C.Room_DestroyRoomDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackRoom_Event
func callbackRoom_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewRoomFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *Room) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.Room_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackRoom_EventFilter
func callbackRoom_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewRoomFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *Room) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.Room_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackRoom_ChildEvent
func callbackRoom_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewRoomFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *Room) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Room_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackRoom_ConnectNotify
func callbackRoom_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewRoomFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Room) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Room_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackRoom_CustomEvent
func callbackRoom_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewRoomFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *Room) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Room_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackRoom_DeleteLater
func callbackRoom_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewRoomFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *Room) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.Room_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackRoom_Destroyed
func callbackRoom_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackRoom_DisconnectNotify
func callbackRoom_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewRoomFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *Room) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.Room_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackRoom_ObjectNameChanged
func callbackRoom_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackRoom_TimerEvent
func callbackRoom_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewRoomFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *Room) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Room_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}
