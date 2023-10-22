package hgtp

import (
	"fmt"
	"github.com/yanis-fourel/go-avial"
	"github.com/yanis-fourel/go-avial/hgtp"
	"log"
	"os"
	"testing"
	"time"
)

const rootAuthToken = "636d384c-54d5-4bab-99b0-9b29236cff59"
const KEYLOG_FILE = "/tmp/avesterra-keylog"

func teardown(client *hgtp.Client, token hgtp.Token) {
	rootAuth := hgtp.MustParseToken(rootAuthToken)

	err := client.DeauthorizeEntity(hgtp.ArgsDeauthorizeEntity{
		Entity:        hgtp.NewEntity(0, 0, 0),
		Authority:     token,
		Authorization: rootAuth,
	})
	if err != nil {
		fmt.Println("ERROR - Could not deauthorize token: ", err)
	}

	client.Close()
}

func setup(t *testing.T) (*hgtp.Client, hgtp.Token) {
	var client *hgtp.Client
	token := hgtp.RandomToken()

	ok := false
	defer func() {
		if !ok {
			teardown(client, token)
		}
	}()

	keyLogWriter, err := os.Create(KEYLOG_FILE)
	if err != nil {
		log.Print("Error: ", err)
		return nil, hgtp.Token{}
	}

	config := hgtp.ClientConfig{
		Address:      "localhost",
		KeyLogWriter: keyLogWriter,
	}

	client, err = hgtp.NewClient(&config)
	if err != nil {
		t.Fatal("Failed to create client: ", err)
	}

	rootAuth := hgtp.MustParseToken(rootAuthToken)

	err = client.AuthorizeEntity(hgtp.ArgsAuthorizeEntity{
		Entity:        hgtp.NewEntity(0, 0, 0),
		Authority:     token,
		Authorization: rootAuth,
	})
	if err != nil {
		t.Fatal("Failed to authorize test token: ", err)
	}

	objectOutlet := hgtp.Entity{Pid: 0, Hid: 0, Uid: 11}
	err = client.AuthorizeEntity(hgtp.ArgsAuthorizeEntity{
		Entity:        objectOutlet,
		Authority:     token,
		Authorization: rootAuth,
	})
	if err != nil {
		t.Fatal("Failed to authorize object outlet: ", err)
	}

	ok = true
	return client, token
}

