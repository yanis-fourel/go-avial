package hgtp

func (args *ArgsCreateEntity) toHGTPMessage() (hgtpMessage, error) {
	var msg hgtpMessage

	msg.PackCommand(CommandCreate)
	msg.PackOutlet(args.Outlet)
	msg.PackAncillary(args.Server)
	err := msg.PackName(args.Name)
	if err != nil {
		return msg, err
	}
	err = msg.PackKey(args.Key)
	if err != nil {
		return msg, err
	}
	msg.PackContext(args.Context)
	msg.PackCategory(args.Category)
	msg.PackClass(args.Class)
	msg.PackMethod(args.Method)
	msg.PackAttribute(args.Attribute)
	msg.PackEvent(args.Event)
	msg.PackPrecedence(args.Precedence)
	msg.PackTimeout(args.Timeout)
	msg.PackAuthorization(args.Authorization)

	return msg, nil
}

func (args *ArgsDeleteEntity) toHGTPMessage() hgtpMessage {
	var msg hgtpMessage

	msg.PackCommand(CommandDelete)
	msg.PackEntity(args.Entity)
	msg.PackAuthorization(args.Authorization)

	return msg
}

func (args *ArgsInvokeEntity) toHGTPMessage() (hgtpMessage, error) {
	var msg hgtpMessage

	msg.PackCommand(CommandInvoke)
	msg.PackEntity(args.Entity)
	msg.PackAuxiliary(args.Auxiliary)
	msg.PackAncillary(args.Ancillary)
	msg.PackMethod(args.Method)
	msg.PackAttribute(args.Attribute)
	msg.PackInstance(args.Instance)
	msg.PackOffset(args.Offset)
	err := msg.PackName(args.Name)
	if err != nil {
		return msg, err
	}
	err = msg.PackKey(args.Key)
	if err != nil {
		return msg, err
	}
	msg.PackBytes(args.Bytes)
	msg.PackParameter(args.Parameter)
	msg.PackResultant(args.Resultant)
	msg.PackIndex(args.Index)
	msg.PackCount(args.Count)
	msg.PackAspect(args.Aspect)
	msg.PackContext(args.Context)
	msg.PackCategory(args.Category)
	msg.PackClass(args.Class)
	msg.PackEvent(args.Event)
	msg.PackMode(args.Mode)
	msg.PackState(args.State)
	msg.PackCondition(args.Condition)
	msg.PackPrecedence(args.Precedence)
	msg.PackTag(args.Tag)
	msg.PackTimeout(args.Timeout)
	msg.PackAuthority(args.Authority)
	msg.PackAuthorization(args.Authorization)
	msg.PackTime(args.Time)

	return msg, nil
}

func (args *ArgsInquireEntity) toHGTPMessage() (hgtpMessage, error) {
	var msg hgtpMessage

	msg.PackCommand(CommandInquire)
	msg.PackEntity(args.Entity)
	msg.PackAuxiliary(args.Auxiliary)
	msg.PackAncillary(args.Ancillary)
	msg.PackAttribute(args.Attribute)
	msg.PackInstance(args.Instance)
	msg.PackOffset(args.Offset)
	err := msg.PackName(args.Name)
	if err != nil {
		return msg, err
	}
	err = msg.PackKey(args.Key)
	if err != nil {
		return msg, err
	}
	msg.PackBytes(args.Bytes)
	msg.PackParameter(args.Parameter)
	msg.PackResultant(args.Resultant)
	msg.PackIndex(args.Index)
	msg.PackCount(args.Count)
	msg.PackAspect(args.Aspect)
	msg.PackContext(args.Context)
	msg.PackCategory(args.Category)
	msg.PackClass(args.Class)
	msg.PackEvent(args.Event)
	msg.PackMode(args.Mode)
	msg.PackState(args.State)
	msg.PackCondition(args.Condition)
	msg.PackPrecedence(args.Precedence)
	msg.PackTag(args.Tag)
	msg.PackTime(args.Time)
	msg.PackTimeout(args.Timeout)
	msg.PackAuthority(args.Authority)
	msg.PackAuthorization(args.Authorization)

	return msg, nil
}

func (args *ArgsReferenceEntity) toHGTPMessage() hgtpMessage {
	var msg hgtpMessage

	msg.PackCommand(CommandReference)
	msg.PackEntity(args.Entity)
	msg.PackCount(args.Count)
	msg.PackAuthorization(args.Authorization)

	return msg
}

