

#pragma once

#ifndef GO_MOC_H
#define GO_MOC_H

#include <stdint.h>

#ifdef __cplusplus
class Room;
void Room_Room_QRegisterMetaTypes();
extern "C" {
#endif

struct Moc_PackedString { char* data; long long len; };
struct Moc_PackedList { void* data; long long len; };
void Room_ConnectSetAvatar(void* ptr);
void Room_DisconnectSetAvatar(void* ptr);
void Room_SetAvatar(void* ptr, uintptr_t IMGdata);
int Room_Room_QRegisterMetaType();
int Room_Room_QRegisterMetaType2(char* typeName);
int Room_Room_QmlRegisterType();
int Room_Room_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName);
void* Room___dynamicPropertyNames_atList(void* ptr, int i);
void Room___dynamicPropertyNames_setList(void* ptr, void* i);
void* Room___dynamicPropertyNames_newList(void* ptr);
void* Room___findChildren_atList2(void* ptr, int i);
void Room___findChildren_setList2(void* ptr, void* i);
void* Room___findChildren_newList2(void* ptr);
void* Room___findChildren_atList3(void* ptr, int i);
void Room___findChildren_setList3(void* ptr, void* i);
void* Room___findChildren_newList3(void* ptr);
void* Room___findChildren_atList(void* ptr, int i);
void Room___findChildren_setList(void* ptr, void* i);
void* Room___findChildren_newList(void* ptr);
void* Room___children_atList(void* ptr, int i);
void Room___children_setList(void* ptr, void* i);
void* Room___children_newList(void* ptr);
void* Room_NewRoom(void* parent);
void Room_DestroyRoom(void* ptr);
void Room_DestroyRoomDefault(void* ptr);
char Room_EventDefault(void* ptr, void* e);
char Room_EventFilterDefault(void* ptr, void* watched, void* event);
void Room_ChildEventDefault(void* ptr, void* event);
void Room_ConnectNotifyDefault(void* ptr, void* sign);
void Room_CustomEventDefault(void* ptr, void* event);
void Room_DeleteLaterDefault(void* ptr);
void Room_DisconnectNotifyDefault(void* ptr, void* sign);
void Room_TimerEventDefault(void* ptr, void* event);
;

#ifdef __cplusplus
}
#endif

#endif