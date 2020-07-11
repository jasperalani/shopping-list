import React from 'react';

const JsonTable = require('ts-react-json-table');

class List extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      data: null
    };
  }

  componentDidMount() {
    fetch('http://localhost:10000/').
        then(response => response.json()).
        then(data => this.setState({data}));

  }

  render() {

    return (
        <JsonTable className={'items-table json-table'} rows={this.state.data}/>
    );

  }

}

export default List;
