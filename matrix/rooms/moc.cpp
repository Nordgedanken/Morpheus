

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QChildEvent>
#include <QEvent>
#include <QHBoxLayout>
#include <QLayout>
#include <QLayoutItem>
#include <QList>
#include <QMetaMethod>
#include <QObject>
#include <QPixmap>
#include <QRect>
#include <QSize>
#include <QSizePolicy>
#include <QSpacerItem>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QWidget>


class Room: public QHBoxLayout
{
Q_OBJECT
public:
	Room() : QHBoxLayout() {qRegisterMetaType<quintptr>("quintptr");Room_Room_QRegisterMetaType();Room_Room_QRegisterMetaTypes();callbackRoom_Constructor(this);};
	Room(QWidget *parent) : QHBoxLayout(parent) {qRegisterMetaType<quintptr>("quintptr");Room_Room_QRegisterMetaType();Room_Room_QRegisterMetaTypes();callbackRoom_Constructor(this);};
	void Signal_SetAvatar(QPixmap roomAvatar) { callbackRoom_SetAvatar(this, new QPixmap(roomAvatar)); };
	QLayoutItem * takeAt(int index) { return static_cast<QLayoutItem*>(callbackRoom_TakeAt(this, index)); };
	void addItem(QLayoutItem * item) { callbackRoom_AddItem(this, item); };
	void invalidate() { callbackRoom_Invalidate(this); };
	void setGeometry(const QRect & r) { callbackRoom_SetGeometry(this, const_cast<QRect*>(&r)); };
	QLayoutItem * itemAt(int index) const { return static_cast<QLayoutItem*>(callbackRoom_ItemAt(const_cast<void*>(static_cast<const void*>(this)), index)); };
	QSize maximumSize() const { return *static_cast<QSize*>(callbackRoom_MaximumSize(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize minimumSize() const { return *static_cast<QSize*>(callbackRoom_MinimumSize(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackRoom_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::Orientations expandingDirections() const { return static_cast<Qt::Orientation>(callbackRoom_ExpandingDirections(const_cast<void*>(static_cast<const void*>(this)))); };
	bool hasHeightForWidth() const { return callbackRoom_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int count() const { return callbackRoom_Count(const_cast<void*>(static_cast<const void*>(this))); };
	int heightForWidth(int w) const { return callbackRoom_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	int minimumHeightForWidth(int w) const { return callbackRoom_MinimumHeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	QLayout * layout() { return static_cast<QLayout*>(callbackRoom_Layout(this)); };
	void childEvent(QChildEvent * e) { callbackRoom_ChildEvent(this, e); };
	QRect geometry() const { return *static_cast<QRect*>(callbackRoom_Geometry(const_cast<void*>(static_cast<const void*>(this)))); };
	QSizePolicy::ControlTypes controlTypes() const { return static_cast<QSizePolicy::ControlType>(callbackRoom_ControlTypes(const_cast<void*>(static_cast<const void*>(this)))); };
	bool isEmpty() const { return callbackRoom_IsEmpty(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int indexOf(QWidget * widget) const { return callbackRoom_IndexOf(const_cast<void*>(static_cast<const void*>(this)), widget); };
	bool event(QEvent * e) { return callbackRoom_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackRoom_EventFilter(this, watched, event) != 0; };
	void connectNotify(const QMetaMethod & sign) { callbackRoom_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackRoom_CustomEvent(this, event); };
	void deleteLater() { callbackRoom_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackRoom_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackRoom_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackRoom_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackRoom_TimerEvent(this, event); };
	
	QSpacerItem * spacerItem() { return static_cast<QSpacerItem*>(callbackRoom_SpacerItem(this)); };
	QWidget * widget() { return static_cast<QWidget*>(callbackRoom_Widget(this)); };
signals:
	void SetAvatar(QPixmap roomAvatar);
public slots:
private:
};

Q_DECLARE_METATYPE(Room*)


void Room_Room_QRegisterMetaTypes() {
}

void Room_ConnectSetAvatar(void* ptr)
{
	QObject::connect(static_cast<Room*>(ptr), static_cast<void (Room::*)(QPixmap)>(&Room::SetAvatar), static_cast<Room*>(ptr), static_cast<void (Room::*)(QPixmap)>(&Room::Signal_SetAvatar));
}

void Room_DisconnectSetAvatar(void* ptr)
{
	QObject::disconnect(static_cast<Room*>(ptr), static_cast<void (Room::*)(QPixmap)>(&Room::SetAvatar), static_cast<Room*>(ptr), static_cast<void (Room::*)(QPixmap)>(&Room::Signal_SetAvatar));
}

void Room_SetAvatar(void* ptr, void* roomAvatar)
{
	static_cast<Room*>(ptr)->SetAvatar(*static_cast<QPixmap*>(roomAvatar));
}

int Room_Room_QRegisterMetaType()
{
	return qRegisterMetaType<Room*>();
}

int Room_Room_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<Room*>(const_cast<const char*>(typeName));
}

int Room_Room_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<Room>();
#else
	return 0;
#endif
}

int Room_Room_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<Room>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* Room___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void Room___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* Room___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* Room___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void Room___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Room___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* Room___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void Room___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Room___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* Room___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void Room___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Room___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* Room___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void Room___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Room___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

void* Room_NewRoom()
{
	return new Room();
}

void* Room_NewRoom2(void* parent)
{
		return new Room(static_cast<QWidget*>(parent));
}

void Room_DestroyRoom(void* ptr)
{
	static_cast<Room*>(ptr)->~Room();
}

void* Room_TakeAtDefault(void* ptr, int index)
{
	return static_cast<Room*>(ptr)->QHBoxLayout::takeAt(index);
}

void Room_AddItemDefault(void* ptr, void* item)
{
	static_cast<Room*>(ptr)->QHBoxLayout::addItem(static_cast<QLayoutItem*>(item));
}

void Room_InvalidateDefault(void* ptr)
{
	static_cast<Room*>(ptr)->QHBoxLayout::invalidate();
}

void Room_SetGeometryDefault(void* ptr, void* r)
{
	static_cast<Room*>(ptr)->QHBoxLayout::setGeometry(*static_cast<QRect*>(r));
}

void* Room_ItemAtDefault(void* ptr, int index)
{
	return static_cast<Room*>(ptr)->QHBoxLayout::itemAt(index);
}

void* Room_MaximumSizeDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<Room*>(ptr)->QHBoxLayout::maximumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* Room_MinimumSizeDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<Room*>(ptr)->QHBoxLayout::minimumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* Room_SizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<Room*>(ptr)->QHBoxLayout::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

long long Room_ExpandingDirectionsDefault(void* ptr)
{
	return static_cast<Room*>(ptr)->QHBoxLayout::expandingDirections();
}

char Room_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<Room*>(ptr)->QHBoxLayout::hasHeightForWidth();
}

int Room_CountDefault(void* ptr)
{
	return static_cast<Room*>(ptr)->QHBoxLayout::count();
}

int Room_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<Room*>(ptr)->QHBoxLayout::heightForWidth(w);
}

int Room_MinimumHeightForWidthDefault(void* ptr, int w)
{
	return static_cast<Room*>(ptr)->QHBoxLayout::minimumHeightForWidth(w);
}

void* Room_LayoutDefault(void* ptr)
{
	return static_cast<Room*>(ptr)->QHBoxLayout::layout();
}

void Room_ChildEventDefault(void* ptr, void* e)
{
	static_cast<Room*>(ptr)->QHBoxLayout::childEvent(static_cast<QChildEvent*>(e));
}

void* Room_GeometryDefault(void* ptr)
{
	return ({ QRect tmpValue = static_cast<Room*>(ptr)->QHBoxLayout::geometry(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

long long Room_ControlTypesDefault(void* ptr)
{
	return static_cast<Room*>(ptr)->QHBoxLayout::controlTypes();
}

char Room_IsEmptyDefault(void* ptr)
{
	return static_cast<Room*>(ptr)->QHBoxLayout::isEmpty();
}

int Room_IndexOfDefault(void* ptr, void* widget)
{
	return static_cast<Room*>(ptr)->QHBoxLayout::indexOf(static_cast<QWidget*>(widget));
}

char Room_EventDefault(void* ptr, void* e)
{
	return static_cast<Room*>(ptr)->QHBoxLayout::event(static_cast<QEvent*>(e));
}

char Room_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<Room*>(ptr)->QHBoxLayout::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void Room_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<Room*>(ptr)->QHBoxLayout::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void Room_CustomEventDefault(void* ptr, void* event)
{
	static_cast<Room*>(ptr)->QHBoxLayout::customEvent(static_cast<QEvent*>(event));
}

void Room_DeleteLaterDefault(void* ptr)
{
	static_cast<Room*>(ptr)->QHBoxLayout::deleteLater();
}

void Room_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<Room*>(ptr)->QHBoxLayout::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void Room_TimerEventDefault(void* ptr, void* event)
{
	static_cast<Room*>(ptr)->QHBoxLayout::timerEvent(static_cast<QTimerEvent*>(event));
}



void* Room_SpacerItemDefault(void* ptr)
{
	return static_cast<Room*>(ptr)->QHBoxLayout::spacerItem();
}

void* Room_WidgetDefault(void* ptr)
{
	return static_cast<Room*>(ptr)->QHBoxLayout::widget();
}

#include "moc_moc.h"
