$(document).ready(function() {
    removeError();
    $("#password").on("focusin", function() {
        removeError();
    });
    $("#password").on("focusout", function() {
        validatePasswordNoEmail(this);
    });
    $("#retryPassword").on("focusin", function() {
        validatePasswordNoEmail($("#password").val());
    });
    $("#retryPassword").on("focusout", function() {
        isPasswordSameNoEmail();
    });
    $("#fname").on("focusiin", function() {
        isPasswordSameNoEmail();
    });
    $("#fname").on("focusout", function() {
        if($(this).val() == "") showError("First Name cannot be blank");
    });
    $("#lname").on("focusin", function() {
        isPasswordSameNoEmail();
        if($(this).val() == "") showError("First Name cannot be blank");

    });
    $("#lname").on("focusout", function() {
        if($(this).val() == "") showError("Last Name cannot be blank");
    });
    $("#submit").on("mousedown", function(e) {
            validatePasswordNoEmail($("#password"));
            if(!$(".errorMessage").is(":visible")) {
                isPasswordSameNoEmail();
                if(!$(".errorMessage").is(":visible")) {
                    if($("#fname").val() == "") {
                        showError("First Name cannot be blank");
                    } else {
                        removeError();
                        if($("#lname").val() == "") {
                            showError("Last Name cannot be blank");
                        } else {
                            removeError();
                        }
                    }
                }
            }

    });
    $("#resetBtn").on("click", function() {
        removeError();
        $('signUpForm').reset()
    });
});
