package listLayouts

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
	std_widgets "github.com/therecipe/qt/widgets"
)

func cGoUnpackString(s C.struct_Moc_PackedString) string {
	if len := int(s.len); len == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}

type QRoomLayout_ITF interface {
	std_widgets.QHBoxLayout_ITF
	QRoomLayout_PTR() *QRoomLayout
}

func (ptr *QRoomLayout) QRoomLayout_PTR() *QRoomLayout {
	return ptr
}

func (ptr *QRoomLayout) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QHBoxLayout_PTR().Pointer()
	}
	return nil
}

func (ptr *QRoomLayout) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QHBoxLayout_PTR().SetPointer(p)
	}
}

func PointerFromQRoomLayout(ptr QRoomLayout_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRoomLayout_PTR().Pointer()
	}
	return nil
}

func NewQRoomLayoutFromPointer(ptr unsafe.Pointer) *QRoomLayout {
	var n *QRoomLayout
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(QRoomLayout)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *QRoomLayout:
			n = deduced

		case *std_widgets.QHBoxLayout:
			n = &QRoomLayout{QHBoxLayout: *deduced}

		default:
			n = new(QRoomLayout)
			n.SetPointer(ptr)
		}
	}
	return n
}

//export callbackQRoomLayout_Constructor
func callbackQRoomLayout_Constructor(ptr unsafe.Pointer) {
	gPtr := NewQRoomLayoutFromPointer(ptr)
	qt.Register(ptr, gPtr)
}

//export callbackQRoomLayout_SetAvatar
func callbackQRoomLayout_SetAvatar(ptr unsafe.Pointer, roomID C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "SetAvatar"); signal != nil {
		signal.(func(string))(cGoUnpackString(roomID))
	}

}

func (ptr *QRoomLayout) ConnectSetAvatar(f func(roomID string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "SetAvatar") {
			C.QRoomLayout_ConnectSetAvatar(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "SetAvatar"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "SetAvatar", func(roomID string) {
				signal.(func(string))(roomID)
				f(roomID)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "SetAvatar", f)
		}
	}
}

func (ptr *QRoomLayout) DisconnectSetAvatar() {
	if ptr.Pointer() != nil {
		C.QRoomLayout_DisconnectSetAvatar(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "SetAvatar")
	}
}

func (ptr *QRoomLayout) SetAvatar(roomID string) {
	if ptr.Pointer() != nil {
		var roomIDC *C.char
		if roomID != "" {
			roomIDC = C.CString(roomID)
			defer C.free(unsafe.Pointer(roomIDC))
		}
		C.QRoomLayout_SetAvatar(ptr.Pointer(), C.struct_Moc_PackedString{data: roomIDC, len: C.longlong(len(roomID))})
	}
}

func QRoomLayout_QRegisterMetaType() int {
	return int(int32(C.QRoomLayout_QRoomLayout_QRegisterMetaType()))
}

func (ptr *QRoomLayout) QRegisterMetaType() int {
	return int(int32(C.QRoomLayout_QRoomLayout_QRegisterMetaType()))
}

func QRoomLayout_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.QRoomLayout_QRoomLayout_QRegisterMetaType2(typeNameC)))
}

func (ptr *QRoomLayout) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.QRoomLayout_QRoomLayout_QRegisterMetaType2(typeNameC)))
}

func QRoomLayout_QmlRegisterType() int {
	return int(int32(C.QRoomLayout_QRoomLayout_QmlRegisterType()))
}

func (ptr *QRoomLayout) QmlRegisterType() int {
	return int(int32(C.QRoomLayout_QRoomLayout_QmlRegisterType()))
}

func QRoomLayout_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.QRoomLayout_QRoomLayout_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QRoomLayout) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.QRoomLayout_QRoomLayout_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QRoomLayout) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQByteArrayFromPointer(C.QRoomLayout___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QRoomLayout) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QRoomLayout___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *QRoomLayout) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QRoomLayout___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QRoomLayout) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.QRoomLayout___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QRoomLayout) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QRoomLayout___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QRoomLayout) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QRoomLayout___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QRoomLayout) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.QRoomLayout___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QRoomLayout) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QRoomLayout___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QRoomLayout) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QRoomLayout___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QRoomLayout) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.QRoomLayout___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QRoomLayout) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QRoomLayout___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QRoomLayout) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QRoomLayout___findChildren_newList(ptr.Pointer()))
}

func (ptr *QRoomLayout) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.QRoomLayout___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QRoomLayout) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QRoomLayout___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QRoomLayout) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QRoomLayout___children_newList(ptr.Pointer()))
}

func NewQRoomLayout() *QRoomLayout {
	var tmpValue = NewQRoomLayoutFromPointer(C.QRoomLayout_NewQRoomLayout())
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQRoomLayout2(parent std_widgets.QWidget_ITF) *QRoomLayout {
	var tmpValue = NewQRoomLayoutFromPointer(C.QRoomLayout_NewQRoomLayout2(std_widgets.PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QRoomLayout) DestroyQRoomLayout() {
	if ptr.Pointer() != nil {
		C.QRoomLayout_DestroyQRoomLayout(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQRoomLayout_TakeAt
func callbackQRoomLayout_TakeAt(ptr unsafe.Pointer, index C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "takeAt"); signal != nil {
		return std_widgets.PointerFromQLayoutItem(signal.(func(int) *std_widgets.QLayoutItem)(int(int32(index))))
	}

	return std_widgets.PointerFromQLayoutItem(NewQRoomLayoutFromPointer(ptr).TakeAtDefault(int(int32(index))))
}

func (ptr *QRoomLayout) TakeAtDefault(index int) *std_widgets.QLayoutItem {
	if ptr.Pointer() != nil {
		return std_widgets.NewQLayoutItemFromPointer(C.QRoomLayout_TakeAtDefault(ptr.Pointer(), C.int(int32(index))))
	}
	return nil
}

//export callbackQRoomLayout_AddItem
func callbackQRoomLayout_AddItem(ptr unsafe.Pointer, item unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "addItem"); signal != nil {
		signal.(func(*std_widgets.QLayoutItem))(std_widgets.NewQLayoutItemFromPointer(item))
	} else {
		NewQRoomLayoutFromPointer(ptr).AddItemDefault(std_widgets.NewQLayoutItemFromPointer(item))
	}
}

func (ptr *QRoomLayout) AddItemDefault(item std_widgets.QLayoutItem_ITF) {
	if ptr.Pointer() != nil {
		C.QRoomLayout_AddItemDefault(ptr.Pointer(), std_widgets.PointerFromQLayoutItem(item))
	}
}

//export callbackQRoomLayout_Invalidate
func callbackQRoomLayout_Invalidate(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "invalidate"); signal != nil {
		signal.(func())()
	} else {
		NewQRoomLayoutFromPointer(ptr).InvalidateDefault()
	}
}

func (ptr *QRoomLayout) InvalidateDefault() {
	if ptr.Pointer() != nil {
		C.QRoomLayout_InvalidateDefault(ptr.Pointer())
	}
}

//export callbackQRoomLayout_SetGeometry
func callbackQRoomLayout_SetGeometry(ptr unsafe.Pointer, r unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setGeometry"); signal != nil {
		signal.(func(*std_core.QRect))(std_core.NewQRectFromPointer(r))
	} else {
		NewQRoomLayoutFromPointer(ptr).SetGeometryDefault(std_core.NewQRectFromPointer(r))
	}
}

func (ptr *QRoomLayout) SetGeometryDefault(r std_core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QRoomLayout_SetGeometryDefault(ptr.Pointer(), std_core.PointerFromQRect(r))
	}
}

//export callbackQRoomLayout_ItemAt
func callbackQRoomLayout_ItemAt(ptr unsafe.Pointer, index C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemAt"); signal != nil {
		return std_widgets.PointerFromQLayoutItem(signal.(func(int) *std_widgets.QLayoutItem)(int(int32(index))))
	}

	return std_widgets.PointerFromQLayoutItem(NewQRoomLayoutFromPointer(ptr).ItemAtDefault(int(int32(index))))
}