func (args *ArgsDereferenceEntity) toHGTPMessage() hgtpMessage {
	var msg hgtpMessage

	msg.PackCommand(CommandDereference)
	msg.PackEntity(args.Entity)
	msg.PackCount(args.Count)
	msg.PackAuthorization(args.Authorization)

	return msg
}

func (args *ArgsConnectMethod) toHGTPMessage() hgtpMessage {
	var msg hgtpMessage

	msg.PackCommand(CommandConnect)
	msg.PackEntity(args.Entity)
	msg.PackOutlet(args.Outlet)
	msg.PackMethod(args.Method)
	msg.PackPrecedence(args.Precedence)
	msg.PackTimeout(args.Timeout)
	msg.PackAuthorization(args.Authorization)

	return msg
}

func (args *ArgsDisconnectMethod) toHGTPMessage() hgtpMessage {
	var msg hgtpMessage

	msg.PackCommand(CommandDisconnect)
	msg.PackEntity(args.Entity)
	msg.PackOutlet(args.Outlet)
	msg.PackMethod(args.Method)
	msg.PackPrecedence(args.Precedence)
	msg.PackAuthorization(args.Authorization)

	return msg
}

func (args *ArgsEntityConnection) toHGTPMessage() hgtpMessage {
	var msg hgtpMessage

	msg.PackCommand(CommandConnection)
	msg.PackEntity(args.Entity)
	msg.PackIndex(args.Index)
	msg.PackAuthorization(args.Authorization)

	return msg
}

func (args *ArgsAttachAttribute) toHGTPMessage() hgtpMessage {
	var msg hgtpMessage

	msg.PackCommand(CommandAttach)
	msg.PackEntity(args.Entity)
	msg.PackOutlet(args.Outlet)
	msg.PackAttribute(args.Attribute)
	msg.PackPrecedence(args.Precedence)
	msg.PackTimeout(args.Timeout)
	msg.PackAuthorization(args.Authorization)

	return msg
}

func (args *ArgsDetachAttribute) toHGTPMessage() hgtpMessage {
	var msg hgtpMessage

	msg.PackCommand(CommandDetach)
	msg.PackEntity(args.Entity)
	msg.PackOutlet(args.Outlet)
	msg.PackAttribute(args.Attribute)
	msg.PackPrecedence(args.Precedence)
	msg.PackAuthorization(args.Authorization)

	return msg
}

func (args *ArgsEntityAttachment) toHGTPMessage() hgtpMessage {
	var msg hgtpMessage

	msg.PackCommand(CommandAttachment)
	msg.PackEntity(args.Entity)
	msg.PackIndex(args.Index)
	msg.PackAuthorization(args.Authorization)

	return msg
}

func (args *ArgsActivateOutlet) toHGTPMessage() hgtpMessage {
	var msg hgtpMessage

	msg.PackCommand(CommandActivate)
	msg.PackOutlet(args.Outlet)
	msg.PackAuthorization(args.Authorization)

	return msg
}

func (args *ArgsDeactivateOutlet) toHGTPMessage() hgtpMessage {
	var msg hgtpMessage

	msg.PackCommand(CommandDeactivate)
	msg.PackOutlet(args.Outlet)
	msg.PackAuthorization(args.Authorization)

	return msg
}

func (args *ArgsAdaptOutlet) toHGTPMessage() hgtpMessage {
	var msg hgtpMessage

	msg.PackCommand(CommandAdapt)
	msg.PackOutlet(args.Outlet)
	msg.PackTimeout(args.Timeout)
	msg.PackAuthorization(args.Authorization)

	return msg
}

func NewArgsAdapterCallback(msg hgtpMessage) ArgsAdapterCallback {
	return ArgsAdapterCallback{
		Entity:        msg.UnpackEntity(),
		Method:        msg.UnpackMethod(),
		Auxiliary:     msg.UnpackAuxiliary(),
		Ancillary:     msg.UnpackAncillary(),
		Attribute:     msg.UnpackAttribute(),
		Instance:      msg.UnpackInstance(),
		Offset:        msg.UnpackOffset(),
		Name:          msg.UnpackName(),
		Key:           msg.UnpackKey(),
		Parameter:     msg.UnpackParameter(),
		Resultant:     msg.UnpackResultant(),
		Index:         msg.UnpackIndex(),
		Count:         msg.UnpackCount(),
		Aspect:        msg.UnpackAspect(),
		Context:       msg.UnpackContext(),
		Category:      msg.UnpackCategory(),
		Class:         msg.UnpackClass(),
		Event:         msg.UnpackEvent(),
		Mode:          msg.UnpackMode(),
		State:         msg.UnpackState(),
		Condition:     msg.UnpackCondition(),
		Precedence:    msg.UnpackPrecedence(),
		Tag:           msg.UnpackTag(),
		Bytes:         msg.UnpackBytes(),
		Time:          msg.UnpackTime(),
		Timeout:       msg.UnpackTimeout(),
		Authority:     msg.UnpackAuthority(),
		Authorization: msg.UnpackAuthorization(),
	}
}

