package domain

import( 
	eh "github.com/looplab/eventhorizon"
)

func init() {

		eh.RegisterEvent(func() eh.Event { return &StoryCreated{} })
		eh.RegisterEvent(func() eh.Event { return &StoryAccepted{} })

}

const (
	StoryCreatedEvent eh.EventType = "StoryCreated"
	StoryAcceptedEvent eh.EventType = "StoryAccepted"
)

// StoryCreated is an event for when an Story has been created.

type StoryCreated struct {
	StoryId eh.UUID `bson:"story_id"`

	Name string `bson:"name"`

	Content string `bson:"content"`
}

func (c StoryCreated) AggregateID() eh.UUID { return c.StoryId }

func (c StoryCreated) AggregateType() eh.AggregateType { return StoryAggregateType }

func (c StoryCreated) EventType() eh.EventType { return StoryCreatedEvent }

// StoryAccepted is an event for when an Story has been accepted.

type StoryAccepted struct {
	StoryId eh.UUID `bson:"story_id"`
}

func (c StoryAccepted) AggregateID() eh.UUID { return c.StoryId }

func (c StoryAccepted) AggregateType() eh.AggregateType { return StoryAggregateType }

func (c StoryAccepted) EventType() eh.EventType { return StoryAcceptedEvent }
