pragma solidity ^0.4.23;

import "./ERC721/ERC721.sol";
import "./ERC721/ERC721Basic.sol";
import "./ERC721/ERC721Receiver.sol";
import "./math/SafeMath.sol";
import "./AddressUtils.sol";

contract StampAccessControl {
    event ContractUpgrade(address newContract);

    address public ceoAddress;
    address public cfoAddress;
    address public cooAddress;

    bool public paused = false;

    modifier onlyCEO() {
        require(msg.sender == ceoAddress);
        _;
    }

    modifier onlyCFO() {
        require(msg.sender == cfoAddress);
        _;
    }

    modifier onlyCOO() {
        require(msg.sender == cooAddress);
        _;
    }

    modifier onlyCLevel() {
        require(
            msg.sender == cooAddress ||
            msg.sender == ceoAddress ||
            msg.sender == cfoAddress
        );
        _;
    }

    function setCEO(address _newCEO) external onlyCEO {
        require(_newCEO != address(0));

        ceoAddress = _newCEO;
    }

    function setCFO(address _newCFO) external onlyCEO {
        require(_newCFO != address(0));

        cfoAddress = _newCFO;
    }

    function setCOO(address _newCOO) external onlyCEO {
        require(_newCOO != address(0));

        cooAddress = _newCOO;
    }

    modifier whenNotPaused() {
        require(!paused);
        _;
    }

    modifier whenPaused {
        require(paused);
        _;
    }

    function pause() external onlyCLevel whenNotPaused {
        paused = true;
    }

    function unpause() public onlyCEO whenPaused {
        paused = false;
    }
}

contract StampDataSource {
    /*** DATA TYPES ***/
    struct StampInfo {
        uint32 setId;//set ID
        uint32 typeId;//Unique stamp ID
        uint32 totalAmount;//The total amount of stamps issued
        uint32 remainingAmount;//The remainder of the stamp issue
        uint16 year;//issuing year of Stamp 
        uint16 idOfSet;//ID in a set of stamps
        bytes32 name;//the name of Stamp 
    }
    //Information about a stamp
    struct Stamp {
        uint32 typeId;
        uint256 price;
        uint8 appearance;
    }
    //Information on the stamp section
    struct StampSet {
        uint16 setId;
        uint32[] typeIds;
    }

    /*** STORAGE ***/
    StampInfo[] stampInfos;
    mapping (uint32 => StampInfo) public TypeIdToStampInfo;//typeId -> StampInfo
    
    Stamp[] stamps;
    
    StampSet[] stampSets;
    mapping (uint16 => StampSet) public SetIdToStampSet;//
    mapping (uint32 => uint16) public TypeIdToSetId;//
    mapping (uint16 => bool) public isExistSet;
    mapping (uint32 => bool) public isExistStampType;
}



