var fs = require('fs');
var express = require('express');
var router = express.Router();
var config = require('../config');
var moment = require('moment');

var file = fs.createWriteStream(config.logDestination , {
    flags: 'a'
});

var elasticsearch = require('elasticsearch');
var client = new elasticsearch.Client({
    host: config.es.host,
    log: 'trace'
});

router.route('/:from/:to').get(function(req,res){
    client.search({
        index: 'irc-leet',
        body: {
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
                                "gte": req.params.from,
                                "lte": req.params.to
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
        }
    }, function(error, result, status) {
        if (status != 200) {
            res.json({
                success: false,
                reason: "Something wrong happened"
            });
        } else {
            res.json({
                success: true,
                result: result
            });
        }
    });
});

router.route('/').post(function(req,res){
    if (config.endpointKey != req.body.EndpointKey) {
        console.error("EndpointKey missmatch: "+req.body.EndpointKey +" vs "+ config.endpointKey +" from "+ req.connection.remoteAddress);

        res.json({
            success: false,
            reason: "Something wrong happened. Check your settings"
        });
        return;
    }

    file.write(
        req.body.Time +' '+
        req.body.Channel +' '+
        req.body.Status +' '+
        req.body.User.Nick +'\n'
    );

    res.json({
        success: true
    })
});

router.route('/top').get(function(req, res){
    if (config.endpointKey != req.query.EndpointKey) {
        console.error("EndpointKey missmatch: "+req.query.EndpointKey +" vs "+ config.endpointKey +" from "+ req.connection.remoteAddress);

        res.json({
            success: false,
            reason: "Something wrong happened. Check your settings"
        });
        return;
    }

    var date = moment().format('YYYY-MM-DD');

    if (typeof req.query.date != 'undefined') {
        if (moment(req.query.date).isValid()) {
            date = moment(req.query.date).format('YYYY-MM-DD');
             console.log(date);
        } else {
            res.json({
                success: false,
                reason: "Invalid date format"
            });
            return;
        }
    }

    client.search({
        index: 'irc-leet',
        body: {
            "query": {
                "filtered": {
                    "query": {
                        "bool": {
                            "must": [
                                {
                                    "terms": {
                                        "status": [
                                            "0"
                                        ]
                                    }
                                }
                            ]
                        }
                    },
                    "filter": {
                        "range": {
                            "@timestamp": {
                                "gte": date,
                                "lte": date
                            }
                        }
                    }
                }
            },
            "size": 3,
            "sort": [
                {
                    "@timestamp": {
                        "order": "asc"
                    }
                }
            ]
        }
    }, function(error, result, status) {
        if (status != 200) {
            res.json({
                success: false,
                reason: "Something wrong happened"
            });
        } else {
            if (result.hits.hits.length == 0) {
                res.json({
                    success: false,
                    reason: "No leets returned </3"
                });
                return;
            }

            var leets = [];
            for (var i = 0; i < result.hits.hits.length; i++) {
                leets.push({
                    "nick": result.hits.hits[i]._source.nick,
                    "delay": result.hits.hits[i]._source.delay
                });
            }

            res.json({
                success: true,
                result: leets
            });
        }
    });
});

module.exports = router;
