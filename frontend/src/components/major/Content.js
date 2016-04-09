import React from 'react';

import ResultList from './../minor/ResultList';

var Content = React.createClass({
    componentDidMount: function() {
        this.getData();
    },

    getInitialState: function() {
        return {leets: []};
    },

    getData: function() {
        var from = moment.tz("00 00", "HH mm", "Europe/Oslo");
        var to = moment.tz("23 59", "HH mm", "Europe/Oslo");

        var request = new XMLHttpRequest();
        request.open('GET', 'http://46.101.16.97/api/leet/'+from+'/'+to, true);

        var that = this;

        request.onload = function() {
            if (request.status >= 200 && request.status < 400) {
                // Success!
                var data = JSON.parse(request.responseText);
                if (data.success == true) {
                    that.setState({"leets": data.result.hits.hits});
                } else {
                    // There was a connection error of some sort
                }
            } else {
                // We reached our target server, but it returned an error

            }
        };

        request.onerror = function() {
            // There was a connection error of some sort
        };

        request.send();
    },

    render: function() {
        return (
            <div>
                {/*<div>
                    <Time />
                </div>*/}
                <div>
                    <ResultList leets={this.state.leets} />
                </div>
            </div>
        );
    }
});

export default Content;