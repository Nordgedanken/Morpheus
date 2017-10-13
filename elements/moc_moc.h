/****************************************************************************
** Meta object code from reading C++ file 'moc.cpp'
**
** Created by: The Qt Meta Object Compiler version 67 (Qt 5.9.1)
**
** WARNING! All changes made in this file will be lost!
*****************************************************************************/

#include <QtCore/qbytearray.h>
#include <QtCore/qmetatype.h>
#if !defined(Q_MOC_OUTPUT_REVISION)
#error "The header file 'moc.cpp' doesn't include <QObject>."
#elif Q_MOC_OUTPUT_REVISION != 67
#error "This file was generated using the moc from 5.9.1. It"
#error "cannot be used with the include files from this version of Qt."
#error "(The moc has changed too much.)"
#endif

QT_BEGIN_MOC_NAMESPACE
QT_WARNING_PUSH
QT_WARNING_DISABLE_DEPRECATED
struct qt_meta_stringdata_QVBoxLayoutWithTriggerSlot_t {
    QByteArrayData data[5];
    char stringdata0[62];
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
QT_MOC_LITERAL(4, 55, 6) // "sender"

    },
    "QVBoxLayoutWithTriggerSlot\0TriggerMessage\0"
    "\0messageBody\0sender"
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
       0,       // signalCount

 // slots: name, argc, parameters, tag, flags
       1,    2,   19,    2, 0x0a /* Public */,

 // slots: parameters
    QMetaType::Void, QMetaType::QString, QMetaType::QString,    3,    4,

       0        // eod
};

void QVBoxLayoutWithTriggerSlot::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        QVBoxLayoutWithTriggerSlot *_t = static_cast<QVBoxLayoutWithTriggerSlot *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->TriggerMessage((*reinterpret_cast< QString(*)>(_a[1])),(*reinterpret_cast< QString(*)>(_a[2]))); break;
        default: ;
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
        return static_cast<void*>(const_cast< QVBoxLayoutWithTriggerSlot*>(this));
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
QT_WARNING_POP
QT_END_MOC_NAMESPACE
