import React, { Component } from 'react';
import './App.css';
import './index.css';
import Layout from './components/layout/Layout';
import {BrowserRouter, Switch, Route} from 'react-router-dom';

class App extends Component {
    render() {
        return (
            <BrowserRouter>
            <Switch>
                <Route exact path="/" component={Layout} />
                <Route exact path="/order" component={Layout} />
                <Route exact path="/category" component={Layout} />
            </Switch>
            </BrowserRouter>        
        )
    }
}

export default App