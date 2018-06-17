var StampCollection = artifacts.require("./StampCollection.sol");

contract('StampCollection', function(accounts) {
  var stamp;

  it("发布合约，验证发布的一些初始化信息", function() {
    return StampCollection.deployed().then(function(instance) {
      stamp = instance;
      return stamp.STAMP_SET_RELEASE_LIMIT.call();
    }).then(function(setLimit) {
      assert.equal(setLimit, '3000', "套票组个数限制为3000");
      return stamp.STAMP_TYPE_RELEASE_LIMIT.call();
    }).then(function(typeLimit) {
      assert.equal(typeLimit, '30000', "邮票类型限制为30000");
      return stamp.paused.call();
    }).then(function(paused) {
      assert.equal(paused, false, "初始状态为未暂停");
      return stamp.ceoAddress.call();
    }).then(function(ceoAddress) {
      assert.equal(ceoAddress, accounts[0], "初始状态为未暂停");
      return stamp.cooAddress.call();
    }).then(function(cooAddress) {
      assert.equal(cooAddress, accounts[0], "初始状态为未暂停");
      return stamp.cfoAddress.call();
    }).then(function(cfoAddress) {
      assert.equal(cfoAddress, 0x0000000000000000000000000000000000000000, "初始状态为未暂停");
    });
  });

//发布一组邮票
  it("发布一组邮票，发布这组有三种邮票，发行三枚邮票并且可以交易", function() {
    return stamp.releaseNewStampSet(1, [11,22,33]).then(function(instance) {
      //发布一种邮票信息
        return stamp.releaseNewStampInfo(1, 11, 1000, 1989, 1, "0x666f757069616f");
    }).then(function() {
      //发布一种邮票信息
        return stamp.releaseNewStampInfo(1, 22, 1000, 1989, 2, "0x666f757069616f");
    }).then(function() {
      //发布一种邮票信息
        return stamp.releaseNewStampInfo(1, 33, 1000, 1989, 3, "0x666f757069616f");
    }).then(function() {
      //发行一枚邮票信息，并创建交易
      return stamp.releaseNewStampToTransaction(11, 2);
    }).then(function() {
      return stamp.releaseNewStampToTransaction(22, 3);
    }).then(function() {
      return stamp.releaseNewStampToTransaction(33, 4);
    }).then(function() {
      //获取邮票tokenId为2的邮票拥有者地址
      return stamp.ownerOf.call(2);
    }).then(function(ownerAddress) {
      //判断拥有着是不是合约地址
      assert.equal(ownerAddress, stamp.address, "tokenId为2的邮票应该归合约所有");
    });
  });


  //此处的测试数据依据上面的测试数据
  it("ERC721验证信息功能测试", function() {
    //获取项目名称
    return stamp.name.call().then(function(name) {
      assert.equal(name, "StampToken", "name应该为StampToken");
      //获取ID标识
      return stamp.symbol.call();
    }).then(function(symbol) {
      assert.equal(symbol, "ST", "symbol应该为ST");
      //获取token的总量
      return stamp.totalSupply.call();
    }).then(function(totalSupply) {
      assert.equal(totalSupply, 3, "totalSupply应该为3");
      //获取stamp合约地址index为2的tokenId
      return stamp.tokenOfOwnerByIndex.call(stamp.address, 2);
    }).then(function(tokenId) {
      assert.equal(tokenId, 2, "index为2的tokenId为2");
      //获取index为1的tokenId
      return stamp.tokenByIndex.call(1);
    }).then(function(tokenId) {
      assert.equal(tokenId, 1, "index为1的tokenId为1");
      //获取tokenId为0的用户地址
      return stamp.ownerOf.call(0);
    }).then(function(ownerAddress) {
      //判断拥有着是不是合约地址
      assert.equal(ownerAddress, stamp.address, "tokenId为0的邮票应该归合约所有");
      //判断tokenId为0是否存在
      return stamp.exists.call(0);
    }).then(function(havaExisted) {
      assert.equal(havaExisted, true, "tokenId为0的邮票已经发布");
      return stamp.balanceOf(stamp.address);
    }).then(function(balance) {
      assert.equal(balance, 3, "合约地址拥有的邮票个数不为3");
    });
  });

  it("自己实现的相关功能测试", function() {
    return stamp.isExistSet.call(1).then(function(existed) {
      assert.equal(existed, true, "setId = 1已经发布，但是isExistSet返回false");
      return stamp.isExistStampType.call(11);
    }).then(function(existed) {
      assert.equal(existed, true, "typeId = 11已经发布，但是isExistStampType返回false");
      return stamp.isExistStampType.call(22);
    }).then(function(existed) {
      assert.equal(existed, true, "typeId = 22已经发布，但是isExistStampType返回false");
      return stamp.isExistStampType.call(33);
    }).then(function(existed) {
      assert.equal(existed, true, "typeId = 33已经发布，但是isExistStampType返回false");
      return stamp.totalSupply.call();
    }).then(function(totalCount) {
      //判断邮票总量和创建的邮票个数相等
      assert.equal(totalCount.toNumber(), 3, "创建的个数和totalSupply不一致");
      //获取stamp合约地址的index为0的tokenId
      return stamp.tokenOfOwnerByIndex.call(stamp.address,0);
    }).then(function(tokenId) {
      //判断tokenId应该为0
      assert.equal(tokenId, 0, "拍卖合约地址的第一个tokenId应该为0");
      //获取index为0的tokenId
      return stamp.tokenByIndex.call(0);
    }).then(function(tokenId) {
      //判断index为0的tokenId为0
      assert.equal(tokenId, 0, "tokenId和Index对应关系不对，tokenId为1index也应该为0");
      //获取tokenId为0的stampInfo
      return stamp.getStamp.call(0);
    }).then(function(result) {
      //打印tokenId为0的发行年year
      console.log('appearance = '+result[1]);
      return stamp.exists.call(0);
    }).then(function(havaExisted) {
      //判断tokenId为0的已经存在
      assert.equal(havaExisted, true, "tokenId为0的邮票已经发布");
      return stamp.balanceOf(stamp.address);
    }).then(function(balance) {
      console.log('balance = '+balance);
      //合约地址拥有的邮票个数应该为1枚
      assert.equal(balance, 3, "合约地址拥有的邮票个数不为3");
    });
  });


    it("ERC721授权功能验证", function() {
    return stamp.transactionInfo(0).then(function(result) {
      console.log('tokenId 0 price = ' + result[0]);
      return stamp.buy(0,{from:accounts[1], value:result[0]});
    }).then(function() {
        return stamp.ownerOf.call(0);
    }).then(function(ownerAddress) {
      assert.equal(ownerAddress, accounts[1], "使用accounts[3]购买，但是tokenId为0的所有者部位accounts[3]");
      return stamp.setApprovalForAll(accounts[2],true,{from:accounts[1]})
    }).then(function(result) {
      //发行邮票，并且创建交易
      return stamp.isApprovedForAll(accounts[1],accounts[2]);
    }).then(function(isApproved) {
        assert.equal(isApproved, true, "合约地址授权给accounts[2]失败");
        return stamp.approve(accounts[2], 0, {from:accounts[1]});
    }).then(function() {
      return stamp.getApproved(0);
    }).then(function(approvedAddress) {
      assert.equal(approvedAddress, accounts[2], "使用accounts[2]购买，但是tokenId为0的所有者部位accounts[2]");
    });
  });


  it("购买一枚邮票", function() {
    return StampCollection.deployed().then(function(instance) {
      stamp = instance;
      return stamp.transactionInfo(1);
    }).then(function(result) {
      console.log('price = ' + result[0]);
      //发行邮票，并且创建交易
        return stamp.buy(1,{from:accounts[3], value:result[0]});
    }).then(function() {
        return stamp.ownerOf.call(1);
    }).then(function(ownerAddress) {
      assert.equal(ownerAddress, accounts[3], "使用accounts[3]购买，但是tokenId为0的所有者部位accounts[3]");
      return stamp.transactionInfo(1);
    }).then(function(result) {
      console.log('result = ' + JSON.stringify(result));
      var price = result[0].toNumber();
        console.log('price = ' + result[0]);
        console.log('seller = ' + result[1]);
        assert.equal(result[0], 0, "未删除交易");
        assert.equal(result[1], 0x0000000000000000000000000000000000000000, "未删除交易");
    });
  });

//此处的测试数据依据上面的测试数据, accounts[2]拥有tokenId为0的邮票
  it("用户创建一个交易，", function() {
    return StampCollection.deployed().then(function(instance) {
      stamp = instance;
      return stamp.createTransaction(1, 80000, accounts[3], {from:accounts[3]});
    }).then(function() {                
        return stamp.transactionInfo(1);
    }).then(function(result) {
      assert.equal(result[0], 80000, "发布的交易价格不是800000");
      assert.equal(result[1], accounts[3], "seller不是accounts[2]");
      return stamp.buy(1,{from:accounts[4], value:result[0]});
    }).then(function() {
      return stamp.ownerOf.call(1);
    }).then(function(ownerAddress) {
      assert.equal(ownerAddress, accounts[4], "使用accounts[3]购买，但是tokenId为0的所有者部位accounts[3]");
    });
  });

//此处的测试数据依据上面的测试数据
  it("用户发起使用邮票换元宝交易--系统回购", function() {
    return stamp.repoIngotsTransaction(1, 123456, {from:accounts[4]}).then(function() {
      return stamp.repoIngotsInfo(1);
    }).then(function(result) {
      assert.equal(result[1], 123456, "换购的元宝数量不对");
      assert.equal(result[2], accounts[4], "换购的卖方地址不对");
    });
  });

  it("验证CFO提取以太币", function() {
    return stamp.setCFO(accounts[5],{from:accounts[0]}).then(function() {
      return stamp.cfoAddress.call();
    }).then(function(cfoAddress) {
      assert.equal(cfoAddress, accounts[5], "初始状态为未暂停");
      return stamp.withdrawBalance({from:accounts[5]});
    }).then(function(result) {
      return web3.fromWei(web3.eth.getBalance(accounts[5]));
    }).then(function(balance) {
      console.log("cfoAddress-balance = "+balance);
    });
  });
});
