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
    if(pwd == "") {
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
    $(".errorMessage").css("color", "#D8000C");
    $(".errorMessage").css("backgroundColor", "#FFBABA");
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

/**
 * Checks if the password and retry password entered are same.
 */
function isPasswordSame(displaySuccess) {
    if($("#password").val() != $("#retryPassword").val()) {
        showError("Password should be same");
    } else {
        console.log("lol1");
        if($("#retryPassword").val() == "") showError("Password cannot be empty");
        else {
            if (displaySuccess) showSuccess("Passwords match");
        }
    }
}

function showSuccess(message) {
    $(".errorMessage").css("color", "#436a3d");
    $(".errorMessage").css("backgroundColor", "#90f780");
    $(".errorMessage").text(message);
    $(".errorMessage").show();

}