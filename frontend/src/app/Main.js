/**
 * In this file, we create a React component
 * which incorporates components providedby material-ui.
 */

import React from 'react';

import TopBar from './../components/major/TopBar';
import Content from './../components/major/Content';

class Main extends React.Component {
    render() {
        return (
            <div>
                <TopBar />
                <Content />
            </div>
        );
    }
}

export default Main;
