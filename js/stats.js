function GetValidsForUser(username)
{
    var api = "http://thorium.skriveleif.com:9200/irc-leet/0/_search?search_type=count";

    $.ajax(api, {
        data: JSON.stringify({
            "query": {
                "query_string": {
                    "default_field": "nick",
                    "query": username
                }
            },
            "facets": {
                "delay": {
                    "statistical": {
                        "field": "delay"
                    }
                }
            }
        }),
        success: function(data)
        {
            ShowStats(data);
        },
        type: 'POST'
    });
}

function ShowStats(data)
{
    var d = data.facets.delay;
    
    
    
    
}