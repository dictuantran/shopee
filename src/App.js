import React, { useState, useEffect } from 'react';
import { Admin, Resource } from 'react-admin';
import './App.css';

import fakeServerFactory from './components/fakeServer';
import dataProviderFactory from './components/dataProvider';
import { Layout } from './components/layout';
import orders from './components/orders';

const App = () => {
    const [dataProvider, setDataProvider] = useState(null);

    useEffect(() => {
        let restoreFetch;

        const fetchDataProvider = async () => {       
            setDataProvider(
                await dataProviderFactory(process.env.REACT_APP_DATA_PROVIDER)
            );            
        }

        fetchDataProvider();

        return restoreFetch;
    }, []);

    if (!dataProvider) {
        return (
            <div className="loader-container">
                <div className="loader">Loading...</div>
            </div>
        );
    }

    return (
        <Admin 
            dataProvider={dataProvider}
            layout={Layout}
        >
            <Resource name="Orders" {...orders}></Resource>
        </Admin>
    );
}

export default App;