func TestGetters(t *testing.T) {
	client, token := setup(t)
	defer teardown(client, token)

	objectOutlet := hgtp.Entity{Pid: 0, Hid: 0, Uid: 11}

	tag, bytes, err := client.InvokeEntity(hgtp.ArgsInvokeEntity{
		Entity:        objectOutlet,
		Method:        hgtp.Method(av.MethodCreate),
		Name:          "I am an object omg",
		Key:           "object_01",
		Class:         hgtp.Class(av.ClassAdjective),
		Category:      hgtp.Category(av.CategoryAirplane),
		Authorization: token,
	})
	if err != nil {
		t.Fatal(err)
	}

	if av.Tag(tag) != av.TagEntity {
		t.Fatal(fmt.Errorf("Expected tag to be TagEntity, got %s", av.Tag(tag)))
	}
	myEntity, err := hgtp.ParseEntity(string(bytes))
	if err != nil {
		t.Fatal(err)
	}

	err = client.ReferenceEntity(hgtp.ArgsReferenceEntity{
		Entity:        myEntity,
		Authorization: token,
	})
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		err := client.DereferenceEntity(hgtp.ArgsDereferenceEntity{
			Entity:        myEntity,
			Authorization: token,
		})
		if err != nil {
			fmt.Println("WARNING: Failed to dereference entity: ", err.Error())
		}
	}()

	name, err := client.EntityName(hgtp.ArgsEntityName{
		Entity:        myEntity,
		Authorization: token,
	})
	if err != nil {
		t.Fatal(err)
	}
	if name != "I am an object omg" {
		t.Fatal(fmt.Errorf("Expected name to be 'I am an object omg', got '%s'", name))
	}

	class, err := client.EntityClass(hgtp.ArgsEntityClass{
		Entity:        myEntity,
		Authorization: token,
	})
	if err != nil {
		t.Fatal(err)
	}
	if av.Class(class) != av.ClassAdjective {
		t.Fatal(fmt.Errorf("Expected class to be ClassAdjective (393), got '%d'", class))
	}

	category, err := client.EntityCategory(hgtp.ArgsEntityCategory{
		Entity:        myEntity,
		Authorization: token,
	})
	if err != nil {
		t.Fatal(err)
	}
	if av.Category(category) != av.CategoryAirplane {
		t.Fatal(fmt.Errorf("Expected category to be CategoryAirplane (145), got '%d'", category))
	}

	server, err := client.LocalServer(hgtp.ArgsLocalServer{
		Entity:        myEntity,
		Authorization: token,
	})
	if err != nil {
		t.Fatal(err)
	}
	if server != hgtp.NewEntity(0, 0, 0) {
		t.Fatal(fmt.Errorf("Expected server to be <0|0|0>, got '%s'", server))
	}

	redirection, err := client.EntityRedirection(hgtp.ArgsEntityRedirection{
		Entity:        myEntity,
		Authorization: token,
	})
	if err != nil {
		t.Fatal(err)
	}
	if redirection != myEntity {
		t.Fatal(fmt.Errorf("Expected redirection to be myEntity (%s), got '%s'", myEntity, redirection))
	}

	timestamp, err := client.EntityTimestamp(hgtp.ArgsEntityTimestamp{
		Entity:        myEntity,
		Authorization: token,
	})
	if err != nil {
		t.Fatal(err)
	}
	if timestamp == time.Unix(0, 0) { // lol
		t.Fatal(fmt.Errorf("Expected timestamp to not be zero, got '%s'", timestamp))
	}

	available, err := client.EntityAvailable(hgtp.ArgsEntityAvailable{
		Entity:        myEntity,
		Authorization: token,
	})
	if err != nil {
		t.Fatal(err)
	}
	if available != true {
		t.Fatal(fmt.Errorf("Expected available to be true, got '%t'", available))
	}

	extinct, err := client.EntityExtinct(hgtp.ArgsEntityExtinct{
		Entity:        myEntity,
		Authorization: token,
	})
	if err != nil {
		t.Fatal(err)
	}
	if extinct != false {
		t.Fatal(fmt.Errorf("Expected extinct to be false, got '%t'", extinct))
	}

	activated, err := client.EntityExtinct(hgtp.ArgsEntityExtinct{
		Entity:        myEntity,
		Authorization: token,
	})
	if err != nil {
		t.Fatal(err)
	}
	if activated != false {
		t.Fatal(fmt.Errorf("Expected activated to be false, got '%t'", extinct))
	}

	locked, err := client.EntityLocked(hgtp.ArgsEntityLocked{
		Entity:        myEntity,
		Authorization: token,
	})
	if err != nil {
		t.Fatal(err)
	}
	if locked != false {
		t.Fatal(fmt.Errorf("Expected locked to be false, got '%t'", locked))
	}

	pending, err := client.Pending(hgtp.ArgsPending{
		Entity:        myEntity,
		Authorization: token,
	})
	if err != nil {
		t.Fatal(err)
	}
	if pending != 0 {
		t.Fatal(fmt.Errorf("Expected pending to be false, got '%d'", pending))
	}

	rendezvous, err := client.EntityRendezvous(hgtp.ArgsEntityRendezvous{
		Entity:        myEntity,
		Authorization: token,
	})
	if err != nil {
		t.Fatal(err)
	}
	if rendezvous != "NULL" {
		t.Fatal(fmt.Errorf("Expected Rendezvous to be 'NULL', got '%s'", rendezvous))
	}

	references, err := client.EntityReferences(hgtp.ArgsEntityReferences{
		Entity:        myEntity,
		Authorization: token,
	})
	if err != nil {
		t.Fatal(err)
	}
	if references != 1 {
		t.Fatal(fmt.Errorf("Expected References to be 1, got '%d'", references))
	}

	connections, err := client.ConnectionCount(hgtp.ArgsConnectionCount{
		Entity:        myEntity,
		Authorization: token,
	})
	if err != nil {
		t.Fatal(err)
	}
	if connections != 1 { // we're connected to the object adapter
		t.Fatal(fmt.Errorf("Expected Connections to be 1, got '%d'", connections))
	}

	attachments, err := client.AttachmentCount(hgtp.ArgsAttachmentCount{
		Entity:        myEntity,
		Authorization: token,
	})
	if err != nil {
		t.Fatal(err)
	}
	if attachments != 0 {
		t.Fatal(fmt.Errorf("Expected Attachments to be 0, got '%d'", attachments))
	}

	subscriptions, err := client.SubscriptionCount(hgtp.ArgsSubscriptionCount{
		Entity:        myEntity,
		Authorization: token,
	})
	if err != nil {
		t.Fatal(err)
	}
	if subscriptions != 0 {
		t.Fatal(fmt.Errorf("Expected Subscriptions to be 0, got '%d'", subscriptions))
	}

	authorizations, err := client.EntityAuthorizations(hgtp.ArgsEntityAuthorizations{
		Entity:        myEntity,
		Authorization: token,
	})
	if err != nil {
		t.Fatal(err)
	}
	if authorizations != 0 {
		t.Fatal(fmt.Errorf("Expected Authorizations to be 0, got '%d'", authorizations))
	}
}

