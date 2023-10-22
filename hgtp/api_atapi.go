package hgtp

import (
	"fmt"
	"time"
)

/*

Missing calls:

ATAPI:
- Active  (EntityActive)                   report active, WHAT IS THIS?
- Adapting (EntityAdapting)                report adapting, number of adapters in queue
- Address (LocalAddress)                   get IPV4 of server as decimal like 2886991882
- Arm (ArmOutlet)                          WHAT IS THIS
- Armed  (EntityArmed)                     report Armed, WHAT IS THIS?
- Attached (AttributeAttached)             Get whether an specific attachmend exists?
- Authorization (EntityAuthorization)      Get one authorization at given index
- Blocking (EntityBlocking)
- Busy  (EntityBusy)
- Call
- Clear (ClearCondition)
- Clock (LocalClock)
- Condition (EntityCondition)
- Conditions (EntityConditions)
- Connected (MethodConnected)
- Context  (EntityContext)
- Disarm (DisarmOutlet)
- Elapsed (EntityElapsed)
- Element  (EntityElement)
- Elements (EntityElements)
- Erase  (EraseElement)
- Execute (ExecuteOutlet)
- Fetch (FetchEntity)
- Finalize
- Find  (FindElement)
- Flush (FlushEvents, FlushOutlet)
- Gateway (LocalGateway)
- Halt (HaltOutlet)
- Hostname (LocalHostname)
- Initialize
- Insert   (InsertElement)
- Internet (LocalInternet)     Gets IPv4 of server as string like 127.20.0.10
- Invoking (EntityInvoking)    report invoking, number of invokers in queue
- Key  (EntityKey)
- Length  (ElementCount)
- Notify
- Redirect (RedirectEntity)
- Remove  (RemoveElement)
- Replace (ReplaceElement)
- Schedule (ScheduleOutlet)
- Server (EntityServer)
- Set (SetCondition)
- Start (StartOutlet)
- State (EntityState)
- Status (LocalStatus)      "2023-07-27 01:54:20 PHYSICAL 30"
- Stop (StopOutlet)
- Subscribed (EntitySubscribed)
- Sync (SynchronizeOutlet)
- Test (TestCondition)
- Timer (EntityTimer)
- Version (LocalVersion)
- Waiting (EntityWaiting)


AVAPI:
- save_entity / restore_entity
- purge_entity
- store_entity
- retrieve_entity
- index_entity
- search_entity
- refresh_entity
- entity_ready  (invoke echo)
- create_outlet
- delete_outlet
- read_data
- write_data
- purge_data
- data_count
... and many more


ATAPI but never used in AVAPI:
- Call
- Notify


*/

///////////////////////////////////////////////////////////////////////////////
//
//
//                              Entity operations
//
//
///////////////////////////////////////////////////////////////////////////////

