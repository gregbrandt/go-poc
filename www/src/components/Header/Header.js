import React from 'react'
import { IndexLink, Link } from 'react-router'
import './Header.scss'

export const Header = () => (
  <div>
    <h1>React Redux Starter Kit</h1>
    <IndexLink to='/' activeClassName='route--active'>
      Home
    </IndexLink>
    {' · '}
    <Link to='/story' activeClassName='route--active'>
      Stories
    </Link>
    {' · '}
    <Link to='/story/story-form' activeClassName='route--active'>
      Create
    </Link>
    {' · '}
    <Link to='/counter' activeClassName='route--active'>
      Counter
    </Link>
  </div>
)

export default Header