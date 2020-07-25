import React from 'react';
import ItemCreator from './ItemCreator';

class TopJacker extends React.Component {

  returnToTop = () => {
    window.scrollTo(0, 0);
  }

  render() {
    return (
        <>
          <small className={"return-to-top"} onClick={this.returnToTop}>Top</small>
        </>
    );
  }

}

export default TopJacker;