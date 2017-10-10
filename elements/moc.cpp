

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QChildEvent>
#include <QEvent>
#include <QGridLayout>
#include <QLayout>
#include <QLayoutItem>
#include <QList>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QRect>
#include <QSize>
#include <QSizePolicy>
#include <QSpacerItem>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QWidget>


class QGridLayoutWithTriggerSlot: public QGridLayout
{
Q_OBJECT
public:
	QGridLayoutWithTriggerSlot() : QGridLayout() {qRegisterMetaType<quintptr>("quintptr");QGridLayoutWithTriggerSlot_QGridLayoutWithTriggerSlot_QRegisterMetaType();QGridLayoutWithTriggerSlot_QGridLayoutWithTriggerSlot_QRegisterMetaTypes();callbackQGridLayoutWithTriggerSlot_Constructor(this);};
	QGridLayoutWithTriggerSlot(QWidget *parent) : QGridLayout(parent) {qRegisterMetaType<quintptr>("quintptr");QGridLayoutWithTriggerSlot_QGridLayoutWithTriggerSlot_QRegisterMetaType();QGridLayoutWithTriggerSlot_QGridLayoutWithTriggerSlot_QRegisterMetaTypes();callbackQGridLayoutWithTriggerSlot_Constructor(this);};
	QLayoutItem * takeAt(int index) { return static_cast<QLayoutItem*>(callbackQGridLayoutWithTriggerSlot_TakeAt(this, index)); };
	void invalidate() { callbackQGridLayoutWithTriggerSlot_Invalidate(this); };
	void setGeometry(const QRect & rect) { callbackQGridLayoutWithTriggerSlot_SetGeometry(this, const_cast<QRect*>(&rect)); };
	QLayoutItem * itemAt(int index) const { return static_cast<QLayoutItem*>(callbackQGridLayoutWithTriggerSlot_ItemAt(const_cast<void*>(static_cast<const void*>(this)), index)); };
	QSize maximumSize() const { return *static_cast<QSize*>(callbackQGridLayoutWithTriggerSlot_MaximumSize(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize minimumSize() const { return *static_cast<QSize*>(callbackQGridLayoutWithTriggerSlot_MinimumSize(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQGridLayoutWithTriggerSlot_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::Orientations expandingDirections() const { return static_cast<Qt::Orientation>(callbackQGridLayoutWithTriggerSlot_ExpandingDirections(const_cast<void*>(static_cast<const void*>(this)))); };
	bool hasHeightForWidth() const { return callbackQGridLayoutWithTriggerSlot_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int count() const { return callbackQGridLayoutWithTriggerSlot_Count(const_cast<void*>(static_cast<const void*>(this))); };
	int heightForWidth(int w) const { return callbackQGridLayoutWithTriggerSlot_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	int minimumHeightForWidth(int w) const { return callbackQGridLayoutWithTriggerSlot_MinimumHeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	QLayout * layout() { return static_cast<QLayout*>(callbackQGridLayoutWithTriggerSlot_Layout(this)); };
	void addItem(QLayoutItem * item) { callbackQGridLayoutWithTriggerSlot_AddItem(this, item); };
	void childEvent(QChildEvent * e) { callbackQGridLayoutWithTriggerSlot_ChildEvent(this, e); };
	QRect geometry() const { return *static_cast<QRect*>(callbackQGridLayoutWithTriggerSlot_Geometry(const_cast<void*>(static_cast<const void*>(this)))); };
	QSizePolicy::ControlTypes controlTypes() const { return static_cast<QSizePolicy::ControlType>(callbackQGridLayoutWithTriggerSlot_ControlTypes(const_cast<void*>(static_cast<const void*>(this)))); };
	bool isEmpty() const { return callbackQGridLayoutWithTriggerSlot_IsEmpty(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int indexOf(QWidget * widget) const { return callbackQGridLayoutWithTriggerSlot_IndexOf(const_cast<void*>(static_cast<const void*>(this)), widget); };
	bool event(QEvent * e) { return callbackQGridLayoutWithTriggerSlot_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQGridLayoutWithTriggerSlot_EventFilter(this, watched, event) != 0; };
	void connectNotify(const QMetaMethod & sign) { callbackQGridLayoutWithTriggerSlot_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQGridLayoutWithTriggerSlot_CustomEvent(this, event); };
	void deleteLater() { callbackQGridLayoutWithTriggerSlot_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQGridLayoutWithTriggerSlot_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQGridLayoutWithTriggerSlot_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQGridLayoutWithTriggerSlot_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQGridLayoutWithTriggerSlot_TimerEvent(this, event); };
	
	QSpacerItem * spacerItem() { return static_cast<QSpacerItem*>(callbackQGridLayoutWithTriggerSlot_SpacerItem(this)); };
	QWidget * widget() { return static_cast<QWidget*>(callbackQGridLayoutWithTriggerSlot_Widget(this)); };
signals:
public slots:
	void TriggerMessage(QString messageBody) { QByteArray t59bf16 = messageBody.toUtf8(); Moc_PackedString messageBodyPacked = { const_cast<char*>(t59bf16.prepend("WHITESPACE").constData()+10), t59bf16.size()-10 };callbackQGridLayoutWithTriggerSlot_TriggerMessage(this, messageBodyPacked); };
private:
};

Q_DECLARE_METATYPE(QGridLayoutWithTriggerSlot*)


void QGridLayoutWithTriggerSlot_QGridLayoutWithTriggerSlot_QRegisterMetaTypes() {
}

void QGridLayoutWithTriggerSlot_TriggerMessage(void* ptr, struct Moc_PackedString messageBody)
{
	QMetaObject::invokeMethod(static_cast<QGridLayoutWithTriggerSlot*>(ptr), "TriggerMessage", Q_ARG(QString, QString::fromUtf8(messageBody.data, messageBody.len)));
}

int QGridLayoutWithTriggerSlot_QGridLayoutWithTriggerSlot_QRegisterMetaType()
{
	return qRegisterMetaType<QGridLayoutWithTriggerSlot*>();
}

int QGridLayoutWithTriggerSlot_QGridLayoutWithTriggerSlot_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<QGridLayoutWithTriggerSlot*>(const_cast<const char*>(typeName));
}

int QGridLayoutWithTriggerSlot_QGridLayoutWithTriggerSlot_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QGridLayoutWithTriggerSlot>();
#else
	return 0;
#endif
}

int QGridLayoutWithTriggerSlot_QGridLayoutWithTriggerSlot_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QGridLayoutWithTriggerSlot>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* QGridLayoutWithTriggerSlot___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QGridLayoutWithTriggerSlot___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QGridLayoutWithTriggerSlot___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QGridLayoutWithTriggerSlot___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QGridLayoutWithTriggerSlot___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGridLayoutWithTriggerSlot___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QGridLayoutWithTriggerSlot___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QGridLayoutWithTriggerSlot___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGridLayoutWithTriggerSlot___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QGridLayoutWithTriggerSlot___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QGridLayoutWithTriggerSlot___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGridLayoutWithTriggerSlot___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QGridLayoutWithTriggerSlot___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QGridLayoutWithTriggerSlot___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGridLayoutWithTriggerSlot___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

void* QGridLayoutWithTriggerSlot_NewQGridLayoutWithTriggerSlot2()
{
	return new QGridLayoutWithTriggerSlot();
}

void* QGridLayoutWithTriggerSlot_NewQGridLayoutWithTriggerSlot(void* parent)
{
		return new QGridLayoutWithTriggerSlot(static_cast<QWidget*>(parent));
}

void QGridLayoutWithTriggerSlot_DestroyQGridLayoutWithTriggerSlot(void* ptr)
{
	static_cast<QGridLayoutWithTriggerSlot*>(ptr)->~QGridLayoutWithTriggerSlot();
}

void* QGridLayoutWithTriggerSlot_TakeAtDefault(void* ptr, int index)
{
	return static_cast<QGridLayoutWithTriggerSlot*>(ptr)->QGridLayout::takeAt(index);
}

void QGridLayoutWithTriggerSlot_InvalidateDefault(void* ptr)
{
	static_cast<QGridLayoutWithTriggerSlot*>(ptr)->QGridLayout::invalidate();
}

void QGridLayoutWithTriggerSlot_SetGeometryDefault(void* ptr, void* rect)
{
	static_cast<QGridLayoutWithTriggerSlot*>(ptr)->QGridLayout::setGeometry(*static_cast<QRect*>(rect));
}

void* QGridLayoutWithTriggerSlot_ItemAtDefault(void* ptr, int index)
{
	return static_cast<QGridLayoutWithTriggerSlot*>(ptr)->QGridLayout::itemAt(index);
}

void* QGridLayoutWithTriggerSlot_MaximumSizeDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QGridLayoutWithTriggerSlot*>(ptr)->QGridLayout::maximumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QGridLayoutWithTriggerSlot_MinimumSizeDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QGridLayoutWithTriggerSlot*>(ptr)->QGridLayout::minimumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QGridLayoutWithTriggerSlot_SizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QGridLayoutWithTriggerSlot*>(ptr)->QGridLayout::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

long long QGridLayoutWithTriggerSlot_ExpandingDirectionsDefault(void* ptr)
{
	return static_cast<QGridLayoutWithTriggerSlot*>(ptr)->QGridLayout::expandingDirections();
}

char QGridLayoutWithTriggerSlot_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<QGridLayoutWithTriggerSlot*>(ptr)->QGridLayout::hasHeightForWidth();
}

int QGridLayoutWithTriggerSlot_CountDefault(void* ptr)
{
	return static_cast<QGridLayoutWithTriggerSlot*>(ptr)->QGridLayout::count();
}

int QGridLayoutWithTriggerSlot_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<QGridLayoutWithTriggerSlot*>(ptr)->QGridLayout::heightForWidth(w);
}

int QGridLayoutWithTriggerSlot_MinimumHeightForWidthDefault(void* ptr, int w)
{
	return static_cast<QGridLayoutWithTriggerSlot*>(ptr)->QGridLayout::minimumHeightForWidth(w);
}

void* QGridLayoutWithTriggerSlot_LayoutDefault(void* ptr)
{
	return static_cast<QGridLayoutWithTriggerSlot*>(ptr)->QGridLayout::layout();
}

void QGridLayoutWithTriggerSlot_AddItemDefault(void* ptr, void* item)
{
	static_cast<QGridLayoutWithTriggerSlot*>(ptr)->QGridLayout::addItem(static_cast<QLayoutItem*>(item));
}

void QGridLayoutWithTriggerSlot_ChildEventDefault(void* ptr, void* e)
{
	static_cast<QGridLayoutWithTriggerSlot*>(ptr)->QGridLayout::childEvent(static_cast<QChildEvent*>(e));
}

void* QGridLayoutWithTriggerSlot_GeometryDefault(void* ptr)
{
	return ({ QRect tmpValue = static_cast<QGridLayoutWithTriggerSlot*>(ptr)->QGridLayout::geometry(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

long long QGridLayoutWithTriggerSlot_ControlTypesDefault(void* ptr)
{
	return static_cast<QGridLayoutWithTriggerSlot*>(ptr)->QGridLayout::controlTypes();
}

char QGridLayoutWithTriggerSlot_IsEmptyDefault(void* ptr)
{
	return static_cast<QGridLayoutWithTriggerSlot*>(ptr)->QGridLayout::isEmpty();
}

int QGridLayoutWithTriggerSlot_IndexOfDefault(void* ptr, void* widget)
{
	return static_cast<QGridLayoutWithTriggerSlot*>(ptr)->QGridLayout::indexOf(static_cast<QWidget*>(widget));
}

char QGridLayoutWithTriggerSlot_EventDefault(void* ptr, void* e)
{
	return static_cast<QGridLayoutWithTriggerSlot*>(ptr)->QGridLayout::event(static_cast<QEvent*>(e));
}

char QGridLayoutWithTriggerSlot_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QGridLayoutWithTriggerSlot*>(ptr)->QGridLayout::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QGridLayoutWithTriggerSlot_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QGridLayoutWithTriggerSlot*>(ptr)->QGridLayout::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGridLayoutWithTriggerSlot_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QGridLayoutWithTriggerSlot*>(ptr)->QGridLayout::customEvent(static_cast<QEvent*>(event));
}

void QGridLayoutWithTriggerSlot_DeleteLaterDefault(void* ptr)
{
	static_cast<QGridLayoutWithTriggerSlot*>(ptr)->QGridLayout::deleteLater();
}

void QGridLayoutWithTriggerSlot_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QGridLayoutWithTriggerSlot*>(ptr)->QGridLayout::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGridLayoutWithTriggerSlot_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QGridLayoutWithTriggerSlot*>(ptr)->QGridLayout::timerEvent(static_cast<QTimerEvent*>(event));
}



void* QGridLayoutWithTriggerSlot_SpacerItemDefault(void* ptr)
{
	return static_cast<QGridLayoutWithTriggerSlot*>(ptr)->QGridLayout::spacerItem();
}

void* QGridLayoutWithTriggerSlot_WidgetDefault(void* ptr)
{
	return static_cast<QGridLayoutWithTriggerSlot*>(ptr)->QGridLayout::widget();
}

#include "moc_moc.h"
