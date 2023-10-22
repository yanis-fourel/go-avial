package av

import (
	"github.com/yanis-fourel/go-avial/hgtp"
	"time"
)

type Entity struct {
	hgtp.Entity
}

var EntityNull = NewEntity(0, 0, 0)

func NewEntity(pid uint32, hid uint32, uid uint64) Entity {
	return Entity{hgtp.NewEntity(pid, hid, uid)}
}

func ParseEntity(str string) (Entity, error) {
	e, err := hgtp.ParseEntity(str)
	return Entity{e}, err
}

func (e Entity) Delete(auth Token) error {
	return gClient.DeleteEntity(hgtp.ArgsDeleteEntity{
		Entity:        e.Entity,
		Authorization: auth,
	})
}

func (e Entity) Invoke(args ArgsInvokeEntity) (Value, error) {
	t, b, err := gClient.InvokeEntity(hgtp.ArgsInvokeEntity{
		Entity:        e.Entity,
		Method:        hgtp.Method(args.Method),
		Auxiliary:     args.Auxiliary.Entity,
		Ancillary:     args.Ancillary.Entity,
		Attribute:     hgtp.Attribute(args.Attribute),
		Instance:      args.Instance,
		Offset:        args.Offset,
		Name:          args.Name,
		Key:           args.Key,
		Parameter:     args.Parameter,
		Resultant:     args.Resultant,
		Index:         args.Index,
		Count:         args.Count,
		Aspect:        hgtp.Aspect(args.Aspect),
		Context:       hgtp.Context(args.Context),
		Category:      hgtp.Category(args.Category),
		Class:         hgtp.Class(args.Class),
		Event:         hgtp.Event(args.Event),
		Mode:          hgtp.Mode(args.Mode),
		State:         hgtp.State(args.State),
		Condition:     hgtp.Condition(args.Condition),
		Precedence:    hgtp.Precedence(args.Precedence),
		Tag:           hgtp.Tag(args.Value.Tag),
		Bytes:         args.Value.Bytes,
		Time:          args.Time,
		Timeout:       args.Timeout,
		Authority:     args.Authority,
		Authorization: args.Authorization,
	})
	return Value{Tag(t), b}, err
}

type ArgsInvokeEntity struct {
	Method        Method
	Auxiliary     Entity
	Ancillary     Entity
	Attribute     Attribute
	Instance      uint32
	Offset        uint32
	Name          string
	Key           string
	Parameter     int64
	Resultant     int64
	Index         uint64
	Count         uint64
	Aspect        Aspect
	Context       Context
	Category      Category
	Class         Class
	Event         Event
	Mode          Mode
	State         State
	Condition     Condition
	Precedence    Precedence
	Value         Value
	Time          time.Time
	Timeout       Timeout
	Authority     Token
	Authorization Token
}

func (e Entity) Inquire(args ArgsInquireEntity) (Value, error) {
	t, b, err := gClient.InquireEntity(hgtp.ArgsInquireEntity{
		Entity:     e.Entity,
		Auxiliary:  args.Auxiliary.Entity,
		Ancillary:  args.Ancillary.Entity,
		Attribute:  hgtp.Attribute(args.Attribute),
		Instance:   args.Instance,
		Offset:     args.Offset,
		Name:       args.Name,
		Key:        args.Key,
		Parameter:  args.Parameter,
		Resultant:  args.Resultant,
		Index:      args.Index,
		Count:      args.Count,
		Aspect:     hgtp.Aspect(args.Aspect),
		Context:    hgtp.Context(args.Context),
		Category:   hgtp.Category(args.Category),
		Class:      hgtp.Class(args.Class),
		Event:      hgtp.Event(args.Event),
		Mode:       hgtp.Mode(args.Mode),
		State:      hgtp.State(args.State),
		Condition:  hgtp.Condition(args.Condition),
		Precedence: hgtp.Precedence(args.Precedence),
		Tag:        hgtp.Tag(args.Value.Tag),
		Bytes:      args.Value.Bytes,
		Time:       args.Time,
		Timeout:    args.Timeout,
		Authority:  args.Authority,
	})
	return Value{Tag(t), b}, err
}

