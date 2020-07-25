import React from 'react'
import { ListItem } from './ListItem'
import constants from '../constants'

// import { useAsync } from 'react-async'

class List extends React.Component {
  constructor (props) {
    super(props)

    this.state = {
      data: null,
      lastAddedItem: ''
    }
  }

  finishListItem = (id) => {
    const localisedData = this.state.data

    for (const index in localisedData) {
      if (localisedData.hasOwnProperty(index)) {
        if (id === localisedData[index].id) {
          delete localisedData[index]
        }
      }
    }

    this.setState({data: localisedData})
  }

  removeEmptyFields () {

    for (const index in this.state.data) {
      if (this.state.data.hasOwnProperty(index)) {
        for (const itemField in this.state.data[index]) {
          if (
            this.state.data[index].hasOwnProperty(itemField)
            && this.state.data[index][itemField] === ''
          ) {
            delete this.state.data[index][itemField]
          }
        }
      }
    }

  }

  createList () {

    this.removeEmptyFields()

    const listItems = []
    const addedListItemIndexes = []

    for (const index in this.state.data) {
      if (this.state.data.hasOwnProperty(index)) {

        if (addedListItemIndexes.includes(index)) {
          return
        }

        listItems.push(
          <ListItem
            key={'list-item-key-' + index}
            data={this.state.data[index]}
            finishListItem={this.finishListItem}
          />,
        )
        addedListItemIndexes.push(index)

      }
    }

    return listItems.reverse()

  }

  getJSON () {// eslint-disable-next-line
    fetch(constants.go_endpoint). // eslint-disable-next-line
        then(res => (res.ok ? res : Promise.reject(res))).
        then(res => res.json().then(json => this.setState({ data: json })))
  }

  componentDidMount () {
    this.getJSON()
  }

  componentWillReceiveProps(nextProps) {
    if(nextProps.lastAddedItem !== this.state.lastAddedItem){
      this.getJSON()
    }
  }

  render () {

    return (
      <ul id={'unordered-list'}>
        {this.createList()}
      </ul>
    )

  }

}

export default List;