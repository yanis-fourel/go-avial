package av

func (e Entity) SetProperty(name string, key string, value Value, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodSet,
		Aspect:        AspectProperty,
		Name:          name,
		Key:           key,
		Value:         value,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) GetProperty(key string, auth Token) (Value, error) {
	return e.Invoke(ArgsInvokeEntity{
		Method:        MethodGet,
		Aspect:        AspectProperty,
		Key:           key,
		Authorization: auth,
	})
}

func (e Entity) ClearProperty(key string, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodClear,
		Aspect:        AspectProperty,
		Key:           key,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) PropertyCount(auth Token) (uint64, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodCount,
		Aspect:        AspectProperty,
		Authorization: auth,
	})
	if err != nil {
		return 0, err
	}
	res, err := val.UnwrapInteger()
	return uint64(res), err
}

func (e Entity) InsertProperty(name string, key string, value Value, index uint64, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodInsert,
		Aspect:        AspectProperty,
		Name:          name,
		Key:           key,
		Value:         value,
		Index:         index,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) RemoveProperty(index uint64, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodRemove,
		Aspect:        AspectProperty,
		Index:         index,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) ReplaceProperty(name string, key string, value Value, index uint64, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodReplace,
		Aspect:        AspectProperty,
		Name:          name,
		Key:           key,
		Value:         value,
		Index:         index,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) FindProperty(name string, value Value, index uint64, auth Token) (uint64, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodFind,
		Aspect:        AspectProperty,
		Name:          name,
		Value:         value,
		Index:         index,
		Authorization: auth,
	})
	if err != nil {
		return 0, err
	}
	res, err := val.UnwrapInteger()
	return uint64(res), err
}

func (e Entity) IncludeProperty(name string, key string, value Value, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodInclude,
		Aspect:        AspectProperty,
		Name:          name,
		Key:           key,
		Value:         value,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) ExpludeProperty(key string, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodExclude,
		Aspect:        AspectProperty,
		Key:           key,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) PropertyMember(key string, auth Token) (bool, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodMember,
		Aspect:        AspectProperty,
		Key:           key,
		Authorization: auth,
	})
	if err != nil {
		return false, err
	}
	return val.UnwrapBoolean(), nil
}

func (e Entity) PropertyName(key string, index uint64, auth Token) (string, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodName,
		Aspect:        AspectProperty,
		Key:           key,
		Index:         index,
		Authorization: auth,
	})
	if err != nil {
		return "", err
	}
	return val.UnwrapString(), nil
}

func (e Entity) PropertyKey(index uint64, auth Token) (string, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodKey,
		Aspect:        AspectProperty,
		Index:         index,
		Authorization: auth,
	})
	if err != nil {
		return "", err
	}
	return val.UnwrapString(), nil
}

func (e Entity) PropertyValue(key string, index uint64, auth Token) (Value, error) {
	return e.Invoke(ArgsInvokeEntity{
		Method:        MethodValue,
		Aspect:        AspectProperty,
		Key:           key,
		Index:         index,
		Authorization: auth,
	})
}

func (e Entity) PropertyIndex(key string, auth Token) (uint64, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodIndex,
		Aspect:        AspectProperty,
		Key:           key,
		Authorization: auth,
	})
	if err != nil {
		return 0, err
	}
	res, err := val.UnwrapInteger()
	return uint64(res), err
}

func (e Entity) PropertyAttribute(key string, index uint64, auth Token) (Attribute, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodAttribute,
		Aspect:        AspectProperty,
		Key:           key,
		Index:         index,
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

func (e Entity) SortProperties(parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodSort,
		Aspect:        AspectProperty,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) PurgeProperties(parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodPurge,
		Aspect:        AspectProperty,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) RetrieveProperties(auth Token) (string, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodRetrieve,
		Aspect:        AspectProperty,
		Authorization: auth,
	})
	if err != nil {
		return "", err
	}
	return val.UnwrapInterchange(), nil
	// TODO: Test the unmarshalling
	// var res []Property
	// err = json.Unmarshal([]byte(val.UnwrapString()), &res)
}
