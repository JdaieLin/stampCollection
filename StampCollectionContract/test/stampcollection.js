var StampCollection = artifacts.require("./StampCollection.sol");
var SaleClockAuction = artifacts.require("./SaleClockAuction.sol");

contract('StampCollection', function(accounts) {
  var stamp;
  var saleauction;
  var stampCount_;
  console.log('--stamp--');
  it("创建5枚邮票", function() {
    return StampCollection.deployed().then(function(instance) {
      stamp = instance;
      return SaleClockAuction.deployed();
    }).then(function(instance) {
        saleauction = instance;
        return stamp.setSaleAuctionAddress(saleauction.address);
    }).then(function() {
        return stamp.saleAuction.call();
    }).then(function(saleAutionAddress) {
        assert.equal(saleAutionAddress, saleauction.address, "设置拍卖地址成功");
    }).then(function() {
        return stamp.unpause();
    }).then(function() {
        return stamp.paused.call();
    }).then(function(paused) {
       assert.equal(paused, false, "设置不暂停成功");
    }).then(function() {
        return stamp.setCFO(accounts[1]);
    }).then(function() {
        return stamp.cfoAddress.call();
    }).then(function(cfoAddress) {
        assert.equal(cfoAddress, accounts[1], "设置CFO地址成功");
    }).then(function() {
        return stamp.releaseNewStampToAuction(1,10, 10, "0x686f757069616f",1989, 1, 1,  [0,1,2,3,4]);
    }).then(function() {
      return stamp.ownerOf.call(0);
    }).then(function(ownerAddress) {
      console.log('saleauction = '+saleauction.address);
      console.log('stamp = '+stamp.address);
      console.log('accounts0 = '+accounts[0]);
      console.log('accounts1 = '+accounts[1]);
      console.log('accounts2 = '+accounts[2]);
      console.log('accounts3 = '+accounts[3]);
      console.log('accounts4 = '+accounts[4]);
      assert.equal(ownerAddress, saleauction.address, "token的拥有者应该为拍卖地址");
    }).then(function() {
      return stamp.stampCreatedCount.call();
    }).then(function(stampCount) {
      stampCount_ = stampCount;
      console.log('stampCount = '+stampCount);
      return stamp.totalSupply.call();
    }).then(function(totalCount) {
      console.log('totalCount = '+totalCount);
      console.log('stampCount_ = '+totalCount);
      assert.equal(totalCount.toNumber(), stampCount_, "创建的个数和totalSupply不一致");
      return stamp.tokenOfOwnerByIndex.call(saleauction.address,0);
    }).then(function(tokenId) {
      assert.equal(tokenId, 0, "拍卖合约地址的第一个tokenId应该为0");
      return stamp.tokenByIndex.call(1);
    }).then(function(tokenId) {
      assert.equal(tokenId, 1, "tokenId和Index对应关系不对，tokenId为1index也应该为1");
      return stamp.symbol.call();
    }).then(function(symbol) {
      console.log("symbol = "+symbol.toString());
      assert.equal(symbol.toString(), 'ST', "设置symbol失败，应该为ST");
      return stamp.stampInfo.call(0);
    }).then(function(result) {
      console.log('year = '+result[1]);
      return stamp.STAMP_STARTING_PRICE.call();
    }).then(function(startPrice) {
      console.log('startPrice = '+startPrice);
      assert.equal(startPrice, 10000000000000000, "初始价格设置不正确");
      return stamp.STAMP_AUCTION_DURATION.call();
    }).then(function(duration) {
      console.log('duration = '+duration);
      assert.equal(duration, 604800, "初始时间间隔设置不正确");
      return stamp.isApprovedForAll.call(accounts[0],saleauction.address);
    }).then(function(isApproved) {
      console.log('isApproved = '+isApproved);
      assert.equal(isApproved, true, "ceo创建的邮票已经授权给拍卖合约进行拍卖");
      return stamp.exists.call(0);
    }).then(function(havaExisted) {
      assert.equal(havaExisted, true, "tokenId为0的邮票已经发布");
      // return stamp.exists.call(0);
    });
  });
});
