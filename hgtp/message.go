package hgtp

import (
	"encoding/binary"
	"fmt"
	"time"
	"unicode/utf8"
)

// Most low-level possible representation of an HGTP message while still
// abstracting away how it is send/received over the network
type hgtpMessage struct {
	Frame     [1024]byte
	Unbounded []byte
}

func (msg *hgtpMessage) DebugString() string {
	bytes := msg.UnpackBytes()
	payload_string := string(bytes)
	payload_hex := fmt.Sprintf("% x", bytes)

	return "" +
		fmt.Sprintf("%2d | %2d  [command, error]\n", uint16(msg.UnpackCommand()), uint16(msg.UnpackError())) +
		fmt.Sprintf("%2d | %2d  [version, report]\n", msg.UnpackVersion(), msg.UnpackReport()) +
		fmt.Sprintf("%2d | %2d  [method, attribute]\n", msg.UnpackMethod(), msg.UnpackAttribute()) +
		fmt.Sprintf("%2d | %2d  [event, mode]\n", msg.UnpackEvent(), msg.UnpackMode()) +
		fmt.Sprintf("%2d | %2d  [category, class]\n", msg.UnpackCategory(), msg.UnpackClass()) +
		fmt.Sprintf("%2d | %2d  [context, aspect]\n", msg.UnpackContext(), msg.UnpackAspect()) +
		fmt.Sprintf("%2d | %2d  [states, precedence]\n", msg.UnpackState(), msg.UnpackPrecedence()) +
		fmt.Sprintf("%2d | %2d  [tag, condition]\n", msg.UnpackTag(), msg.UnpackCondition()) +
		fmt.Sprintf("%2d | %2d  [instance, offset]\n", msg.UnpackInstance(), msg.UnpackOffset()) +
		fmt.Sprintf("%2d  (%s) [time]\n", msg.UnpackTime().Unix(), msg.UnpackTime().UTC()) +
		fmt.Sprintf("%2d | %2d  [index, count]\n", msg.UnpackIndex(), msg.UnpackCount()) +
		fmt.Sprintf("%2d | %2d  [extension, parameter]\n", msg.UnpackExtension(), msg.UnpackParameter()) +
		fmt.Sprintf("%2d | %2d  [resultant, timeout]\n", msg.UnpackResultant(), msg.UnpackTimeout()) +
		fmt.Sprintf("%s [authorization]\n", msg.UnpackAuthorization()) +
		fmt.Sprintf("%s [authority]\n", msg.UnpackAuthority()) +
		fmt.Sprintf("%s [entity]\n", msg.UnpackEntity()) +
		fmt.Sprintf("%s [outlet]\n", msg.UnpackOutlet()) +
		fmt.Sprintf("%s [auxiliary]\n", msg.UnpackAuxiliary()) +
		fmt.Sprintf("%s [ancillary]\n", msg.UnpackAncillary()) +
		fmt.Sprintf("'%s' [name]\n", msg.UnpackName()) +
		fmt.Sprintf("'%s' [key]\n", msg.UnpackKey()) +
		fmt.Sprintf("bytes / extension as hex (%d bytes):\n", len(msg.UnpackBytes())) +
		payload_hex + "\n" +
		fmt.Sprintf("bytes / extension as string (%d utf8-encoded glyphs):\n", utf8.RuneCountInString(payload_string)) +
		"'" + payload_string + "'\n"
}

func (msg *hgtpMessage) Validate() error {
	errcode := msg.UnpackError()
	if errcode != ErrorOk {
		return &AvesTerraError{
			Code:    errcode,
			Message: string(msg.UnpackBytes()),
		}
	} else {
		return nil
	}
}

//
// -------------------------------- Unpackers ---------------------------------
//

func (msg *hgtpMessage) UnpackCommand() Command {
	return Command(binary.LittleEndian.Uint16(msg.Frame[0:2]))
}

func (msg *hgtpMessage) UnpackError() Error {
	return Error(binary.LittleEndian.Uint16(msg.Frame[2:4]))
}

func (msg *hgtpMessage) UnpackVersion() uint16 {
	return binary.LittleEndian.Uint16(msg.Frame[4:6])
}

func (msg *hgtpMessage) UnpackReport() Report {
	return Report(binary.LittleEndian.Uint16(msg.Frame[6:8]))
}

