import React, { Component } from 'react';
import './App.css';
//import ItemService from './shared/mock-item-service';
import ItemService from './services/itemService';
import { render } from 'react-dom';

class App extends Component {
    constructor(props) {
        super(props)
        this.itemService = new ItemService();
        
        this.state = {
            showDetails: false,
            editItem: false,
            selectedItem: null,
            newItem: null
        }
    }

    componentDidMount(){
        this.getItems();
        this.getTaxCodes();
    }

    render() {
        const items = this.state.items;
        if(!items) return null;
        const showDetails = this.state.showDetails;
        const selectedItem = this.state.selectedItem;
        const newItem = this.state.newItem;
        const editItem = this.state.editItem;
        const listItems = items.map((item) => 
            <li key={item.link}> 
                <span className="item-name">{item.name}</span>&nbsp;|&nbsp; {item.summary}
            </li>
        )

        return (
            <div className="App">
                <ul className="items">
                    {listItems}
                </ul>
                <br />
                <button type="button" name="button"></button>
                <br />
            </div>
        );
    }

    getItems(){
        this.itemService.retrieveItems().then(items => {
            this.setState({items: items});
        })
    }

    getTaxCodes(){
        console.log("getTaxCodes");
        this.itemService.retrieveTaxcode().then(taxcodes => {
            console.log(taxcodes);
        })
    }
}

export default App;