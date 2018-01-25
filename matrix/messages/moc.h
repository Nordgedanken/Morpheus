

#pragma once

#ifndef GO_MOC_H
#define GO_MOC_H

#include <stdint.h>

#ifdef __cplusplus
class Message;
void Message_Message_QRegisterMetaTypes();
extern "C" {
#endif

struct Moc_PackedString { char* data; long long len; };
struct Moc_PackedList { void* data; long long len; };
void Message_ConnectSetAvatar(void* ptr);
void Message_DisconnectSetAvatar(void* ptr);
void Message_SetAvatar(void* ptr, void* avatar);
int Message_Message_QRegisterMetaType();
int Message_Message_QRegisterMetaType2(char* typeName);
int Message_Message_QmlRegisterType();
int Message_Message_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName);
void* Message___dynamicPropertyNames_atList(void* ptr, int i);
void Message___dynamicPropertyNames_setList(void* ptr, void* i);
void* Message___dynamicPropertyNames_newList(void* ptr);
void* Message___findChildren_atList2(void* ptr, int i);
void Message___findChildren_setList2(void* ptr, void* i);
void* Message___findChildren_newList2(void* ptr);
void* Message___findChildren_atList3(void* ptr, int i);
void Message___findChildren_setList3(void* ptr, void* i);
void* Message___findChildren_newList3(void* ptr);
void* Message___findChildren_atList(void* ptr, int i);
void Message___findChildren_setList(void* ptr, void* i);
void* Message___findChildren_newList(void* ptr);
void* Message___children_atList(void* ptr, int i);
void Message___children_setList(void* ptr, void* i);
void* Message___children_newList(void* ptr);
void* Message_NewMessage(void* parent);
void Message_DestroyMessage(void* ptr);
void Message_DestroyMessageDefault(void* ptr);
char Message_EventDefault(void* ptr, void* e);
char Message_EventFilterDefault(void* ptr, void* watched, void* event);
void Message_ChildEventDefault(void* ptr, void* event);
void Message_ConnectNotifyDefault(void* ptr, void* sign);
void Message_CustomEventDefault(void* ptr, void* event);
void Message_DeleteLaterDefault(void* ptr);
void Message_DisconnectNotifyDefault(void* ptr, void* sign);
void Message_TimerEventDefault(void* ptr, void* event);
;

#ifdef __cplusplus
}
#endif

#endif