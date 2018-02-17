$(document).ready(function () {

             $("#saveSensor").click(function () {
                 var code = prompt("Please enter your access code:", "");
                 $.ajax({
                      url: 'http://alknopfler.ddns.net:8080/admin/'+code,
                      type: "get",
                      dataType: "json",
                      data: '',
                      error: function(data, textStatus, jqXHR) {
                            window.alert("Your Access Password is incorrect...");
                            location.reload();
                      },
                      success: function(data, textStatus, jqXHR) {
                         var sensor = $('#inputjson').val();
                         $.ajax({
                             async: true,
                             crossDomain: true,
                             url: 'http://alknopfler.ddns.net:8080/setup/sensor',
                             type: 'POST',
                             dataType: 'json',
                             data: sensor,
                             success: function (data, textStatus, xhr) {
                                 window.alert("Sensor loaded successfully");
                                 location.reload();
                             },
                             error: function (data){
                                 window.alert("Remember to stop the alarm before...");
                             }
                         });
                      }
                 });
             });
             $("#deleteSensor").click(function () {
                   var code = prompt("Please enter your access code:", "");
                   $.ajax({
                        url: 'http://alknopfler.ddns.net:8080/admin/'+code,
                        type: "get",
                        dataType: "json",
                        data: '',
                        error: function(data, textStatus, jqXHR) {
                              window.alert("Your Access Password is incorrect...");
                              location.reload();
                        },
                        success: function(data, textStatus, jqXHR) {
                          var code = $('#code').val();
                          var urldelete = 'http://alknopfler.ddns.net:8080/setup/sensor/'+code
                          $.ajax({
                              async: true,
                              url: urldelete,
                              type: 'delete',
                              success: function () {
                                 window.alert("Sensor delete successfully");
                                 location.reload();
                              },
                              error: function (data){
                                   window.alert("Remember to stop the alarm before...");
                              }
                          });
                        }
                   });
             });
             $("#example").click(function () {
                                           document.getElementById("inputjson").value = "[{\"Code\":\"3462404\",\"TypeOf\":\"presence\",\"Zone\":\"salon\"}]";
                                       });
             $("#sensortable").show(function (){
                     $.ajax({
                         url: 'http://alknopfler.ddns.net:8080/setup/sensor',
                         type: "get",
                         dataType: "json",
                         data: '',
                         success: function(data, textStatus, jqXHR) {
                             // since we are using jQuery, you don't need to parse response
                             for (var i = 0; i < data.length; i++) {
                                 var row = $("<tr />")
                                 $("#sensortable").append(row); //this will append tr element to table... keep its reference for a while since we will add cels into it
                                 row.append($("<td>" + data[i].Code + "</td>"));
                                 row.append($("<td>" + data[i].TypeOf + "</td>"));
                                 row.append($("<td>" + data[i].Zone + "</td>"));

                             }
                         }
                     });
                 });
 });