$(function() {
  'use strict';
  $("#datatable").dataTable({
    "bProcessing": true,
    "bDeferRender": true,
    "aaData": aaData,
    "aoColumns": aoColumns,
    "sDom": "<'row'<'span6'l><'span6'f>r>t<'row'<'span6'i><'span6'p>>",                                                                                                   
    "sPaginationType": "bootstrap",                                                                                                                                       
    "oLanguage": {                                                                                                                                                        
      "sLengthMenu": "_MENU_ records per page"                                                                                                                            
    }                                                                                                                                                                     
  });
});
