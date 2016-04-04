import React from 'react';
import TableHeaderColumn from 'material-ui/lib/table/table-header-column';
import TableRow from 'material-ui/lib/table/table-row';

const ResultTableHeader = () => (
    <TableRow>
        <TableHeaderColumn>Nick</TableHeaderColumn>
        <TableHeaderColumn>Time</TableHeaderColumn>
    </TableRow>
);

export default ResultTableHeader;