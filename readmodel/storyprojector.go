 package readmodel

import (
     eh "github.com/looplab/eventhorizon"
	story "github.com/gregbrandt/Go-POC/domain"
	readrepository "github.com/looplab/eventhorizon/readrepository/mongodb"
    "github.com/gregbrandt/Go-POC/infrastructure"
    "os"
)

// Story is a read model object for an Story.
type Story struct {
	ID eh.UUID

	Name string

	Content string

	Status string
}

// StoryProjector is a projector that updates the Storys.
type StoryProjector struct {
	repository eh.ReadRepository
}


func init(){
    host := os.Getenv("MONGO_PORT_27017_TCP_ADDR")
	port := os.Getenv("MONGO_PORT_27017_TCP_PORT")

	url := "localhost"
	if host != "" && port != "" {
		url = host + ":" + port
	}
    eventBus := infrastructure.GetEventBus()
	repository,_ := readrepository.NewReadRepository(url, "demo", "story")
	projector := NewStoryProjector(repository)
	eventBus.AddHandler(projector, story.StoryCreatedEvent)
}


// NewStoryProjector creates a new StoryProjector.
func NewStoryProjector(repository eh.ReadRepository) *StoryProjector {

	p := &StoryProjector{

		repository: repository,
	}

	return p
}

// HandlerType implements the HandlerType method of the EventHandler interface.
func (p *StoryProjector) HandlerType() eh.EventHandlerType {

	return eh.EventHandlerType("StoryProjector")

}

// HandleEvent implements the HandleEvent method of the EventHandler interface.
func (p *StoryProjector) HandleEvent(event eh.Event) {

	switch event := event.(type) {

	case *story.StoryCreated:

		i := &Story{

			ID: event.StoryId,

			Name: event.Name,

			Content: event.Content,

			Status: "created",
		}

		p.repository.Save(i.ID, i)

	case *story.StoryAccepted:

		m, _ := p.repository.Find(event.StoryId)

		i := m.(*Story)

		i.Status = "approved"

		p.repository.Save(i.ID, i)

	}

}