func (ptr *QRoomLayout) ItemAtDefault(index int) *std_widgets.QLayoutItem {
	if ptr.Pointer() != nil {
		return std_widgets.NewQLayoutItemFromPointer(C.QRoomLayout_ItemAtDefault(ptr.Pointer(), C.int(int32(index))))
	}
	return nil
}

//export callbackQRoomLayout_MaximumSize
func callbackQRoomLayout_MaximumSize(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "maximumSize"); signal != nil {
		return std_core.PointerFromQSize(signal.(func() *std_core.QSize)())
	}

	return std_core.PointerFromQSize(NewQRoomLayoutFromPointer(ptr).MaximumSizeDefault())
}

func (ptr *QRoomLayout) MaximumSizeDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQSizeFromPointer(C.QRoomLayout_MaximumSizeDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQRoomLayout_MinimumSize
func callbackQRoomLayout_MinimumSize(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSize"); signal != nil {
		return std_core.PointerFromQSize(signal.(func() *std_core.QSize)())
	}

	return std_core.PointerFromQSize(NewQRoomLayoutFromPointer(ptr).MinimumSizeDefault())
}

func (ptr *QRoomLayout) MinimumSizeDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQSizeFromPointer(C.QRoomLayout_MinimumSizeDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQRoomLayout_SizeHint
func callbackQRoomLayout_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return std_core.PointerFromQSize(signal.(func() *std_core.QSize)())
	}

	return std_core.PointerFromQSize(NewQRoomLayoutFromPointer(ptr).SizeHintDefault())
}

func (ptr *QRoomLayout) SizeHintDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQSizeFromPointer(C.QRoomLayout_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQRoomLayout_ExpandingDirections
func callbackQRoomLayout_ExpandingDirections(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "expandingDirections"); signal != nil {
		return C.longlong(signal.(func() std_core.Qt__Orientation)())
	}

	return C.longlong(NewQRoomLayoutFromPointer(ptr).ExpandingDirectionsDefault())
}

func (ptr *QRoomLayout) ExpandingDirectionsDefault() std_core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return std_core.Qt__Orientation(C.QRoomLayout_ExpandingDirectionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQRoomLayout_HasHeightForWidth
func callbackQRoomLayout_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQRoomLayoutFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QRoomLayout) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return C.QRoomLayout_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQRoomLayout_Count
func callbackQRoomLayout_Count(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "count"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(NewQRoomLayoutFromPointer(ptr).CountDefault()))
}

func (ptr *QRoomLayout) CountDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QRoomLayout_CountDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackQRoomLayout_HeightForWidth
func callbackQRoomLayout_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQRoomLayoutFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QRoomLayout) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QRoomLayout_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQRoomLayout_MinimumHeightForWidth
func callbackQRoomLayout_MinimumHeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "minimumHeightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQRoomLayoutFromPointer(ptr).MinimumHeightForWidthDefault(int(int32(w)))))
}

func (ptr *QRoomLayout) MinimumHeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QRoomLayout_MinimumHeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQRoomLayout_Layout
func callbackQRoomLayout_Layout(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "layout"); signal != nil {
		return std_widgets.PointerFromQLayout(signal.(func() *std_widgets.QLayout)())
	}

	return std_widgets.PointerFromQLayout(NewQRoomLayoutFromPointer(ptr).LayoutDefault())
}

func (ptr *QRoomLayout) LayoutDefault() *std_widgets.QLayout {
	if ptr.Pointer() != nil {
		var tmpValue = std_widgets.NewQLayoutFromPointer(C.QRoomLayout_LayoutDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQRoomLayout_ChildEvent
func callbackQRoomLayout_ChildEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(e))
	} else {
		NewQRoomLayoutFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(e))
	}
}

func (ptr *QRoomLayout) ChildEventDefault(e std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRoomLayout_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(e))
	}
}

//export callbackQRoomLayout_Geometry
func callbackQRoomLayout_Geometry(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "geometry"); signal != nil {
		return std_core.PointerFromQRect(signal.(func() *std_core.QRect)())
	}

	return std_core.PointerFromQRect(NewQRoomLayoutFromPointer(ptr).GeometryDefault())
}

func (ptr *QRoomLayout) GeometryDefault() *std_core.QRect {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQRectFromPointer(C.QRoomLayout_GeometryDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

//export callbackQRoomLayout_ControlTypes
func callbackQRoomLayout_ControlTypes(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "controlTypes"); signal != nil {
		return C.longlong(signal.(func() std_widgets.QSizePolicy__ControlType)())
	}

	return C.longlong(NewQRoomLayoutFromPointer(ptr).ControlTypesDefault())
}

func (ptr *QRoomLayout) ControlTypesDefault() std_widgets.QSizePolicy__ControlType {
	if ptr.Pointer() != nil {
		return std_widgets.QSizePolicy__ControlType(C.QRoomLayout_ControlTypesDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQRoomLayout_IsEmpty
func callbackQRoomLayout_IsEmpty(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isEmpty"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQRoomLayoutFromPointer(ptr).IsEmptyDefault())))
}

func (ptr *QRoomLayout) IsEmptyDefault() bool {
	if ptr.Pointer() != nil {
		return C.QRoomLayout_IsEmptyDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQRoomLayout_IndexOf
func callbackQRoomLayout_IndexOf(ptr unsafe.Pointer, widget unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "indexOf"); signal != nil {
		return C.int(int32(signal.(func(*std_widgets.QWidget) int)(std_widgets.NewQWidgetFromPointer(widget))))
	}

	return C.int(int32(NewQRoomLayoutFromPointer(ptr).IndexOfDefault(std_widgets.NewQWidgetFromPointer(widget))))
}

func (ptr *QRoomLayout) IndexOfDefault(widget std_widgets.QWidget_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QRoomLayout_IndexOfDefault(ptr.Pointer(), std_widgets.PointerFromQWidget(widget))))
	}
	return 0
}

//export callbackQRoomLayout_Event
func callbackQRoomLayout_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQRoomLayoutFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *QRoomLayout) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QRoomLayout_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQRoomLayout_EventFilter
func callbackQRoomLayout_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQRoomLayoutFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *QRoomLayout) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QRoomLayout_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQRoomLayout_ConnectNotify
func callbackQRoomLayout_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQRoomLayoutFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QRoomLayout) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QRoomLayout_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQRoomLayout_CustomEvent
func callbackQRoomLayout_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewQRoomLayoutFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *QRoomLayout) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRoomLayout_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackQRoomLayout_DeleteLater
func callbackQRoomLayout_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQRoomLayoutFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QRoomLayout) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QRoomLayout_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQRoomLayout_Destroyed
func callbackQRoomLayout_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackQRoomLayout_DisconnectNotify
func callbackQRoomLayout_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQRoomLayoutFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QRoomLayout) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QRoomLayout_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQRoomLayout_ObjectNameChanged
func callbackQRoomLayout_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQRoomLayout_TimerEvent
func callbackQRoomLayout_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewQRoomLayoutFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QRoomLayout) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRoomLayout_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

//export callbackQRoomLayout_SpacerItem
func callbackQRoomLayout_SpacerItem(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "spacerItem"); signal != nil {
		return std_widgets.PointerFromQSpacerItem(signal.(func() *std_widgets.QSpacerItem)())
	}

	return std_widgets.PointerFromQSpacerItem(NewQRoomLayoutFromPointer(ptr).SpacerItemDefault())
}

