import React from 'react';

import Time from './../minor/Time';
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
        request.open('POST', 'http://thorium.skriveleif.com:9200/irc-leet/_search', true);

        var that = this;

        request.onload = function() {
            if (request.status >= 200 && request.status < 400) {
                // Success!
                var data = JSON.parse(request.responseText);
                that.setState({"leets": data.hits.hits});
            } else {
                // We reached our target server, but it returned an error

            }
        };

        request.onerror = function() {
            // There was a connection error of some sort
        };

        request.send(JSON.stringify({
            "fields": [
                "time",
                "status",
                "nick",
                "msg"
            ],
            "query": {
                "filtered": {
                    "query": {
                        "bool": {
                            "must": [
                                {
                                    "terms": {
                                        "_type": [
                                            "0",
                                            "5",
                                            "6"
                                        ]
                                    }
                                }
                            ]
                        }
                    },
                    "filter": {
                        "range": {
                            "@timestamp": {
                                "gte": from,
                                "lte": to
                            }
                        }
                    }
                }
            },
            "size": 500,
            "sort": [
                {
                    "@timestamp": {
                        "order": "asc"
                    }
                }
            ]
        }));
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