package av

func (e Entity) SetField(attribute Attribute, name string, value Value, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodSet,
		Aspect:        AspectField,
		Attribute:     attribute,
		Name:          name,
		Value:         value,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) GetField(attribute Attribute, name string, auth Token) (Value, error) {
	return e.Invoke(ArgsInvokeEntity{
		Method:        MethodGet,
		Aspect:        AspectField,
		Attribute:     attribute,
		Name:          name,
		Authorization: auth,
	})
}

func (e Entity) ClearField(attribute Attribute, name string, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodClear,
		Aspect:        AspectField,
		Attribute:     attribute,
		Name:          name,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) FieldCount(attribute Attribute, instance uint32, auth Token) (uint64, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodCount,
		Aspect:        AspectField,
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

func (e Entity) InsertField(attribute Attribute, name string, value Value, index uint64, instance uint32, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodInsert,
		Aspect:        AspectField,
		Attribute:     attribute,
		Name:          name,
		Value:         value,
		Index:         index,
		Instance:      instance,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) RemoveField(attribute Attribute, index uint64, instance uint32, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodRemove,
		Aspect:        AspectField,
		Attribute:     attribute,
		Index:         index,
		Instance:      instance,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) ReplaceField(attribute Attribute, name string, value Value, index uint64, instance uint32, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodReplace,
		Aspect:        AspectField,
		Attribute:     attribute,
		Name:          name,
		Value:         value,
		Index:         index,
		Instance:      instance,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) FindField(attribute Attribute, value Value, index uint64, instance uint32, auth Token) (uint64, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodFind,
		Aspect:        AspectField,
		Attribute:     attribute,
		Value:         value,
		Index:         index,
		Instance:      instance,
		Authorization: auth,
	})
	if err != nil {
		return 0, err
	}
	res, err := val.UnwrapInteger()
	return uint64(res), err
}

func (e Entity) IncludeField(attribute Attribute, name string, value Value, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodInclude,
		Aspect:        AspectField,
		Attribute:     attribute,
		Name:          name,
		Value:         value,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) ExcludeField(attribute Attribute, name string, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodExclude,
		Aspect:        AspectField,
		Attribute:     attribute,
		Name:          name,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) FieldMember(attribute Attribute, name string, instance uint32, auth Token) (bool, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodMember,
		Aspect:        AspectField,
		Attribute:     attribute,
		Name:          name,
		Instance:      instance,
		Authorization: auth,
	})
	if err != nil {
		return false, err
	}
	return val.UnwrapBoolean(), nil
}

func (e Entity) FieldName(attribute Attribute, index uint64, instance uint32, auth Token) (string, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodName,
		Aspect:        AspectField,
		Attribute:     attribute,
		Instance:      instance,
		Index:         index,
		Authorization: auth,
	})
	if err != nil {
		return "", err
	}
	return val.UnwrapString(), nil
}

func (e Entity) FieldKey(attribute Attribute, name string, index uint64, instance uint32, auth Token) (string, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodKey,
		Aspect:        AspectField,
		Attribute:     attribute,
		Name:          name,
		Index:         index,
		Instance:      instance,
		Authorization: auth,
	})
	if err != nil {
		return "", err
	}
	return val.UnwrapString(), nil
}

func (e Entity) FieldValue(attribute Attribute, name string, index uint64, instance uint32, auth Token) (Value, error) {
	return e.Invoke(ArgsInvokeEntity{
		Method:        MethodValue,
		Aspect:        AspectField,
		Attribute:     attribute,
		Name:          name,
		Index:         index,
		Instance:      instance,
		Authorization: auth,
	})
}

func (e Entity) FieldIndex(attribute Attribute, name string, instance uint32, auth Token) (uint64, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodIndex,
		Aspect:        AspectField,
		Attribute:     attribute,
		Name:          name,
		Instance:      instance,
		Authorization: auth,
	})
	if err != nil {
		return 0, err
	}
	res, err := val.UnwrapInteger()
	return uint64(res), err
}

func (e Entity) FieldAttribute(name string, index uint64, instance uint32, auth Token) (Attribute, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodAttribute,
		Aspect:        AspectField,
		Name:          name,
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

func (e Entity) SortFields(attribute Attribute, instance uint32, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodSort,
		Aspect:        AspectField,
		Attribute:     attribute,
		Instance:      instance,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) PurgeFields(attribute Attribute, instance uint32, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodPurge,
		Aspect:        AspectField,
		Attribute:     attribute,
		Instance:      instance,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) RetrieveFields(attribute Attribute, instance uint32, auth Token) (string, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodRetrieve,
		Aspect:        AspectField,
		Attribute:     attribute,
		Instance:      instance,
		Authorization: auth,
	})
	if err != nil {
		return "", err
	}
	return val.UnwrapInterchange(), nil
	// TODO: Test the unmarshalling
	// var res []Field
	// err = json.Unmarshal([]byte(val.UnwrapString()), &res)
}
