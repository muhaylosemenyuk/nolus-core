# Governance Proposal Handlers

Governance proposals in Cosmos SDK are used to provide users more honest and better experienced voting.
By using gov proposals many important decisions about the network's future.

There are some proposal types enabled in our network:
* Upgrade proposals
* Distribution proposals
* Parameters change proposals
* Wasm proposals (actions related to smart contracts)

Proposals can be created by any member of the network, but only users who delegated their tokens to validator(-s).

## Algorithm of creating and passing governance proposals:
1. First of all proposal should be created with mentioned type, title, description and proposed changes. E.g. for contract migration:
````shell
nolusd tx gov submit-proposal migrate-contract [contract-address] [uploaded-code-id] [message-for-migration-function] --from [your-account] --title [title] --description [decription] --fees 750000unls --gas 300000000unls
````
Message for calling smart contract's function should be in JSON format and pasted into ' '. In case when the message for your purpose is defined empty you should send only '{}'.
2. After proposal is created it is in deposit period so amount of staked should reach 10000000unls to start the voting period. E.g.:
````shell
nolusd tx gov deposit [proposal-id] [deposit-ammount] --from [your-account] --fees 500unls
````
3. If you still have not delegated your tokens to validator and want to take part in the voting you should stake, E.g.:
````shell
nolusd tx staking delegate [delegators-address] [ammount] --from [your-account] --fees 100000unls
````
The addresses of validators you can find by querying them:
````shell
nolusd query staking validators
````
When you are able to vote for option you are interested in, you can do it by using:
````shell
nolusd tx gov vote [proposal-id] [option] --from [your-account] --fees 500unls
````
After you had voted you can only wait voting period's end to get the result of proposal, you can always check it by:
````shell
nolusd query gov proposal [proposal-id]
````