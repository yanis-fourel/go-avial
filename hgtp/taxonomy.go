package hgtp

import (
	"encoding/json"
	"fmt"
)

type Command uint16

const (
	MaxCommand = CommandCondition

	CommandBye           Command = 0  // BYE
	CommandCreate        Command = 1  // CREATE
	CommandDelete        Command = 2  // DELETE
	CommandConnect       Command = 3  // CONNECT
	CommandDisconnect    Command = 4  // DISCONNECT
	CommandAttach        Command = 5  // ATTACH
	CommandDetach        Command = 6  // DETACH
	CommandInvoke        Command = 7  // INVOKE
	CommandInquire       Command = 8  // INQUIRE
	CommandCall          Command = 9  // CALL
	CommandAdapt         Command = 10 // ADAPT
	CommandSync          Command = 11 // SYNC
	CommandConnection    Command = 12 // CONNECTION
	CommandAttachment    Command = 13 // ATTACHMENT
	CommandReference     Command = 14 // REFERENCE
	CommandDereference   Command = 15 // DEREFERENCE
	CommandActivate      Command = 16 // ACTIVATE
	CommandDeactivate    Command = 17 // DEACTIVATE
	CommandPublish       Command = 18 // PUBLISH
	CommandSubscribe     Command = 19 // SUBSCRIBE
	CommandNotify        Command = 20 // NOTIFY
	CommandUnsubscribe   Command = 21 // UNSUBSCRIBE
	CommandSubscription  Command = 22 // SUBSCRIPTION
	CommandWait          Command = 23 // WAIT
	CommandFlush         Command = 24 // FLUSH
	CommandAttached      Command = 25 // ATTACHED
	CommandConnected     Command = 26 // CONNECTED
	CommandSubscribed    Command = 27 // SUBSCRIBED
	CommandAuthorize     Command = 28 // AUTHORIZE
	CommandDeauthorize   Command = 29 // DEAUTHORIZE
	CommandAuthorized    Command = 30 // AUTHORIZED
	CommandAuthorization Command = 31 // AUTHORIZATION
	CommandReport        Command = 32 // REPORT
	CommandLock          Command = 33 // LOCK
	CommandUnlock        Command = 34 // UNLOCK
	CommandChange        Command = 35 // CHANGE
	CommandInsert        Command = 36 // INSERT
	CommandRemove        Command = 37 // REMOVE
	CommandReplace       Command = 38 // REPLACE
	CommandErase         Command = 39 // ERASE
	CommandFind          Command = 40 // FIND
	CommandElement       Command = 41 // ELEMENT
	CommandLength        Command = 42 // LENGTH
	CommandArm           Command = 43 // ARM
	CommandDisarm        Command = 44 // DISARM
	CommandStart         Command = 45 // START
	CommandStop          Command = 46 // STOP
	CommandReset         Command = 47 // RESET
	CommandFetch         Command = 48 // FETCH
	CommandSet           Command = 49 // SET
	CommandClear         Command = 50 // CLEAR
	CommandTest          Command = 51 // TEST
	CommandExecute       Command = 52 // EXECUTE
	CommandHalt          Command = 53 // HALT
	CommandSchedule      Command = 54 // SCHEDULE
	CommandRedirect      Command = 55 // REDIRECT
	CommandCondition     Command = 56 // CONDITION
)

var MapStringToCommand = func() map[string]Command {
	m := make(map[string]Command)
	for i := Command(0); i <= MaxCommand; i++ {
		m[i.String()] = i
	}
	return m
}()

func (c Command) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

func (c *Command) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	if v, ok := MapStringToCommand[s]; ok {
		*c = v
		return nil
	}
	return fmt.Errorf("invalid Command: %q", s)
}

// ---

type Report uint16

