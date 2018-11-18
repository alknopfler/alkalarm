$(document).ready(function () {
             $("#saveMail").click(function () {
                         var sensor = $('#inputjson').val();
                         $.ajax({
                             async: true,
                             crossDomain: true,
                             url: 'http://{{HOST}}:8080/setup/mail',
                             type: 'POST',
                             dataType: 'json',
                             data: sensor,
                             success: function (data, textStatus, xhr) {
                                 window.alert("Control loaded successfully");
                                 location.reload();
                             },
                             error: function (data){
                                 window.alert("Remember to stop the alarm before...");
                             }
                         });
             });
             $("#deleteMail").click(function () {
                          var receptor = $('#receptor').val();
                          var urldelete = 'http://{{HOST}}:8080/setup/mail/'+receptor;
                          $.ajax({
                              async: true,
                              url: urldelete,
                              type: 'delete',
                              success: function () {
                                 window.alert("Mail delete successfully");
                                 location.reload();
                              },
                              error: function (data){
                                   window.alert("Remember to stop the alarm before...");
                              }
                          });
             });
             $("#example").click(function () {
                   document.getElementById("inputjson").value = "[{\"receptor\":\"alknopfler@gmail.com\"}]";
             });
             $("#mailtable").show(function (){
                     $.ajax({
                         url: 'http://{{HOST}}:8080/mails',
                         type: "get",
                         dataType: "json",
                         data: '',
                         success: function(data, textStatus, jqXHR) {
                             // since we are using jQuery, you don't need to parse response
                             for (var i = 0; i < data.length; i++) {
                                 var row = $("<tr />");
                                 $("#mailtable").append(row); //this will append tr element to table... keep its reference for a while since we will add cels into it
                                 row.append($("<td>" + data[i] + "</td>"));

                             }
                         }
                     });
             });
 });