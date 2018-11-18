$(document).ready(function () {

             $("#saveSensor").click(function () {
                         var sensor = $('#inputjson').val();
                         $.ajax({
                             async: true,
                             crossDomain: true,
                             url: 'http://{{HOST}}:8080/setup/sensor',
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
             });
            $("#scanSensor").click(function () {
                $('#inputjson').val("Scanning...Press button or enable sensor detection!");
                $.ajax({
                    async: true,
                    crossDomain: true,
                    url: 'http://{{HOST}}:8080/scan/sensor',
                    type: "get",
                    dataType: "json",
                    data: '',
                    success: function (data) {
                        $('#inputjson').val('[{"Code":"'+data.Code+'","TypeOf":"presence","Zone":"salon"}]');
                    },
                    error: function (data){
                        window.alert("Sensor not found or key not pressed");
                    },
                });
             });
             $("#deleteSensor").click(function () {
                          var code = $('#code').val();
                          var urldelete = 'http://{{HOST}}:8080/setup/sensor/'+code;
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
             });
             $("#example").click(function () {
                                           document.getElementById("inputjson").value = "[{\"Code\":\"3462404\",\"TypeOf\":\"presence\",\"Zone\":\"salon\"},{\"Code\":\"3462405\",\"TypeOf\":\"aperture\",\"Zone\":\"entrada\"},{\"Code\":\"3462406\",\"TypeOf\":\"other\",\"Zone\":\"hall\"}]";
                                       });
             $("#sensortable").show(function (){
                     $.ajax({
                         url: 'http://{{HOST}}:8080/sensors',
                         type: "get",
                         dataType: "json",
                         data: '',
                         success: function(data, textStatus, jqXHR) {
                             // since we are using jQuery, you don't need to parse response
                             for (var i = 0; i < data.length; i++) {
                                 var row = $("<tr />");
                                 $("#sensortable").append(row); //this will append tr element to table... keep its reference for a while since we will add cels into it
                                 row.append($("<td>" + data[i].Code + "</td>"));
                                 row.append($("<td>" + data[i].TypeOf + "</td>"));
                                 row.append($("<td>" + data[i].Zone + "</td>"));

                             }
                         }
                     });
                 });
 });