// ATAPI: CREATE
func (c *Client) CreateEntity(args ArgsCreateEntity) (Entity, error) {

	msg, err := args.toHGTPMessage()
	if err != nil {
		return Entity{}, fmt.Errorf("Failed to create hgtp message: %s", err)
	}

	msg, err = c.transact(msg, 300*time.Second)
	if err != nil {
		return Entity{}, err
	}

	return msg.UnpackEntity(), nil
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

// ATAPI: DELETE
func (c *Client) DeleteEntity(args ArgsDeleteEntity) error {

	msg := args.toHGTPMessage()

	_, err := c.transact(msg, 300*time.Second)
	return err
}

type ArgsDeleteEntity struct {
	Entity        Entity
	Authorization Token
}

// ATAPI: INVOKE
func (c *Client) InvokeEntity(args ArgsInvokeEntity) (Tag, []byte, error) {

	msg, err := args.toHGTPMessage()
	if err != nil {
		return 0, nil, fmt.Errorf("Failed to create hgtp message: %s", err)
	}

	msg, err = c.transact(msg, 300*time.Second)
	if err != nil {
		return 0, nil, err
	}

	return msg.UnpackTag(), msg.UnpackBytes(), nil
}

type ArgsInvokeEntity struct {
	Entity        Entity
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
	Tag           Tag
	Bytes         []byte
	Time          time.Time
	Timeout       Timeout
	Authority     Token
	Authorization Token
}

// ATAPI: INQUIRE
func (c *Client) InquireEntity(args ArgsInquireEntity) (Tag, []byte, error) {

	msg, err := args.toHGTPMessage()
	if err != nil {
		return 0, nil, fmt.Errorf("Failed to create hgtp message: %s", err)
	}

	msg, err = c.transact(msg, 300*time.Second)
	if err != nil {
		return 0, nil, err
	}

	return msg.UnpackTag(), msg.UnpackBytes(), nil
}

type ArgsInquireEntity struct {
	Entity        Entity
	Attribute     Attribute
	Auxiliary     Entity
	Ancillary     Entity
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
	Tag           Tag
	Bytes         []byte
	Time          time.Time
	Timeout       Timeout
	Authority     Token
	Authorization Token
}

// ATAPI: REFERENCE
func (c *Client) ReferenceEntity(args ArgsReferenceEntity) error {

	msg := args.toHGTPMessage()

	_, err := c.transact(msg, 300*time.Second)
	return err
}

type ArgsReferenceEntity struct {
	Entity Entity

	Count uint64

	Authorization Token
}

// ATAPI: DEREFERENCE
func (c *Client) DereferenceEntity(args ArgsDereferenceEntity) error {

	msg := args.toHGTPMessage()

	_, err := c.transact(msg, 300*time.Second)
	return err
}

type ArgsDereferenceEntity struct {
	Entity        Entity
	Count         uint64
	Authorization Token
}

///////////////////////////////////////////////////////////////////////////////
//
//
//                            Connection operations
//
//
///////////////////////////////////////////////////////////////////////////////

// ATAPI: CONNECT
func (c *Client) ConnectMethod(args ArgsConnectMethod) error {
	msg := args.toHGTPMessage()

	_, err := c.transact(msg, 300*time.Second)
	return err
}

type ArgsConnectMethod struct {
	Entity        Entity
	Outlet        Entity
	Method        Method
	Precedence    Precedence
	Timeout       Timeout
	Authorization Token
}

// ATAPI: DISCONNECT
func (c *Client) DisconnectMethod(args ArgsDisconnectMethod) error {
	msg := args.toHGTPMessage()

	_, err := c.transact(msg, 300*time.Second)
	return err
}

type ArgsDisconnectMethod struct {
	Entity        Entity
	Outlet        Entity
	Method        Method
	Precedence    Precedence
	Authorization Token
}

// ATAPI: CONNECTION
func (c *Client) EntityConnection(args ArgsEntityConnection) (
	outlet Entity,
	method Method,
	precedence Precedence,
	error error,
) {
	msg := args.toHGTPMessage()

	msg, err := c.transact(msg, 300*time.Second)
	if err != nil {
		return Entity{}, 0, 0, err
	}

	return msg.UnpackOutlet(),
		msg.UnpackMethod(),
		msg.UnpackPrecedence(),
		nil
}

type ArgsEntityConnection struct {
	Entity        Entity
	Index         uint64
	Authorization Token
}

///////////////////////////////////////////////////////////////////////////////
//
//
//                            Attachment operations
//
//
///////////////////////////////////////////////////////////////////////////////

// ATAPI: ATTACH
func (c *Client) AttachAttribute(args ArgsAttachAttribute) error {
	msg := args.toHGTPMessage()

	_, err := c.transact(msg, 300*time.Second)
	return err
}

type ArgsAttachAttribute struct {
	Entity        Entity
	Outlet        Entity
	Attribute     Attribute
	Precedence    Precedence
	Timeout       Timeout
	Authorization Token
}

// ATAPI: DETACH
func (c *Client) DetachAttribute(args ArgsDetachAttribute) error {
	msg := args.toHGTPMessage()

	_, err := c.transact(msg, 300*time.Second)
	return err
}

type ArgsDetachAttribute struct {
	Entity        Entity
	Outlet        Entity
	Attribute     Attribute
	Precedence    Precedence
	Authorization Token
}

// ATAPI: ATTACHMENT
func (c *Client) EntityAttachment(args ArgsEntityAttachment) (
	outlet Entity,
	attribute Attribute,
	precedence Precedence,
	error error,
) {
	msg := args.toHGTPMessage()

	msg, err := c.transact(msg, 300*time.Second)
	if err != nil {
		return Entity{}, 0, 0, err
	}

	return msg.UnpackOutlet(),
		msg.UnpackAttribute(),
		msg.UnpackPrecedence(),
		nil
}

type ArgsEntityAttachment struct {
	Entity        Entity
	Index         uint64
	Authorization Token
}

///////////////////////////////////////////////////////////////////////////////
//
//
//                              Outlet operations
//
//
///////////////////////////////////////////////////////////////////////////////

// ATAPI: ACTIVATE
func (c *Client) ActivateOutlet(args ArgsActivateOutlet) error {
	msg := args.toHGTPMessage()

	_, err := c.transact(msg, 300*time.Second)
	return err
}

type ArgsActivateOutlet struct {
	Outlet        Entity
	Authorization Token
}

// ATAPI: DEACTIVATE
func (c *Client) DeactivateOutlet(args ArgsDeactivateOutlet) error {
	msg := args.toHGTPMessage()

	_, err := c.transact(msg, 300*time.Second)
	return err
}

type ArgsDeactivateOutlet struct {
	Outlet        Entity
	Authorization Token
}

// This call will be blocking until it receive an `INVOKE` or `INQUIRE` call,
// after which the given `Callback` will be called with the request's
// arguments, then the value that `Callback` returns is sent back to the
// invoker/inquirer and this call finally returns.
//
// If the `Callback` returns an error, the error is sent back to the
// invoker/inquirer and this call returns with no error
// If the error is not an `AvesTerraError`, it is wrapped in an
// `ErrorApplication` error.
// Two interesting errors are `ErrorBypass` and `ErrorForward`.
// See `Error` taxonomy for more information.
// ATAPI: ADAPT
func (c *Client) AdaptOutlet(args ArgsAdaptOutlet) error {
	msg := args.toHGTPMessage()

	chann, err := c.channPool.Borrow()
	if err != nil {
		return err
	}

	msg, err = chann.SendRecv(msg, 300*time.Second)
	if err != nil {
		c.channPool.Close(chann)
		return err
	}

	err = msg.Validate()
	if err != nil {
		c.channPool.Return(chann)
		return err
	}

	var tag Tag
	var bytes []byte

	if args.Callback != nil {
		tag, bytes, err = args.Callback(NewArgsAdapterCallback(msg))
	}

	if err != nil {
		avtErr, isAvtErr := err.(*AvesTerraError)
		if !isAvtErr {
			avtErr = NewErrorApplication(err.Error())
		}

		msg.PackError(avtErr.Code)
		msg.PackBytes([]byte(avtErr.Message))
	} else {
		msg.PackTag(tag)
		msg.PackBytes(bytes)
	}

	err = chann.Send(msg)
	if err != nil {
		c.channPool.Close(chann)
		return err
	}

	c.channPool.Return(chann)
	return nil
}

type ArgsAdaptOutlet struct {
	Outlet        Entity
	Timeout       Timeout
	Authorization Token

	// If null, an AvesTerra connection/attachment will still be
	// created but no callback will be called when we get an invoke/inquire.
	// Instead, null tag and null bytes will be returned to the
	// invoker/inquirer
	Callback func(ArgsAdapterCallback) (Tag, []byte, error)
}

// Argument of an INVOKE/INQUIRE the adapt callback will receive as argument
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
	Tag           Tag
	Bytes         []byte
	Time          time.Time
	Timeout       Timeout
	Authority     Token
	Authorization Token
}

