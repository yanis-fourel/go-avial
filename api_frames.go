package av

func (e Entity) SetFrame(attribute Attribute, name string, key string, value Value, index uint64, instance uint32, offset uint32, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodSet,
		Aspect:        AspectFrame,
		Attribute:     attribute,
		Name:          name,
		Key:           key,
		Value:         value,
		Index:         index,
		Instance:      instance,
		Offset:        offset,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) GetFrame(attribute Attribute, name string, key string, index uint64, instance uint32, offset uint32, auth Token) (Value, error) {
	return e.Invoke(ArgsInvokeEntity{
		Method:        MethodGet,
		Aspect:        AspectFrame,
		Attribute:     attribute,
		Name:          name,
		Key:           key,
		Index:         index,
		Instance:      instance,
		Offset:        offset,
		Authorization: auth,
	})
}

func (e Entity) ClearFrame(attribute Attribute, name string, key string, index uint64, instance uint32, offset uint32, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodClear,
		Aspect:        AspectFrame,
		Attribute:     attribute,
		Name:          name,
		Key:           key,
		Index:         index,
		Instance:      instance,
		Offset:        offset,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) FrameCount(attribute Attribute, instance uint32, auth Token) (uint64, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodCount,
		Aspect:        AspectFrame,
		Attribute:     attribute,
		Instance:      instance,
		Authorization: auth,
	})
	if err != nil {
		return 0, err
	}
	res, err := val.UnwrapInteger()
	return uint64(res), err
}

func (e Entity) InsertFrame(attribute Attribute, key string, index uint64, instance uint32, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodInsert,
		Aspect:        AspectFrame,
		Attribute:     attribute,
		Key:           key,
		Index:         index,
		Instance:      instance,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) RemoveFrame(attribute Attribute, index uint64, instance uint32, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodRemove,
		Aspect:        AspectFrame,
		Attribute:     attribute,
		Index:         index,
		Instance:      instance,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) ReplaceFrame(attribute Attribute, key string, index uint64, instance uint32, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodReplace,
		Aspect:        AspectFrame,
		Attribute:     attribute,
		Key:           key,
		Index:         index,
		Instance:      instance,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) FindFrame(attribute Attribute, name string, value Value, index uint64, instance uint32, offset uint32, auth Token) (uint64, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodFind,
		Aspect:        AspectFrame,
		Attribute:     attribute,
		Name:          name,
		Value:         value,
		Index:         index,
		Instance:      instance,
		Offset:        offset,
		Authorization: auth,
	})
	if err != nil {
		return 0, err
	}
	res, err := val.UnwrapInteger()
	return uint64(res), err
}

func (e Entity) IncludeFrame(attribute Attribute, key string, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodInclude,
		Aspect:        AspectFrame,
		Attribute:     attribute,
		Key:           key,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) ExcludeFrame(attribute Attribute, key string, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodExclude,
		Aspect:        AspectFrame,
		Attribute:     attribute,
		Key:           key,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) FrameMember(attribute Attribute, key string, instance uint32, auth Token) (bool, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodMember,
		Aspect:        AspectFrame,
		Attribute:     attribute,
		Key:           key,
		Instance:      instance,
		Authorization: auth,
	})
	if err != nil {
		return false, err
	}
	return val.UnwrapBoolean(), nil
}

func (e Entity) FrameName(attribute Attribute, instance uint32, offset uint32, auth Token) (string, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodName,
		Aspect:        AspectFrame,
		Attribute:     attribute,
		Instance:      instance,
		Authorization: auth,
	})
	if err != nil {
		return "", err
	}
	return val.UnwrapString(), nil
}

func (e Entity) FrameKey(attribute Attribute, index uint64, instance uint32, auth Token) (string, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodKey,
		Aspect:        AspectFrame,
		Attribute:     attribute,
		Index:         index,
		Instance:      instance,
		Authorization: auth,
	})
	if err != nil {
		return "", err
	}
	return val.UnwrapString(), nil
}

func (e Entity) FrameValue(attribute Attribute, name string, key string, index uint64, instance uint32, offset uint32, auth Token) (Value, error) {
	return e.Invoke(ArgsInvokeEntity{
		Method:        MethodValue,
		Aspect:        AspectFrame,
		Attribute:     attribute,
		Name:          name,
		Key:           key,
		Index:         index,
		Instance:      instance,
		Offset:        offset,
		Authorization: auth,
	})
}

func (e Entity) FrameIndex(attribute Attribute, key string, instance uint32, auth Token) (uint64, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodIndex,
		Aspect:        AspectFrame,
		Attribute:     attribute,
		Key:           key,
		Instance:      instance,
		Authorization: auth,
	})
	if err != nil {
		return 0, err
	}
	res, err := val.UnwrapInteger()
	return uint64(res), err
}

func (e Entity) FrameAttribute(key string, index uint64, instance uint32, auth Token) (Attribute, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodAttribute,
		Aspect:        AspectFrame,
		Key:           key,
		Index:         index,
		Instance:      instance,
		Authorization: auth,
	})
	if err != nil {
		return AttributeNull, err
	}
	res, err := val.UnwrapInteger()
	if err != nil {
		return AttributeNull, err
	}
	return Attribute(res), nil
}

func (e Entity) SortFrames(attribute Attribute, instance uint32, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodSort,
		Aspect:        AspectFrame,
		Attribute:     attribute,
		Instance:      instance,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) PurgeFrames(attribute Attribute, instance uint32, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodPurge,
		Aspect:        AspectFrame,
		Attribute:     attribute,
		Instance:      instance,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) RetrieveFrames(attribute Attribute, instance uint32, parameter Parameter, auth Token) (string, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodRetrieve,
		Aspect:        AspectFrame,
		Attribute:     attribute,
		Instance:      instance,
		Authorization: auth,
	})
	if err != nil {
		return "", err
	}
	return val.UnwrapInterchange(), nil
	// TODO: Test the unmarshalling
	// var res []Frame
	// err = json.Unmarshal([]byte(val.UnwrapString()), &res)
}
