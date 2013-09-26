$(function() {
  'use strict';
  $("#datatable").dataTable({
    "bProcessing": true,
    "bDeferRender": true,
    "aaData": aaData,
    "aoColumns": aoColumns,
    "sDom": "<'row'<'span6'l><pf>r>t<'row'<'span6'i>>",                                                                                                   
    "sPaginationType": "bootstrap",                                                                                                                                       
    "oLanguage": {                                                                                                                                                        
      "sLengthMenu": "_MENU_ records per page"                                                                                                                            
    }                                                                                                                                                                     
  });
});
