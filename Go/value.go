package av

import (
	"encoding/json"
	"fmt"
	"github.com/yanis-fourel/go-avial/Go/hgtp"
	"strconv"
	"time"
)

type Value struct {
	Tag   Tag
	Bytes []byte
}

type Variable struct {
	Name  string
	Value Value
}

func ValueAvesterra(val string) Value {
	return Value{Tag: TagAvesterra, Bytes: []byte(val)}
}

func ValueBoolean(val bool) Value {
	if val {
		return Value{Tag: TagBoolean, Bytes: []byte("TRUE")}
	}
	return Value{Tag: TagBoolean, Bytes: []byte("FALSE")}
}

func ValueCharacter(val byte) Value {
	return Value{Tag: TagCharacter, Bytes: []byte{val}}
}

// ascii string
func ValueString(val string) Value {
	return Value{Tag: TagString, Bytes: []byte(val)}
}

// utf-8 encoded string
func ValueText(val string) Value {
	return Value{Tag: TagText, Bytes: []byte(val)}
}

func ValueInteger(val int64) Value {
	str := strconv.FormatInt(val, 10)
	return Value{Tag: TagInteger, Bytes: []byte(str)}
}

func ValueFloat(val float64) Value {
	str := strconv.FormatFloat(val, 'E', 17, 64)
	return Value{Tag: TagFloat, Bytes: []byte(str)}
}

func ValueEntity(val Entity) Value {
	return Value{Tag: TagEntity, Bytes: []byte(val.String())}
}

func ValueTime(val time.Time) Value {
	// `time.Time` default value is Jan 1, year 1. We want to default to Unix epoch.
	if val.IsZero() {
		val = time.Unix(0, 0)
	}
	str := strconv.FormatInt(val.Unix(), 10)
	return Value{Tag: TagTime, Bytes: []byte(str)}
}

func ValueWeb(val string) Value {
	return Value{Tag: TagWeb, Bytes: []byte(val)}
}

func ValueInterchange(val string) Value {
	return Value{Tag: TagInterchange, Bytes: []byte(val)}
}

func ValueData(val []byte) Value {
	return Value{Tag: TagData, Bytes: val}
}

func ValueException(val AvesTerraError) Value {
	tmp := struct {
		ERROR   string `json:"ERROR"`
		MESSAGE string `json:"MESSAGE"`
	}{
		ERROR:   val.Code.String(),
		MESSAGE: val.Message,
	}
	b, err := json.Marshal(tmp)
	if err != nil {
		panic(err)
	}
	return Value{Tag: TagException, Bytes: b}
}

func ValueOperator(val Entity) Value {
	return Value{Tag: TagOperator, Bytes: []byte(val.String())}
}

func ValueFunction(val Entity) Value {
	return Value{Tag: TagFunction, Bytes: []byte(val.String())}
}

func ValueMeasurement(val interface{}) Value {
	panic("TODO: implement, need the measurement type")
	// return Value{Tag: TagMeasurement, Val: []byte(val)}
}

func ValueLocutor(val interface{}) Value {
	panic("TODO: implement, need the locutor type")
	// return Value{Tag: TagLocutor, Val: []byte(val)}
}

func ValueAuthorization(val Token) Value {
	var b []byte
	err := val.UnmarshalBinary(b)
	if err != nil {
		panic(err)
	}

	return Value{Tag: TagAuthorization, Bytes: b}
}

func ValueDuration(val float64) Value {
	str := strconv.FormatFloat(val, 'f', 9, 64)
	return Value{Tag: TagDuration, Bytes: []byte(str)}
}

func ValueVariable(val Variable) Value {
	panic("TODO: implement")
}

func ValueAggregate(val []Variable) Value {
	panic("TODO: implement")
}

func ValueArray(val []Value) Value {
	panic("TODO: implement")
}

//
// ---
//

// Get string value from a String, Text, Web, Interchange or AvesTerra value
func (v *Value) GetString() (string, error) {
	switch v.Tag {
	case TagAvesterra | TagString | TagText | TagWeb | TagInterchange:
		return string(v.Bytes), nil
	default:
		return "", fmt.Errorf("Invalid tag: %s is not readable as a string", v.Tag.String())
	}
}

// Panics if tag is not Avesterra
func (v *Value) UnwrapAvesterra() string {
	if v.Tag != TagAvesterra {
		panic("Invalid tag: Expected value to be Avesterra but it is " + v.Tag.String())
	}
	return string(v.Bytes)
}

// Panics if tag is not Boolean
func (v *Value) UnwrapBoolean() bool {
	if v.Tag != TagBoolean {
		panic("Invalid tag: Expected value to be Boolean but it is " + v.Tag.String())
	}
	return string(v.Bytes) == "TRUE"
}

// Panics if tag is not Character
func (v *Value) UnwrapCharacter() byte {
	if v.Tag != TagCharacter {
		panic("Invalid tag: Expected value to be Character but it is " + v.Tag.String())
	}
	return v.Bytes[0]
}

// Panics if tag is not String
func (v *Value) UnwrapString() string {
	if v.Tag != TagString {
		panic("Invalid tag: Expected value to be String but it is " + v.Tag.String())
	}
	return string(v.Bytes)
}

// Panics if tag is not Text
func (v *Value) UnwrapText() string {
	if v.Tag != TagText {
		panic("Invalid tag: Expected value to be Text but it is " + v.Tag.String())
	}
	return string(v.Bytes)
}

// Panics if tag is not Integer
func (v *Value) UnwrapInteger() (int64, error) {
	if v.Tag != TagInteger {
		panic("Invalid tag: Expected value to be Integer but it is " + v.Tag.String())
	}
	return strconv.ParseInt(string(v.Bytes), 10, 0)
}

