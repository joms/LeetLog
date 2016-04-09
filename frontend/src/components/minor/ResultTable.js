import React from 'react';
import Table from 'material-ui/lib/table/table';
import TableHeader from 'material-ui/lib/table/table-header';
import TableBody from 'material-ui/lib/table/table-body';

import ResultTableHeader from './resulttableheader';
import ResultTableRow from './resulttablerow';

var ResultTable = React.createClass({
    render: function() {
        var rows = this.props.leets.map(function(leet) {
            return <ResultTableRow key={leet._id} leet={leet}/>
        });

        return (
            <Table>
                <TableHeader>
                    <ResultTableHeader />
                </TableHeader>
                <TableBody>
                    {rows}
                </TableBody>
            </Table>
        );
    }
});

export default ResultTable;