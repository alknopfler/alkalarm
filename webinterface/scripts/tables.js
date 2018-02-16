$("#alarm").click(function () {
$.ajax({
    url: 'http://alknopfler.ddns.net:8080/alarm',
    type: "get",
    dataType: "json",
    data: '',
    success: function(data, textStatus, jqXHR) {
        // since we are using jQuery, you don't need to parse response
        drawTable(data);
    }
});

function drawTable(data) {
    for (var i = 0; i < data.length; i++) {
        drawRow(data[i]);
    }
}

function drawRow(rowData) {
    var row = $("<tr />")
    $("#alarm").append(row); //this will append tr element to table... keep its reference for a while since we will add cels into it
    row.append($("<td>" + rowData.Date + "</td>"));
    row.append($("<td>" + rowData.Sensor + "</td>"));
}
