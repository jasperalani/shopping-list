import React from 'react';
import logo from './res/shopping-bag-logo-512.png';
import List from './components/List';

// const AppName = 'Shopping List';
const app_name = 'shopping-list';

function App() {
  return (
      <div className={app_name + "-app"}>
        <header className={app_name + "-header"}>
          <img src={logo} id={app_name + "-logo"} alt="logo"/>
          <h1>Shopping List</h1>
        </header>

        <main id={"main"}>

          <section id={"list"}>

            <List></List>

          </section>

        </main>
      </div>
  );
}

export default App;
