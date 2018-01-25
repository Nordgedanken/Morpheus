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
	std_widgets "github.com/therecipe/qt/widgets"
)

func cGoUnpackString(s C.struct_Moc_PackedString) string {
	if len := int(s.len); len == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}

type Room_ITF interface {
	std_widgets.QHBoxLayout_ITF
	Room_PTR() *Room
}

func (ptr *Room) Room_PTR() *Room {
	return ptr
}

func (ptr *Room) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QHBoxLayout_PTR().Pointer()
	}
	return nil
}

func (ptr *Room) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QHBoxLayout_PTR().SetPointer(p)
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

		case *std_widgets.QHBoxLayout:
			n = &Room{QHBoxLayout: *deduced}

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

func NewRoom() *Room {
	var tmpValue = NewRoomFromPointer(C.Room_NewRoom())
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewRoom2(parent std_widgets.QWidget_ITF) *Room {
	var tmpValue = NewRoomFromPointer(C.Room_NewRoom2(std_widgets.PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *Room) DestroyRoom() {
	if ptr.Pointer() != nil {
		C.Room_DestroyRoom(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackRoom_TakeAt
func callbackRoom_TakeAt(ptr unsafe.Pointer, index C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "takeAt"); signal != nil {
		return std_widgets.PointerFromQLayoutItem(signal.(func(int) *std_widgets.QLayoutItem)(int(int32(index))))
	}

	return std_widgets.PointerFromQLayoutItem(NewRoomFromPointer(ptr).TakeAtDefault(int(int32(index))))
}

func (ptr *Room) TakeAtDefault(index int) *std_widgets.QLayoutItem {
	if ptr.Pointer() != nil {
		return std_widgets.NewQLayoutItemFromPointer(C.Room_TakeAtDefault(ptr.Pointer(), C.int(int32(index))))
	}
	return nil
}

//export callbackRoom_AddItem
func callbackRoom_AddItem(ptr unsafe.Pointer, item unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "addItem"); signal != nil {
		signal.(func(*std_widgets.QLayoutItem))(std_widgets.NewQLayoutItemFromPointer(item))
	} else {
		NewRoomFromPointer(ptr).AddItemDefault(std_widgets.NewQLayoutItemFromPointer(item))
	}
}

func (ptr *Room) AddItemDefault(item std_widgets.QLayoutItem_ITF) {
	if ptr.Pointer() != nil {
		C.Room_AddItemDefault(ptr.Pointer(), std_widgets.PointerFromQLayoutItem(item))
	}
}

//export callbackRoom_Invalidate
func callbackRoom_Invalidate(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "invalidate"); signal != nil {
		signal.(func())()
	} else {
		NewRoomFromPointer(ptr).InvalidateDefault()
	}
}

func (ptr *Room) InvalidateDefault() {
	if ptr.Pointer() != nil {
		C.Room_InvalidateDefault(ptr.Pointer())
	}
}

//export callbackRoom_SetGeometry
func callbackRoom_SetGeometry(ptr unsafe.Pointer, r unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setGeometry"); signal != nil {
		signal.(func(*std_core.QRect))(std_core.NewQRectFromPointer(r))
	} else {
		NewRoomFromPointer(ptr).SetGeometryDefault(std_core.NewQRectFromPointer(r))
	}
}

func (ptr *Room) SetGeometryDefault(r std_core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.Room_SetGeometryDefault(ptr.Pointer(), std_core.PointerFromQRect(r))
	}
}

//export callbackRoom_ItemAt
func callbackRoom_ItemAt(ptr unsafe.Pointer, index C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemAt"); signal != nil {
		return std_widgets.PointerFromQLayoutItem(signal.(func(int) *std_widgets.QLayoutItem)(int(int32(index))))
	}

	return std_widgets.PointerFromQLayoutItem(NewRoomFromPointer(ptr).ItemAtDefault(int(int32(index))))
}

func (ptr *Room) ItemAtDefault(index int) *std_widgets.QLayoutItem {
	if ptr.Pointer() != nil {
		return std_widgets.NewQLayoutItemFromPointer(C.Room_ItemAtDefault(ptr.Pointer(), C.int(int32(index))))
	}
	return nil
}

//export callbackRoom_MaximumSize
func callbackRoom_MaximumSize(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "maximumSize"); signal != nil {
		return std_core.PointerFromQSize(signal.(func() *std_core.QSize)())
	}

	return std_core.PointerFromQSize(NewRoomFromPointer(ptr).MaximumSizeDefault())
}

func (ptr *Room) MaximumSizeDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQSizeFromPointer(C.Room_MaximumSizeDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackRoom_MinimumSize
func callbackRoom_MinimumSize(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSize"); signal != nil {
		return std_core.PointerFromQSize(signal.(func() *std_core.QSize)())
	}

	return std_core.PointerFromQSize(NewRoomFromPointer(ptr).MinimumSizeDefault())
}

func (ptr *Room) MinimumSizeDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQSizeFromPointer(C.Room_MinimumSizeDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackRoom_SizeHint
func callbackRoom_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return std_core.PointerFromQSize(signal.(func() *std_core.QSize)())
	}

	return std_core.PointerFromQSize(NewRoomFromPointer(ptr).SizeHintDefault())
}