func (ptr *QRoomLayout) SpacerItemDefault() *std_widgets.QSpacerItem {
	if ptr.Pointer() != nil {
		return std_widgets.NewQSpacerItemFromPointer(C.QRoomLayout_SpacerItemDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQRoomLayout_Widget
func callbackQRoomLayout_Widget(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "widget"); signal != nil {
		return std_widgets.PointerFromQWidget(signal.(func() *std_widgets.QWidget)())
	}

	return std_widgets.PointerFromQWidget(NewQRoomLayoutFromPointer(ptr).WidgetDefault())
}

func (ptr *QRoomLayout) WidgetDefault() *std_widgets.QWidget {
	if ptr.Pointer() != nil {
		var tmpValue = std_widgets.NewQWidgetFromPointer(C.QRoomLayout_WidgetDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

type QRoomVBoxLayoutWithTriggerSlot_ITF interface {
	std_widgets.QVBoxLayout_ITF
	QRoomVBoxLayoutWithTriggerSlot_PTR() *QRoomVBoxLayoutWithTriggerSlot
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) QRoomVBoxLayoutWithTriggerSlot_PTR() *QRoomVBoxLayoutWithTriggerSlot {
	return ptr
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QVBoxLayout_PTR().Pointer()
	}
	return nil
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QVBoxLayout_PTR().SetPointer(p)
	}
}

func PointerFromQRoomVBoxLayoutWithTriggerSlot(ptr QRoomVBoxLayoutWithTriggerSlot_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRoomVBoxLayoutWithTriggerSlot_PTR().Pointer()
	}
	return nil
}

func NewQRoomVBoxLayoutWithTriggerSlotFromPointer(ptr unsafe.Pointer) *QRoomVBoxLayoutWithTriggerSlot {
	var n *QRoomVBoxLayoutWithTriggerSlot
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(QRoomVBoxLayoutWithTriggerSlot)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *QRoomVBoxLayoutWithTriggerSlot:
			n = deduced

		case *std_widgets.QVBoxLayout:
			n = &QRoomVBoxLayoutWithTriggerSlot{QVBoxLayout: *deduced}

		default:
			n = new(QRoomVBoxLayoutWithTriggerSlot)
			n.SetPointer(ptr)
		}
	}
	return n
}

//export callbackQRoomVBoxLayoutWithTriggerSlot_Constructor
func callbackQRoomVBoxLayoutWithTriggerSlot_Constructor(ptr unsafe.Pointer) {
	gPtr := NewQRoomVBoxLayoutWithTriggerSlotFromPointer(ptr)
	qt.Register(ptr, gPtr)
}

//export callbackQRoomVBoxLayoutWithTriggerSlot_TriggerRoom
func callbackQRoomVBoxLayoutWithTriggerSlot_TriggerRoom(ptr unsafe.Pointer, roomID C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "TriggerRoom"); signal != nil {
		signal.(func(string))(cGoUnpackString(roomID))
	}

}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) ConnectTriggerRoom(f func(roomID string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "TriggerRoom") {
			C.QRoomVBoxLayoutWithTriggerSlot_ConnectTriggerRoom(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "TriggerRoom"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "TriggerRoom", func(roomID string) {
				signal.(func(string))(roomID)
				f(roomID)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "TriggerRoom", f)
		}
	}
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) DisconnectTriggerRoom() {
	if ptr.Pointer() != nil {
		C.QRoomVBoxLayoutWithTriggerSlot_DisconnectTriggerRoom(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "TriggerRoom")
	}
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) TriggerRoom(roomID string) {
	if ptr.Pointer() != nil {
		var roomIDC *C.char
		if roomID != "" {
			roomIDC = C.CString(roomID)
			defer C.free(unsafe.Pointer(roomIDC))
		}
		C.QRoomVBoxLayoutWithTriggerSlot_TriggerRoom(ptr.Pointer(), C.struct_Moc_PackedString{data: roomIDC, len: C.longlong(len(roomID))})
	}
}

//export callbackQRoomVBoxLayoutWithTriggerSlot_ChangeRoom
func callbackQRoomVBoxLayoutWithTriggerSlot_ChangeRoom(ptr unsafe.Pointer, roomID C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "ChangeRoom"); signal != nil {
		signal.(func(string))(cGoUnpackString(roomID))
	}

}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) ConnectChangeRoom(f func(roomID string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "ChangeRoom") {
			C.QRoomVBoxLayoutWithTriggerSlot_ConnectChangeRoom(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "ChangeRoom"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "ChangeRoom", func(roomID string) {
				signal.(func(string))(roomID)
				f(roomID)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "ChangeRoom", f)
		}
	}
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) DisconnectChangeRoom() {
	if ptr.Pointer() != nil {
		C.QRoomVBoxLayoutWithTriggerSlot_DisconnectChangeRoom(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "ChangeRoom")
	}
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) ChangeRoom(roomID string) {
	if ptr.Pointer() != nil {
		var roomIDC *C.char
		if roomID != "" {
			roomIDC = C.CString(roomID)
			defer C.free(unsafe.Pointer(roomIDC))
		}
		C.QRoomVBoxLayoutWithTriggerSlot_ChangeRoom(ptr.Pointer(), C.struct_Moc_PackedString{data: roomIDC, len: C.longlong(len(roomID))})
	}
}

func QRoomVBoxLayoutWithTriggerSlot_QRegisterMetaType() int {
	return int(int32(C.QRoomVBoxLayoutWithTriggerSlot_QRoomVBoxLayoutWithTriggerSlot_QRegisterMetaType()))
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) QRegisterMetaType() int {
	return int(int32(C.QRoomVBoxLayoutWithTriggerSlot_QRoomVBoxLayoutWithTriggerSlot_QRegisterMetaType()))
}

func QRoomVBoxLayoutWithTriggerSlot_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.QRoomVBoxLayoutWithTriggerSlot_QRoomVBoxLayoutWithTriggerSlot_QRegisterMetaType2(typeNameC)))
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.QRoomVBoxLayoutWithTriggerSlot_QRoomVBoxLayoutWithTriggerSlot_QRegisterMetaType2(typeNameC)))
}

func QRoomVBoxLayoutWithTriggerSlot_QmlRegisterType() int {
	return int(int32(C.QRoomVBoxLayoutWithTriggerSlot_QRoomVBoxLayoutWithTriggerSlot_QmlRegisterType()))
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) QmlRegisterType() int {
	return int(int32(C.QRoomVBoxLayoutWithTriggerSlot_QRoomVBoxLayoutWithTriggerSlot_QmlRegisterType()))
}

func QRoomVBoxLayoutWithTriggerSlot_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.QRoomVBoxLayoutWithTriggerSlot_QRoomVBoxLayoutWithTriggerSlot_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.QRoomVBoxLayoutWithTriggerSlot_QRoomVBoxLayoutWithTriggerSlot_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQByteArrayFromPointer(C.QRoomVBoxLayoutWithTriggerSlot___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QRoomVBoxLayoutWithTriggerSlot___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QRoomVBoxLayoutWithTriggerSlot___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.QRoomVBoxLayoutWithTriggerSlot___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QRoomVBoxLayoutWithTriggerSlot___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QRoomVBoxLayoutWithTriggerSlot___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.QRoomVBoxLayoutWithTriggerSlot___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QRoomVBoxLayoutWithTriggerSlot___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QRoomVBoxLayoutWithTriggerSlot___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.QRoomVBoxLayoutWithTriggerSlot___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QRoomVBoxLayoutWithTriggerSlot___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QRoomVBoxLayoutWithTriggerSlot___findChildren_newList(ptr.Pointer()))
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.QRoomVBoxLayoutWithTriggerSlot___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QRoomVBoxLayoutWithTriggerSlot___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QRoomVBoxLayoutWithTriggerSlot___children_newList(ptr.Pointer()))
}

