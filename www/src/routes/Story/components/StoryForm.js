import React from 'react'
import { connect } from 'react-redux'
import { create, save } from '../modules/story'
import { Control, Form, actions } from 'react-redux-form';

export const StoryForm = (props) => (
  <div style={{ margin: '0 auto' }} >
    <Form model="currentstory"
        onSubmit={(storyForm) => this.handleSubmit(currentstory)}>
         <label>Story:</label>
        <Control.text model="currentstory.name" />

        <label>Last name:</label>
        <Control.text model="currentstory.content" />
      <button className='btn btn-default' type="submit">
        save
      </button>
    </Form>
  </div>
)

StoryForm.propTypes = {
  currentstory: React.PropTypes.object,
  save: React.PropTypes.func.isRequired,
  onFieldChanged: React.PropTypes.func.isRequired
}

const mapDispatchToProps = {
  save: () => save(),
  onFieldChanged: function (event) {
    this.setState({
      [event.target.name]: event.target.value
    });
  }
}

const mapStateToProps = (state) => ({
  currentstory: state.story.storyForm.currentstory
})


export default connect(mapStateToProps, mapDispatchToProps)(StoryForm)
