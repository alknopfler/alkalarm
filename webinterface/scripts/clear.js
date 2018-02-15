 $.ajax({
     async: true,
     crossDomain: true,
     url: 'http://alknopfler.ddns.net:8080/alarm',
     type: 'DELETE',
     success: function (data, textStatus, xhr) {
         $('#table').reload()
     },
     error: function (xhr, textStatus, errorThrown) {
         console.log('Error in Operation');
     }
 });
