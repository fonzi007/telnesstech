import React, { Component } from 'react';
import Form from "./Form";

class App extends Component {
  state = {
    characters: []
  };

  removeCharacter = index => {
    const { characters } = this.state;

    this.setState({
      characters: characters.filter((character, i) => {
        return i !== index;
      })
    });
  }

  handleSubmit = character => {
    this.setState({characters: [...this.state.characters, character]});
  }

  render() {
    const { characters } = this.state;

    return (
        <div className="container">
          <h1>Phone numbers repertoire</h1>
          <p>Search for a phone number by msisdn. You can add filters (grade, type)</p>
          <h3>Search for phone number:</h3>
          <Form/>
        </div>
    );
  }
}

export default App;