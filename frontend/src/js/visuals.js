$(document).ready(function(){
    var d = new Date();
    d.setHours(13);
    d.setMinutes(37);
    d.setSeconds(00);

    for (var i = 0; i <= 5; i++)
    {

        switch (i)
        {
            case 0:
                $('.daylist-divider').before("<li id='"+ d+"' onclick='GetData(this)'><a>Today</a></li>");
                break;
            case 1:
                $('.daylist-divider').before("<li id='"+d+"' onclick='GetData(this)'><a>Yesterday</a></li>");
                break;
            default:
                $('.dropdown-menu').append("<li id='"+d+"' onclick='GetData(this)'><a>"+
                    addZ(d.getDate()) +" "+ addZ(d.getMonth()) +" "+ d.getFullYear()
                    +"</a></li>");
        }

        d.setDate(d.getDate() - 1);
    }
});

function addZ(n){return n<10? '0'+n:''+n;}
