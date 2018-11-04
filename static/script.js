  /*$(document).ready(function(){
        //$("button").hover(function(){
            $.getJSON('/api', { get_param: 'fname' }, function(data) {

                var items = [];
                console.log(data)
                $.each( data, function( key, val ) {
                    items.push( "<li id='" + key + "'>" + val + "</li>" );
                });
                console.log(items);

                $( "<ul/>", {
                    "class": "my-new-list",
                    html: items.join( "" )
                }).appendTo( "body" );
            });
        //});
    });
    */
    function addToList(){
        var param = $("#input1").val()
        console.log(param);
        $.getJSON('/api', { get_param: param }, function(data) {

            var items = [];
            $.each( data, function( key, val ) {
              console.log(key, val)
                items.push( "<li class='id='" + key + "'>" + val + "</li>" );
            });

            $( "<ul/>", {
                "class": "my-new-list",
                html: items.join( "" )
            }).appendTo( "body" );
        });
    }


    $(document).ready(function(){
        $("button").click(function(){
            $("li").remove(".person, .demo");
        });
    });
