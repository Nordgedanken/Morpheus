package elements

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

type QGridLayoutWithTriggerSlot_ITF interface {
	std_widgets.QGridLayout_ITF
	QGridLayoutWithTriggerSlot_PTR() *QGridLayoutWithTriggerSlot
}

func (ptr *QGridLayoutWithTriggerSlot) QGridLayoutWithTriggerSlot_PTR() *QGridLayoutWithTriggerSlot {
	return ptr
}

func (ptr *QGridLayoutWithTriggerSlot) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QGridLayout_PTR().Pointer()
	}
	return nil
}

func (ptr *QGridLayoutWithTriggerSlot) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QGridLayout_PTR().SetPointer(p)
	}
}

func PointerFromQGridLayoutWithTriggerSlot(ptr QGridLayoutWithTriggerSlot_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGridLayoutWithTriggerSlot_PTR().Pointer()
	}
	return nil
}

func NewQGridLayoutWithTriggerSlotFromPointer(ptr unsafe.Pointer) *QGridLayoutWithTriggerSlot {
	var n *QGridLayoutWithTriggerSlot
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(QGridLayoutWithTriggerSlot)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *QGridLayoutWithTriggerSlot:
			n = deduced

		case *std_widgets.QGridLayout:
			n = &QGridLayoutWithTriggerSlot{QGridLayout: *deduced}

		default:
			n = new(QGridLayoutWithTriggerSlot)
			n.SetPointer(ptr)
		}
	}
	return n
}

//export callbackQGridLayoutWithTriggerSlot_Constructor
func callbackQGridLayoutWithTriggerSlot_Constructor(ptr unsafe.Pointer) {
	gPtr := NewQGridLayoutWithTriggerSlotFromPointer(ptr)
	qt.Register(ptr, gPtr)
}

//export callbackQGridLayoutWithTriggerSlot_TriggerMessage
func callbackQGridLayoutWithTriggerSlot_TriggerMessage(ptr unsafe.Pointer, messageBody C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "TriggerMessage"); signal != nil {
		signal.(func(string))(cGoUnpackString(messageBody))
	}

}

func (ptr *QGridLayoutWithTriggerSlot) ConnectTriggerMessage(f func(messageBody string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "TriggerMessage"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "TriggerMessage", func(messageBody string) {
				signal.(func(string))(messageBody)
				f(messageBody)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "TriggerMessage", f)
		}
	}
}

func (ptr *QGridLayoutWithTriggerSlot) DisconnectTriggerMessage() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "TriggerMessage")
	}
}

func (ptr *QGridLayoutWithTriggerSlot) TriggerMessage(messageBody string) {
	if ptr.Pointer() != nil {
		var messageBodyC *C.char
		if messageBody != "" {
			messageBodyC = C.CString(messageBody)
			defer C.free(unsafe.Pointer(messageBodyC))
		}
		C.QGridLayoutWithTriggerSlot_TriggerMessage(ptr.Pointer(), C.struct_Moc_PackedString{data: messageBodyC, len: C.longlong(len(messageBody))})
	}
}

func QGridLayoutWithTriggerSlot_QRegisterMetaType() int {
	return int(int32(C.QGridLayoutWithTriggerSlot_QGridLayoutWithTriggerSlot_QRegisterMetaType()))
}

func (ptr *QGridLayoutWithTriggerSlot) QRegisterMetaType() int {
	return int(int32(C.QGridLayoutWithTriggerSlot_QGridLayoutWithTriggerSlot_QRegisterMetaType()))
}

func QGridLayoutWithTriggerSlot_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.QGridLayoutWithTriggerSlot_QGridLayoutWithTriggerSlot_QRegisterMetaType2(typeNameC)))
}

func (ptr *QGridLayoutWithTriggerSlot) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.QGridLayoutWithTriggerSlot_QGridLayoutWithTriggerSlot_QRegisterMetaType2(typeNameC)))
}