func (args *ArgsResetOutlet) toHGTPMessage() hgtpMessage {
	var msg hgtpMessage

	msg.PackCommand(CommandReset)
	msg.PackOutlet(args.Outlet)
	msg.PackAuthorization(args.Authorization)

	return msg
}

func (args *ArgsPublishEvent) toHGTPMessage() (hgtpMessage, error) {
	var msg hgtpMessage

	msg.PackCommand(CommandInvoke)
	msg.PackEntity(args.Entity)
	msg.PackAuxiliary(args.Auxiliary)
	msg.PackAncillary(args.Ancillary)
	msg.PackMethod(args.Method)
	msg.PackAttribute(args.Attribute)
	msg.PackInstance(args.Instance)
	msg.PackOffset(args.Offset)
	err := msg.PackName(args.Name)
	if err != nil {
		return msg, err
	}
	err = msg.PackKey(args.Key)
	if err != nil {
		return msg, err
	}
	msg.PackBytes(args.Bytes)
	msg.PackParameter(args.Parameter)
	msg.PackResultant(args.Resultant)
	msg.PackIndex(args.Index)
	msg.PackCount(args.Count)
	msg.PackAspect(args.Aspect)
	msg.PackContext(args.Context)
	msg.PackCategory(args.Category)
	msg.PackClass(args.Class)
	msg.PackEvent(args.Event)
	msg.PackMode(args.Mode)
	msg.PackState(args.State)
	msg.PackCondition(args.Condition)
	msg.PackPrecedence(args.Precedence)
	msg.PackTag(args.Tag)
	msg.PackTimeout(args.Timeout)
	msg.PackAuthority(args.Authority)
	msg.PackAuthorization(args.Authorization)

	return msg, nil
}

func (args *ArgsSubscribeEvent) toHGTPMessage() hgtpMessage {
	var msg hgtpMessage

	msg.PackCommand(CommandSubscribe)
	msg.PackEntity(args.Entity)
	msg.PackOutlet(args.Outlet)
	msg.PackEvent(args.Event)
	msg.PackPrecedence(args.Precedence)
	msg.PackTimeout(args.Timeout)
	msg.PackAuthorization(args.Authorization)

	return msg
}

func (args *ArgsUnsubscribeEvent) toHGTPMessage() hgtpMessage {
	var msg hgtpMessage

	msg.PackCommand(CommandUnsubscribe)
	msg.PackEntity(args.Entity)
	msg.PackOutlet(args.Outlet)
	msg.PackEvent(args.Event)
	msg.PackPrecedence(args.Precedence)
	msg.PackAuthorization(args.Authorization)

	return msg
}

func (args *ArgsWaitEvent) toHGTPMessage() hgtpMessage {
	var msg hgtpMessage

	msg.PackCommand(CommandWait)
	msg.PackOutlet(args.Outlet)
	msg.PackTimeout(args.Timeout)
	msg.PackAuthorization(args.Authorization)

	return msg
}

func NewArgsWaitCallback(msg hgtpMessage) ArgsWaitCallback {
	return ArgsWaitCallback{
		Entity:        msg.UnpackEntity(),
		Method:        msg.UnpackMethod(),
		Auxiliary:     msg.UnpackAuxiliary(),
		Ancillary:     msg.UnpackAncillary(),
		Attribute:     msg.UnpackAttribute(),
		Instance:      msg.UnpackInstance(),
		Offset:        msg.UnpackOffset(),
		Name:          msg.UnpackName(),
		Key:           msg.UnpackKey(),
		Parameter:     msg.UnpackParameter(),
		Resultant:     msg.UnpackResultant(),
		Index:         msg.UnpackIndex(),
		Count:         msg.UnpackCount(),
		Aspect:        msg.UnpackAspect(),
		Context:       msg.UnpackContext(),
		Category:      msg.UnpackCategory(),
		Class:         msg.UnpackClass(),
		Event:         msg.UnpackEvent(),
		Mode:          msg.UnpackMode(),
		State:         msg.UnpackState(),
		Condition:     msg.UnpackCondition(),
		Precedence:    msg.UnpackPrecedence(),
		Tag:           msg.UnpackTag(),
		Bytes:         msg.UnpackBytes(),
		Time:          msg.UnpackTime(),
		Timeout:       msg.UnpackTimeout(),
		Authority:     msg.UnpackAuthority(),
		Authorization: msg.UnpackAuthorization(),
	}
}