contract ERC721BasicToken is ERC721Basic, StampAccessControl {
  using SafeMath for uint256;
  using AddressUtils for address;

  bytes4 constant ERC721_RECEIVED = 0xf0b9e5ba;

  // Mapping from token ID to owner
  mapping (uint256 => address) internal tokenOwner;

  // Mapping from token ID to approved address
  mapping (uint256 => address) internal tokenApprovals;

  // Mapping from owner to number of owned token
  mapping (address => uint256) internal ownedTokensCount;

  // Mapping from owner to operator approvals
  mapping (address => mapping (address => bool)) internal operatorApprovals;

  /**
   * @dev Guarantees msg.sender is owner of the given token
   * @param _tokenId uint256 ID of the token to validate its ownership belongs to msg.sender
   */
  modifier onlyOwnerOf(uint256 _tokenId) {
    require(ownerOf(_tokenId) == msg.sender);
    _;
  }

  /**
   * @dev Checks msg.sender can transfer a token, by being owner, approved, or operator
   * @param _tokenId uint256 ID of the token to validate
   */
  modifier canTransfer(uint256 _tokenId) {
    require(isApprovedOrOwner(msg.sender, _tokenId));
    _;
  }

  /**
   * @dev Gets the balance of the specified address
   * @param _owner address to query the balance of
   * @return uint256 representing the amount owned by the passed address
   */
  function balanceOf(address _owner) public view returns (uint256) {
    require(_owner != address(0));
    return ownedTokensCount[_owner];
  }

  /**
   * @dev Gets the owner of the specified token ID
   * @param _tokenId uint256 ID of the token to query the owner of
   * @return owner address currently marked as the owner of the given token ID
   */
  function ownerOf(uint256 _tokenId) public view returns (address) {
    address owner = tokenOwner[_tokenId];
    require(owner != address(0));
    return owner;
  }

  /**
   * @dev Returns whether the specified token exists
   * @param _tokenId uint256 ID of the token to query the existance of
   * @return whether the token exists
   */
  function exists(uint256 _tokenId) public view returns (bool) {
    address owner = tokenOwner[_tokenId];
    return owner != address(0);
  }

  /**
   * @dev Approves another address to transfer the given token ID
   * @dev The zero address indicates there is no approved address.
   * @dev There can only be one approved address per token at a given time.
   * @dev Can only be called by the token owner or an approved operator.
   * @param _to address to be approved for the given token ID
   * @param _tokenId uint256 ID of the token to be approved
   */
  function approve(address _to, uint256 _tokenId) public whenNotPaused{
    address owner = ownerOf(_tokenId);
    require(_to != owner);
    require(msg.sender == owner || isApprovedForAll(owner, msg.sender));

    if (getApproved(_tokenId) != address(0) || _to != address(0)) {
      tokenApprovals[_tokenId] = _to;
      emit Approval(owner, _to, _tokenId);
    }
  }

  /**
   * @dev Gets the approved address for a token ID, or zero if no address set
   * @param _tokenId uint256 ID of the token to query the approval of
   * @return address currently approved for a the given token ID
   */
  function getApproved(uint256 _tokenId) public view returns (address) {
    return tokenApprovals[_tokenId];
  }

  /**
   * @dev Sets or unsets the approval of a given operator
   * @dev An operator is allowed to transfer all tokens of the sender on their behalf
   * @param _to operator address to set the approval
   * @param _approved representing the status of the approval to be set
   */
  function setApprovalForAll(address _to, bool _approved) public whenNotPaused{
    require(_to != msg.sender);
    operatorApprovals[msg.sender][_to] = _approved;
    emit ApprovalForAll(msg.sender, _to, _approved);
  }

  /**
   * @dev Tells whether an operator is approved by a given owner
   * @param _owner owner address which you want to query the approval of
   * @param _operator operator address which you want to query the approval of
   * @return bool whether the given operator is approved by the given owner
   */
  function isApprovedForAll(address _owner, address _operator) public view returns (bool approved) {
   approved = operatorApprovals[_owner][_operator];
  }

  /**
   * @dev Transfers the ownership of a given token ID to another address
   * @dev Usage of this method is discouraged, use `safeTransferFrom` whenever possible
   * @dev Requires the msg sender to be the owner, approved, or operator
   * @param _from current owner of the token
   * @param _to address to receive the ownership of the given token ID
   * @param _tokenId uint256 ID of the token to be transferred
  */
  function transferFrom(address _from, address _to, uint256 _tokenId) public whenNotPaused canTransfer(_tokenId){
    require(_from != address(0));
    require(_to != address(0));

    clearApproval(_from, _tokenId);
    removeTokenFrom(_from, _tokenId);
    addTokenTo(_to, _tokenId);

    emit Transfer(_from, _to, _tokenId);
  }

  /**
   * @dev Safely transfers the ownership of a given token ID to another address
   * @dev If the target address is a contract, it must implement `onERC721Received`,
   *  which is called upon a safe transfer, and return the magic value
   *  `bytes4(keccak256("onERC721Received(address,uint256,bytes)"))`; otherwise,
   *  the transfer is reverted.
   * @dev Requires the msg sender to be the owner, approved, or operator
   * @param _from current owner of the token
   * @param _to address to receive the ownership of the given token ID
   * @param _tokenId uint256 ID of the token to be transferred
  */
  function safeTransferFrom(
    address _from,
    address _to,
    uint256 _tokenId
  )
    public
    whenNotPaused
    canTransfer(_tokenId)
  {
    // solium-disable-next-line arg-overflow
    safeTransferFrom(_from, _to, _tokenId, "");
  }

  /**
   * @dev Safely transfers the ownership of a given token ID to another address
   * @dev If the target address is a contract, it must implement `onERC721Received`,
   *  which is called upon a safe transfer, and return the magic value
   *  `bytes4(keccak256("onERC721Received(address,uint256,bytes)"))`; otherwise,
   *  the transfer is reverted.
   * @dev Requires the msg sender to be the owner, approved, or operator
   * @param _from current owner of the token
   * @param _to address to receive the ownership of the given token ID
   * @param _tokenId uint256 ID of the token to be transferred
   * @param _data bytes data to send along with a safe transfer check
   */
  function safeTransferFrom(
    address _from,
    address _to,
    uint256 _tokenId,
    bytes _data
  )
    public
    whenNotPaused
    canTransfer(_tokenId)
  {
    transferFrom(_from, _to, _tokenId);
    // solium-disable-next-line arg-overflow
    require(checkAndCallSafeTransfer(_from, _to, _tokenId, _data));
  }

  /**
   * @dev Returns whether the given spender can transfer a given token ID
   * @param _spender address of the spender to query
   * @param _tokenId uint256 ID of the token to be transferred
   * @return bool whether the msg.sender is approved for the given token ID,
   *  is an operator of the owner, or is the owner of the token
   */
  function isApprovedOrOwner(address _spender, uint256 _tokenId) internal view returns (bool) {
    address owner = ownerOf(_tokenId);
    return _spender == owner || getApproved(_tokenId) == _spender || isApprovedForAll(owner, _spender);
  }

  /**
   * @dev Internal function to mint a new token
   * @dev Reverts if the given token ID already exists
   * @param _to The address that will own the minted token
   * @param _tokenId uint256 ID of the token to be minted by the msg.sender
   */
  function _mint(address _to, uint256 _tokenId) internal {
    require(_to != address(0));
    addTokenTo(_to, _tokenId);
    emit Transfer(address(0), _to, _tokenId);
  }

  /**
   * @dev Internal function to burn a specific token
   * @dev Reverts if the token does not exist
   * @param _tokenId uint256 ID of the token being burned by the msg.sender
   */
  function _burn(address _owner, uint256 _tokenId) internal {
    clearApproval(_owner, _tokenId);
    removeTokenFrom(_owner, _tokenId);
    emit Transfer(_owner, address(0), _tokenId);
  }

  /**
   * @dev Internal function to clear current approval of a given token ID
   * @dev Reverts if the given address is not indeed the owner of the token
   * @param _owner owner of the token
   * @param _tokenId uint256 ID of the token to be transferred
   */
  function clearApproval(address _owner, uint256 _tokenId) internal {
    require(ownerOf(_tokenId) == _owner);
    if (tokenApprovals[_tokenId] != address(0)) {
      tokenApprovals[_tokenId] = address(0);
      emit Approval(_owner, address(0), _tokenId);
    }
  }

  /**
   * @dev Internal function to add a token ID to the list of a given address
   * @param _to address representing the new owner of the given token ID
   * @param _tokenId uint256 ID of the token to be added to the tokens list of the given address
   */
  function addTokenTo(address _to, uint256 _tokenId) internal {
    require(tokenOwner[_tokenId] == address(0));
    tokenOwner[_tokenId] = _to;
    ownedTokensCount[_to] = ownedTokensCount[_to].add(1);
  }

  /**
   * @dev Internal function to remove a token ID from the list of a given address
   * @param _from address representing the previous owner of the given token ID
   * @param _tokenId uint256 ID of the token to be removed from the tokens list of the given address
   */
  function removeTokenFrom(address _from, uint256 _tokenId) internal {
    require(ownerOf(_tokenId) == _from);
    ownedTokensCount[_from] = ownedTokensCount[_from].sub(1);
    tokenOwner[_tokenId] = address(0);
  }

  /**
   * @dev Internal function to invoke `onERC721Received` on a target address
   * @dev The call is not executed if the target address is not a contract
   * @param _from address representing the previous owner of the given token ID
   * @param _to target address that will receive the tokens
   * @param _tokenId uint256 ID of the token to be transferred
   * @param _data bytes optional data to send along with the call
   * @return whether the call correctly returned the expected magic value
   */
  function checkAndCallSafeTransfer(
    address _from,
    address _to,
    uint256 _tokenId,
    bytes _data
  )
    internal
    returns (bool)
  {
    if (!_to.isContract()) {
      return true;
    }
    bytes4 retval = ERC721Receiver(_to).onERC721Received(_from, _tokenId, _data);
    return (retval == ERC721_RECEIVED);
  }
}

