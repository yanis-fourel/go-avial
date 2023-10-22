package av

import (
	"encoding/json"
	"fmt"
)

type AvialModel struct {
	Name string
	Key  string

	// Length of the data in bytes
	Data uint64

	Properties []Property
	Facts      []Fact
}

func (i *AvialModel) UnmashalJSON(data []byte) error {
	type Alias AvialModel
	var tmp Alias
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	if len(tmp.Facts) == 0 {
		tmp.Facts = nil
	}
	if len(tmp.Properties) == 0 {
		tmp.Properties = nil
	}
	*i = AvialModel(tmp)
	return nil
}

//
// Properties
//

type Property struct {
	Name        string
	Key         string
	Value       Value
	Annotations AnnotationMap
}

type AnnotationMap map[Attribute]Value

func (p *Property) UnmarshalJSON(data []byte) error {
	var arr []json.RawMessage
	err := json.Unmarshal(data, &arr)
	if err != nil {
		return err
	}

	if len(arr) != 3 && len(arr) != 4 {
		return fmt.Errorf("Invalid Property JSON: expected array of length 3 or 4, got %d", len(arr))
	}

	if err = json.Unmarshal(arr[0], &p.Name); err != nil {
		return err
	}
	if err = json.Unmarshal(arr[1], &p.Key); err != nil {
		return err
	}
	if err = json.Unmarshal(arr[2], &p.Value); err != nil {
		return err
	}

	if len(arr) == 4 {
		if err = json.Unmarshal(arr[3], &p.Annotations); err != nil {
			return err
		}
	}

	if err = p.validate(); err != nil {
		return err
	}

	return nil
}

func (p Property) MarshalJSON() ([]byte, error) {
	arr := []interface{}{p.Name, p.Key, p.Value}
	if len(p.Annotations) > 0 {
		arr = append(arr, p.Annotations)
	}
	return json.Marshal(arr)
}

func (p *Property) validate() error {
	if len(p.Name) > 255 {
		return fmt.Errorf("Invalid Property: Name too long (max 255 bytes)")
	}
	if len(p.Key) > 255 {
		return fmt.Errorf("Invalid Property: Key too long (max 255 bytes)")
	}
	return nil
}

//
// Facts
//

type Fact struct {
	Attribute Attribute
	Value     Value

	Facets   []Facet
	Features []Feature
	Fields   []Field
	Frames   []Frame
}

func (f *Fact) UnmarshalJSON(data []byte) error {
	var arr []json.RawMessage

	err := json.Unmarshal(data, &arr)
	if err != nil {
		return err
	}
	if len(arr) != 6 {
		return fmt.Errorf("Invalid Fact JSON: expected array of length 6 (attribute,value,facets,features,fields,frames), got %d", len(arr))
	}

	if err = json.Unmarshal(arr[0], &f.Attribute); err != nil {
		return err
	}

	if err = json.Unmarshal(arr[1], &f.Value); err != nil {
		return err
	}

	if err = json.Unmarshal(arr[2], &f.Facets); err != nil {
		return err
	}
	if len(f.Facets) == 0 {
		f.Facets = nil
	}

	if err = json.Unmarshal(arr[3], &f.Features); err != nil {
		return err
	}
	if len(f.Features) == 0 {
		f.Features = nil
	}

	if err = json.Unmarshal(arr[4], &f.Fields); err != nil {
		return err
	}
	if len(f.Fields) == 0 {
		f.Fields = nil
	}

	if err = json.Unmarshal(arr[5], &f.Frames); err != nil {
		return err
	}
	if len(f.Frames) == 0 {
		f.Frames = nil
	}

	return f.validate()
}

func (f *Fact) MarshalJSON() ([]byte, error) {
	arr := []interface{}{
		f.Attribute,
		f.Value,
		f.Facets,
		f.Features,
		f.Fields,
		f.Frames,
	}
	return json.Marshal(arr)
}

func (f *Fact) validate() error {
	if len(f.Fields) != len(f.Frames) {
		return fmt.Errorf("Invalid Fact: Different length for Fields (%d) and Frames (%d)", len(f.Fields), len(f.Frames))
	}
	return nil
}

type Facet struct {
	Name    string
	Value   Value
	Factors FactorList
}

func (f *Facet) UnmarshalJSON(data []byte) error {
	var arr []json.RawMessage
	err := json.Unmarshal(data, &arr)
	if err != nil {
		return err
	}

	if len(arr) != 3 {
		return fmt.Errorf("Invalid Facet JSON: expected array of length 3 (name, value, factors), got %d", len(arr))
	}

	if err = json.Unmarshal(arr[0], &f.Name); err != nil {
		return err
	}
	if err = json.Unmarshal(arr[1], &f.Value); err != nil {
		return err
	}
	if err = json.Unmarshal(arr[2], &f.Factors); err != nil {
		return err
	}

	if err = f.validate(); err != nil {
		return err
	}

	return nil
}

