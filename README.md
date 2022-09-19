# Art Gobblers contest details
- $95,000 USDC main award pot
- $5,000 USDC gas optimization award pot
- Join [C4 Discord](https://discord.gg/code4rena) to register
- Submit findings [using the C4 form](https://code4rena.com/contests/2022-09-artgobblers-contest/submit)
- [Read our guidelines for more details](https://docs.code4rena.com/roles/wardens)
- Starts September 20, 2022 16:00 UTC
- Ends Septemnber 27, 2022 16:00 UTC

## Contact us

#### Frankie
- **Discord** : frankie#6284
- **Telegram** : @FrankieIsLost
- **Twitter** : @FrankieIsLost

#### Transmissions11

- **Discord** : t11s#9672
- **Telegram** : @t11s_paradigm
- **Twitter** : @transmissions11

## Key Links 

- [Art Gobblers Overview](https://www.paradigm.xyz/2022/09/artgobblers)
- [VRGDAs](https://www.paradigm.xyz/2022/08/vrgda)
- [GOO](https://www.paradigm.xyz/2022/09/goo)
- [Website](https://artgobblers.com/)
- [Project Twitter](https://twitter.com/artgobblers)



## Contest Scope

The scope of the contest includes all the contracts in the src and script directories, as well as all src contracts in lib/VRGDAs and lib/goo-issuance. It also includes some contracts from lib/solmate. Specifically, the full list of contracts is as follows: 

File|blank|comment|code
:-------|-------:|-------:|-------:
src/ArtGobblers.sol|191|370|357
lib/solmate/src/utils/FixedPointMathLib.sol|40|69|144
lib/solmate/src/tokens/ERC721.sol|58|36|137
src/utils/token/GobblersERC1155B.sol|39|40|115
lib/solmate/src/utils/SignedWadMath.sol|32|73|112
src/utils/token/GobblersERC721.sol|46|44|105
src/utils/token/PagesERC721.sol|49|34|105
script/deploy/DeployBase.s.sol|14|11|81
src/Pages.sol|47|148|75
src/utils/rand/ChainlinkV1RandProvider.sol|18|30|30
lib/VRGDAs/src/LogisticToLinearVRGDA.sol|12|31|28
lib/solmate/src/utils/MerkleProofLib.sol|8|15|24
script/deploy/DeployRinkeby.s.sol|6|11|24
src/Goo.sol|15|79|24
lib/VRGDAs/src/LogisticVRGDA.sol|11|31|23
lib/solmate/src/utils/LibString.sol|12|22|21
lib/VRGDAs/src/VRGDA.sol|10|31|19
lib/goo-issuance/src/LibGOO.sol|6|17|19
lib/solmate/src/auth/Owned.sol|12|15|17
lib/VRGDAs/src/LinearVRGDA.sol|7|21|16
src/utils/GobblerReserve.sol|7|19|16
src/utils/rand/RandProvider.sol|4|12|6
--------|--------|--------|--------
SUM:|644|1159|1498

## Attack Surfaces

### Related to VRGDA Auctions:
* Could an adversary use specific inputs or timing conditions to severely underpay for a Gobbler/Page?
* Could an innocent user be front-run in a way that our slippage protections wouldn't catch and overpay for a Gobbler/Page?
* Could auctions be messed up for all users by becoming bricked and/or return a price far too low or high?

### Related to Page Switchover:
* Can the VRGDA switchover be delayed or forced to occur early by an adversary?
* Can an adversary use post-switch or pre-switch pricing early?
* Are prices messed up during or around the switch?

### Related to Reserves/Community:
* Can an adversary stop gobblers from being minted to a reserve?
* Can reserves stop functioning and not allow minting/withdrawing gobblers to them?
* Can too many gobblers be minted to the reserves?

### Related to GOO emissions:
* Can a user manipulate their gobbler holdings in such a way that they receive a disproportionate amount of goo emissions?
* Can goo emissions be bricked?
* Can an innocent user receive too few goo emissions relative to their gobbler holdings?
* Can goo balances overflow?
* Can an innocent user receive too few goo emissions due to a rounding error?
* Can an adversary use rounding error to gain extra goo, especially when depositing and withdrawing?
* Can transferring gobblers allow the recipient to receive a disproportionate emission multiple or result in the sender losing a disproportionate emission multiple?

### Related to Legendary Auctions:
* Can legendary auctions be started too early?
* Can an adversary underpay for a legendary gobbler?
* Can an innocent user be forced to overpay for a legendary gobbler?
* Can legendary gobbler auctions be bricked?
* Can legendary auctions return messed-up prices?
* Can more than 10 legendary auctions occur?
* Can a user mint a legendary gobbler with a higher than expected emission multiple?

### Related to Gobbler reveals:
* Can an adversary take advantage of the way seeds are requested from Chainlink or use specific inputs when revealing their Gobblers in a way that would allow them to control the randomness used to generate Gobble metadata and attributes?
* Can reveals be halted by an adversary or messed up in a way that the attributes assigned are not random?

### Related to GOO:
* Is all access control sound? Can an adversary mint goo for themselves?
* Can an innocent user accidentally burn goo without recourse?
* Can an adversary burn goo from another user's wallet?

## Development & Testing

You will need a copy of [Foundry](https://github.com/foundry-rs/foundry) installed before proceeding. See the [installation guide](https://github.com/foundry-rs/foundry#installation) for details.

To build the contracts:

```sh
git clone https://github.com/artgobblers/art-gobblers.git
cd art-gobblers
forge install
```

### Run Tests

In order to run unit tests, run: 

```sh
forge test
```

For longer fuzz campaigns, run: 

```sh
FOUNDRY_PROFILE="intense" forge test
```

For differential fuzzing against a python implementation, see [here](./analysis/README.md).



# Project Overview

The following is taken from our [overview post](https://www.paradigm.xyz/2022/09/artgobblers), and added here for convenience. 

## Overview

Art Gobblers is a decentralized art factory owned by aliens.
As artists make cool art, Gobblers gains cultural relevance, making collectors want the art more, incentivizing artists to make cooler art.
It's also an on-chain game.

## Systems

Art Gobblers are called Art Gobblers because they gobble art. In particular, they eat art that artists draw using our [draw tool](https://artgobblers.com/draw) and turn into 1/1 NFTs using in-game resources. All the artworks a Gobbler eats belong to it on-chain and are displayed in its belly gallery forever.
![](https://www.paradigm.xyz/static/art_gobblers_eating_art.png)
Art Gobblers produce [Goo](https://www.paradigm.xyz/2022/09/goo) tokens, which are used to produce the blank pages needed to make art. Gobblers love the smell of their own goo, so the more Goo they have in their Goo tanks, the faster they squirt out more new Goo.

The supply of Goo grows faster every day, starting at hundreds and eventually reaching billions and beyond. So, the game can't be balanced by giving items fixed prices in Goo. Instead, a mechanism called [VRGDA](https://www.paradigm.xyz/2022/08/vrgda) automatically adjusts prices over time to target a desired issuance schedule, adjusting prices up when sales are ahead of schedule, and down when sales are behind schedule.

The protocol targets an initial Blank Page VRGDA issuance rate of 69 per day to foster experimentation, but eventually slows down to a constant 10 per day to ensure a high bar and focused attention from the community.

The initial (free!) Gobbler mint consists of 2,000 fully animated Gobblers. Over the next 10 years, players will spend Goo to mint 8,000 more Gobblers using VRGDA. Issuance is relatively fast at first to bootstrap growth, but slows and eventually stops to preserve exclusivity.

These systems can lead to some interesting gameplay, as players must decide when to trade Goo for Gobblers, changing the VRGDA price for everyone else. Legendary Gobblers, Goo-boosting 1/1s that can only be obtained by burning huge numbers of regular Gobblers, encourage large-scale collaboration.
![](https://www.paradigm.xyz/static/art_gobblers_mechanism.png)
## Intentions

The mechanisms making up the Art Gobblers experiment will be final at the time of mint, and no more will be built. They are intended to power a self-sustaining ecosystem capable of creating and collecting the coolest art in the universe without the need for human intervention.

We hope:

As artists make cool art, the cultural relevance of Art Gobblers will grow.

As cultural relevance grows, Gobbler art will be in higher demand from collectors.

As artists see higher collector demand, they will produce cooler art.

Gobble on.
![](https://www.paradigm.xyz/static/art_gobblers_flywheel.png)
# Design

## Art

Art Gobblers exists to facilitate the creation and collection of the coolest art in the universe.

Artists can create hand-drawn art using our [draw tool](https://artgobblers.com/draw/dad78c97-6ca9-4636-83cb-45e662b1d923), which works on desktop, iPad, or Cintiq.

They’ve been making some pretty cool stuff already.
![](https://www.paradigm.xyz/static/art_gobblers_artworks.png)
The draw tool records the entire process of creation, and when you visit a work of art in the Art Gobblers webapp, you can play back the drawing process stroke by stroke.

## Pages

Once Art Gobblers is live, artists will be able to mint their own drawings as 1/1 NFTs using in-game resources, a process we call *glamination*. These NFTs, which belong to the Art Gobblers Pages collection, are ERC721 NFTs that belong to the artist. Like any other NFT, they can be transferred or sold by the artist at will.

Producing these art Pages is the main point of the Art Gobblers ecosystem. Because the number of Pages that can be created in a given period of time is limited by a [VRGDA](https://www.paradigm.xyz/2022/08/vrgda) mechanism, any art glaminated onto a Page will benefit from the cultural relevance of Art Gobblers and the attention of the Art Gobblers community, which has been seeded with some of the most accomplished collectors in the 1/1 space.

In order to glaminate your drawing onto a Page, you need a *Blank Page* NFT, which can be created using *Goo.*

## Goo

Goo, as described in [the GOO paper](https://www.paradigm.xyz/2022/09/goo), is an ERC20 Ethereum token emitted by Art Gobblers NFTs. Because Goo is needed to make new Blank Pages, Art Gobblers ultimately determine what art can be created in the ecosystem, and in that way serve as co-curators of a decentralized art gallery.

Goo can also be used to create new Gobblers, giving players an interesting set of strategic decisions that we will discuss further in the Details section.

## Gobblers

Art Gobblers themselves are fully animated ERC1155 NFTs.

As well as emitting Goo, they can, of course, gobble art.

In particular, let’s say you own both an Art Gobbler and some cool art on a Page. If you feed that Page to your Gobbler, ownership of the Page transfers on-chain to the Art Gobblers contract, which has a mapping indicating which particular Gobbler that piece of art belongs to.

That Page then becomes a permanent part of that Gobbler’s belly gallery, viewable through the Art Gobblers app. If you transfer your Gobbler, all of its gobbled works go with it.

Collectors can curate their Gobblers however they choose. We can imagine a Gobbler containing only works by Justin, a Gobbler containing only pictures of dogs, a Gobbler full of autographs, a Gobbler containing a collaborative manga, and countless other possibilities.

## Collecting Without a Gobbler

Given their limited supply and their in-game uses, Art Gobblers may be hard to come by. Collectors who can’t get their hands on a Gobbler will still be able to contribute to the ecosystem by collecting pages with art on them.

Because Blank Pages are rare, by choosing which artists to support and even directly commissioning artists to draw on Blank Pages, even collectors without Gobblers will be able to make a serious contribution to the curation of the decentralized Gallery that is Art Gobblers.

## Legendary Gobblers

The ten *Legendary Gobblers* are extravagantly rare one of one Gobblers that represent the supreme rulers of Gobbler civilization. They will arrive at specific predestined times over the next ten years, providing punctuated drama and a sense of structure over time to the game.

In order to obtain a Legendary Gobbler, you must sacrifice a huge number of normal Gobblers. The first Legendary Gobbler will start at a price of 69 Gobblers, and will become cheaper using a standard Dutch Auction mechanism until it is purchased. Each successive Legendary Gobbler will start at a price in ordinary Gobblers equal to twice what the last Legendary Gobbler sold for.

A new Legendary Gobbler appears each time an additional 10% of the total supply of Gobblers is issued via VRGDA, and each Legendary Auction is scheduled to end by the time the next Legendary Gobbler appears (or when all Gobblers are issued, in the case of the final Legendary).

To compensate for the many Gobbler lives lost, Legendary Gobblers produce Goo at twice the rate of the combined Gobblers sacrificed to summon them.

## Roadmap

There isn’t one.

Art Gobblers will be launched as a finished product, designed to bootstrap a self-sustaining ecosystem.

Neither Justin, nor Paradigm, nor the Art Gobblers team plan to build anything net new after the upcoming free mint. Art Gobblers is not stage one of the Gobblers metaverse. It is a whole and complete piece of alien technology that we will be unleashing upon an unsuspecting populace.

What it does next, and how the inhabitants of Earth choose to aid it in its mission, is anyone’s guess.

# Details

## Goo

From [the GOO Paper](https://www.paradigm.xyz/2022/09/goo):

Art Gobblers, from our upcoming NFT project, are NFTs that produce an Ethereum token called Goo. The more Goo a Gobbler has, the faster it generates more Goo. This means the total Goo supply will increase faster and faster every day, going from thousands to millions and beyond.


![](https://www.paradigm.xyz/static/goo_supply.png)

Hoarding Goo tokens without owning any Gobbler NFTs is a very bad strategy, as everyone else will be generating goo and your share of the total Goo supply will rapidly dwindle to nothing. On the other hand, if you own many Gobblers but little Goo, your Goo production will lag compared to other players.

But, let's say you maintain ownership of Gobbler NFTs with a combined Goo production capacity of 1% of the total and never remove your Goo. No matter how much Goo you start with, you will eventually end up with at least 1% of the total Goo supply! This ensures Goo stays in NFT holder control over the long term.

![](https://www.paradigm.xyz/static/goo_supply_share_chart.png)

Mathematically speaking, instantaneous Goo issuance is equal to $\sqrt{\texttt{mult}\cdot\texttt{goo in tank}}$, where a Gobbler’s $\texttt{mult}$, or multiplier, is its base speed of squirting out Goo. We auto-compound Goo issuance over time using a differential equation, and also automatically balance Goo between gobblers if you have more than one.
 
In this system, due to some extremely lucky math, it turns out that having many Gobblers with multipliers that add to $\texttt{total mult}$ is the same as having a single Gobbler with a multiplier of $\texttt{total mult}$, which means the game stays fair as some players obtain more Gobblers.

![](https://www.paradigm.xyz/static/gobbler_sum_graphic.png) 

## VRGDA

Because the Goo supply is increasing faster and faster every day, having fixed Goo prices for Gobblers or Pages would make no sense. A fair price for a Gobbler might be 69 immediately after mint, but 69,000 when the Goo supply has grown by several orders of magnitude.

Instead, prices are set dynamically to make issuance of Gobblers and Pages track pre-determined schedules using a mechanism called VRGDA.

From [the VRGDA paper](https://www.paradigm.xyz/2022/08/vrgda):

Variable Rate GDAs (VRGDAs), designed for [Art Gobblers](http://www.artgobblers.com/) and used in [0xMonaco](https://twitter.com/transmissions11/status/1561100140160593920), let you sell tokens close to a custom schedule over time by raising prices when sales are ahead of schedule and lowering prices when sales are behind schedule — a generalization of the [GDA](https://www.paradigm.xyz/2022/04/gda) mechanism.

[…]

Imagine a simple schedule where we want to sell 10 NFTs per day. We set a starting price of 1 token for the first NFT.

Suppose it is currently day 5, so we should have sold 50 NFTs. However, demand has been high, and we have sold 70. We weren’t supposed to sell 70 NFTs until day 7, so we are two days ahead of schedule.

As a result, we want to charge a higher price going forward. We use an exponential curve to determine how much higher. This can vary based on parameters, but in this case, let’s say we use $2^\text{days ahead of schedule}$, so that we increase our price by a factor $2^2=4$, so since our initial price was 1 token, the new price will be 4 tokens, making it harder to buy more NFTs.

![](https://www.paradigm.xyz/static/vrgda_graphic.png)

Ten days later, on day 15, we should have sold 150 NFTs, but users have only bought 120, the amount they should have bought by day 12, meaning we are three days *behind* schedule. We adjust the price to $2^{-3}=0.125$, making it easier for users to buy more NFTs.

Because your rate of Goo production depends on both how much Goo you have and how many Gobblers you have, spending Goo to get Gobblers presents a tradeoff. Furthermore, every time you spend your Goo to get Gobblers, you increase the Gobbler VRGDA price for every other player.

This presents a complex strategy space that we saw play out on our 30x speed test net run (to be released soon). Once you add in the possibilities of DAOs and other forms of collaboration, especially as they interact with Legendary Gobblers, the scope for action becomes quite wide.

## Issuance Schedules

### Gobblers

![](/static/gobbler_schedule.png)

Art Gobblers themselves begin at a supply of 2,000, issued as a free mint. Of these, 300 go to the core team and the rest go to the community, a mixture of collaborators, artists, collectors, builders, contest winners, and others.

Over the next ten years, an additional 8,000 Gobblers will be minted using a [logistic schedule](https://www.paradigm.xyz/2022/08/vrgda#logistic-issuance-schedule). At first, the protocol will issue roughly 200 Gobblers per month, but this will slow down over time.

This pattern allows the community to grow quickly and bootstrap at first while avoiding over-inflation in the long term.

Similarly to the [Nouns](https://nouns.wtf/) model, one in ten newly minted Gobblers will accrue to the team. An additional one in ten will go to a vault to be distributed to the community.

### Blank Pages

![](https://www.paradigm.xyz/static/gobbler_schedule.png)

No Blank Pages will exist at mint, but they can be created shortly after using Goo.

The schedule calls for an initial creation rate of 69 pages per day to allow the community to experiment with art styles, memes, and more. Over the course of about 8 months, this issuance rate will slow down along a logistic curve until it hits a constant rate of 10 pages per day, which the protocol maintains forever.

No Pages accrue directly to the team, but one in ten newly created pages do go to a vault to be distributed to artists in the community.

## Parameter Spreadsheet

A spreadsheet containing parameters and issuance schedules can be found [here](https://docs.google.com/spreadsheets/d/1i8hYuWyAymjbwx54fA1HcEMEwfEEyluX4tPssOXoyf4/edit?usp=sharing).
