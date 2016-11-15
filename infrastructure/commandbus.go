package infrastructure

import (
    "fmt"
	eh "github.com/looplab/eventhorizon"
    ehcommandbus "github.com/looplab/eventhorizon/commandbus/local"
	eventstore "github.com/looplab/eventhorizon/eventstore/mongodb"
	eventbus "github.com/looplab/eventhorizon/eventbus/local"
    "github.com/gregbrandt/Go-POC/domain"
    "os"
)

var commandbus *ehcommandbus.CommandBus
var eventBus *eventbus.EventBus


func init(){
    host := os.Getenv("MONGO_PORT_27017_TCP_ADDR")
	port := os.Getenv("MONGO_PORT_27017_TCP_PORT")

	url := "localhost"
	if host != "" && port != "" {
		url = host + ":" + port
	}

	// Create the event store.
	eventStore, err := eventstore.NewEventStore(url, "demo")
    
	// Create the event bus that distributes events.
	eventBus = eventbus.NewEventBus()
	eventBus.AddObserver(&EventBusLogger{})

	// Create the aggregate repository.
	repository, err := eh.NewEventSourcingRepository(eventStore, eventBus)
	if err != nil {
		panic(fmt.Errorf("could not create repository: %s", err))
	}

	// Create the aggregate command handler.
	handler, err := eh.NewAggregateCommandHandler(repository)
	if err != nil {
		panic(fmt.Errorf("could not create command handler: %s", err))
	}

	err = handler.SetAggregate(domain.StoryAggregateType, domain.CreateStoryCommand)
    
	if err != nil {
        panic(fmt.Errorf("could not set command handler aggregate: %s \n", err))
	}
	err = handler.SetAggregate(domain.StoryAggregateType, domain.AcceptStoryCommand)
    
	if err != nil {
		panic(fmt.Errorf("could not set command handler aggregate: %s", err))
	}

    commandbus = ehcommandbus.NewCommandBus()

	err = commandbus.SetHandler(handler, domain.CreateStoryCommand)
    
	if err != nil {
		panic(fmt.Errorf("could not set command handler: %s", err))
	}

	err = commandbus.SetHandler(handler, domain.AcceptStoryCommand)

	if err != nil {
		panic(fmt.Errorf("could not set command handler: %s", err))
	}

}


type EventBusLogger struct{}

func (l *EventBusLogger) Notify(event eh.Event) {
}

func GetCommandBus() *ehcommandbus.CommandBus {
	return commandbus
}

func GetEventBus() *eventbus.EventBus {
	return eventBus
}