func TestAdaptInvoke(t *testing.T) {
	client, token := setup(t)
	defer teardown(client, token)

	myEntity, err := client.CreateEntity(hgtp.ArgsCreateEntity{
		Name:          "Adapt invoke entity",
		Authorization: token,
	})
	if err != nil {
		t.Fatal()
	}

	myOutlet, err := client.CreateEntity(hgtp.ArgsCreateEntity{
		Name:          "Adapt invoke outlet",
		Authorization: token,
	})
	if err != nil {
		t.Fatal()
	}

	err = client.ActivateOutlet(hgtp.ArgsActivateOutlet{
		Outlet:        myOutlet,
		Authorization: token,
	})
	if err != nil {
		t.Fatal(fmt.Errorf("Failed to activate outlet"))
	}

	err = client.ConnectMethod(hgtp.ArgsConnectMethod{
		Entity:        myEntity,
		Outlet:        myOutlet,
		Precedence:    4,
		Authorization: token,
	})
	if err != nil {
		t.Fatal(fmt.Errorf("Failed to connect entity to outlet"))
	}

	testfail := make(chan error, 100)

	callback := func(args hgtp.ArgsAdapterCallback) (hgtp.Tag, []byte, error) {
		if args.Ancillary != hgtp.NewEntity(4, 5, 6) {
			testfail <- fmt.Errorf("Expected ancillary to be <4|5|6>, got %v", args.Ancillary)
		}
		if av.Aspect(args.Aspect) != av.AspectFolder {
			testfail <- fmt.Errorf("Expected aspect to be %d, got %d", av.AspectFolder, args.Aspect)
		}
		if args.Attribute != 9 {
			testfail <- fmt.Errorf("Expected attribute to be 4, got %d", args.Attribute)
		}
		if args.Authority != token {
			testfail <- fmt.Errorf("Expected authority to be %s, got %s", token, args.Authority)
		}
		if args.Authorization != token {
			testfail <- fmt.Errorf("Expected authorization to be %s, got %s", token, args.Authorization)
		}
		if args.Auxiliary != hgtp.NewEntity(1, 2, 3) {
			testfail <- fmt.Errorf("Expected auxiliary to be <1|2|3>, got %v", args.Auxiliary)
		}
		if av.Category(args.Category) != av.CategoryAuditorium {
			testfail <- fmt.Errorf("Expected category to be %d, got %d", av.CategoryAuditorium, args.Category)
		}
		if av.Class(args.Class) != av.ClassAffiliation {
			testfail <- fmt.Errorf("Expected class to be %d, got %d", av.ClassAffiliation, args.Class)
		}
		if av.Condition(args.Condition) != av.ConditionHolding {
			testfail <- fmt.Errorf("Expected condition to be %d, got %d", av.ConditionHolding, args.Condition)
		}
		if av.Context(args.Context) != av.ContextAvian {
			testfail <- fmt.Errorf("Expected context to be %d, got %d", av.ContextAvian, args.Context)
		}
		if args.Count != 1023 {
			testfail <- fmt.Errorf("Expected count to be 1023, got %d", args.Count)
		}
		if args.Entity != myEntity {
			testfail <- fmt.Errorf("Expected entity to be %d, got %d", myEntity, args.Entity)
		}
		if av.Event(args.Event) != av.EventExport {
			testfail <- fmt.Errorf("Expected event to be %d, got %d", av.EventExport, args.Event)
		}
		if args.Index != 832 {
			testfail <- fmt.Errorf("Expected index to be 832, got %d", args.Index)
		}
		if args.Instance != 53 {
			testfail <- fmt.Errorf("Expected instance to be 53, got %d", args.Instance)
		}
		if args.Key != "the key" {
			testfail <- fmt.Errorf("Expected key to be 'the key', got %s", args.Key)
		}
		if args.Method != 3 {
			testfail <- fmt.Errorf("Expected method to be 3, got %d", args.Method)
		}
		if av.Mode(args.Mode) != av.ModeReplace {
			testfail <- fmt.Errorf("Expected mode to be %d, got %d", av.ModeReplace, args.Mode)
		}
		if args.Name != "the name" {
			testfail <- fmt.Errorf("Expected name to be 'the name', got %s", args.Name)
		}
		if args.Offset != 38 {
			testfail <- fmt.Errorf("Expected offset to be 38, got %d", args.Offset)
		}
		if args.Parameter != 89 {
			testfail <- fmt.Errorf("Expected parameter to be 89, got %d", args.Parameter)
		}
		if args.Precedence != 4 {
			testfail <- fmt.Errorf("Expected precedence to be 4, got %d", args.Precedence)
		}
		if args.Resultant != 204 {
			testfail <- fmt.Errorf("Expected resultant to be 204, got %d", args.Resultant)
		}
		if av.Tag(args.Tag) != av.TagString {
			testfail <- fmt.Errorf("Expected tag to be %d, got %d", av.TagString, args.Tag)
		}
		val := av.Value{
			Tag:   av.Tag(args.Tag),
			Bytes: args.Bytes,
		}
		if s := val.UnwrapString(); s != "I'm invoking you!" {
			testfail <- fmt.Errorf("Expected tag text to be 'I'm invoking you!, got %s", s)
		}
		if args.Time != time.Unix(98432, 0) {
			testfail <- fmt.Errorf("Expected time to be 98432, got %d", args.Time.Unix())
		}

		return hgtp.Tag(av.TagText), []byte("Hello from golang!"), nil
	}

	go func() {
		err := client.AdaptOutlet(hgtp.ArgsAdaptOutlet{
			Outlet:        myOutlet,
			Authorization: token,
			Callback:      callback,
		})
		if err != nil {
			testfail <- fmt.Errorf("Aadpt failed")
		}
	}()

	// Waiting for the adapt call to reach the server before we invoke
	time.Sleep(100 * time.Millisecond)

	tag, bytes, err := client.InvokeEntity(hgtp.ArgsInvokeEntity{
		Entity:        myEntity,
		Authorization: token,

		Ancillary:  hgtp.NewEntity(4, 5, 6),
		Aspect:     hgtp.Aspect(av.AspectFolder),
		Attribute:  9,
		Authority:  hgtp.MustParseToken("abde2178-d288-4545-a5d2-52dd76924997"),
		Auxiliary:  hgtp.NewEntity(1, 2, 3),
		Category:   hgtp.Category(av.CategoryAuditorium),
		Class:      hgtp.Class(av.ClassAffiliation),
		Condition:  hgtp.Condition(av.ConditionHolding),
		Context:    hgtp.Context(av.ContextAvian),
		Count:      1023,
		Event:      hgtp.Event(av.EventExport),
		Index:      832,
		Instance:   53,
		Key:        "the key",
		Method:     3,
		Mode:       hgtp.Mode(av.ModeReplace),
		Name:       "the name",
		Offset:     38,
		Parameter:  89,
		Precedence: 4,
		Resultant:  204,
		State:      hgtp.State(av.StateDanger),
		Tag:        hgtp.Tag(av.TagString),
		Bytes:      []byte("I'm invoking you!"),
		Time:       time.Unix(98432, 0),
	})

	close(testfail)
	failed := len(testfail) > 0
	for e := range testfail {
		fmt.Println(e)
	}
	if failed {
		t.FailNow()
	}

	if err != nil {
		t.Fatal(fmt.Errorf("Invoke failed"))
	}

	if av.Tag(tag) != av.TagText {
		t.Fatal(fmt.Errorf("Expected tag to be %d, got %d", av.TagText, tag))
	}
	val := av.Value{
		Tag:   av.Tag(tag),
		Bytes: bytes,
	}
	if s := val.UnwrapText(); s != "Hello from golang!" {
		t.Fatal(fmt.Errorf("Expected bytes to be 'Hello from golang!', got %s", s))
	}
}
