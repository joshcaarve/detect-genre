import React, { Component } from 'react'; // let's also import Component
export class App extends Component {

  // The tick function sets the current state. TypeScript will let us know
  // which ones we are allowed to set.

  // After the component did mount, we set the state each second.
  componentDidMount() {
    
  }

  // render will know everything!
  render() {
    return <p> "Hello" </p>
  }
}

export default App
