 package readmodel

// import ( eh "github.com/looplab/eventhorizon"
// )

// // Story is a read model object for an Story.
// type Story struct {
// 	ID eh.UUID

// 	Name string

// 	Content string

// 	Status string
// }

// // StoryProjector is a projector that updates the Storys.
// type StoryProjector struct {
// 	repository eh.ReadRepository
// }

// // NewStoryProjector creates a new StoryProjector.
// func NewStoryProjector(repository eh.ReadRepository) *StoryProjector {

// 	p := &StoryProjector{

// 		repository: repository,
// 	}

// 	return p
// }

// // HandlerType implements the HandlerType method of the EventHandler interface.
// func (p *StoryProjector) HandlerType() eh.EventHandlerType {

// 	return eh.EventHandlerType("StoryProjector")

// }

// // HandleEvent implements the HandleEvent method of the EventHandler interface.
// func (p *StoryProjector) HandleEvent(event eh.Event) {

// 	switch event := event.(type) {

// 	case *StoryCreated:

// 		i := &Story{

// 			ID: event.StoryId,

// 			Name: event.Name,

// 			Content: event.Content,

// 			Status: "created",
// 		}

// 		p.repository.Save(i.ID, i)

// 	case *StoryAccepted:

// 		m, _ := p.repository.Find(event.StoryId)

// 		i := m.(*Story)

// 		i.Status = "approved"

// 		p.repository.Save(i.ID, i)

// 	}

// }