func (args *ArgsEntitySubscription) toHGTPMessage() hgtpMessage {
	var msg hgtpMessage

	msg.PackCommand(CommandSubscription)
	msg.PackEntity(args.Entity)
	msg.PackIndex(args.Index)
	msg.PackAuthorization(args.Authorization)

	return msg
}

func (args *ArgsLockOutlet) toHGTPMessage() hgtpMessage {
	var msg hgtpMessage

	msg.PackCommand(CommandLock)
	msg.PackOutlet(args.Outlet)
	msg.PackTimeout(args.Timeout)
	msg.PackAuthorization(args.Authorization)

	return msg
}

func (args *ArgsUnlockOutlet) toHGTPMessage() hgtpMessage {
	var msg hgtpMessage

	msg.PackCommand(CommandUnlock)
	msg.PackOutlet(args.Outlet)
	msg.PackAuthorization(args.Authorization)

	return msg
}

func (args *ArgsAuthorizeEntity) toHGTPMessage() hgtpMessage {
	var msg hgtpMessage

	msg.PackCommand(CommandAuthorize)
	msg.PackEntity(args.Entity)
	msg.PackAuthority(args.Authority)
	msg.PackParameter(args.Restriction)
	msg.PackAuthorization(args.Authorization)

	return msg
}

func (args *ArgsDeauthorizeEntity) toHGTPMessage() hgtpMessage {
	var msg hgtpMessage

	msg.PackCommand(CommandDeauthorize)
	msg.PackEntity(args.Entity)
	msg.PackAuthority(args.Authority)
	msg.PackAuthorization(args.Authorization)

	return msg
}

func (args *ArgsEntityAuthorized) toHGTPMessage() hgtpMessage {
	var msg hgtpMessage

	msg.PackCommand(CommandAuthorized)
	msg.PackEntity(args.Entity)
	msg.PackAuthority(args.Authority)
	msg.PackAuthorization(args.Authorization)

	return msg
}

func (args *ArgsEntityAuthority) toHGTPMessage() hgtpMessage {
	var msg hgtpMessage

	msg.PackCommand(CommandReport)
	msg.PackReport(ReportAuthority)
	msg.PackEntity(args.Entity)
	msg.PackAuthorization(args.Authorization)

	return msg
}

func (args *ArgsEntityName) toHGTPMessage() hgtpMessage {
	var msg hgtpMessage

	msg.PackCommand(CommandReport)
	msg.PackReport(ReportName)
	msg.PackEntity(args.Entity)
	msg.PackAuthorization(args.Authorization)

	return msg
}

func (args *ArgsEntityClass) toHGTPMessage() hgtpMessage {
	var msg hgtpMessage

	msg.PackCommand(CommandReport)
	msg.PackReport(ReportClass)
	msg.PackEntity(args.Entity)
	msg.PackAuthorization(args.Authorization)

	return msg
}

func (args *ArgsEntityCategory) toHGTPMessage() hgtpMessage {
	var msg hgtpMessage

	msg.PackCommand(CommandReport)
	msg.PackReport(ReportCategory)
	msg.PackEntity(args.Entity)
	msg.PackAuthorization(args.Authorization)

	return msg
}

func (args *ArgsLocalServer) toHGTPMessage() hgtpMessage {
	var msg hgtpMessage

	msg.PackCommand(CommandReport)
	msg.PackReport(ReportServer)
	msg.PackEntity(args.Entity)
	msg.PackAuthorization(args.Authorization)

	return msg
}

func (args *ArgsEntityRedirection) toHGTPMessage() hgtpMessage {
	var msg hgtpMessage

	msg.PackCommand(CommandReport)
	msg.PackReport(ReportRedirection)
	msg.PackEntity(args.Entity)
	msg.PackAuthorization(args.Authorization)

	return msg
}

