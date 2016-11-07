package domain

import (
	eh "github.com/looplab/eventhorizon"
)

const (
	CreateStoryCommand eh.CommandType = "CreateStory"

	AcceptStoryCommand eh.CommandType = "AcceptStory"
)

// CreateStory is a command for creating Stories.

type CreateStory struct {
	StoryId eh.UUID

	Name string

	Content string `eh:"optional"`
}

func (c CreateStory) AggregateID() eh.UUID { return c.StoryId }

func (c CreateStory) AggregateType() eh.AggregateType { return StoryAggregateType }

func (c CreateStory) CommandType() eh.CommandType { return CreateStoryCommand }

// AcceptStory is a command for accepting Stories.
type AcceptStory struct {
	StoryId eh.UUID
}

func (c AcceptStory) AggregateID() eh.UUID { return c.StoryId }

func (c AcceptStory) AggregateType() eh.AggregateType { return StoryAggregateType }

func (c AcceptStory) CommandType() eh.CommandType { return AcceptStoryCommand }