func NewQRoomVBoxLayoutWithTriggerSlot() *QRoomVBoxLayoutWithTriggerSlot {
	var tmpValue = NewQRoomVBoxLayoutWithTriggerSlotFromPointer(C.QRoomVBoxLayoutWithTriggerSlot_NewQRoomVBoxLayoutWithTriggerSlot())
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQRoomVBoxLayoutWithTriggerSlot2(parent std_widgets.QWidget_ITF) *QRoomVBoxLayoutWithTriggerSlot {
	var tmpValue = NewQRoomVBoxLayoutWithTriggerSlotFromPointer(C.QRoomVBoxLayoutWithTriggerSlot_NewQRoomVBoxLayoutWithTriggerSlot2(std_widgets.PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) DestroyQRoomVBoxLayoutWithTriggerSlot() {
	if ptr.Pointer() != nil {
		C.QRoomVBoxLayoutWithTriggerSlot_DestroyQRoomVBoxLayoutWithTriggerSlot(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQRoomVBoxLayoutWithTriggerSlot_TakeAt
func callbackQRoomVBoxLayoutWithTriggerSlot_TakeAt(ptr unsafe.Pointer, index C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "takeAt"); signal != nil {
		return std_widgets.PointerFromQLayoutItem(signal.(func(int) *std_widgets.QLayoutItem)(int(int32(index))))
	}

	return std_widgets.PointerFromQLayoutItem(NewQRoomVBoxLayoutWithTriggerSlotFromPointer(ptr).TakeAtDefault(int(int32(index))))
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) TakeAtDefault(index int) *std_widgets.QLayoutItem {
	if ptr.Pointer() != nil {
		return std_widgets.NewQLayoutItemFromPointer(C.QRoomVBoxLayoutWithTriggerSlot_TakeAtDefault(ptr.Pointer(), C.int(int32(index))))
	}
	return nil
}

//export callbackQRoomVBoxLayoutWithTriggerSlot_AddItem
func callbackQRoomVBoxLayoutWithTriggerSlot_AddItem(ptr unsafe.Pointer, item unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "addItem"); signal != nil {
		signal.(func(*std_widgets.QLayoutItem))(std_widgets.NewQLayoutItemFromPointer(item))
	} else {
		NewQRoomVBoxLayoutWithTriggerSlotFromPointer(ptr).AddItemDefault(std_widgets.NewQLayoutItemFromPointer(item))
	}
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) AddItemDefault(item std_widgets.QLayoutItem_ITF) {
	if ptr.Pointer() != nil {
		C.QRoomVBoxLayoutWithTriggerSlot_AddItemDefault(ptr.Pointer(), std_widgets.PointerFromQLayoutItem(item))
	}
}

//export callbackQRoomVBoxLayoutWithTriggerSlot_Invalidate
func callbackQRoomVBoxLayoutWithTriggerSlot_Invalidate(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "invalidate"); signal != nil {
		signal.(func())()
	} else {
		NewQRoomVBoxLayoutWithTriggerSlotFromPointer(ptr).InvalidateDefault()
	}
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) InvalidateDefault() {
	if ptr.Pointer() != nil {
		C.QRoomVBoxLayoutWithTriggerSlot_InvalidateDefault(ptr.Pointer())
	}
}

//export callbackQRoomVBoxLayoutWithTriggerSlot_SetGeometry
func callbackQRoomVBoxLayoutWithTriggerSlot_SetGeometry(ptr unsafe.Pointer, r unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setGeometry"); signal != nil {
		signal.(func(*std_core.QRect))(std_core.NewQRectFromPointer(r))
	} else {
		NewQRoomVBoxLayoutWithTriggerSlotFromPointer(ptr).SetGeometryDefault(std_core.NewQRectFromPointer(r))
	}
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) SetGeometryDefault(r std_core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QRoomVBoxLayoutWithTriggerSlot_SetGeometryDefault(ptr.Pointer(), std_core.PointerFromQRect(r))
	}
}

//export callbackQRoomVBoxLayoutWithTriggerSlot_ItemAt
func callbackQRoomVBoxLayoutWithTriggerSlot_ItemAt(ptr unsafe.Pointer, index C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemAt"); signal != nil {
		return std_widgets.PointerFromQLayoutItem(signal.(func(int) *std_widgets.QLayoutItem)(int(int32(index))))
	}

	return std_widgets.PointerFromQLayoutItem(NewQRoomVBoxLayoutWithTriggerSlotFromPointer(ptr).ItemAtDefault(int(int32(index))))
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) ItemAtDefault(index int) *std_widgets.QLayoutItem {
	if ptr.Pointer() != nil {
		return std_widgets.NewQLayoutItemFromPointer(C.QRoomVBoxLayoutWithTriggerSlot_ItemAtDefault(ptr.Pointer(), C.int(int32(index))))
	}
	return nil
}

//export callbackQRoomVBoxLayoutWithTriggerSlot_MaximumSize
func callbackQRoomVBoxLayoutWithTriggerSlot_MaximumSize(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "maximumSize"); signal != nil {
		return std_core.PointerFromQSize(signal.(func() *std_core.QSize)())
	}

	return std_core.PointerFromQSize(NewQRoomVBoxLayoutWithTriggerSlotFromPointer(ptr).MaximumSizeDefault())
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) MaximumSizeDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQSizeFromPointer(C.QRoomVBoxLayoutWithTriggerSlot_MaximumSizeDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQRoomVBoxLayoutWithTriggerSlot_MinimumSize
func callbackQRoomVBoxLayoutWithTriggerSlot_MinimumSize(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSize"); signal != nil {
		return std_core.PointerFromQSize(signal.(func() *std_core.QSize)())
	}

	return std_core.PointerFromQSize(NewQRoomVBoxLayoutWithTriggerSlotFromPointer(ptr).MinimumSizeDefault())
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) MinimumSizeDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQSizeFromPointer(C.QRoomVBoxLayoutWithTriggerSlot_MinimumSizeDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQRoomVBoxLayoutWithTriggerSlot_SizeHint
func callbackQRoomVBoxLayoutWithTriggerSlot_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return std_core.PointerFromQSize(signal.(func() *std_core.QSize)())
	}

	return std_core.PointerFromQSize(NewQRoomVBoxLayoutWithTriggerSlotFromPointer(ptr).SizeHintDefault())
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) SizeHintDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQSizeFromPointer(C.QRoomVBoxLayoutWithTriggerSlot_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQRoomVBoxLayoutWithTriggerSlot_ExpandingDirections
func callbackQRoomVBoxLayoutWithTriggerSlot_ExpandingDirections(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "expandingDirections"); signal != nil {
		return C.longlong(signal.(func() std_core.Qt__Orientation)())
	}

	return C.longlong(NewQRoomVBoxLayoutWithTriggerSlotFromPointer(ptr).ExpandingDirectionsDefault())
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) ExpandingDirectionsDefault() std_core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return std_core.Qt__Orientation(C.QRoomVBoxLayoutWithTriggerSlot_ExpandingDirectionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQRoomVBoxLayoutWithTriggerSlot_HasHeightForWidth
func callbackQRoomVBoxLayoutWithTriggerSlot_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQRoomVBoxLayoutWithTriggerSlotFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return C.QRoomVBoxLayoutWithTriggerSlot_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQRoomVBoxLayoutWithTriggerSlot_Count
func callbackQRoomVBoxLayoutWithTriggerSlot_Count(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "count"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(NewQRoomVBoxLayoutWithTriggerSlotFromPointer(ptr).CountDefault()))
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) CountDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QRoomVBoxLayoutWithTriggerSlot_CountDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackQRoomVBoxLayoutWithTriggerSlot_HeightForWidth
func callbackQRoomVBoxLayoutWithTriggerSlot_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQRoomVBoxLayoutWithTriggerSlotFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QRoomVBoxLayoutWithTriggerSlot_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQRoomVBoxLayoutWithTriggerSlot_MinimumHeightForWidth
func callbackQRoomVBoxLayoutWithTriggerSlot_MinimumHeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "minimumHeightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQRoomVBoxLayoutWithTriggerSlotFromPointer(ptr).MinimumHeightForWidthDefault(int(int32(w)))))
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) MinimumHeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QRoomVBoxLayoutWithTriggerSlot_MinimumHeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQRoomVBoxLayoutWithTriggerSlot_Layout
func callbackQRoomVBoxLayoutWithTriggerSlot_Layout(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "layout"); signal != nil {
		return std_widgets.PointerFromQLayout(signal.(func() *std_widgets.QLayout)())
	}

	return std_widgets.PointerFromQLayout(NewQRoomVBoxLayoutWithTriggerSlotFromPointer(ptr).LayoutDefault())
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) LayoutDefault() *std_widgets.QLayout {
	if ptr.Pointer() != nil {
		var tmpValue = std_widgets.NewQLayoutFromPointer(C.QRoomVBoxLayoutWithTriggerSlot_LayoutDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQRoomVBoxLayoutWithTriggerSlot_ChildEvent
func callbackQRoomVBoxLayoutWithTriggerSlot_ChildEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(e))
	} else {
		NewQRoomVBoxLayoutWithTriggerSlotFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(e))
	}
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) ChildEventDefault(e std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRoomVBoxLayoutWithTriggerSlot_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(e))
	}
}

