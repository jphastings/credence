# Credence

A platform for publishing and analysing belief in the truth of short, factual statements on the internet. For example, a published "cred" could be interpreted like this:

> I believe the fact in the following phrase to be false:
>
> "The longest consecutive crowd wave occurred in 2002 at a Denver broncos game. The wave circled the stadium 492 times and lasted over 3 hrs." ([source](https://twitter.com/RealFakeFacts/status/409062831355486208))
>
> I cite this [proof](http://www.guinnessworldrecords.com/world-records/longest-mexican-wave-%28timed%29). Signed, me.

Which woud be represented like this (in a compressed form) and transmitted to other users.

```json
{
  "keys": [
    "https://twitter.com/RealFakeFacts/status/409062831355486208"
  ],
  "timestamp": "1452109802",
  "assertion": "IS_FALSE",
  "proof_uri": "http://www.guinnessworldrecords.com/world-records/longest-mexican-wave-%28timed%29",
  "signature": "(base64 encoded RSA signature)",
  "human_readable": {
    "statement": "The longest consecutive crowd wave occurred in 2002 at a Denver broncos game. The wave circled the stadium 492 times and lasted over 3 hrs."
  }
}
```

## How?

Creds are broadcast in a gossip-like fashion so any other user of the network can request what credence others give to specific phrases, websites, etc. Relevant creds are requested from the network and aggregated to create an indication of truth that can suppliment the user's own intuition (see [Why?](#why)).

Some sources are trustworthy, some are best ignored, and which source is which is different for individuals. If a user publishes their public key, which allows others to recognise creds coming from them and assign a trust weighting to them.

I might choose to weight creds from [Brian Cox](https://en.wikipedia.org/wiki/Brian_Cox_(physicist)) or [Sir David Attenborough](https://en.wikipedia.org/wiki/David_Attenborough) at 1000 times more relevant to me than an anonymous user. Though the algorithm is complex, this broadly equates to "it would take 1000 anonymous people telling me a statement is false to balance Sir David telling me it is true".

## Why?

The internet has allowed very fast, very broad dispersal of information and in particular between people who are effectively anonymous to each other. However there are limited mechanisms for evaluating the information _sources_ and for doing fact checking.

The provenance of information used to be very clear -- "this article was written by this journalist who writes all the articles on that matter") -- but now its more often that the author of an article isn't known, let alone that their biases and factual reliability are understood.

Credence aims to be a mechanism for bringing our own networks of belief-in-accuracy to the internet. In an isolated community an individual quickly learns whose declarations can be trusted and who is less reliable using their own experiences and the reputation already earned with others in the community. Credence facilitates this functionality for massively distributed communication, and tries to do only that with the aim of becoming a platform for more specific tools.

## Usage

Credence is a background service written in go. Once you have go installed you should clone this repo and use the Makefile to compile it.

```bash
brew install libsodium czmq
brew install --devel protobuf

cd $GOPATH/src
git clone https://github.com/jphastings/credence.git github.com/jphastings/credence
cd github.com/jphastings/credence
make bootstrap
```

A `credence` binary will now exist and running it will give you all the instructions you need.
