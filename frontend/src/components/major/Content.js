import React from 'react';

import Time from './../minor/Time';
import ResultTable from './../minor/ResultTable';

class Content extends React.Component {

    render() {
        return (
            <div>
                <div>
                    <Time />
                </div>
                <div>
                    <ResultTable />
                </div>
            </div>
        );
    }
}

export default Content;
