

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QCamera>
#include <QCameraImageCapture>
#include <QChildEvent>
#include <QDBusPendingCall>
#include <QDBusPendingCallWatcher>
#include <QEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QLayout>
#include <QList>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMetaMethod>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDevice>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QQuickItem>
#include <QRadioData>
#include <QSignalSpy>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QWidget>
#include <QWindow>


class Room: public QObject
{
Q_OBJECT
public:
	Room(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");Room_Room_QRegisterMetaType();Room_Room_QRegisterMetaTypes();callbackRoom_Constructor(this);};
	void Signal_SetAvatar(quintptr IMGdata) { callbackRoom_SetAvatar(this, IMGdata); };
	 ~Room() { callbackRoom_DestroyRoom(this); };
	bool event(QEvent * e) { return callbackRoom_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackRoom_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackRoom_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackRoom_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackRoom_CustomEvent(this, event); };
	void deleteLater() { callbackRoom_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackRoom_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackRoom_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackRoom_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackRoom_TimerEvent(this, event); };
	
signals:
	void SetAvatar(quintptr IMGdata);
public slots:
private:
};

Q_DECLARE_METATYPE(Room*)


void Room_Room_QRegisterMetaTypes() {
}

void Room_ConnectSetAvatar(void* ptr)
{
	QObject::connect(static_cast<Room*>(ptr), static_cast<void (Room::*)(quintptr)>(&Room::SetAvatar), static_cast<Room*>(ptr), static_cast<void (Room::*)(quintptr)>(&Room::Signal_SetAvatar));
}

void Room_DisconnectSetAvatar(void* ptr)
{
	QObject::disconnect(static_cast<Room*>(ptr), static_cast<void (Room::*)(quintptr)>(&Room::SetAvatar), static_cast<Room*>(ptr), static_cast<void (Room::*)(quintptr)>(&Room::Signal_SetAvatar));
}

void Room_SetAvatar(void* ptr, uintptr_t IMGdata)
{
	static_cast<Room*>(ptr)->SetAvatar(IMGdata);
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

void* Room_NewRoom(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new Room(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new Room(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new Room(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new Room(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new Room(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new Room(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new Room(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new Room(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new Room(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new Room(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new Room(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new Room(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new Room(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new Room(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new Room(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new Room(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new Room(static_cast<QWindow*>(parent));
	} else {
		return new Room(static_cast<QObject*>(parent));
	}
}

void Room_DestroyRoom(void* ptr)
{
	static_cast<Room*>(ptr)->~Room();
}

void Room_DestroyRoomDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char Room_EventDefault(void* ptr, void* e)
{
	return static_cast<Room*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char Room_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<Room*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void Room_ChildEventDefault(void* ptr, void* event)
{
	static_cast<Room*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void Room_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<Room*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void Room_CustomEventDefault(void* ptr, void* event)
{
	static_cast<Room*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void Room_DeleteLaterDefault(void* ptr)
{
	static_cast<Room*>(ptr)->QObject::deleteLater();
}

void Room_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<Room*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void Room_TimerEventDefault(void* ptr, void* event)
{
	static_cast<Room*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}



#include "moc_moc.h"