func (msg *hgtpMessage) UnpackMethod() Method {

	return Method(binary.LittleEndian.Uint16(msg.Frame[8:10]))
}

func (msg *hgtpMessage) UnpackAttribute() Attribute {
	return Attribute(binary.LittleEndian.Uint16(msg.Frame[10:12]))
}

func (msg *hgtpMessage) UnpackEvent() Event {
	return Event(binary.LittleEndian.Uint16(msg.Frame[12:14]))
}

func (msg *hgtpMessage) UnpackMode() Mode {
	return Mode(binary.LittleEndian.Uint16(msg.Frame[14:16]))
}

func (msg *hgtpMessage) UnpackCategory() Category {
	return Category(binary.LittleEndian.Uint16(msg.Frame[16:18]))
}

func (msg *hgtpMessage) UnpackClass() Class {
	return Class(binary.LittleEndian.Uint16(msg.Frame[18:20]))
}

func (msg *hgtpMessage) UnpackContext() Context {
	return Context(binary.LittleEndian.Uint16(msg.Frame[20:22]))
}

func (msg *hgtpMessage) UnpackAspect() Aspect {
	return Aspect(binary.LittleEndian.Uint16(msg.Frame[22:24]))
}

func (msg *hgtpMessage) UnpackState() State {
	return State(binary.LittleEndian.Uint16(msg.Frame[24:26]))
}

func (msg *hgtpMessage) UnpackPrecedence() Precedence {
	return Precedence(binary.LittleEndian.Uint16(msg.Frame[26:28]))
}

func (msg *hgtpMessage) UnpackTag() Tag {
	return Tag(binary.LittleEndian.Uint16(msg.Frame[28:30]))
}

func (msg *hgtpMessage) UnpackCondition() Condition {
	return Condition(binary.LittleEndian.Uint16(msg.Frame[30:32]))
}

func (msg *hgtpMessage) UnpackInstance() uint32 {
	return binary.LittleEndian.Uint32(msg.Frame[32:36])
}

func (msg *hgtpMessage) UnpackOffset() uint32 {
	return binary.LittleEndian.Uint32(msg.Frame[36:40])
}

func (msg *hgtpMessage) UnpackTime() time.Time {
	return time.Unix(int64(binary.LittleEndian.Uint64(msg.Frame[40:48])), 0)
}

func (msg *hgtpMessage) UnpackIndex() uint64 {
	return binary.LittleEndian.Uint64(msg.Frame[48:56])
}

func (msg *hgtpMessage) UnpackCount() uint64 {
	return binary.LittleEndian.Uint64(msg.Frame[56:64])
}

func (msg *hgtpMessage) UnpackExtension() uint64 {
	return binary.LittleEndian.Uint64(msg.Frame[64:72])
}

func (msg *hgtpMessage) UnpackParameter() int64 {
	return int64(binary.LittleEndian.Uint64(msg.Frame[72:80]))
}

func (msg *hgtpMessage) UnpackResultant() int64 {
	return int64(binary.LittleEndian.Uint64(msg.Frame[80:88]))
}

func (msg *hgtpMessage) UnpackTimeout() Timeout {
	return Timeout(binary.LittleEndian.Uint64(msg.Frame[88:96]))
}

// -- 64 reserved bytes --

func (msg *hgtpMessage) UnpackEntity() Entity {
	return Entity{
		Pid: binary.LittleEndian.Uint32(msg.Frame[160:164]),
		Hid: binary.LittleEndian.Uint32(msg.Frame[164:168]),
		Uid: binary.LittleEndian.Uint64(msg.Frame[168:176]),
	}
}

func (msg *hgtpMessage) UnpackOutlet() Entity {
	return Entity{
		Pid: binary.LittleEndian.Uint32(msg.Frame[176:180]),
		Hid: binary.LittleEndian.Uint32(msg.Frame[180:184]),
		Uid: binary.LittleEndian.Uint64(msg.Frame[184:192]),
	}
}

func (msg *hgtpMessage) UnpackAuxiliary() Entity {
	return Entity{
		Pid: binary.LittleEndian.Uint32(msg.Frame[192:196]),
		Hid: binary.LittleEndian.Uint32(msg.Frame[196:200]),
		Uid: binary.LittleEndian.Uint64(msg.Frame[200:208]),
	}
}

