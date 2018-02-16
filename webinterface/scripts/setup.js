$(document).ready(function () {
 var code = prompt("Please enter your code:", "");

 jQuery.get('../../.htaccess', function(data) {
 if (code ==data) {

             $("#saveSensor").click(function () {
                 var sensor = $('#sensor').val();
                 $.ajax({
                     async: true,
                     crossDomain: true,
                     url: 'http://alknopfler.ddns.net:8080/setup/sensor',
                     type: 'POST',
                     dataType: 'json',
                     data: sensor,
                     success: function (data, textStatus, xhr) {
                         console.log(data);
                     },
                 });
             });
             $("#saveControl").click(function () {
                  var control = $('#control').val();
                  $.ajax({
                      async: true,
                      crossDomain: true,
                      url: 'http://alknopfler.ddns.net:8080/setup/control',
                      type: 'POST',
                      dataType: 'json',
                      data: control,
                      success: function (data, textStatus, xhr) {
                          console.log(data);
                      },
                  });
             });
             $("#saveMailer").click(function () {
                   var mailer = $('#mailer').val();
                   $.ajax({
                       async: true,
                       crossDomain: true,
                       url: 'http://alknopfler.ddns.net:8080/setup/mail',
                       type: 'POST',
                       dataType: 'json',
                       data: mailer,
                       success: function (data, textStatus, xhr) {
                           console.log(data);
                       },
                   });
             });

 }
 });
 });