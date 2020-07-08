import React, { Component } from 'react';
import axios from 'axios';

class OrderDetail extends Component {
    state = {
        selectedFile: null
    };

    onFileChange = event => {
        this.setState({ selectedFile: event.target.files[0] });
    }

    onFileUpload = () => {
        const formData = new FormData();

        formData.append(
            "myFile",
            this.state.selectedFile,
            this.state.selectedFile.name
        )          
     
        var apiURL = "http://localhost:3001/order/ImportOrder"; 
        axios.post(apiURL, formData);
    }

    render(){
        return (
            <div>
                <h3>File upload using ReactJS</h3>
                <div>
                    <input type="file" onChange={this.onFileChange} />
                    <button onClick={this.onFileUpload}>Upload</button>
                </div>
            </div>
        )
    }
}

export default OrderDetail;