//export callbackQRoomVBoxLayoutWithTriggerSlot_Geometry
func callbackQRoomVBoxLayoutWithTriggerSlot_Geometry(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "geometry"); signal != nil {
		return std_core.PointerFromQRect(signal.(func() *std_core.QRect)())
	}

	return std_core.PointerFromQRect(NewQRoomVBoxLayoutWithTriggerSlotFromPointer(ptr).GeometryDefault())
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) GeometryDefault() *std_core.QRect {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQRectFromPointer(C.QRoomVBoxLayoutWithTriggerSlot_GeometryDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

//export callbackQRoomVBoxLayoutWithTriggerSlot_ControlTypes
func callbackQRoomVBoxLayoutWithTriggerSlot_ControlTypes(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "controlTypes"); signal != nil {
		return C.longlong(signal.(func() std_widgets.QSizePolicy__ControlType)())
	}

	return C.longlong(NewQRoomVBoxLayoutWithTriggerSlotFromPointer(ptr).ControlTypesDefault())
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) ControlTypesDefault() std_widgets.QSizePolicy__ControlType {
	if ptr.Pointer() != nil {
		return std_widgets.QSizePolicy__ControlType(C.QRoomVBoxLayoutWithTriggerSlot_ControlTypesDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQRoomVBoxLayoutWithTriggerSlot_IsEmpty
func callbackQRoomVBoxLayoutWithTriggerSlot_IsEmpty(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isEmpty"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQRoomVBoxLayoutWithTriggerSlotFromPointer(ptr).IsEmptyDefault())))
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) IsEmptyDefault() bool {
	if ptr.Pointer() != nil {
		return C.QRoomVBoxLayoutWithTriggerSlot_IsEmptyDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQRoomVBoxLayoutWithTriggerSlot_IndexOf
func callbackQRoomVBoxLayoutWithTriggerSlot_IndexOf(ptr unsafe.Pointer, widget unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "indexOf"); signal != nil {
		return C.int(int32(signal.(func(*std_widgets.QWidget) int)(std_widgets.NewQWidgetFromPointer(widget))))
	}

	return C.int(int32(NewQRoomVBoxLayoutWithTriggerSlotFromPointer(ptr).IndexOfDefault(std_widgets.NewQWidgetFromPointer(widget))))
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) IndexOfDefault(widget std_widgets.QWidget_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QRoomVBoxLayoutWithTriggerSlot_IndexOfDefault(ptr.Pointer(), std_widgets.PointerFromQWidget(widget))))
	}
	return 0
}

//export callbackQRoomVBoxLayoutWithTriggerSlot_Event
func callbackQRoomVBoxLayoutWithTriggerSlot_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQRoomVBoxLayoutWithTriggerSlotFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QRoomVBoxLayoutWithTriggerSlot_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQRoomVBoxLayoutWithTriggerSlot_EventFilter
func callbackQRoomVBoxLayoutWithTriggerSlot_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQRoomVBoxLayoutWithTriggerSlotFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QRoomVBoxLayoutWithTriggerSlot_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQRoomVBoxLayoutWithTriggerSlot_ConnectNotify
func callbackQRoomVBoxLayoutWithTriggerSlot_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQRoomVBoxLayoutWithTriggerSlotFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QRoomVBoxLayoutWithTriggerSlot_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQRoomVBoxLayoutWithTriggerSlot_CustomEvent
func callbackQRoomVBoxLayoutWithTriggerSlot_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewQRoomVBoxLayoutWithTriggerSlotFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRoomVBoxLayoutWithTriggerSlot_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackQRoomVBoxLayoutWithTriggerSlot_DeleteLater
func callbackQRoomVBoxLayoutWithTriggerSlot_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQRoomVBoxLayoutWithTriggerSlotFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QRoomVBoxLayoutWithTriggerSlot_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQRoomVBoxLayoutWithTriggerSlot_Destroyed
func callbackQRoomVBoxLayoutWithTriggerSlot_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackQRoomVBoxLayoutWithTriggerSlot_DisconnectNotify
func callbackQRoomVBoxLayoutWithTriggerSlot_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQRoomVBoxLayoutWithTriggerSlotFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QRoomVBoxLayoutWithTriggerSlot_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQRoomVBoxLayoutWithTriggerSlot_ObjectNameChanged
func callbackQRoomVBoxLayoutWithTriggerSlot_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQRoomVBoxLayoutWithTriggerSlot_TimerEvent
func callbackQRoomVBoxLayoutWithTriggerSlot_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewQRoomVBoxLayoutWithTriggerSlotFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QRoomVBoxLayoutWithTriggerSlot_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

//export callbackQRoomVBoxLayoutWithTriggerSlot_SpacerItem
func callbackQRoomVBoxLayoutWithTriggerSlot_SpacerItem(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "spacerItem"); signal != nil {
		return std_widgets.PointerFromQSpacerItem(signal.(func() *std_widgets.QSpacerItem)())
	}

	return std_widgets.PointerFromQSpacerItem(NewQRoomVBoxLayoutWithTriggerSlotFromPointer(ptr).SpacerItemDefault())
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) SpacerItemDefault() *std_widgets.QSpacerItem {
	if ptr.Pointer() != nil {
		return std_widgets.NewQSpacerItemFromPointer(C.QRoomVBoxLayoutWithTriggerSlot_SpacerItemDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQRoomVBoxLayoutWithTriggerSlot_Widget
func callbackQRoomVBoxLayoutWithTriggerSlot_Widget(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "widget"); signal != nil {
		return std_widgets.PointerFromQWidget(signal.(func() *std_widgets.QWidget)())
	}

	return std_widgets.PointerFromQWidget(NewQRoomVBoxLayoutWithTriggerSlotFromPointer(ptr).WidgetDefault())
}

func (ptr *QRoomVBoxLayoutWithTriggerSlot) WidgetDefault() *std_widgets.QWidget {
	if ptr.Pointer() != nil {
		var tmpValue = std_widgets.NewQWidgetFromPointer(C.QRoomVBoxLayoutWithTriggerSlot_WidgetDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

type QVBoxLayoutWithTriggerSlot_ITF interface {
	std_widgets.QVBoxLayout_ITF
	QVBoxLayoutWithTriggerSlot_PTR() *QVBoxLayoutWithTriggerSlot
}

func (ptr *QVBoxLayoutWithTriggerSlot) QVBoxLayoutWithTriggerSlot_PTR() *QVBoxLayoutWithTriggerSlot {
	return ptr
}

func (ptr *QVBoxLayoutWithTriggerSlot) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QVBoxLayout_PTR().Pointer()
	}
	return nil
}

func (ptr *QVBoxLayoutWithTriggerSlot) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QVBoxLayout_PTR().SetPointer(p)
	}
}

func PointerFromQVBoxLayoutWithTriggerSlot(ptr QVBoxLayoutWithTriggerSlot_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVBoxLayoutWithTriggerSlot_PTR().Pointer()
	}
	return nil
}

func NewQVBoxLayoutWithTriggerSlotFromPointer(ptr unsafe.Pointer) *QVBoxLayoutWithTriggerSlot {
	var n *QVBoxLayoutWithTriggerSlot
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(QVBoxLayoutWithTriggerSlot)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *QVBoxLayoutWithTriggerSlot:
			n = deduced

		case *std_widgets.QVBoxLayout:
			n = &QVBoxLayoutWithTriggerSlot{QVBoxLayout: *deduced}

		default:
			n = new(QVBoxLayoutWithTriggerSlot)
			n.SetPointer(ptr)
		}
	}
	return n
}

