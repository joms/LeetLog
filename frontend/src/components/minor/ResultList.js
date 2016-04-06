import React from 'react';

import List from 'material-ui/lib/lists/list';
import ResultListItem from './resultlistitem';

var ResultList = React.createClass({
    render: function() {
        var i = 1;
        var rows = this.props.leets.map(function(leet, index) {
            if (leet.fields.status[0] == 0) {
                var score = i;
                i++;
            } else {
                var score = "X";
            }
            return <ResultListItem key={leet._id} leet={leet} score={score}/>
        });

        return (
            <List>
                {rows}
            </List>
        );
    }
});

export default ResultList;