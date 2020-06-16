import React, { Fragment } from 'react';

import ItemService from '../../services/itemService';

class TabbedDatagrid extends React.Component {   
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
        this.getTaxCodes();
    }

    render() {
        const items = this.state.items;        

      /*   const listItems = items.map((item) => 
            <li key={item.link}> 
                <span className="item-name">{item.name}</span>&nbsp;|&nbsp; {item.summary}
            </li>
        )
    */
        return (
            <div className="App">
                <ul className="items">
                   {/*  {listItems} */}
                </ul>
                <br />
                <button type="button" name="button"></button>
                <br />
            </div>
        );
    }

    getTaxCodes(){          
        this.itemService.retrieveTaxcode().then(taxcodes => {
            console.log(taxcodes);
        })
    }
}

const StyledTabbedDatagrid = props => {
    return <TabbedDatagrid {...props} />;
};

const OrderList = ({ classes, ...props }) => (
    <StyledTabbedDatagrid />
);

export default OrderList;
