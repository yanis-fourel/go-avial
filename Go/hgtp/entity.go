package hgtp

import (
	"fmt"
)

// The most fundamental concept in AvesTerra is the entity.
// An entity represent *a thing*, its semantic representation is defined by
// the application.
// Each entity is identified by a unique entity identifier (EUID), which is
// what this `Entity` type represents.
type Entity struct {
	// Public Identifier: Identifies the public network where the AvesTerra
	// server that is managing the entity resides
	Pid uint32

	// Host Identifier: Identifies the specific host within the public network
	// The Combinaison of Pid and Hid uniquely identifies the AvesTerra server
	Hid uint32

	// Unique Identifier: Identifies the entity itself on the designated host
	Uid uint64
}

func NewEntity(pid uint32, hid uint32, uid uint64) Entity {
	return Entity{
		Pid: pid,
		Hid: hid,
		Uid: uid,
	}
}

// For textual representation, an EUID is displayed as three decimal number
// groups separated by the '|' (pipe) character enclosed in angle brackets in
// the form <pid|hid|uid>. Several of the AvesTerra tools allow the shorthand
// forms <hid|uid> and <uid> where the missing pid and hid fields are defaulted
// to that of the current server.
// All three of "<pid|hid|uid>", "<hid|uid>", and "<uid>" are valid and
// and correctly parsed by this function.
func ParseEntity(str string) (Entity, error) {
	var entity Entity

	_, err := fmt.Sscanf(str, "<%d>", &entity.Uid)
	if err == nil {
		return entity, nil
	}

	_, err = fmt.Sscanf(str, "<%d|%d>", &entity.Uid, &entity.Uid)
	if err == nil {
		return entity, nil
	}

	_, err = fmt.Sscanf(str, "<%d|%d|%d>", &entity.Pid, &entity.Hid, &entity.Uid)
	if err == nil {
		return entity, nil
	}

	return Entity{}, fmt.Errorf("invalid entity: '%s'", str)
}

func (e Entity) String() string {
	return fmt.Sprintf("<%d|%d|%d>", e.Pid, e.Hid, e.Uid)
}
