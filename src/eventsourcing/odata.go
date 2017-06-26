package eventsourcing
import (
	"fmt"
	"gopkg.in/yaml.v2"
)

// An aggregate implementation representing odata
type Odata struct {
	baseAggregate
	properties map[string]interface{}
}

// Make sure it implements Aggregate
var _ Aggregate = (*Odata)(nil)

type Processor interface {
    Process(Odata) []Event
}

// @see Aggregate.applyEvents
func (a *Odata) applyEvents(events []Event) {
	for _, e := range events {
		switch event := e.(type){
        case Applier:
            event.Apply(a)
		default:
			panic(fmt.Sprintf("Unknown event %#v", event))
		}
	}
	a.Version = len(events)
}

// @see Aggregate.processCommand
func (a Odata) processCommand(command Command) []Event {
	var events []Event
	switch c := command.(type){
    case Processor:
        events = c.Process(a)
	default:
		panic(fmt.Sprintf("Unknown command %#v", c))
	}
    for _, event := range(events){
	    event.SetGuid(command.GetGuid())
    }
	return events
}


// Helper function to restore odata according to persisted state in event store
func RestoreOdata(guid guid, store EventStore) *Odata {
	a:= NewOdata()
	RestoreAggregate(guid, a, store)
	return a
}

// create new Odata in an initial state
func NewOdata() *Odata {
	return &Odata{properties: make(map[string]interface{}) }
}

// pretty print in YAML
func (a Odata) String() string {
	yaml, _ := yaml.Marshal(&a)
	return fmt.Sprintf("Odata:\n%v", string(yaml))
}

