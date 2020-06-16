import React, { Fragment } from 'react';

import ItemService from '../../services/itemService';

class TabbedDatagrid extends React.Component {   
    constructor(props) {
        super(props)
        this.itemService = new ItemService();
        
        this.state = {
            taxCodes: null
        }
    }

    componentDidMount(){
        this.getTaxCodes();
    }

    render() {
        const taxCodes = this.state.taxCodes;        
        
        if(!taxCodes) return null;

        const taxCodeList = taxCodes.map((taxcode) => 
            <li key={taxcode.tax_code_id}> 
                <span className="name">{taxcode.tax_code_id}</span>&nbsp;|&nbsp; {taxcode.name}
            </li>
        )
    
        return (
            <div>
                <ul className="taxcodes">
                {taxCodeList}
                </ul>
            </div>
        );
    }

    getTaxCodes(){          
        this.itemService.retrieveTaxcode().then(taxcodes => {
            this.setState({taxCodes : taxcodes});
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