func (f Facet) MarshalJSON() ([]byte, error) {
	arr := []interface{}{f.Name, f.Value, f.Factors}
	return json.Marshal(arr)
}

func (f *Facet) validate() error {
	if len(f.Name) > 255 {
		return fmt.Errorf("Invalid Facet: Name too long (max 255 bytes)")
	}
	return nil
}

type FactorList []Factor

type Factor struct {
	Key   string
	Value Value
}

func (f *Factor) UnmarshalJSON(data []byte) error {
	var arr []json.RawMessage
	err := json.Unmarshal(data, &arr)
	if err != nil {
		return err
	}

	if len(arr) != 2 {
		return fmt.Errorf("Invalid Factor JSON: expected array of length 2 (key, value), got %d", len(arr))
	}

	if err = json.Unmarshal(arr[0], &f.Key); err != nil {
		return err
	}
	if err = json.Unmarshal(arr[1], &f.Value); err != nil {
		return err
	}

	if err = f.validate(); err != nil {
		return err
	}
	return nil
}

func (f Factor) MarshalJSON() ([]byte, error) {
	arr := []interface{}{f.Key, f.Value}
	return json.Marshal(arr)
}

func (f *Factor) validate() error {
	if len(f.Key) > 255 {
		return fmt.Errorf("Invalid Factor: Key too long (max 255 bytes)")
	}
	return nil
}

type Feature struct {
	Name  string
	Key   string
	Value Value
}

func (f *Feature) UnmarshalJSON(data []byte) error {
	var arr []json.RawMessage
	err := json.Unmarshal(data, &arr)
	if err != nil {
		return err
	}

	if len(arr) != 3 {
		return fmt.Errorf("Invalid Factor JSON: expected array of length 3 (name, key, value), got %d", len(arr))
	}

	if err = json.Unmarshal(arr[0], &f.Name); err != nil {
		return err
	}
	if err = json.Unmarshal(arr[1], &f.Key); err != nil {
		return err
	}
	if err = json.Unmarshal(arr[2], &f.Value); err != nil {
		return err
	}

	if err = f.validate(); err != nil {
		return err
	}
	return nil
}

func (f Feature) MarshalJSON() ([]byte, error) {
	arr := []interface{}{f.Name, f.Key, f.Value}
	return json.Marshal(arr)
}

func (f *Feature) validate() error {
	if len(f.Name) > 255 {
		return fmt.Errorf("Invalid Feature: Name too long (max 255 bytes)")
	}
	if len(f.Key) > 255 {
		return fmt.Errorf("Invalid Feature: Key too long (max 255 bytes)")
	}
	return nil
}

type Field struct {
	Name         string
	DefaultValue Value
}

func (f *Field) UnmarshalJSON(data []byte) error {
	var arr []json.RawMessage
	err := json.Unmarshal(data, &arr)
	if err != nil {
		return err
	}

	if len(arr) != 2 {
		return fmt.Errorf("Invalid Field JSON: expected array of length 2 (name, defaultvalue), got %d", len(arr))
	}

	if err = json.Unmarshal(arr[0], &f.Name); err != nil {
		return err
	}
	if err = json.Unmarshal(arr[1], &f.DefaultValue); err != nil {
		return err
	}

	if err = f.validate(); err != nil {
		return err
	}

	return nil
}

func (f Field) MarshalJSON() ([]byte, error) {
	arr := []interface{}{f.Name, f.DefaultValue}
	return json.Marshal(arr)
}

func (f *Field) validate() error {
	if len(f.Name) > 255 {
		return fmt.Errorf("Invalid Field: Name too long (max 255 bytes)")
	}
	return nil
}

type Frame struct {
	Key    string
	Values []Value
}

func (f *Frame) UnmarshalJSON(data []byte) error {
	var arr []json.RawMessage
	err := json.Unmarshal(data, &arr)
	if err != nil {
		return err
	}

	if len(arr) != 2 {
		return fmt.Errorf("Invalid Frame JSON: expected array of length 2 (key, values), got %d", len(arr))
	}

	if err = json.Unmarshal(arr[0], &f.Key); err != nil {
		return err
	}
	if err = json.Unmarshal(arr[1], &f.Values); err != nil {
		return err
	}

	if err = f.validate(); err != nil {
		return err
	}
	return nil
}

func (f Frame) MarshalJSON() ([]byte, error) {
	arr := []interface{}{f.Key, f.Values}
	return json.Marshal(arr)
}

func (f *Frame) validate() error {
	if len(f.Key) > 255 {
		return fmt.Errorf("Invalid Frame: Key too long (max 255 bytes)")
	}
	return nil
}
