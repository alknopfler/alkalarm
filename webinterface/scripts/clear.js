$("#clear").click(function () {
                     $.ajax({
                         async: true,
                         crossDomain: true,
                         url: 'http://alknopfler.ddns.net:8080/alarm',
                         type: 'DELETE',
                         success: function () {
                            //if (data.success){
                                   $("#table").load();
                                    console.log("entra success")
                            //}
                         }
                     });
                 });