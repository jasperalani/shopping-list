import React from 'react'

class ItemAction extends React.Component {
  constructor (props) {
    super(props)
    this.state = {
      id: 0
    }
  }

  componentDidMount () {
    this.setState({ id: this.props.id })
  }
}

export class CompleteItem extends ItemAction {

  toggleButtonState = () => {

    const localData = this.props.data;
    localData.completed = true;
    const json = JSON.stringify(localData);
    const headers = {
      method: "PUT",
      body: json
    }

    fetch("http://localhost:10000/" + localData.id, headers)
    .then(res => (res.ok ? res : Promise.reject(res)))
    .then(res => res.json().then(json => {
      console.log(json.response)
      if(json.response === "item_updated"){
        this.props.finishItem(this.state.id)
      }
    }))

  };

  render () {
    return (
      <i className="fas fa-check" onClick={this.toggleButtonState}/>
    )
  }
}

export class RemoveItem extends ItemAction {

  toggleButtonState = () => {

    fetch("http://localhost:10000/" + this.state.id, {method: "DELETE"})
    .then(res => (res.ok ? res : Promise.reject(res)))
    .then(res => res.json().then(json => {
      if(json.response === "item_deleted"){
        this.props.finishItem(this.state.id)
      }
    }))

  };

  render () {
    return (
      <i className="fas fa-times" onClick={this.toggleButtonState}/>
    )
  }

}