var StampCollection = artifacts.require("./StampCollection.sol");

contract('StampCollection', function(accounts) {
  var stamp;
  var stampCount_;
  it("发行1枚邮票", function() {
    //部署StampCollection合约
    return StampCollection.deployed().then(function(instance) {
      //获取StampCollection合约实例
      stamp = instance;
    }).then(function() {
      //取消暂停
        return stamp.unpause();
    }).then(function() {
      //获取暂停值
        return stamp.paused.call();
    }).then(function(paused) {
      //判断暂停值
       assert.equal(paused, false, "设置不暂停失败");
    }).then(function() {
      //设置CFO地址
        return stamp.setCFO(accounts[1]);
    }).then(function() {
      //获取CFO地址
        return stamp.cfoAddress.call();
    }).then(function(cfoAddress) {
      //判断CFO地址是否和accounts[1]相等
        assert.equal(cfoAddress, accounts[1], "设置CFO地址失败");
    }).then(function() {
      //发行邮票，并且创建交易
        return stamp.releaseNewStampToTransaction(1,10,10, "0x686f757069616f",1989, 1, 1, 4);
    }).then(function() {
      //获取tokenId为0的用户地址
      return stamp.ownerOf.call(0);
    }).then(function(ownerAddress) {
      console.log('stamp = '+stamp.address);
      console.log('accounts0 = '+accounts[0]);
      console.log('accounts1 = '+accounts[1]);
      console.log('accounts2 = '+accounts[2]);
      console.log('accounts3 = '+accounts[3]);
      console.log('accounts4 = '+accounts[4]);
      //判断tokenId为0的用户地址应该为stamp合约地址
      assert.equal(ownerAddress, stamp.address, "token的拥有者应该为拍卖地址");
    }).then(function() {
      //获取总共发行的邮票个数
      return stamp.stampCreatedCount.call();
    }).then(function(stampCount) {
      //保存邮票个数
      stampCount_ = stampCount;
      console.log('stampCount = '+stampCount);
      //获取发行的邮票总量
      return stamp.totalSupply.call();
    }).then(function(totalCount) {
      console.log('totalCount = '+totalCount);
      console.log('stampCount_ = '+totalCount);
      //判断邮票总量和创建的邮票个数相等
      assert.equal(totalCount.toNumber(), stampCount_, "创建的个数和totalSupply不一致");
      //获取stamp合约地址的index为0的tokenId
      return stamp.tokenOfOwnerByIndex.call(stamp.address,0);
    }).then(function(tokenId) {
      //判断tokenId应该为0
      assert.equal(tokenId, 0, "拍卖合约地址的第一个tokenId应该为1");
      //获取index为0的tokenId
      return stamp.tokenByIndex.call(0);
    }).then(function(tokenId) {
      //判断index为0的tokenId为0
      assert.equal(tokenId, 0, "tokenId和Index对应关系不对，tokenId为1index也应该为1");
      //获取symbol
      return stamp.symbol.call();
    }).then(function(symbol) {
      console.log("symbol = "+symbol.toString());
      //判断symbol是否为ST
      assert.equal(symbol.toString(), 'ST', "设置symbol失败，应该为ST");
      //获取tokenId为0的stampInfo
      return stamp.stampInfo.call(0);
    }).then(function(result) {
      //打印tokenId为0的发行年year
      console.log('year = '+result[1]);
      return stamp.exists.call(0);
    }).then(function(havaExisted) {
      //判断tokenId为0的已经存在
      assert.equal(havaExisted, true, "tokenId为0的邮票已经发布");
      return stamp.balanceOf(stamp.address);
    }).then(function(balance) {
      console.log('balance = '+balance);
      //合约地址拥有的邮票个数应该为1枚
      assert.equal(balance, 1, "合约地址拥有的邮票个数不为1");
    });
  });

  it("购买一枚邮票", function() {
    return StampCollection.deployed().then(function(instance) {
      stamp = instance;
      return stamp.transactionInfo(0);
    }).then(function(result) {
      console.log('price = ' + result[0]);
      //发行邮票，并且创建交易
        return stamp.buy(0,{from:accounts[2], value:result[0]});
    }).then(function() {
        return stamp.ownerOf.call(0);
    }).then(function(ownerAddress) {
      assert.equal(ownerAddress, accounts[2], "使用accounts[2]购买，但是tokenId为0的所有者部位accounts[2]");
      return stamp.transactionInfo(0);
    }).then(function(result) {
        console.log('price = ' + result[0]);
        console.log('seller = ' + result[1]);
        assert.equal(result[0], 0, "未删除交易");
        assert.equal(result[1], 0x0000000000000000000000000000000000000000, "未删除交易");
    });
  });

  it("用户创建一个交易，", function() {
    return StampCollection.deployed().then(function(instance) {
      stamp = instance;
      return stamp.createTransaction(0, 800000, accounts[2], {from:accounts[2]});
    }).then(function() {
        return stamp.transactionInfo(0);
    }).then(function(result) {
      assert.equal(result[0], 800000, "发布的交易价格不是800000");
      assert.equal(result[1], accounts[2], "seller不是accounts[2]");
      return stamp.buy(0,{from:accounts[3], value:result[0]});
    }).then(function() {
      return stamp.ownerOf.call(0);
    }).then(function(ownerAddress) {
      assert.equal(ownerAddress, accounts[3], "使用accounts[3]购买，但是tokenId为0的所有者部位accounts[3]");
    });
  });
});
