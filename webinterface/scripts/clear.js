$("#clear").click(function () {
 $.ajax({
     async: true,
     crossDomain: true,
     url: 'http://alknopfler.ddns.net:8080/alarm',
     type: 'DELETE',
     success: function () {
         $('#table').reload()
     }
 });
 });
