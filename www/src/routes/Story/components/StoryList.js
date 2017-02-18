import React from 'react'

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

// StoryList.contextTypes = {
//     router: PropTypes.object,
//   };

export default StoryList