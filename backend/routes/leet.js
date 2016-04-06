var express = require('express');
var router = express.Router();

var elasticsearch = require('elasticsearch');
var client = new elasticsearch.Client({
    host: 'thorium.skriveleif.com:9200',
    log: 'trace'
});

router.route('/:from/:to').get(function(req,res){
    client.search({
        index: 'irc-leet',
        // type: 'tweets',
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

module.exports = router;