//export callbackQVBoxLayoutWithTriggerSlot_Constructor
func callbackQVBoxLayoutWithTriggerSlot_Constructor(ptr unsafe.Pointer) {
	gPtr := NewQVBoxLayoutWithTriggerSlotFromPointer(ptr)
	qt.Register(ptr, gPtr)
}

//export callbackQVBoxLayoutWithTriggerSlot_TriggerMessage
func callbackQVBoxLayoutWithTriggerSlot_TriggerMessage(ptr unsafe.Pointer, messageBody C.struct_Moc_PackedString, sender C.struct_Moc_PackedString, timestamp C.longlong) {
	if signal := qt.GetSignal(ptr, "TriggerMessage"); signal != nil {
		signal.(func(string, string, int64))(cGoUnpackString(messageBody), cGoUnpackString(sender), int64(timestamp))
	}

}

func (ptr *QVBoxLayoutWithTriggerSlot) ConnectTriggerMessage(f func(messageBody string, sender string, timestamp int64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "TriggerMessage") {
			C.QVBoxLayoutWithTriggerSlot_ConnectTriggerMessage(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "TriggerMessage"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "TriggerMessage", func(messageBody string, sender string, timestamp int64) {
				signal.(func(string, string, int64))(messageBody, sender, timestamp)
				f(messageBody, sender, timestamp)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "TriggerMessage", f)
		}
	}
}

func (ptr *QVBoxLayoutWithTriggerSlot) DisconnectTriggerMessage() {
	if ptr.Pointer() != nil {
		C.QVBoxLayoutWithTriggerSlot_DisconnectTriggerMessage(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "TriggerMessage")
	}
}

func (ptr *QVBoxLayoutWithTriggerSlot) TriggerMessage(messageBody string, sender string, timestamp int64) {
	if ptr.Pointer() != nil {
		var messageBodyC *C.char
		if messageBody != "" {
			messageBodyC = C.CString(messageBody)
			defer C.free(unsafe.Pointer(messageBodyC))
		}
		var senderC *C.char
		if sender != "" {
			senderC = C.CString(sender)
			defer C.free(unsafe.Pointer(senderC))
		}
		C.QVBoxLayoutWithTriggerSlot_TriggerMessage(ptr.Pointer(), C.struct_Moc_PackedString{data: messageBodyC, len: C.longlong(len(messageBody))}, C.struct_Moc_PackedString{data: senderC, len: C.longlong(len(sender))}, C.longlong(timestamp))
	}
}

func QVBoxLayoutWithTriggerSlot_QRegisterMetaType() int {
	return int(int32(C.QVBoxLayoutWithTriggerSlot_QVBoxLayoutWithTriggerSlot_QRegisterMetaType()))
}

func (ptr *QVBoxLayoutWithTriggerSlot) QRegisterMetaType() int {
	return int(int32(C.QVBoxLayoutWithTriggerSlot_QVBoxLayoutWithTriggerSlot_QRegisterMetaType()))
}

func QVBoxLayoutWithTriggerSlot_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.QVBoxLayoutWithTriggerSlot_QVBoxLayoutWithTriggerSlot_QRegisterMetaType2(typeNameC)))
}

func (ptr *QVBoxLayoutWithTriggerSlot) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.QVBoxLayoutWithTriggerSlot_QVBoxLayoutWithTriggerSlot_QRegisterMetaType2(typeNameC)))
}

func QVBoxLayoutWithTriggerSlot_QmlRegisterType() int {
	return int(int32(C.QVBoxLayoutWithTriggerSlot_QVBoxLayoutWithTriggerSlot_QmlRegisterType()))
}

func (ptr *QVBoxLayoutWithTriggerSlot) QmlRegisterType() int {
	return int(int32(C.QVBoxLayoutWithTriggerSlot_QVBoxLayoutWithTriggerSlot_QmlRegisterType()))
}

func QVBoxLayoutWithTriggerSlot_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.QVBoxLayoutWithTriggerSlot_QVBoxLayoutWithTriggerSlot_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QVBoxLayoutWithTriggerSlot) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.QVBoxLayoutWithTriggerSlot_QVBoxLayoutWithTriggerSlot_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QVBoxLayoutWithTriggerSlot) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQByteArrayFromPointer(C.QVBoxLayoutWithTriggerSlot___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QVBoxLayoutWithTriggerSlot) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QVBoxLayoutWithTriggerSlot___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *QVBoxLayoutWithTriggerSlot) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QVBoxLayoutWithTriggerSlot___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QVBoxLayoutWithTriggerSlot) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.QVBoxLayoutWithTriggerSlot___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QVBoxLayoutWithTriggerSlot) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QVBoxLayoutWithTriggerSlot___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QVBoxLayoutWithTriggerSlot) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QVBoxLayoutWithTriggerSlot___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QVBoxLayoutWithTriggerSlot) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.QVBoxLayoutWithTriggerSlot___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QVBoxLayoutWithTriggerSlot) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QVBoxLayoutWithTriggerSlot___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QVBoxLayoutWithTriggerSlot) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QVBoxLayoutWithTriggerSlot___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QVBoxLayoutWithTriggerSlot) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.QVBoxLayoutWithTriggerSlot___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QVBoxLayoutWithTriggerSlot) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QVBoxLayoutWithTriggerSlot___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QVBoxLayoutWithTriggerSlot) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QVBoxLayoutWithTriggerSlot___findChildren_newList(ptr.Pointer()))
}

func (ptr *QVBoxLayoutWithTriggerSlot) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.QVBoxLayoutWithTriggerSlot___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QVBoxLayoutWithTriggerSlot) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QVBoxLayoutWithTriggerSlot___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QVBoxLayoutWithTriggerSlot) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QVBoxLayoutWithTriggerSlot___children_newList(ptr.Pointer()))
}