contract ERC721TokenImplement is StampDataSource, ERC721, ERC721BasicToken {
  
    // Token name
    string internal name_ = "StampToken";

    // Token symbol
    string internal symbol_ = "ST";

  // Mapping from owner to list of owned token IDs
  mapping (address => uint256[]) internal ownedTokens;

  // Mapping from token ID to index of the owner tokens list
  mapping(uint256 => uint256) internal ownedTokensIndex;

  // Array with all token ids, used for enumeration
  uint256[] internal allTokens;

  // Mapping from token id to position in the allTokens array
  mapping(uint256 => uint256) internal allTokensIndex;

  // Optional mapping for token URIs
  mapping(uint256 => string) internal tokenURIs;

  /**
   * @dev Gets the token name
   * @return string representing the token name
   */
  function name() public view returns (string) {
    return name_;
  }

  /**
   * @dev Gets the token symbol
   * @return string representing the token symbol
   */
  function symbol() public view returns (string) {
    return symbol_;
  }

  /**
   * @dev Returns an URI for a given token ID
   * @dev Throws if the token ID does not exist. May return an empty string.
   * @param _tokenId uint256 ID of the token to query
   */
  function tokenURI(uint256 _tokenId) public view returns (string) {
    require(exists(_tokenId));
    return tokenURIs[_tokenId];
  }

  /**
   * @dev Gets the total amount of tokens stored by the contract
   * @return uint256 representing the total amount of tokens
   */
  function totalSupply() public view returns (uint256) {
    return allTokens.length;
  }
  
  /**
   * @dev Gets the token ID at a given index of the tokens list of the requested owner
   * @param _owner address owning the tokens list to be accessed
   * @param _index uint256 representing the index to be accessed of the requested tokens list
   * @return uint256 token ID at the given index of the tokens list owned by the requested address
   */
  function tokenOfOwnerByIndex(address _owner, uint256 _index) public view returns (uint256) {
    require(_index < balanceOf(_owner));
    return ownedTokens[_owner][_index];
  }

  

  /**
   * @dev Gets the token ID at a given index of all the tokens in this contract
   * @dev Reverts if the index is greater or equal to the total number of tokens
   * @param _index uint256 representing the index to be accessed of the tokens list
   * @return uint256 token ID at the given index of the tokens list
   */
  function tokenByIndex(uint256 _index) public view returns (uint256) {
    require(_index < totalSupply());
    return allTokens[_index];
  }

  /**
   * @dev Internal function to set the token URI for a given token
   * @dev Reverts if the token ID does not exist
   * @param _tokenId uint256 ID of the token to set its URI
   * @param _uri string URI to assign
   */
  function _setTokenURI(uint256 _tokenId, string _uri) internal {
    require(exists(_tokenId));
    tokenURIs[_tokenId] = _uri;
  }

  /**
   * @dev Internal function to add a token ID to the list of a given address
   * @param _to address representing the new owner of the given token ID
   * @param _tokenId uint256 ID of the token to be added to the tokens list of the given address
   */
  function addTokenTo(address _to, uint256 _tokenId) internal {
    super.addTokenTo(_to, _tokenId);
    uint256 length = ownedTokens[_to].length;
    ownedTokens[_to].push(_tokenId);
    ownedTokensIndex[_tokenId] = length;
  }

  /**
   * @dev Internal function to remove a token ID from the list of a given address
   * @param _from address representing the previous owner of the given token ID
   * @param _tokenId uint256 ID of the token to be removed from the tokens list of the given address
   */
  function removeTokenFrom(address _from, uint256 _tokenId) internal {
    super.removeTokenFrom(_from, _tokenId);

    uint256 tokenIndex = ownedTokensIndex[_tokenId];
    uint256 lastTokenIndex = ownedTokens[_from].length.sub(1);
    uint256 lastToken = ownedTokens[_from][lastTokenIndex];

    ownedTokens[_from][tokenIndex] = lastToken;
    ownedTokens[_from][lastTokenIndex] = 0;
    // Note that this will handle single-element arrays. In that case, both tokenIndex and lastTokenIndex are going to
    // be zero. Then we can make sure that we will remove _tokenId from the ownedTokens list since we are first swapping
    // the lastToken to the first position, and then dropping the element placed in the last position of the list

    ownedTokens[_from].length--;
    ownedTokensIndex[_tokenId] = 0;
    ownedTokensIndex[lastToken] = tokenIndex;
  }

  /**
   * @dev Internal function to mint a new token
   * @dev Reverts if the given token ID already exists
   * @param _to address the beneficiary that will own the minted token
   * @param _tokenId uint256 ID of the token to be minted by the msg.sender
   */
  function _mint(address _to, uint256 _tokenId) internal {
    super._mint(_to, _tokenId);

    allTokensIndex[_tokenId] = allTokens.length;
    allTokens.push(_tokenId);
  }

  /**
   * @dev Internal function to burn a specific token
   * @dev Reverts if the token does not exist
   * @param _owner owner of the token to burn
   * @param _tokenId uint256 ID of the token being burned by the msg.sender
   */
  function _burn(address _owner, uint256 _tokenId) internal {
    super._burn(_owner, _tokenId);

    // Clear metadata (if any)
    if (bytes(tokenURIs[_tokenId]).length != 0) {
      delete tokenURIs[_tokenId];
    }

    // Reorg all tokens array
    uint256 tokenIndex = allTokensIndex[_tokenId];
    uint256 lastTokenIndex = allTokens.length.sub(1);
    uint256 lastToken = allTokens[lastTokenIndex];

    allTokens[tokenIndex] = lastToken;
    allTokens[lastTokenIndex] = 0;

    allTokens.length--;
    allTokensIndex[_tokenId] = 0;
    allTokensIndex[lastToken] = tokenIndex;
  }
}

