import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import TreeComponent from './components/tree/tree';
import ButtonComponent from './components/button/button';
import reportWebVitals from './reportWebVitals';
import FormComponent from './components/form/form';

/*
    <ButtonComponent 
        border="none"
        color="pink"
        height = "200px"
        onClick={() => console.log("You clicked on the pink circle!")}
        radius = "50%"
        width = "200px"
        children = "I'm a pink circle!"
      />
          
*/
ReactDOM.render(
  <React.StrictMode>
    <div>
      <TreeComponent
        name=""
      />
      <FormComponent
        name="foo"
        value="bar"
      />
    </div>
  </React.StrictMode>,
  document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