type ArgsInquireEntity struct {
	Auxiliary     Entity
	Ancillary     Entity
	Attribute     Attribute
	Instance      uint32
	Offset        uint32
	Name          string
	Key           string
	Parameter     int64
	Resultant     int64
	Index         uint64
	Count         uint64
	Aspect        Aspect
	Context       Context
	Category      Category
	Class         Class
	Event         Event
	Mode          Mode
	State         State
	Condition     Condition
	Precedence    Precedence
	Value         Value
	Time          time.Time
	Timeout       Timeout
	Authority     Token
	Authorization Token
}

func (e Entity) Reference(count uint64, auth Token) error {
	return gClient.ReferenceEntity(hgtp.ArgsReferenceEntity{
		Entity:        e.Entity,
		Count:         count,
		Authorization: auth,
	})
}

func (e Entity) Dereference(count uint64, auth Token) error {
	return gClient.DereferenceEntity(hgtp.ArgsDereferenceEntity{
		Entity:        e.Entity,
		Count:         count,
		Authorization: auth,
	})
}

func (e Entity) ConnectMethod(args ArgsConnectMethod) error {
	return gClient.ConnectMethod(hgtp.ArgsConnectMethod{
		Entity:        e.Entity,
		Outlet:        args.Outlet.Entity,
		Method:        hgtp.Method(args.Method),
		Precedence:    hgtp.Precedence(args.Precedence),
		Timeout:       args.Timeout,
		Authorization: args.Authorization,
	})
}

type ArgsConnectMethod struct {
	Outlet        Entity
	Method        Method
	Precedence    Precedence
	Timeout       Timeout
	Authorization Token
}

func (e Entity) Disconnect(args ArgsDisconnect) error {
	return gClient.DisconnectMethod(hgtp.ArgsDisconnectMethod{
		Entity:        e.Entity,
		Outlet:        args.Outlet.Entity,
		Method:        hgtp.Method(args.Method),
		Precedence:    hgtp.Precedence(args.Precedence),
		Authorization: args.Authorization,
	})
}

type ArgsDisconnect struct {
	Outlet        Entity
	Method        Method
	Precedence    Precedence
	Authorization Token
}

func (e Entity) Connection(index uint64, auth Token) (
	outlet Entity,
	method Method,
	precedence Precedence,
	err error,
) {
	o, m, p, err := gClient.EntityConnection(hgtp.ArgsEntityConnection{
		Entity:        e.Entity,
		Index:         index,
		Authorization: auth,
	})

	return Entity{o}, Method(m), Precedence(p), err
}

func (e Entity) AttachAttribute(args ArgsAttachAttribute) error {
	return gClient.AttachAttribute(hgtp.ArgsAttachAttribute{
		Entity:        e.Entity,
		Outlet:        args.Outlet.Entity,
		Attribute:     hgtp.Attribute(args.Attribute),
		Precedence:    hgtp.Precedence(args.Precedence),
		Timeout:       args.Timeout,
		Authorization: args.Authorization,
	})
}

type ArgsAttachAttribute struct {
	Outlet        Entity
	Attribute     Attribute
	Precedence    Precedence
	Timeout       Timeout
	Authorization Token
}

func (e Entity) DetachAttribute(args ArgsDetachAttribute) error {
	return gClient.DetachAttribute(hgtp.ArgsDetachAttribute{
		Entity:        e.Entity,
		Outlet:        args.Outlet.Entity,
		Attribute:     hgtp.Attribute(args.Attribute),
		Precedence:    hgtp.Precedence(args.Precedence),
		Authorization: args.Authorization,
	})
}

type ArgsDetachAttribute struct {
	Outlet        Entity
	Attribute     Attribute
	Precedence    Precedence
	Authorization Token
}

func (e Entity) Attachment(index uint64, auth Token) (
	outlet Entity,
	attribute Attribute,
	precedence Precedence,
	err error,
) {
	o, a, p, err := gClient.EntityAttachment(hgtp.ArgsEntityAttachment{
		Entity:        e.Entity,
		Index:         index,
		Authorization: auth,
	})
	return Entity{o}, Attribute(a), Precedence(p), err
}

