$(document).ready(function () {
             $("#saveControl").click(function () {
                     var sensor = $('#inputjson').val();
                     $.ajax({
                         async: true,
                         crossDomain: true,
                         url: 'http://{{HOST}}:8080/setup/control',
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
                $("#scanControl").click(function () {
                    $('#inputjson').val("Scanning...Press any Control button!");
                    $.ajax({
                        async: true,
                        crossDomain: true,
                        url: 'http://{{HOST}}:8080/scan/control',
                        type: "get",
                        dataType: "json",
                        data: '',
                        success: function (data) {
                            $('#inputjson').val('[{"Code":"'+data.Code+'","Description":"mando1","TypeOf":"inactive"}]');
                        },
                        error: function (data){
                            window.alert("Control key not pressed");
                        },
                    });
                });
             $("#deleteControl").click(function () {
                  var code = $('#code').val();
                  var urldelete = 'http://{{HOST}}:8080/setup/control/'+code;
                  $.ajax({
                      async: true,
                      url: urldelete,
                      type: 'delete',
                      success: function () {
                         window.alert("Control delete successfully");
                         location.reload();
                      },
                      error: function (data){
                           window.alert("Remember to stop the alarm before...");
                      }
                  });
             });
             $("#example").click(function () {
                              document.getElementById("inputjson").value = "[{\"Code\":\"3462412\",\"Description\":\"mando1\",\"TypeOf\":\"inactive\"},{\"Code\":\"3462448\",\"Description\":\"mando1\",\"TypeOf\":\"full\"},{\"Code\":\"3462592\",\"Description\":\"mando1\",\"TypeOf\":\"partial\"}]";
                          });
             $("#controltable").show(function (){
                     $.ajax({
                         url: 'http://{{HOST}}:8080/controls',
                         type: "get",
                         dataType: "json",
                         data: '',
                         success: function(data, textStatus, jqXHR) {
                             // since we are using jQuery, you don't need to parse response
                             for (var i = 0; i < data.length; i++) {
                                 var row = $("<tr />");
                                 $("#controltable").append(row); //this will append tr element to table... keep its reference for a while since we will add cels into it
                                 row.append($("<td>" + data[i].Code + "</td>"));
                                 row.append($("<td>" + data[i].Description + "</td>"));
                                 row.append($("<td>" + data[i].TypeOf + "</td>"));

                             }
                         }
                     });
                 });
 });