// ATAPI: RESET
func (c *Client) ResetOutlet(args ArgsResetOutlet) error {
	msg := args.toHGTPMessage()

	_, err := c.transact(msg, 300*time.Second)
	return err
}

type ArgsResetOutlet struct {
	Outlet        Entity
	Authorization Token
}

///////////////////////////////////////////////////////////////////////////////
//
//
//                              Events operations
//
//
///////////////////////////////////////////////////////////////////////////////

// ATAPI: PUBLISH
func (c *Client) PublishEvent(args ArgsPublishEvent) error {
	msg, err := args.toHGTPMessage()
	if err != nil {
		return fmt.Errorf("failed to create HGTP message: %s", err)
	}

	_, err = c.transact(msg, 300*time.Second)
	return err
}

type ArgsPublishEvent struct {
	Entity        Entity
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
	Tag           Tag
	Bytes         []byte
	Time          time.Time
	Timeout       Timeout
	Authority     Token
	Authorization Token
}

// ATAPI: SUBSCRIBE
func (c *Client) SubscribeEvent(args ArgsSubscribeEvent) error {
	msg := args.toHGTPMessage()

	_, err := c.transact(msg, 300*time.Second)
	return err
}

type ArgsSubscribeEvent struct {
	Entity        Entity
	Outlet        Entity
	Event         Event
	Precedence    Precedence
	Timeout       Timeout
	Authorization Token
}