func QGridLayoutWithTriggerSlot_QmlRegisterType() int {
	return int(int32(C.QGridLayoutWithTriggerSlot_QGridLayoutWithTriggerSlot_QmlRegisterType()))
}

func (ptr *QGridLayoutWithTriggerSlot) QmlRegisterType() int {
	return int(int32(C.QGridLayoutWithTriggerSlot_QGridLayoutWithTriggerSlot_QmlRegisterType()))
}

func QGridLayoutWithTriggerSlot_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.QGridLayoutWithTriggerSlot_QGridLayoutWithTriggerSlot_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QGridLayoutWithTriggerSlot) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.QGridLayoutWithTriggerSlot_QGridLayoutWithTriggerSlot_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *QGridLayoutWithTriggerSlot) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQByteArrayFromPointer(C.QGridLayoutWithTriggerSlot___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QGridLayoutWithTriggerSlot) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QGridLayoutWithTriggerSlot___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *QGridLayoutWithTriggerSlot) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QGridLayoutWithTriggerSlot___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QGridLayoutWithTriggerSlot) __findChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.QGridLayoutWithTriggerSlot___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGridLayoutWithTriggerSlot) __findChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGridLayoutWithTriggerSlot___findChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QGridLayoutWithTriggerSlot) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QGridLayoutWithTriggerSlot___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QGridLayoutWithTriggerSlot) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.QGridLayoutWithTriggerSlot___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGridLayoutWithTriggerSlot) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGridLayoutWithTriggerSlot___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QGridLayoutWithTriggerSlot) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QGridLayoutWithTriggerSlot___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QGridLayoutWithTriggerSlot) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.QGridLayoutWithTriggerSlot___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGridLayoutWithTriggerSlot) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGridLayoutWithTriggerSlot___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QGridLayoutWithTriggerSlot) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QGridLayoutWithTriggerSlot___findChildren_newList(ptr.Pointer()))
}

func (ptr *QGridLayoutWithTriggerSlot) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQObjectFromPointer(C.QGridLayoutWithTriggerSlot___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QGridLayoutWithTriggerSlot) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QGridLayoutWithTriggerSlot___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *QGridLayoutWithTriggerSlot) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QGridLayoutWithTriggerSlot___children_newList(ptr.Pointer()))
}

