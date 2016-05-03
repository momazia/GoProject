$(document).ready(function() {
    $("#username").on("focusout", function() {
        var email = $(this).val();
        console.log("Email: " + email);
        var isEmail = validateEmail(email);
        console.log("IsEmail: " + typeof isEmail);
        if (!isEmail) {
            console.log("here 1!!");
            $(".errorMessage").text("Username not valid");
            $(".errorMessage").show();
            $('.login').bind('click', function(e){
                e.preventDefault();
            })
        } else {
            $(".errorMessage").hide();
            $(".login").unbind('click')
        }
    });
});
function validateEmail(email) {
    var regex = /^([a-zA-Z0-9_.+-])+\@(([a-zA-Z0-9-])+\.)+([a-zA-Z0-9]{2,4})+$/;
    return regex.test(email);
}