import React from 'react';
import Card from 'material-ui/lib/card/card';
import CardHeader from 'material-ui/lib/card/card-header';
import CardText from 'material-ui/lib/card/card-text';

var moment = require('moment');

const style = {
    "maxWidth" : "200px"
};

var Time = React.createClass({
    getInitialState: function() {
        return {countdown: this.getTime()};
    },

    componentDidMount: function() {
        setInterval(this.updateTime, 500);
    },

    updateTime: function() {
        this.setState({countdown: this.getTime()});
    },

    getTime: function() {
        var t = moment();
        return t.format("HH:mm:ss");
    },

    render: function() {
        return (
            <div>
                <Card style={style} >
                    <CardHeader
                        title="Next Leet"
                    />
                    <CardText>
                        {this.state.countdown}
                    </CardText>
                </Card>
            </div>
        );
    }
});

export default Time;
