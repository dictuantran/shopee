import React, { Component } from 'react';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Title from '../title/Title';

class Order extends Component {
    constructor(props) {
        super(props)

        this.state = {        
            items: []
        }      
    }       

    componentDidMount() {
        const apiURL = "http://localhost:3001/order/GetOrder";       
        fetch(apiURL)
            .then(response => {
                if(!response.ok) {                                    
                }                                
                return response.json()
            })
            .then(json => {                                
                this.setState({items: json.data});                        
            })
            .catch(error=> {
                console.log("error" + error);
            })            
    }

    render() {        
        const rows = this.state.items;

        return (
            <React.Fragment>
                <Title>Recent Orders</Title>
                <Table size="small">
                    <TableHead>
                        <TableRow>
                            <TableCell>
                                Order ID
                            </TableCell>
                            <TableCell>
                                Order Date
                            </TableCell>
                            <TableCell>
                                Order Status
                            </TableCell>
                            <TableCell>
                                Created By
                            </TableCell>
                            <TableCell>
                                Payment Method
                            </TableCell>
                            <TableCell>
                                Payment Status
                            </TableCell>
                            <TableCell>
                                Price
                            </TableCell>
                        </TableRow>                        
                    </TableHead>
                       
                    <TableBody>
                    {rows.map((row) => (
                        <TableRow key={row.OrderId}>
                            <TableCell>
                                {row.OrderId}
                            </TableCell>
                            <TableCell>
                                {row.OrderDate}
                            </TableCell>
                            <TableCell>
                                {row.OrderStatus}
                            </TableCell>
                            <TableCell>
                                {row.CreatedBy}
                            </TableCell>
                            <TableCell>
                                {row.PaymentMethod}
                            </TableCell>
                            <TableCell>
                                {row.PaymentStatus}
                            </TableCell>
                            <TableCell>
                                {row.Price}
                            </TableCell>
                        </TableRow>   
                    ))}
                       
                    </TableBody>                  
                </Table>
            </React.Fragment>
        )
    }    
}

export default Order;