contract StampBase is ERC721TokenImplement{
    event CreateNewStamp(uint32 typeId, uint8 appearance);

    function _createNewStamp(
        uint32 _typeId,
        uint256 _price,
        uint8 _appearance
    )
        internal
        returns (uint newTokenId)
    {
        require(_typeId == uint256(uint32(_typeId)));

        Stamp memory _stamp = Stamp({
            typeId:_typeId,
            price:_price,
            appearance:_appearance
        });
        newTokenId = stamps.push(_stamp) - 1;

        require(newTokenId == uint256(uint32(newTokenId)));

        emit CreateNewStamp(_typeId, _appearance);
    }
    
    function _owns(address _claimant, uint256 _tokenId) internal view returns (bool) {
        return tokenOwner[_tokenId] == _claimant;
    }
}

contract StampTransaction is StampBase {
     
    struct Trasaction {
        uint256 startedAt;
        uint128 price;
        address seller;
    }
    
    mapping (uint256 => Trasaction) tokenIdToTransaction;
    
    event TransactionCreated(uint256 tokenId, uint256 price);
    event TransactionSuccessful(uint256 tokenId, uint256 price, address buyer);
    event TransactionCancelled(uint256 tokenId);
    event TransactionTimeOut(uint256 tokenId, uint256 startedAt, address seller);
    
    function _addTransaction(uint256 _tokenId, Trasaction _tracsaction) internal {
        tokenIdToTransaction[_tokenId] = _tracsaction;

        emit TransactionCreated(
            uint256(_tokenId),
            uint256(_tracsaction.price)
        );
    }

    function _removeTransaction(uint256 _tokenId) internal {
        delete tokenIdToTransaction[_tokenId];
    }
    
    function cancelTransaction(uint256 _tokenId, address _seller) internal {
        _removeTransaction(_tokenId);
        transferFrom(address(this), _seller, _tokenId);
        emit TransactionCancelled(_tokenId);
    }

    function _bid(uint256 _tokenId, uint256 _buyPrice) internal
    {
        Trasaction storage transactionn = tokenIdToTransaction[_tokenId];

        address seller = transactionn.seller;
        uint128 price = transactionn.price;
        require(_buyPrice == price);
        
        _removeTransaction(_tokenId);
        seller.transfer(_computePrice(_buyPrice));

        emit TransactionSuccessful(_tokenId, _buyPrice, msg.sender);
    }
    
    function _computePrice(uint256 _price) internal pure returns (uint256) {
        uint256 price = _price * 9900 / 10000;
        if (price > 500000000000000000) {
            price = 500000000000000000;
        }
        return price;
    }
    
    function createTransaction(uint256 _tokenId, uint256 _price, address _seller, uint256 _startedAt) public whenNotPaused
    {
        require(_price == uint256(uint128(_price)));
        require(ownerOf(_tokenId) == msg.sender);

        transferFrom(msg.sender, address(this), _tokenId);
        Trasaction memory tracsaction = Trasaction(
            _startedAt,
            uint128(_price),
            _seller
        );
        _addTransaction(_tokenId, tracsaction);
    }

    function buy(uint256 _tokenId) external payable whenNotPaused
    {
        _bid(_tokenId, msg.value);
        operatorApprovals[address(this)][msg.sender] = true;
        approve(msg.sender, _tokenId);
        transferFrom(address(this), msg.sender, _tokenId);
    }
    function transactionInfo(uint256 _tokenId)
        external
        view
        returns (uint128 price, address seller) 
    {
        Trasaction memory trasaction = tokenIdToTransaction[_tokenId];
        price = trasaction.price;
        seller = trasaction.seller;
    }
}