func NewQGridLayoutWithTriggerSlot2() *QGridLayoutWithTriggerSlot {
	var tmpValue = NewQGridLayoutWithTriggerSlotFromPointer(C.QGridLayoutWithTriggerSlot_NewQGridLayoutWithTriggerSlot2())
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQGridLayoutWithTriggerSlot(parent std_widgets.QWidget_ITF) *QGridLayoutWithTriggerSlot {
	var tmpValue = NewQGridLayoutWithTriggerSlotFromPointer(C.QGridLayoutWithTriggerSlot_NewQGridLayoutWithTriggerSlot(std_widgets.PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QGridLayoutWithTriggerSlot) DestroyQGridLayoutWithTriggerSlot() {
	if ptr.Pointer() != nil {
		C.QGridLayoutWithTriggerSlot_DestroyQGridLayoutWithTriggerSlot(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQGridLayoutWithTriggerSlot_TakeAt
func callbackQGridLayoutWithTriggerSlot_TakeAt(ptr unsafe.Pointer, index C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "takeAt"); signal != nil {
		return std_widgets.PointerFromQLayoutItem(signal.(func(int) *std_widgets.QLayoutItem)(int(int32(index))))
	}

	return std_widgets.PointerFromQLayoutItem(NewQGridLayoutWithTriggerSlotFromPointer(ptr).TakeAtDefault(int(int32(index))))
}

func (ptr *QGridLayoutWithTriggerSlot) TakeAtDefault(index int) *std_widgets.QLayoutItem {
	if ptr.Pointer() != nil {
		return std_widgets.NewQLayoutItemFromPointer(C.QGridLayoutWithTriggerSlot_TakeAtDefault(ptr.Pointer(), C.int(int32(index))))
	}
	return nil
}

//export callbackQGridLayoutWithTriggerSlot_Invalidate
func callbackQGridLayoutWithTriggerSlot_Invalidate(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "invalidate"); signal != nil {
		signal.(func())()
	} else {
		NewQGridLayoutWithTriggerSlotFromPointer(ptr).InvalidateDefault()
	}
}

func (ptr *QGridLayoutWithTriggerSlot) InvalidateDefault() {
	if ptr.Pointer() != nil {
		C.QGridLayoutWithTriggerSlot_InvalidateDefault(ptr.Pointer())
	}
}

//export callbackQGridLayoutWithTriggerSlot_SetGeometry
func callbackQGridLayoutWithTriggerSlot_SetGeometry(ptr unsafe.Pointer, rect unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setGeometry"); signal != nil {
		signal.(func(*std_core.QRect))(std_core.NewQRectFromPointer(rect))
	} else {
		NewQGridLayoutWithTriggerSlotFromPointer(ptr).SetGeometryDefault(std_core.NewQRectFromPointer(rect))
	}
}

func (ptr *QGridLayoutWithTriggerSlot) SetGeometryDefault(rect std_core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QGridLayoutWithTriggerSlot_SetGeometryDefault(ptr.Pointer(), std_core.PointerFromQRect(rect))
	}
}

//export callbackQGridLayoutWithTriggerSlot_ItemAt
func callbackQGridLayoutWithTriggerSlot_ItemAt(ptr unsafe.Pointer, index C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "itemAt"); signal != nil {
		return std_widgets.PointerFromQLayoutItem(signal.(func(int) *std_widgets.QLayoutItem)(int(int32(index))))
	}

	return std_widgets.PointerFromQLayoutItem(NewQGridLayoutWithTriggerSlotFromPointer(ptr).ItemAtDefault(int(int32(index))))
}

func (ptr *QGridLayoutWithTriggerSlot) ItemAtDefault(index int) *std_widgets.QLayoutItem {
	if ptr.Pointer() != nil {
		return std_widgets.NewQLayoutItemFromPointer(C.QGridLayoutWithTriggerSlot_ItemAtDefault(ptr.Pointer(), C.int(int32(index))))
	}
	return nil
}

//export callbackQGridLayoutWithTriggerSlot_MaximumSize
func callbackQGridLayoutWithTriggerSlot_MaximumSize(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "maximumSize"); signal != nil {
		return std_core.PointerFromQSize(signal.(func() *std_core.QSize)())
	}

	return std_core.PointerFromQSize(NewQGridLayoutWithTriggerSlotFromPointer(ptr).MaximumSizeDefault())
}

func (ptr *QGridLayoutWithTriggerSlot) MaximumSizeDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQSizeFromPointer(C.QGridLayoutWithTriggerSlot_MaximumSizeDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQGridLayoutWithTriggerSlot_MinimumSize
func callbackQGridLayoutWithTriggerSlot_MinimumSize(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "minimumSize"); signal != nil {
		return std_core.PointerFromQSize(signal.(func() *std_core.QSize)())
	}

	return std_core.PointerFromQSize(NewQGridLayoutWithTriggerSlotFromPointer(ptr).MinimumSizeDefault())
}

func (ptr *QGridLayoutWithTriggerSlot) MinimumSizeDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQSizeFromPointer(C.QGridLayoutWithTriggerSlot_MinimumSizeDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQGridLayoutWithTriggerSlot_SizeHint
func callbackQGridLayoutWithTriggerSlot_SizeHint(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "sizeHint"); signal != nil {
		return std_core.PointerFromQSize(signal.(func() *std_core.QSize)())
	}

	return std_core.PointerFromQSize(NewQGridLayoutWithTriggerSlotFromPointer(ptr).SizeHintDefault())
}

