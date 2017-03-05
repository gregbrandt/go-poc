import { combineReducers } from 'redux'
import { combineForms } from 'react-redux-form'
import { injectReducer } from '../../store/reducers'

const CreateReducers = (store) => {
  const storyreducer = require('./modules/story').default
  //const storyForm = { name: null, content: null }


  // const unclaimed = combineForms({
  //   inviteManager: initialState,
  // }, 'shared.unclaimed'); // must state exact path of deep reducer


  // const store = createStore(combineReducers({
  //   shared: combineReducers({
  //     storyForm,
  //   }),
  // }), applyMiddleware(createLogger(), thunk));


  const storyForm = ({
    currentstory: { name: null, content: null },
  }, 'story.storyForm'); // must state exact path of deep reducer

  const reducer = combineReducers({ entities: storyreducer, storyForm })
  injectReducer(store, { key: 'story', reducer })
}

export const StoryRoute = (store) => ({
  path: 'story',

  /*  Async getComponent is only invoked when route matches   */
  getComponent(nextState, cb) {
    /*  Webpack - use 'require.ensure' to create a split point
        and embed an async module loader (jsonp) when bundling   */
    require.ensure([], (require) => {
      /*  Webpack - use require callback to define
          dependencies for bundling   */

      const StoryList = require('./components/StoryList').default
      CreateReducers(store)

      /*  Return getComponent   */
      cb(null, StoryList)

      /* Webpack named bundle   */
    }, 'story')
  }
})

export const StoryFormRoute = (store) => {
  return {
    path: 'story/story-form',
    /*  Async getComponent is only invoked when route matches   */
    getComponent(nextState, cb) {
      /*  Webpack - use 'require.ensure' to create a split point
          and embed an async module loader (jsonp) when bundling   */
      require.ensure([], (require) => {
        /*  Webpack - use require callback to define
            dependencies for bundling   */
        const StoryForm = require('./components/StoryForm').default
        CreateReducers(store)
        /*  Return getComponent   */
        cb(null, StoryForm)

        /* Webpack named bundle   */
      }, 'storyform')
    }
  };
}
