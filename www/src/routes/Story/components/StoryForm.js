import React from 'react'

export const StoryForm = (props) => (
  <div style={{ margin: '0 auto' }} >
    <h2>Story: {props.currentstory.name}</h2>
    <button className='btn btn-default' onClick={props.save}>
      save
    </button>
  </div>
)

StoryForm.propTypes = {
  currentstory: React.PropTypes.object,
  save: React.PropTypes.func.isRequired
}

export default StoryForm