func (ptr *QGridLayoutWithTriggerSlot) SizeHintDefault() *std_core.QSize {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQSizeFromPointer(C.QGridLayoutWithTriggerSlot_SizeHintDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QSize).DestroyQSize)
		return tmpValue
	}
	return nil
}

//export callbackQGridLayoutWithTriggerSlot_ExpandingDirections
func callbackQGridLayoutWithTriggerSlot_ExpandingDirections(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "expandingDirections"); signal != nil {
		return C.longlong(signal.(func() std_core.Qt__Orientation)())
	}

	return C.longlong(NewQGridLayoutWithTriggerSlotFromPointer(ptr).ExpandingDirectionsDefault())
}

func (ptr *QGridLayoutWithTriggerSlot) ExpandingDirectionsDefault() std_core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return std_core.Qt__Orientation(C.QGridLayoutWithTriggerSlot_ExpandingDirectionsDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQGridLayoutWithTriggerSlot_HasHeightForWidth
func callbackQGridLayoutWithTriggerSlot_HasHeightForWidth(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "hasHeightForWidth"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGridLayoutWithTriggerSlotFromPointer(ptr).HasHeightForWidthDefault())))
}

func (ptr *QGridLayoutWithTriggerSlot) HasHeightForWidthDefault() bool {
	if ptr.Pointer() != nil {
		return C.QGridLayoutWithTriggerSlot_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQGridLayoutWithTriggerSlot_Count
func callbackQGridLayoutWithTriggerSlot_Count(ptr unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "count"); signal != nil {
		return C.int(int32(signal.(func() int)()))
	}

	return C.int(int32(NewQGridLayoutWithTriggerSlotFromPointer(ptr).CountDefault()))
}

func (ptr *QGridLayoutWithTriggerSlot) CountDefault() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGridLayoutWithTriggerSlot_CountDefault(ptr.Pointer())))
	}
	return 0
}

//export callbackQGridLayoutWithTriggerSlot_HeightForWidth
func callbackQGridLayoutWithTriggerSlot_HeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "heightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQGridLayoutWithTriggerSlotFromPointer(ptr).HeightForWidthDefault(int(int32(w)))))
}

func (ptr *QGridLayoutWithTriggerSlot) HeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGridLayoutWithTriggerSlot_HeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQGridLayoutWithTriggerSlot_MinimumHeightForWidth
func callbackQGridLayoutWithTriggerSlot_MinimumHeightForWidth(ptr unsafe.Pointer, w C.int) C.int {
	if signal := qt.GetSignal(ptr, "minimumHeightForWidth"); signal != nil {
		return C.int(int32(signal.(func(int) int)(int(int32(w)))))
	}

	return C.int(int32(NewQGridLayoutWithTriggerSlotFromPointer(ptr).MinimumHeightForWidthDefault(int(int32(w)))))
}

func (ptr *QGridLayoutWithTriggerSlot) MinimumHeightForWidthDefault(w int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGridLayoutWithTriggerSlot_MinimumHeightForWidthDefault(ptr.Pointer(), C.int(int32(w)))))
	}
	return 0
}

//export callbackQGridLayoutWithTriggerSlot_Layout
func callbackQGridLayoutWithTriggerSlot_Layout(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "layout"); signal != nil {
		return std_widgets.PointerFromQLayout(signal.(func() *std_widgets.QLayout)())
	}

	return std_widgets.PointerFromQLayout(NewQGridLayoutWithTriggerSlotFromPointer(ptr).LayoutDefault())
}

