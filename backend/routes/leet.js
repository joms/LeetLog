var express = require('express');
var router = express.Router();

var elasticsearch = require('elasticsearch');
var client = new elasticsearch.Client({
    host: 'thorium.skriveleif.com:9200',
    log: 'trace'
});

router.route('').get(function(req,res){
    res.json({
        success: "true"
    });
});

module.exports = router;