contract RepoTransaction is StampBase{
    struct RepoIngots {
        address seller;
        uint256 tokenId;
        uint256 repoCount;
    }
    
    mapping (uint256 => RepoIngots) tokenIdToRepo;
    
    event RepoIngotsSuccessful(uint256 tokenId, uint256 repoCount, address seller);

    function repoIngotsTransaction(uint256 _tokenId, uint256 _repoCount) public whenNotPaused{
        require(_owns(msg.sender, _tokenId));
        transferFrom(msg.sender, address(this), _tokenId);
        RepoIngots memory repoIngotsInfo = RepoIngots(
            msg.sender,
            _tokenId,
            _repoCount
        );
        tokenIdToRepo[_tokenId] = repoIngotsInfo;
        emit RepoIngotsSuccessful(_tokenId, _repoCount, msg.sender);
    }
    
    function repoIngotsInfo(uint256 _tokenId) public view returns (uint256 tokenId, uint256 repoCount, address seller)
    {
         RepoIngots memory repoIngots = tokenIdToRepo[_tokenId];
         tokenId = repoIngots.tokenId;
         repoCount = repoIngots.repoCount;
         seller = repoIngots.seller;
    }
}

contract StampMinting is StampTransaction, RepoTransaction {
    uint256 public constant STAMP_SET_RELEASE_LIMIT = 3000;
    uint256 public constant STAMP_TYPE_RELEASE_LIMIT = 30000;
    
    uint256 public setReleasedCount;
    uint256 public stampTypeReleasedCount;

    function releaseNewStampSet(uint16 _setId, uint32[] _typeIds) external whenNotPaused onlyCOO{
        require(setReleasedCount < STAMP_SET_RELEASE_LIMIT);
        require(!isExistSet[_setId]);
        require(_setId > 0);
        require(_typeIds.length > 0);
        StampSet memory _stampSet = StampSet({
            setId:_setId,
            typeIds:_typeIds
        });
        
        for(uint8 i = 0; i < _typeIds.length; i++) {
            TypeIdToSetId[_typeIds[i]] = _setId;
        }
        
        stampSets.push(_stampSet);
        SetIdToStampSet[_setId] = _stampSet;
        isExistSet[_setId] = true;
        setReleasedCount++;
    }
    
    function releaseNewStampInfo(uint32 _setId, 
    uint32 _typeId,  uint32 _totalAmount, 
    uint16 _year, uint16 _idOfSet, bytes32 _name) external whenNotPaused onlyCOO returns (uint32){
        require(stampTypeReleasedCount < STAMP_TYPE_RELEASE_LIMIT);
        require(!isExistStampType[_typeId]);
        StampInfo memory stampInfo = StampInfo({
            setId:_setId,
            typeId:_typeId,
            totalAmount:_totalAmount,
            remainingAmount:_totalAmount,
            year:_year,
            idOfSet:_idOfSet,
            name:_name
        });
        stampInfos.push(stampInfo);
        TypeIdToStampInfo[_typeId] = stampInfo;
        isExistStampType[_typeId] = true;
        return _typeId;
    }

    function releaseNewStampToTransaction(
        uint32 _typeId,
        uint8 _appearance
        ) 
        external 
        whenNotPaused
        onlyCOO 
        returns (uint256)
    {
        StampInfo memory stampInfo = TypeIdToStampInfo[_typeId];
        uint32 remainingAmount = stampInfo.remainingAmount;
        require(remainingAmount>0);
        require(isExistStampType[_typeId]);

        uint32 totalAmount = stampInfo.totalAmount;
        uint16 year = stampInfo.year;
        uint256 price = _computeStampPrice(year, totalAmount, _appearance);
        uint256 tokenId = _createNewStamp(_typeId, price, _appearance);
        _mint(address(this), tokenId);
        
        Trasaction memory tracsaction = Trasaction(
            now,
            uint128(price),
            address(this)
        );
        _addTransaction(tokenId, tracsaction);
        
        stampInfo.remainingAmount--;
        return tokenId;
    }

    function _computeStampPrice(uint16 _year, uint256 _totalCirculation, uint8 _appearance) internal pure returns (uint256 price) {
        price= 33*(10**3)/((_year - 1959)*_totalCirculation);
        if (_appearance == 0) {
            price = 60000*price;
        }
        if (_appearance == 1) {
            price = 1016*price;
        }
        if (_appearance == 2) {
            price = 600*price;
        }
        if (_appearance == 3) {
            price = 250*price;
        }
    }
}

