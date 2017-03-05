import React from 'react'
import { connect } from 'react-redux'
import { create, save } from '../modules/story'

export const StoryList = (props) => (
  <div style={{ margin: '0 auto' }} >
    <button className='btn btn-default' onClick={props.create}>
      Create
    </button>
    <ul>
      {typeof props.stories.map === 'function' && props.stories.map(story =>
        <li>
          {story.name}
        </li>
      )}
    </ul>
  </div>
)

StoryList.propTypes = {
  stories: React.PropTypes.array.isRequired,
  create: React.PropTypes.func.isRequired
}

const mapDispatchToProps = {
  create: () => create()
}

const mapStateToProps = (state) => (
  {
    stories: state.story.entities.stories
  }
)


//export default StoryList

export default connect(mapStateToProps, mapDispatchToProps)(StoryList)