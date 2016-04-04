import React from 'react';
import TableRow from 'material-ui/lib/table/table-row';
import TableRowColumn from 'material-ui/lib/table/table-row-column';
import Badge from 'material-ui/lib/badge';

function addZ(n){return n < 10 ? '0'+n : ''+n;}
function addZ2(n){return n < 100 ? '0'+n : ''+n;}

var ResultTableRow = React.createClass({
    render: function() {
        console.log(this.props.leet);

        var date = new Date(this.props.leet.sort[0]);

        var c = (date.getMinutes() == 37) ? "alert-success" : "alert-danger";
        if (this._type == 3) c = "alert-danger";

        var str = addZ(date.getSeconds()) +"."+ addZ2(date.getMilliseconds());

        if (date.getMinutes() != 37)
        {
            str = (date.getMinutes() <= 36 ? "- ":"+ ") + date.getMinutes()+"."+str;

            var x = date.getMinutes() +"."+ date.getSeconds();
            if (x != 36.59 && x != 38.00) str = str.substr(0, str.length - 4);
        }

        var style = {
            height: "40px",
            width: "50px"
        };


        return (
            <TableRow>
                <TableRowColumn>{this.props.leet.fields.nick[0]}</TableRowColumn>
                <TableRowColumn>
                    <Badge
                        badgeContent={str}
                        secondary={true}
                        badgeStyle={style}
                    />
                </TableRowColumn>
            </TableRow>
        );
    }
});

export default ResultTableRow;