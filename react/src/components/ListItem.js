import React, {useState} from 'react';
import axios from 'axios';

import {RemoveItem} from './ItemActions';
import ModalListItem from './modals/ListItem';
import constants from '../constants';

// import DateDisplayer from './DateDisplayer'

export class ListItem extends React.Component {

  constructor(props) {
    super(props);

    this.state = {
      id: this.props.data.id,
      data: this.props.data,
      show: false,
      pictures: [],
      uploading: false,
      imageId: 0,
    };
    this.onDrop = this.onDrop.bind(this);
  }

  onDrop(picture) {
    this.setState({
      pictures: this.state.pictures.concat(picture),
    });
  }

  componentDidMount() {
    this.setState({
      id: this.props.data.id,
    });
  }

  displayUrl() {
    // console.log(this.props.data.url)
    if (this.props.data.url !== undefined) {
      return <>
        <a className={'url'} href={this.props.data.url}
           key={'li-' + this.props.data.id + '-url'}>
          Link</a>
      </>;
    }
  }

  onSaveChanges = () => {

    // Save picture to database
    const pictures = this.state.pictures;
    if (pictures.length > 0 && pictures.length < 2) {
      const data = new FormData();
      data.append('file', pictures[0]);
      axios.post(constants.php_post, data).then(async res => {
        console.log("saving image:" + res)
        await this.saveImageIdToDatabase(res);
      });
    }

  };

  saveImageIdToDatabase = async (imageid) => {

    const data = new FormData();
    data.append('id', this.state.id);
    data.append('image_id', imageid.data.image_id);
    axios.put(constants.go_endpoint, data).then(async res => {
      console.log("saving id:" + res)
    });

  };

  handleClose = () => {
    this.setState({show: false});
  };

  handleShow = () => {
    this.setState({show: true});
  };

  render() {

    return (
        <div className={'list-item-wrapper'}>

          <li className={'list-item reduced'}
              key={'li-' + this.props.data.id}
              onClick={this.handleShow}>

            <div className={'d-flex list-item-inner'}>
              <p className={'name'}
                 key={'li-' + this.props.data.id +
                 '-name'}>{this.props.data.name}</p>
              <RemoveItem id={this.props.data.id}
                          finishItem={this.props.finishListItem}/>
            </div>

          </li>


          <div id={'list-item-modal'}>

            <ModalListItem show={this.state.show}
              data={this.props.data} close={this.handleClose}
              drop={this.onDrop} click={this.onSaveChanges}
            />

          </div>

        </div>

    );
  }

}

// {/*<p>{this.props.data.id}</p>*/}
//
// {/*<img className={'image'} key={'li-' + this.props.data.id + '-image'}*/}
// {/*     src={this.props.data.image_url} alt={this.props.data.name}/>*/}
//
// {/*<div className={'division-one'}>*/}
// {/*  {this.displayUrl()}*/}
// {/*  <p className={'quantity'}*/}
// {/*     key={'li-' + this.props.data.id + '-quantity'}>*/}
// {/*    Qty: {this.props.data.quantity}</p>*/}
// {/*</div>*/}
//
// {/*<div className={'division-two'}>*/}
// {/*  <p className={'person'}*/}
// {/*     key={'li-' + this.props.data.id + '-person'}>*/}
// {/*    <small>Requested by:</small><br/>*/}
// {/*    {this.props.data.person}*/}
// {/*  </p>*/}
// {/*  <p className={'created'}*/}
// {/*     key={'li-' + this.props.data.id + '-created'}>*/}
// {/*    <small>Requested on:</small><br/>*/}
// {/*    <DateDisplayer date={this.props.data.created} />*/}
//
// {/*  </p>*/}
// {/*</div>*/}