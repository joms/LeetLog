var express = require('express');
var bodyParser = require('body-parser');
var app = express();

var leet = require('./routes/leet.js');

app.set('port', process.env.PORT || 8000);

var server = app.listen(app.get('port'), function() {
    console.log('Express server listening on port ' + server.address().port);
});

app.use(bodyParser.urlencoded({extended: false}));
app.use(bodyParser.json());

app.use("/api/leet", leet);