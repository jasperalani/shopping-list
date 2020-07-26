import React, {useState} from 'react';
import {RemoveItem} from './ItemActions';
import Modal from 'react-bootstrap/Modal';
import DateDisplayer from './DateDisplayer';
import ImageUploader from 'react-images-upload';



// import DateDisplayer from './DateDisplayer'

export class ListItem extends React.Component {

  constructor(props) {
    super(props);

    this.state = {
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

  onChange = () => {
    const file = this.state.pictures[0];
    const formData = new FormData()
    formData.append('file', file);

    const headers = {
      method: 'POST',
      body: formData,
      mode: 'cors'
    };

    fetch('http://localhost:8888/image', headers)
        .then(res => console.log(res.headers));
    // .then(res => res.json())
    // .then(image => {
    //   // console.log(image)
    //   // this.setState({
    //   //   uploading: false,
    //   //   imageId
    //   // })
    // })
  }

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

            <Modal show={this.state.show} onHide={this.handleClose}>
              <div className={'list-item-modal'}>
                <Modal.Header closeButton>
                  <Modal.Title>{this.props.data.name}</Modal.Title>
                </Modal.Header>
                <Modal.Body>

                  <form>

                    <ImageUploader
                        withIcon={true}
                        buttonText='Choose image'
                        onChange={this.onDrop}
                        imgExtension={['.jpg', '.jpeg', '.gif', '.png']}
                        maxFileSize={5242880}
                    />

                    {/*<img className={'oranges-image'}*/}
                    {/*     src={'https://lh3.googleusercontent.com/proxy/9aY-f3Q_yhoI2lPBmTU9UZSmlGye0sfxmX9QQcI-n-3pmpbgcGsjxPbeAonfJMKUocUCOt8U1XyvLmf0rF520TsbNWdf2PQpm8yCOoIWj7VhmJPReS4p'}*/}
                    {/*     alt={this.props.data.name}/>*/}

                    <div className={'d-flex item-attribute requested-by'}>
                      <p>Requested By:</p>
                      <input></input>
                    </div>


                    <div className={'d-flex item-attribute quantity'}>
                      <p>Quantity:</p>
                      <input></input>
                    </div>

                    <p>Date Requested: <DateDisplayer
                        date={this.props.data.created}/></p>

                    <div className={'d-flex flex-column item-attribute notes'}>
                      <p>Notes</p>
                      <textarea></textarea>
                    </div>

                  </form>

                </Modal.Body>
                <Modal.Footer>
                  <button className={'btn-light'} onClick={this.handleClose}>
                    Close
                  </button>
                  <button className={'btn-light'} onClick={this.onChange}>
                    Save Changes
                  </button>
                </Modal.Footer>
              </div>
            </Modal>

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