func (args *ArgsChangeEntity) toHGTPMessage() (hgtpMessage, error) {
	var msg hgtpMessage

	msg.PackCommand(CommandChange)
	msg.PackEntity(args.Entity)
	err := msg.PackName(args.Name)
	if err != nil {
		return msg, err
	}
	err = msg.PackKey(args.Key)
	if err != nil {
		return msg, err
	}
	msg.PackContext(args.Context)
	msg.PackCategory(args.Category)
	msg.PackClass(args.Class)
	msg.PackState(args.State)
	msg.PackAuthority(args.Authority)
	msg.PackAuthorization(args.Authorization)

	return msg, nil
}

func (args *ArgsEntityTimestamp) toHGTPMessage() hgtpMessage {
	var msg hgtpMessage

	msg.PackCommand(CommandReport)
	msg.PackReport(ReportTimestamp)
	msg.PackEntity(args.Entity)
	msg.PackAuthorization(args.Authorization)

	return msg
}

func (args *ArgsEntityAvailable) toHGTPMessage() hgtpMessage {
	var msg hgtpMessage

	msg.PackCommand(CommandReport)
	msg.PackReport(ReportAvailable)
	msg.PackEntity(args.Entity)
	msg.PackAuthorization(args.Authorization)

	return msg
}

func (args *ArgsEntityExtinct) toHGTPMessage() hgtpMessage {
	var msg hgtpMessage

	msg.PackCommand(CommandReport)
	msg.PackReport(ReportExtinct)
	msg.PackEntity(args.Entity)
	msg.PackAuthorization(args.Authorization)

	return msg
}

func (args *ArgsEntityActivated) toHGTPMessage() hgtpMessage {
	var msg hgtpMessage

	msg.PackCommand(CommandReport)
	msg.PackReport(ReportActivated)
	msg.PackEntity(args.Entity)
	msg.PackAuthorization(args.Authorization)

	return msg
}

func (args *ArgsEntityLocked) toHGTPMessage() hgtpMessage {
	var msg hgtpMessage

	msg.PackCommand(CommandReport)
	msg.PackReport(ReportLocked)
	msg.PackEntity(args.Entity)
	msg.PackAuthorization(args.Authorization)

	return msg
}

func (args *ArgsPending) toHGTPMessage() hgtpMessage {
	var msg hgtpMessage

	msg.PackCommand(CommandReport)
	msg.PackReport(ReportPending)
	msg.PackEntity(args.Entity)
	msg.PackAuthorization(args.Authorization)

	return msg
}

func (args *ArgsEntityRendezvous) toHGTPMessage() hgtpMessage {
	var msg hgtpMessage

	msg.PackCommand(CommandReport)
	msg.PackReport(ReportRendezvous)
	msg.PackEntity(args.Entity)
	msg.PackAuthorization(args.Authorization)

	return msg
}

func (args *ArgsEntityReferences) toHGTPMessage() hgtpMessage {
	var msg hgtpMessage

	msg.PackCommand(CommandReport)
	msg.PackReport(ReportReferences)
	msg.PackEntity(args.Entity)
	msg.PackAuthorization(args.Authorization)

	return msg
}

func (args *ArgsConnectionCount) toHGTPMessage() hgtpMessage {
	var msg hgtpMessage

	msg.PackCommand(CommandReport)
	msg.PackReport(ReportConnections)
	msg.PackEntity(args.Entity)
	msg.PackAuthorization(args.Authorization)

	return msg
}

func (args *ArgsAttachmentCount) toHGTPMessage() hgtpMessage {
	var msg hgtpMessage

	msg.PackCommand(CommandReport)
	msg.PackReport(ReportAttachments)
	msg.PackEntity(args.Entity)
	msg.PackAuthorization(args.Authorization)

	return msg
}

func (args *ArgsSubscriptionCount) toHGTPMessage() hgtpMessage {
	var msg hgtpMessage

	msg.PackCommand(CommandReport)
	msg.PackReport(ReportSubscriptions)
	msg.PackEntity(args.Entity)
	msg.PackAuthorization(args.Authorization)

	return msg
}

func (args *ArgsEntityAuthorizations) toHGTPMessage() hgtpMessage {
	var msg hgtpMessage

	msg.PackCommand(CommandReport)
	msg.PackReport(ReportAuthorizations)
	msg.PackEntity(args.Entity)
	msg.PackAuthorization(args.Authorization)

	return msg
}
