package av

func (e Entity) SetFactor(attribute Attribute, name string, key string, value Value, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodSet,
		Aspect:        AspectFactor,
		Attribute:     attribute,
		Name:          name,
		Key:           key,
		Value:         value,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) GetFactor(attribute Attribute, name string, key string, auth Token) (Value, error) {
	return e.Invoke(ArgsInvokeEntity{
		Method:        MethodGet,
		Aspect:        AspectFactor,
		Attribute:     attribute,
		Name:          name,
		Key:           key,
		Authorization: auth,
	})
}

func (e Entity) ClearFactor(attribute Attribute, name string, key string, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodClear,
		Aspect:        AspectFactor,
		Attribute:     attribute,
		Name:          name,
		Key:           key,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) FactorCount(attribute Attribute, name string, instance uint32, offset uint32, auth Token) (uint64, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodCount,
		Aspect:        AspectFactor,
		Attribute:     attribute,
		Name:          name,
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

func (e Entity) InsertFactor(attribute Attribute, name string, key string, value Value, index uint64, instance uint32, offset uint32, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodInsert,
		Aspect:        AspectFactor,
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

func (e Entity) RemoveFactor(attribute Attribute, name string, index uint64, instance uint32, offset uint32, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodRemove,
		Aspect:        AspectFactor,
		Attribute:     attribute,
		Name:          name,
		Index:         index,
		Instance:      instance,
		Offset:        offset,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) ReplaceFactor(attribute Attribute, name string, key string, value Value, index uint64, instance uint32, offset uint32, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodReplace,
		Aspect:        AspectFactor,
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

func (e Entity) FindFactor(attribute Attribute, name string, value Value, index uint64, instance uint32, offset uint32, auth Token) (uint64, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodFind,
		Aspect:        AspectFactor,
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

func (e Entity) IncludeFactor(attribute Attribute, name string, key string, value Value, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodInclude,
		Aspect:        AspectFactor,
		Attribute:     attribute,
		Name:          name,
		Key:           key,
		Value:         value,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) ExcludeFactor(attribute Attribute, name string, key string, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodExclude,
		Aspect:        AspectFactor,
		Attribute:     attribute,
		Name:          name,
		Key:           key,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) FactorMember(attribute Attribute, name string, key string, instance uint32, offset uint32, auth Token) (bool, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodMember,
		Aspect:        AspectFactor,
		Attribute:     attribute,
		Name:          name,
		Key:           key,
		Instance:      instance,
		Offset:        offset,
		Authorization: auth,
	})
	if err != nil {
		return false, err
	}
	return val.UnwrapBoolean(), nil
}

func (e Entity) FactorName(attribute Attribute, name string, index uint64, instance uint32, offset uint32, auth Token) (string, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodName,
		Aspect:        AspectFactor,
		Attribute:     attribute,
		Name:          name,
		Instance:      instance,
		Index:         index,
		Authorization: auth,
	})
	if err != nil {
		return "", err
	}
	return val.UnwrapString(), nil
}

func (e Entity) FactorKey(attribute Attribute, name string, index uint64, instance uint32, offset uint32, auth Token) (string, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodKey,
		Aspect:        AspectFactor,
		Attribute:     attribute,
		Name:          name,
		Index:         index,
		Instance:      instance,
		Offset:        offset,
		Authorization: auth,
	})
	if err != nil {
		return "", err
	}
	return val.UnwrapString(), nil
}

func (e Entity) FactorValue(attribute Attribute, name string, key string, index uint64, instance uint32, offset uint32, auth Token) (Value, error) {
	return e.Invoke(ArgsInvokeEntity{
		Method:        MethodValue,
		Aspect:        AspectFactor,
		Attribute:     attribute,
		Name:          name,
		Key:           key,
		Index:         index,
		Instance:      instance,
		Offset:        offset,
		Authorization: auth,
	})
}

func (e Entity) FactorIndex(attribute Attribute, name string, key string, instance uint32, offset uint32, auth Token) (uint64, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodIndex,
		Aspect:        AspectFactor,
		Attribute:     attribute,
		Name:          name,
		Key:           key,
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

func (e Entity) FactorAttribute(name string, key string, index uint64, instance uint32, offset uint32, auth Token) (Attribute, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodAttribute,
		Aspect:        AspectFactor,
		Name:          name,
		Key:           key,
		Index:         index,
		Instance:      instance,
		Offset:        offset,
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

func (e Entity) SortFactors(attribute Attribute, name string, instance uint32, offset uint32, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodSort,
		Aspect:        AspectFactor,
		Attribute:     attribute,
		Name:          name,
		Instance:      instance,
		Offset:        offset,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) PurgeFactors(attribute Attribute, name string, instance uint32, offset uint32, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodPurge,
		Aspect:        AspectFactor,
		Attribute:     attribute,
		Name:          name,
		Instance:      instance,
		Offset:        offset,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) RetrieveFactors(attribute Attribute, name string, instance uint32, offset uint32, parameter Parameter, auth Token) (string, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodRetrieve,
		Aspect:        AspectFactor,
		Attribute:     attribute,
		Name:          name,
		Instance:      instance,
		Offset:        offset,
		Authorization: auth,
	})
	if err != nil {
		return "", err
	}
	return val.UnwrapInterchange(), nil
	// TODO: Test the unmarshalling
	// var res []Factor
	// err = json.Unmarshal([]byte(val.UnwrapString()), &res)
}
