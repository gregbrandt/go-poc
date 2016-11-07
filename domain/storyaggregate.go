package domain

import (
	"fmt"

	eh "github.com/looplab/eventhorizon"

)

func init() {

	eh.RegisterAggregate(func(id eh.UUID) eh.Aggregate {

		return NewStoryAggregate(id)

	})

}

// StoryAggregateType is the type name of the aggregate.
const StoryAggregateType eh.AggregateType = "Story"

// StoryAggregate is the root aggregate.

//

// The aggregate root will guard that the invitation can only be accepted OR

// declined, but not both.

type StoryAggregate struct {

	// AggregateBase implements most of the eventhorizon.Aggregate interface.

	*eh.AggregateBase

	name string

	content string

	status int
}

// NewStoryAggregate creates a new InvitationAggregate with an ID.
func NewStoryAggregate(id eh.UUID) *StoryAggregate {

	return &StoryAggregate{

		AggregateBase: eh.NewAggregateBase(id),
	}

}

// AggregateType implements the AggregateType method of the Aggregate interface.
func (i *StoryAggregate) AggregateType() eh.AggregateType {

	return StoryAggregateType

}

// HandleCommand implements the HandleCommand method of the Aggregate interface.
func (i *StoryAggregate) HandleCommand(command eh.Command) error {

	switch command := command.(type) {

	case *CreateStory:

		i.StoreEvent(&StoryCreated{command.StoryId, command.Name, command.Content})

		return nil

	case *AcceptStory:

		return nil

	}

	return fmt.Errorf("couldn't handle command")

}

// ApplyEvent implements the ApplyEvent method of the Aggregate interface.
func (i *StoryAggregate) ApplyEvent(event eh.Event) {

	switch event := event.(type) {

	case *StoryCreated:

		i.name = event.Name

		i.content = event.Content

	case *StoryAccepted:

		i.status = 2
    }
}
