var express = require('express');
var bodyParser = require('body-parser');
var app = express();

var leet = require('./routes/leet.js');

var endpointKey = "abc123";

app.set('port', process.env.PORT || 8000);

app.use(function (req, res, next) {
    res.header("Access-Control-Allow-Origin", "*");
    res.header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept");
    next();
});

var server = app.listen(app.get('port'), function() {
    console.log('Express server listening on port ' + server.address().port);
});

app.use(bodyParser.urlencoded({extended: false}));
app.use(bodyParser.json());

app.use("/api/leet", leet);