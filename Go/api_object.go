package av

import (
	"encoding/json"
	"github.com/yanis-fourel/go-avial/hgtp"
)

func CreateObject(args ArgsCreateObject) (Entity, error) {
	if args.Outlet == EntityNull {
		args.Outlet = OUTLET_OBJECTS
	}

	val, err := args.Outlet.Invoke(ArgsInvokeEntity{
		Method:        MethodCreate,
		Name:          args.Name,
		Key:           args.Key,
		Value:         args.Value,
		Context:       args.Context,
		Category:      args.Category,
		Class:         args.Class,
		Mode:          args.Mode,
		Ancillary:     args.Server,
		Authorization: args.Authorization,
	})
	if err != nil {
		return EntityNull, err
	}
	return val.UnwrapEntity()
}

func (e Entity) DeleteObject(auth Token) error {
	_, err := e.Invoke(ArgsInvokeEntity{
		Method:        MethodDelete,
		Authorization: auth,
	})
	return err
}

type ArgsCreateObject struct {
	Name          string
	Key           string
	Value         Value
	Context       Context
	Category      Category
	Class         Class
	Mode          Mode
	Outlet        Entity
	Server        Entity
	Authorization Token
}

func (e Entity) Retrieve(auth Token) (AvialModel, error) {
	_, bytes, err := gClient.InvokeEntity(hgtp.ArgsInvokeEntity{
		Entity:        e.Entity,
		Method:        hgtp.Method(MethodRetrieve),
		Authorization: auth,
	})

	result := AvialModel{}
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(bytes, &result)
	return result, err
}

func (e Entity) Store(model AvialModel, mode Mode, auth Token) error {
	modelAsJson, err := json.Marshal(model)
	if err != nil {
		return err
	}
	_, _, err = gClient.InvokeEntity(hgtp.ArgsInvokeEntity{
		Entity:        e.Entity,
		Method:        hgtp.Method(MethodStore),
		Mode:          hgtp.Mode(mode),
		Tag:           hgtp.Tag(TagInterchange),
		Bytes:         []byte(modelAsJson),
		Authorization: auth,
	})

	return err
}
