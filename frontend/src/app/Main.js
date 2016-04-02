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
