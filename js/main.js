var app = angular.module("sample", ["ngResource"]);
app.controller("MainController", function($scope, $resource){
  $scope.categories = [{"id" : "1", "name" : "野菜"}];
  var List = $resource("/item/list");
  var Entry = $resource("/store/entry");
  $scope.changeCategory = function() {
    $scope.items = List.query({id : $scope.entryForm.categoryid}, function(){
      console.log("success list");
    }, function(){
      console.log("error list");
    });
  };
  $scope.submit = function($event) {
    Entry.save($scope.entryForm, function(){
      console.log("success entry");
    }, function(){
      console.log("error entry");
    });
  };
});
