# NiftyRewards Server

Backend server created in Golang to provide an API interface for interaction with Database.

## Database Design

TODO: Insert Image here

### Entities

Merchants(**merchant_id**, merchant_name)

NFTs (**collection_address**, collection_name)

Users (**address_w3a**, address_b)

### Relationships

Campaign (**campaign_id**, _merchant_id_, _collection_address_, start_time, duration, totalSupply)

Which merchant is collaborating with which NFT

Rewards (**rewards_id**, _merchant_id_, _collection_address_, _ownership_id_, description, quantity, quantity_used)

## Endpoints

### User

[**POST**] /users/bind/{address_w3a}/{address_to_bind}

Maps the web3Auth wallet address to a wallet that the user owns (one that is used to collect NFTS for example)

AUTH: Caller must be the owner of the account associated to address_w3a

Replaces current binded wallet with address_to_bind.

[**GET**] /user/nfts/{address_w3a}

Gets all the NFTS

Returns a list of user’s NFTS in the wallet.

Data is populated from request to Tatum

[https://apidoc.tatum.io/#operation/NftGetTokensByAddressErc721]

[**POST**] /user/rewards/redeem
Redeems rewards based on rewards_id

Only owner of rewards_id can redeem

Reduce quantity by 1
Increase quantity_used by 1

### Merchant

[**POST**] /merchant/startcampaign

Request to start a campaign with a collection.

```txt
body:
collection_address: String
```

```txt
- Collection owner gets account address of the merchant
- Collection owner logs into the web app
- Collection owner goes to “approve merchant” page
- Collection owner inputs merchant address collected in step 1 & collection address
- Collection owner signs an typed sign with the merchant address and collection address
- Backend verifies if the collection owner is indeed the ownerOf collection address
- If ownerOf() Collection address is equals to the recovered address of the signature, the collection will be binded.
```

[**GET**] /merchant/campaigns
Retrieves all campaigns that a merchant has

```txt
returns:
campaigns: Campaign[]
```

[**POST**] /merchant/addrewards

Add a reward for the collection address.

Collection must be approved by collection owner.

```txt
body:
collection_address: String
description: String
```

### Collection Owner

[**POST**] /collectionowner/add

Adds collection to merchant. Only owner of collection can do this
The owner of the collection will be able to approve the merchant to use it to apply discounts. Approval can only happen after merchant has requested for a collaboration

```txt
body:
merchant_address: String
```
