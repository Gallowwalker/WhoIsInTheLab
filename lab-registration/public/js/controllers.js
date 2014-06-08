var labRegistration = angular.module('lab-registration', ['ngSanitize', 'ui.select']);


labRegistration.controller('LabRegCtrl', ['$scope', '$http', function ($scope, $http) {

	$scope.device = {};
	$scope.user = {};
	$scope.user.selected = undefined;
	$scope.owner = {};
	$scope.macValid = true;
	$scope.hasApiError = false;
	$scope.error = {}

	$http.get('/users').success(function(users) {
		$scope.users = users;
	}).error(function(error) {
		$scope.hasErrors(true, error.message);
	});

	$http.get('/mac').success(function(data) {
		$scope.hasErrors(false, {});
		$scope.device.MAC = data.mac;
	}).error(function(error) {
		$scope.hasErrors(true, error.message);
		$scope.macValid = false;
		$scope.device.MAC = error.message;
	});


	$scope.hasErrors = function(flag, errorMsg) {
		$scope.hasApiError = flag;
		$scope.success = false;
		$scope.errors = errorMsg;
	};

	$scope.submitDevice = function(device) {
		if($scope.user.selected === undefined) {
			$http.post('/users', $scope.owner).success(function(response) {
				$scope.hasErrors(false, {});
				$scope.apiAddDevice(device, response.id);
			}).error(function(error) {
				$scope.hasError(true, error.message);
			});
		} else {
			$http.put('/users/' + $scope.user.selected.id, $scope.owner).success(function(response) {
				$scope.hasErrors(false, {});
				$scope.apiAddDevice(device, response.id);
			}).error(function(error) {
				$scope.hasError(true, error.message);
			});
		}
	};


	$scope.apiAddDevice = function(device, userId) {
		$http.post('/users/' + userId + '/devices', device).success(function(response) {
			$scope.hasErrors(false, {});
			$scope.success = true;
		}).error(function(error) {
			$scope.hasErrors(true, error.message);
		});
	};

	$scope.ownerSelected = function(owner) {
		$http.get('/users/' + owner.selected.id).success(function(owner) {
			$scope.owner = owner;
			$scope.hasErrors(false, {});
		}).error(function(error) {
			$scope.hasErrors(true, error.message);
		});
	};
}]);
