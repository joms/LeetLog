import React from 'react';
import Card from 'material-ui/lib/card/card';
import CardHeader from 'material-ui/lib/card/card-header';
import CardText from 'material-ui/lib/card/card-text';

const style = {
    marginBottom: "20px"
};

var Time = React.createClass({
    getInitialState: function() {
        return {countdown: this.getTime()};
    },

    componentDidMount: function() {
        setInterval(this.updateTime, 1000);
    },

    updateTime: function() {
        this.setState({countdown: this.getTime()});
    },

    getTime: function() {
        var now = moment().tz("Europe/Oslo");
        var next = moment.tz("13 37", "HH mm", "Europe/Oslo");

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
