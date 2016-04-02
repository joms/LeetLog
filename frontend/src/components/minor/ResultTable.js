import React from 'react';
import Table from 'material-ui/lib/table/table';
import TableHeader from 'material-ui/lib/table/table-header';
import TableBody from 'material-ui/lib/table/table-body';

import ResultTableHeader from './resulttableheader';
import ResultTableRow from './resulttablerow';

const ResultTable = () => (
    <Table>
        <TableHeader>
            <ResultTableHeader />
        </TableHeader>
        <TableBody>
            <ResultTableRow />
        </TableBody>
    </Table>
);

export default ResultTable;