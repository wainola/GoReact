import React, { Component } from 'react';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import Navegacion from './components/navegacion';
import '../node_modules/bootstrap/dist/css/bootstrap.min.css';
import MainContent from './containers/main_content';
import SideContent from './containers/side_content';

class App extends Component {
  render() {
    return (
      <div className="App">
        <MuiThemeProvider>
          <div className="row">
            <div className="container-fluid">
              <Navegacion />
            </div>
          </div>
          <div className="container-fluid">
            <div className="row">
              <div className="col-8">
                <MainContent />
              </div>
              <div className="col-4">
                <SideContent />
              </div>
            </div>
          </div>
        </MuiThemeProvider>
      </div>
    );
  }
}

export default App;
