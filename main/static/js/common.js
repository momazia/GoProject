/**
 * Used for validating email address for correct and empty
 * @param username
 */
function validateEmail(username) {
    var email = $(username).val();
    if (email != "") {
        var isEmail = isValidEmail(email);
        if (!isEmail) {
            showError("Username not valid");
        } else {
            removeError();
        }
    } else {
        showError("Username cannot be blank");
    }
}

/**
 * Validating the password field for empty only
 * @param password
 */
function validatePassword(password) {
    var pwd = $(password).val();
    if($("#email").val().length ==0)
    {
        showError("Email cannot be empty");
    }
    else if(pwd == "") {
        showError("Password cannot be empty");
    } else {
        removeError();
    }
}
/**
 * Validating the password field for empty only
 * @param password
 */
function validatePasswordNoEmail(password) {
    var pwd = $("#password").val().length;
    console.log("value is:"+pwd);
    if(pwd == 0) {
        showError("Password cannot be empty");
    } else {
        removeError();
    }
}

/**
 * Checks whether email is correctly typed as a email or not
 * @param email
 * @returns {boolean}
 */
function isValidEmail(email) {
    var regex = /^([a-zA-Z0-9_.+-])+\@(([a-zA-Z0-9-])+\.)+([a-zA-Z0-9]{2,4})+$/;
    return regex.test(email);
}

/**
 * Displays error passed in as parameter and disables the submit button
 * @param message
 */
function showError(message) {
    $(".errorMessage").text(message);
    $(".errorMessage").show();
    $('.submit').bind('click', function (e) {
        e.preventDefault();
    });
}

/**
 * Removes the error message and make submit clickable
 */
function removeError() {
    $(".errorMessage").hide();
    $(".submit").unbind('click');
}

function isPasswordSame() {
    validatePassword($("#password").val());
    if($("#password").val() != $("#retryPassword").val()) {
        showError("Password should be same");
    } else {
        if($("#retryPassword").val() == "") showError("Password cannot be empty");
        else {
            showError("Passwords match");
            removeError();
        }
    }
}
function isPasswordSameNoEmail() {
    validatePasswordNoEmail($("#password").val());
    if($("#password").val() != $("#retryPassword").val()) {
        showError("Password should be same");
    } else {
        if($("#retryPassword").val() == "") showError("Password cannot be empty");
        else {
            showError("Passwords match");
            removeError();
        }
    }
}