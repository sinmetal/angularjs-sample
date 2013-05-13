var app = angular.module("sample", ["ngResource"]);
app.factory("List", function($resource){
  return $resource("/item/list");
});
app.factory("Entry", function($resource){
  return $resource("/store/entry");
});
app.controller("MainController", function($scope, $resource, List, Entry){
  $scope.categories = [{"id" : "1", "name" : "野菜"}];
  $scope.changeCategory = function() {
    List.query({id : $scope.entryForm.categoryid}, function(data){
      $scope.items = data;
    }, function(){
      console.log("error");
    });
  };
  $scope.submit = function() {
    Entry.save($scope.entryForm, function(){
      console.log("success");
    }, function(){
      console.log("error");
    });
  };
});
