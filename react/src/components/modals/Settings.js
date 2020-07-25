import React, {useState} from 'react';
import Modal from 'react-bootstrap/Modal';

function Settings() {
  const [show, setShow] = useState(false);

  const handleClose = () => setShow(false);
  const handleShow = () => setShow(true);

  return (
      <>
        <div id={'settings-modal'}>
          <button onClick={handleShow}>
            <i className="fas fa-cog"></i>
          </button>

          <Modal show={show} onHide={handleClose}>
            <Modal.Header closeButton>
              <Modal.Title>Settings</Modal.Title>
            </Modal.Header>
            <Modal.Body>
              <p>Here are some settings, warning! no undo's</p>
              <button>Clear all items </button>
            </Modal.Body>
            <Modal.Footer>
              <button onClick={handleClose}>
                Close
              </button>
            </Modal.Footer>
          </Modal>
        </div>
      </>
  );
}

export default Settings;