func NewQVBoxLayoutWithTriggerSlot() *QVBoxLayoutWithTriggerSlot {
	var tmpValue = NewQVBoxLayoutWithTriggerSlotFromPointer(C.QVBoxLayoutWithTriggerSlot_NewQVBoxLayoutWithTriggerSlot())
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQVBoxLayoutWithTriggerSlot2(parent std_widgets.QWidget_ITF) *QVBoxLayoutWithTriggerSlot {
	var tmpValue = NewQVBoxLayoutWithTriggerSlotFromPointer(C.QVBoxLayoutWithTriggerSlot_NewQVBoxLayoutWithTriggerSlot2(std_widgets.PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QVBoxLayoutWithTriggerSlot) DestroyQVBoxLayoutWithTriggerSlot() {
	if ptr.Pointer() != nil {
		C.QVBoxLayoutWithTriggerSlot_DestroyQVBoxLayoutWithTriggerSlot(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQVBoxLayoutWithTriggerSlot_TakeAt
func callbackQVBoxLayoutWithTriggerSlot_TakeAt(ptr unsafe.Pointer, index C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "takeAt"); signal != nil {
		return std_widgets.PointerFromQLayoutItem(signal.(func(int) *std_widgets.QLayoutItem)(int(int32(index))))
	}

	return std_widgets.PointerFromQLayoutItem(NewQVBoxLayoutWithTriggerSlotFromPointer(ptr).TakeAtDefault(int(int32(index))))
}

func (ptr *QVBoxLayoutWithTriggerSlot) TakeAtDefault(index int) *std_widgets.QLayoutItem {
	if ptr.Pointer() != nil {
		return std_widgets.NewQLayoutItemFromPointer(C.QVBoxLayoutWithTriggerSlot_TakeAtDefault(ptr.Pointer(), C.int(int32(index))))
	}
	return nil
}

//export callbackQVBoxLayoutWithTriggerSlot_AddItem
func callbackQVBoxLayoutWithTriggerSlot_AddItem(ptr unsafe.Pointer, item unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "addItem"); signal != nil {
		signal.(func(*std_widgets.QLayoutItem))(std_widgets.NewQLayoutItemFromPointer(item))
	} else {
		NewQVBoxLayoutWithTriggerSlotFromPointer(ptr).AddItemDefault(std_widgets.NewQLayoutItemFromPointer(item))
	}
}

func (ptr *QVBoxLayoutWithTriggerSlot) AddItemDefault(item std_widgets.QLayoutItem_ITF) {
	if ptr.Pointer() != nil {
		C.QVBoxLayoutWithTriggerSlot_AddItemDefault(ptr.Pointer(), std_widgets.PointerFromQLayoutItem(item))
	}
}

//export callbackQVBoxLayoutWithTriggerSlot_Invalidate
func callbackQVBoxLayoutWithTriggerSlot_Invalidate(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "invalidate"); signal != nil {
		signal.(func())()
	} else {
		NewQVBoxLayoutWithTriggerSlotFromPointer(ptr).InvalidateDefault()
	}
}

func (ptr *QVBoxLayoutWithTriggerSlot) InvalidateDefault() {
	if ptr.Pointer() != nil {
		C.QVBoxLayoutWithTriggerSlot_InvalidateDefault(ptr.Pointer())
	}
}

//export callbackQVBoxLayoutWithTriggerSlot_SetGeometry
func callbackQVBoxLayoutWithTriggerSlot_SetGeometry(ptr unsafe.Pointer, r unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setGeometry"); signal != nil {
		signal.(func(*std_core.QRect))(std_core.NewQRectFromPointer(r))
	} else {
		NewQVBoxLayoutWithTriggerSlotFromPointer(ptr).SetGeometryDefault(std_core.NewQRectFromPointer(r))
	}
}

func (ptr *QVBoxLayoutWithTriggerSlot) SetGeometryDefault(r std_core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QVBoxLayoutWithTriggerSlot_SetGeometryDefault(ptr.Pointer(), std_core.PointerFromQRect(r))
	}
}

//export callbackQVBoxLayoutWithTriggerSlot_ItemAt
func callbackQVBoxLayoutWithTriggerSlot_ItemAt(ptr unsafe.Pointer, index C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemAt"); signal != nil {
		return std_widgets.PointerFromQLayoutItem(signal.(func(int) *std_widgets.QLayoutItem)(int(int32(index))))
	}

	return std_widgets.PointerFromQLayoutItem(NewQVBoxLayoutWithTriggerSlotFromPointer(ptr).ItemAtDefault(int(int32(index))))
}

func (ptr *QVBoxLayoutWithTriggerSlot) ItemAtDefault(index int) *std_widgets.QLayoutItem {
	if ptr.Pointer() != nil {
		return std_widgets.NewQLayoutItemFromPointer(C.QVBoxLayoutWithTriggerSlot_ItemAtDefault(ptr.Pointer(), C.int(int32(index))))
	}
	return nil
}

//export callbackQVBoxLayoutWithTriggerSlot_MaximumSize
func callbackQVBoxLayoutWithTriggerSlot_MaximumSize(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "maximumSize"); signal != nil {
		return std_core.PointerFromQSize(signal.(func() *std_core.QSize)())
	}

	return std_core.PointerFromQSize(NewQVBoxLayoutWithTriggerSlotFromPointer(ptr).MaximumSizeDefault())
}

func (ptr *QVBoxLayoutWithTriggerSlot) MaximumSizeDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQSizeFromPointer(C.QVBoxLayoutWithTriggerSlot_MaximumSizeDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQVBoxLayoutWithTriggerSlot_MinimumSize
func callbackQVBoxLayoutWithTriggerSlot_MinimumSize(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSize"); signal != nil {
		return std_core.PointerFromQSize(signal.(func() *std_core.QSize)())
	}

	return std_core.PointerFromQSize(NewQVBoxLayoutWithTriggerSlotFromPointer(ptr).MinimumSizeDefault())
}

func (ptr *QVBoxLayoutWithTriggerSlot) MinimumSizeDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQSizeFromPointer(C.QVBoxLayoutWithTriggerSlot_MinimumSizeDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQVBoxLayoutWithTriggerSlot_SizeHint
func callbackQVBoxLayoutWithTriggerSlot_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return std_core.PointerFromQSize(signal.(func() *std_core.QSize)())
	}

	return std_core.PointerFromQSize(NewQVBoxLayoutWithTriggerSlotFromPointer(ptr).SizeHintDefault())
}

func (ptr *QVBoxLayoutWithTriggerSlot) SizeHintDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQSizeFromPointer(C.QVBoxLayoutWithTriggerSlot_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQVBoxLayoutWithTriggerSlot_ExpandingDirections
func callbackQVBoxLayoutWithTriggerSlot_ExpandingDirections(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "expandingDirections"); signal != nil {
		return C.longlong(signal.(func() std_core.Qt__Orientation)())
	}

	return C.longlong(NewQVBoxLayoutWithTriggerSlotFromPointer(ptr).ExpandingDirectionsDefault())
}

func (ptr *QVBoxLayoutWithTriggerSlot) ExpandingDirectionsDefault() std_core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return std_core.Qt__Orientation(C.QVBoxLayoutWithTriggerSlot_ExpandingDirectionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQVBoxLayoutWithTriggerSlot_HasHeightForWidth
func callbackQVBoxLayoutWithTriggerSlot_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQVBoxLayoutWithTriggerSlotFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QVBoxLayoutWithTriggerSlot) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return C.QVBoxLayoutWithTriggerSlot_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQVBoxLayoutWithTriggerSlot_Count
func callbackQVBoxLayoutWithTriggerSlot_Count(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "count"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(NewQVBoxLayoutWithTriggerSlotFromPointer(ptr).CountDefault()))
}

func (ptr *QVBoxLayoutWithTriggerSlot) CountDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVBoxLayoutWithTriggerSlot_CountDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackQVBoxLayoutWithTriggerSlot_HeightForWidth
func callbackQVBoxLayoutWithTriggerSlot_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQVBoxLayoutWithTriggerSlotFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QVBoxLayoutWithTriggerSlot) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVBoxLayoutWithTriggerSlot_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQVBoxLayoutWithTriggerSlot_MinimumHeightForWidth
func callbackQVBoxLayoutWithTriggerSlot_MinimumHeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "minimumHeightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQVBoxLayoutWithTriggerSlotFromPointer(ptr).MinimumHeightForWidthDefault(int(int32(w)))))
}

