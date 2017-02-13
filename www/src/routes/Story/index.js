import { injectReducer } from '../../store/reducers'

export default (store) => ({
  path: 'story',
  childRoutes: [
    CreateStoryFormRoute(store)
  ],

  /*  Async getComponent is only invoked when route matches   */
  getComponent(nextState, cb) {
    /*  Webpack - use 'require.ensure' to create a split point
        and embed an async module loader (jsonp) when bundling   */
    require.ensure([], (require) => {
      /*  Webpack - use require callback to define
          dependencies for bundling   */
      const StoryList = require('./containers/StoryListContainer').default
      const reducer = require('./modules/story').default

      /*  Add the reducer to the store on key 'counter'  */
      injectReducer(store, { key: 'story', reducer })

      /*  Return getComponent   */
      cb(null, StoryList)

      /* Webpack named bundle   */
    }, 'story')
  }
})

export function CreateStoryFormRoute(store) {
  return {
    path: 'story-form',
    /*  Async getComponent is only invoked when route matches   */
    getComponent(nextState, cb) {
      /*  Webpack - use 'require.ensure' to create a split point
          and embed an async module loader (jsonp) when bundling   */
      require.ensure([], (require) => {
        /*  Webpack - use require callback to define
            dependencies for bundling   */
        const StoryForm = require('./containers/StoryFormContainer').default
        const reducer = require('./modules/story').default

        /*  Add the reducer to the store on key 'counter'  */
        injectReducer(store, { key: 'story', reducer })

        /*  Return getComponent   */
        cb(null, StoryForm)

        /* Webpack named bundle   */
      }, 'storyform')
    }
  };
}