func (ptr *Room) SizeHintDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQSizeFromPointer(C.Room_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackRoom_ExpandingDirections
func callbackRoom_ExpandingDirections(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "expandingDirections"); signal != nil {
		return C.longlong(signal.(func() std_core.Qt__Orientation)())
	}

	return C.longlong(NewRoomFromPointer(ptr).ExpandingDirectionsDefault())
}

func (ptr *Room) ExpandingDirectionsDefault() std_core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return std_core.Qt__Orientation(C.Room_ExpandingDirectionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackRoom_HasHeightForWidth
func callbackRoom_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewRoomFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *Room) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return C.Room_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackRoom_Count
func callbackRoom_Count(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "count"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(NewRoomFromPointer(ptr).CountDefault()))
}

func (ptr *Room) CountDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.Room_CountDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackRoom_HeightForWidth
func callbackRoom_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewRoomFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *Room) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.Room_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackRoom_MinimumHeightForWidth
func callbackRoom_MinimumHeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "minimumHeightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewRoomFromPointer(ptr).MinimumHeightForWidthDefault(int(int32(w)))))
}

func (ptr *Room) MinimumHeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.Room_MinimumHeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackRoom_Layout
func callbackRoom_Layout(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "layout"); signal != nil {
		return std_widgets.PointerFromQLayout(signal.(func() *std_widgets.QLayout)())
	}

	return std_widgets.PointerFromQLayout(NewRoomFromPointer(ptr).LayoutDefault())
}

func (ptr *Room) LayoutDefault() *std_widgets.QLayout {
	if ptr.Pointer() != nil {
		var tmpValue = std_widgets.NewQLayoutFromPointer(C.Room_LayoutDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackRoom_ChildEvent
func callbackRoom_ChildEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(e))
	} else {
		NewRoomFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(e))
	}
}

func (ptr *Room) ChildEventDefault(e std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.Room_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(e))
	}
}

//export callbackRoom_Geometry
func callbackRoom_Geometry(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "geometry"); signal != nil {
		return std_core.PointerFromQRect(signal.(func() *std_core.QRect)())
	}

	return std_core.PointerFromQRect(NewRoomFromPointer(ptr).GeometryDefault())
}

func (ptr *Room) GeometryDefault() *std_core.QRect {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQRectFromPointer(C.Room_GeometryDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

//export callbackRoom_ControlTypes
func callbackRoom_ControlTypes(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "controlTypes"); signal != nil {
		return C.longlong(signal.(func() std_widgets.QSizePolicy__ControlType)())
	}

	return C.longlong(NewRoomFromPointer(ptr).ControlTypesDefault())
}

func (ptr *Room) ControlTypesDefault() std_widgets.QSizePolicy__ControlType {
	if ptr.Pointer() != nil {
		return std_widgets.QSizePolicy__ControlType(C.Room_ControlTypesDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackRoom_IsEmpty
func callbackRoom_IsEmpty(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isEmpty"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewRoomFromPointer(ptr).IsEmptyDefault())))
}

func (ptr *Room) IsEmptyDefault() bool {
	if ptr.Pointer() != nil {
		return C.Room_IsEmptyDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackRoom_IndexOf
func callbackRoom_IndexOf(ptr unsafe.Pointer, widget unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "indexOf"); signal != nil {
		return C.int(int32(signal.(func(*std_widgets.QWidget) int)(std_widgets.NewQWidgetFromPointer(widget))))
	}

	return C.int(int32(NewRoomFromPointer(ptr).IndexOfDefault(std_widgets.NewQWidgetFromPointer(widget))))
}

func (ptr *Room) IndexOfDefault(widget std_widgets.QWidget_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.Room_IndexOfDefault(ptr.Pointer(), std_widgets.PointerFromQWidget(widget))))
	}
	return 0
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

//export callbackRoom_SpacerItem
func callbackRoom_SpacerItem(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "spacerItem"); signal != nil {
		return std_widgets.PointerFromQSpacerItem(signal.(func() *std_widgets.QSpacerItem)())
	}

	return std_widgets.PointerFromQSpacerItem(NewRoomFromPointer(ptr).SpacerItemDefault())
}

func (ptr *Room) SpacerItemDefault() *std_widgets.QSpacerItem {
	if ptr.Pointer() != nil {
		return std_widgets.NewQSpacerItemFromPointer(C.Room_SpacerItemDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackRoom_Widget
func callbackRoom_Widget(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "widget"); signal != nil {
		return std_widgets.PointerFromQWidget(signal.(func() *std_widgets.QWidget)())
	}

	return std_widgets.PointerFromQWidget(NewRoomFromPointer(ptr).WidgetDefault())
}

func (ptr *Room) WidgetDefault() *std_widgets.QWidget {
	if ptr.Pointer() != nil {
		var tmpValue = std_widgets.NewQWidgetFromPointer(C.Room_WidgetDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}
