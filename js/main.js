(function() {
  var app = angular.module('sample', ['ngResource']).
    config(function($routeProvider) {
      $routeProvider.
        when('/', {controller:ListController, templateUrl:'list.html'}).
        when('/entry', {controller:EntryController, templateUrl:'entry.html'});
    });

  function ListController($scope, $resource) {
    var Store = $resource("/store");
    $scope.stores = Store.query(function() {
      console.log("success store query");
    }, function(){
      console.log("error store query");
    });
  }

  function EntryController($scope, $resource){
    $scope.categories = [{"id" : "1", "name" : "野菜"}];
    var List = $resource("/item/list");
    var Store = $resource("/store");

    $scope.changeCategory = function() {
      $scope.items = List.query({id : $scope.entryForm.categoryid}, function(){
        console.log("success list");
      }, function(){
        console.log("error list");
      });
    };

    $scope.submit = function($event) {
      console.log($scope.entryForm);
      Store.save($scope.entryForm, function(){
        console.log("success entry");
      }, function(){
        console.log("error entry");
      });
    };
  }
})();