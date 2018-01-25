

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
#include <QPixmap>
#include <QQuickItem>
#include <QRadioData>
#include <QSignalSpy>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QWidget>
#include <QWindow>


class Message: public QObject
{
Q_OBJECT
public:
	Message(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");Message_Message_QRegisterMetaType();Message_Message_QRegisterMetaTypes();callbackMessage_Constructor(this);};
	void Signal_SetAvatar(QPixmap avatar) { callbackMessage_SetAvatar(this, new QPixmap(avatar)); };
	 ~Message() { callbackMessage_DestroyMessage(this); };
	bool event(QEvent * e) { return callbackMessage_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackMessage_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackMessage_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackMessage_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackMessage_CustomEvent(this, event); };
	void deleteLater() { callbackMessage_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackMessage_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackMessage_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackMessage_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackMessage_TimerEvent(this, event); };
	
signals:
	void SetAvatar(QPixmap avatar);
public slots:
private:
};

Q_DECLARE_METATYPE(Message*)


void Message_Message_QRegisterMetaTypes() {
}

void Message_ConnectSetAvatar(void* ptr)
{
	QObject::connect(static_cast<Message*>(ptr), static_cast<void (Message::*)(QPixmap)>(&Message::SetAvatar), static_cast<Message*>(ptr), static_cast<void (Message::*)(QPixmap)>(&Message::Signal_SetAvatar));
}

void Message_DisconnectSetAvatar(void* ptr)
{
	QObject::disconnect(static_cast<Message*>(ptr), static_cast<void (Message::*)(QPixmap)>(&Message::SetAvatar), static_cast<Message*>(ptr), static_cast<void (Message::*)(QPixmap)>(&Message::Signal_SetAvatar));
}

void Message_SetAvatar(void* ptr, void* avatar)
{
	static_cast<Message*>(ptr)->SetAvatar(*static_cast<QPixmap*>(avatar));
}

int Message_Message_QRegisterMetaType()
{
	return qRegisterMetaType<Message*>();
}

int Message_Message_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<Message*>(const_cast<const char*>(typeName));
}

int Message_Message_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<Message>();
#else
	return 0;
#endif
}

int Message_Message_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<Message>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* Message___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void Message___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* Message___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* Message___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void Message___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Message___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* Message___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void Message___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Message___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* Message___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void Message___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Message___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* Message___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void Message___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Message___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

void* Message_NewMessage(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new Message(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new Message(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new Message(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new Message(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new Message(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new Message(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new Message(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new Message(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new Message(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new Message(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new Message(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new Message(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new Message(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new Message(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new Message(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new Message(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new Message(static_cast<QWindow*>(parent));
	} else {
		return new Message(static_cast<QObject*>(parent));
	}
}

void Message_DestroyMessage(void* ptr)
{
	static_cast<Message*>(ptr)->~Message();
}

void Message_DestroyMessageDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char Message_EventDefault(void* ptr, void* e)
{
	return static_cast<Message*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char Message_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<Message*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void Message_ChildEventDefault(void* ptr, void* event)
{
	static_cast<Message*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void Message_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<Message*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void Message_CustomEventDefault(void* ptr, void* event)
{
	static_cast<Message*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void Message_DeleteLaterDefault(void* ptr)
{
	static_cast<Message*>(ptr)->QObject::deleteLater();
}

void Message_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<Message*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void Message_TimerEventDefault(void* ptr, void* event)
{
	static_cast<Message*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}



#include "moc_moc.h"