// ATAPI: UNSUBSCRIBE
func (c *Client) UnsubscribeEvent(args ArgsUnsubscribeEvent) error {
	msg := args.toHGTPMessage()

	_, err := c.transact(msg, 300*time.Second)
	return err
}

type ArgsUnsubscribeEvent struct {
	Entity        Entity
	Outlet        Entity
	Event         Event
	Precedence    Precedence
	Authorization Token
}

// ATAPI: WAIT
func (c *Client) WaitEvent(args ArgsWaitEvent) error {
	msg := args.toHGTPMessage()

	msg, err := c.transact(msg, 300*time.Second)
	if err != nil {
		return err
	}

	if args.Callback != nil {
		args.Callback(NewArgsWaitCallback(msg))
	}

	return nil
}

type ArgsWaitEvent struct {
	Outlet        Entity
	Timeout       Timeout
	Authorization Token

	// If null, an AvesTerra subscription will still be created but
	// no callback will be called when we get an event
	Callback func(ArgsWaitCallback)
}

type ArgsWaitCallback struct {
	// The entity the event has been published on
	Entity Entity
	// The outlet the rendezvous has taken place in
	Outlet     Entity
	Method     Method
	Auxiliary  Entity
	Ancillary  Entity
	Attribute  Attribute
	Instance   uint32
	Offset     uint32
	Name       string
	Key        string
	Parameter  int64
	Resultant  int64
	Index      uint64
	Count      uint64
	Aspect     Aspect
	Context    Context
	Category   Category
	Class      Class
	Event      Event
	Mode       Mode
	State      State
	Condition  Condition
	Precedence Precedence
	Tag        Tag
	Bytes      []byte
	Time       time.Time
	Timeout    Timeout
	Authority  Token
	// The authorization the event was published with
	Authorization Token
}

// ATAPI: SUBSCRIPTION
func (c *Client) EntitySubscription(args ArgsEntitySubscription) (
	outlet Entity,
	event Event,
	precedence Precedence,
	expiration time.Time,
	err error,
) {
	msg := args.toHGTPMessage()

	msg, err = c.transact(msg, 300*time.Second)
	if err != nil {
		return Entity{}, 0, 0, time.Time{}, err
	}

	return msg.UnpackOutlet(), msg.UnpackEvent(), msg.UnpackPrecedence(), msg.UnpackTime(), nil
}

type ArgsEntitySubscription struct {
	Entity        Entity
	Index         uint64
	Authorization Token
}

///////////////////////////////////////////////////////////////////////////////
//
//
//                               Lock operations
//
//
///////////////////////////////////////////////////////////////////////////////

// ATAPI: LOCK
func (c *Client) LockOutlet(args ArgsLockOutlet) error {
	msg := args.toHGTPMessage()

	_, err := c.transact(msg, 300*time.Second)
	return err
}

type ArgsLockOutlet struct {
	Outlet        Entity
	Timeout       Timeout
	Authorization Token
}

// ATAPI: UNLOCK
func (c *Client) UnlockOutlet(args ArgsUnlockOutlet) error {
	msg := args.toHGTPMessage()

	_, err := c.transact(msg, 300*time.Second)
	return err
}

type ArgsUnlockOutlet struct {
	Outlet        Entity
	Authorization Token
}

///////////////////////////////////////////////////////////////////////////////
//
//
//                           Authorization operations
//
//
///////////////////////////////////////////////////////////////////////////////

// ATAPI: AUTHORIZE
func (c *Client) AuthorizeEntity(args ArgsAuthorizeEntity) error {
	msg := args.toHGTPMessage()

	_, err := c.transact(msg, 300*time.Second)
	return err
}

type ArgsAuthorizeEntity struct {
	Entity        Entity
	Authority     Token
	Authorization Token

	Restriction int64
}

// ATAPI: DEAUTHORIZE
func (c *Client) DeauthorizeEntity(args ArgsDeauthorizeEntity) error {
	msg := args.toHGTPMessage()

	_, err := c.transact(msg, 300*time.Second)
	return err
}

