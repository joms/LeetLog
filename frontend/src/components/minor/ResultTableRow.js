import React from 'react';
import TableRow from 'material-ui/lib/table/table-row';
import TableRowColumn from 'material-ui/lib/table/table-row-column';

var ResultTableRow = React.createClass({
    render: function() {
        return (
            <TableRow>
                <TableRowColumn>{this.props.leet._id}</TableRowColumn>
                <TableRowColumn>{this.props.leet.fields.nick[0]}</TableRowColumn>
                <TableRowColumn>Employed</TableRowColumn>
            </TableRow>
        );
    }
});

export default ResultTableRow;