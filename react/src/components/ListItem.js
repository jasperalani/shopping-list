import React from 'react'
import { CompleteItem, RemoveItem } from './ItemActions'
import DateDisplayer from './DateDisplayer'

export class ListItem extends React.Component {

  componentDidMount () {
    this.setState({
      id: this.props.data.id,
    })
  }

  displayUrl () {
    // console.log(this.props.data.url)
    if (this.props.data.url !== undefined) {
      return <>
        <a className={'url'} href={this.props.data.url}
           key={'li-' + this.props.data.id + '-url'}>
          Link</a>
      </>
    }
  }

  /**
   * @param this.props.data.image_url
   * @param this.props.data.person
   * @param this.props.data.quantity
   * @param this.props.data.created
   */
  render () {

    // console.log(this.props.data.image_url)

    return (
      <div className={'d-flex list-item-parent'}>

        <CompleteItem data={this.props.data} finishItem={this.props.finishListItem} />

        <li className={'list-item'} key={'li-' + this.props.data.id}>

          {/*<p>{this.props.data.id}</p>*/}
          <p className={'name'}
             key={'li-' + this.props.data.id +
             '-name'}>{this.props.data.name}</p>
          <img className={'image'} key={'li-' + this.props.data.id + '-image'}
               src={this.props.data.image_url} alt={this.props.data.name}/>

          <div className={'division-one'}>
            {this.displayUrl()}
            <p className={'quantity'}
               key={'li-' + this.props.data.id + '-quantity'}>
              Qty: {this.props.data.quantity}</p>
          </div>

          <div className={'division-two'}>
            <p className={'person'}
               key={'li-' + this.props.data.id + '-person'}>
              <small>Requested by:</small><br/>
              {this.props.data.person}
            </p>
            <p className={'created'}
               key={'li-' + this.props.data.id + '-created'}>
              <small>Requested on:</small><br/>
              <DateDisplayer date={this.props.data.created} />

            </p>
          </div>

        </li>

        <RemoveItem id={this.props.data.id} finishItem={this.props.finishListItem}/>

      </div>
    )
  }

}