contract StampCollection is StampMinting {
    address public newContractAddress;

    constructor() public {
        paused = false;
        ceoAddress = msg.sender;
        cooAddress = msg.sender;
    }
    function() external payable {
    }
    function setNewAddress(address _v2Address) external onlyCEO whenPaused {
        newContractAddress = _v2Address;
        emit ContractUpgrade(_v2Address);
    }

    function getStampInfo(uint16 _typeId) public view returns(uint32 setId, 
    uint32 typeId,  uint32 totalAmount, uint32 remainingAmount, 
    uint16 year, uint16 idOfSet, bytes32 name){
        StampInfo memory _stampInfo = TypeIdToStampInfo[_typeId];
        setId = _stampInfo.setId;
        typeId = _stampInfo.typeId;
        totalAmount = _stampInfo.totalAmount;
        remainingAmount = _stampInfo.remainingAmount;
        year = _stampInfo.year;
        idOfSet = _stampInfo.idOfSet;
        name = _stampInfo.name;
    }
    
    function getStamp(uint256 _typeId)
        external
        view
        returns (
        uint32 typeId,
        uint8 appearance,
        uint256 price
    ) 
    {
        Stamp storage stamp = stamps[_typeId];
        typeId = stamp.typeId;
        appearance = stamp.appearance;
        price = stamp.price;
    }

    function getStampSetInfo(uint16 _setId) public view returns(uint16 setId, uint32[] typeIds){
        StampSet memory stampSet = SetIdToStampSet[_setId];
        setId = stampSet.setId;
        typeIds = stampSet.typeIds;
    }
    
    function existStampSet(uint16 _setId) public view returns (bool) {
        StampSet memory stampSet = SetIdToStampSet[_setId];
        return stampSet.setId != 0;
    }
    
    function existStampInfo(uint16 _typeId) public view returns (bool) {
        StampInfo memory stampInfo = TypeIdToStampInfo[_typeId];
        return stampInfo.setId != 0 && stampInfo.typeId != 0;
    }
    
    function unpause() public onlyCEO whenPaused {
        require(newContractAddress == address(0));
        super.unpause();
    }

    function withdrawBalance() external onlyCFO {
        uint256 balance = address(this).balance;
        cfoAddress.transfer(balance);
    }
}