func (ptr *QGridLayoutWithTriggerSlot) LayoutDefault() *std_widgets.QLayout {
	if ptr.Pointer() != nil {
		var tmpValue = std_widgets.NewQLayoutFromPointer(C.QGridLayoutWithTriggerSlot_LayoutDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQGridLayoutWithTriggerSlot_AddItem
func callbackQGridLayoutWithTriggerSlot_AddItem(ptr unsafe.Pointer, item unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "addItem"); signal != nil {
		signal.(func(*std_widgets.QLayoutItem))(std_widgets.NewQLayoutItemFromPointer(item))
	} else {
		NewQGridLayoutWithTriggerSlotFromPointer(ptr).AddItemDefault(std_widgets.NewQLayoutItemFromPointer(item))
	}
}

func (ptr *QGridLayoutWithTriggerSlot) AddItemDefault(item std_widgets.QLayoutItem_ITF) {
	if ptr.Pointer() != nil {
		C.QGridLayoutWithTriggerSlot_AddItemDefault(ptr.Pointer(), std_widgets.PointerFromQLayoutItem(item))
	}
}

//export callbackQGridLayoutWithTriggerSlot_ChildEvent
func callbackQGridLayoutWithTriggerSlot_ChildEvent(ptr unsafe.Pointer, e unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*std_core.QChildEvent))(std_core.NewQChildEventFromPointer(e))
	} else {
		NewQGridLayoutWithTriggerSlotFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(e))
	}
}

func (ptr *QGridLayoutWithTriggerSlot) ChildEventDefault(e std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGridLayoutWithTriggerSlot_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(e))
	}
}

//export callbackQGridLayoutWithTriggerSlot_Geometry
func callbackQGridLayoutWithTriggerSlot_Geometry(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "geometry"); signal != nil {
		return std_core.PointerFromQRect(signal.(func() *std_core.QRect)())
	}

	return std_core.PointerFromQRect(NewQGridLayoutWithTriggerSlotFromPointer(ptr).GeometryDefault())
}

func (ptr *QGridLayoutWithTriggerSlot) GeometryDefault() *std_core.QRect {
	if ptr.Pointer() != nil {
		var tmpValue = std_core.NewQRectFromPointer(C.QGridLayoutWithTriggerSlot_GeometryDefault(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*std_core.QRect).DestroyQRect)
		return tmpValue
	}
	return nil
}

//export callbackQGridLayoutWithTriggerSlot_ControlTypes
func callbackQGridLayoutWithTriggerSlot_ControlTypes(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "controlTypes"); signal != nil {
		return C.longlong(signal.(func() std_widgets.QSizePolicy__ControlType)())
	}

	return C.longlong(NewQGridLayoutWithTriggerSlotFromPointer(ptr).ControlTypesDefault())
}

func (ptr *QGridLayoutWithTriggerSlot) ControlTypesDefault() std_widgets.QSizePolicy__ControlType {
	if ptr.Pointer() != nil {
		return std_widgets.QSizePolicy__ControlType(C.QGridLayoutWithTriggerSlot_ControlTypesDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQGridLayoutWithTriggerSlot_IsEmpty
func callbackQGridLayoutWithTriggerSlot_IsEmpty(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isEmpty"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGridLayoutWithTriggerSlotFromPointer(ptr).IsEmptyDefault())))
}

func (ptr *QGridLayoutWithTriggerSlot) IsEmptyDefault() bool {
	if ptr.Pointer() != nil {
		return C.QGridLayoutWithTriggerSlot_IsEmptyDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQGridLayoutWithTriggerSlot_IndexOf
func callbackQGridLayoutWithTriggerSlot_IndexOf(ptr unsafe.Pointer, widget unsafe.Pointer) C.int {
	if signal := qt.GetSignal(ptr, "indexOf"); signal != nil {
		return C.int(int32(signal.(func(*std_widgets.QWidget) int)(std_widgets.NewQWidgetFromPointer(widget))))
	}

	return C.int(int32(NewQGridLayoutWithTriggerSlotFromPointer(ptr).IndexOfDefault(std_widgets.NewQWidgetFromPointer(widget))))
}

func (ptr *QGridLayoutWithTriggerSlot) IndexOfDefault(widget std_widgets.QWidget_ITF) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QGridLayoutWithTriggerSlot_IndexOfDefault(ptr.Pointer(), std_widgets.PointerFromQWidget(widget))))
	}
	return 0
}

//export callbackQGridLayoutWithTriggerSlot_Event
func callbackQGridLayoutWithTriggerSlot_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QEvent) bool)(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGridLayoutWithTriggerSlotFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *QGridLayoutWithTriggerSlot) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGridLayoutWithTriggerSlot_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQGridLayoutWithTriggerSlot_EventFilter
func callbackQGridLayoutWithTriggerSlot_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*std_core.QObject, *std_core.QEvent) bool)(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQGridLayoutWithTriggerSlotFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *QGridLayoutWithTriggerSlot) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGridLayoutWithTriggerSlot_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQGridLayoutWithTriggerSlot_ConnectNotify
func callbackQGridLayoutWithTriggerSlot_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGridLayoutWithTriggerSlotFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGridLayoutWithTriggerSlot) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGridLayoutWithTriggerSlot_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGridLayoutWithTriggerSlot_CustomEvent
func callbackQGridLayoutWithTriggerSlot_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*std_core.QEvent))(std_core.NewQEventFromPointer(event))
	} else {
		NewQGridLayoutWithTriggerSlotFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *QGridLayoutWithTriggerSlot) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGridLayoutWithTriggerSlot_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackQGridLayoutWithTriggerSlot_DeleteLater
