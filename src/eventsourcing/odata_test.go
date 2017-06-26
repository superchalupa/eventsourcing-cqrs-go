package eventsourcing
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var _ = assert.Equal

func TestOdataRestore(t *testing.T) {
	oid := NewOdata()
	oid.applyEvents([]Event{
        &NewOdataEvent{},
        &NewOdataPropertyEvent{PropertyName: "test1", PropertyValue: "foobar1"},
        &NewOdataPropertyEvent{PropertyName: "test2", PropertyValue: "foobar2"},
        &UpdatedOdataPropertyEvent{PropertyName: "test2", PropertyValue: "foobar-update"},
	})
	assert.Equal(t, "foobar1", oid.properties["test1"])
	assert.Equal(t, "foobar-update", oid.properties["test2"])
}

func TestOdataCommand(t *testing.T) {
	oid := NewOdata()
	e := oid.processCommand(&NewOdataCommand{})
	assert.Equal(t, []Event{&NewOdataEvent{}}, e)

	oid.applyEvents(e)
}
