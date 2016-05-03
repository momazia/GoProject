$(document).ready(function() {

	$("#email").on("focusout", function() {
		var username = $(this).val();
		if (username != "") {
			$.ajax({
				url : "/signup/isusertaken",
				data : {
					"username" : username
				},
				success : function(result) {
					if (result.includes('true')) {
						$('.errorMessage').show();
						$('.errorMessage').text('User is already taken');
					} else {
						$('.errorMessage').hide();
						$('.errorMessage').text('');
					}
				}
			});
		}
	});

	$("#resetBtn").on("click", function() {
		$('.errorMessage').hide();
		$('.errorMessage').text('');
		$('signUpForm').reset()
	});
});