func (e Entity) Activate(auth Token) error {
	return gClient.ActivateOutlet(hgtp.ArgsActivateOutlet{
		Outlet:        e.Entity,
		Authorization: auth,
	})
}

func (e Entity) Deactivate(auth Token) error {
	return gClient.DeactivateOutlet(hgtp.ArgsDeactivateOutlet{
		Outlet:        e.Entity,
		Authorization: auth,
	})
}

func (e Entity) Adapt(args ArgsAdapt) error {
	return gClient.AdaptOutlet(hgtp.ArgsAdaptOutlet{
		Outlet:        e.Entity,
		Timeout:       args.Timeout,
		Authorization: args.Authorization,
		Callback: func(invokeArgs hgtp.ArgsAdapterCallback) (hgtp.Tag, []byte, error) {
			value, err := args.Callback(ArgsAdapterCallback{
				Entity:        Entity{invokeArgs.Entity},
				Method:        Method(invokeArgs.Method),
				Outlet:        Entity{invokeArgs.Outlet},
				Auxiliary:     Entity{invokeArgs.Auxiliary},
				Ancillary:     Entity{invokeArgs.Ancillary},
				Attribute:     Attribute(invokeArgs.Attribute),
				Instance:      invokeArgs.Instance,
				Offset:        invokeArgs.Offset,
				Name:          invokeArgs.Name,
				Key:           invokeArgs.Key,
				Parameter:     invokeArgs.Parameter,
				Resultant:     invokeArgs.Resultant,
				Index:         invokeArgs.Index,
				Count:         invokeArgs.Count,
				Aspect:        Aspect(invokeArgs.Aspect),
				Context:       Context(invokeArgs.Context),
				Category:      Category(invokeArgs.Category),
				Class:         Class(invokeArgs.Class),
				Event:         Event(invokeArgs.Event),
				Mode:          Mode(invokeArgs.Mode),
				State:         State(invokeArgs.State),
				Condition:     Condition(invokeArgs.Condition),
				Precedence:    Precedence(invokeArgs.Precedence),
				Value:         Value{Tag(invokeArgs.Tag), invokeArgs.Bytes},
				Time:          invokeArgs.Time,
				Timeout:       invokeArgs.Timeout,
				Authority:     invokeArgs.Authority,
				Authorization: invokeArgs.Authorization,
			})
			return hgtp.Tag(value.Tag), value.Bytes, err
		},
	})
}

type ArgsAdapt struct {
	Timeout       Timeout
	Authorization Token
	Callback      func(ArgsAdapterCallback) (Value, error)
}

type ArgsAdapterCallback struct {
	Entity Entity
	// If call was INQUIRE, the method is always null method (`0`)
	Method Method
	// The outlet the rendezvous has taken place in
	Outlet        Entity
	Auxiliary     Entity
	Ancillary     Entity
	Attribute     Attribute
	Instance      uint32
	Offset        uint32
	Name          string
	Key           string
	Parameter     int64
	Resultant     int64
	Index         uint64
	Count         uint64
	Aspect        Aspect
	Context       Context
	Category      Category
	Class         Class
	Event         Event
	Mode          Mode
	State         State
	Condition     Condition
	Precedence    Precedence
	Value         Value
	Time          time.Time
	Timeout       Timeout
	Authority     Token
	Authorization Token
}

func (e Entity) Reset(auth Token) error {
	return gClient.ResetOutlet(hgtp.ArgsResetOutlet{
		Outlet:        e.Entity,
		Authorization: auth,
	})
}

