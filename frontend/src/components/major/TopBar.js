/**
 * In this file, we create a React component
 * which incorporates components providedby material-ui.
 */

import React from 'react';
import AppBar from '../../../node_modules/material-ui/lib/app-bar';

const style={
    "marginBottom" : "25px"
}

class TopBar extends React.Component {

    render() {
        return (
            <AppBar
                title="LeetLog"
                style={style}
            />
        );
    }
}

export default TopBar;
