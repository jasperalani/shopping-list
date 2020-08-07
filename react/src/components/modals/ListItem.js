import Modal from 'react-bootstrap/Modal';
import ImageUploader from 'react-images-upload';
import DateDisplayer from '../DateDisplayer';
import React from 'react';

class ModalListItem extends React.Component {

  constructor(props) {
    super(props);

    this.state = {
      data: [],
      name: '',
      requested_by: '',
      quantity: 0,
      notes: '',
      created: ''
    }
  }

  componentDidMount() {
    console.log(this.props.data)
    this.setState({
      name: this.props.data.name ?? '',
      requested_by: this.props.data.person ?? '',
      quantity: this.props.data.quantity ?? '',
      notes: this.props.data.notes ?? '',
      created: this.props.data.created ?? ''
    })
  }

  handleChange = (event) => {
    console.log([event.target.name])
    this.setState({
      [event.target.name]: event.target.value
    });
  }

  render() {
    return (
        <Modal show={this.props.show} onHide={this.props.close}>
          <div className={'list-item-modal'}>
            <Modal.Header closeButton>
              <Modal.Title>
                <input value={this.state.name} name={"name"} onChange={this.handleChange}/>
              </Modal.Title>
            </Modal.Header>
            <Modal.Body>

              <form>

                <ImageUploader
                    withIcon={true}
                    buttonText='Choose image'
                    onChange={this.props.drop}
                    imgExtension={['.jpg', '.jpeg', '.gif', '.png']}
                    maxFileSize={5242880}
                />

                <div className={'d-flex item-attribute requested-by'}>
                  <p>Requested By:</p>
                  <input value={this.state.requested_by} name={"requested_by"} onChange={this.handleChange}/>
                </div>

                <div className={'d-flex item-attribute quantity'}>
                  <p>Quantity:</p>
                  <input value={this.state.quantity} name={"quantity"} onChange={this.handleChange}/>
                </div>

                <p>Date Requested: <DateDisplayer
                    date={this.state.created}/></p>

                <div className={'d-flex flex-column item-attribute notes'}>
                  <p>Notes</p>
                  <textarea value={this.state.notes} name={"notes"} onChange={this.handleChange}/>
                </div>

              </form>

            </Modal.Body>
            <Modal.Footer>
              <button className={'btn-light'} onClick={this.props.close}>
                Close
              </button>
              <button className={'btn-light'} onClick={this.props.click}>
                Save Changes
              </button>
            </Modal.Footer>
          </div>
        </Modal>
    )
  }

}

export default ModalListItem;