const (
	MaxReport = ReportAuthority

	ReportNull           Report = 0  // NULL
	ReportAvailable      Report = 1  // AVAILABLE
	ReportName           Report = 2  // NAME
	ReportCategory       Report = 3  // CATEGORY
	ReportClass          Report = 4  // CLASS
	ReportActivated      Report = 5  // ACTIVATED
	ReportConnections    Report = 6  // CONNECTIONS
	ReportAttachments    Report = 7  // ATTACHMENTS
	ReportSubscriptions  Report = 8  // SUBSCRIPTIONS
	ReportAuthorizations Report = 9  // AUTHORIZATIONS
	ReportRedirection    Report = 10 // REDIRECTION
	ReportReferences     Report = 11 // REFERENCES
	ReportTimestamp      Report = 12 // TIMESTAMP
	ReportServer         Report = 13 // SERVER
	ReportLocal          Report = 14 // LOCAL
	ReportGateway        Report = 15 // GATEWAY
	ReportHostname       Report = 16 // HOSTNAME
	ReportVersion        Report = 17 // VERSION
	ReportExtinct        Report = 18 // EXTINCT
	ReportPending        Report = 19 // PENDING
	ReportLocked         Report = 20 // LOCKED
	ReportRendezvous     Report = 21 // RENDEZVOUS
	ReportStatus         Report = 22 // STATUS
	ReportContext        Report = 23 // CONTEXT
	ReportClock          Report = 24 // CLOCK
	ReportTimer          Report = 25 // TIMER
	ReportElapsed        Report = 26 // ELAPSED
	ReportArmed          Report = 27 // ARMED
	ReportActive         Report = 28 // ACTIVE
	ReportState          Report = 29 // STATE
	ReportConditions     Report = 30 // CONDITIONS
	ReportElements       Report = 31 // ELEMENTS
	ReportKey            Report = 32 // KEY
	ReportBusy           Report = 33 // BUSY
	ReportBlocking       Report = 34 // BLOCKING
	ReportWaiting        Report = 35 // WAITING
	ReportAdapting       Report = 36 // ADAPTING
	ReportInvoking       Report = 37 // INVOKING
	ReportInternet       Report = 38 // INTERNET
	ReportAuthority      Report = 39 // AUTHORITY
)

var MapStringToReport = func() map[string]Report {
	m := make(map[string]Report)
	for i := Report(0); i <= MaxReport; i++ {
		m[i.String()] = i
	}
	return m
}()

func (r Report) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r *Report) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	if v, ok := MapStringToReport[s]; ok {
		*r = v
		return nil
	}
	return fmt.Errorf("invalid Report: %q", s)
}

// ----

type Error uint16

const (
	MaxError = ErrorApplication

	ErrorOk        Error = 0 // OK
	ErrorAvesterra Error = 1 // AVESTERRA

	// Raised when an invalid entity operation is requested.
	ErrorEntity Error = 2 // ENTITY
	// Raised when an invalid outlet operation is requested.
	ErrorOutlet Error = 3 // OUTLET

	// Raised when an operation cannot be completed due to network failure
	ErrorNetwork Error = 4 // NETWORK

	// Raised when an invocation/inquistion fails due to timer expiration
	ErrorTimeout Error = 5 // TIMEOUT

	// Raised when an invalid or insufficient authorization is specified.
	ErrorAuthorization Error = 6 // AUTHORIZATION

	// Raised when an error in an adapter is detected
	ErrorAdapter Error = 7 // ADAPTER

	// Raised when an error in a subscriber is detected
	ErrorSubscriber Error = 8 // SUBSCRIBER

	ErrorApplication Error = 9 // APPLICATION

	// Raised when an adapter detects an error in the data.
	// Both `ErrorBypass` and `ErrorForward` can be returned from an ADAPT // BYPASS
	// callback to signal the server managing the entity that the INVOKE
	// request is to be passed to the next lower precedence level adapter.
	// These errors are captured by the server(s) and are
	// never returned to the originator of the invocation request.
	//
	// The `ErrorBypass` error indicates the adapter did not process the // BYPASS
	// invocation request, and is deferring the operation to the next lower
	// precedence level adapter, if any, so that it may also process the
	// request
	ErrorBypass Error = 10 // BYPASS

	// Raised when an adapter detects an error in the data.
	// Both `ErrorBypass` and `ErrorForward` can be returned from an ADAPT // BYPASS
	// callback to signal the server managing the entity that the INVOKE
	// request is to be passed to the next lower precedence level adapter.
	// These errors are captured by the server(s) and are
	// never returned to the originator of the invocation request.
	//
	// The `ErrorForward` error indicates that the adapter did successfully // FORWARD
	// fulfil the invocation request, but that it is forwarding it on to the
	// next lower precedence level adapter, if any, so that it may also process
	// the request.
	ErrorForward Error = 11 // FORWARD

	ErrorValue     Error = 12 // VALUE
	ErrorMessage   Error = 13 // MESSAGE
	ErrorExecution Error = 14 // EXECUTION
)

var MapStringToError = func() map[string]Error {
	m := make(map[string]Error)
	for i := Error(0); i <= MaxError; i++ {
		m[i.String()] = i
	}
	return m
}()

func (e Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

func (e *Error) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	if v, ok := MapStringToError[s]; ok {
		*e = v
		return nil
	}
	return fmt.Errorf("invalid Error: %q", s)
}
