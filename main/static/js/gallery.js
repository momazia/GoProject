$(document).ready(function() {
    $("#file").change(function () {
        validateFileFormat(this);
    });
    $(".reset").on("click", function() {
       removeError();
    });
    $(".submit").on("mousedown", function(event) {
        validateFileFormat($("#file"));
        if(!$(".errorMessage").is(":visible")) {
            $("#galleryForm").submit();
        }
    });
});