func (msg *hgtpMessage) UnpackAncillary() Entity {
	return Entity{
		Pid: binary.LittleEndian.Uint32(msg.Frame[208:212]),
		Hid: binary.LittleEndian.Uint32(msg.Frame[212:216]),
		Uid: binary.LittleEndian.Uint64(msg.Frame[216:224]),
	}
}

func (msg *hgtpMessage) UnpackAuthorization() Token {
	var authorization Token
	copy(authorization[:], msg.Frame[224:240])
	return authorization
}

func (msg *hgtpMessage) UnpackAuthority() Token {
	var authority Token
	copy(authority[:], msg.Frame[240:256])
	return authority
}

func (msg *hgtpMessage) UnpackName() string {
	name_length := msg.Frame[256]
	name := string(msg.Frame[257 : 257+int(name_length)])
	return name
}

func (msg *hgtpMessage) UnpackKey() string {
	key_length := msg.Frame[512]
	key := string(msg.Frame[513 : 513+int(key_length)])
	return key
}

func (msg *hgtpMessage) UnpackBytes() []byte {
	if msg.UnpackExtension() > 0 {
		return msg.Unbounded
	} else {
		bytes_length := msg.Frame[768]
		return msg.Frame[769 : 769+int(bytes_length)]
	}
}

//
// --------------------------------- Packers ----------------------------------
//

func (msg *hgtpMessage) PackCommand(command Command) {
	binary.LittleEndian.PutUint16(msg.Frame[0:2], uint16(command))
}

func (msg *hgtpMessage) PackError(error Error) {
	binary.LittleEndian.PutUint16(msg.Frame[2:4], uint16(error))
}

func (msg *hgtpMessage) PackVersion(version uint16) {
	binary.LittleEndian.PutUint16(msg.Frame[4:6], version)
}

func (msg *hgtpMessage) PackReport(report Report) {
	binary.LittleEndian.PutUint16(msg.Frame[6:8], uint16(report))
}

func (msg *hgtpMessage) PackMethod(method Method) {
	binary.LittleEndian.PutUint16(msg.Frame[8:10], uint16(method))
}

func (msg *hgtpMessage) PackAttribute(attribute Attribute) {
	binary.LittleEndian.PutUint16(msg.Frame[10:12], uint16(attribute))
}

func (msg *hgtpMessage) PackEvent(event Event) {
	binary.LittleEndian.PutUint16(msg.Frame[12:14], uint16(event))
}

func (msg *hgtpMessage) PackMode(mode Mode) {
	binary.LittleEndian.PutUint16(msg.Frame[14:16], uint16(mode))
}

func (msg *hgtpMessage) PackCategory(category Category) {
	binary.LittleEndian.PutUint16(msg.Frame[16:18], uint16(category))
}

func (msg *hgtpMessage) PackClass(class Class) {
	binary.LittleEndian.PutUint16(msg.Frame[18:20], uint16(class))
}

func (msg *hgtpMessage) PackContext(context Context) {
	binary.LittleEndian.PutUint16(msg.Frame[20:22], uint16(context))
}

func (msg *hgtpMessage) PackAspect(aspect Aspect) {
	binary.LittleEndian.PutUint16(msg.Frame[22:24], uint16(aspect))
}

func (msg *hgtpMessage) PackState(state State) {
	binary.LittleEndian.PutUint16(msg.Frame[24:26], uint16(state))
}

func (msg *hgtpMessage) PackPrecedence(precedence Precedence) {
	binary.LittleEndian.PutUint16(msg.Frame[26:28], uint16(precedence))
}

func (msg *hgtpMessage) PackTag(tag Tag) {
	binary.LittleEndian.PutUint16(msg.Frame[28:30], uint16(tag))
}

func (msg *hgtpMessage) PackCondition(condition Condition) {
	binary.LittleEndian.PutUint16(msg.Frame[30:32], uint16(condition))
}

func (msg *hgtpMessage) PackInstance(instance uint32) {
	binary.LittleEndian.PutUint32(msg.Frame[32:36], instance)
}

func (msg *hgtpMessage) PackOffset(offset uint32) {
	binary.LittleEndian.PutUint32(msg.Frame[36:40], offset)
}