func (e Entity) PublishEvent(args ArgsPublishEvent) error {
	return gClient.PublishEvent(hgtp.ArgsPublishEvent{
		Entity:        e.Entity,
		Event:         hgtp.Event(args.Event),
		Auxiliary:     args.Auxiliary.Entity,
		Ancillary:     args.Ancillary.Entity,
		Method:        hgtp.Method(args.Method),
		Attribute:     hgtp.Attribute(args.Attribute),
		Instance:      args.Instance,
		Offset:        args.Offset,
		Name:          args.Name,
		Key:           args.Key,
		Parameter:     args.Parameter,
		Resultant:     args.Resultant,
		Index:         args.Index,
		Count:         args.Count,
		Aspect:        hgtp.Aspect(args.Aspect),
		Context:       hgtp.Context(args.Context),
		Category:      hgtp.Category(args.Category),
		Class:         hgtp.Class(args.Class),
		Mode:          hgtp.Mode(args.Mode),
		State:         hgtp.State(args.State),
		Condition:     hgtp.Condition(args.Condition),
		Precedence:    hgtp.Precedence(args.Precedence),
		Tag:           hgtp.Tag(args.Value.Tag),
		Bytes:         args.Value.Bytes,
		Time:          args.Time,
		Timeout:       args.Timeout,
		Authority:     args.Authority,
		Authorization: args.Authorization,
	})
}

type ArgsPublishEvent struct {
	Event         Event
	Auxiliary     Entity
	Ancillary     Entity
	Method        Method
	Attribute     Attribute
	Instance      uint32
	Offset        uint32
	Name          string
	Key           string
	Parameter     int64
	Resultant     int64
	Index         uint64
	Count         uint64
	Aspect        Aspect
	Context       Context
	Category      Category
	Class         Class
	Mode          Mode
	State         State
	Condition     Condition
	Precedence    Precedence
	Value         Value
	Time          time.Time
	Timeout       Timeout
	Authority     Token
	Authorization Token
}

func (e Entity) SubscribeEvent(args ArgsSubscribeEvent) error {
	return gClient.SubscribeEvent(hgtp.ArgsSubscribeEvent{
		Entity:        e.Entity,
		Outlet:        args.Outlet.Entity,
		Event:         hgtp.Event(args.Event),
		Precedence:    hgtp.Precedence(args.Precedence),
		Timeout:       args.Timeout,
		Authorization: args.Authorization,
	})
}

type ArgsSubscribeEvent struct {
	Outlet        Entity
	Event         Event
	Precedence    Precedence
	Timeout       Timeout
	Authorization Token
}

func (e Entity) UnsubscribeEvent(args ArgsUnsubscribeEvent) error {
	return gClient.UnsubscribeEvent(hgtp.ArgsUnsubscribeEvent{
		Entity:        e.Entity,
		Outlet:        args.Outlet.Entity,
		Event:         hgtp.Event(args.Event),
		Precedence:    hgtp.Precedence(args.Precedence),
		Authorization: args.Authorization,
	})
}

type ArgsUnsubscribeEvent struct {
	Outlet        Entity
	Event         Event
	Precedence    Precedence
	Authorization Token
}

func (e Entity) WaitEvent(args ArgsWaitEvent) error {
	return gClient.WaitEvent(hgtp.ArgsWaitEvent{
		Outlet:        e.Entity,
		Timeout:       args.Timeout,
		Authorization: args.Authorization,
		Callback: func(publishArgs hgtp.ArgsWaitCallback) {
			args.Callback(ArgsWaitCallback{
				Entity:        Entity{publishArgs.Entity},
				Outlet:        Entity{publishArgs.Outlet},
				Method:        Method(publishArgs.Method),
				Auxiliary:     Entity{publishArgs.Auxiliary},
				Ancillary:     Entity{publishArgs.Ancillary},
				Attribute:     Attribute(publishArgs.Attribute),
				Instance:      publishArgs.Instance,
				Offset:        publishArgs.Offset,
				Name:          publishArgs.Name,
				Key:           publishArgs.Key,
				Parameter:     publishArgs.Parameter,
				Resultant:     publishArgs.Resultant,
				Index:         publishArgs.Index,
				Count:         publishArgs.Count,
				Aspect:        Aspect(publishArgs.Aspect),
				Context:       Context(publishArgs.Context),
				Category:      Category(publishArgs.Category),
				Class:         Class(publishArgs.Class),
				Event:         Event(publishArgs.Event),
				Mode:          Mode(publishArgs.Mode),
				State:         State(publishArgs.State),
				Condition:     Condition(publishArgs.Condition),
				Precedence:    Precedence(publishArgs.Precedence),
				Value:         Value{Tag(publishArgs.Tag), publishArgs.Bytes},
				Time:          publishArgs.Time,
				Timeout:       publishArgs.Timeout,
				Authority:     publishArgs.Authority,
				Authorization: publishArgs.Authorization,
			})
		},
	})
}

