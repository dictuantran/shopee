import React, { Component } from 'react';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import EditIcon from '@material-ui/icons/Edit';
import DeleteIcon from '@material-ui/icons/Delete';
import KeyboardArrowDownIcon from '@material-ui/icons/KeyboardArrowDown';
import KeyboardArrowUpIcon from '@material-ui/icons/KeyboardArrowUp';
import Title from '../title/Title';
import { IconButton } from '@material-ui/core';

const [open, setOpen] = React.useState(false);

const collapseExpandIcon = (
    <IconButton aria-label="expand row" size="small" onClick={() => setOpen(!open)}>
        {open ? <KeyboardArrowUpIcon /> : <KeyboardArrowDownIcon />}
    </IconButton>
);

const editIcon = (
    <IconButton onClick={console.log("click edit")}>
        <EditIcon color="primary" />
    </IconButton>
);

const deleteIcon = (
    <IconButton onClick={console.log("click delete")}>
        <DeleteIcon color="primary" />
    </IconButton>
);

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
                            </TableCell>
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
                            <TableCell>
                                Edit
                            </TableCell>
                        </TableRow>                        
                    </TableHead>
                       
                    <TableBody>
                    {rows.map((row) => (
                        <TableRow key={row.OrderId}>
                            <TableCell>
                                {collapseExpandIcon}
                            </TableCell>
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
                            <TableCell>                             
                                {editIcon}
                                {deleteIcon}
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