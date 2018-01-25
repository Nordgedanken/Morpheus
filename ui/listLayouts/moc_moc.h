/****************************************************************************
** Meta object code from reading C++ file 'moc.cpp'
**
** Created by: The Qt Meta Object Compiler version 67 (Qt 5.9.3)
**
** WARNING! All changes made in this file will be lost!
*****************************************************************************/

#include <QtCore/qbytearray.h>
#include <QtCore/qmetatype.h>
#if !defined(Q_MOC_OUTPUT_REVISION)
#error "The header file 'moc.cpp' doesn't include <QObject>."
#elif Q_MOC_OUTPUT_REVISION != 67
#error "This file was generated using the moc from 5.9.3. It"
#error "cannot be used with the include files from this version of Qt."
#error "(The moc has changed too much.)"
#endif

QT_BEGIN_MOC_NAMESPACE
QT_WARNING_PUSH
QT_WARNING_DISABLE_DEPRECATED
struct qt_meta_stringdata_QRoomVBoxLayoutWithTriggerSlot_t {
    QByteArrayData data[5];
    char stringdata0[62];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_QRoomVBoxLayoutWithTriggerSlot_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_QRoomVBoxLayoutWithTriggerSlot_t qt_meta_stringdata_QRoomVBoxLayoutWithTriggerSlot = {
    {
QT_MOC_LITERAL(0, 0, 30), // "QRoomVBoxLayoutWithTriggerSlot"
QT_MOC_LITERAL(1, 31, 11), // "TriggerRoom"
QT_MOC_LITERAL(2, 43, 0), // ""
QT_MOC_LITERAL(3, 44, 6), // "roomID"
QT_MOC_LITERAL(4, 51, 10) // "ChangeRoom"

    },
    "QRoomVBoxLayoutWithTriggerSlot\0"
    "TriggerRoom\0\0roomID\0ChangeRoom"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_QRoomVBoxLayoutWithTriggerSlot[] = {

 // content:
       7,       // revision
       0,       // classname
       0,    0, // classinfo
       2,   14, // methods
       0,    0, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       2,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   24,    2, 0x06 /* Public */,
       4,    1,   27,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void, QMetaType::QString,    3,
    QMetaType::Void, QMetaType::QString,    3,

       0        // eod
};

void QRoomVBoxLayoutWithTriggerSlot::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        QRoomVBoxLayoutWithTriggerSlot *_t = static_cast<QRoomVBoxLayoutWithTriggerSlot *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->TriggerRoom((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 1: _t->ChangeRoom((*reinterpret_cast< QString(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            typedef void (QRoomVBoxLayoutWithTriggerSlot::*_t)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&QRoomVBoxLayoutWithTriggerSlot::TriggerRoom)) {
                *result = 0;
                return;
            }
        }
        {
            typedef void (QRoomVBoxLayoutWithTriggerSlot::*_t)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&QRoomVBoxLayoutWithTriggerSlot::ChangeRoom)) {
                *result = 1;
                return;
            }
        }
    }
}

const QMetaObject QRoomVBoxLayoutWithTriggerSlot::staticMetaObject = {
    { &QVBoxLayout::staticMetaObject, qt_meta_stringdata_QRoomVBoxLayoutWithTriggerSlot.data,
      qt_meta_data_QRoomVBoxLayoutWithTriggerSlot,  qt_static_metacall, nullptr, nullptr}
};


const QMetaObject *QRoomVBoxLayoutWithTriggerSlot::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *QRoomVBoxLayoutWithTriggerSlot::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_QRoomVBoxLayoutWithTriggerSlot.stringdata0))
        return static_cast<void*>(this);
    return QVBoxLayout::qt_metacast(_clname);
}

int QRoomVBoxLayoutWithTriggerSlot::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QVBoxLayout::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 2)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 2;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 2)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 2;
    }
    return _id;
}

