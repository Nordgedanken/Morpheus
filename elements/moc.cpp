

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QChildEvent>
#include <QEvent>
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
#include <QVBoxLayout>
#include <QWidget>


class QRoomVBoxLayoutWithTriggerSlot: public QVBoxLayout
{
Q_OBJECT
public:
	QRoomVBoxLayoutWithTriggerSlot() : QVBoxLayout() {qRegisterMetaType<quintptr>("quintptr");QRoomVBoxLayoutWithTriggerSlot_QRoomVBoxLayoutWithTriggerSlot_QRegisterMetaType();QRoomVBoxLayoutWithTriggerSlot_QRoomVBoxLayoutWithTriggerSlot_QRegisterMetaTypes();callbackQRoomVBoxLayoutWithTriggerSlot_Constructor(this);};
	QRoomVBoxLayoutWithTriggerSlot(QWidget *parent) : QVBoxLayout(parent) {qRegisterMetaType<quintptr>("quintptr");QRoomVBoxLayoutWithTriggerSlot_QRoomVBoxLayoutWithTriggerSlot_QRegisterMetaType();QRoomVBoxLayoutWithTriggerSlot_QRoomVBoxLayoutWithTriggerSlot_QRegisterMetaTypes();callbackQRoomVBoxLayoutWithTriggerSlot_Constructor(this);};
	QLayoutItem * takeAt(int index) { return static_cast<QLayoutItem*>(callbackQRoomVBoxLayoutWithTriggerSlot_TakeAt(this, index)); };
	void addItem(QLayoutItem * item) { callbackQRoomVBoxLayoutWithTriggerSlot_AddItem(this, item); };
	void invalidate() { callbackQRoomVBoxLayoutWithTriggerSlot_Invalidate(this); };
	void setGeometry(const QRect & r) { callbackQRoomVBoxLayoutWithTriggerSlot_SetGeometry(this, const_cast<QRect*>(&r)); };
	QLayoutItem * itemAt(int index) const { return static_cast<QLayoutItem*>(callbackQRoomVBoxLayoutWithTriggerSlot_ItemAt(const_cast<void*>(static_cast<const void*>(this)), index)); };
	QSize maximumSize() const { return *static_cast<QSize*>(callbackQRoomVBoxLayoutWithTriggerSlot_MaximumSize(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize minimumSize() const { return *static_cast<QSize*>(callbackQRoomVBoxLayoutWithTriggerSlot_MinimumSize(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQRoomVBoxLayoutWithTriggerSlot_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::Orientations expandingDirections() const { return static_cast<Qt::Orientation>(callbackQRoomVBoxLayoutWithTriggerSlot_ExpandingDirections(const_cast<void*>(static_cast<const void*>(this)))); };
	bool hasHeightForWidth() const { return callbackQRoomVBoxLayoutWithTriggerSlot_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int count() const { return callbackQRoomVBoxLayoutWithTriggerSlot_Count(const_cast<void*>(static_cast<const void*>(this))); };
	int heightForWidth(int w) const { return callbackQRoomVBoxLayoutWithTriggerSlot_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	int minimumHeightForWidth(int w) const { return callbackQRoomVBoxLayoutWithTriggerSlot_MinimumHeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	QLayout * layout() { return static_cast<QLayout*>(callbackQRoomVBoxLayoutWithTriggerSlot_Layout(this)); };
	void childEvent(QChildEvent * e) { callbackQRoomVBoxLayoutWithTriggerSlot_ChildEvent(this, e); };
	QRect geometry() const { return *static_cast<QRect*>(callbackQRoomVBoxLayoutWithTriggerSlot_Geometry(const_cast<void*>(static_cast<const void*>(this)))); };
	QSizePolicy::ControlTypes controlTypes() const { return static_cast<QSizePolicy::ControlType>(callbackQRoomVBoxLayoutWithTriggerSlot_ControlTypes(const_cast<void*>(static_cast<const void*>(this)))); };
	bool isEmpty() const { return callbackQRoomVBoxLayoutWithTriggerSlot_IsEmpty(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int indexOf(QWidget * widget) const { return callbackQRoomVBoxLayoutWithTriggerSlot_IndexOf(const_cast<void*>(static_cast<const void*>(this)), widget); };
	bool event(QEvent * e) { return callbackQRoomVBoxLayoutWithTriggerSlot_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQRoomVBoxLayoutWithTriggerSlot_EventFilter(this, watched, event) != 0; };
	void connectNotify(const QMetaMethod & sign) { callbackQRoomVBoxLayoutWithTriggerSlot_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQRoomVBoxLayoutWithTriggerSlot_CustomEvent(this, event); };
	void deleteLater() { callbackQRoomVBoxLayoutWithTriggerSlot_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQRoomVBoxLayoutWithTriggerSlot_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQRoomVBoxLayoutWithTriggerSlot_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQRoomVBoxLayoutWithTriggerSlot_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQRoomVBoxLayoutWithTriggerSlot_TimerEvent(this, event); };
	
	QSpacerItem * spacerItem() { return static_cast<QSpacerItem*>(callbackQRoomVBoxLayoutWithTriggerSlot_SpacerItem(this)); };
	QWidget * widget() { return static_cast<QWidget*>(callbackQRoomVBoxLayoutWithTriggerSlot_Widget(this)); };
signals:
public slots:
	void TriggerRoom(QString roomID) { QByteArray tdf131c = roomID.toUtf8(); Moc_PackedString roomIDPacked = { const_cast<char*>(tdf131c.prepend("WHITESPACE").constData()+10), tdf131c.size()-10 };callbackQRoomVBoxLayoutWithTriggerSlot_TriggerRoom(this, roomIDPacked); };
private:
};

Q_DECLARE_METATYPE(QRoomVBoxLayoutWithTriggerSlot*)


void QRoomVBoxLayoutWithTriggerSlot_QRoomVBoxLayoutWithTriggerSlot_QRegisterMetaTypes() {
}

class QVBoxLayoutWithTriggerSlot: public QVBoxLayout
{
Q_OBJECT
public:
	QVBoxLayoutWithTriggerSlot() : QVBoxLayout() {qRegisterMetaType<quintptr>("quintptr");QVBoxLayoutWithTriggerSlot_QVBoxLayoutWithTriggerSlot_QRegisterMetaType();QVBoxLayoutWithTriggerSlot_QVBoxLayoutWithTriggerSlot_QRegisterMetaTypes();callbackQVBoxLayoutWithTriggerSlot_Constructor(this);};
	QVBoxLayoutWithTriggerSlot(QWidget *parent) : QVBoxLayout(parent) {qRegisterMetaType<quintptr>("quintptr");QVBoxLayoutWithTriggerSlot_QVBoxLayoutWithTriggerSlot_QRegisterMetaType();QVBoxLayoutWithTriggerSlot_QVBoxLayoutWithTriggerSlot_QRegisterMetaTypes();callbackQVBoxLayoutWithTriggerSlot_Constructor(this);};
	QLayoutItem * takeAt(int index) { return static_cast<QLayoutItem*>(callbackQVBoxLayoutWithTriggerSlot_TakeAt(this, index)); };
	void addItem(QLayoutItem * item) { callbackQVBoxLayoutWithTriggerSlot_AddItem(this, item); };
	void invalidate() { callbackQVBoxLayoutWithTriggerSlot_Invalidate(this); };
	void setGeometry(const QRect & r) { callbackQVBoxLayoutWithTriggerSlot_SetGeometry(this, const_cast<QRect*>(&r)); };
	QLayoutItem * itemAt(int index) const { return static_cast<QLayoutItem*>(callbackQVBoxLayoutWithTriggerSlot_ItemAt(const_cast<void*>(static_cast<const void*>(this)), index)); };
	QSize maximumSize() const { return *static_cast<QSize*>(callbackQVBoxLayoutWithTriggerSlot_MaximumSize(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize minimumSize() const { return *static_cast<QSize*>(callbackQVBoxLayoutWithTriggerSlot_MinimumSize(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQVBoxLayoutWithTriggerSlot_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::Orientations expandingDirections() const { return static_cast<Qt::Orientation>(callbackQVBoxLayoutWithTriggerSlot_ExpandingDirections(const_cast<void*>(static_cast<const void*>(this)))); };
	bool hasHeightForWidth() const { return callbackQVBoxLayoutWithTriggerSlot_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int count() const { return callbackQVBoxLayoutWithTriggerSlot_Count(const_cast<void*>(static_cast<const void*>(this))); };
	int heightForWidth(int w) const { return callbackQVBoxLayoutWithTriggerSlot_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	int minimumHeightForWidth(int w) const { return callbackQVBoxLayoutWithTriggerSlot_MinimumHeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	QLayout * layout() { return static_cast<QLayout*>(callbackQVBoxLayoutWithTriggerSlot_Layout(this)); };
	void childEvent(QChildEvent * e) { callbackQVBoxLayoutWithTriggerSlot_ChildEvent(this, e); };
	QRect geometry() const { return *static_cast<QRect*>(callbackQVBoxLayoutWithTriggerSlot_Geometry(const_cast<void*>(static_cast<const void*>(this)))); };
	QSizePolicy::ControlTypes controlTypes() const { return static_cast<QSizePolicy::ControlType>(callbackQVBoxLayoutWithTriggerSlot_ControlTypes(const_cast<void*>(static_cast<const void*>(this)))); };
	bool isEmpty() const { return callbackQVBoxLayoutWithTriggerSlot_IsEmpty(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int indexOf(QWidget * widget) const { return callbackQVBoxLayoutWithTriggerSlot_IndexOf(const_cast<void*>(static_cast<const void*>(this)), widget); };
	bool event(QEvent * e) { return callbackQVBoxLayoutWithTriggerSlot_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQVBoxLayoutWithTriggerSlot_EventFilter(this, watched, event) != 0; };
	void connectNotify(const QMetaMethod & sign) { callbackQVBoxLayoutWithTriggerSlot_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQVBoxLayoutWithTriggerSlot_CustomEvent(this, event); };
	void deleteLater() { callbackQVBoxLayoutWithTriggerSlot_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQVBoxLayoutWithTriggerSlot_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQVBoxLayoutWithTriggerSlot_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQVBoxLayoutWithTriggerSlot_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQVBoxLayoutWithTriggerSlot_TimerEvent(this, event); };
	
	QSpacerItem * spacerItem() { return static_cast<QSpacerItem*>(callbackQVBoxLayoutWithTriggerSlot_SpacerItem(this)); };
	QWidget * widget() { return static_cast<QWidget*>(callbackQVBoxLayoutWithTriggerSlot_Widget(this)); };
signals:
public slots:
	void TriggerMessage(QString messageBody, QString sender, qint64 timestamp) { QByteArray t59bf16 = messageBody.toUtf8(); Moc_PackedString messageBodyPacked = { const_cast<char*>(t59bf16.prepend("WHITESPACE").constData()+10), t59bf16.size()-10 };QByteArray tacc6a3 = sender.toUtf8(); Moc_PackedString senderPacked = { const_cast<char*>(tacc6a3.prepend("WHITESPACE").constData()+10), tacc6a3.size()-10 };callbackQVBoxLayoutWithTriggerSlot_TriggerMessage(this, messageBodyPacked, senderPacked, timestamp); };
private:
};

Q_DECLARE_METATYPE(QVBoxLayoutWithTriggerSlot*)


void QVBoxLayoutWithTriggerSlot_QVBoxLayoutWithTriggerSlot_QRegisterMetaTypes() {
}

void QRoomVBoxLayoutWithTriggerSlot_TriggerRoom(void* ptr, struct Moc_PackedString roomID)
{
	QMetaObject::invokeMethod(static_cast<QRoomVBoxLayoutWithTriggerSlot*>(ptr), "TriggerRoom", Q_ARG(QString, QString::fromUtf8(roomID.data, roomID.len)));
}

int QRoomVBoxLayoutWithTriggerSlot_QRoomVBoxLayoutWithTriggerSlot_QRegisterMetaType()
{
	return qRegisterMetaType<QRoomVBoxLayoutWithTriggerSlot*>();
}

int QRoomVBoxLayoutWithTriggerSlot_QRoomVBoxLayoutWithTriggerSlot_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<QRoomVBoxLayoutWithTriggerSlot*>(const_cast<const char*>(typeName));
}

int QRoomVBoxLayoutWithTriggerSlot_QRoomVBoxLayoutWithTriggerSlot_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QRoomVBoxLayoutWithTriggerSlot>();
#else
	return 0;
#endif
}

int QRoomVBoxLayoutWithTriggerSlot_QRoomVBoxLayoutWithTriggerSlot_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QRoomVBoxLayoutWithTriggerSlot>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* QRoomVBoxLayoutWithTriggerSlot___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QRoomVBoxLayoutWithTriggerSlot___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QRoomVBoxLayoutWithTriggerSlot___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QRoomVBoxLayoutWithTriggerSlot___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QRoomVBoxLayoutWithTriggerSlot___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QRoomVBoxLayoutWithTriggerSlot___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QRoomVBoxLayoutWithTriggerSlot___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QRoomVBoxLayoutWithTriggerSlot___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QRoomVBoxLayoutWithTriggerSlot___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QRoomVBoxLayoutWithTriggerSlot___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QRoomVBoxLayoutWithTriggerSlot___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QRoomVBoxLayoutWithTriggerSlot___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QRoomVBoxLayoutWithTriggerSlot___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QRoomVBoxLayoutWithTriggerSlot___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QRoomVBoxLayoutWithTriggerSlot___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

void* QRoomVBoxLayoutWithTriggerSlot_NewQRoomVBoxLayoutWithTriggerSlot()
{
	return new QRoomVBoxLayoutWithTriggerSlot();
}

void* QRoomVBoxLayoutWithTriggerSlot_NewQRoomVBoxLayoutWithTriggerSlot2(void* parent)
{
		return new QRoomVBoxLayoutWithTriggerSlot(static_cast<QWidget*>(parent));
}

void QRoomVBoxLayoutWithTriggerSlot_DestroyQRoomVBoxLayoutWithTriggerSlot(void* ptr)
{
	static_cast<QRoomVBoxLayoutWithTriggerSlot*>(ptr)->~QRoomVBoxLayoutWithTriggerSlot();
}

void* QRoomVBoxLayoutWithTriggerSlot_TakeAtDefault(void* ptr, int index)
{
	return static_cast<QRoomVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::takeAt(index);
}

void QRoomVBoxLayoutWithTriggerSlot_AddItemDefault(void* ptr, void* item)
{
	static_cast<QRoomVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::addItem(static_cast<QLayoutItem*>(item));
}

void QRoomVBoxLayoutWithTriggerSlot_InvalidateDefault(void* ptr)
{
	static_cast<QRoomVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::invalidate();
}

void QRoomVBoxLayoutWithTriggerSlot_SetGeometryDefault(void* ptr, void* r)
{
	static_cast<QRoomVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::setGeometry(*static_cast<QRect*>(r));
}

void* QRoomVBoxLayoutWithTriggerSlot_ItemAtDefault(void* ptr, int index)
{
	return static_cast<QRoomVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::itemAt(index);
}

void* QRoomVBoxLayoutWithTriggerSlot_MaximumSizeDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QRoomVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::maximumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QRoomVBoxLayoutWithTriggerSlot_MinimumSizeDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QRoomVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::minimumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QRoomVBoxLayoutWithTriggerSlot_SizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QRoomVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

long long QRoomVBoxLayoutWithTriggerSlot_ExpandingDirectionsDefault(void* ptr)
{
	return static_cast<QRoomVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::expandingDirections();
}

char QRoomVBoxLayoutWithTriggerSlot_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<QRoomVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::hasHeightForWidth();
}

int QRoomVBoxLayoutWithTriggerSlot_CountDefault(void* ptr)
{
	return static_cast<QRoomVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::count();
}

int QRoomVBoxLayoutWithTriggerSlot_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<QRoomVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::heightForWidth(w);
}

int QRoomVBoxLayoutWithTriggerSlot_MinimumHeightForWidthDefault(void* ptr, int w)
{
	return static_cast<QRoomVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::minimumHeightForWidth(w);
}

void* QRoomVBoxLayoutWithTriggerSlot_LayoutDefault(void* ptr)
{
	return static_cast<QRoomVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::layout();
}

void QRoomVBoxLayoutWithTriggerSlot_ChildEventDefault(void* ptr, void* e)
{
	static_cast<QRoomVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::childEvent(static_cast<QChildEvent*>(e));
}

void* QRoomVBoxLayoutWithTriggerSlot_GeometryDefault(void* ptr)
{
	return ({ QRect tmpValue = static_cast<QRoomVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::geometry(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

long long QRoomVBoxLayoutWithTriggerSlot_ControlTypesDefault(void* ptr)
{
	return static_cast<QRoomVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::controlTypes();
}

char QRoomVBoxLayoutWithTriggerSlot_IsEmptyDefault(void* ptr)
{
	return static_cast<QRoomVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::isEmpty();
}

int QRoomVBoxLayoutWithTriggerSlot_IndexOfDefault(void* ptr, void* widget)
{
	return static_cast<QRoomVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::indexOf(static_cast<QWidget*>(widget));
}

char QRoomVBoxLayoutWithTriggerSlot_EventDefault(void* ptr, void* e)
{
	return static_cast<QRoomVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::event(static_cast<QEvent*>(e));
}

char QRoomVBoxLayoutWithTriggerSlot_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QRoomVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QRoomVBoxLayoutWithTriggerSlot_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QRoomVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QRoomVBoxLayoutWithTriggerSlot_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QRoomVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::customEvent(static_cast<QEvent*>(event));
}

void QRoomVBoxLayoutWithTriggerSlot_DeleteLaterDefault(void* ptr)
{
	static_cast<QRoomVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::deleteLater();
}

void QRoomVBoxLayoutWithTriggerSlot_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QRoomVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QRoomVBoxLayoutWithTriggerSlot_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QRoomVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::timerEvent(static_cast<QTimerEvent*>(event));
}



void* QRoomVBoxLayoutWithTriggerSlot_SpacerItemDefault(void* ptr)
{
	return static_cast<QRoomVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::spacerItem();
}

void* QRoomVBoxLayoutWithTriggerSlot_WidgetDefault(void* ptr)
{
	return static_cast<QRoomVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::widget();
}

void QVBoxLayoutWithTriggerSlot_TriggerMessage(void* ptr, struct Moc_PackedString messageBody, struct Moc_PackedString sender, long long timestamp)
{
	QMetaObject::invokeMethod(static_cast<QVBoxLayoutWithTriggerSlot*>(ptr), "TriggerMessage", Q_ARG(QString, QString::fromUtf8(messageBody.data, messageBody.len)), Q_ARG(QString, QString::fromUtf8(sender.data, sender.len)), Q_ARG(qint64, timestamp));
}

int QVBoxLayoutWithTriggerSlot_QVBoxLayoutWithTriggerSlot_QRegisterMetaType()
{
	return qRegisterMetaType<QVBoxLayoutWithTriggerSlot*>();
}

int QVBoxLayoutWithTriggerSlot_QVBoxLayoutWithTriggerSlot_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<QVBoxLayoutWithTriggerSlot*>(const_cast<const char*>(typeName));
}

int QVBoxLayoutWithTriggerSlot_QVBoxLayoutWithTriggerSlot_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QVBoxLayoutWithTriggerSlot>();
#else
	return 0;
#endif
}

int QVBoxLayoutWithTriggerSlot_QVBoxLayoutWithTriggerSlot_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QVBoxLayoutWithTriggerSlot>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* QVBoxLayoutWithTriggerSlot___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QVBoxLayoutWithTriggerSlot___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QVBoxLayoutWithTriggerSlot___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QVBoxLayoutWithTriggerSlot___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QVBoxLayoutWithTriggerSlot___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QVBoxLayoutWithTriggerSlot___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QVBoxLayoutWithTriggerSlot___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QVBoxLayoutWithTriggerSlot___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QVBoxLayoutWithTriggerSlot___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QVBoxLayoutWithTriggerSlot___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QVBoxLayoutWithTriggerSlot___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QVBoxLayoutWithTriggerSlot___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QVBoxLayoutWithTriggerSlot___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QVBoxLayoutWithTriggerSlot___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QVBoxLayoutWithTriggerSlot___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

void* QVBoxLayoutWithTriggerSlot_NewQVBoxLayoutWithTriggerSlot()
{
	return new QVBoxLayoutWithTriggerSlot();
}

void* QVBoxLayoutWithTriggerSlot_NewQVBoxLayoutWithTriggerSlot2(void* parent)
{
		return new QVBoxLayoutWithTriggerSlot(static_cast<QWidget*>(parent));
}

void QVBoxLayoutWithTriggerSlot_DestroyQVBoxLayoutWithTriggerSlot(void* ptr)
{
	static_cast<QVBoxLayoutWithTriggerSlot*>(ptr)->~QVBoxLayoutWithTriggerSlot();
}

void* QVBoxLayoutWithTriggerSlot_TakeAtDefault(void* ptr, int index)
{
	return static_cast<QVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::takeAt(index);
}

void QVBoxLayoutWithTriggerSlot_AddItemDefault(void* ptr, void* item)
{
	static_cast<QVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::addItem(static_cast<QLayoutItem*>(item));
}

void QVBoxLayoutWithTriggerSlot_InvalidateDefault(void* ptr)
{
	static_cast<QVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::invalidate();
}

void QVBoxLayoutWithTriggerSlot_SetGeometryDefault(void* ptr, void* r)
{
	static_cast<QVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::setGeometry(*static_cast<QRect*>(r));
}

void* QVBoxLayoutWithTriggerSlot_ItemAtDefault(void* ptr, int index)
{
	return static_cast<QVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::itemAt(index);
}

void* QVBoxLayoutWithTriggerSlot_MaximumSizeDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::maximumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QVBoxLayoutWithTriggerSlot_MinimumSizeDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::minimumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QVBoxLayoutWithTriggerSlot_SizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

long long QVBoxLayoutWithTriggerSlot_ExpandingDirectionsDefault(void* ptr)
{
	return static_cast<QVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::expandingDirections();
}

char QVBoxLayoutWithTriggerSlot_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<QVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::hasHeightForWidth();
}

int QVBoxLayoutWithTriggerSlot_CountDefault(void* ptr)
{
	return static_cast<QVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::count();
}

int QVBoxLayoutWithTriggerSlot_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<QVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::heightForWidth(w);
}

int QVBoxLayoutWithTriggerSlot_MinimumHeightForWidthDefault(void* ptr, int w)
{
	return static_cast<QVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::minimumHeightForWidth(w);
}

void* QVBoxLayoutWithTriggerSlot_LayoutDefault(void* ptr)
{
	return static_cast<QVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::layout();
}

void QVBoxLayoutWithTriggerSlot_ChildEventDefault(void* ptr, void* e)
{
	static_cast<QVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::childEvent(static_cast<QChildEvent*>(e));
}

void* QVBoxLayoutWithTriggerSlot_GeometryDefault(void* ptr)
{
	return ({ QRect tmpValue = static_cast<QVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::geometry(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

long long QVBoxLayoutWithTriggerSlot_ControlTypesDefault(void* ptr)
{
	return static_cast<QVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::controlTypes();
}

char QVBoxLayoutWithTriggerSlot_IsEmptyDefault(void* ptr)
{
	return static_cast<QVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::isEmpty();
}

int QVBoxLayoutWithTriggerSlot_IndexOfDefault(void* ptr, void* widget)
{
	return static_cast<QVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::indexOf(static_cast<QWidget*>(widget));
}

char QVBoxLayoutWithTriggerSlot_EventDefault(void* ptr, void* e)
{
	return static_cast<QVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::event(static_cast<QEvent*>(e));
}

char QVBoxLayoutWithTriggerSlot_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QVBoxLayoutWithTriggerSlot_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QVBoxLayoutWithTriggerSlot_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::customEvent(static_cast<QEvent*>(event));
}

void QVBoxLayoutWithTriggerSlot_DeleteLaterDefault(void* ptr)
{
	static_cast<QVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::deleteLater();
}

void QVBoxLayoutWithTriggerSlot_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QVBoxLayoutWithTriggerSlot_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::timerEvent(static_cast<QTimerEvent*>(event));
}



void* QVBoxLayoutWithTriggerSlot_SpacerItemDefault(void* ptr)
{
	return static_cast<QVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::spacerItem();
}

void* QVBoxLayoutWithTriggerSlot_WidgetDefault(void* ptr)
{
	return static_cast<QVBoxLayoutWithTriggerSlot*>(ptr)->QVBoxLayout::widget();
}

#include "moc_moc.h"
