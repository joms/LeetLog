$(document).ready(function(){
    GetData();
});

var data;

function GetData(day)
{
    var api = "http://thorium.skriveleif.com:9200/irc-leet/_search";

    var date = new Date();
    if (day)
    {
        date = new Date(day.id);

        console.log(day);
        $('.dropdown-toggle').text($(day).text());
    }
    date.setMilliseconds(0);
    date.setSeconds(0);
    date.setMinutes(30);
    date.setHours(13);

    $.ajax(api, {
        data: JSON.stringify({
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
                            "should": [
                                {
                                    "query_string": {
                                        "query": "_type=0"
                                    }
                                },
                                {
                                    "query_string": {
                                        "query": "_type=2"
                                    }
                                },
                                {
                                    "query_string": {
                                        "query": "_type=3"
                                    }
                                },
                                {
                                    "query_string": {
                                        "query": "_type=5"
                                    }
                                },
                                {
                                    "query_string": {
                                        "query": "_type=6"
                                    }
                                }
                            ]
                        }
                    },
                    "filter": {
                        "range": {
                            "@timestamp": {
                                "gte": date.getTime(),
                                "lte": date.getTime(date.setMinutes(45))
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
        }),
        success: function(_data)
        {
            $('.list-group').empty();
            data = _data.hits;
            AddScore(data);
        },
        type: 'POST'
    });
}

function AddScore(d)
{
    if (d.total > 0)
    {
        $.each(d.hits, function(index){
            var date = new Date(this.sort[0]);

            var c = (date.getMinutes() == 37) ? "alert-success" : "alert-danger";

            var str = date.getSeconds() +"."+ date.getMilliseconds();
            var miss = "";
            if (date.getMinutes() < 37)
            {
                miss = "-"+date.getHours()+"."+date.getMinutes()+".";
                str = str.substr(0, str.length - 4);
            }
            if (date.getMinutes() > 37)
            {
                miss = "+"+date.getHours()+"."+date.getMinutes()+".";
                str = str.substr(0, str.length - 4);
            }

            var f = "<li class='list-group-item'> <span class='badge "+c+"'>"+
                miss + str
                +"</span>"+this.fields.nick[0]+"</li>";
            $('.list-group').append(f);

        });
    } else {
        $('#ErrorModalLabel').text('No leet for you!');
        $('.modal-body').text('Please select another day!');
        $('#ErrorModal').modal();
    }
}