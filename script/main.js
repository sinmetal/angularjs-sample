(function () {
    var app = angular.module('sample', ['ngResource']).
        config(function ($routeProvider) {
            $routeProvider.
                when('/', {controller: 'TopController', templateUrl: '/html/top/top.html'}).
                when('/guestbook/', {controller: 'GuestBookListController', templateUrl: '/html/guestbook/list.html'}).
                when('/guestbook/entry', {controller: 'GuestBookEntryController', templateUrl: '/html/guestbook/entry.html'});
        });


    app.directive('navitop', function () {
        return {
            priority: 0,
            templateUrl: '/html/common/topmenu.html',
            replace: false,
            transclude: false,
            restrict: 'E',
            scope: false,
            controller: ['$scope', '$route', '$location', function ($scope, $route, $location) {
                $scope.$on('$routeChangeSuccess', function () {
                    var urlFlagments = $location.$$path.split("/");
                    $scope.currentPage = urlFlagments[1];
                });
            }]
        };
    });

    app.controller('TopController', ['$scope', '$resource', function ($scope, $resource) {
    }]);

    app.controller('GuestBookListController', ['$scope', '$resource', function ($scope, $resource) {
        var Favorite = $resource("/favorite");
        $scope.favos = Favorite.query(function () {
            console.log("success favorite query");
            console.log($scope.favos);
        }, function () {
            console.log("error favorite query");
        });
    }]);

    app.controller('GuestBookEntryController', ['$scope', '$location', '$resource', function ($scope, $location, $resource) {
        $scope.elementTypes = [
            {"id": "1", "name": "くさ"},
            {"id": "2", "name": "ほのお"},
            {"id": "3", "name": "みず"}
        ];
        var Pokemon = $resource("/pokemon");
        var Favorite = $resource("/favorite");

        $scope.changeElementType = function () {
            console.log($scope.entryForm.elementTypeId);
            $scope.pokemons = Pokemon.query({id: $scope.entryForm.elementTypeId}, function () {
                console.log("success list");
            }, function () {
                console.log("error list");
            });

        }

        $scope.submit = function ($event) {
            console.log($scope.entryForm);
            var value = {
                pokemonName: $scope.entryForm.pokemon.name,
                nickname: $scope.entryForm.nickname,
                email: $scope.entryForm.email
            }
            console.log(value);
            Favorite.save(value, function () {
                console.log("success entry");
                $location.path(encodeURI('/guestbook/'));
            }, function () {
                console.log("error entry");
            });
        };
    }]);
})();