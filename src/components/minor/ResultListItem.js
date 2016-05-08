import React from 'react';

import Badge from 'material-ui/Badge';
import ListItem from 'material-ui/List/ListItem';
import FontIcon from 'material-ui/FontIcon';

function addZ(n){return n < 10 ? '0'+n : ''+n;}
function addZ2(n){return n < 100 ? '0'+n : ''+n;}

var ResultListItem = React.createClass({
    render: function() {

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

        var numberStyle = {
            top: "18px"
        };

        var itemStyle = {
            paddingLeft: "150px"
        };

        var style = {
            borderTop: "1px solid rgba(0, 0, 0, 0.1)"
        };

        return (
            <ListItem
                leftIcon={
                    <FontIcon
                        style={numberStyle}
                    >
                        {this.props.score}
                    </FontIcon>
                }
                secondaryText={str}
                primaryText={this.props.leet.fields.nick[0]}
                innerDivStyle={itemStyle}
                style={this.props.id != 0 ? style : undefined}
            />
        );
    }
});

export default ResultListItem;