package eventsourcing

/*
// common properties for all customer facing services
type Service struct {
	commandChannel chan Command
	store EventStore
}

// Getter for command channel - will allow others to post commands
func (a Service) CommandChannel() chan<- Command {
	return a.commandChannel
}
*/

// Odata Service - allows simple Odata management (open, credit, debit)
type OdataService struct {
	Service
}

// Create Odata service - initialize command channel, and let it use passed event store
func NewOdataService(store EventStore) *OdataService{
	acc := &OdataService{
		Service:Service{
			commandChannel:make(chan Command),
			store:store,
		},
	}
	return acc
}

// Reads from command channel,
// restores an aggregate,
// processes the command and
// persists received events.
// This method *blocks* until command is available,
// therefore should run in a goroutine
func (a *OdataService) HandleCommands() {
	for {
		c := <- a.commandChannel
		acc := RestoreOdata(c.GetGuid(), a.store)
		a.store.Update(c.GetGuid(), acc.Version, acc.processCommand(c))

	}
}

// Open a new Odata
// Returns Odatas GUID
func (a OdataService) NewOdata(oid string) guid {
    thisguid := guid(oid)
	c := &NewOdataCommand{
		withGuid:withGuid{Guid: thisguid},
	}
	a.commandChannel <- c
	return thisguid
}

// New property
func (a OdataService) NewOdataProperty(oid string, p string, v string) {
    thisguid := guid(oid)
    c := &NewOdataPropertyEvent{PropertyName: p, PropertyValue: v,
		withGuid:withGuid{Guid:thisguid},
	}
	a.commandChannel <- c
}

