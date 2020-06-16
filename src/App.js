import React, { useState, useEffect } from 'react';
import { Admin, Resource } from 'react-admin';
import './App.css';

const App = () => {
    return (
        <Admin>
            <Resource name="Orders"></Resource>
        </Admin>
    );
}

export default App;