// SIGNAL 0
void QRoomVBoxLayoutWithTriggerSlot::TriggerRoom(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void QRoomVBoxLayoutWithTriggerSlot::ChangeRoom(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}
struct qt_meta_stringdata_QVBoxLayoutWithTriggerSlot_t {
    QByteArrayData data[6];
    char stringdata0[72];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_QVBoxLayoutWithTriggerSlot_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_QVBoxLayoutWithTriggerSlot_t qt_meta_stringdata_QVBoxLayoutWithTriggerSlot = {
    {
QT_MOC_LITERAL(0, 0, 26), // "QVBoxLayoutWithTriggerSlot"
QT_MOC_LITERAL(1, 27, 14), // "TriggerMessage"
QT_MOC_LITERAL(2, 42, 0), // ""
QT_MOC_LITERAL(3, 43, 11), // "messageBody"
QT_MOC_LITERAL(4, 55, 6), // "sender"
QT_MOC_LITERAL(5, 62, 9) // "timestamp"

    },
    "QVBoxLayoutWithTriggerSlot\0TriggerMessage\0"
    "\0messageBody\0sender\0timestamp"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_QVBoxLayoutWithTriggerSlot[] = {

 // content:
       7,       // revision
       0,       // classname
       0,    0, // classinfo
       1,   14, // methods
       0,    0, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       1,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    3,   19,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void, QMetaType::QString, QMetaType::QString, QMetaType::LongLong,    3,    4,    5,

       0        // eod
};

void QVBoxLayoutWithTriggerSlot::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        QVBoxLayoutWithTriggerSlot *_t = static_cast<QVBoxLayoutWithTriggerSlot *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->TriggerMessage((*reinterpret_cast< QString(*)>(_a[1])),(*reinterpret_cast< QString(*)>(_a[2])),(*reinterpret_cast< qint64(*)>(_a[3]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            typedef void (QVBoxLayoutWithTriggerSlot::*_t)(QString , QString , qint64 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&QVBoxLayoutWithTriggerSlot::TriggerMessage)) {
                *result = 0;
                return;
            }
        }
    }
}

const QMetaObject QVBoxLayoutWithTriggerSlot::staticMetaObject = {
    { &QVBoxLayout::staticMetaObject, qt_meta_stringdata_QVBoxLayoutWithTriggerSlot.data,
      qt_meta_data_QVBoxLayoutWithTriggerSlot,  qt_static_metacall, nullptr, nullptr}
};


const QMetaObject *QVBoxLayoutWithTriggerSlot::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *QVBoxLayoutWithTriggerSlot::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_QVBoxLayoutWithTriggerSlot.stringdata0))
        return static_cast<void*>(this);
    return QVBoxLayout::qt_metacast(_clname);
}

int QVBoxLayoutWithTriggerSlot::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QVBoxLayout::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 1)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 1;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 1)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 1;
    }
    return _id;
}

// SIGNAL 0
void QVBoxLayoutWithTriggerSlot::TriggerMessage(QString _t1, QString _t2, qint64 _t3)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)), const_cast<void*>(reinterpret_cast<const void*>(&_t2)), const_cast<void*>(reinterpret_cast<const void*>(&_t3)) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}
struct qt_meta_stringdata_QRoomLayout_t {
    QByteArrayData data[4];
    char stringdata0[30];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_QRoomLayout_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_QRoomLayout_t qt_meta_stringdata_QRoomLayout = {
    {
QT_MOC_LITERAL(0, 0, 11), // "QRoomLayout"
QT_MOC_LITERAL(1, 12, 9), // "SetAvatar"
QT_MOC_LITERAL(2, 22, 0), // ""
QT_MOC_LITERAL(3, 23, 6) // "roomID"

    },
    "QRoomLayout\0SetAvatar\0\0roomID"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_QRoomLayout[] = {

 // content:
       7,       // revision
       0,       // classname
       0,    0, // classinfo
       1,   14, // methods
       0,    0, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       1,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   19,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void, QMetaType::QString,    3,

       0        // eod
};

void QRoomLayout::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        QRoomLayout *_t = static_cast<QRoomLayout *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->SetAvatar((*reinterpret_cast< QString(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            typedef void (QRoomLayout::*_t)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&QRoomLayout::SetAvatar)) {
                *result = 0;
                return;
            }
        }
    }
}

const QMetaObject QRoomLayout::staticMetaObject = {
    { &QHBoxLayout::staticMetaObject, qt_meta_stringdata_QRoomLayout.data,
      qt_meta_data_QRoomLayout,  qt_static_metacall, nullptr, nullptr}
};


const QMetaObject *QRoomLayout::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *QRoomLayout::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_QRoomLayout.stringdata0))
        return static_cast<void*>(this);
    return QHBoxLayout::qt_metacast(_clname);
}

int QRoomLayout::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QHBoxLayout::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 1)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 1;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 1)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 1;
    }
    return _id;
}

// SIGNAL 0
void QRoomLayout::SetAvatar(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(&_t1)) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}
QT_WARNING_POP
QT_END_MOC_NAMESPACE
