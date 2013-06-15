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
        templateUrl: '/html/common/topmenu.html',
        replace: false,
        transclude: false,
        restrict: 'E',
        scope: false,
        controller: ['$scope', '$location', function($scope, $location) {
          var urlFragments = $location.absUrl().split('/');
          // TODO ここでtop menu のactiveを判断する要素を作成する
          console.log(urlFragments);
          $scope.current = urlFragments[5];
        }]
      };
      return directiveDefinitionObject
    });
  };

  var app = angular.module('sample', ['ngResource']).
    config(function($routeProvider) {
      $routeProvider.
        when('/', {controller:'TopController', templateUrl:'/html/top/top.html'}).
        when('/guestbook/', {controller:'GuestBookListController', templateUrl:'/html/guestbook/list.html'}).
        when('/guestbook/entry', {controller:'GuestBookEntryController', templateUrl:'/html/guestbook/entry.html'});
    });

  setNaviTopDirective(app);

  app.directive('highlighttab', ['$location', function(location) {

    return {
        restrict: 'C',
        link: function($scope, $element, $attrs) {
            console.log($element.parent());

            var elementPath = $attrs.href.substring(1);
            $scope.$location = location;
            $scope.$watch('$location.path()', function(locationPath) {
                if (elementPath === locationPath) {
                    $element.parent().addClass("active");
                } else {
                    $element.parent().removeClass("active");
                }
            });
        }
    };
  }]);

  app.controller('TopController', ['$scope', '$resource', function($scope, $resource) {
  }]);

  app.controller('GuestBookListController', ['$scope', '$resource', function($scope, $resource) {
    var Favorite = $resource("/favorite");
    $scope.favos = Favorite.query(function() {
      console.log("success favorite query");
      console.log($scope.favos);
    }, function(){
      console.log("error favorite query");
    });
  }]);

  app.controller('GuestBookEntryController', ['$scope', '$location', '$resource', function($scope, $location, $resource) {
    $scope.elementTypes = [{"id" : "1", "name" : "くさ"},
                           {"id" : "2", "name" : "ほのお"},
                           {"id" : "3", "name" : "みず"}];
    var Pokemon = $resource("/pokemon");
    var Favorite = $resource("/favorite");

    $scope.changeElementType = function() {
      console.log($scope.entryForm.elementTypeId);
      $scope.pokemons = Pokemon.query({id : $scope.entryForm.elementTypeId}, function(){
        console.log("success list");
      }, function(){
        console.log("error list");
      });

    }

    $scope.submit = function($event) {
      console.log($scope.entryForm);
      var value = {
        pokemonName : $scope.entryForm.pokemon.name,
        nickname : $scope.entryForm.nickname,
        email : $scope.entryForm.email
      }
      console.log(value);
      Favorite.save(value, function(){
        console.log("success entry");
        $location.path(encodeURI('/guestbook/'));
      }, function(){
        console.log("error entry");
      });
    };
  }]);
})();