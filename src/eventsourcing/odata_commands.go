package eventsourcing

// Open a new account with specified initial balance
type NewOdataCommand struct {
	withGuid
}

func (c NewOdataCommand)Process(o Odata) []Event {
    return []Event{ &NewOdataEvent{}, }
}

type NewOdataPropertyCommand struct {
	withGuid
    PropertyName string
    PropertyValue string
}

func (c NewOdataPropertyCommand)Process(o Odata) []Event {
    return []Event{ &NewOdataPropertyEvent{PropertyName: c.PropertyName, PropertyValue: c.PropertyValue, withGuid: c.withGuid}, }
}

type UpdateOdataPropertyCommand struct {
	withGuid
    PropertyName string
    PropertyValue string
}

func (c UpdateOdataPropertyCommand)Process(o Odata) []Event {
    return []Event{ &UpdatedOdataPropertyEvent{PropertyName: c.PropertyName, PropertyValue: c.PropertyValue, withGuid: c.withGuid}, }
}

