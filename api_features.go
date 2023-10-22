package av

func (e Entity) SetFeature(attribute Attribute, name string, key string, value Value, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodSet,
		Aspect:        AspectFeature,
		Attribute:     attribute,
		Name:          name,
		Key:           key,
		Value:         value,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) GetFeature(attribute Attribute, name string, key string, auth Token) (Value, error) {
	return e.Invoke(ArgsInvokeEntity{
		Method:        MethodGet,
		Aspect:        AspectFeature,
		Attribute:     attribute,
		Name:          name,
		Key:           key,
		Authorization: auth,
	})
}

func (e Entity) ClearFeature(attribute Attribute, name string, key string, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodClear,
		Aspect:        AspectFeature,
		Attribute:     attribute,
		Name:          name,
		Key:           key,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) FeatureCount(attribute Attribute, name string, instance uint32, auth Token) (uint64, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodCount,
		Aspect:        AspectFeature,
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

func (e Entity) InsertFeature(attribute Attribute, name string, key string, value Value, index uint64, instance uint32, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodInsert,
		Aspect:        AspectFeature,
		Attribute:     attribute,
		Name:          name,
		Key:           key,
		Value:         value,
		Index:         index,
		Instance:      instance,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) RemoveFeature(attribute Attribute, name string, index uint64, instance uint32, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodRemove,
		Aspect:        AspectFeature,
		Attribute:     attribute,
		Name:          name,
		Index:         index,
		Instance:      instance,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) ReplaceFeature(attribute Attribute, name string, key string, value Value, index uint64, instance uint32, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodReplace,
		Aspect:        AspectFeature,
		Attribute:     attribute,
		Name:          name,
		Key:           key,
		Value:         value,
		Index:         index,
		Instance:      instance,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) FindFeature(attribute Attribute, name string, value Value, index uint64, instance uint32, auth Token) (uint64, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodFind,
		Aspect:        AspectFeature,
		Attribute:     attribute,
		Name:          name,
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

func (e Entity) IncludeFeature(attribute Attribute, name string, key string, value Value, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodInclude,
		Aspect:        AspectFeature,
		Attribute:     attribute,
		Name:          name,
		Key:           key,
		Value:         value,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) ExcludeFeature(attribute Attribute, name string, key string, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodExclude,
		Aspect:        AspectFeature,
		Attribute:     attribute,
		Name:          name,
		Key:           key,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) FeatureMember(attribute Attribute, name string, key string, instance uint32, auth Token) (bool, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodMember,
		Aspect:        AspectFeature,
		Attribute:     attribute,
		Name:          name,
		Key:           key,
		Instance:      instance,
		Authorization: auth,
	})
	if err != nil {
		return false, err
	}
	return val.UnwrapBoolean(), nil
}

func (e Entity) FeatureName(attribute Attribute, name string, index uint64, instance uint32, auth Token) (string, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodName,
		Aspect:        AspectFeature,
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

func (e Entity) FeatureKey(attribute Attribute, name string, index uint64, instance uint32, auth Token) (string, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodKey,
		Aspect:        AspectFeature,
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

func (e Entity) FeatureValue(attribute Attribute, name string, key string, index uint64, instance uint32, auth Token) (Value, error) {
	return e.Invoke(ArgsInvokeEntity{
		Method:        MethodValue,
		Aspect:        AspectFeature,
		Attribute:     attribute,
		Name:          name,
		Key:           key,
		Index:         index,
		Instance:      instance,
		Authorization: auth,
	})
}

func (e Entity) FeatureIndex(attribute Attribute, name string, key string, instance uint32, auth Token) (uint64, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodIndex,
		Aspect:        AspectFeature,
		Attribute:     attribute,
		Name:          name,
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

func (e Entity) FeatureAttribute(name string, key string, index uint64, instance uint32, auth Token) (Attribute, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodAttribute,
		Aspect:        AspectFeature,
		Name:          name,
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

func (e Entity) SortFeatures(attribute Attribute, name string, instance uint32, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodSort,
		Aspect:        AspectFeature,
		Attribute:     attribute,
		Name:          name,
		Instance:      instance,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) PurgeFeatures(attribute Attribute, name string, instance uint32, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodPurge,
		Aspect:        AspectFeature,
		Attribute:     attribute,
		Name:          name,
		Instance:      instance,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) RetrieveFeatures(attribute Attribute, name string, instance uint32, parameter Parameter, auth Token) (string, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodRetrieve,
		Aspect:        AspectFeature,
		Attribute:     attribute,
		Name:          name,
		Instance:      instance,
		Authorization: auth,
	})
	if err != nil {
		return "", err
	}
	return val.UnwrapInterchange(), nil
	// TODO: Test the unmarshalling
	// var res []Feature
	// err = json.Unmarshal([]byte(val.UnwrapString()), &res)
}