func (ptr *QVBoxLayoutWithTriggerSlot) MinimumHeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVBoxLayoutWithTriggerSlot_MinimumHeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQVBoxLayoutWithTriggerSlot_Layout
func callbackQVBoxLayoutWithTriggerSlot_Layout(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "layout"); signal != nil {
		return std_widgets.PointerFromQLayout(signal.(func() *std_widgets.QLayout)())
	}

	return std_widgets.PointerFromQLayout(NewQVBoxLayoutWithTriggerSlotFromPointer(ptr).LayoutDefault())
}

func (ptr *QVBoxLayoutWithTriggerSlot) LayoutDefault() *std_widgets.QLayout {
	if ptr.Pointer() != nil {
		var tmpValue = std_widgets.NewQLayoutFromPointer(C.QVBoxLayoutWithTriggerSlot_LayoutDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQVBoxLayoutWithTriggerSlot_ChildEvent
func callbackQVBoxLayoutWithTriggerSlot_ChildEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(e))
	} else {
		NewQVBoxLayoutWithTriggerSlotFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(e))
	}
}

func (ptr *QVBoxLayoutWithTriggerSlot) ChildEventDefault(e std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QVBoxLayoutWithTriggerSlot_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(e))
	}
}

//export callbackQVBoxLayoutWithTriggerSlot_Geometry
func callbackQVBoxLayoutWithTriggerSlot_Geometry(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "geometry"); signal != nil {
		return std_core.PointerFromQRect(signal.(func() *std_core.QRect)())
	}

	return std_core.PointerFromQRect(NewQVBoxLayoutWithTriggerSlotFromPointer(ptr).GeometryDefault())
}

func (ptr *QVBoxLayoutWithTriggerSlot) GeometryDefault() *std_core.QRect {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQRectFromPointer(C.QVBoxLayoutWithTriggerSlot_GeometryDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

//export callbackQVBoxLayoutWithTriggerSlot_ControlTypes
func callbackQVBoxLayoutWithTriggerSlot_ControlTypes(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "controlTypes"); signal != nil {
		return C.longlong(signal.(func() std_widgets.QSizePolicy__ControlType)())
	}

	return C.longlong(NewQVBoxLayoutWithTriggerSlotFromPointer(ptr).ControlTypesDefault())
}

func (ptr *QVBoxLayoutWithTriggerSlot) ControlTypesDefault() std_widgets.QSizePolicy__ControlType {
	if ptr.Pointer() != nil {
		return std_widgets.QSizePolicy__ControlType(C.QVBoxLayoutWithTriggerSlot_ControlTypesDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQVBoxLayoutWithTriggerSlot_IsEmpty
func callbackQVBoxLayoutWithTriggerSlot_IsEmpty(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isEmpty"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQVBoxLayoutWithTriggerSlotFromPointer(ptr).IsEmptyDefault())))
}

func (ptr *QVBoxLayoutWithTriggerSlot) IsEmptyDefault() bool {
	if ptr.Pointer() != nil {
		return C.QVBoxLayoutWithTriggerSlot_IsEmptyDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQVBoxLayoutWithTriggerSlot_IndexOf
func callbackQVBoxLayoutWithTriggerSlot_IndexOf(ptr unsafe.Pointer, widget unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "indexOf"); signal != nil {
		return C.int(int32(signal.(func(*std_widgets.QWidget) int)(std_widgets.NewQWidgetFromPointer(widget))))
	}

	return C.int(int32(NewQVBoxLayoutWithTriggerSlotFromPointer(ptr).IndexOfDefault(std_widgets.NewQWidgetFromPointer(widget))))
}

func (ptr *QVBoxLayoutWithTriggerSlot) IndexOfDefault(widget std_widgets.QWidget_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QVBoxLayoutWithTriggerSlot_IndexOfDefault(ptr.Pointer(), std_widgets.PointerFromQWidget(widget))))
	}
	return 0
}

//export callbackQVBoxLayoutWithTriggerSlot_Event
func callbackQVBoxLayoutWithTriggerSlot_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQVBoxLayoutWithTriggerSlotFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *QVBoxLayoutWithTriggerSlot) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QVBoxLayoutWithTriggerSlot_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQVBoxLayoutWithTriggerSlot_EventFilter
func callbackQVBoxLayoutWithTriggerSlot_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQVBoxLayoutWithTriggerSlotFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *QVBoxLayoutWithTriggerSlot) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QVBoxLayoutWithTriggerSlot_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQVBoxLayoutWithTriggerSlot_ConnectNotify
func callbackQVBoxLayoutWithTriggerSlot_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQVBoxLayoutWithTriggerSlotFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QVBoxLayoutWithTriggerSlot) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QVBoxLayoutWithTriggerSlot_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQVBoxLayoutWithTriggerSlot_CustomEvent
func callbackQVBoxLayoutWithTriggerSlot_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewQVBoxLayoutWithTriggerSlotFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *QVBoxLayoutWithTriggerSlot) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QVBoxLayoutWithTriggerSlot_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackQVBoxLayoutWithTriggerSlot_DeleteLater
func callbackQVBoxLayoutWithTriggerSlot_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQVBoxLayoutWithTriggerSlotFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QVBoxLayoutWithTriggerSlot) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QVBoxLayoutWithTriggerSlot_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQVBoxLayoutWithTriggerSlot_Destroyed
func callbackQVBoxLayoutWithTriggerSlot_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackQVBoxLayoutWithTriggerSlot_DisconnectNotify
func callbackQVBoxLayoutWithTriggerSlot_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQVBoxLayoutWithTriggerSlotFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QVBoxLayoutWithTriggerSlot) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QVBoxLayoutWithTriggerSlot_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQVBoxLayoutWithTriggerSlot_ObjectNameChanged
func callbackQVBoxLayoutWithTriggerSlot_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQVBoxLayoutWithTriggerSlot_TimerEvent
func callbackQVBoxLayoutWithTriggerSlot_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewQVBoxLayoutWithTriggerSlotFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QVBoxLayoutWithTriggerSlot) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QVBoxLayoutWithTriggerSlot_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

//export callbackQVBoxLayoutWithTriggerSlot_SpacerItem
func callbackQVBoxLayoutWithTriggerSlot_SpacerItem(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "spacerItem"); signal != nil {
		return std_widgets.PointerFromQSpacerItem(signal.(func() *std_widgets.QSpacerItem)())
	}

	return std_widgets.PointerFromQSpacerItem(NewQVBoxLayoutWithTriggerSlotFromPointer(ptr).SpacerItemDefault())
}

func (ptr *QVBoxLayoutWithTriggerSlot) SpacerItemDefault() *std_widgets.QSpacerItem {
	if ptr.Pointer() != nil {
		return std_widgets.NewQSpacerItemFromPointer(C.QVBoxLayoutWithTriggerSlot_SpacerItemDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQVBoxLayoutWithTriggerSlot_Widget
func callbackQVBoxLayoutWithTriggerSlot_Widget(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "widget"); signal != nil {
		return std_widgets.PointerFromQWidget(signal.(func() *std_widgets.QWidget)())
	}

	return std_widgets.PointerFromQWidget(NewQVBoxLayoutWithTriggerSlotFromPointer(ptr).WidgetDefault())
}

func (ptr *QVBoxLayoutWithTriggerSlot) WidgetDefault() *std_widgets.QWidget {
	if ptr.Pointer() != nil {
		var tmpValue = std_widgets.NewQWidgetFromPointer(C.QVBoxLayoutWithTriggerSlot_WidgetDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}
