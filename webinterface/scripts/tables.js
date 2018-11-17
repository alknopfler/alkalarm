$(document).ready(function () {
    $("#alarm").show(function (){
        $.ajax({
            url: 'http://{{HOST}}:8080/alarm',
            type: "get",
            dataType: "json",
            data: '',
            success: function(data, textStatus, jqXHR) {
                // since we are using jQuery, you don't need to parse response
                for (var i = 0; i < data.length; i++) {
                    var row = $("<tr />");
                    $("#alarm").append(row); //this will append tr element to table... keep its reference for a while since we will add cels into it
                    row.append($("<td>" + data[i].Date + "</td>"));
                    row.append($("<td>" + data[i].Sensor + "</td>"));
                }
            }
        });
    });
    $("#clearAlarm").click(function () {
                         $.ajax({
                             async: true,
                             url: 'http://{{HOST}}:8080/alarm',
                             type: 'delete',
                             success: function () {
                                $("#alarm").load(" #alarm")
                             }
                         });
    });

});