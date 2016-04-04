import React from 'react';

import List from 'material-ui/lib/lists/list';
import ResultListItem from './resultlistitem';

var ResultList = React.createClass({
    render: function() {
        var rows = this.props.leets.map(function(leet, index) {
            return <ResultListItem key={leet._id} leet={leet} id={index}/>
        });

        return (
            <List>
                {rows}
            </List>
        );
    }
});

export default ResultList;