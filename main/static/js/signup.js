$(document).ready(function() {

	$("#email").on("focusout", function() {
		validateEmail(this);
		if(!$(".errorMessage").is(":visible")) {
			isUserTaken($("#email").val());
		}
	});
	$("#email").on("focusin", function() {
		$(".errorMessage").hide();
	});
	$("#password").on("focusout", function() {
		validatePassword(this);
	});
	$("#password").on("focusin", function() {
		validateEmail($("#email").val());
	});
	$("#retryPassword").on("focusin", function() {
		validatePassword($("#password").val());
	});
	$("#retryPassword").on("focusout", function() {
		isPasswordSame();
	});
	$("#fname").on("focusiin", function() {
		isPasswordSame();
	});
	$("#fname").on("focusout", function() {
		if($(this).val() == "") showError("First Name cannot be blank");
	});
	$("#lname").on("focusin", function() {
		isPasswordSame();
		if($(this).val() == "") showError("First Name cannot be blank");

	});
	$("#lname").on("focusout", function() {
		if($(this).val() == "") showError("Last Name cannot be blank");
	});
	$("#submit").on("mousedown", function(e) {
		validateEmail($("#email"));
		if(!$(".errorMessage").is(":visible")) {
			validatePassword($("#password"));
			if(!$(".errorMessage").is(":visible")) {
				isPasswordSame();
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
		}
	});
	$("#resetBtn").on("click", function() {
		removeError();
		$('signUpForm').reset()
	});
});

function isUserTaken(username) {
	$.ajax({
		url : "/signup/isusertaken",
		data : {
			"username" : username
		},
		success : function(result) {
			if (result.includes('true')) {
				$('.errorMessage').show();
				$('.errorMessage').text('User is already taken !!');
			} else {
				$('.errorMessage').show();
				$('.errorMessage').text('Username available !!');
			}
		}
	});
}