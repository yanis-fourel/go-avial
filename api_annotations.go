package av

func (e Entity) SetAnnotation(attribute Attribute, key string, value Value, parameter Parameter, auth Token) {
	e.Invoke(ArgsInvokeEntity{
		Method:        MethodSet,
		Aspect:        AspectAnnotation,
		Attribute:     attribute,
		Key:           key,
		Value:         value,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
}

func (e Entity) GetAnnotation(attribute Attribute, key string, auth Token) (Value, error) {
	return e.Invoke(ArgsInvokeEntity{
		Method:        MethodGet,
		Aspect:        AspectAnnotation,
		Attribute:     attribute,
		Key:           key,
		Authorization: auth,
	})
}

func (e Entity) ClearAnnotation(attribute Attribute, key string, parameter Parameter, auth Token) {
	e.Invoke(ArgsInvokeEntity{
		Method:        MethodClear,
		Aspect:        AspectAnnotation,
		Attribute:     attribute,
		Key:           key,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
}

func (e Entity) AnnotationCount(key string, instance uint32, auth Token) (uint64, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodCount,
		Aspect:        AspectAnnotation,
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

func (e Entity) InsertAnnotation(attribute Attribute, key string, value Value, index uint64, instance uint32, parameter Parameter, auth Token) {
	e.Invoke(ArgsInvokeEntity{
		Method:        MethodInsert,
		Aspect:        AspectAnnotation,
		Attribute:     attribute,
		Key:           key,
		Value:         value,
		Index:         index,
		Instance:      instance,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
}

func (e Entity) RemoveAnnotation(key string, index uint64, instance uint32, parameter Parameter, auth Token) {
	e.Invoke(ArgsInvokeEntity{
		Method:        MethodRemove,
		Aspect:        AspectAnnotation,
		Key:           key,
		Index:         index,
		Instance:      instance,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
}

func (e Entity) ReplaceAnnotation(attribute Attribute, key string, value Value, index uint64, instance uint32, parameter Parameter, auth Token) {
	e.Invoke(ArgsInvokeEntity{
		Method:        MethodReplace,
		Aspect:        AspectAnnotation,
		Attribute:     attribute,
		Key:           key,
		Value:         value,
		Index:         index,
		Instance:      instance,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
}

func (e Entity) FindAnnotation(key string, value Value, index uint64, instance uint32, auth Token) (uint64, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodFind,
		Aspect:        AspectAnnotation,
		Key:           key,
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

func (e Entity) IncludeAnnotation(attribute Attribute, key string, value Value, parameter Parameter, auth Token) {
	e.Invoke(ArgsInvokeEntity{
		Method:        MethodInclude,
		Aspect:        AspectAnnotation,
		Attribute:     attribute,
		Key:           key,
		Value:         value,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
}

func (e Entity) ExcludeAnnotation(attribute Attribute, key string, parameter Parameter, auth Token) {
	e.Invoke(ArgsInvokeEntity{
		Method:        MethodExclude,
		Aspect:        AspectAnnotation,
		Attribute:     attribute,
		Key:           key,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
}

func (e Entity) AnnotationMember(attribute Attribute, key string, instance uint32, auth Token) (bool, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodMember,
		Aspect:        AspectAnnotation,
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

func (e Entity) AnnotationName(attribute Attribute, key string, index uint64, instance uint32, auth Token) (string, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodName,
		Aspect:        AspectAnnotation,
		Attribute:     attribute,
		Key:           key,
		Instance:      instance,
		Index:         index,
		Authorization: auth,
	})
	if err != nil {
		return "", err
	}
	return val.UnwrapString(), nil
}

func (e Entity) AnnotationKey(attribute Attribute, index uint64, instance uint32, auth Token) (string, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodKey,
		Aspect:        AspectAnnotation,
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

func (e Entity) AnnotationValue(attribute Attribute, key string, index uint64, instance uint32, auth Token) (Value, error) {
	return e.Invoke(ArgsInvokeEntity{
		Method:        MethodValue,
		Aspect:        AspectAnnotation,
		Attribute:     attribute,
		Key:           key,
		Index:         index,
		Instance:      instance,
		Authorization: auth,
	})
}

func (e Entity) AnnotationIndex(attribute Attribute, key string, instance uint32, auth Token) (uint64, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodIndex,
		Aspect:        AspectAnnotation,
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

func (e Entity) AnnotationAttribute(key string, index uint64, instance uint32, auth Token) (Attribute, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodAttribute,
		Aspect:        AspectAnnotation,
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

func (e Entity) SortAnnotations(key string, instance uint32, parameter Parameter, auth Token) {
	e.Invoke(ArgsInvokeEntity{
		Method:        MethodSort,
		Aspect:        AspectAnnotation,
		Key:           key,
		Instance:      instance,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
}

func (e Entity) PurgeAnnotations(key string, instance uint32, parameter Parameter, auth Token) {
	e.Invoke(ArgsInvokeEntity{
		Method:        MethodPurge,
		Aspect:        AspectAnnotation,
		Key:           key,
		Instance:      instance,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
}

func (e Entity) RetrieveAnnotations(key string, instance uint32, auth Token) (string, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodRetrieve,
		Aspect:        AspectAnnotation,
		Key:           key,
		Instance:      instance,
		Authorization: auth,
	})
	if err != nil {
		return "", err
	}
	return val.UnwrapString(), nil
	// TODO: Test the unmarshalling
	// var res []Annotation
	// err = json.Unmarshal([]byte(val.UnwrapString()), &res)
}