type ArgsWaitEvent struct {
	Timeout       Timeout
	Authorization Token
	Callback      func(ArgsWaitCallback)
}

type ArgsWaitCallback struct {
	// The entity the event has been published on
	Entity Entity
	// The outlet the rendezvous has taken place in
	Outlet        Entity
	Method        Method
	Auxiliary     Entity
	Ancillary     Entity
	Attribute     Attribute
	Instance      uint32
	Offset        uint32
	Name          string
	Key           string
	Parameter     int64
	Resultant     int64
	Index         uint64
	Count         uint64
	Aspect        Aspect
	Context       Context
	Category      Category
	Class         Class
	Event         Event
	Mode          Mode
	State         State
	Condition     Condition
	Precedence    Precedence
	Value         Value
	Time          time.Time
	Timeout       Timeout
	Authority     Token
	Authorization Token
}

func (e Entity) Subscription(index uint64, auth Token) (
	outlet Entity,
	event Event,
	precedence Precedence,
	expiration time.Time,
	err error,
) {
	o, ev, p, ex, err := gClient.EntitySubscription(hgtp.ArgsEntitySubscription{
		Entity:        e.Entity,
		Index:         index,
		Authorization: auth,
	})
	return Entity{o}, Event(ev), Precedence(p), ex, err
}

func (e Entity) LockOutlet(timeout Timeout, auth Token) error {
	return gClient.LockOutlet(hgtp.ArgsLockOutlet{
		Outlet:        e.Entity,
		Timeout:       timeout,
		Authorization: auth,
	})
}

func (e Entity) UnlockOutlet(auth Token) error {
	return gClient.UnlockOutlet(hgtp.ArgsUnlockOutlet{
		Outlet:        e.Entity,
		Authorization: auth,
	})
}

func (e Entity) Authorize(args ArgsAuthorize) error {
	return gClient.AuthorizeEntity(hgtp.ArgsAuthorizeEntity{
		Entity:        e.Entity,
		Authority:     args.Authority,
		Authorization: args.Authorization,
		Restriction:   args.Restriction,
	})
}

type ArgsAuthorize struct {
	Authority     Token
	Authorization Token

	Restriction int64
}

func (e Entity) Deauthorize(args ArgsDeauthorize) error {
	return gClient.DeauthorizeEntity(hgtp.ArgsDeauthorizeEntity{
		Entity:        e.Entity,
		Authority:     args.Authority,
		Authorization: args.Authorization,
	})
}

type ArgsDeauthorize struct {
	Authority     Token
	Authorization Token
}

func (e Entity) Authorized(args ArgsEntityAuthorized) (bool, error) {
	return gClient.EntityAuthorized(hgtp.ArgsEntityAuthorized{
		Entity:        e.Entity,
		Authority:     args.Authority,
		Authorization: args.Authorization,
	})
}

type ArgsEntityAuthorized struct {
	Authority     Token
	Authorization Token
}

func (e Entity) Authority(auth Token) (Token, error) {
	return gClient.EntityAuthority(hgtp.ArgsEntityAuthority{
		Entity:        e.Entity,
		Authorization: auth,
	})
}

func (e Entity) Name(auth Token) (string, error) {
	return gClient.EntityName(hgtp.ArgsEntityName{
		Entity:        e.Entity,
		Authorization: auth,
	})
}

func (e Entity) Class(auth Token) (Class, error) {
	res, err := gClient.EntityClass(hgtp.ArgsEntityClass{
		Entity:        e.Entity,
		Authorization: auth,
	})
	return Class(res), err
}

