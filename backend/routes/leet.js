var express = require('express');
var router = express.Router();
var config = require('../config');

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

    res.json({
        success: true
    })
});

module.exports = router;
