import React from 'react';

import Header from './../components/major/Header';
import Content from './../components/major/Content';

class Main extends React.Component {
    render() {
        return (
            <div>
                <Header
                    title={"LeetLog \u2014 #scene.no"}
                />
                <Content />
            </div>
        );
    }
}

export default Main;