func (e Entity) Category(auth Token) (Category, error) {
	res, err := gClient.EntityCategory(hgtp.ArgsEntityCategory{
		Entity:        e.Entity,
		Authorization: auth,
	})
	return Category(res), err
}

func (e Entity) LocalServer(auth Token) (Entity, error) {
	res, err := gClient.LocalServer(hgtp.ArgsLocalServer{
		Entity:        e.Entity,
		Authorization: auth,
	})
	return Entity{res}, err
}

func (e Entity) Redirection(auth Token) (Entity, error) {
	res, err := gClient.EntityRedirection(hgtp.ArgsEntityRedirection{
		Entity:        e.Entity,
		Authorization: auth,
	})
	return Entity{res}, err
}

func (e Entity) Change(args ArgsChangeEntity) error {
	return gClient.ChangeEntity(hgtp.ArgsChangeEntity{
		Entity:        e.Entity,
		Name:          args.Name,
		Key:           args.Key,
		Context:       hgtp.Context(args.Context),
		Category:      hgtp.Category(args.Category),
		Class:         hgtp.Class(args.Class),
		State:         hgtp.State(args.State),
		Authority:     args.Authority,
		Authorization: args.Authorization,
	})
}

type ArgsChangeEntity struct {
	Name          string
	Key           string
	Context       Context
	Category      Category
	Class         Class
	State         State
	Authority     Token
	Authorization Token
}

func (e Entity) Timestamp(auth Token) (time.Time, error) {
	return gClient.EntityTimestamp(hgtp.ArgsEntityTimestamp{
		Entity:        e.Entity,
		Authorization: auth,
	})
}

func (e Entity) Available(auth Token) (bool, error) {
	return gClient.EntityAvailable(hgtp.ArgsEntityAvailable{
		Entity:        e.Entity,
		Authorization: auth,
	})
}

func (e Entity) Extinct(auth Token) (bool, error) {
	return gClient.EntityExtinct(hgtp.ArgsEntityExtinct{
		Entity:        e.Entity,
		Authorization: auth,
	})
}

func (e Entity) Activated(auth Token) (bool, error) {
	return gClient.EntityActivated(hgtp.ArgsEntityActivated{
		Entity:        e.Entity,
		Authorization: auth,
	})
}

func (e Entity) Locked(auth Token) (bool, error) {
	return gClient.EntityLocked(hgtp.ArgsEntityLocked{
		Entity:        e.Entity,
		Authorization: auth,
	})
}

func (e Entity) Pending(auth Token) (uint64, error) {
	return gClient.Pending(hgtp.ArgsPending{
		Entity:        e.Entity,
		Authorization: auth,
	})
}

func (e Entity) RendezVous(auth Token) (string, error) {
	return gClient.EntityRendezvous(hgtp.ArgsEntityRendezvous{
		Entity:        e.Entity,
		Authorization: auth,
	})
}

func (e Entity) References(auth Token) (uint64, error) {
	return gClient.EntityReferences(hgtp.ArgsEntityReferences{
		Entity:        e.Entity,
		Authorization: auth,
	})
}

func (e Entity) ConnectionCount(auth Token) (uint64, error) {
	return gClient.ConnectionCount(hgtp.ArgsConnectionCount{
		Entity:        e.Entity,
		Authorization: auth,
	})
}

func (e Entity) AttachmentCount(auth Token) (uint64, error) {
	return gClient.AttachmentCount(hgtp.ArgsAttachmentCount{
		Entity:        e.Entity,
		Authorization: auth,
	})
}

func (e Entity) SubscriptionCount(auth Token) (uint64, error) {
	return gClient.SubscriptionCount(hgtp.ArgsSubscriptionCount{
		Entity:        e.Entity,
		Authorization: auth,
	})
}

func (e Entity) AuthorizationsCount(auth Token) (uint64, error) {
	return gClient.EntityAuthorizations(hgtp.ArgsEntityAuthorizations{
		Entity:        e.Entity,
		Authorization: auth,
	})
}