func (msg *hgtpMessage) PackTime(t time.Time) {
	// `time.Time` default value is Jan 1, year 1. We want to default to Unix epoch.
	if t.IsZero() {
		t = time.Unix(0, 0)
	}
	binary.LittleEndian.PutUint64(msg.Frame[40:48], uint64(t.Unix()))
}

func (msg *hgtpMessage) PackIndex(index uint64) {
	binary.LittleEndian.PutUint64(msg.Frame[48:56], index)
}

func (msg *hgtpMessage) PackCount(count uint64) {
	binary.LittleEndian.PutUint64(msg.Frame[56:64], count)
}

// Internally managed by `PackBytes`
func (msg *hgtpMessage) packExtension(extension uint64) {
	binary.LittleEndian.PutUint64(msg.Frame[64:72], extension)
}

func (msg *hgtpMessage) PackParameter(parameter int64) {
	binary.LittleEndian.PutUint64(msg.Frame[72:80], uint64(parameter))
}

func (msg *hgtpMessage) PackResultant(resultant int64) {
	binary.LittleEndian.PutUint64(msg.Frame[80:88], uint64(resultant))
}

func (msg *hgtpMessage) PackTimeout(timeout Timeout) {
	binary.LittleEndian.PutUint64(msg.Frame[88:96], uint64(timeout))
}

// -- 64 reserved bytes --

func (msg *hgtpMessage) PackEntity(entity Entity) {
	binary.LittleEndian.PutUint32(msg.Frame[160:164], entity.Pid)
	binary.LittleEndian.PutUint32(msg.Frame[164:168], entity.Hid)
	binary.LittleEndian.PutUint64(msg.Frame[168:176], entity.Uid)
}

func (msg *hgtpMessage) PackOutlet(outlet Entity) {
	binary.LittleEndian.PutUint32(msg.Frame[176:180], outlet.Pid)
	binary.LittleEndian.PutUint32(msg.Frame[180:184], outlet.Hid)
	binary.LittleEndian.PutUint64(msg.Frame[184:192], outlet.Uid)
}

func (msg *hgtpMessage) PackAuxiliary(auxiliary Entity) {
	binary.LittleEndian.PutUint32(msg.Frame[192:196], auxiliary.Pid)
	binary.LittleEndian.PutUint32(msg.Frame[196:200], auxiliary.Hid)
	binary.LittleEndian.PutUint64(msg.Frame[200:208], auxiliary.Uid)
}

func (msg *hgtpMessage) PackAncillary(ancillary Entity) {
	binary.LittleEndian.PutUint32(msg.Frame[208:212], ancillary.Pid)
	binary.LittleEndian.PutUint32(msg.Frame[212:216], ancillary.Hid)
	binary.LittleEndian.PutUint64(msg.Frame[216:224], ancillary.Uid)
}

func (msg *hgtpMessage) PackAuthorization(authorization Token) {
	copy(msg.Frame[224:240], authorization[:])
}

func (msg *hgtpMessage) PackAuthority(authority Token) {
	copy(msg.Frame[240:256], authority[:])
}

func (msg *hgtpMessage) PackName(name string) error {
	if len(name) > 255 {
		return fmt.Errorf("Invalid name: Length exceeds maximum length of 255 : %d", len(name))
	}
	msg.Frame[256] = byte(len(name))
	copy(msg.Frame[257:257+len(name)], name)
	return nil
}

func (msg *hgtpMessage) PackKey(key string) error {
	if len(key) > 255 {
		return fmt.Errorf("Invalid key: Length exceeds maximum length of 255 : %d", len(key))
	}
	msg.Frame[512] = byte(len(key))
	copy(msg.Frame[513:513+len(key)], key)
	return nil
}

func (msg *hgtpMessage) PackBytes(bytes []byte) {
	if len(bytes) > 255 {
		if msg.Frame[768] > 0 {
			for i := 0; i < int(msg.Frame[768]); i++ {
				msg.Frame[769+i] = 0
			}
			msg.Frame[768] = 0
		}
		msg.packExtension(uint64(len(bytes)))
		msg.Unbounded = bytes
	} else {
		msg.Frame[768] = byte(len(bytes))
		copy(msg.Frame[769:769+len(bytes)], bytes)
		msg.Unbounded = nil
		msg.packExtension(0)
	}
}
