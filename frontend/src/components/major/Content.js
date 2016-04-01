import React from 'react';

import Time from './../minor/Time';
import ResultTable from './../minor/ResultTable';

class Content extends React.Component {

    render() {
        return (
            <div>
                <Time />
                <ResultTable />
            </div>
        );
    }
}

export default Content;
