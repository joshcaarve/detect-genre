import * as React from 'react';
import { useState, useEffect } from 'react';
//import './App.css';

interface IAppProps {
}

const App: React.FC<IAppProps> = props => {

  const [ name, setName ] = useState<string>('');

  const getName = async () => {
    await fetch("/api/hello")
    .then(response => response.json())
    .then(responseJson => {
      setName(responseJson.data)
    })
  }

  useEffect(() => {
    getName();
  }, []);

  return (
    <main className="container">
      <h1> {name} </h1>
    </main>
  )
}
export default App;
