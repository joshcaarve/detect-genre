import React from 'react';
import './App.css';
import Tree from 'react-d3-tree';

export default class App extends React.Component {

  render() {
    let data = {
      "name": "Electronic",
      "children": [
        {
          "name": "DnB",
          "children": [
            {
              "name": "Deep",
              "children": []
            },
            {
              "name": "Jungle",
              "children": []
            }
          ]
        },
        {
          "name": "House",
          "children": [
            {
              "name": "Deep",
              "children": []
            },
            {
              "name": "BigRoom",
              "children": []
            }
          ]
        },
        {
          "name": "Dubstep",
          "children": []
        }
      ]
    }
    return (
    <h1> "Genre Tree"
      <div>
        <Tree
          data={data}
          orientation={'vertical'}
        />
      </div>
    </h1>
    )
  }
}
