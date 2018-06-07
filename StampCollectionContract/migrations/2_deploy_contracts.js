var ClockAuctionBase = artifacts.require("ClockAuctionBase");
var ClockAuction = artifacts.require("ClockAuction");
var SaleClockAuction = artifacts.require("SaleClockAuction");
var Ownable = artifacts.require("Ownable");
var Pausable = artifacts.require("Pausable");
var ERC721Metadata = artifacts.require("ERC721Metadata");
var ERC721Enumerable = artifacts.require("ERC721Enumerable");
var ERC721Basic = artifacts.require("ERC721Basic");
var ERC721BasicToken = artifacts.require("ERC721BasicToken");
var ERC721 = artifacts.require("ERC721");
var ERC721TokenImplement = artifacts.require("ERC721TokenImplement");
var StampAccessControl = artifacts.require("StampAccessControl");
var StampBase = artifacts.require("StampBase");
var StampAuction = artifacts.require("StampAuction");
var StampMinting = artifacts.require("StampMinting");
var StampCollection = artifacts.require("StampCollection");

module.exports = function(deployer) {
	var ERC721Basic_, ERC721Enumerable_, ERC721Metadata_, ERC721_, ERC721BasicToken_, ERC721TokenImplement_;
	var StampAccessControl_, StampBase_, StampAuction_, StampMinting_, StampCollection_;
	var ClockAuctionBase_, Pausable_, ClockAuction_;
  deployer.deploy(StampCollection).then(function() {
   return deployer.deploy(SaleClockAuction, StampCollection.address, 5000);
  });
 //  	deployer.then(function() {
 //  		return ERC721Basic.new();
	// }).then(function(instance) {
	// 	ERC721BasicAdd = instance;
	// 	return deployer.deploy(ERC721Enumerable, ERC721Basic_.address);
	// }).then(function(instance) {
 //  		ERC721Enumerable_ = instance;
 //  		return deployer.deploy(ERC721Metadata, ERC721Basic_.address);
	// }).then(function(instance) {
 //  		ERC721Metadata_ = instance;
 //  		return deployer.deploy(ERC721, ERC721Basic_.address, ERC721Enumerable_.address, ERC721Metadata_.address);
	// }).then(function(instance) {
 //  		ERC721_ = instance;
 //  		return deployer.deploy(ERC721BasicToken, ERC721Basic_.address);
	// }).then(function(instance) {
 //  		ERC721BasicToken_ = instance;
 //  		return deployer.deploy(ERC721TokenImplement, ERC721Basic_.address, ERC721BasicToken_.address);
	// }).then(function(instance) {
 //  		ERC721TokenImplement_ = instance;
 //  		return deployer.deploy(StampAccessControl);
	// }).then(function(instance) {
 //  		StampAccessControl_ = instance;
 //  		return deployer.deploy(StampBase, StampAccessControl_.address, ERC721TokenImplement_.address);
	// }).then(function(instance) {
 //  		StampBase_ = instance;
 //  		return deployer.deploy(StampAuction, StampBase_.address);
	// }).then(function(instance) {
 //  		StampAuction_ = instance;
 //  		return deployer.deploy(StampMinting, StampAuction_.address);
	// }).then(function(instance) {
 //  		StampMinting_ = instance;
 //  		return deployer.deploy(StampCollection, StampMinting_.address);
	// }).then(function(instance) {
 //  		StampMinting_ = instance;
 //  		return deployer.deploy(ClockAuctionBase);
	// }).then(function(instance) {
 //  		ClockAuctionBase_ = instance;
 //  		return deployer.deploy(Pausable);
	// }).then(function(instance) {
 //  		Pausable_ = instance;
 //  		return deployer.deploy(ClockAuction, Pausable_.address, ClockAuctionBase_.address);
	// }).then(function(instance) {
 //  		ClockAuction_ = instance;
 //  		return deployer.deploy(SaleClockAuction, ClockAuction_.address);
	// });

};
