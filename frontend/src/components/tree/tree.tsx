import * as React from 'react';
import { useState, useEffect } from 'react';
import Tree from 'react-d3-tree'
import './tree.css';


interface IAppProps {
  name: string
}

interface RawNodeDatum {
  name: string;
  attributes?: Record<string, string | number | boolean>;
  children?: RawNodeDatum[];
}

const TreeComponent: React.FC<IAppProps> = props => {

  const [tree, setTree] = useState<RawNodeDatum>({ name: props.name });

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
    <div id="treeWrapper" style={{ width: '100em', height: '50em' }}>
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
export default TreeComponent;