// Panics if tag is not Float
func (v *Value) UnwrapFloat() (float64, error) {
	if v.Tag != TagFloat {
		panic("Invalid tag: Expected value to be Float but it is " + v.Tag.String())
	}
	return strconv.ParseFloat(string(v.Bytes), 64)
}

// Panics if tag is not Entity
func (v *Value) UnwrapEntity() (Entity, error) {
	if v.Tag != TagEntity {
		panic("Invalid tag: Expected value to be Entity but it is " + v.Tag.String())
	}
	return ParseEntity(string(v.Bytes))
}

// Panics if tag is not Time
func (v *Value) UnwrapTime() (time.Time, error) {
	if v.Tag != TagTime {
		panic("Invalid tag: Expected value to be Time but it is " + v.Tag.String())
	}
	i, err := strconv.ParseInt(string(v.Bytes), 10, 64)
	return time.Unix(i, 0), err
}

// Panics if tag is not Web
func (v *Value) UnwrapWeb() string {
	if v.Tag != TagWeb {
		panic("Invalid tag: Expected value to be Web but it is " + v.Tag.String())
	}
	return string(v.Bytes)
}

// Panics if tag is not Interchange
func (v *Value) UnwrapInterchange() string {
	if v.Tag != TagInterchange {
		panic("Invalid tag: Expected value to be Interchange but it is " + v.Tag.String())
	}
	return string(v.Bytes)
}

// Panics if tag is not Data
func (v *Value) UnwrapData() []byte {
	if v.Tag != TagData {
		panic("Invalid tag: Expected value to be Data but it is " + v.Tag.String())
	}
	return v.Bytes
}

// Panics if tag is not Exception
func (v *Value) UnwrapException() (AvesTerraError, error) {
	if v.Tag != TagException {
		panic("Invalid tag: Expected value to be Exception but it is " + v.Tag.String())
	}

	res := struct {
		ERROR   hgtp.Error `json:"ERROR"`
		MESSAGE string     `json:"MESSAGE"`
	}{}
	err := json.Unmarshal(v.Bytes, &res)
	return AvesTerraError{
		Code:    res.ERROR,
		Message: res.MESSAGE,
	}, err
}

// Panics if tag is not Operator
func (v *Value) UnwrapOperator() (Entity, error) {
	if v.Tag != TagOperator {
		panic("Invalid tag: Expected value to be Operator but it is " + v.Tag.String())
	}
	return ParseEntity(string(v.Bytes))
}

// Panics if tag is not Function
func (v *Value) UnwrapFunction() (Entity, error) {

	// Panics if tag is not Tag
	if v.Tag != TagFunction {

		// Panics if tag is not c("
		panic("Invalid tag: Expected value to be Function but it is " + v.Tag.String())
	}
	return ParseEntity(string(v.Bytes))
}

// Panics if tag is not Measurement
func (v *Value) UnwrapMeasurement() (Measurement, error) {
	if v.Tag != TagMeasurement {
		panic("Invalid tag: Expected value to be Measurement but it is " + v.Tag.String())
	}
	res := Measurement{}
	err := json.Unmarshal(v.Bytes, &res)
	return res, err
}

// Panics if tag is not Locutor
func (v *Value) UnwrapLocutor() (Locutor, error) {
	if v.Tag != TagLocutor {
		panic("Invalid tag: Expected value to be Locutor but it is " + v.Tag.String())
	}
	res := Locutor{}
	err := json.Unmarshal(v.Bytes, &res)
	return res, err
}

// Panics if tag is not Authorization
func (v *Value) UnwrapAuthorization() (Token, error) {
	if v.Tag != TagAuthorization {
		panic("Invalid tag: Expected value to be Authorization but it is " + v.Tag.String())
	}
	res := Token{}
	err := json.Unmarshal(v.Bytes, &res)
	return res, err
}

// Panics if tag is not Duration
func (v *Value) UnwrapDuration() (float64, error) {
	if v.Tag != TagDuration {
		panic("Invalid tag: Expected value to be Duration but it is " + v.Tag.String())
	}
	return strconv.ParseFloat(string(v.Bytes), 64)
}

// Panics if tag is not Variable
func (v *Value) UnwrapVariable() (Variable, error) {
	if v.Tag != TagVariable {
		panic("Invalid tag: Expected value to be Variable but it is " + v.Tag.String())
	}
	panic("TODO: implement")
}

// Panics if tag is not Aggregate
func (v *Value) UnwrapAggregate() ([]Variable, error) {
	if v.Tag != TagAggregate {
		panic("Invalid tag: Expected value to be Aggregate but it is " + v.Tag.String())
	}
	panic("TODO: implement")
}

// Panics if tag is not Array
func (v *Value) UnwrapArray() ([]Value, error) {
	if v.Tag != TagArray {
		panic("Invalid tag: Expected value to be Array but it is " + v.Tag.String())
	}
	panic("TODO: implement")
}

func (v Value) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		v.Tag.String(): string(v.Bytes),
	})
}

func (v *Value) UnmarshalJSON(data []byte) error {
	var res map[string]string
	err := json.Unmarshal(data, &res)
	if err != nil {
		return err
	}
	if len(res) != 1 {
		return fmt.Errorf("Invalid Value JSON: more than one tag-value pair")
	}

	for key, val := range res {
		t, ok := MapStringToTag[key]
		if !ok {
			return fmt.Errorf("Invalid Value JSON: unknown tag '%s'", key)
		}

		var b []byte
		if len(val) > 0 {
			b = []byte(val)
		}

		*v = Value{
			Tag:   t,
			Bytes: b,
		}
	}
	return nil
}
