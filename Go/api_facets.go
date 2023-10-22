package av

func (e Entity) SetFacet(attribute Attribute, name string, value Value, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodSet,
		Aspect:        AspectFacet,
		Attribute:     attribute,
		Name:          name,
		Value:         value,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) GetFacet(attribute Attribute, name string, auth Token) (Value, error) {
	return e.Invoke(ArgsInvokeEntity{
		Method:        MethodGet,
		Aspect:        AspectFacet,
		Attribute:     attribute,
		Name:          name,
		Authorization: auth,
	})
}

func (e Entity) ClearFacet(attribute Attribute, name string, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodClear,
		Aspect:        AspectFacet,
		Attribute:     attribute,
		Name:          name,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) FacetCount(attribute Attribute, instance uint32, auth Token) (uint64, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodCount,
		Aspect:        AspectFacet,
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

func (e Entity) InsertFacet(attribute Attribute, name string, value Value, index uint64, instance uint32, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodInsert,
		Aspect:        AspectFacet,
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

func (e Entity) RemoveFacet(attribute Attribute, index uint64, instance uint32, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodRemove,
		Aspect:        AspectFacet,
		Attribute:     attribute,
		Index:         index,
		Instance:      instance,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) ReplaceFacet(attribute Attribute, name string, value Value, index uint64, instance uint32, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodReplace,
		Aspect:        AspectFacet,
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

func (e Entity) FindFacet(attribute Attribute, value Value, index uint64, instance uint32, auth Token) (uint64, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodFind,
		Aspect:        AspectFacet,
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

func (e Entity) IncludeFacet(attribute Attribute, name string, value Value, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodInclude,
		Aspect:        AspectFacet,
		Attribute:     attribute,
		Name:          name,
		Value:         value,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) ExcludeFacet(attribute Attribute, name string, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodExclude,
		Aspect:        AspectFacet,
		Attribute:     attribute,
		Name:          name,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) FacetMember(attribute Attribute, name string, instance uint32, auth Token) (bool, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodMember,
		Aspect:        AspectFacet,
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

func (e Entity) FacetName(attribute Attribute, index uint64, instance uint32, auth Token) (string, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodName,
		Aspect:        AspectFacet,
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

func (e Entity) FacetKey(attribute Attribute, name string, index uint64, instance uint32, auth Token) (string, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodKey,
		Aspect:        AspectFacet,
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

func (e Entity) FacetValue(attribute Attribute, name string, index uint64, instance uint32, auth Token) (Value, error) {
	return e.Invoke(ArgsInvokeEntity{
		Method:        MethodValue,
		Aspect:        AspectFacet,
		Attribute:     attribute,
		Name:          name,
		Index:         index,
		Instance:      instance,
		Authorization: auth,
	})
}

func (e Entity) FacetIndex(attribute Attribute, name string, instance uint32, auth Token) (uint64, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodIndex,
		Aspect:        AspectFacet,
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

func (e Entity) FacetAttribute(name string, index uint64, instance uint32, auth Token) (Attribute, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodAttribute,
		Aspect:        AspectFacet,
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

func (e Entity) SortFacets(attribute Attribute, instance uint32, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodSort,
		Aspect:        AspectFacet,
		Attribute:     attribute,
		Instance:      instance,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) PurgeFacets(attribute Attribute, instance uint32, parameter Parameter, auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodPurge,
		Aspect:        AspectFacet,
		Attribute:     attribute,
		Instance:      instance,
		Parameter:     int64(parameter),
		Authorization: auth,
	})
	return err
}

func (e Entity) RetrieveFacets(attribute Attribute, instance uint32, auth Token) (string, error) {
	val, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodRetrieve,
		Aspect:        AspectFacet,
		Attribute:     attribute,
		Instance:      instance,
		Authorization: auth,
	})
	if err != nil {
		return "", err
	}
	return val.UnwrapInterchange(), nil
	// TODO: Test the unmarshalling
	// var res []Facet
	// err = json.Unmarshal([]byte(val.UnwrapString()), &res)
}
