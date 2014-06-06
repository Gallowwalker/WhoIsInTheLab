var labRegistration = angular.module('lab-registration', []);

labRegistration.controller('LabRegCtrl', ['$scope', '$http', function ($scope, $http) {

	$http.get('/users').success(function(data) {
		$scope.users = data;
	});
}]);
