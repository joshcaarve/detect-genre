import * as React from 'react';
import { useState, useEffect } from 'react';
import Tree from 'react-d3-tree'
import './custom-tree.css';
import './App.css';

interface IAppProps {
}

interface RawNodeDatum {
  name: string;
  attributes?: Record<string, string | number | boolean>;
  children?: RawNodeDatum[];
}

const App: React.FC<IAppProps> = props => {

  const [tree, setTree] = useState<RawNodeDatum>({name:'init'});

  const getTree = async () => {
    await fetch("/api/tree")
      .then(response => response.json())
      .then(responseJson => {
        setTree(responseJson.data)
      })
  }

  useEffect(() => {
    getTree();
  }, []);

  return (
    <div id="treeWrapper" style={{ width: '1000em', height: '200em' }}>
      <h1>Genre Tree</h1>
      <Tree
        data={tree}
        orientation={"vertical"}
        rootNodeClassName="node__root"
        branchNodeClassName="node__branch"
        leafNodeClassName="node__leaf"
      />
    </div>
  )
}
export default App;
