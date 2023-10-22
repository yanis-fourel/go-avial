package av

func CreateFile(args ArgsCreateFile) (Entity, error) {
	if args.Outlet == EntityNull {
		args.Outlet = OUTLET_FILES
	}

	val, err := args.Outlet.Invoke(ArgsInvokeEntity{
		Method:        MethodCreate,
		Name:          args.Name,
		Key:           args.Key,
		Mode:          args.Mode,
		Ancillary:     args.Server,
		Authorization: args.Authorization,
	})
	if err != nil {
		return EntityNull, err
	}
	return val.UnwrapEntity()
}

type ArgsCreateFile struct {
	Name          string
	Key           string
	Mode          Mode
	Outlet        Entity
	Server        Entity
	Authorization Token
}

func (e Entity) DeleteFile(auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodDelete,
		Authorization: auth,
	})
	return err
}