func callbackQGridLayoutWithTriggerSlot_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQGridLayoutWithTriggerSlotFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QGridLayoutWithTriggerSlot) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QGridLayoutWithTriggerSlot_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQGridLayoutWithTriggerSlot_Destroyed
func callbackQGridLayoutWithTriggerSlot_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*std_core.QObject))(std_core.NewQObjectFromPointer(obj))
	}

}

//export callbackQGridLayoutWithTriggerSlot_DisconnectNotify
func callbackQGridLayoutWithTriggerSlot_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*std_core.QMetaMethod))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQGridLayoutWithTriggerSlotFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QGridLayoutWithTriggerSlot) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QGridLayoutWithTriggerSlot_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQGridLayoutWithTriggerSlot_ObjectNameChanged
func callbackQGridLayoutWithTriggerSlot_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQGridLayoutWithTriggerSlot_TimerEvent
func callbackQGridLayoutWithTriggerSlot_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*std_core.QTimerEvent))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewQGridLayoutWithTriggerSlotFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGridLayoutWithTriggerSlot) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QGridLayoutWithTriggerSlot_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

//export callbackQGridLayoutWithTriggerSlot_SpacerItem
func callbackQGridLayoutWithTriggerSlot_SpacerItem(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "spacerItem"); signal != nil {
		return std_widgets.PointerFromQSpacerItem(signal.(func() *std_widgets.QSpacerItem)())
	}

	return std_widgets.PointerFromQSpacerItem(NewQGridLayoutWithTriggerSlotFromPointer(ptr).SpacerItemDefault())
}

func (ptr *QGridLayoutWithTriggerSlot) SpacerItemDefault() *std_widgets.QSpacerItem {
	if ptr.Pointer() != nil {
		return std_widgets.NewQSpacerItemFromPointer(C.QGridLayoutWithTriggerSlot_SpacerItemDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQGridLayoutWithTriggerSlot_Widget
func callbackQGridLayoutWithTriggerSlot_Widget(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "widget"); signal != nil {
		return std_widgets.PointerFromQWidget(signal.(func() *std_widgets.QWidget)())
	}

	return std_widgets.PointerFromQWidget(NewQGridLayoutWithTriggerSlotFromPointer(ptr).WidgetDefault())
}

func (ptr *QGridLayoutWithTriggerSlot) WidgetDefault() *std_widgets.QWidget {
	if ptr.Pointer() != nil {
		var tmpValue = std_widgets.NewQWidgetFromPointer(C.QGridLayoutWithTriggerSlot_WidgetDefault(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}