type ArgsDeauthorizeEntity struct {
	Entity        Entity
	Authority     Token
	Authorization Token
}

// ATAPI: AUTHORIZED
func (c *Client) EntityAuthorized(args ArgsEntityAuthorized) (bool, error) {
	msg := args.toHGTPMessage()

	msg, err := c.transact(msg, 300*time.Second)
	if err != nil {
		return false, err
	}

	return msg.UnpackResultant() != 0, nil
}

type ArgsEntityAuthorized struct {
	Entity        Entity
	Authority     Token
	Authorization Token
}

// ATAPI: AUTHORITY
func (c *Client) EntityAuthority(args ArgsEntityAuthority) (Token, error) {
	msg := args.toHGTPMessage()

	msg, err := c.transact(msg, 300*time.Second)
	if err != nil {
		return NullToken, err
	}

	return msg.UnpackAuthority(), nil
}

type ArgsEntityAuthority struct {
	Entity        Entity
	Authorization Token
}

///////////////////////////////////////////////////////////////////////////////
//
//
//                             Auxiliary operations
//
//
///////////////////////////////////////////////////////////////////////////////

// ATAPI: NAME
func (c *Client) EntityName(args ArgsEntityName) (string, error) {
	msg := args.toHGTPMessage()

	msg, err := c.transact(msg, 300*time.Second)
	if err != nil {
		return "", err
	}

	return msg.UnpackName(), nil
}

type ArgsEntityName struct {
	Entity        Entity
	Authorization Token
}

// ATAPI: CLASS
func (c *Client) EntityClass(args ArgsEntityClass) (Class, error) {
	msg := args.toHGTPMessage()

	msg, err := c.transact(msg, 300*time.Second)
	if err != nil {
		return 0, err
	}

	return msg.UnpackClass(), nil
}

type ArgsEntityClass struct {
	Entity        Entity
	Authorization Token
}

// ATAPI: CATEGORY
func (c *Client) EntityCategory(args ArgsEntityCategory) (Category, error) {
	msg := args.toHGTPMessage()

	msg, err := c.transact(msg, 300*time.Second)
	if err != nil {
		return 0, err
	}

	return msg.UnpackCategory(), nil
}

type ArgsEntityCategory struct {
	Entity        Entity
	Authorization Token
}

// ATAPI: SERVER
func (c *Client) LocalServer(args ArgsLocalServer) (Entity, error) {
	msg := args.toHGTPMessage()

	msg, err := c.transact(msg, 300*time.Second)
	if err != nil {
		return Entity{}, err
	}

	return msg.UnpackEntity(), nil
}

type ArgsLocalServer struct {
	Entity        Entity
	Authorization Token
}

// ATAPI: REDIRECTION
func (c *Client) EntityRedirection(args ArgsEntityRedirection) (Entity, error) {
	msg := args.toHGTPMessage()

	msg, err := c.transact(msg, 300*time.Second)
	if err != nil {
		return Entity{}, err
	}

	return msg.UnpackEntity(), nil
}

type ArgsEntityRedirection struct {
	Entity        Entity
	Authorization Token
}

// ATAPI: CHANGE
func (c *Client) ChangeEntity(args ArgsChangeEntity) error {
	msg, err := args.toHGTPMessage()
	if err != nil {
		return fmt.Errorf("Failed to create hgtp message: %s", err)
	}

	_, err = c.transact(msg, 300*time.Second)
	return err
}

type ArgsChangeEntity struct {
	Entity        Entity
	Name          string
	Key           string
	Context       Context
	Category      Category
	Class         Class
	State         State
	Authority     Token
	Authorization Token
}

// ATAPI: TIMESTAMP
func (c *Client) EntityTimestamp(args ArgsEntityTimestamp) (time.Time, error) {
	msg := args.toHGTPMessage()

	msg, err := c.transact(msg, 300*time.Second)
	if err != nil {
		return time.Time{}, err
	}

	return msg.UnpackTime(), nil
}

type ArgsEntityTimestamp struct {
	Entity        Entity
	Authorization Token
}

// ATAPI: AVAILABLE
func (c *Client) EntityAvailable(args ArgsEntityAvailable) (bool, error) {
	msg := args.toHGTPMessage()

	msg, err := c.transact(msg, 300*time.Second)
	if err != nil {
		return false, err
	}

	return msg.UnpackResultant() != 0, nil
}

