// ------------------------------------
// Constants
// ------------------------------------
export const STORY_CREATE = 'STORY_CREATE'
export const STORY_SAVE = 'STORY_SAVE'

// ------------------------------------
// Actions
// ------------------------------------
export function create() {
  return {
    type: STORY_CREATE,
    story: { name: "", content: "" }
  }
}

/*  This is a thunk, meaning it is a function that immediately
    returns a function for lazy evaluation. It is incredibly useful for
    creating async actions, especially when combined with redux-thunk!

    NOTE: This is solely for demonstration purposes. In a real application,
    you'd probably want to dispatch an action of COUNTER_DOUBLE and let the
    reducer take care of this logic.  */

export const save = (story) => {
  return (dispatch, getState) => {
    return new Promise((resolve) => {
      setTimeout(() => {
        dispatch({
          type: STORY_CREATE,
          story: story
        })
        resolve()
      }, 200)
    })
  }
}


export const actions = {
  create,
  save
}

// ------------------------------------
// Action Handlers
// ------------------------------------
const ACTION_HANDLERS = {
  [STORY_CREATE]: (state, action) => { return{ ...state, currentstory : action.story }},
  [STORY_SAVE]: (state, action) => { return{ ...state,  currentstory : action.story }}
}

// ------------------------------------
// Reducer
// ------------------------------------
const initialState = {
  currentstory: null,
  stories: []
}
export default function storyReducer(state = initialState, action) {
  const handler = ACTION_HANDLERS[action.type]

  return handler ? handler(state, action) : state
}
