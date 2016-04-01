import React from 'react';
import Card from 'material-ui/lib/card/card';
import CardHeader from 'material-ui/lib/card/card-header';
import CardText from 'material-ui/lib/card/card-text';

var moment = require('moment');

const style = {
    // "maxWidth" : "250px"
    marginBottom: "20px"
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
        var now = moment();
        var next = new moment("13 37", "HH mm");

        if (next -  now  < 0) {
            next = next.add(1, 'days');
        }

        return moment(next).fromNow();
    },

    render: function() {
        return (
            <div>
                <Card style={style} >
                    <CardHeader
                        title="Next Leet is in..."
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
