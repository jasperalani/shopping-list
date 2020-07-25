import React from 'react';
import constants from '../constants';

class ItemCreator extends React.Component {

  constructor(props) {
    super(props);

    // this.state = {
    //   renderForm: false,
    //   formHTML: '',
    //   actionButtonText: 'Add new',
    // };

    this.state = {value: ''};
  }

  renderForm = () => {

    let form = '', actionButtonText = 'Add new';
    if (!this.state.renderForm) {
      form = <>
        <form id={'item-creator-form'}>
          <label>
            Name *:
            <input type="text" name="name"/>
          </label>

          <label>
            URL:
            <input type="text" name="url"/>
          </label>

          <label>
            Image URL:
            <input type="text" name="imageurl"/>
          </label>

          <label>
            Person *:
            <input type="text" name="person"/>
          </label>

          <label>
            Quantity:
            <input type="text" name="quantity"/>
          </label>

          <input type="submit" value="Submit"/>
        </form>
      </>;
      actionButtonText = 'Cancel';
    }

    this.setState({
      renderForm: !this.state.renderForm,
      formHTML: form,
      actionButtonText: actionButtonText,
    });
  };

  handleChange = (event) => {
    this.setState({value: event.target.value});
  }

  addItem = () => {

    const newItem = {
      name: this.state.value
    }

    if(newItem.name === '' || newItem.name === undefined){
      return false;
    }

    const headers = {
      method: "POST",
      body: JSON.stringify(newItem)
    }

    fetch(constants.go_endpoint, headers)
    .then(res => (res.ok ? res : Promise.reject(res)))
    .then(res => res.json().then(json => {
      console.log(json)
      // console.log(json.response === "item_created")
      if(json.response === "item_created" || json.response === "quantity_increased"){
        // console.log(newItem)
        this.props.refreshHandler(newItem.name)
        // this.forceUpdate();
      }
    }))

  }

  _handleKeyDown = (e) => {
    if (e.key === 'Enter') {
      this.addItem()
    }
  }

  render() {
    return (
        <div id={'item-creator'} className={'container'}>
          <div className={'row'}>
            <div className={'col'}>
              {/*{this.state.formHTML}*/}
              <input placeholder={'What would you like to add?'} value={this.state.value} onChange={this.handleChange} onKeyDown={this._handleKeyDown}/>
              <button className={'form-action-button'} onClick={this.addItem}>
                Add
              </button>
            </div>
          </div>
        </div>
    );
  }
}

export default ItemCreator;