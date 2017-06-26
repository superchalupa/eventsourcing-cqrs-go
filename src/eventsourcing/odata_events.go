package eventsourcing

type Applier interface{
    Apply(*Odata)
}

// An account was opened with given initial balance
type NewOdataEvent struct {
	withGuid
}

func (e *NewOdataEvent) Apply(o *Odata){
    // do nothing
}

type NewOdataPropertyEvent struct {
	withGuid
    PropertyName string
    PropertyValue string
}

func (e *NewOdataPropertyEvent) Apply(o *Odata){
    o.properties[e.PropertyName] = e.PropertyValue
}

type UpdatedOdataPropertyEvent struct {
	withGuid
    PropertyName string
    PropertyValue string
}

func (e *UpdatedOdataPropertyEvent) Apply(o *Odata){
    o.properties[e.PropertyName] = e.PropertyValue
}
