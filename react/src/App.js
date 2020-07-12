import React from 'react';
import logo from './res/shopping-bag-logo-512.png';
import {List} from './components/List';

// const AppName = 'Shopping List';
const app_name = 'shopping-list';

export class App extends React.Component {

  render () {
    return (
      <div id={app_name + "-app"}>

        <header id={"header"}>
          <img src={logo} id={"logo"} alt="logo"/>
          <h1>Shopping List</h1>
        </header>

        <main id={"main"}>
          <section id={"list"}>
            <List/>
          </section>
        </main>

      </div>
    );
  }

}

export default App;
