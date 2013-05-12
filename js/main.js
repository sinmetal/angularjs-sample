function MainController($scope, $http) {
	$scope.init = function() {
		$scope.categories = [{"id" : "1", "name" : "野菜"}];
	};

	$scope.changeCategory = function() {
		$http({
			method: 'GET', 
			url: '/item/list',
			param: $scope.categoryid
		}).
  		success(function(data, status, headers, config) {
  			$scope.items = data;
  		}).
		error(function(data, status, headers, config) {
			console.log('error');
		});
	};

	$scope.submit = function() {
		$http({
			method: 'POST', 
			url: '/store/entry',
			param: $scope.entryForm
		}).
  		success(function(data, status, headers, config) {
  			console.log('success');
  		}).
		error(function(data, status, headers, config) {
			console.log('error');
		});
	};
}