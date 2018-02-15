
         $(document).ready(function () {
             var code = prompt("Please enter your code:", "");

             jQuery.get('../../.htaccess', function(data) {
             if (code ==data) {
                 $.get("http://alknopfler.ddns.net:8080/status", function(respuesta){
                    if respuesta == "full"{
                         $("#activarFullAlarm").css({'color':'blue'});
                         $("#activarPartialAlarm").css({'color':'grey'});
                         $('#desactivarAlarm').css({'color':'grey'});
                   }else if respuesta == "partial"{
                         $("#activarFullAlarm").css({'color':'grey'});
                         $("#activarPartialAlarm").css({'color':'blue'});
                         $('#desactivarAlarm').css({'color':'grey'});
                   }else{
                         $("#activarFullAlarm").css({'color':'grey'});
                         $('#desactivarAlarm').css({'color':'blue'});
                         $("#activarPartialAlarm").css({'color':'grey'});
                   }
                 })
                 $("#activarFullAlarm").click(function () {
                     $.ajax({
                         async: true,
                         crossDomain: true,
                         url: 'http://alknopfler.ddns.net:8080/activate/full',
                         type: 'POST',
                         success: function () {
                            //if (data.success){
                                   $("#activarFullAlarm").css({'color':'blue'});
                                   $("#activarPartialAlarm").css({'color':'grey'});
                                   $('#desactivarAlarm').css({'color':'grey'});
                            //}
                         }
                     });
                 });
                 $("#activarPartialAlarm").click(function () {
                     $.ajax({
                         async: true,
                         crossDomain: true,
                         url: 'http://alknopfler.ddns.net:8080/activate/partial',
                         type: 'POST',
                         success: function () {
                            //if (data.success){
                                 $("#activarFullAlarm").css({'color':'grey'});
                                 $("#activarPartialAlarm").css({'color':'blue'});
                                 $('#desactivarAlarm').css({'color':'grey'});

                            //}
                         }
                     });
                 });
                 $("#desactivarAlarm").click(function () {
                     $.ajax({
                         async: true,
                         crossDomain: true,
                         url: 'http://alknopfler.ddns.net:8080/deactivate',
                         type: 'POST',
                         success: function () {
                            //if (data.success){
                                 $("#activarFullAlarm").css({'color':'grey'});
                                 $('#desactivarAlarm').css({'color':'blue'});
                                 $("#activarPartialAlarm").css({'color':'grey'});


                            //}
                         }
                     });
                 });
             }
             });
         });