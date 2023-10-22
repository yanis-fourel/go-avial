package av

import (
	"errors"
	"github.com/yanis-fourel/go-avial/hgtp"
)

var gClient *hgtp.Client

type ClientConfig = hgtp.ClientConfig

func Initialize(conf *ClientConfig) error {
	if gClient != nil {
		return errors.New("avial client has already been initialized")
	}

	client, err := hgtp.NewClient(conf)
	gClient = client
	return err
}

func Finalize() error {
	if gClient == nil {
		return errors.New("avial client has not been initialized")
	}

	gClient.Close()
	return nil
}

func CreateEntity(args ArgsCreateEntity) (Entity, error) {
	e, err := gClient.CreateEntity(hgtp.ArgsCreateEntity{
		Authorization: args.Authorization,
		Outlet:        args.Outlet.Entity,
		Server:        args.Server.Entity,
		Name:          args.Name,
		Key:           args.Key,
		Context:       hgtp.Context(args.Context),
		Category:      hgtp.Category(args.Category),
		Class:         hgtp.Class(args.Class),
		Method:        hgtp.Method(args.Method),
		Attribute:     hgtp.Attribute(args.Attribute),
		Event:         hgtp.Event(args.Event),
		Precedence:    hgtp.Precedence(args.Precedence),
		Timeout:       args.Timeout,
	})

	return Entity{e}, err
}

type ArgsCreateEntity struct {
	Authorization Token
	Outlet        Entity
	Server        Entity
	Name          string
	Key           string
	Context       Context
	Category      Category
	Class         Class
	Method        Method
	Attribute     Attribute
	Event         Event
	Precedence    Precedence
	Timeout       Timeout
}

func RedirectToken(to Token, from Token, auth Token) error {
	_, err := NewEntity(0, 0, 0).Invoke(ArgsInvokeEntity{
		Method:        MethodAvesterra,
		Name:          "ENABLE",
		Value:         ValueString(from.String()),
		Authority:     to,
		Authorization: auth,
	})
	return err
}

func StopRedirectToken(token Token, auth Token) error {
	_, err := NewEntity(0, 0, 0).Invoke(ArgsInvokeEntity{
		Method:        MethodAvesterra,
		Name:          "DISABLE",
		Value:         ValueString(token.String()),
		Authorization: auth,
	})
	return err
}
