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
func callbackQVBoxLayoutWithTriggerSlot_TriggerMessage(ptr unsafe.Pointer, messageBody C.struct_Moc_PackedString, sender C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "TriggerMessage"); signal != nil {
		signal.(func(string, string))(cGoUnpackString(messageBody), cGoUnpackString(sender))
	}

}

func (ptr *QVBoxLayoutWithTriggerSlot) ConnectTriggerMessage(f func(messageBody string, sender string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "TriggerMessage"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "TriggerMessage", func(messageBody string, sender string) {
				signal.(func(string, string))(messageBody, sender)
				f(messageBody, sender)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "TriggerMessage", f)
		}
	}
}

func (ptr *QVBoxLayoutWithTriggerSlot) DisconnectTriggerMessage() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "TriggerMessage")
	}
}

func (ptr *QVBoxLayoutWithTriggerSlot) TriggerMessage(messageBody string, sender string) {
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
		C.QVBoxLayoutWithTriggerSlot_TriggerMessage(ptr.Pointer(), C.struct_Moc_PackedString{data: messageBodyC, len: C.longlong(len(messageBody))}, C.struct_Moc_PackedString{data: senderC, len: C.longlong(len(sender))})
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
