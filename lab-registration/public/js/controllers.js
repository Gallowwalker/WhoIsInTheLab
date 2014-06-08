var labRegistration = angular.module('lab-registration', ['ngSanitize', 'ui.select']);


labRegistration.controller('LabRegCtrl', ['$scope', '$http', function ($scope, $http) {

	$scope.device = {};
	$scope.user = {};
	$scope.user.selected = undefined;

	$http.get('/users').success(function(users) {
		$scope.users = users;
	});

	$http.get('/mac').success(function(data) {
		$scope.device.mac = data.mac;
	}).error(function(error) {
		$scope.device.mac = error.message;
	});

	$scope.submitDevice = function(device) {
		// TODO: call add device api
	};
}]);
