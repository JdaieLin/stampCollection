var StampCollection = artifacts.require("./StampCollection.sol");
var SaleClockAuction = artifacts.require("./SaleClockAuction.sol");

contract('StampCollection', function(accounts) {
  var stamp;
  var saleauction;
  console.log('--stamp--');
  it("创建10枚邮票", function() {
    return StampCollection.deployed().then(function(instance) {
      stamp = instance;
      console.log('stamp:'+stamp);
      return SaleClockAuction.deployed();
    }).then(function(instance) {
        saleauction = instance;
        console.log('saleauction:'+saleauction);
        return stamp.setSaleAuctionAddress(saleauction.address);
    }).then(function() {
        return stamp.unpause();
    }).then(function() {
        return stamp.setCFO(accounts[1]);
    }).then(function() {
        return stamp.createNewStampAuction(1,10, 10, "0x686f757069616f",1989, 1, 1,  [0,1,2,3,4,0,1,2,3,4,0,1,2,3,4,0,1,2,3,4,0,1,2,3,4,0]);
    }).then(function() {
      return stamp.ownerOf.call(0);
    }).then(function(ownerAddress) {
      console.log('stampCreatedCount = '+saleauction.address);
      console.log('stampCreatedCount = '+stamp.address);
      console.log('accounts0 = '+accounts[0]);
      console.log('accounts1 = '+accounts[1]);
      console.log('accounts2 = '+accounts[2]);
      console.log('accounts3 = '+accounts[3]);
      console.log('accounts4 = '+accounts[4]);
        assert.equal(ownerAddress, stamp.address, "Amount wasn't correctly taken from the sender");
    }).then(function() {
      return stamp.stampCreatedCount.call();
    }).then(function(stampCount) {
      console.log('stampCount = '+stampCount);
    });
  });
});
