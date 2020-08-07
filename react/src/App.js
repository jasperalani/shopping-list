import React from 'react';
// import {useFetch} from 'react-async';

import logo from './res/shopping-bag-logo-512.png';

import List from './components/List';
import ItemCreator from './components/ItemCreator';
import TopJacker from './components/TopJacker';
import Settings from './components/modals/Settings';


import constants from './constants';

export class App extends React.Component {

  constructor(props) {
    super(props);

    this.state = {
      data: [],
      createdItem: false,
    };

  }

  getData = () => {
    const that = this;
    fetch(constants.go_endpoint, {method: 'GET'}).
        then(res => (res.ok ? res : Promise.reject(res))).
        then(res => res.json().then(json => {
          // console.log(json);
          that.setState({data: json});
        }));
  };

  componentDidMount() {
    this.getData();
  }

  remountComponent = (newItem) => {
    this.setState({
      lastAddedItem: newItem,
    });
  };

  render() {

    return (
        <div id={constants.app_name + '-app'}>


          <div className={'background'}/>

          <header id={'header'}>
            <img src={logo} id={'logo'} alt="logo"/>
            <h1>{constants.app_display_name}</h1>
          </header>

          <main id={'main'}>

            <div className={'modals'}>
              <Settings/>
            </div>

            <ItemCreator refreshHandler={this.remountComponent}/>

            <section id={'list'}>

              <div className={'container'}>
                <div className={'row'}>
                  <div className={'col'}>

                    <List lastAddedItem={this.state.lastAddedItem}/>

                  </div>
                </div>
              </div>

            </section>
          </main>

          <footer>
            <small>
              {constants.footer_copyright}
            </small>
            <TopJacker/>
          </footer>



        </div>
    );
  }

}

export default App;