type ArgsEntityAvailable struct {
	Entity        Entity
	Authorization Token
}

// ATAPI: EXTINCT
func (c *Client) EntityExtinct(args ArgsEntityExtinct) (bool, error) {
	msg := args.toHGTPMessage()

	msg, err := c.transact(msg, 300*time.Second)
	if err != nil {
		return false, err
	}

	return msg.UnpackResultant() != 0, nil
}

type ArgsEntityExtinct struct {
	Entity        Entity
	Authorization Token
}

// ATAPI: ACTIVATED
func (c *Client) EntityActivated(args ArgsEntityActivated) (bool, error) {
	msg := args.toHGTPMessage()

	msg, err := c.transact(msg, 300*time.Second)
	if err != nil {
		return false, err
	}

	return msg.UnpackResultant() != 0, nil
}

type ArgsEntityActivated struct {
	Entity        Entity
	Authorization Token
}

// ATAPI: LOCKED
func (c *Client) EntityLocked(args ArgsEntityLocked) (bool, error) {
	msg := args.toHGTPMessage()

	msg, err := c.transact(msg, 300*time.Second)
	if err != nil {
		return false, err
	}

	return msg.UnpackResultant() != 0, nil
}

type ArgsEntityLocked struct {
	Entity        Entity
	Authorization Token
}

// ATAPI: PENDING
func (c *Client) Pending(args ArgsPending) (uint64, error) {
	msg := args.toHGTPMessage()

	msg, err := c.transact(msg, 300*time.Second)
	if err != nil {
		return 0, err
	}

	return msg.UnpackCount(), nil
}

type ArgsPending struct {
	Entity        Entity
	Authorization Token
}

// ATAPI: RENDEZVOUS
func (c *Client) EntityRendezvous(args ArgsEntityRendezvous) (string, error) {
	msg := args.toHGTPMessage()

	msg, err := c.transact(msg, 300*time.Second)
	if err != nil {
		return "", err
	}

	return string(msg.UnpackBytes()), nil
}

type ArgsEntityRendezvous struct {
	Entity        Entity
	Authorization Token
}

// ATAPI: REFERENCES
func (c *Client) EntityReferences(args ArgsEntityReferences) (uint64, error) {
	msg := args.toHGTPMessage()

	msg, err := c.transact(msg, 300*time.Second)
	if err != nil {
		return 0, err
	}

	return msg.UnpackCount(), nil
}

type ArgsEntityReferences struct {
	Entity        Entity
	Authorization Token
}

// ATAPI: CONNECTIONS
func (c *Client) ConnectionCount(args ArgsConnectionCount) (uint64, error) {
	msg := args.toHGTPMessage()

	msg, err := c.transact(msg, 300*time.Second)
	if err != nil {
		return 0, err
	}

	return msg.UnpackCount(), nil
}

type ArgsConnectionCount struct {
	Entity        Entity
	Authorization Token
}

// ATAPI: ATTACHMENTS
func (c *Client) AttachmentCount(args ArgsAttachmentCount) (uint64, error) {
	msg := args.toHGTPMessage()

	msg, err := c.transact(msg, 300*time.Second)
	if err != nil {
		return 0, err
	}

	return msg.UnpackCount(), nil
}

type ArgsAttachmentCount struct {
	Entity        Entity
	Authorization Token
}

// ATAPI: SUBSCRIPTIONS
func (c *Client) SubscriptionCount(args ArgsSubscriptionCount) (uint64, error) {
	msg := args.toHGTPMessage()

	msg, err := c.transact(msg, 300*time.Second)
	if err != nil {
		return 0, err
	}

	return msg.UnpackCount(), nil
}

type ArgsSubscriptionCount struct {
	Entity        Entity
	Authorization Token
}

// ATAPI: AUTHORIZATIONS
func (c *Client) EntityAuthorizations(args ArgsEntityAuthorizations) (uint64, error) {
	msg := args.toHGTPMessage()

	msg, err := c.transact(msg, 300*time.Second)
	if err != nil {
		return 0, err
	}

	return msg.UnpackCount(), nil
}

type ArgsEntityAuthorizations struct {
	Entity        Entity
	Authorization Token
}
