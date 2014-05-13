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

        $('.dropdown-toggle').text($(day).text());
        $('.dropdown-toggle').append("<b class='caret'></b>");
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
                            "must": [
                                {
                                    "terms": {
                                        "_type": [
                                            "0",
                                            "2",
                                            "3",
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
        $('.list-group').empty();
        $.each(d.hits, function(index){
            var date = new Date(this.sort[0]);

            var c = (date.getMinutes() == 37) ? "alert-success" : "alert-danger";

            var str = addZ(date.getSeconds()) +"."+ addZ2(date.getMilliseconds());

            if (date.getMinutes() != 37)
            {
                str = (date.getMinutes() <= 36 ? "- ":"+ ") + date.getMinutes()+"."+str;

                var x = date.getMinutes() +"."+ date.getSeconds();
                if (x != 36.59 && x != 38.00) str = str.substr(0, str.length - 4);
            }

            var f = "<li class='list-group-item'> <span class='badge "+c+"'>"+
                str
                +"</span>"+this.fields.nick[0]+"</li>";
            $('.list-group').append(f);

        });
    } else {
        $('#ErrorModalLabel').text('No leet for you!');
        $('.modal-body').text('Please select another day!');
        $('#ErrorModal').modal();
    }
}