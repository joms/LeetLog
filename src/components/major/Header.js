import React from 'react';
import AppBar from '../../../node_modules/material-ui/AppBar';

const style={
    "marginBottom" : "25px"
}

class Header extends React.Component {

    render() {
        return (
            <AppBar
                showMenuIconButton={false}
                title={this.props.title}
                style={style}
            />
        );
    }
}

export default Header;
