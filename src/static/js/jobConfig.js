$(function(){
    $('#myFormSubmit').click(function(e){
      console.log("my xuan fu ye");
      e.preventDefault();
      alert($('#myField').val());
      /*
      $.post('http://path/to/post', 
         $('#myForm').serialize(), 
         function(data, status, xhr){
           // do something here with response;
         });
      */
    });
});
