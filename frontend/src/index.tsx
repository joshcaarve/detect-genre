import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import TreeComponent from './components/tree/tree';
import reportWebVitals from './reportWebVitals';
import FormComponent from './components/form/form';

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
