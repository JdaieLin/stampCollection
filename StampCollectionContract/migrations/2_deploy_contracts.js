
var StampCollection = artifacts.require("StampCollection");

module.exports = function(deployer) {
  deployer.deploy(StampCollection).then(function() {
  });
};
