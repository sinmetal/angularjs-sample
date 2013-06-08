(function() {
  /**
   * 上部ナビゲーション設定
   * 
   * @param app angular module
   */
  var setNaviTopDirective = function(app) {
    app.directive('navitop', function() {
      var directiveDefinitionObject = {
        priority: 0,
        templateUrl: '/topmenu.html',
        replace: false,
        transclude: false,
        restrict: 'E',
        scope: false,
        controller: ['$scope', '$location', function($scope, $location) {
          // var urlFragments = $location.absUrl().split('/');
          // var menuCategoryPathPosition = 1;
          // $scope.current = urlFragments[menuCategoryPathPosition];
        }]
      };
      return directiveDefinitionObject
    });
  };

  var app = angular.module('sample', ['ngResource']).
    config(function($routeProvider) {
      $routeProvider.
        when('/', {controller:'ListController', templateUrl:'list.html'}).
        when('/entry', {controller:'EntryController', templateUrl:'entry.html'});
    });

  setNaviTopDirective(app);

  app.controller('ListController', ['$scope', '$resource', function($scope, $resource) {
    var Store = $resource("/store");
    $scope.stores = Store.query(function() {
      console.log("success store query");
    }, function(){
      console.log("error store query");
    });
  }]);

  app.controller('EntryController', ['$scope', '$location', '$resource', function($scope, $location, $resource) {
    $scope.categories = [{"id" : "1", "name" : "野菜"}];
    var List = $resource("/item/list");
    var Store = $resource("/store");

    $scope.changeCategory = function() {
      $scope.items = List.query({id : $scope.entryForm.categoryid}, function(){
        console.log("success list");
      }, function(){
        console.log("error list");
      });
    }

    $scope.submit = function($event) {
      console.log($scope.entryForm);
      Store.save($scope.entryForm, function(){
        console.log("success entry");
        $location.path('/');
      }, function(){
        console.log("error entry");
      });
    };
  }]);
})();